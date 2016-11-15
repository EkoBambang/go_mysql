package main

import "fmt"
import "database/sql"
import "os"
import _ "github.com/go-sql-driver/mysql"

var user, pas, table, nama_table, exit string

func connect() *sql.DB {

	var db, err = sql.Open("mysql", user+":"+pas+"@/"+table)
	err = db.Ping()
	if err != nil {
		fmt.Println("Gagal Membuka Database")
		fmt.Println("Database Tidak Bisa Dihubungi")
		os.Exit(0)
	}
	return db
}

func outputsql() {
	var db = connect()
	defer db.Close()

	var id_mhs, nama_mhs string

	fmt.Println(" ")
	fmt.Println("Menampilkan Data :")
	fmt.Println(" ")

	rows, _ := db.Query("select * from " + nama_table)

	for rows.Next() {
		rows.Scan(&id_mhs, &nama_mhs)
		fmt.Println("id_mhs :" + id_mhs)
		fmt.Println("nama_mhs :" + nama_mhs)
		fmt.Println(" ")

	}
	fmt.Println("exit press enter")
	fmt.Scanln(&exit)
}

func main() {
	fmt.Print("Masukkan User mysql = ")
	fmt.Scanln(&user)
	fmt.Print("Masukkan Password mysql = ")
	fmt.Scanln(&pas)
	fmt.Print("Masukan Nama Database = ")
	fmt.Scanln(&table)
	fmt.Print("Masukan Nama Tabel = ")
	fmt.Scanln(&nama_table)
	fmt.Println(" ")
	outputsql()

}