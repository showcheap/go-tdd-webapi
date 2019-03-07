package main

func main() {
	a := App{}

	a.Initialize("prod.db")
	defer a.DB.Close()

	a.Run(":8080")
}
