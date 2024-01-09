package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/danielzinhors/go-migrate/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDb(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	// o nil é o isolation level
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rolback: %v, original errpe %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCatetgory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCatetgory.ID,
			Name:        argsCatetgory.Name,
			Description: argsCatetgory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCatetgory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	//queries := db.New(dbConn)
	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "GO",
		Description: sql.NullString{String: "Go Course", Valid: true},
		Price:       10.95,
	}
	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend Course", Valid: true},
	}
	courseDb := NewCourseDb(dbConn)
	err = courseDb.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	if err != nil {
		panic(err)
	}

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

	// err = queries.DeleteCategory(ctx, "fec2f257-a066-4dca-be5e-3ae23857a821")
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
}
