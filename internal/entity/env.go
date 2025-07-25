package entity

type Env struct {
	SERVER_APP_NAME                  string `mapstructure:"SERVER_APP_NAME"`
	SERVER_SERVER_NAME               string `mapstructure:"SERVER_SERVER_NAME"`
	SERVER_PORT                      string `mapstructure:"SERVER_PORT"`
	SERVER_HOST                      string `mapstructure:"SERVER_HOST"`
	SERVER_PREFORK                   bool   `mapstructure:"SERVER_PREFORK"`
	SERVER_APP_ENV                   string `mapstructure:"SERVER_APP_ENV"`
	DATABASE_HOST                    string `mapstructure:"DATABASE_HOST"`
	DATABASE_PORT                    string `mapstructure:"DATABASE_PORT"`
	DATABASE_USER                    string `mapstructure:"DATABASE_USER"`
	DATABASE_PASS                    string `mapstructure:"DATABSE_PASS"`
	DATABASE_NAME                    string `mapstructure:"DATABASE_NAME"`
	DATABASE_TIMEZONE                string `mapstructure:"DATABASE_TIMEZONE"`
	DATABASE_SSLMODE                 string `mapstructure:"DATABASE_SSLMODE"`
	DATABASE_POOL_IDLE               int    `mapstructure:"DATABASE_POOL_IDLE"`
	DATABASE_MAX_CONNECTIONS         int    `mapstructure:"DATABASE_MAX_CONNECTIONS"`
	DATABASE_MAXLIFETIME_CONNECTIONS int    `mapstructure:"DATABASE_MAXLIFETIME_CONNECTIONS"`
	LOG_LEVEL                        int    `mapstructure:"LOG_LEVEL"`
	ALLOWED_ORIGINS                  string `mapstructure:"ALLOWED_ORIGINS"`
	REFRESH_COOKIE_NAME              string `mapstructure:"REFRESH_COOKIE_NAME"`
	REFRESH_COOKIE_DOMAIN            string `mapstructure:"REFRESH_COOKIE_DOMAIN"`
	REFRESH_COOKIE_PATH              string `mapstructure:"REFRESH_COOKIE_PATH"`
	REFRESH_COOKIE_SECURE            bool   `mapstructure:"REFRESH_COOKIE_SECURE"`
	REFRESH_COOKIE_HTTPONLY          bool   `mapstructure:"REFRESH_COOKIE_HTTPONLY"`
	REFRESH_COOKIE_SAMESITE          string `mapstructure:"REFRESH_COOKIE_SAMESITE"`
	REFRESH_COOKIE_MAXAGE            int    `mapstructure:"REFRESH_COOKIE_MAXAGE"`
	ACCESS_COOKIE_NAME               string `mapstructure:"ACCESS_COOKIE_NAME"`
	ACCESS_COOKIE_DOMAIN             string `mapstructure:"ACCESS_COOKIE_DOMAIN"`
	ACCESS_COOKIE_PATH               string `mapstructure:"ACCESS_COOKIE_PATH"`
	ACCESS_COOKIE_SECURE             bool   `mapstructure:"ACCESS_COOKIE_SECURE"`
	ACCESS_COOKIE_HTTPONLY           bool   `mapstructure:"ACCESS_COOKIE_HTTPONLY"`
	ACCESS_COOKIE_SAMESITE           string `mapstructure:"ACCESS_COOKIE_SAMESITE"`
	ACCESS_COOKIE_MAXAGE             int    `mapstructure:"ACCESS_COOKIE_MAXAGE"`
}
