package main

import (
	"fmt"
	"time"

	// go get github.com/patrickmn/go-cache
	cache "github.com/patrickmn/go-cache"
)

func main() {

	// 기본 만료 4분 짜리 캐쉬를 만들고, 만료된 항목은 30초마다 퍼지하도록 설정
	c := cache.New(5*time.Minute, 30*time.Second)

	// 캐쉬에 키와 값을 넣는다.
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	// 건전성 체크 후 키와 값을 스탠다드아웃풋으로 출력한다.
	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}
}
