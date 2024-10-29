package translation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAcceptLanguage_Simple(t *testing.T) {
	langs := ParseAcceptLanguage("de")
	assert.Equal(t, []AcceptLanguage{{Lang: "de", Base: "de", Region: "", Quality: 1}}, langs)
}

func TestParseAcceptLanguage_WithQuality(t *testing.T) {
	langs := ParseAcceptLanguage("de-DE;q=0.1")
	assert.Equal(t, []AcceptLanguage{{Lang: "de-DE", Base: "de", Region: "DE", Quality: 0.1}}, langs)
}

func TestParseAcceptLanguage_WithMalformedQuality(t *testing.T) {
	langs := ParseAcceptLanguage("de-DE;0.1")
	assert.Equal(t, []AcceptLanguage{{Lang: "de-DE", Base: "de", Region: "DE", Quality: 0}}, langs)

	langs = ParseAcceptLanguage("de-DE;xx")
	assert.Equal(t, []AcceptLanguage{{Lang: "de-DE", Base: "de", Region: "DE", Quality: 0}}, langs)

	langs = ParseAcceptLanguage("de-DE;q=xx")
	assert.Equal(t, []AcceptLanguage{{Lang: "de-DE", Base: "de", Region: "DE", Quality: 0}}, langs)
}

func TestParseAcceptLanguage_Multiple(t *testing.T) {
	langs := ParseAcceptLanguage("de, de-AT,fr")
	assert.Equal(t, []AcceptLanguage{
		{Lang: "de", Base: "de", Region: "", Quality: 1},
		{Lang: "de-AT", Base: "de", Region: "AT", Quality: 1},
		{Lang: "fr", Base: "fr", Region: "", Quality: 1}}, langs)
}

func TestParseAcceptLanguage_MultipleWithQuality(t *testing.T) {
	langs := ParseAcceptLanguage("de;q=1,de-AT;q=0.8,fr;q=0.2")
	assert.Equal(t, []AcceptLanguage{
		{Lang: "de", Base: "de", Region: "", Quality: 1},
		{Lang: "de-AT", Base: "de", Region: "AT", Quality: 0.8},
		{Lang: "fr", Base: "fr", Region: "", Quality: 0.2}}, langs)
}

func TestParseAcceptLanguage_MultipleWithQualitySort(t *testing.T) {
	langs := ParseAcceptLanguage("fr;q=0.2,de-AT;q=0.8,de;q=1,*;q=0.1")
	assert.Equal(t, []AcceptLanguage{
		{Lang: "de", Base: "de", Region: "", Quality: 1},
		{Lang: "de-AT", Base: "de", Region: "AT", Quality: 0.8},
		{Lang: "fr", Base: "fr", Region: "", Quality: 0.2},
		{Lang: "*", Base: "*", Region: "", Quality: 0.1}}, langs)
}

func TestParseAcceptLanguage_Empty(t *testing.T) {
	langs := ParseAcceptLanguage("")
	assert.Empty(t, langs)
}

func TestGetBaseLanguage(t *testing.T) {
	assert.Equal(t, "de", GetBaseLanguage(AcceptLanguage{Lang: "de-DE", Base: "de", Region: "DE"}))
	assert.Equal(t, "de", GetBaseLanguage(AcceptLanguage{Lang: "de", Base: "de"}))
}
