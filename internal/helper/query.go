// package helper

// import (
// 	"database/sql"
// 	"log"
// )

// // ExecuteQuery executes the query on the given DB and returns rows
// func ExecuteQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
// 	rows, err := db.Query(query, args...)
// 	if err != nil {
// 		log.Println("==== QUERY ERROR ====")
// 		log.Printf("Query: %s\n", query)
// 		log.Printf("Args: %v\n", args)
// 		log.Printf("Error: %v\n", err)

// 		return nil, err
// 	}
// 	return rows, nil
// }

package helper

import (
	"database/sql"
	"log"
)

// ExecuteQuery executes the query on the given DB and returns rows
func ExecuteQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	// Log the query and its arguments for debugging
	log.Printf("Executing Query: %s\n", query)
	log.Printf("Arguments: %v\n", args)

	// Use db.Query for executing SELECT queries (PostgreSQL supports the same method)
	rows, err := db.Query(query, args...)
	if err != nil {
		// Log the error details for debugging
		log.Println("==== QUERY ERROR ====")
		log.Printf("Query: %s\n", query)
		log.Printf("Args: %v\n", args)
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	// Return the rows for further processing
	return rows, nil
}
