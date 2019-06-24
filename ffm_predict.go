package main

import (
	"go-ffm/server"
	"go-ffm/util"
)

func init() {
	util.ParsePredictOption()
	server.LoadFfmModel()
}

func main() {
	testData := "0:0:1 1:14:1 2:39:1 3:42:1 4:55:1"

}
