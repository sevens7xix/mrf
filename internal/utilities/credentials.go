package utilities

import (
	"github.com/spf13/viper"
)

func GetCredentials() ([]string, error) {
	// Here we're going to retreive the credentials on the .env file using viper
	var credentials []string

	viper.SetConfigName("credentials")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/mrf")
	viper.AddConfigPath("/app")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	// First we're going to get the client id, and later the client secret key
	credentials = append(credentials, viper.GetString("CLIENT_ID"))
	credentials = append(credentials, viper.GetString("CLIENT_SECRET"))

	return credentials, nil
}
