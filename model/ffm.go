package model

type FfmNode struct {
	Field   int     // field index
	Feature int     // feature index
	Value   float32 // value
}

type FfmModel struct {
	FeaturesNumber      int32
	FieldsNumber        int32
	LatentFactorsNumber int32
	W                   []float32
	Normalization       bool
}

type FfmParameter struct {
	Eta                 float64 // learning rate
	Lambda              float64 // regularization parameter
	NrIters             int
	LatentFactorsNumber int // number of latent factors
	Normalization       bool
	AutoStop            bool
}
