package main

func main() {
	CreateAdminsTable()
	CreatePlayersTable()
	CreateAdmin("elliot", "bug")
	CreatePlayer("victor", "password")
	ReadAdminsTable()
	ServeWebApp()
}
