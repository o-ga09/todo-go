package config

type Config struct {
	// TODO : 環境変数設定
}
func New() (*Config, error) {
	return &Config{}, nil
}