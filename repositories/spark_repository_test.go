package repositories

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSparkCouchDBRepository_Get(t *testing.T) {
	rep := NewSparkRepository()
	id := "c9c09ef539dac8d52531a77d6100705e"
	spark := rep.Get(id)
	data, _ := json.Marshal(spark)
	fmt.Println(string(data))
}
