package models

import (
	"database/sql"
	_"database/sql"
	"fmt"
	"repo/test_go_echo/db"
)

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"Name"`
	Salary string `json: "Salary"`
	Age    string `json : "Age"`
}
type Employees struct {
	Employees []Employee `json:"employee"`
}
var con *sql.DB

func GetEmployee() Employees {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT id,Name, Salary, Age FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result

}
