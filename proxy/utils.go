package proxy

import "time"

func GetEndpointByRestService(svc_name string) string {
	second := time.Now().Second()

	var divisorToString map[int]string
	switch svc_name {
		case "crypto":
			divisorToString = map[int]string{
				9: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				7: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				5: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				3: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				2: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
			}
		case "national":
			divisorToString = map[int]string{
				9: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				7: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				5: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				3: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				2: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
			}
		default:
			divisorToString = map[int]string{
				9: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				7: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				5: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				3: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
				2: "https://qywok-cryptocurrency-prediction.hf.space/crypto",
			}
	}

	var result string
	for divisor, str := range divisorToString {
		if (second % divisor) == 0 {
			result += str
			break
		}
	}

	return result
}