package main

import (
	"github.com/lunny/log"
	"go-ffm/util"
	"path"
)

func init() {

}

func TrainOnDisk() {
	trainBinPath := path.Base(util.TrainOption.TrainPath) + ".bin"
	varBinPath := ""
	if len(util.TrainOption.VarPath) > 0 {
		varBinPath = path.Base(util.TrainOption.VarPath) + ".bin"

	}
	log.Info(trainBinPath, varBinPath)

}

func main() {
	util.ParseTrainOption()

	if util.TrainOption.Quiet {
		log.SetOutputLevel(log.Lerror)
		log.Info("不打印日志")
	}
	if util.TrainOption.Param.AutoStop && len(util.TrainOption.TrainPath) == 0 {
		log.Info("To use auto-stop, you need to assign a validation set")
	}
}
