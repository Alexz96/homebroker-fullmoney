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

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPostion = append(i.AssetPostion, assetPosition)
}

func (i *Investor) AdjustAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)

	if assetPosition == nil {
		i.AssetPostion = append(i.AssetPostion, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.AddShares(qtdShares)
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPostion {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, qtdShares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  qtdShares,
	}
}

func (iap *InvestorAssetPosition) AddShares(qtd int) {
	iap.Shares += qtd
}
