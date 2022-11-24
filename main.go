package main

func main() {
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
