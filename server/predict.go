package server

import (
	"go-ffm/core"
	"go-ffm/model"
	"go-ffm/util"
)

var (
	ffmModel *model.FfmModel
)

func LoadFfmModel() {
	var err error
	ffmModel, err = core.FfmLoadJsonModel(util.PredictOption.ModelPath)
	if err != nil {
		panic(err)
	}
}

func FfmPredictOne(nodes []*model.FfmNode) float64 {
	return core.FfmPredict(nodes, ffmModel)
}
