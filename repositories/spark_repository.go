package repositories

import (
	"context"
	"github.com/go-kivik/kivik"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/models"
)

type SparkRepository interface {
	//根据id获取spark
	Get(id string) *models.Spark
	GetImg(id string) *kivik.Attachment
}

type sparkCouchDBRepository struct {
	template *kivik.DB
}

func NewSparkRepository() SparkRepository {

	return &sparkCouchDBRepository{template: core.GetCouchDBTemplate()}

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
