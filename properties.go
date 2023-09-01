package azuretexttospeech

// AudioOutput types represent the supported audio encoding formats for the text-to-speech endpoint.
// This type is required when requesting to azuretexttospeech.Synthesize text-to-speed request.
// Each incorporates a bitrate and encoding type. The Speech service supports 24 kHz, 16 kHz, and 8 kHz audio outputs.
// See: https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-text-to-speech#audio-outputs
type AudioOutput int

const (
	AudioRIFF8Bit8kHzMonoPCM AudioOutput = iota
	AudioRIFF16Bit16kHzMonoPCM
	AudioRIFF16khz16kbpsMonoSiren
	AudioRIFF24khz16bitMonoPcm
	AudioRAW8Bit8kHzMonoMulaw
	AudioRAW16Bit16kHzMonoMulaw
	AudioRAW24khz16bitMonoPcm
	AudioSsml16khz16bitMonoTts
	Audio16khz16kbpsMonoSiren
	Audio16khz32kbitrateMonoMp3
	Audio6khz64kbitrateMonoMp3
	Audio16khz128kbitrateMonoMp3
	Audio24khz48kbitrateMonoMp3
	Audio24khz96kbitrateMonoMp3
	Audio24Khz160KBitRateMonoMp3 
	Raw16Khz16BitMonoTrueSilk 
	Riff16Khz16BitMonoPcm 
	Riff8Khz16BitMonoPcm 
	Riff24Khz16BitMonoPcm 
	Riff8Khz8BitMonoMULaw 
	Raw16Khz16BitMonoPcm 
	Raw24Khz16BitMonoPcm 
	Raw8Khz16BitMonoPcm 
	Ogg16Khz16BitMonoOpus 
	Ogg24Khz16BitMonoOpus 
	Raw48Khz16BitMonoPcm 
	Riff48Khz16BitMonoPcm 
	Audio48Khz96KBitRateMonoMp3 
	Audio48Khz192KBitRateMonoMp3 
	Ogg48Khz16BitMonoOpus 
	Webm16Khz16BitMonoOpus 
	Webm24Khz16BitMonoOpus 
	Raw24Khz16BitMonoTrueSilk 
	Raw8Khz8BitMonoALaw 
	Riff8Khz8BitMonoALaw 
	Webm24Khz16Bit24KbpsMonoOpus 
	Audio16Khz16Bit32KbpsMonoOpus 
	Audio24Khz16Bit48KbpsMonoOpus 
	Audio24Khz16Bit24KbpsMonoOpus 
	Raw22050Hz16BitMonoPcm 
	Riff22050Hz16BitMonoPcm 
	Raw44100Hz16BitMonoPcm 
	Riff44100Hz16BitMonoPcm 
)

func (a AudioOutput) String() string {
	return []string{"riff-8khz-8bit-mono-mulaw",
		"riff-16khz-16bit-mono-pcm",
		"riff-16khz-16kbps-mono-siren",
		"riff-24khz-16bit-mono-pcm",
		"raw-8khz-8bit-mono-mulaw",
		"raw-16khz-16bit-mono-pcm",
		"raw-24khz-16bit-mono-pcm",
		"ssml-16khz-16bit-mono-tts",
		"audio-16khz-16kbps-mono-siren",
		"audio-16khz-32kbitrate-mono-mp3",
		"audio-16khz-64kbitrate-mono-mp3",
		"audio-16khz-128kbitrate-mono-mp3",
		"audio-24khz-48kbitrate-mono-mp3",
		"audio-24khz-96kbitrate-mono-mp3",
		"audio-24khz-160kbitrate-mono-mp3",
		"raw-16khz-16bit-mono-truesilk",
		"riff-16khz-16bit-mono-pcm",
		"riff-8khz-16bit-mono-pcm",
		"riff-24khz-16bit-mono-pcm",
		"riff-8khz-8bit-mono-mulaw",
		"raw-16khz-16bit-mono-pcm",
		"raw-24khz-16bit-mono-pcm",
		"raw-8khz-16bit-mono-pcm",
		"ogg-16khz-16bit-mono-opus",
		"ogg-24khz-16bit-mono-opus",
		"raw-48khz-16bit-mono-pcm",
		"riff-48khz-16bit-mono-pcm",
		"audio-48khz-96kbitrate-mono-mp3",
		"audio-48khz-192kbitrate-mono-mp3",
		"ogg-48khz-16bit-mono-opus",
		"webm-16khz-16bit-mono-opus",
		"webm-24khz-16bit-mono-opus",
		"raw-24khz-16bit-mono-truesilk",
		"raw-8khz-8bit-mono-alaw",
		"riff-8khz-8bit-mono-alaw",
		"webm-24khz-16bit-24kbps-mono-opus",
		"audio-16khz-16bit-32kbps-mono-opus",
		"audio-24khz-16bit-48kbps-mono-opus",
		"audio-24khz-16bit-24kbps-mono-opus",
		"raw-22050hz-16bit-mono-pcm",
		"riff-22050hz-16bit-mono-pcm",
		"raw-44100hz-16bit-mono-pcm",
		"riff-44100hz-16bit-mono-pcm",
	}[a]
}

// Gender type for the digitized language
//go:generate enumer -type=Gender -linecomment -json
type Gender int

const (
	// GenderMale , GenderFemale are the static Gender constants for digitized voices.
	// See Gender in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices for breakdown
	GenderMale   Gender = iota // Male
	GenderFemale               // Female
)

// Locale references the language or locale for text-to-speech.
// See "locale" in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices
//go:generate enumer -type=Locale -linecomment -json
type Locale int

const (
	LocaleArEG Locale = iota //ar-EG
	LocaleArSA               // ar-SA
	LocaleBgBG               // bg-BG
	LocaleCaES               // ca-ES
	LocaleCsCZ               // cs-CZ
	LocaleDaDK               // da-DK
	LocaleDeAT               // de-AT
	LocaleDeCH               // de-CH
	LocaleDeDE               // de-DE
	LocaleElGR               // el-GR
	LocaleEnAU               // en-AU
	LocaleEnCA               // en-CA
	LocaleEnGB               // en-GB
	LocaleEnIE               // en-IE
	LocaleEnIN               // en-IN
	LocaleEnUS               // en-US
	LocaleEsES               // es-ES
	LocaleEsMX               // es-MX
	LocaleEtEE               // et-EE
	LocaleFiFI               // fi-FI
	LocaleFrCA               // fr-CA
	LocaleFrCH               // fr-CH
	LocaleFrFR               // fr-FR
	LocaleGaIE               // ga-IE
	LocaleHeIL               // he-IL
	LocaleHiIN               // hi-IN
	LocaleHrHR               // hr-HR
	LocaleHuHU               // hu-HU
	LocaleIdID               // id-ID
	LocaleItIT               // it-IT
	LocaleJaJP               // ja-JP
	LocaleKoKR               // ko-KR
	LocaleLtLT               // lt-LT
	LocaleLvLV               // lv-LV
	LocaleMtMT               // mt-MT
	LocaleMrIN               // mr-IN
	LocaleMsMY               // ms-MY
	LocaleNbNO               // nb-NO
	LocaleNlNL               // nl-NL
	LocalePlPL               // pl-PL
	LocalePtBR               // pt-BR
	LocalePtPT               // pt-PT
	LocaleRoRO               // ro-RO
	LocaleRuRU               // ru-RU
	LocaleSkSK               // sk-SK
	LocaleSlSI               // sl-SI
	LocaleSvSE               // sv-SE
	LocaleTaIN               // ta-IN
	LocaleTeIN               // te-IN
	LocaleThTH               // th-TH
	LocaleTrTR               // tr-TR
	LocaleViVN               // vi-VN
	LocaleZhCN               // zh-CN
	LocaleZhHK               // zh-HK
	LocaleZhTW               // zh-TW
)

// Region references the locations of the availability of standard voices.
// See https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/regions#standard-voices
type Region int

const (
	// Azure regions and their endpoints that support the Text To Speech service.
	RegionAustraliaEast Region = iota
	RegionBrazilSouth
	RegionCanadaCentral
	RegionCentralUS
	RegionEastAsia
	RegionEastUS
	RegionEastUS2
	RegionFranceCentral
	RegionIndiaCentral
	RegionJapanEast
	RegionJapanWest
	RegionKoreaCentral
	RegionNorthCentralUS
	RegionNorthEurope
	RegionSouthCentralUS
	RegionSoutheastAsia
	RegionUKSouth
	RegionWestEurope
	RegionWestUS
	RegionWestUS2
)

func (t Region) String() string {
	return [...]string{
		"australiaeast",
		"brazilsouth",
		"canadacentral",
		"centralus",
		"eastasia",
		"eastus",
		"eastus2",
		"francecentral",
		"indiacentral",
		"japaneast",
		"japanwest",
		"koreacentral",
		"northcentralus",
		"northeurope",
		"southcentralus",
		"southeastasia",
		"uksouth",
		"westeurope",
		"westus",
		"westus2",
	}[t]

}
