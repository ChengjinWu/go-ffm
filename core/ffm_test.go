package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
	"go-ffm/model"
	"math"
	"os"
	"testing"
	"unsafe"
)

func GetLatentFactorsNumberAlignedJson(latentFactorsNumber int) int {
	return int(math.Ceil(float64(latentFactorsNumber)/kALIGN) * kALIGN)
}

func GetWSizeJson(ffmModel *model.FfmModel) int {
	latentFactorsNumberAligned := GetLatentFactorsNumberAligned(ffmModel.LatentFactorsNumber)
	return ffmModel.FeaturesNumber * ffmModel.FieldsNumber * latentFactorsNumberAligned * 2
}

type FfmModelJson struct {
	FeaturesNumber      int32
	FieldsNumber        int32
	LatentFactorsNumber int32
	W                   []float32
	Normalization       bool
}

func FfmLoadModelToJson(modelPath string) (*FfmModelJson, error) {
	ffmModel := FfmModelJson{}
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
func TestModelParse(t *testing.T) {
	FfmLoadModelToJson("../data/train.log.model")
}

func TestSizeof(t *testing.T) {
	a := "abc"
	b := len(a)
	c := unsafe.Sizeof(a)
	fmt.Println(int(c))
	fmt.Println(a, b, c)
}
