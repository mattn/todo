package main

import (
	"database/sql"
	"fmt"
	"time"
)

type TodoItem struct {
	Todo   string
	Done   int
	Id     int
	Status string
}

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, fmt.Errorf("database doesnt exist")
	}
	return db, nil
}

func CreateTable(db *sql.DB, project string) error {
	// create table if not exists
	sql_table := "CREATE TABLE IF NOT EXISTS " + project + "(todo TEXT,done INTEGER,status TEXT, id INTEGER PRIMARY KEY AUTOINCREMENT)"
	_, err := db.Exec(sql_table)
	if err != nil {
		return err
	}
	return nil
}

func CheckTodo(db *sql.DB, project, todo string) error {
	sql_checktodo := "UPDATE " + project + " SET done=1 WHERE id =" + todo

	_, err2 := db.Exec(sql_checktodo)
	if err2 != nil {
		return err2
	}
	return nil
}

func UnCheckTodo(db *sql.DB, project, todo string) error {

	sql_checktodo := "UPDATE " + project + " SET done=0 WHERE id =" + todo

	_, err2 := db.Exec(sql_checktodo)
	if err2 != nil {
		return fmt.Errorf("Error on execute for CheckTodo")
	}
	return nil
}

func StoreTodo(db *sql.DB, project, todo string) error {
	sql_addtodo := `
	INSERT OR REPLACE INTO ` + project + `(
		todo,
		done,
		status
	) values(?,?,?)
	`

	stmt, err := db.Prepare(sql_addtodo)
	if err != nil {
		return fmt.Errorf("Error on prepare todo storetodo")
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(todo, 0, fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05")))
	if err2 != nil {
		return fmt.Errorf("Error on execute for StoreTodo")
	}
	return nil
}

func ReadTodos(db *sql.DB, project string) ([]TodoItem, error) {
	sql_readall := "SELECT * FROM " + project + " order by id ASC"
	rows, err := db.Query(sql_readall)
	if err != nil {
		return nil, fmt.Errorf("Error on querying todo description")
	}
	defer rows.Close()

	var result []TodoItem
	for rows.Next() {
		item := TodoItem{}

		err2 := rows.Scan(&item.Todo, &item.Done, &item.Status, &item.Id)
		if err2 != nil {
			return nil, err2
		}

		result = append(result, item)
	}
	return result, nil
}

func DeleteTodo(db *sql.DB, project, todo string) error {
	sql_deltodo := "DELETE FROM " + project + " WHERE id = " + todo
	_, err2 := db.Exec(sql_deltodo)

	if err2 != nil {
		return fmt.Errorf("Error on execute for DeleteTodo")
	}
	return nil

}

func DeleteAllTodos(db *sql.DB, project string) error {
	sql_deltodo := "DELETE FROM  " + project

	_, err2 := db.Exec(sql_deltodo)

	if err2 != nil {
		return fmt.Errorf("Error on execute for DeleteTodo")
	}
	sql_delautoinc := "delete from sqlite_sequence where name ='" + project + "'"
	_, err2 = db.Exec(sql_delautoinc)

	if err2 != nil {
		return fmt.Errorf("Error on execute for DeleteTodo")
	}
	return nil

}

func DeleteProject(db *sql.DB, project string) error {
	sql_deletetable := "DROP TABLE " + project
	_, err2 := db.Exec(sql_deletetable)

	if err2 != nil {
		return fmt.Errorf("Error on execute for DeleteTodo")
	}
	return nil

}
