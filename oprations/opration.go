package operation

import (
	"encoding/json"
	"fmt"
	"log"

	data "todoList/datbase"
)

type Items struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

func Add(listName string, itemList []Items) {
	itemsJSON, err := json.Marshal(itemList)
	if err != nil {
		log.Fatalf("Error serializing itemList: %v", err)
	}

	query := "INSERT INTO list (list_name, items, user_id) VALUES (?, ?, ?)"

	result := data.Db.Exec(query, listName, string(itemsJSON), 1)
	if result.Error != nil {
		log.Fatalf("Error inserting data into the database: %v", result.Error)
	} else {
		fmt.Printf("Successfully added list: %s\n", listName)
	}

	for _, item := range itemList {
		fmt.Printf("Item Name: %s, Quantity: %d\n", item.Name, item.Qty)
	}
}

func Update(listID int, listName string, itemList []Items) {
	itemsJSON, err := json.Marshal(itemList)
	if err != nil {
		log.Fatalf("Error serializing itemList: %v", err)
	}
	query := "UPDATE list SET list_name = ?, items = ?, user_id = ? WHERE id = ?"

	result := data.Db.Exec(query, listName, string(itemsJSON), 2, listID)
	if result.Error != nil {
		log.Fatalf("Error updating data in the database: %v", result.Error)
	} else {
		fmt.Printf("Successfully updated list with ID: %d\n", listID)
	}

}

func Delete(listID int) {

	query := "DELETE FROM list WHERE id = ?"

	result := data.Db.Exec(query, listID)
	if result.Error != nil {
		log.Fatalf("Error deleting data from the database: %v", result.Error)
	} else {
		fmt.Printf("Successfully deleted list with ID: %d\n", listID)
	}

}
