package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {

	// 현재 디렉토리에 embedded.db 자료 파일을 연다.
	// 없으면 생성함
	db, err := bolt.Open("embedded.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 자료를 위한 "bucket"을 볼트DB 파일에 생성한다.
	if err := db.Update(
		func(tx *bolt.Tx) error {
			_, err := tx.CreateBucket([]byte("MyBucket"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		return nil
		});
		err != nil {
		log.Fatal(err)
	}

	// 볼트DB 파일에 키와 값의 맵을 넣는다.
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("mykey"), []byte("myvalue"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	// Output the keys and values in the embedded
	// BoltDB file to standard out.
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
