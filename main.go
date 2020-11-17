package main

func main() {
	db := getDB()
	readALL(db)
	db.Close()

	// fmt.Println(db.String())

	// u := URL{short: "aws", long: "https://aws.amazon.com"}

	// err := dbUpdate(db, &u)
	// if err != nil {
	// 	fmt.Println("Unable to update the db")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Updated")
	// }

	httpServer()
}
