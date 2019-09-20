package config

var config =  map[string]interface{}{
	"envrionment": "test",
}

func Init(){
	config["database"] = "mysql"
}

func GetConfig() map[string]interface{} {
 	return config
}