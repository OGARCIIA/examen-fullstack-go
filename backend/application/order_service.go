package application

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/OGARCIIA/examen-backend/domain"
	"github.com/OGARCIIA/examen-backend/infrastructure"
)

func CreateOrder(productID uint, quantity int) (*domain.Order, error) {
	var product domain.Product

	err := infrastructure.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses().First(&product, productID).Error; err != nil {
			return err
		}

		// Verificamos si hay suficiente stock
		if product.Stock < quantity {
			return errors.New("No hay suficiente stock.")
		}

		// Descontamos el stock
		product.Stock -= quantity
		if err := tx.Save(&product).Error; err != nil {
			return err
		}

		// Creamos la orden
		order := domain.Order{
			ProductID: productID,
			Quantity:  quantity,
			Total:     float64(quantity) * product.Price,
			Date:      time.Now(),
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Volvemos a obtener la orden (para devolverla completa)
	var createdOrder domain.Order
	if err := infrastructure.DB.Last(&createdOrder).Error; err != nil {
		return nil, err
	}

	return &createdOrder, nil
}
