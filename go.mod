module github.com/gertcuykens/main

go 1.14

require (
	github.com/gertcuykens/main/hello v0.0.0-00010101000000-000000000000
	github.com/gertcuykens/module/v5 v5.0.0
)

replace github.com/gertcuykens/main/hello => ./hello
