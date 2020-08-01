package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

var jsonContent = `{
  "estadoPorDdd": {
    "11": "SP",
    "12": "SP",
    "13": "SP",
    "14": "SP",
    "15": "SP",
    "16": "SP",
    "17": "SP",
    "18": "SP",
    "19": "SP",
    "21": "RJ",
    "22": "RJ",
    "24": "RJ",
    "27": "ES",
    "28": "ES",
    "31": "MG",
    "32": "MG",
    "33": "MG",
    "34": "MG",
    "35": "MG",
    "37": "MG",
    "38": "MG",
    "41": "PR",
    "42": "PR",
    "43": "PR",
    "44": "PR",
    "45": "PR",
    "46": "PR",
    "47": "SC",
    "48": "SC",
    "49": "SC",
    "51": "RS",
    "53": "RS",
    "54": "RS",
    "55": "RS",
    "61": "DF",
    "62": "GO",
    "63": "TO",
    "64": "GO",
    "65": "MT",
    "66": "MT",
    "67": "MS",
    "68": "AC",
    "69": "RO",
    "71": "BA",
    "73": "BA",
    "74": "BA",
    "75": "BA",
    "77": "BA",
    "79": "SE",
    "81": "PE",
    "82": "AL",
    "83": "PB",
    "84": "RN",
    "85": "CE",
    "86": "PI",
    "87": "PE",
    "88": "CE",
    "89": "PI",
    "91": "PA",
    "92": "AM",
    "93": "PA",
    "94": "PA",
    "95": "RR",
    "96": "AP",
    "97": "AM",
    "98": "MA",
    "99": "MA"
  }
}`

//TODO Finish TestGetStateByDDD
func TestGetStateByDDD(t *testing.T) {
	phoneList := []string{
		"51 9999999",
		"41 9999999",
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantDdd string
	}{
		//test cases.
		{
			name: "Returns correct string according to the passed phone number.",
			args: args{
				phone: phoneList[0],
			},
			wantDdd: "51",
		},
		{
			name: "Returns correct string according to the passed phone number. //same case",
			args: args{
				phone: phoneList[1],
			},
			wantDdd: "41",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDdd := GetStateByDDD(tt.args.phone); gotDdd != tt.wantDdd {
				t.Errorf("GetStateByDDD() = %v, want %v", gotDdd, tt.wantDdd)
			}
		})
	}
}

//TODO Finish TestGetStateMap
func TestGetStateMap(t *testing.T) {
	var ddd DDD
	_ = json.Unmarshal([]byte(jsonContent), &ddd)

	tests := []struct {
		name string
		want map[string]string
	}{
		//test cases.
		{
			name: "Checks the structure of the Json object.",
			want: ddd.DDDToState,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStateMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStateMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
