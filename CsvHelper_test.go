package utility

import (
	"io"
	"os"
	"testing"
)

func Test_CsvHelper1(t *testing.T) {
	if test1Csv, openErr := os.Open("./TestData/test1.csv"); openErr == nil {
		defer test1Csv.Close()

		csvHelper := NewCsvReader(test1Csv, true)
		for {
			recvr := map[string]interface{}{}
			if retErr := csvHelper.DecodeLine(&recvr); retErr == nil || retErr == io.EOF {
				t.Logf("decoded: %v", recvr)
				if retErr == io.EOF {
					break
				}
			} else {
				t.Fatalf("fail to decode record: %v", retErr)
				break
			}
		}
	} else {
		t.Fatalf("fail to open test data file: %v", openErr)
	}
}
