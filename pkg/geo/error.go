package geo

import "errors"

var (
	ErrFailedToGetAccessToYOLP = errors.New("geo: failed to get access to YOLP")
	ErrUnableToDecodeResponse  = errors.New("geo: unable to decode response of YOLP to JSON format")
)
