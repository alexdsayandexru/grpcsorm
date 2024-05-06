package configs

import (
	"crypto/tls"
	"crypto/x509"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/errors"
	"os"
)

func GetTLS(caCert string, isRequired bool) *tls.Config {
	if !isRequired {
		return nil
	}

	var err error
	var cert []byte

	if len(caCert) == 0 {
		panic(errors.NotFoundCACertificateError().Error())
	} else if len(caCert) < 256 {
		cert, err = os.ReadFile(caCert)
		if err != nil {
			panic(errors.NotFoundFileCACertificateError(err.Error()).Error())
		}
	} else {
		cert = []byte(caCert)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	tlsConfig := tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}

	return &tlsConfig
}
