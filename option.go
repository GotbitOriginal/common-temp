package common

type Option struct {
	Name  string
	Value interface{}
}

const (
	KEY        = "KEY"
	SECRET     = "SECRET"
	ACCOUNT_ID = "ACCOUNT_ID"
	PASSPHRASE = "PASSPHRASE"
)

func Set(name string, value interface{}) Option {
	return Option{name, value}
}

func SetKey(key string) Option {
	return Set(KEY, key)
}

func SetSecret(secret string) Option {
	return Set(SECRET, secret)
}

func SetAccountID(accountId string) Option {
	return Set(ACCOUNT_ID, accountId)
}

func SetPassphrase(passphrase string) Option {
	return Set(PASSPHRASE, passphrase)
}
