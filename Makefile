.PHONY: app tidy lampp klampp

app:
	@clear && go run src/cmd/main.go

tidy:
	@clear && go mod tidy

lampp:
	@clear && sudo /opt/lampp/lampp start

klampp:
	@clear && sudo /opt/lampp/lampp stop

