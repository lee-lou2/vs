package main

import "vs/contents/request_rest_api/apps/golang"

func main() {
	go golang.MainFiber()
	golang.MainGin()
}
