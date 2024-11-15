module github.com/joaogabriel01/gochestra/example1

go 1.21.6

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/joaogabriel01/gochestra v0.0.0-20240901201103-a4d2c067e515
	github.com/lib/pq v1.10.9
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

replace github.com/joaogabriel01/gochestra => ../..
