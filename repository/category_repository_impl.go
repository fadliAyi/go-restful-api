package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restful-api/helper"
	"go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "INSERT INTO category(name) values (?)"
	result, err := tx.ExecContext(ctx, Sql, category.Name)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	category.Id = id
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "UPDATE category SET name = ? where id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Name, category.Id)

	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	Sql := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Id)

	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (domain.Category, error) {
	Sql := "SELECT id, name FROM category where id = ?"
	rows, err := tx.QueryContext(ctx, Sql, categoryId)
	helper.PanicIfError(err)
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category {
	Sql := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, Sql)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories

}
