package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
	"go-ffm/model"
	"os"
)

const (
	kALIGNByte   = 4
	kALIGN       = kALIGNByte / 4
	kCHUNK_SIZE  = 10000000
	kMaxLineSize = 100000
)

func GetLatentFactorsNumberAligned(latentFactorsNumber int32) int32 {
	return int32(float32((latentFactorsNumber)/kALIGN) * kALIGN)
}

func GetWSize(ffmModel *model.FfmModel) int {
	latentFactorsNumberAligned := GetLatentFactorsNumberAligned(ffmModel.LatentFactorsNumber)
	return int(ffmModel.FeaturesNumber * ffmModel.FieldsNumber * latentFactorsNumberAligned * 2)
}

func FfmReadProblemToDisk(txtPath, binPath string) {

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
	ffmModel.W = make([]float32, wSize)
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
