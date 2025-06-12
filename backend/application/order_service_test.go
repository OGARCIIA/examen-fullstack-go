package application

import (
	"testing"

	"github.com/OGARCIIA/examen-backend/domain"
	"github.com/OGARCIIA/examen-backend/infrastructure"
)

func setupTestDBOrders(t *testing.T) {
	infrastructure.InitDatabase()

	// AutoMigrar las tablas (muy importante en GitHub Actions)
	err := infrastructure.DB.AutoMigrate(&domain.Product{}, &domain.Order{})
	if err != nil {
		t.Fatalf("Error al hacer AutoMigrate: %v", err)
	}

	// Limpiar las tablas
	err = infrastructure.DB.Exec("DELETE FROM orders").Error
	if err != nil {
		t.Fatalf("Error al limpiar tabla orders: %v", err)
	}

	err = infrastructure.DB.Exec("DELETE FROM products").Error
	if err != nil {
		t.Fatalf("Error al limpiar tabla products: %v", err)
	}
}

func TestCreateOrder(t *testing.T) {
	setupTestDBOrders(t)

	product := domain.Product{
		Name:        "Producto para orden",
		Description: "Desc prueba",
		Price:       100.0,
		Stock:       10,
	}

	err := CreateProduct(&product)
	if err != nil {
		t.Fatalf("Error al crear el producto: %v", err)
	}

	order, err := CreateOrder(product.ID, 3)
	if err != nil {
		t.Fatalf("Error al crear la orden: %v", err)
	}

	if order.Quantity != 3 {
		t.Errorf("Se esperaba Quantity=3, se obtuvo %d", order.Quantity)
	}

	if order.Total != 300.0 {
		t.Errorf("Se esperaba Total=300.0, se obtuvo %.2f", order.Total)
	}

	var updatedProduct domain.Product
	res := infrastructure.DB.First(&updatedProduct, product.ID)
	if res.Error != nil {
		t.Fatalf("Error al obtener el producto actualizado: %v", res.Error)
	}

	if updatedProduct.Stock != 7 {
		t.Errorf("Se esperaba Stock=7 después de la orden, se obtuvo %d", updatedProduct.Stock)
	}
}
