package main

import "bufiodb/handlers"

func main() {
	// handlers.List()
	// handlers.ListByID(2)
	// handlers.InsertClient(models.Client{
	//   Id: 0,
	//   Name: "Saul",
	//   Mail: "saul@saul.com",
	//   Phone: "2468013579",
	// })
	// handlers.UpdateClient(models.Client{
	//   Id: 0,
	//   Name: "Fernando",
	//   Mail: "fer@fer.com",
	//   Phone: "1357924680",
	// }, 3)
	// handlers.DeleteClient(3)

	handlers.Run()
}
