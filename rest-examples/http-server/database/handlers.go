package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

func (db *DefaultInMemoryDb) GetProducts(start, count int, ctx context.Context) ([]*Product, error) {
	span := opentracing.StartSpan(
		"get-products-db",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()

	query := fmt.Sprintf("all-start[%d]-count[%d]", start, count)
	span.SetTag("query", query)
	span.SetTag("count", count)
	span.SetTag("start", start)
	span.LogKV(
		"query", query,
		"count", count,
		"start", start,
	)

	products := make([]*Product, 0)
	for _, prod := range db.products {
		products = append(products, prod)
	}

	span.SetTag("products-found", len(products))
	span.LogKV("products-found", len(products))

	return products, nil
}

func (db *DefaultInMemoryDb) GetProduct(product *Product, ctx context.Context) *Product {
	span := opentracing.StartSpan(
		"get-product-db",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()

	span.SetTag("product-id", product.ID)
	span.LogKV("product-id", product.ID)

	return db.products[product.ID]
}

func (db *DefaultInMemoryDb) CreateProduct(product *Product, ctx context.Context) error {
	span := opentracing.StartSpan(
		"create-product-db",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()

	span.SetTag("product", product.String())
	span.LogKV("product", product.String())

	newUuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	product.ID = newUuid.String()
	db.products[newUuid.String()] = product
	return nil
}

func (db *DefaultInMemoryDb) UpdateProduct(product *Product, ctx context.Context) error {
	span := opentracing.StartSpan(
		"update-product-db",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()

	span.SetTag("product", product.String())
	span.LogKV("product", product.String())

	if product != nil {
		if product.ID != "" {
			db.products[product.ID] = product
		} else {
			return errors.New("updated product ID cannot be empty")
		}
	} else {
		return errors.New("updated product cannot be empty")
	}
	return nil
}

func (db *DefaultInMemoryDb) DeleteProduct(productId string, ctx context.Context) error {
	span := opentracing.StartSpan(
		"delete-product-db",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()

	span.SetTag("product-id", productId)
	span.LogKV("product-id", productId)

	if productId != "" {
		db.products[productId] = nil
	} else {
		return errors.New("product ID cannot be empty")
	}
	return nil
}
