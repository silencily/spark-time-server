package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/kataras/golog"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/models"
	"io"
	"strings"
	"time"
)

const (
	SPARK_COUCHDB_VIEW_SPARK string = "spark"
	SPARK_COUCHDB_DDOC_SPARK string = "spark"
	SPARK_COUCHDB_VIEW_DYING string = "dying"
)

func getLogger() *golog.Logger {
	return core.GetLogger("SparkRepository")
}

type SparkRepository interface {
	//根据id获取spark
	Get(id string) *models.Spark
	GetImg(id string) *kivik.Attachment
	List(query map[string]interface{}) []models.Spark
	Save(spark *models.Spark, imgReader io.Reader) (string, error)
	Clean(query map[string]interface{}) error
}

func NewSparkRepository() SparkRepository {
	return &sparkCouchDBRepository{template: core.GetCouchDBTemplate()}
}

type sparkCouchDBRepository struct {
	template *kivik.DB
}

func (rep *sparkCouchDBRepository) Clean(query map[string]interface{}) error {
	errorCount := 0
	rows, err := rep.template.Query(context.TODO(), SPARK_COUCHDB_DDOC_SPARK, SPARK_COUCHDB_VIEW_DYING, query)
	if err != nil {
		errorCount++
		getLogger().Errorf("Query error:%s", err.Error())
		return err
	}
	for hasNext := rows.Next(); hasNext; hasNext = rows.Next() {
		docId := rows.ID()
		var docRev string
		err = rows.ScanValue(&docRev)
		if err != nil {
			errorCount++
			getLogger().Errorf("ScanValue error-[docId:%s]", docId)
			continue
		}
		_, err = rep.template.Delete(context.TODO(), docId, docRev)
		if err != nil {
			errorCount++
			getLogger().Errorf("Delete error-[docId:%s,docRev:%s]", docId, docRev)
			continue
		}
	}
	if errorCount > 0 {
		errMsg := fmt.Sprintf("Clean dying spark occurs %d errors ", errorCount)
		return errors.New(errMsg)
	}
	return nil
}

func (rep *sparkCouchDBRepository) List(query map[string]interface{}) []models.Spark {
	rows, err := rep.template.Query(context.TODO(), SPARK_COUCHDB_DDOC_SPARK, SPARK_COUCHDB_VIEW_SPARK, query)
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

func (rep *sparkCouchDBRepository) Save(spark *models.Spark, imgReader io.Reader) (string, error) {
	var stamp = fmt.Sprintf("%s", time.Time(spark.CreatedTime).Format("2006-01-02 15:04:05"))
	sparkDoc := map[string]interface{}{
		"content":      spark.Content,
		"created_time": stamp,
		"img_name":     spark.ImgName,
	}
	docId, rev, err := rep.template.CreateDoc(context.TODO(), sparkDoc)
	if err != nil {
		return "", err
	}
	var contentType string
	if strings.Contains(spark.ImgName, ".png") {
		contentType = "image/png"
	} else {
		contentType = "image/jpeg"
	}
	attachment, err := couchdb.NewAttachment(spark.ImgName, contentType, imgReader)
	_, err = rep.template.PutAttachment(context.TODO(), docId, rev, attachment)
	if err != nil {
		return "", err
	}
	return docId, nil
}
