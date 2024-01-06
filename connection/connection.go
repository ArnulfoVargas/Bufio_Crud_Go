package connection

import (
	"database/sql"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
) 

var DB *sql.DB 
func Connect() { 
  err := godotenv.Load() 
  if err != nil { panic(err) } 

  sb := strings.Builder{}

  sb.WriteString(os.Getenv("DB_USER") + ":")
  sb.WriteString(os.Getenv("DB_PASSWORD") + "@")
  sb.WriteString(os.Getenv("DB_SERVER") + ":")
  sb.WriteString(os.Getenv("DB_PORT") + "/")
  sb.WriteString(os.Getenv("DB_NAME"))
  
  connection, err := sql.Open("postgres", sb.String())

  if err != nil { panic(err) }


  DB = connection;
}

func CloseConnection()  {
  err := DB.Close()
  
  if err != nil { panic(err) }
}
