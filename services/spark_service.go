package services

import (
	"github.com/silencily/sparktime/models"
	"github.com/silencily/sparktime/repositories"
	"io"
	"io/ioutil"
)

type SparkService interface {
	GetImg(id string) (img []byte, contentType string)
	List() []models.Spark
	Save(spark *models.Spark, imgReader io.Reader) error
	Clean() error
}

func NewSparkService() SparkService {
	return &sparkServiceImpl{
		sparkRepository: repositories.NewSparkRepository(),
	}
}

type sparkServiceImpl struct {
	sparkRepository repositories.SparkRepository
}

func (s *sparkServiceImpl) Clean() error {
	query := map[string]interface{}{
		"limit": 100,
	}
	err := s.sparkRepository.Clean(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *sparkServiceImpl) List() []models.Spark {
	query := map[string]interface{}{
		"limit": 30,
	}
	sparks := s.sparkRepository.List(query)
	return sparks
}

func (s *sparkServiceImpl) GetImg(id string) ([]byte, string) {
	attachment := s.sparkRepository.GetImg(id)
	if attachment == nil {
		return nil, ""
	}
	closer := attachment.Content
	defer closer.Close()
	result, err := ioutil.ReadAll(closer)
	if err != nil {
		return nil, ""
	}
	return result, attachment.ContentType
}

func (s *sparkServiceImpl) Save(spark *models.Spark, imgReader io.Reader) error {
	_, err := s.sparkRepository.Save(spark, imgReader)
	if err != nil {
		return err
	}
	return nil
}
