package store

import (
	"fmt"

	"github.com/15110102/phuongpt3-market-server/src/model"
)

func (s Store) CreateOrder(order *model.Order) (*model.Order, error) {
	query := fmt.Sprintf("INSERT INTO Orders(AppUser, AppTransId, ZpTransToken, Item, CreateAt, TotalPrice, Status) VALUES ('%s', '%s','%s', '%s', %d, %d, '%s');", order.AppUser, order.AppTransId, "", order.Item, order.CreateAt, order.TotalPrice, order.Status)
	res, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	order.Id = id
	return order, nil
}

func (s Store) UpdateStatusOrderByTrans(transId string, status string) (bool, error) {
	updateQuery := fmt.Sprintf("Update Orders Set Status = '%s' Where AppTransId = '%s'", status, transId)
	updateStatus, err := db.Query(updateQuery)
	if err != nil {
		return false, err
	}
	defer updateStatus.Close()
	return true, nil
}

func (s Store) UpdateZpTransTokenOrderById(orderId int64, zpTransToken string) (bool, error) {
	updateQuery := fmt.Sprintf("Update Orders Set ZpTransToken = '%s' Where Id = %d", zpTransToken, orderId)
	updateZpTransToken, err := db.Query(updateQuery)
	if err != nil {
		return false, err
	}
	defer updateZpTransToken.Close()
	return true, nil
}

func (s Store) GetOrder(orderId int64) (*model.Order, error) {
	var order model.Order
	err := db.QueryRow("SELECT * FROM Orders WHERE Id = ?", orderId).Scan(&order.Id, &order.AppUser, &order.AppTransId, &order.ZpTransToken, &order.Item, &order.CreateAt, &order.TotalPrice, &order.Status)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
