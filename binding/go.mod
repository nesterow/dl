module binding

go 1.22.6

replace github.com/nesterow/dal/pkg/adapter => ../pkg/adapter

replace github.com/nesterow/dal/pkg/proto => ../pkg/proto

replace github.com/nesterow/dal/pkg/builder => ../pkg/builder

replace github.com/nesterow/dal/pkg/filters => ../pkg/filters

replace github.com/nesterow/dal/pkg/utils => ../pkg/utils

replace github.com/nesterow/dal/pkg/handler => ../pkg/handler

replace github.com/nesterow/dal/pkg/facade => ../pkg/facade

require (
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/nesterow/dal/pkg/facade v0.0.0-20240820192515-7a408c994181
)

require (
	github.com/nesterow/dal/pkg/adapter v0.0.0-20240820192515-7a408c994181 // indirect
	github.com/nesterow/dal/pkg/builder v0.0.0-20240820192515-7a408c994181 // indirect
	github.com/nesterow/dal/pkg/filters v0.0.0-20240820192515-7a408c994181 // indirect
	github.com/nesterow/dal/pkg/handler v0.0.0-20240820192515-7a408c994181 // indirect
	github.com/nesterow/dal/pkg/proto v0.0.0-20240820192515-7a408c994181 // indirect
	github.com/nesterow/dal/pkg/utils v0.0.0-20240820192515-7a408c994181 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
)
