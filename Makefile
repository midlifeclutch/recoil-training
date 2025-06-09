.DEFAULT_GOAL := all

ak_chest:
	@GOTMPDIR=/var/tmp go run main.go

ak_head:
	@GOTMPDIR=/var/tmp go run main.go -t head -g 55

all:
	@GOTMPDIR=/var/tmp go run main.go -s 2 -r 15 -w ak47
