package main

import (
	"database/sql"
	"fmt"
)

type TodoItem struct {
	Description string
	Done        int
	Index       int
}

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	fmt.Println(filepath)
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
	sql_table := `
	CREATE TABLE IF NOT EXISTS ` + project + `(
		todo TEXT,
		done INTEGER,
		ind  INTEGER AUTOINCREMENT
	); 
	`
	_, err := db.Exec(sql_table)
	if err != nil {
		return err
	}
	return nil
}

func CheckTodo(db *sql.DB, project, todo string) error {
	sql_checktodo := `
	UPDATE ` + project + `
	SET done=1
	WHERE ind =` + todo + `
	`

	stmt, err := db.Prepare(sql_checktodo)
	if err != nil {
		return fmt.Errorf("Error on prepare checktodo")
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(sql_checktodo)
	if err2 != nil {
		return fmt.Errorf("Error on execute for CheckTodo")
	}
	return nil
}

func UnCheckTodo(db *sql.DB, project, todo string) error {
	sql_checktodo := `
	UPDATE ` + project + `
	SET done=0
	WHERE ind =` + todo + `
	`

	stmt, err := db.Prepare(sql_checktodo)
	if err != nil {
		return fmt.Errorf("Error on prepare checktodo")
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(sql_checktodo)
	if err2 != nil {
		return fmt.Errorf("Error on execute for CheckTodo")
	}
	return nil
}

func StoreTodo(db *sql.DB, project, todo string) error {
	sql_addtodo := `
	INSERT OR REPLACE INTO ` + project + `(
		todo,
		done
	) values(?,?)
	`

	stmt, err := db.Prepare(sql_addtodo)
	if err != nil {
		return fmt.Errorf("Error on prepare todo storetodo")
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(todo, 0)
	if err2 != nil {
		return fmt.Errorf("Error on execute for StoreTodo")
	}
	return nil
}

func ReadTodos(db *sql.DB, project string) ([]TodoItem, error) {
	sql_readall := `
	SELECT todo,done FROM ` + project + `
	 ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		return nil, fmt.Errorf("Error on querying todo description")
	}
	defer rows.Close()

	var result []TodoItem
	for rows.Next() {
		item := TodoItem{}
		err2 := rows.Scan(&item.Description, &item.Done)
		if err2 != nil {
			return nil, fmt.Errorf("Error on Scan todo description")
		}
		result = append(result, item)
	}
	return result, nil
}

func DeleteTodo(db *sql.DB, project, todo string) error {
	sql_deltodo := `
	DELETE FROM ` + project + `
	WHERE ind = ` + todo + `
	`

	stmt, err := db.Prepare(sql_deltodo)
	if err != nil {
		return fmt.Errorf("Error on prepare for DeleteTodo")
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(sql_deltodo)

	if err2 != nil {
		return fmt.Errorf("Error on execute for DeleteTodo")
	}
	return nil

}

func DeleteAllTodos(db *sql.DB, project string) error {
	sql_deltodo := `
	DELETE FROM ` + project + ``

	stmt, err := db.Prepare(sql_deltodo)
	if err != nil {
		return fmt.Errorf("Error on prepare for DeleteAllTodos")
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(sql_deltodo)

	if err2 != nil {
		return fmt.Errorf("Error on execute for DeleteTodo")
	}
	return nil

}

func DeleteProject(db *sql.DB, project string) error {
	sql_deletetable := "DROP TABLE " + project
	stmt, err := db.Prepare(sql_deletetable)
	if err != nil {
		return fmt.Errorf("Error on prepare for DeleteTodo")
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(sql_deletetable)

	if err2 != nil {
		return fmt.Errorf("Error on execute for DeleteTodo")
	}
	return nil

}
