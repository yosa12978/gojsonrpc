MAIN_FILE = main.go
OUT_FILE = server.exe

run:
	go run ${MAIN_FILE}

build:
	go build -o ${OUT_FILE} ${MAIN_FILE}