package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/20-CleanArch/configs"
	"github.com/devfullcycle/20-CleanArch/internal/event/handler"
	"github.com/devfullcycle/20-CleanArch/internal/infra/graph"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/service"
	"github.com/devfullcycle/20-CleanArch/internal/infra/web/webserver"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println("=== CONFIGS ===")
	fmt.Println("DB_DRIVER:", configs.DBDriver)
	fmt.Println("DB_HOST:", configs.DBHost)
	fmt.Println("DB_PORT:", configs.DBPort)
	fmt.Println("DB_USER:", configs.DBUser)
	fmt.Println("DB_NAME:", configs.DBName)
	fmt.Println("WEB_SERVER_PORT:", configs.WebServerPort)
	fmt.Println("GRPC_SERVER_PORT:", configs.GRPCServerPort)
	fmt.Println("GRAPHQL_SERVER_PORT:", configs.GraphQLServerPort)
	fmt.Println("RABBITMQ_URL:", configs.RabbitMQURL) // ✅ AQUI!
	fmt.Println("================")

	db := connectDB(configs)
	defer db.Close()

	if err := migrate(db); err != nil {
		panic(fmt.Errorf("erro ao migrar tabela orders: %w", err))
	}

	rabbitMQChannel := connectRabbitMQ(configs)
	defer rabbitMQChannel.Close()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db)
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webserver.AddHandler("/order", webOrderHandler.Create)
	webserver.AddHandler("/orders", webOrderHandler.List)

	var wg sync.WaitGroup

	// REST Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Starting web server on port", configs.WebServerPort)
		if err := webserver.Start(); err != nil {
			fmt.Printf("Web server error: %v\n", err)
		}
	}()

	// gRPC Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer := grpc.NewServer()
		createOrderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
		pb.RegisterOrderServiceServer(grpcServer, createOrderService)
		reflection.Register(grpcServer)

		fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
		if err != nil {
			fmt.Printf("gRPC listen error: %v\n", err)
			return
		}
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Printf("gRPC server error: %v\n", err)
		}
	}()

	// GraphQL Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{
				CreateOrderUseCase: *createOrderUseCase,
				ListOrderUseCase:   *listOrderUseCase,
			},
		}))

		mux := http.NewServeMux()
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
		mux.Handle("/query", srv)

		fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
		if err := http.ListenAndServe(":"+configs.GraphQLServerPort, mux); err != nil {
			fmt.Printf("GraphQL server error: %v\n", err)
		}
	}()
	wg.Wait()
}

func connectDB(cfg *configs.Conf) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open(cfg.DBDriver, dsn)
		if err == nil {
			if err := db.Ping(); err == nil {
				fmt.Println("✅ Connected to MySQL")
				return db
			}
		}
		fmt.Printf("⏳ Tentando conectar ao MySQL... (tentativa %d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	panic(fmt.Errorf("erro ao conectar ao banco após 10 tentativas: %w", err))
}

func connectRabbitMQ(cfg *configs.Conf) *amqp.Channel {
	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial(cfg.RabbitMQURL)
		if err == nil {
			ch, err := conn.Channel()
			if err == nil {
				fmt.Println("✅ Connected to RabbitMQ")
				return ch
			}
		}
		fmt.Printf("⏳ Tentando conectar ao RabbitMQ... (tentativa %d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	panic(fmt.Errorf("erro ao conectar ao RabbitMQ após 10 tentativas: %w", err))
}

func migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id VARCHAR(255) PRIMARY KEY,
		price FLOAT NOT NULL,
		tax FLOAT NOT NULL,
		final_price FLOAT NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}
