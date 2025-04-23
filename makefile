TRAITS_VERSION = 0.0.8

updatev:
		git tag v${TRAITS_VERSION} && git push origin v${TRAITS_VERSION}



sand:
	go run -v ./cmd/sandbox/main.go