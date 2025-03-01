//Without Using Goroutine Channel:
//------------------------------------

package main

import (
	"fmt"
	"time"
)

// Cleaning steps for different areas
func cleanRoom(name string) {
	fmt.Printf("Starting to clean %s...\n", name)
	time.Sleep(2 * time.Second) // Simulating cleaning time
	fmt.Printf("Sweeping %s...\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Mopping %s...\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Dusting %s...\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s is now shining clean! ✨\n\n", name)
}

func main() {
	fmt.Println("🏠 Starting the house cleaning process...\n")

	cleanRoom("Room 1")
	cleanRoom("Room 2")
	cleanRoom("Hall")

	fmt.Println("🏡 The entire house is now spotless and shining! ✨")
}
