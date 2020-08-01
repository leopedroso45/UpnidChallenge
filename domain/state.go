package domain

import (
	"encoding/json"
	"fmt"
	"github.com/leopedroso45/UpnidChallenge/controller"
	"io/ioutil"
	"os"
	"strings"
)

/*DDD: It's an application structure which has the DDD information for each Brazilian state.*/
type DDD struct {
	DDDToState map[string]string `json:"estadoPorDDD"`
}

/*GetStateByDDD: Method that receives a phone number in string type and returns only the DDD.*/
func GetStateByDDD(phone string) (ddd string) {
	ddd = phone[:strings.IndexByte(phone, ' ')]
	return
}

/*GetStateMap: Method opens a json file containing the DDD and state mapping and returns it.*/
func GetStateMap() map[string]string {
	// Open our jsonFile
	//../utils/ddd.json
	jsonFile, err := os.Open(`/app/utils/ddd.json`)
	// if we os.Open returns an error then handle it
	controller.CheckError(err)
	fmt.Println("Successfully Opened ddd.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ddd DDD

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	_ = json.Unmarshal(byteValue, &ddd)
	return ddd.DDDToState
}
