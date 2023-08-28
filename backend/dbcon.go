package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	con *sql.DB
}

func NewDatabase() (*Database, error) {
	dbpath := filepath.Join(".", "data", "data.db")
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		fmt.Println("No database file was found, creating a new one")
	}

	conn, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}

	return &Database{conn}, nil
}

func (db *Database) Close() error {
	return db.con.Close()
}

// Executes "filename" sql script located in "sql-scripts" folder
func (db *Database) ExecuteScript(filename string) error {
	file := filepath.Join(".", "sql-scripts", filename)
	sqlScript, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	tx, err := db.con.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(string(sqlScript)); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Inserts new item, returns itemID.
func (db *Database) InsertNewItem(item *Item) (int, error) {
	q := "INSERT INTO item(itemName, itemDescription, itemStatus) VALUES(?,?,?)"
	tx, err := db.con.Begin()
	if err != nil {
		return -1, err
	}
	res, err := tx.Exec(q, item.Name, item.Description, 0)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	id, _ := res.LastInsertId()
	return int(id), tx.Commit()
}

// Inserts a new list, returns listID.
func (db *Database) InsertNewList(list *List) (int, error) {
	q := "INSERT INTO list(listName, listDescription, listStatus, createDate) VALUES(?, ?, ?, datetime('now', 'localtime'))"
	tx, err := db.con.Begin()
	if err != nil {
		return -1, err
	}
	res, err := tx.Exec(q, list.Name, list.Description, list.Status)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	id, _ := res.LastInsertId()
	return int(id), tx.Commit()
}

// Inserts list 2 item relationship.
func (db *Database) InsertListItemRel(listID, itemID int) error {
	q := "INSERT INTO listItem(listID, itemID) VALUES(?, ?)"
	tx, err := db.con.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(q, listID, itemID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (db *Database) getAllLists() (*[]List, error) {
	q := "SELECT * FROM list"
	rows, err := db.con.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var lists []List
	for rows.Next() {
		l := List{}
		err := rows.Scan(&l.ListID, &l.Name, &l.Description, &l.Status, &l.CreateDate, &l.CompleteDate)
		if err != nil {
			println(err.Error())
		}
		lists = append(lists, l)
	}

	return &lists, nil
}

func (db *Database) getListItems(listID int) (*[]Item, error) {
	q := "SELECT * FROM item NATURAL JOIN listItem WHERE LISTID=?"
	rows, err := db.con.Query(q, listID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []Item
	for rows.Next() {
		i := Item{}
		err := rows.Scan(&i.ItemID, &i.Name, &i.Description, &i.Status, &i.ListID)
		if err != nil {
			println(err.Error())
		}
		items = append(items, i)
	}

	return &items, nil
}

func (db *Database) getListByID(listID int) (*List, error) {
	q := "SELECT * FROM list WHERE LISTID=?"
	rows, err := db.con.Query(q, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	l := List{}
	err = rows.Scan(&l.ListID, &l.Name, &l.Description, &l.Status, &l.CreateDate, &l.CompleteDate)
	return &l, err
}

func (db *Database) getItemByID(itemID int) (*Item, error) {
	q := "SELECT * FROM item NATURAL JOIN listItem WHERE itemID=?"
	rows, err := db.con.Query(q, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	it := Item{}
	err = rows.Scan(&it.ItemID, &it.Name, &it.Description, &it.ListID, &it.Status)
	return &it, err
}

func (db *Database) updateItem(it *Item) {
	q := "UPDATE item SET itemName=?, itemDescription=? WHERE itemID=?"
	tx, err := db.con.Begin()
	if err != nil {
		println(err.Error())
		return
	}

	_, err = tx.Exec(q, it.Name, it.Description, it.ItemID)
	if err != nil {
		println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
}

func (db *Database) delteItem(it *Item) {
	tx, err := db.con.Begin()
	if err != nil {
		println(err.Error())
		return
	}

	q1 := "DELETE FROM listItem WHERE itemID=?"
	q2 := "DELETE FROM item WHERE itemID=?"
	_, err = tx.Exec(q1, it.ItemID)
	if err != nil {
		println(err.Error())
		tx.Rollback()
		return
	}
	_, err = tx.Exec(q2, it.ItemID)
	if err != nil {
		println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()

}

func (db *Database) completeItem(it *Item) {
	q := "UPDATE item SET itemStatus=? WHERE itemID=?"
	tx, err := db.con.Begin()
	if err != nil {
		println(err.Error())
		return
	}

	_, err = tx.Exec(q, it.Status, it.ItemID)
	if err != nil {
		println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
}

func (db *Database) updateList(li *List) {
	q := "UPDATE list SET listName=?, listDescription=? WHERE listID=?"
	tx, err := db.con.Begin()
	if err != nil {
		println(err.Error())
		return
	}

	_, err = tx.Exec(q, li.Name, li.Description, li.ListID)
	if err != nil {
		println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
}
func (db *Database) deleteList(li *List) {
	tx, err := db.con.Begin()
	if err != nil {
		println(err.Error())
		return
	}

	q := "DELETE FROM list WHERE listID=?"
	_, err = tx.Exec(q, li.ListID)
	if err != nil {
		println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
}

func (db *Database) changeListStatus(li *List) {
	q := "UPDATE list SET listStatus=? WHERE listID=?"
	tx, err := db.con.Begin()
	if err != nil {
		println(err.Error())
		return
	}

	_, err = tx.Exec(q, li.Status, li.ListID)
	if err != nil {
		println(err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
}
