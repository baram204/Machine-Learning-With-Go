package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the CSV file.
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Create a histogram for each of the feature columns in the dataset.
	for _, colName := range irisDF.Names() {

		// 열 이름이 species 가 아니면 히스토그램을 만든다.
		if colName != "species" {

			// Create a plotter.Values value and fill it with the
			// values from the respective column of the dataframe.
			//  플로터 밸류를 데이터프레임에 있는 총 행 수로 채운다.
			v := make(plotter.Values, irisDF.Nrow())
			// 행이름의 각 행을 float 타입으로 긁어서 플롯 밸류에 채운다.
			// 위에서 플로터 밸류의 크기가 데이터 프레임의 행의 개수만큼이므로..
			// 열의 모든 원소를 채워 넣으면 딱 맞게 된다.
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// 플롯을 만들고 제목을 정한다.
			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			// 정규 분포로부터 drawn 한 우리의 값의 히스토그램을 생성한다.
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// 히스토그램을 정규화(normalize)한다.
			h.Normalize(1)

			// 플롯에 히스토그램을 더한다.
			p.Add(h)

			// 플롯을 PNG 파일로 저장한다..
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
