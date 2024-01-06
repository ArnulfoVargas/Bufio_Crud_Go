package handlers

import (
	"bufio"
	"bufiodb/connection"
	"bufiodb/models"
	"fmt"
	"os"
	"strconv"
)

func List() {
	connection.ConnectToDB()

  query := "SELECT id, name, mail, phone FROM clients;"

  data, err := connection.DB.Query(query)
  if err != nil { panic(err) }
  defer connection.CloseConnectionToDB()

  clients := models.Clients{}

  for data.Next() {
    client := models.Client{}
    if err := data.Scan(&client.Id, &client.Name, &client.Mail, &client.Phone); err != nil {
      fmt.Println(err)
      break
    }

    clients = append(clients, client)
  }

  if len(clients) > 0 {
    fmt.Println("---------------------------------------------------------------------------------------------------------")
    for _, value := range clients{
      fmt.Printf("| ID: %d \t\t | Name: %s \t\t | Mail: %s \t\t | Phone: %s \t|\n", value.Id, value.Name, value.Mail,value.Phone)
      fmt.Println("---------------------------------------------------------------------------------------------------------")
    }
  }
}

func ListByID(id int) {
  if id <= 0 {
    fmt.Println("----------------")
    fmt.Println("| ID NOT VALID |")
    fmt.Println("----------------")

    return
  }
  
	connection.ConnectToDB()

  query := "SELECT id, name, mail, phone FROM clients WHERE id=$1;"

  data := connection.DB.QueryRow(query, id)
  defer connection.CloseConnectionToDB()

  client := models.Client{}
  data.Scan(&client.Id, &client.Name, &client.Mail, &client.Phone)

  if client.Id == 0 {
    fmt.Println("--------------------")
    fmt.Println("| CLIENT NOT FOUND |")
    fmt.Println("--------------------")

    return
  }

  fmt.Println("---------------------------------------------------------------------------------------------------------")
  fmt.Printf("| ID: %d \t\t | Name: %s \t\t | Mail: %s \t\t | Phone: %s \t|\n", client.Id, client.Name, client.Mail,client.Phone)
  fmt.Println("---------------------------------------------------------------------------------------------------------")
}

func InsertClient(client models.Client) {
  connection.ConnectToDB()

  exec := "INSERT INTO clients(name, mail, phone, date) VALUES ($1, $2, $3, NOW());"

  _, err := connection.DB.Exec(exec, client.Name, client.Mail, client.Phone)
  defer connection.CloseConnectionToDB()

  if err != nil { panic(err) }

  fmt.Println("Client Added")
}

func UpdateClient(client models.Client, id int) {
  connection.ConnectToDB()

  exec := "UPDATE clients SET (name, mail, phone) = ($1, $2, $3) WHERE id=$4;"

  _, err := connection.DB.Exec(exec, client.Name, client.Mail, client.Phone, id)
  defer connection.CloseConnectionToDB()

  if err != nil { panic(err) }

  fmt.Println("Client Updated")
}

func DeleteClient(id int) {
  connection.ConnectToDB()

  exec := "DELETE FROM clients WHERE id=$1;"

  _, err := connection.DB.Exec(exec, id)
  defer connection.CloseConnectionToDB()

  if err != nil { panic(err) }

  fmt.Println("Client Deleted")
}

// Utils
// var id int;
var name, mail, phone string;

func Run() {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("\n--- SELECT AN OPTION ---\n\n")
    fmt.Println("1- List Clients")
    fmt.Println("2- Search Client by ID")
    fmt.Println("3- Add a Client")
    fmt.Println("4- Edit a Client")
    fmt.Println("5- Delete a Client")

    fmt.Print("\nYour Option: ")

    if scanner.Scan() {
      switch {
      case scanner.Text() == "1":
        optionListAll()
        return
      case scanner.Text() == "2":
        optionListByID(scanner)
        return
      case scanner.Text() == "3":
        optionAddClient(scanner)
        continue
      case scanner.Text() == "4":
        optionEditClient(scanner)
        continue
      case scanner.Text() == "5":
        optionDeleteClient(scanner)
        continue
      default:
        fmt.Print("Error: Not valid option\n")
        continue
      }
    }
  }
}

func optionListAll() {
  List()
}

func optionListByID(scanner *bufio.Scanner) {
  fmt.Print("Type the client ID: ")
  if scanner.Scan() {
    id, err := strconv.Atoi(scanner.Text())
    
    if err != nil { panic("Cannot get id")}

    ListByID(id)
  }
}

func optionAddClient(scanner *bufio.Scanner) {
  var name, mail, phone string;
  
  fmt.Print("Type the name: ")
  if scanner.Scan() {
    name = scanner.Text()
  }

  fmt.Print("Type the e-mail: ")
  if scanner.Scan() {
    mail = scanner.Text()
  }
  
  fmt.Print("Type the phone number: ")
  if scanner.Scan() {
    phone = scanner.Text()
  }

  client := models.Client{
    Id: 0,
    Name: name,
    Mail: mail,
    Phone: phone,
  }

  InsertClient(client) 
}

func optionEditClient(scanner *bufio.Scanner) {
  var name, mail, phone string;
  var id int = -1;

  fmt.Print("Type the client ID: ")
  if scanner.Scan() {
    dummyId, err := strconv.Atoi(scanner.Text())
    
    if err != nil { panic("Cannot get id")}

    id = dummyId
  }

  if id < 0 { panic("Cannot get id") }

  fmt.Print("Type the name: ")
  if scanner.Scan() {
    name = scanner.Text()
  }

  fmt.Print("Type the e-mail: ")
  if scanner.Scan() {
    mail = scanner.Text()
  }
  
  fmt.Print("Type the phone number: ")
  if scanner.Scan() {
    phone = scanner.Text()
  }

  client := models.Client{
    Id: 0,
    Name: name,
    Mail: mail,
    Phone: phone,
  }

  UpdateClient(client, id) 
}

func optionDeleteClient(scanner *bufio.Scanner) {
  fmt.Print("Type the client ID: ")
  if scanner.Scan() {
    id, err := strconv.Atoi(scanner.Text())
    
    if err != nil { panic("Cannot get id")}

    DeleteClient(id)
  }
}
