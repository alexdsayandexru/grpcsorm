package configs

import (
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/errors"
	"os"
)

const (
	XApiKey                  = "X_API_KEY"
	XApiKeyValue             = "X_API_KEY_VALUE"
	EnabledValidation        = "ENABLED_VALIDATION"
	EnabledDelivery          = "ENABLED_DELIVERY"
	EnabledSecurity          = "ENABLED_SECURITY"
	EnabledTls               = "ENABLED_TLS"
	KafkaBrokers             = "KAFKA_BROKERS"
	TopicUserRegistration    = "TOPIC_USER_REGISTRATION"
	TopicUserLogin           = "TOPIC_USER_LOGIN"
	TopicUserLogout          = "TOPIC_USER_LOGOUT"
	TopicDeleteUserAccount   = "TOPIC_DELETE_USER_ACCOUNT"
	TopicUserUpdateData      = "TOPIC_USER_UPDATE_DATA"
	TopicUserDeleteAdmin     = "TOPIC_USER_DELETE_ADMIN"
	TopicUserUpdateAdmin     = "TOPIC_USER_UPDATE_ADMIN"
	TopicDeleteAccount       = "TOPIC_DELETE_ACCOUNT"
	TopicUserAccountRecovery = "TOPIC_USER_ACCOUNT_RECOVERY"
	TopicDirectory           = "TOPIC_DIRECTORY"
	TopicMigrateUser         = "TOPIC_MIGRATE_USER"
	CaCert                   = "CA_CERT"
	SaslUserName             = "SASL_USER_NAME"
	SaslPassword             = "SASL_PASSWORD"
)

func GetConfig() *Config {
	return &Config{
		XApiKey:                  GetEnv(XApiKey),
		XApiKeyValue:             GetEnv(XApiKeyValue),
		EnabledValidation:        GetEnvBool(EnabledValidation),
		EnabledDelivery:          GetEnvBool(EnabledDelivery),
		EnabledSecurity:          GetEnvBool(EnabledSecurity),
		EnabledTls:               GetEnvBool(EnabledTls),
		KafkaBrokers:             GetEnv(KafkaBrokers),
		TopicUserRegistration:    GetEnv(TopicUserRegistration),
		TopicUserLogin:           GetEnv(TopicUserLogin),
		TopicUserLogout:          GetEnv(TopicUserLogout),
		TopicDeleteUserAccount:   GetEnv(TopicDeleteUserAccount),
		TopicUserUpdateData:      GetEnv(TopicUserUpdateData),
		TopicUserDeleteAdmin:     GetEnv(TopicUserDeleteAdmin),
		TopicUserUpdateAdmin:     GetEnv(TopicUserUpdateAdmin),
		TopicDeleteAccount:       GetEnv(TopicDeleteAccount),
		TopicUserAccountRecovery: GetEnv(TopicUserAccountRecovery),
		TopicDirectory:           GetEnv(TopicDirectory),
		TopicMigrateUser:         GetEnv(TopicMigrateUser),
		CaCert:                   TryGetEnv(CaCert),
		SaslUserName:             TryGetEnv(SaslUserName),
		SaslPassword:             TryGetEnv(SaslPassword),
	}
}

type Config struct {
	XApiKey                  string
	XApiKeyValue             string
	EnabledValidation        bool
	EnabledDelivery          bool
	EnabledSecurity          bool
	EnabledTls               bool
	KafkaBrokers             string
	TopicUserRegistration    string
	TopicUserLogin           string
	TopicUserLogout          string
	TopicDeleteUserAccount   string
	TopicUserUpdateData      string
	TopicUserDeleteAdmin     string
	TopicUserUpdateAdmin     string
	TopicDeleteAccount       string
	TopicUserAccountRecovery string
	TopicDirectory           string
	TopicMigrateUser         string
	CaCert                   string
	SaslUserName             string
	SaslPassword             string
}

func TryGetEnv(name string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return ""
}

func GetEnv(name string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	panic(errors.NotFoundKeyInEnvironmentError(name).Error())
}

func GetEnvBool(name string) bool {
	if value, ok := os.LookupEnv(name); ok {
		return value == "on"
	}
	panic(errors.NotFoundKeyInEnvironmentError(name).Error())
}
