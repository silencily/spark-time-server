package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSpark(t *testing.T) {
	spark := &Spark{
		ID:          "1",
		Content:     "foo",
		CreatedTime: Time(time.Now()),
		ImgName:     "hello.jpg",
	}

	spark_json, err := json.Marshal(spark)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(string(spark_json))

	var spark2 Spark

	err2 := json.Unmarshal(spark_json, &spark2)
	if err2 != nil {
		t.Log(err2)
		t.Fail()
	}
	fmt.Println(time.Time(spark2.CreatedTime).Format("2006-01-02 15:04:05"))

}
