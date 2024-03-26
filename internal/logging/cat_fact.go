package logging

import (
	"context"
	"fmt"
	"go-microservice/internal/service"
	"go-microservice/model"
	"time"
)

type CatFactLogging struct {
	// optional adding logger or create by your self logging
	next service.CatFactService
}

func NewCatFactLogging(next service.CatFactService) service.CatFactService {
	return &CatFactLogging{
		next: next,
	}
}

func (c *CatFactLogging) GetCatFact(ctx context.Context) (fact *model.CatFact, err error) {
	defer func(start time.Time) {
		if err != nil {
			fmt.Printf("fact=%s err=%v took%v\n", fact.Fact, err, time.Since(start))
		} else {
			fmt.Printf("fact=%s took%v\n", fact.Fact, time.Since(start))
		}
	}(time.Now())

	return c.next.GetCatFact(ctx)
}
