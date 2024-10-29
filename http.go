package translation

import (
	"sort"
	"strconv"
	"strings"
)

type AcceptLanguage struct {
	Lang    string // de-AT, de
	Base    string // de, de
	Region  string // AT, ""
	Quality float64
}

// ParseAcceptLanguage parses an Accept-Language header into a slice of AcceptLanguage, sorted by quality in descending order.
func ParseAcceptLanguage(headerLine string) []AcceptLanguage {
	ret := make([]AcceptLanguage, 0)

	for _, lq := range strings.Split(headerLine, ",") {
		lq = strings.Trim(lq, " ")

		langQuality := strings.SplitN(lq, ";", 2)

		if langQuality[0] == "" {
			continue
		}

		al := AcceptLanguage{Lang: langQuality[0]}

		langRegion := strings.Split(al.Lang, "-")
		al.Base = langRegion[0]
		if len(langRegion) > 1 {
			al.Region = langRegion[1]
		}

		quality := "1"

		if len(langQuality) > 1 {
			qVal := strings.Split(langQuality[1], "=")
			if len(qVal) < 2 {
				quality = "0"
			} else {
				quality = strings.Trim(qVal[1], " ")
			}
		}

		qFloat, err := strconv.ParseFloat(quality, 64)
		if err != nil {
			qFloat = 0
		}

		al.Quality = qFloat

		ret = append(ret, al)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Quality > ret[j].Quality
	})

	return ret
}

// GetBaseLanguage extracts the base language from the given AcceptLanguage parameter.
func GetBaseLanguage(acceptLanguage AcceptLanguage) string {
	return strings.Split(acceptLanguage.Lang, "-")[0]
}
