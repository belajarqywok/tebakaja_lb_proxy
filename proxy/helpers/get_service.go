package helpers

import "math/rand"

var serviceUrls = map[string][]string{
	"crypto": {
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
	},
	"national": {
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
	},
	"stock": {
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
		"https://qywok-cryptocurrency-prediction.hf.space/crypto",
	},
}

func GetEndpointService(svc_name string) string {
	var selectedUrls []string
	switch svc_name {
	case "crypto":
		selectedUrls = serviceUrls["crypto"]
	case "national":
		selectedUrls = serviceUrls["national"]
	default:
		selectedUrls = serviceUrls["stock"]
	}

	randomIndex := rand.Intn(len(selectedUrls))

	return selectedUrls[randomIndex]
}
