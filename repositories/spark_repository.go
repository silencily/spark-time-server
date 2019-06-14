package repositories

import (
	"context"
	"github.com/go-kivik/kivik"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/models"
)

const (
	SPARK_COUCHDB_VIEW string = "spark"
	SPARK_COUCHDB_DDOC string = "spark"
)

type SparkRepository interface {
	//根据id获取spark
	Get(id string) *models.Spark
	GetImg(id string) *kivik.Attachment
	List(query map[string]interface{}) []models.Spark
}

func NewSparkRepository() SparkRepository {
	return &sparkCouchDBRepository{template: core.GetCouchDBTemplate()}
}

type sparkCouchDBRepository struct {
	template *kivik.DB
}

func (rep *sparkCouchDBRepository) List(query map[string]interface{}) []models.Spark {
	rows, err := rep.template.Query(context.TODO(), SPARK_COUCHDB_DDOC, SPARK_COUCHDB_VIEW, query)
	if err != nil {
		return nil
	}
	sparks := make([]models.Spark, 0)
	for hasNext := rows.Next(); hasNext; hasNext = rows.Next() {
		spark := new(models.Spark)
		err := rows.ScanValue(spark)
		if err != nil {
			continue
		}
		sparks = append(sparks, *spark)
	}
	return sparks
}

func (rep *sparkCouchDBRepository) Get(id string) *models.Spark {
	row := rep.template.Get(context.TODO(), id)
	var spark models.Spark
	err := row.ScanDoc(&spark)
	if err != nil {
		return nil
	}
	return &spark
}

func (rep *sparkCouchDBRepository) GetImg(id string) *kivik.Attachment {
	row := rep.template.Get(context.TODO(), id)
	var spark models.Spark
	err := row.ScanDoc(&spark)
	if err != nil {
		return nil
	}
	attachment, err := rep.template.GetAttachment(context.TODO(), id, spark.ImgName)
	if err != nil {
		return nil
	}
	return attachment
}
