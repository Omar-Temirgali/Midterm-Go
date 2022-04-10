package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/Omar-Temirgali/go-service/models"
)

var (
	mu    sync.Mutex
	Kvmap = make(map[string]string)
)

func jsonToMap() {

	jsonFile, err := os.Open("C:/MyFiles/SDU/3rd COURSE/6th SEMESTER/The Go programming language [INF 368]/go-service/repository/colors.json")

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &Kvmap)
}

func init() {
	RepoCreate(models.KVmap{Key: "apr", Value: "April"})
	RepoCreate(models.KVmap{Key: "mar", Value: "March"})
	RepoCreate(models.KVmap{Key: "may", Value: "May"})
	RepoCreate(models.KVmap{Key: "jan", Value: "January"})
	RepoCreate(models.KVmap{Key: "feb", Value: "February"})
	RepoCreate(models.KVmap{Key: "jun", Value: "June"})
	RepoCreate(models.KVmap{Key: "jul", Value: "July"})
	RepoCreate(models.KVmap{Key: "sep", Value: "September"})
	// jsonToMap()
}

func RepoFind(key string) models.KVmap {
	mu.Lock()
	defer mu.Unlock()
	if value, ok := Kvmap[key]; ok {
		return models.KVmap{Key: key, Value: value}
	}
	return models.KVmap{}
}

func RepoCreate(p models.KVmap) models.KVmap {
	mu.Lock()
	defer mu.Unlock()
	Kvmap[p.Key] = p.Value
	return p
}

func RepoUpdateAndInsert(p models.KVmap) models.KVmap {
	mu.Lock()
	defer mu.Unlock()
	Kvmap[p.Key] = p.Value
	return models.KVmap{Key: p.Key, Value: p.Value}
}
