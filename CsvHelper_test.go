package utility

import (
	"io"
	"os"
	"testing"
)

func Test_CsvHelper1_1(t *testing.T) {
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

func Test_CsvHelper1_2(t *testing.T) {
	if test1Csv, openErr := os.Open("./TestData/test1.csv"); openErr == nil {
		defer test1Csv.Close()

		type Recvr struct {
			Value1 string `csv:"to"`
			Value2 string `csv:"cc"`
			Value3 string `csv:"bcc"`
			Value4 string `csv:"template"`
			Value5 string `csv:"params"`
		}

		csvHelper := NewCsvReader(test1Csv, true)
		for {
			recvr := Recvr{}
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

func Test_CsvHelper1_3(t *testing.T) {
	if test1Csv, openErr := os.Open("./TestData/test1.csv"); openErr == nil {
		defer test1Csv.Close()

		csvHelper := NewCsvReader(test1Csv, true)
		for {
			recvr := []interface{}{}
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

func Test_CsvHelper2_1(t *testing.T) {
	if testCsv, openErr := os.Open("./TestData/test2.csv"); openErr == nil {
		defer testCsv.Close()

		csvHelper := NewCsvReader(testCsv, false)
		for {
			recvr := map[int]interface{}{}
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

func Test_CsvHelper2_2(t *testing.T) {
	if test1Csv, openErr := os.Open("./TestData/test2.csv"); openErr == nil {
		defer test1Csv.Close()

		type Recvr struct {
			Value1 string `csv:"to"`
			Value2 string `csv:"cc"`
			Value3 string `csv:"bcc"`
			Value4 string `csv:"template"`
			Value5 string `csv:"params"`
		}

		csvHelper := NewCsvReader(test1Csv, false)
		for {
			recvr := Recvr{}
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

func Test_CsvHelper2_3(t *testing.T) {
	if test1Csv, openErr := os.Open("./TestData/test2.csv"); openErr == nil {
		defer test1Csv.Close()

		csvHelper := NewCsvReader(test1Csv, false)
		for {
			recvr := []interface{}{}
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
