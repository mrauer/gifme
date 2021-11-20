build:
	docker build -t gifme:latest .

dev:
	docker build -t gifme:latest . && docker run -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/gifme gifme:latest && docker exec -it gifme:latest

binary:
	env GOOS=linux GOARCH=amd64 go build -o gifme
