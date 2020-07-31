package domain

import (
	"encoding/json"
	"fmt"
	"github.com/leopedroso45/UpnidChallenge/controller"
	"io/ioutil"
	"os"
	"strings"
)

type DDD struct {
	StateToDDD map[string]string `json:"estadoPorDDD"`
}

func GetStateByDDD(phone string) (ddd string) {
	ddd = phone[:strings.IndexByte(phone, ' ')]
	return
}

func GetStateMap() map[string]string {
	// Open our jsonFile
	jsonFile, err := os.Open(`../utils/ddd.json`)
	// if we os.Open returns an error then handle it
	controller.CheckError(err)
	fmt.Println("Successfully Opened ddd.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ddd DDD

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &ddd)
	return ddd.StateToDDD
}
