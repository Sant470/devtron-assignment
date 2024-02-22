package services

import (
	"fmt"
	"log"

	"github.com/sant470/devetron/types/req"
	"github.com/sant470/devetron/types/res"
	"github.com/sant470/devetron/utils/errors"
)

type SearchService struct {
	lgr *log.Logger
}

func NewSearchService(lgr *log.Logger) *SearchService {
	return &SearchService{lgr}
}

func (searchSvc *SearchService) Search(searchReq *req.SearchReq) (*res.SearchResult, *errors.AppError) {
	fmt.Println("got here")
	return nil, nil
}
