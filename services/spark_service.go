package services

import (
	"github.com/silencily/sparktime/repositories"
	"io/ioutil"
)

type SparkService interface {
	GetImg(id string) (img []byte, contentType string)
}

type sparkServiceImpl struct {
	sparkRepository repositories.SparkRepository
}

func NewSparkService() SparkService {
	return &sparkServiceImpl{
		sparkRepository: repositories.NewSparkRepository(),
	}
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
