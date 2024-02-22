package req

import "time"

type SearchReq struct {
	SearchKeyword string        `json:"search_keyword"`
	From          time.Duration `json:"from"`
	To            time.Duration `json:"to"`
}
