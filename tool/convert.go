package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
	"math"
	"os"
)

const (
	kALIGNByte   = 4
	kALIGN       = kALIGNByte / 4
	kCHUNK_SIZE  = 10000000
	kMaxLineSize = 100000
)

func GetLatentFactorsNumberAlignedJson(latentFactorsNumber int32) int32 {
	return int32(math.Ceil(float64(latentFactorsNumber)/kALIGN) * kALIGN)
}

func GetWSizeJson(ffmModel *FfmModelJson) int32 {
	latentFactorsNumberAligned := GetLatentFactorsNumberAlignedJson(ffmModel.LatentFactorsNumber)
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
	wSize := GetWSizeJson(&ffmModel)
	ffmModel.W = make([]float32, wSize)
	fmt.Println(wSize)
	var offset int32 = 0
	for offset < wSize {
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

func min(x, y int32) int32 {
	if x > y {
		return y
	}
	return x
}
func main() {
	FfmLoadModelToJson("../data/train.log.model")
}
