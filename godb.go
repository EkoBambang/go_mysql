package main
import (
"fmt"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)

func main() {

// Create the database handle, confirm driver is present
db, _ := sql.Open("mysql", "root:@/praweb")
defer db.Close()

// Connect and check the server version
var version string 
db.QueryRow("SELECT VERSION()").Scan(&version)
fmt.Println("Connected to:", version)
}
