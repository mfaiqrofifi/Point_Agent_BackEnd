package history

import (
	"Final_Project/business/history"
	"Final_Project/drivers/databases/product"
	"Final_Project/drivers/databases/redem"
	"Final_Project/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type HistoryDB struct {
	Id            int `gorm:"primaryKey"`
	ProductDBID   *int
	UsersID       int
	RedemDBID     *int
	Amount        int
	Status        string
	PoinItems     int
	Type          string
	ImageRequest  string
	NomoHp        string
	NomorRekening string
	Product       product.ProductDB `gorm:"foreignKey:ProductDBID"`
	User          users.Users       `gorm:"foreignKey:UsersID"`
	Reward        redem.RedemDB     `gorm:"foreignKey:RedemDBID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (historyDB *HistoryDB) ToDomain() history.History {
	return history.History{
		Id:            historyDB.Id,
		IdUser:        historyDB.UsersID,
		IdProduct:     historyDB.ProductDBID,
		IdReward:      historyDB.RedemDBID,
		Amount:        historyDB.Amount,
		Status:        historyDB.Status,
		PoinItems:     historyDB.PoinItems,
		Type:          historyDB.Type,
		ImageRequest:  historyDB.ImageRequest,
		NomoHp:        historyDB.NomoHp,
		NomorRekening: historyDB.NomorRekening,
		CreatedAt:     historyDB.CreatedAt,
		UpdatedAt:     historyDB.UpdatedAt,
		Product:       historyDB.Product.ToDomain(),
		User:          historyDB.User.ToDomain(),
		Reward:        historyDB.Reward.ToDomain(),
	}
}

func FromDomain(domain history.History) HistoryDB {
	return HistoryDB{
		Id:            domain.Id,
		UsersID:       domain.IdUser,
		ProductDBID:   domain.IdProduct,
		RedemDBID:     domain.IdReward,
		Amount:        domain.Amount,
		Status:        domain.Status,
		PoinItems:     domain.PoinItems,
		Type:          domain.Type,
		ImageRequest:  domain.ImageRequest,
		NomoHp:        domain.NomoHp,
		NomorRekening: domain.NomorRekening,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		Product:       product.FromDomain(domain.Product),
		User:          users.FromDomain(domain.User),
		Reward:        redem.FromDomain(domain.Reward),
	}
}

func ToListDeteil(data []HistoryDB) (result []history.History) {
	result = []history.History{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
