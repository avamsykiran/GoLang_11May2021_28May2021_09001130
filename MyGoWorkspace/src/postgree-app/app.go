package main

import (
	"fmt"
	"postgree-app/db"
)

func main() {
	//db.VerifyConnection()

	/*d := model.Dept{101, "AUDIT", "NEW DELHI"}
	db.AddDept(d)*/

	/*d := model.Dept{101, "AUDIT DEPT", "DELHI"}
	db.UpdateDept(d)*/

	//db.DelDept(101)

	//fmt.Println(db.GetDeptById(30))
	fmt.Println(db.GetAllDepts())
}
