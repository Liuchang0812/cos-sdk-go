package cos

import (
    "github.com/tencentyun/cos-sdk-go/cos"
    "github.com/tencentyun/cos-sdk-go/cos/client"
    "github.com/tencentyun/cos-sdk-go/cos/client/metadata"
    "github.com/tencentyun/cos-sdk-go/cos/request"
    "github.com/tencentyun/cos-sdk-go/cos/protocol/restxml"

type Cos struct {
    *client.Client
}

const ServiceName = "s3"


func New(cfg cos.Config) *Cos {
    svc := &Cos{
        Client: client.New(
              cfg,
              metadata.ClientInfo{
                ServiceName: ServiceName,
                SigningName: cfg.SigningName,
                SigningRegion: cfg.SigningRegion,
                Endpoint: cfg.Endpoint,
                APIVersion: "2006-03-01",
              },
              handlers,
        ),
    }
	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

    return svc
}

// newRequest creates a new request for a S3 operation and runs any
// custom request initialization.
func (c *S3) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}
