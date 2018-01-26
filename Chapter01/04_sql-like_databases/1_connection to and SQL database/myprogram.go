package main

import (
	"database/sql"
	"log"
	"os"

	// 포스트그레스 sql 을 database/sql 을 사용해서
	// 접속하게 해주는 도구가 pq
	_ "github.com/lib/pq"
)

func main() {

	// 시스템 환경 변수에 정의해 좋은 포스트그레스 접속 URL 을 읽어온다.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	/*

	postgres 접속 방법

	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	db, err := sql.Open("postgres", "dbname=dat_test user=dat password=!test host=localhost sslmode=disable")
	db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")

	https://stackoverflow.com/a/32710012/5443084

	변수를 설정해 놓고, Sprintf로 처리하는 예
	https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html
	*/

	// Open a database value.  Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sql.Open() 는 단지 나중에 접속할 수 있도록 준비를 할 뿐이고, 접속을 만들지는 않는다.
	// 정말 접속이 되는지 확인을 하려면 db.Ping() 으로 핑을 날려봐야 함
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
