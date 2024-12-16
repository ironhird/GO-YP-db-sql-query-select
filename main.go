package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	// напишите код здесь

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return sales, err
	}
	defer db.Close()

	rows, err := db.Query("select product, volume, date from sales where client = :client_id", sql.Named("client_id", client))
	if err != nil {
		return sales, err
	}
	defer rows.Close()

	for rows.Next() {
		var sale Sale

		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			return sales, err
		}

		sales = append(sales, sale)
	}

	if err := rows.Err(); err != nil {
		return sales, err
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
