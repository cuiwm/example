package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Args struct {
	Time time.Time
}

var sList = []string{
	`{"Time": "2018-09-15T11:38:27.0Z"}`,   //ok
	`{"Time": "2018-10-15T11:38:00Z"}`,     //ok
	`{"Time": "2018-10-15T11:38:27.000Z"}`, //ok
	`{"Time": "2018-10-15T11:38:27Z"}`,     //ok
	`{"Time": "2018-10-15T11:38:27"}`,      //  parsing time ""2018-10-15T11:38:27"" as ""2006-01-02T15:04:05Z07:00"": cannot parse """ as "Z07:00"
	`{"Time": null}`,                       //ok
	`{"Time": ""}`,                         // parsing time """" as ""2006-01-02T15:04:05Z07:00"": cannot parse """ as "2006"
}

func main() {
	for i, s := range sList {
		a := Args{}
		err := json.Unmarshal([]byte(s), &a)
		fmt.Println(i, a.Time.String(), " err: ", err)
	}
}
