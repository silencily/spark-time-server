package models

import (
	"fmt"
	"strings"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	dataStr := strings.Replace(string(data), "\"", "", -1)
	stamp, err := time.Parse("2006-01-02 15:04:05", dataStr)
	if err != nil {
		return err
	}
	*t = Time(stamp)
	return nil
}

type Spark struct {
	ID          string `json:"_id"`
	Version     string `json:"_rev"`
	Content     string `json:"content"`
	CreatedTime Time   `json:"created_time"`
	ImgName     string `json:"img_name"`
}
