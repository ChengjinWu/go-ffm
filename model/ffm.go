package model

type FfmNode struct {
	Field   int     // field index
	Feature int     // feature index
	Value   float64 // value
}

type FfmModel struct {
	FeaturesNumber      int
	FieldsNumber        int
	LatentFactorsNumber int
	W                   []float64
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
