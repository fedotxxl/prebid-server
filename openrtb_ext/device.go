package openrtb_ext

import (
	"errors"
	"strconv"

	"github.com/buger/jsonparser"
)

// ExtDevice defines the contract for bidrequest.device.ext
type ExtDevice struct {
	Prebid ExtDevicePrebid `json:"prebid"`
}

// Pointer to interstitial so we do not force it to exist
type ExtDevicePrebid struct {
	Interstitial *ExtDeviceInt `json:"interstitial"`
}

type ExtDeviceInt struct {
	MinWidthPerc  int64 `json:"minwidtheperc"`
	MinHeightPerc int64 `json:"minheightperc"`
}

func (edi *ExtDeviceInt) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return errors.New("request.device.ext.prebid.interstitial must have some data in it")
	}
	if value, dataType, _, _ := jsonparser.Get(b, "minwidthperc"); dataType != jsonparser.Number {
		return errors.New(`request.device.ext.prebid.interstitial.minwidthperc must be a number between 0 and 100`)
	} else {
		perc, err := strconv.Atoi(string(value))
		if err != nil || perc < 0 || perc > 100 {
			return errors.New(`request.device.ext.prebid.interstitial.minwidthperc must be a number between 0 and 100`)
		}
		edi.MinWidthPerc = int64(perc)
	}
	if value, dataType, _, _ := jsonparser.Get(b, "minheightperc"); dataType != jsonparser.Number {
		return errors.New(`request.device.ext.prebid.interstitial.minheightperc must be a number between 0 and 100`)
	} else {
		perc, err := strconv.Atoi(string(value))
		if err != nil || perc < 0 || perc > 100 {
			return errors.New(`request.device.ext.prebid.interstitial.minheightperc must be a number between 0 and 100`)
		}
		edi.MinHeightPerc = int64(perc)
	}
	return nil
}
