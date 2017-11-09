package gofiles

import (
	"testing"
	"time"
)

func TestAdd2d(t *testing.T) {
	var data2d Data2d
	data2d.Data = make(map[string]map[string]float64)

	//sets date
	data2d.Date = time.Now().Format("2006-01-02")

	out := []string{"test1", "test2"}
	outF := 4.20

	Add2d(data2d.Data, out[0], out[1], outF)

	if data2d.Data[out[0]] == nil {
		t.Fatalf("ERROR, expected: %s%v, got %s", out[1], outF, "nil")
	}

	if data2d.Data[out[0]][out[1]] != outF {
		t.Fatalf("ERROR, expected: %v, got %v\n", outF, data2d.Data[out[0]][out[1]])
	}

}

func TestGetValue(t *testing.T) {
	out := []string{"NOK", "EUR"}
	date := time.Now()
	dateCopy := date.Format("2006-01-02")

	// Set up the database
	db := SetupDB()
	db.Init()

	// Get today's currencies for today
	data2d, ok := db.GetLatest(dateCopy)

	// If there isn't any data in the db
	if ok == false {
		date = date.AddDate(0,0,-1)
		dateCopy = date.Format("2006-01-02")
		data2d, ok = db.GetLatest(dateCopy)
		if ok == false {
			t.Fatalf("ERROR could not retrieve data from db")
		}
	}

	realValue := data2d.Data[out[0]][out[1]]

	testValue := GetValue(out[0], out[1])

	if realValue != testValue {
		t.Fatalf("ERROR, expected %v, got %v\n", realValue, testValue)
	}
}
