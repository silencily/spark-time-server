package controllers

type SparkController struct {
}

func (c *SparkController) Get() map[string]interface{} {
	result := map[string]interface{}{
		"name": "foo",
		"time": 2,
	}
	return result
}
