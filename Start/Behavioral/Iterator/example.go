package main

import "fmt"

// Iterate using a callback function
func main() {
	// // TODO: use IterateBooks to iterate via a callback function
	// lib.IterateBooks(myBookCallback);

	// // TODO: Use IterateBooks to iterate via anonymous function
	// lib.IterateBooks(func (b Book) error {
	// 	fmt.Println("Book Author:", b.Author);
	// 	return nil;
	// } )

	// TODO: create a BookIterator
	iter := lib.createIterator();
	for iter.hasNext() {
		b := iter.next();
		fmt.Printf("Book %+v\n", b);
	}
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
