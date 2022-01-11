package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/repository/repository"
	"context"
	"database/sql"
	"errors"
)

type FavouriteRepositoryImpl struct {

}

func NewFavouriteRepositoy() repository.FavouriteRepository {
	return &FavouriteRepositoryImpl{}
}

func (repository *FavouriteRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, userId int) []domain.Favourite {
	SQL := "SELECT * FROM favourite WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var favourites []domain.Favourite
	for rows.Next() {
		favourite := domain.Favourite{}
		err := rows.Scan(&favourite.Id, &favourite.MarketId, &favourite.UserId)
		helper.PanicIfError(err)
		favourites = append(favourites, favourite)
	}
	return favourites
}

func (repository *FavouriteRepositoryImpl) Check(ctx context.Context, tx *sql.Tx, favourite domain.Favourite) (domain.Favourite, error) {
	SQL := "SELECT * FROM favourite WHERE market_id = ? AND user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, favourite.MarketId, favourite.UserId)
	helper.PanicIfError(err)
	defer rows.Close()

	favouriteData := domain.Favourite{}
	if rows.Next() {
		err := rows.Scan(&favouriteData.Id, &favouriteData.MarketId, &favouriteData.UserId)
		helper.PanicIfError(err)
		return favouriteData, nil
	} else {
		return favouriteData, errors.New("Data Favorit Tida Ditemukan")
	}
}

func (repository *FavouriteRepositoryImpl) Add(ctx context.Context, tx *sql.Tx, favourite domain.Favourite) domain.Favourite {
	SQL := "INSERT INTO favourite(market_id, user_id) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, SQL, favourite.MarketId, favourite.UserId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	favourite.Id = int(id)
	return favourite
}

func (repository *FavouriteRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, favourite domain.Favourite) {
	SQL := "DELETE FROM favourite WHERE market_id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, favourite.MarketId, favourite.UserId)
	helper.PanicIfError(err)
}