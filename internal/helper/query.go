package helper

import (
	"database/sql"
	"log"
)

// ExecuteQuery executes the query on the given DB and returns rows
func ExecuteQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("==== QUERY ERROR ====")
		log.Printf("Query: %s\n", query)
		log.Printf("Args: %v\n", args)
		log.Printf("Error: %v\n", err)

		return nil, err
	}
	return rows, nil
}
