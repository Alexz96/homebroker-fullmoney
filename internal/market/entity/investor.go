package entity

type Investor struct {
	ID           string
	Name         string
	AssetPostion []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:           id,
		AssetPostion: []*InvestorAssetPosition{},
	}
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}
