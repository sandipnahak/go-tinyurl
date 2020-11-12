package main

func main() {
	db := getDB()
	readALL(db)
	db.Close()

	// fmt.Println(db.String())

	// u := URL{short: "go", long: "http://google.com/"}

	// err := dbUpdate(db, &u)
	// if err != nil {
	// 	fmt.Println("Unable to update the db")
	// } else {
	// 	fmt.Println("Updated")
	// }

	httpServer()
}
