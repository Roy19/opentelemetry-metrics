package service_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"signoz-test/dto"
	"signoz-test/errors"
	"signoz-test/metrics"
	"signoz-test/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestService_AddItemToCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := service.NewService(db)
	ctx := context.Background()
	cartName := "cart1"
	itemName := "item1"
	cartItem := dto.AddToCart{CartName: &cartName, ItemName: &itemName}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select id from public.carts where name = $1")).
		WithArgs(cartName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec(regexp.QuoteMeta("insert into public.items (name, cart_id) values ($1, $2)")).
		WithArgs(itemName, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.AddItemToCart(ctx, cartItem)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_AddItemToCart_CartNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := service.NewService(db)
	ctx := context.Background()
	cartName := "cart1"
	itemName := "item1"
	cartItem := dto.AddToCart{CartName: &cartName, ItemName: &itemName}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select id from public.carts where name = $1")).
		WithArgs(cartName).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	err = s.AddItemToCart(ctx, cartItem)
	assert.ErrorIs(t, err, errors.ErrCartDoesNotExists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_AddItemToCart_AddItemError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := service.NewService(db)
	ctx := context.Background()
	cartName := "cart1"
	itemName := "item1"
	cartItem := dto.AddToCart{CartName: &cartName, ItemName: &itemName}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select id from public.carts where name = $1")).
		WithArgs(cartName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec(regexp.QuoteMeta("insert into public.items (name, cart_id) values ($1, $2)")).
		WithArgs(itemName, 1).
		WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()

	err = s.AddItemToCart(ctx, cartItem)
	assert.ErrorIs(t, err, errors.ErrFailedCartAdd)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_GetItemsInCart(t *testing.T) {
	metrics.InitMeters()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := service.NewService(db)
	ctx := context.Background()
	cartName := "cart1"

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select id from public.carts where name = $1")).
		WithArgs(cartName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectQuery(regexp.QuoteMeta(`select it.id, it.name from 
		public.items it 
		inner join public.carts c 
			on it.cart_id = c.id 
		where it.cart_id = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "item1").AddRow(2, "item2"))
	mock.ExpectCommit()

	items, err := s.GetItemsInCart(ctx, cartName)
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, []string{"item1", "item2"}, items.Items)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_GetItemsInCart_CartNotFound(t *testing.T) {
	metrics.InitMeters()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := service.NewService(db)
	ctx := context.Background()
	cartName := "cart1"

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select id from public.carts where name = $1")).
		WithArgs(cartName).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	items, err := s.GetItemsInCart(ctx, cartName)
	assert.ErrorIs(t, err, errors.ErrCartDoesNotExists)
	assert.Nil(t, items)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_GetItemsInCart_GetItemsError(t *testing.T) {
	metrics.InitMeters()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := service.NewService(db)
	ctx := context.Background()
	cartName := "cart1"

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("select id from public.carts where name = $1")).
		WithArgs(cartName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectQuery(regexp.QuoteMeta(`select it.id, it.name from 
		public.items it 
		inner join public.carts c 
			on it.cart_id = c.id 
		where it.cart_id = $1`)).
		WithArgs(1).
		WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()

	items, err := s.GetItemsInCart(ctx, cartName)
	assert.ErrorIs(t, err, errors.ErrItemsGet)
	assert.Nil(t, items)
	assert.NoError(t, mock.ExpectationsWereMet())
}
