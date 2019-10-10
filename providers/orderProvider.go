package providers

import (
	"log"

	"github.com/MeNoln/golangapi/db"
	"github.com/MeNoln/golangapi/models"
)

//GetAllOrders ...
func GetAllOrders() ([]models.OrderResponseModel, error) {
	var orders []models.OrderResponseModel
	db := db.GetDb()
	defer db.Close()

	query := `select orders.id, orders.ordermessage, orders.orderamount, st.orderstatusname 
			  from orders
			  join orderstatus as st on st.id = orders.id`

	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatalln("Failed tp query")
		return nil, err
	}
	for rows.Next() {
		var order models.OrderResponseModel
		err = rows.StructScan(&order)
		if err != nil {
			log.Fatalln("Failed to get data")
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

//GetCurrentOrder ...
func GetCurrentOrder(id int) (models.OrderResponseModel, error) {
	var order models.OrderResponseModel
	db := db.GetDb()
	defer db.Close()

	query := `select orders.id, orders.ordermessage, orders.orderamount, st.orderstatusname 
	from orders
	join orderstatus as st on st.id = orders.id
	where orders.id = $1`
	queryArgs := []interface{}{
		id,
	}

	err := db.QueryRowx(query, queryArgs...).StructScan(&order)
	if err != nil {
		log.Fatalln("Failed to get data")
		return models.OrderResponseModel{}, err
	}

	return order, nil
}

//CreateNewOrder ...
func CreateNewOrder(data *models.OrderResponseModel) error {
	db := db.GetDb()
	defer db.Close()
	var ID int

	queryID := `select orderstatus.id from orderstatus where orderstatus.orderstatusname = $1`
	err := db.QueryRowx(queryID, "Buy").Scan(&ID)
	if err != nil {
		log.Fatalln("Failed to get id")
		return err
	}

	queryData := `insert into orders (ordermessage, orderamount, orderstatusid) values ($1, $2, $3)`
	queryArgs := []interface{}{
		data.OrderMessage,
		data.OrderAmount,
		ID,
	}
	_, err = db.Exec(queryData, queryArgs...)
	if err != nil {
		log.Fatalln("Failed to insert data")
		return err
	}
	return nil
}

//UpdateOrder ...
func UpdateOrder(data *models.OrderResponseModel) error {
	db := db.GetDb()
	defer db.Close()

	query := `update orders set ordermessage = $1, orderamount = $2 where id = $3`
	queryArgs := []interface{}{
		data.OrderMessage,
		data.OrderAmount,
		data.ID,
	}

	_, err := db.Exec(query, queryArgs...)
	if err != nil {
		log.Fatalln("Failed to update order")
		return err
	}
	return nil
}

//DeleteOrder ...
func DeleteOrder(id int) error {
	db := db.GetDb()
	defer db.Close()

	query := `delete from orders where id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatalln("Failed to update order")
		return err
	}
	return nil
}
