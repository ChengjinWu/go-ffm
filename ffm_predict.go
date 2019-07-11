package main

import (
	"fmt"
	"go-ffm/model"
	"go-ffm/server"
	"go-ffm/util"
)

func init() {
	util.ParsePredictOption()
	server.LoadFfmModel()
}

func main() {
	nodes := []*model.FfmNode{
		&model.FfmNode{
			Field:   0,
			Feature: 0,
			Value:   1,
		},
		&model.FfmNode{
			Field:   1,
			Feature: 14,
			Value:   1,
		},
		&model.FfmNode{
			Field:   2,
			Feature: 39,
			Value:   1,
		},
		&model.FfmNode{
			Field:   3,
			Feature: 42,
			Value:   1,
		},
		&model.FfmNode{
			Field:   4,
			Feature: 55,
			Value:   1,
		},
	}
	ctr := server.FfmPredictOne(nodes)
	fmt.Println(ctr)
}
