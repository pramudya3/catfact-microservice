package service

import (
	"context"
	"encoding/json"
	"go-microservice/model"
	"net/http"
)

type CatFactService interface {
	GetCatFact(ctx context.Context) (*model.CatFact, error)
}

type CatFact struct {
	url string
}

func NewCatFactService(url string) CatFactService {
	return &CatFact{
		url: url,
	}
}

func (c *CatFact) GetCatFact(ctx context.Context) (*model.CatFact, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fact := &model.CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
