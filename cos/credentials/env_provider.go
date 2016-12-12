package credentials

import (
	"os"

	"github.com/tencentyun/cos-sdk-go/cos/coserr"
)

// EnvProviderName provides a name of Env provider
const EnvProviderName = "EnvProvider"

var (
	// ErrAccessKeyIDNotFound is returned when the COS Access Key ID can't be
	// found in the process's environment.
	//
	// @readonly
	ErrAccessKeyIDNotFound = coserr.New("EnvAccessKeyNotFound", "COS_ACCESS_KEY_ID or COS_ACCESS_KEY not found in environment", nil)

	// ErrSecretAccessKeyNotFound is returned when the COS Secret Access Key
	// can't be found in the process's environment.
	//
	// @readonly
	ErrSecretAccessKeyNotFound = coserr.New("EnvSecretNotFound", "COS_SECRET_ACCESS_KEY or COS_SECRET_KEY not found in environment", nil)

    ErrAppIDNotFount = coserr.New("AppIDNotFound", "COS_APPID not found in environment", nil)
)

// A EnvProvider retrieves credentials from the environment variables of the
// running process. Environment credentials never expire.
//
// Environment variables used:
//
// * Access Key ID:     COS_ACCESS_KEY_ID or COS_ACCESS_KEY
// * Secret Access Key: COS_SECRET_ACCESS_KEY or COS_SECRET_KEY
// * Application ID:    COS_APPID
type EnvProvider struct {
	retrieved bool
}

// NewEnvCredentials returns a pointer to a new Credentials object
// wrapping the environment variable provider.
func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
	e.retrieved = false

	id := os.Getenv("COS_ACCESS_KEY_ID")
	if id == "" {
		id = os.Getenv("COS_ACCESS_KEY")
	}

	secret := os.Getenv("COS_SECRET_ACCESS_KEY")
	if secret == "" {
		secret = os.Getenv("COS_SECRET_KEY")
	}

    appid := os.Getenv("COS_APPID")

	if id == "" {
		return Value{ProviderName: EnvProviderName}, ErrAccessKeyIDNotFound
	}

	if secret == "" {
		return Value{ProviderName: EnvProviderName}, ErrSecretAccessKeyNotFound
	}

    if appid == "" {
        return Value{ProviderName: EnvProviderName}, ErrAppIDNotFount
    }

	e.retrieved = true
	return Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
        AppID:           appid,
		ProviderName:    EnvProviderName,
	}, nil
}

// IsExpired returns if the credentials have been retrieved.
func (e *EnvProvider) IsExpired() bool {
	return !e.retrieved
}
