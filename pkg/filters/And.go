package filters

import (
	"fmt"
	"strings"
)

type And struct {
	And []string `json:"$and"`
}

func (f And) ToSQLPart(ctx Context) string {
	if f.And == nil {
		return ""
	}
	value := strings.Join(f.And, " AND ")
	return fmt.Sprintf("(%s)", value)
}

func (a And) FromJSON(data interface{}) Filter {
	return FromJson[And](data)
}
