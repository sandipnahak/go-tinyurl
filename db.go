package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var bucket = []byte("ShortBucket")

// URL struct
type URL struct {
	short string
	long  string
}

func getDB() *bolt.DB {

	db, err := bolt.Open("urlshort.db", 0600, nil)

	if err != nil {
		return nil
	}

	return db
}

func dbUpdate(db *bolt.DB, u *URL) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return b.Put([]byte(u.short), []byte(u.long))

	})
	return err
}

func readALL(db *bolt.DB) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	return err
}

func getLongURL(db *bolt.DB, short string) (string, error) {
	v := []byte{}
	err := db.View(func(tx *bolt.Tx) error {
		v = tx.Bucket(bucket).Get([]byte(short))
		return nil
	})
	return string(v), err
}
