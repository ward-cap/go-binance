package binance

//go:generate easyjson -all models.go

//easyjson:json
type GetAllAssetsResponse struct {
	Data []struct {
		AssetCode  string `json:"assetCode"`
		LogoUrl    string `json:"logoUrl"`
		AssetDigit int64  `json:"assetDigit"`
		Trading    bool   `json:"trading"`
	} `json:"data"`
	Success bool `json:"success"`
}
