package main

import (
	"context"
	"database/sql"

	"github.com/danielzinhors/go-migrate/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	queries := db.New(dbConn)
	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }
	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "fec2f257-a066-4dca-be5e-3ae23857a821",
	// 	Name:        "Backend Update",
	// 	Description: sql.NullString{String: "Backend description update", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	err = queries.DeleteCategory(ctx, "fec2f257-a066-4dca-be5e-3ae23857a821")
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
