package usecase

import (
	"errors"
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/assert"
)

// Mock do OrderRepositoryInterface
type OrderRepositoryMock struct {
	orders []entity.Order
	err    error
}

func (m *OrderRepositoryMock) Save(order *entity.Order) error {
	return nil
}

func (m *OrderRepositoryMock) GetAllOrders() ([]entity.Order, error) {
	return m.orders, m.err
}

func TestListOrderUseCase_Execute_ReturnsOrders(t *testing.T) {
	orders := []entity.Order{
		{ID: "1", Price: 10.0, Tax: 2.0, FinalPrice: 12.0},
		{ID: "2", Price: 20.0, Tax: 3.0, FinalPrice: 23.0},
	}
	repo := &OrderRepositoryMock{orders: orders}
	usecase := NewListOrderUseCase(repo)

	result, err := usecase.Execute()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, 10.0, result[0].Price)
	assert.Equal(t, 2.0, result[0].Tax)
	assert.Equal(t, 12.0, result[0].FinalPrice)
	assert.Equal(t, "2", result[1].ID)
}

func TestListOrderUseCase_Execute_ReturnsError(t *testing.T) {
	repo := &OrderRepositoryMock{err: errors.New("db error")}
	usecase := NewListOrderUseCase(repo)

	result, err := usecase.Execute()

	assert.Error(t, err)
	assert.Nil(t, result)
}
