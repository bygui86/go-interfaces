// +build integration

package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	productName  = "sample"
	productPrice = 42.42

	productNewName  = "new-sample"
	productNewPrice = 9.90
)

func TestGetProducts_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &Product{Name: productName, Price: productPrice}
	insertErr := CreateProduct(db, product, ctx)
	require.NoError(t, insertErr)

	products, err := GetProducts(db, 0, 10, ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 1)
	assert.GreaterOrEqual(t, products[0].ID, 0)
	assert.Equal(t, productName, products[0].Name)
	assert.Equal(t, productPrice, products[0].Price)

	DeleteProducts(db, ctx)
}

func TestGetProduct_Unit_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	sourceProd := &Product{Name: productName, Price: productPrice}
	insertErr := CreateProduct(db, sourceProd, ctx)
	require.NoError(t, insertErr)

	targetProd := &Product{ID: sourceProd.ID}
	err := GetProduct(db, targetProd, ctx)
	assert.NoError(t, err)
	assert.Equal(t, sourceProd.ID, targetProd.ID)
	assert.Equal(t, sourceProd.Name, targetProd.Name)
	assert.Equal(t, sourceProd.Price, targetProd.Price)

	DeleteProducts(db, ctx)
}

func TestCreateProduct_Integr_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &Product{Name: productName, Price: productPrice}
	err := CreateProduct(db, product, ctx)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, product.ID, 0)
	assert.Equal(t, productName, product.Name)
	assert.Equal(t, productPrice, product.Price)

	DeleteProducts(db, ctx)
}

func TestUpdateProduct_Unit_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	insert := &Product{Name: productName, Price: productPrice}
	insertErr := CreateProduct(db, insert, ctx)
	require.NoError(t, insertErr)

	update := &Product{ID: insert.ID, Name: productNewName, Price: productNewPrice}
	err := UpdateProduct(db, update, ctx)
	assert.NoError(t, err)
	assert.Equal(t, insert.ID, update.ID)
	assert.Equal(t, productNewName, update.Name)
	assert.Equal(t, productNewPrice, update.Price)
	assert.NotEqual(t, insert.Name, update.Name)
	assert.NotEqual(t, insert.Price, update.Price)

	DeleteProducts(db, ctx)
}

func TestDeleteProduct_Unit_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &Product{Name: productName, Price: productPrice}
	insertErr := CreateProduct(db, product, ctx)
	require.NoError(t, insertErr)

	getErr := GetProduct(db, product, ctx)
	require.NoError(t, getErr)

	err := DeleteProduct(db, product.ID, ctx)
	assert.NoError(t, err)

	DeleteProducts(db, ctx)
}

func TestDeleteProducts_Unit_Success(t *testing.T) {
	ctx := context.Background()

	db := initConnAndTable(t)

	product := &Product{Name: "one", Price: 1.10}
	insertErr := CreateProduct(db, product, ctx)
	require.NoError(t, insertErr)

	product2 := &Product{Name: "two", Price: 2.20}
	insert2Err := CreateProduct(db, product2, ctx)
	require.NoError(t, insert2Err)

	product3 := &Product{Name: "three", Price: 3.30}
	insert3Err := CreateProduct(db, product3, ctx)
	require.NoError(t, insert3Err)

	productsBefore, getErrBefore := GetProducts(db, 0, 10, ctx)
	require.NoError(t, getErrBefore)
	require.Len(t, productsBefore, 3)

	err := DeleteProducts(db, ctx)
	assert.NoError(t, err)

	productsAfter, getErrAfter := GetProducts(db, 0, 10, ctx)
	assert.NoError(t, getErrAfter)
	assert.Len(t, productsAfter, 0)

	DeleteProducts(db, ctx)
}
