package facade

import (
	"reflect"

	"github.com/nesterow/dal/pkg/adapter"
	"github.com/nesterow/dal/pkg/proto"
)

var db adapter.DBAdapter

func InitSQLite(pragmas []string) {
	if db.Type == "" {
		adapter.RegisterDialect("sqlite3", adapter.CommonDialect{})
		db = adapter.DBAdapter{
			Type: "sqlite3",
		}
		db.AfterOpen("PRAGMA journal_mode=WAL")
	}
	for _, pragma := range pragmas {
		if pragma == "" {
			continue
		}
		db.AfterOpen(pragma)
	}
}

func HandleQuery(input *[]byte, output *[]byte) int {
	InitSQLite([]string{})
	req := proto.Request{}
	_, err := req.UnmarshalMsg(*input)
	if err != nil {
		res := proto.Response{
			Id:           0,
			RowsAffected: -1,
			LastInsertId: -1,
			Msg:          "failed to unmarshal request",
		}
		*output, _ = res.MarshalMsg(nil)
		return 0
	}
	query, err := req.Parse(adapter.GetDialect(db.Type))
	if err != nil {
		res := proto.Response{
			Id:           0,
			RowsAffected: -1,
			LastInsertId: -1,
			Msg:          err.Error(),
		}
		*output, _ = res.MarshalMsg(nil)
		return 0
	}
	if query.Exec {
		result, err := db.Exec(query)
		if err != nil {
			res := proto.Response{
				Id:           0,
				RowsAffected: -1,
				LastInsertId: -1,
				Msg:          err.Error(),
			}
			*output, _ = res.MarshalMsg(nil)
			return 0
		}
		ra, _ := result.RowsAffected()
		la, _ := result.LastInsertId()
		res := proto.Response{
			Id:           0,
			RowsAffected: ra,
			LastInsertId: la,
		}
		*output, _ = res.MarshalMsg(nil)
		return 0
	}
	rows, err := db.Query(query)
	if err != nil {
		res := proto.Response{
			Id:           0,
			RowsAffected: -1,
			LastInsertId: -1,
			Msg:          err.Error(),
		}
		*output, _ = res.MarshalMsg(nil)
		return 0
	}
	columns, _ := rows.Columns()
	types, _ := rows.ColumnTypes()
	cols, _ := proto.MarshalRow(columns)
	*output = append(*output, cols...)
	for rows.Next() {
		data := make([]interface{}, len(columns))
		for i := range data {
			typ := reflect.New(types[i].ScanType()).Interface()
			data[i] = &typ
		}
		rows.Scan(data...)
		cols, _ := proto.MarshalRow(data)
		*output = append(*output, cols...)
	}
	return 0
}
