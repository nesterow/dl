module l12.xyz/dal

go 1.22.6

require l12.xyz/dal/utils v0.0.0 // indirect

replace l12.xyz/dal/utils v0.0.0 => ../utils

require l12.xyz/dal/filters v0.0.0

require github.com/pkg/errors v0.9.1 // indirect

replace l12.xyz/dal/filters v0.0.0 => ../filters
