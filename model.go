package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// added features
func getProductNames(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var names []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	return names, nil
}

func (cheapest *product) getCheapestProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price, id FROM products WHERE price = (SELECT MIN(price) FROM products)").
		Scan(&cheapest.Name, &cheapest.Price, &cheapest.ID)
}

func (mostExpensive *product) getMostExpensiveProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price, id FROM products WHERE price = (SELECT MAX(price) FROM products)").
		Scan(&mostExpensive.Name, &mostExpensive.Price, &mostExpensive.ID)
}

func (p *product) searchProduct(db *sql.DB) ([]product, error) {
	rows, err := db.Query("SELECT * FROM products WHERE name iLIKE '%" + p.Name + "%'")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []product

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
