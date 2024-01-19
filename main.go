package main

import "gin-gorm-oj/routes"

func main() {
	r := routes.Router()
	r.Run()
}
