package configs

var EnvConfigs = map[string]map[string]interface{}{
	"development": {
		"db": map[string]interface{}{
			"host":     "localhost",
			"port":     27017,
			"username": "root",
			"password": "root",
			"database": "hepsi-ads-campaign",
		},
	},
}
