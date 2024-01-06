package connection

import (
	"database/sql"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
) 

var DB *sql.DB 
func ConnectToDB() { 
  err := godotenv.Load() 
  if err != nil { panic(err) } 

  sb := strings.Builder{}

  sb.WriteString("user="+ os.Getenv("DB_USER") + " ")
  sb.WriteString("dbname=" + os.Getenv("DB_NAME") + " ")
  sb.WriteString("sslmode=" + os.Getenv("SSLMode"))
  
  connection, err := sql.Open("postgres", sb.String())

  if err != nil { panic(err) }


  DB = connection;
}

func CloseConnectionToDB()  {
  err := DB.Close()
  
  if err != nil { panic(err) }
}
