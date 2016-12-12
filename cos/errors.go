package cos

import "github.com/tencentyun/cos-sdk-go/cos/coserr"

var (
	// ErrMissingRegion is an error that is returned if region configuration is
	// not found.
	//
	// @readonly
	ErrMissingRegion = coserr.New("MissingRegion", "could not find region configuration", nil)

	// ErrMissingEndpoint is an error that is returned if an endpoint cannot be
	// resolved for a service.
	//
	// @readonly
	ErrMissingEndpoint = coserr.New("MissingEndpoint", "'Endpoint' configuration is required for this service", nil)
)
