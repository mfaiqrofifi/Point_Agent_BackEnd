package history

import (
	"Final_Project/business/history"
	"Final_Project/drivers/databases/product"
	"Final_Project/drivers/databases/redem"
	"Final_Project/drivers/databases/users"
	"Final_Project/helpers/midtrans"
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MySqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) history.Repository {
	return &MySqlUserRepository{
		Conn: conn,
	}
}
func (rep MySqlUserRepository) RequestProduct(ctx context.Context, idUser int, idProduct *int, amount int, img string) (history.History, error) {
	var history HistoryDB
	var user users.Users
	var produk product.ProductDB

	history.UsersID = idUser
	history.ProductDBID = idProduct
	history.Amount = amount
	history.ImageRequest = img
	history.Status = "Request"
	history.Type = "Product"
	resultProduct := rep.Conn.First(&produk, "id = ?", idProduct)
	history.Product = produk
	if resultProduct.Error != nil {
		return history.ToDomain(), resultProduct.Error
	}
	history.PoinItems = amount * (produk.Poin)

	poinUser := rep.Conn.First(&user, "id=?", idUser)
	history.User = user
	if poinUser.Error != nil {
		return history.ToDomain(), poinUser.Error
	}
	result := rep.Conn.Create(&history)
	if result.Error != nil {
		return history.ToDomain(), result.Error
	}
	return history.ToDomain(), nil
}

func (rep MySqlUserRepository) AllowProduct(ctx context.Context, idUser int, stsatus string) (history.History, error) {
	var history HistoryDB
	var user users.Users
	// resultuser := rep.Conn.Model(&history).Where("id = ?", idUser).Update("status", stsatus)
	// if resultuser.Error != nil {
	// 	return history.ToDomain(), resultuser.Error
	// }
	hasil := rep.Conn.First(&history, "id=?", idUser)
	if  history.Status!= "Request" {
		return history.ToDomain(), errors.New("cannot process twice")
	}
	if hasil.Error != nil {
		return history.ToDomain(), hasil.Error
	}
	resultuser := rep.Conn.Model(&history).Where("id = ?", idUser).Update("status", stsatus)
	if resultuser.Error != nil {
		return history.ToDomain(), resultuser.Error
	}
	poinUser := rep.Conn.First(&user, "id=?", history.UsersID)
	if poinUser.Error != nil {
		return history.ToDomain(), poinUser.Error
	}
	poin := user.Poin + history.PoinItems
	resultPoin := rep.Conn.Model(&user).Where("id = ?", history.UsersID).Update("poin", poin)
	if resultPoin.Error != nil {
		return history.ToDomain(), resultPoin.Error
	}
	return history.ToDomain(), nil
}

func (rep MySqlUserRepository) ViewHistoryUser(ctx context.Context, idUser int) ([]history.History, error) {
	var historydb []HistoryDB

	result := rep.Conn.Preload(clause.Associations).Order("id desc").Find(&historydb, "users_id=?", idUser)
	if result.Error != nil {
		return []history.History{}, result.Error
	}
	return ToListDeteil(historydb), nil
}

func (rep MySqlUserRepository) ViewRequestUser(ctx context.Context) ([]history.History, error) {
	var historydb []HistoryDB
	result := rep.Conn.Preload(clause.Associations).Order("id desc").Find(&historydb, "type=?", "Product")
	// result := rep.Conn.Find(&historydb, "type=?", "Product")
	if result.Error != nil {
		return []history.History{}, result.Error
	}
	return ToListDeteil(historydb), nil
}

func (rep MySqlUserRepository) RequestRedem(ctx context.Context, idUser int, idReward *int, amount int, identity string) (history.History, error) {
	var history HistoryDB
	var user users.Users
	var redem redem.RedemDB
	history.Amount = amount
	history.UsersID = idUser
	history.RedemDBID = idReward
	history.Type = "Reward"
	poinItem := rep.Conn.First(&redem, "id=?", idReward)
	if poinItem.Error != nil {
		return history.ToDomain(), poinItem.Error
	}
	if redem.NameType == "Cash Out" {
		history.NomorRekening = identity
	}
	if redem.NameType == "Pulsa" || redem.NameType == "E-Money" {
		history.NomoHp = identity
	}

	history.PoinItems = amount * (redem.Poin)
	UsersUP := rep.Conn.First(&user, "id=?", idUser)
	if UsersUP.Error != nil {
		return history.ToDomain(), UsersUP.Error
	}
	hsil := user.Poin - history.PoinItems
	if hsil < 0 {
		return history.ToDomain(), errors.New("your poin is not enought")
	}
	resultuser := rep.Conn.Model(&user).Where("id = ?", idUser).Update("poin", hsil)
	if resultuser.Error != nil {
		return history.ToDomain(), resultuser.Error
	}
	
	history.User = user
	history.Reward = redem
	jumlah := history.Amount * redem.NominalReward
	if redem.NameType != "Pulsa" {
		hasil, err := midtrans.GetPembayaranUrl(user.Toko, user.Email, jumlah, history.Id)
		if err != nil {
			return history.ToDomain(), err
		}
		history.MidtransLink = hasil
	}
	result := rep.Conn.Create(&history)
	if result.Error != nil {
		return history.ToDomain(), result.Error
	}
	return history.ToDomain(), nil
}

func (rep MySqlUserRepository) ViewRedem(ctx context.Context) ([]history.History, error) {
	var historydb []HistoryDB
	result := rep.Conn.Preload(clause.Associations).Order("id desc").Find(&historydb, "type=?", "Reward")
	if result.Error != nil {
		return []history.History{}, result.Error
	}
	return ToListDeteil(historydb), nil
}
