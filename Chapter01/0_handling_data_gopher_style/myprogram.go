package __handling_data_gopher_style

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	/*
		1. 고에서는 어떤 작업을 시도하면 오류를 꼭 검출한다.
		2. 고에서는 자료의 타입을 명시한다.
	 */

	// CSV 파일을 열어서, 결과와 오류를 받고
	f, err := os.Open("myfile.csv")
	// 오류 검증해서 있으면 로그 출력
	if err != nil {
		log.Fatal(err)
	}

	// CSV 열린 접속으로 리더를 만들고
	r := csv.NewReader(f)
	// 모두 읽기 시도 후 결과와 오류를 받는다.
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//  최대 값을 정수 열에서 가져온다.
	// 최대정수를 저장할 변수 타입을 정의
	var intMax int

	fmt.Println("모든 레코드를 조사하며 큰 수를 찾습니다")
	// for records 를 순회하면서 인덱스를 _, 값을 record 로 받아서
	for _, record := range records {

		// 문자열을 정수로 변환해서 결과와 오류를 얻는다.
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		// 적절한 최대값으로 교체한다.
		if intVal > intMax {
			intMax = intVal
			fmt.Println("더 큰수 발견", intMax)
		}
	}

	// 최대 수를 출력한다.
	fmt.Println("최종 최대 정수는",intMax)
}
