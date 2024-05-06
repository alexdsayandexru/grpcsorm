package security

import (
	"context"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/errors"
	"google.golang.org/grpc/metadata"
)

type Security struct {
	apiKey      string
	apiKeyValue string
}

func (p *Security) Auth(ctx *context.Context) (bool, error) {
	cApiKeyValue, err := GetApikeyFromIncomingContext(ctx, p.apiKey)
	if err != nil {
		return false, err
	}
	if p.apiKeyValue != cApiKeyValue {
		return false, errors.InvalidApiKeyError()
	}
	return true, nil
}

func GetApikeyFromIncomingContext(ctx *context.Context, apiKey string) (string, error) {
	md, _ := metadata.FromIncomingContext(*ctx)
	values := md.Get(apiKey)
	if len(values) == 0 {
		return "", errors.EmptyContextMetadataError()
	}

	return values[0], nil
}

func NewSecurity(apiKey string, apiKeyValue string) ISecurity {
	return &Security{
		apiKey:      apiKey,
		apiKeyValue: apiKeyValue,
	}
}
