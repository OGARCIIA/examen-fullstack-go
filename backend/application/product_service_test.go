package application

import (
	"testing"

	"github.com/OGARCIIA/examen-backend/domain"
	"github.com/OGARCIIA/examen-backend/infrastructure"
)

func setupTestDB(t *testing.T) {
	infrastructure.InitDatabase()

	// AutoMigrar las tablas (necesario en los tests, ya que el pipeline es un MySQL nuevo)
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

func TestCreateProduct(t *testing.T) {
	setupTestDB(t)

	product := domain.Product{
		Name:        "Producto Test",
		Description: "Descripción de prueba",
		Price:       100.0,
		Stock:       10,
	}

	err := CreateProduct(&product)
	if err != nil {
		t.Fatalf("Error al crear el producto: %v", err)
	}

	if product.ID == 0 {
		t.Fatalf("El producto no tiene ID asignado después de la creación")
	}
}

func TestGetAllProducts(t *testing.T) {
	setupTestDB(t)

	product := domain.Product{
		Name:        "Producto Test 2",
		Description: "Descripción 2",
		Price:       150.0,
		Stock:       5,
	}

	err := CreateProduct(&product)
	if err != nil {
		t.Fatalf("Error al crear el producto: %v", err)
	}

	products, err := GetAllProducts()
	if err != nil {
		t.Fatalf("Error al obtener los productos: %v", err)
	}

	if len(products) == 0 {
		t.Fatalf("No se encontraron productos, se esperaba al menos 1")
	}
}

func TestUpdateProduct(t *testing.T) {
	setupTestDB(t)

	product := domain.Product{
		Name:        "Producto para actualizar",
		Description: "Original",
		Price:       50.0,
		Stock:       20,
	}

	err := CreateProduct(&product)
	if err != nil {
		t.Fatalf("Error al crear el producto: %v", err)
	}

	updatedProduct := domain.Product{
		Name:        "Producto actualizado",
		Description: "Descripción actualizada",
		Price:       75.0,
		Stock:       15,
	}

	err = UpdateProduct(product.ID, &updatedProduct)
	if err != nil {
		t.Fatalf("Error al actualizar el producto: %v", err)
	}

	var result domain.Product
	infrastructure.DB.First(&result, product.ID)

	if result.Name != "Producto actualizado" {
		t.Errorf("Se esperaba Nombre 'Producto actualizado', se obtuvo '%s'", result.Name)
	}

	if result.Stock != 15 {
		t.Errorf("Se esperaba Stock 15, se obtuvo %d", result.Stock)
	}
}

func TestDeleteProduct(t *testing.T) {
	setupTestDB(t)

	product := domain.Product{
		Name:        "Producto para eliminar",
		Description: "Para test delete",
		Price:       200.0,
		Stock:       3,
	}

	err := CreateProduct(&product)
	if err != nil {
		t.Fatalf("Error al crear el producto: %v", err)
	}

	err = DeleteProduct(product.ID)
	if err != nil {
		t.Fatalf("Error al eliminar el producto: %v", err)
	}

	var result domain.Product
	res := infrastructure.DB.First(&result, product.ID)
	if res.Error == nil {
		t.Fatalf("El producto no fue eliminado, aún existe con ID %d", product.ID)
	}
}
