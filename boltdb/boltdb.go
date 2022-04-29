package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	updateErr := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = b.Put([]byte("answer"), []byte("42"))
		return err
	})
	fmt.Println(updateErr)
	tx, _ := db.Begin(false)
	val := tx.Bucket([]byte("MyBucket")).Get([]byte("answer"))

	bucket1 := tx.Bucket([]byte("test"))
	bucket2 := tx.Bucket([]byte("MyBucket"))
	fmt.Println(string(val), bucket1, bucket2)
}
