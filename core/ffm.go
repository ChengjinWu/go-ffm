package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
	"go-ffm/model"
	"io/ioutil"
	"math"
	"os"
)

const (
	kALIGNByte   = 4
	kALIGN       = kALIGNByte / 4
	kCHUNK_SIZE  = 10000000
	kMaxLineSize = 100000
)

func GetLatentFactorsNumberAligned(latentFactorsNumber int) int {
	return int(math.Ceil(float64(latentFactorsNumber)/kALIGN) * kALIGN)
}

func GetWSize(ffmModel *model.FfmModel) int {
	latentFactorsNumberAligned := GetLatentFactorsNumberAligned(ffmModel.LatentFactorsNumber)
	return ffmModel.FeaturesNumber * ffmModel.FieldsNumber * latentFactorsNumberAligned * 2
}

func FfmReadProblemToDisk(txtPath, binPath string) {

}

func FfmLoadJsonModel(modelPath string) (*model.FfmModel, error) {
	ffmModelFile, err := os.Open(modelPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	data, err := ioutil.ReadAll(ffmModelFile)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	ffmModel := model.FfmModel{}
	err = json.Unmarshal(data, &ffmModel)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &ffmModel, nil
}
func FfmLoadModel(modelPath string) (*model.FfmModel, error) {
	ffmModel := model.FfmModel{}
	ffmModelFile, err := os.Open(modelPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = binary.Read(ffmModelFile, binary.LittleEndian, &ffmModel.FeaturesNumber)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = binary.Read(ffmModelFile, binary.LittleEndian, &ffmModel.FieldsNumber)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = binary.Read(ffmModelFile, binary.LittleEndian, &ffmModel.LatentFactorsNumber)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = binary.Read(ffmModelFile, binary.LittleEndian, &ffmModel.Normalization)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	wSize := GetWSize(&ffmModel)
	ffmModel.W = make([]float64, wSize)
	fmt.Println(wSize)
	for offset := 0; offset < wSize; {
		nextOffset := min(wSize, offset+4*kCHUNK_SIZE)
		currW := ffmModel.W[offset:nextOffset]
		err = binary.Read(ffmModelFile, binary.LittleEndian, &currW)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		offset = nextOffset
	}

	jb, _ := json.Marshal(ffmModel)
	fmt.Println(string(jb))
	return nil, nil
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func wTx(nodes []*model.FfmNode, ffmModel *model.FfmModel, r, kappa, eta, lambda float64, doUpdate bool) float64 {
	align0 := 2 * GetLatentFactorsNumberAligned(ffmModel.LatentFactorsNumber)
	align1 := ffmModel.FieldsNumber * align0

	var t float64 = 0
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
			w1Index := feature1*align1 + field2*align0
			w2Index := feature2*align1 + field1*align0
			value := value1 * value2 * r
			if doUpdate {
				wg1Index := w1Index + kALIGN
				wg2Index := w2Index + kALIGN
				var d int = 0
				for ; d < align0; d += kALIGN * 2 {
					g1 := lambda*ffmModel.W[w1Index+d] + kappa*ffmModel.W[w2Index+d]*value
					g2 := lambda*ffmModel.W[w2Index+d] + kappa*ffmModel.W[w1Index+d]*value
					ffmModel.W[wg1Index+d] += g1 * g1
					ffmModel.W[wg2Index+d] += g2 * g2
					ffmModel.W[w1Index+d] -= eta / math.Sqrt(ffmModel.W[wg1Index+d]) * g1
					ffmModel.W[w2Index+d] -= eta / math.Sqrt(ffmModel.W[wg2Index+d]) * g2
				}
			} else {
				for d := 0; d < align0; d += kALIGN * 2 {
					t += ffmModel.W[w1Index+d] * ffmModel.W[w2Index+d] * value
				}
			}
		}
	}
	return t
}

func FfmPredict(nodes []*model.FfmNode, ffmModel *model.FfmModel) float64 {
	var r float64 = 1
	if ffmModel.Normalization {
		r = 0
		for _, node := range nodes {
			r += node.Value * node.Value
		}
		r = 1 / r
	}
	t := wTx(nodes, ffmModel, r, 0, 0, 0, false)
	return 1 / (1 + math.Exp(t))
}
