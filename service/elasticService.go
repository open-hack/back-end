package service

import (
	"context"
	"errors"

	"github.com/open-hack/back-end/model"
	"gopkg.in/olivere/elastic.v6"
)

const (
	indexName = "chat"
	docType   = "log"
	appName   = "back-end"
)

type ElasticService struct {
	ElasticCLI *elastic.Client
}

func (e *ElasticService) SaveToElastic(ctx context.Context, payload model.Message) error {

	exists, err := e.ElasticCLI.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		res, error := e.ElasticCLI.CreateIndex(indexName).Do(ctx)
		if error != nil {
			return error
		}
		if !res.Acknowledged {
			return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
		}
	}

	var user = model.User{
		Name:         payload.User.Name,
		ProfileImage: payload.User.ProfileImage,
	}

	var message = model.Message{
		Text:      payload.Text,
		CreatedAt: payload.CreatedAt,
		User:      user,
	}

	_, error := e.ElasticCLI.Index().
		Index(indexName).
		Type(docType).
		BodyJson(message).
		Do(ctx)

	if error != nil {
		return error
	}

	return nil
}
