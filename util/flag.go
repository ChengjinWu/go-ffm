package util

import (
	"flag"
	"go-ffm/model"
	"path"
)

type trainOption struct {
	TrainPath string
	VarPath   string
	ModelPath string
	Param     model.FfmParameter
	Quiet     bool
	NrThreads int
}

type predictOption struct {
	TestPath   string
	ModelPath  string
	OutputPath string
}

var (
	TrainOption   trainOption
	PredictOption predictOption
)

func init() {
	TrainOption = trainOption{
		Param: model.FfmParameter{
			Eta:                 0.2,
			Lambda:              0.00002,
			NrIters:             15,
			LatentFactorsNumber: 4,
			Normalization:       true,
			AutoStop:            false,
		},
		NrThreads: 1,
	}

	PredictOption = predictOption{}
}
func ParseTrainOption() {
	flag.StringVar(&TrainOption.TrainPath, "tp", "", "need to specify train path after -p")
	flag.StringVar(&TrainOption.VarPath, "vp", "", "need to specify var path after -p")
	flag.StringVar(&TrainOption.ModelPath, "mp", "", "need to specify train path after -p")
	flag.Float64Var(&TrainOption.Param.Eta, "r", 0.2, "need to specify eta after -r")
	flag.Float64Var(&TrainOption.Param.Lambda, "l", 0.00002, "need to specify lambda after -l")
	flag.IntVar(&TrainOption.Param.NrIters, "t", 15, "need to specify number of iterations after -t")
	flag.IntVar(&TrainOption.Param.LatentFactorsNumber, "k", 4, "need to specify number of factors after -k")
	flag.BoolVar(&TrainOption.Param.Normalization, "norm", true, "")
	flag.BoolVar(&TrainOption.Param.AutoStop, "auto-stop", false, "")
	flag.BoolVar(&TrainOption.Quiet, "quiet", false, "")
	flag.IntVar(&TrainOption.NrThreads, "s", 1, "need to specify number of threads after -s")
	flag.Parse()
	if len(TrainOption.VarPath) == 0 || len(TrainOption.TrainPath) == 0 {
		panic("var path or train path must specify")
	}
	if len(TrainOption.ModelPath) == 0 {
		trainFileNameWithSuffix := path.Base(TrainOption.TrainPath)
		TrainOption.ModelPath = trainFileNameWithSuffix + ".model"
	}
}

func ParsePredictOption() {
	flag.StringVar(&PredictOption.TestPath, "tp", "", "test path")
	flag.StringVar(&PredictOption.ModelPath, "mp", "", "model path")
	flag.StringVar(&PredictOption.OutputPath, "op", "", "output path")
	flag.Parse()
}
