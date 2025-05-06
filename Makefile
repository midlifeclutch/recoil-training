.DEFAULT_GOAL := all

all:
	@GOTMPDIR=/var/tmp go run main.go
