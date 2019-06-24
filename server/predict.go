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
	ffmModel, err = core.FfmLoadModel(util.PredictOption.ModelPath)
	if err != nil {
		panic(err)
	}
}

func FfmPredictOne(node model.FfmNode, ffmModel *model.FfmModel) {

}

func FfmPredict(nodes []model.FfmNode, ffmModel *model.FfmModel) {
	var r float32 = 1
	if ffmModel.Normalization {
		r = 0
		for _, node := range nodes {
			r += node.Value * node.Value
		}
		r = 1 / r
	}

}

func wTx(nodes []model.FfmNode, ffmModel *model.FfmModel, r, kappa, eta, lambda float32, doUpdate bool) {
	align0 := 2 * core.GetLatentFactorsNumberAligned(ffmModel.LatentFactorsNumber)
	align1 := ffmModel.FieldsNumber * align0

	var t float32 = 0
	for i, node := range nodes {
		feature1 := node.Feature
		field1 := node.Field
		value1 := node.Value
		if feature1 >= ffmModel.FeaturesNumber || field1 >= ffmModel.FieldsNumber {
			continue
		}
		for j := i; j < len(nodes); j++ {
			feature2 := nodes[j].Feature
			field2 := nodes[j].Field
			value2 := nodes[j].Value
			if feature2 >= ffmModel.FeaturesNumber || field2 >= ffmModel.FieldsNumber {
				continue
			}
			w1 := ffmModel.W[0] + float32(feature1*align1+field2*align0)
			w2 := ffmModel.W[0] + float32(feature2*align1+field1*align0)
			value := value1 * value2 * r

		}
	}

}
