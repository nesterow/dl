module pkg/adapter

go 1.22.6

require pkg/utils v0.0.0-00010101000000-000000000000

replace pkg/utils => ../utils

require github.com/pkg/errors v0.9.1 // indirect
