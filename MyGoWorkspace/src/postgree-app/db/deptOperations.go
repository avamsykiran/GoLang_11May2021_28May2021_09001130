package db

import (
	"database/sql"
	"fmt"

	"postgree-app/model"

	_ "github.com/lib/pq"
)

func AddDept(d model.Dept) model.Dept {
	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	insStmt := "insert into dept values($1,$2,$3)"

	res, err := conn.Exec(insStmt, d.Deptno, d.Dname, d.Loc)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d rows added successfully", count)

	return d
}

func UpdateDept(d model.Dept) model.Dept {
	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	updStmt := "update dept set dname=$1, loc=$2 where deptno=$3"

	res, err := conn.Exec(updStmt, d.Dname, d.Loc, d.Deptno)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d rows updated successfully", count)

	return d
}

func DelDept(deptno int) {
	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	delStmt := "delete from dept where deptno=$1"

	res, err := conn.Exec(delStmt, deptno)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d rows deleted successfully", count)
}

func GetDeptById(deptno int) model.Dept {

	var dept model.Dept

	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	selStmt := "select deptno,dname,loc from dept where deptno=$1"

	row := conn.QueryRow(selStmt, deptno)

	err = row.Scan(&dept.Deptno, &dept.Dname, &dept.Loc)

	if err != nil {
		panic(err)
	}

	return dept
}

func GetAllDepts() []model.Dept {

	var depts []model.Dept

	conn, err := sql.Open("postgres", ConnString())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	selStmt := "select deptno,dname,loc from dept"

	rows, err := conn.Query(selStmt)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var d model.Dept
		rows.Scan(&d.Deptno, &d.Dname, &d.Loc)
		depts = append(depts, d)
	}

	return depts
}
