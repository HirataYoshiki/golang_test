package client

import (
	i "golang_test/dbinterface"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type MySQLClient struct {
	Host string
	Port int
	User string
	Pass string
	DB string
}

func New(message string) DBError {
	return DBError{Message: message}
}

type DBError struct {
	Message string
}

func (err *DBError) Error() string {
	return err.message
}

func (mysql *MySQLClient) getURI (uri string, err error) {
	host := mysql.Host
	if host == "" {
		err := New("environ DBHost not found.")
		return host, err
	}
	
	port := mysql.Port
	if port == "" {
		err := New("environ DBPORT not found.")
		return port, err
	}

	user := mysql.User
	if user == "" {
		err := New("environ MYSQLUSER not found.")
		return user, err
	}

	pass := mysql.Pass
	if pass == "" {
		err := New("environ MYSQLPASS not found.")
		return pass, err
	}

	db := mysql.DB
	if db == "" {
		err := New("environ DB not found.")
		return db, err
	}

	uri = user + ":" + pass + "@" + host + "/" + db
	return uri, nil
}

func (mysql *MySQLClient) GetAll(table string) (response i.clientResponse, err error) {
	database, err := sql.Open("mysql", mysql.getURI())
	if err != nil {
		return nil err
	}

	defer database.Close()

	rows, err := database.Query("SELECT * FROM %s", table)
	if err != nil {
		return nil, err

	}// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}

	result := i.clientResponse{Columns: columns, Values: value}
	return result, nil

	
}