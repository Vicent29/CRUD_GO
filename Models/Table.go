package Models

import (
	"CRUD-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllTables Fetch all Table data
func GetAllTables(table *[]Table) (err error) {
	if err = Config.DB.Find(table).Error; err != nil {
		return err
	}
	return nil
}

// CreateTable ... Insert New data
func CreateTable(table *Table) (err error) {
	if err = Config.DB.Create(table).Error; err != nil {
		return err
	}
	return nil
}

// GetTableByID ... Fetch only one Table by Id
func GetTableByID(table *Table, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(table).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTable ... Update Table
func UpdateTable(table *Table, id string) (err error) {
	fmt.Println(table)
	Config.DB.Save(table)
	return nil
}

// DeleteTable ... Delete Table
func DeleteTable(table *Table, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(table)
	return nil
}

