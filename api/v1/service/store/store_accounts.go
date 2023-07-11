package store

import (
	"context"
	"strconv"

	"github.com/suhas-totality/gomongodemo/api/v1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	Model
}

func NewAccount() *Account {
	return &Account{}
}

func (m *Account) CreateNewAccount(ctx context.Context, db *mongo.Database, collection *mongo.Collection, acc *models.Account) (bool, error) {
	err := m.Create(ctx, db, collection, acc)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *Account) ReadAccountByAccountNumber(ctx context.Context, db *mongo.Database, collection *mongo.Collection, accnum *models.AccNumber) (*models.Account, error) {
	filter := bson.M{"account_number": accnum.Accno}
	var accres *models.Account
	err := m.Read(ctx, db, collection, filter, &accres)
	if err != nil {
		return nil, err
	}
	return accres, nil
}

func (m *Account) UpdateAccountByAccountNumber(ctx context.Context, db *mongo.Database, collection *mongo.Collection, acc *models.Account) (bool, error) {
	filter := bson.M{"account_number": acc.AccNumber}
	update := bson.M{"$set": acc}
	err := m.Update(ctx, db, collection, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *Account) DeleteAccountByAccountNumber(ctx context.Context, db *mongo.Database, collection *mongo.Collection, accnum *models.AccNumber) (bool, error) {
	filter := bson.M{"account_number": accnum.Accno}
	err := m.Delete(ctx, db, collection, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *Account) DeleteAccountByAccountNumberMatView(ctx context.Context, db *mongo.Database, collection *mongo.Collection, accnum *models.AccNumber) (bool, error) {
	filter := bson.M{"account_number": strconv.Itoa(int(accnum.Accno))}
	err := m.Delete(ctx, db, collection, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
