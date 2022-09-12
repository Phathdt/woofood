package sdkcm

import (
	"strings"
)

type OrderBy struct {
	Key    string
	IsDesc bool
}

type Paging struct {
	Cursor      uint32    `json:"-" form:"-"`
	NextCursor  string    `json:"next_cursor" form:"-"`
	CursorStr   string    `json:"cursor" form:"cursor"`
	Limit       int       `json:"limit" form:"limit"`
	Total       int       `json:"total" form:"-"`
	Page        int       `json:"page" form:"page"`
	HasNext     bool      `json:"has_next" form:"-"`
	OrderBy     string    `json:"-" form:"-"`
	OB          []OrderBy `json:"-" form:"-"`
	CursorIsUID bool      `json:"-" form:"-"`
}

func (p *Paging) FullFill() {
	if p.Limit <= 0 {
		p.Limit = 25
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	if strings.TrimSpace(p.OrderBy) == "" {
		p.OrderBy = "id desc"
		p.OB = []OrderBy{{Key: "id", IsDesc: true}}
	} else {
		p.OB = getOrderBy(p.OrderBy)
	}
}

func getOrderBy(ord string) []OrderBy {
	comps := strings.Split(ord, ",")
	result := make([]OrderBy, len(comps))

	for i := range comps {
		kvs := strings.Split(strings.TrimSpace(comps[i]), " ")
		result[i] = OrderBy{Key: kvs[0], IsDesc: len(kvs) == 1 || kvs[1] == "-1"}
	}

	return result
}
