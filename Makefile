#!make

BINARY_NAME = ynab4importer
OUT_DIR = dist

#include ${APP_CONFIG}
#export

build:
	go build -o ${OUT_DIR}/${BINARY_NAME}.exe main.go

# make run in=SK-spends-dec2023.xml out=SK-spends-dec2023.csv
run:
	${OUT_DIR}/${BINARY_NAME} ${OUT_DIR}/$(in) ${OUT_DIR}/$(out)

# test: 
# 	go test -v -cover -timeout 30s jirareports/internal/epic
# 	go test -v -cover -timeout 30s jirareports/internal/uiconfig

# bench:
# 	go test -bench ^BenchmarkTestFromJiraCsvLine$ jirareports/internal/epic -timeout 2h -benchmem

