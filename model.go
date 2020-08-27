package main

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *pgx.Conn) error {
	return db.QueryRow(context.Background(), "SELECT name, price FROM products WHERE id=$1", p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), "UPDATE products SET name=$1, price=$2 WHERE id=$3", p.Name, p.Price, p.ID)
	return err
}

func (p *product) deleteProduct(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), "DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createProduct(db *pgx.Conn) error {
	return db.QueryRow(context.Background(), "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id", p.Name, p.Price).Scan(&p.ID)
}

func getProducts(db *pgx.Conn, start, count int) ([]product, error) {
	rows, err := db.Query(context.Background(), "SELECT id, name,  price FROM products LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Slice of products
	products := []product{}

	// Loop over the rows
	for rows.Next() {
		var p product
		// Scan will map "p product" to the values of the struct
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
