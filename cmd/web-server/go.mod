module github.com/mitchallen/go-monorepo-demo/cmd/web-server

go 1.21.1

require (
	github.com/mitchallen/go-monorepo-demo/pkg/beta v0.0.0
	github.com/mitchallen/go-monorepo-demo/pkg/shared v0.0.0
)

require (
	github.com/mitchallen/coin v0.3.0 // indirect
	github.com/mitchallen/go-monorepo-demo/pkg/alpha v0.0.0 // indirect
)

replace (
	github.com/mitchallen/go-monorepo-demo/pkg/alpha => ../../pkg/alpha
	github.com/mitchallen/go-monorepo-demo/pkg/beta => ../../pkg/beta
	github.com/mitchallen/go-monorepo-demo/pkg/shared => ../../pkg/shared
)
