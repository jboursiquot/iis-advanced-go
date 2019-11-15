cover:
	go test -coverprofile=coverage.out 
	go tool cover -html=coverage.out

bench:
	go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
	go tool pprof profile.out

pprof:
	pprof -http=:8080 profile.out