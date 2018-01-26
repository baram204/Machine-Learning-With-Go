package main

import (
	// json encoding 을 위한 go 표준 라이브러리
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 도시바이크 정류장 상태를 json 으로 나타낸 주소
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// 정거장자료는 위의 citiBikeURL 부터 반환받은 JSON 문서를 언마샬 하는 것으로 사용될 수 있다.
// stationData is used to unmarshal the JSON document returned form citiBikeURL.
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// 정거장은 정거장 자료 내부에서, 각 정거장 문서를 언마샬 함으로써 사용될 수 있다.
type station struct {

	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

// 1. 고 이디엄을 따르기 위해서, json 에서 밑줄을 사용한 스네이크 케이스를 고의 이미엄(카멜+파스칼 케이스)로 바꿨다.
// 2. json 구조 태그를 예상되는 대응 필드에 붙였다. 그래서 언마샬하는 json 라이브러리가 알아먹을 수 있게.
// 3. JSON 데이터가 제대로 해석되려면 구조체 필드 이름의 첫 글자가 대문자여야 한다.
//    encoding/json 은 필드가 exported(대문화 되어서 외부에서 접근 가능하게 되는 상태)되기 전 까지는 reflect 를 사용해서
//    볼 수가 없기 때문이다.

func main() {

	// URL 로부터 응답을 가져온다.
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}

	// 응답 본문 닫기를 지연걸어놓는다.
	defer response.Body.Close()

	// 응답 본문을  []byte 로 읽는다.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 정거장 자료를 담을 변수를 선언한다.
	var sd stationData

	// JSON 자료를 주어진 변수에 담을 수 있도록 언마샬한다
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}

	// 첫 역을 출력한다.
	fmt.Printf("%+v\n\n", sd.Data.Stations[0])
}
