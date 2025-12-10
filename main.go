package main

import (
	"fmt"
	"log"
)

func main() {
	// 1. Start the DB
	db, err := NewDB("apex.db")
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}
	defer db.Close()

	fmt.Println("🚀 ApexDB Started...")

	// 2. Write Data
	fmt.Println("Writing data...")
	err = db.Set("username", "parth_bhanti")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Set("role", "systems_engineer")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Read Data
	fmt.Println("Reading data...")
	val, err := db.Get("username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✅ Found Key 'username': %s\n", val)

	val2, err := db.Get("role")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✅ Found Key 'role': %s\n", val2)
}