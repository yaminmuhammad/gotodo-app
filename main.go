package main

import (
	"gotodo-app/delivery"
)

func main() {
	delivery.NewServer().Run()
}

// TODO:
/*
1. Unit testing -> Repository -> (db: mockDb -> DATA DOG)
2. Unit testing -> UseCase
3. Unit testing -> Controller
*/
