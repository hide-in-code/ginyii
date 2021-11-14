package config

import (
	"encoding/json"
	"log"
	"os"
)

type CommonConf struct {
	Config struct {
		Timezone string `json:"timezone"`
		LogDebug string `json:"log_debug"`
		Server struct {
			HOST string `json:"HOST"`
			PORT string `json:"PORT"`
			ENV string `json:"ENV"`
		} `json:"server"`
	} `json:"config"`
}

type MysqlConf struct {
	Config struct {
		Mysql struct {
			HOST string `json:"DB_HOST"`
			PORT string `json:"DB_PORT"`
			NAME string `json:"DB_NAME"`
			USER string `json:"DB_USER"`
			PWD string `json:"DB_PWD"`
			CHARSET string `json:"DB_CHARSET"`
		} `json:"mysql"`
	} `json:"config"`
}

type RedisConf struct {
	Config struct {
		Redis struct {
			HOST string `json:"REDIS_HOST"`
			PORT string `json:"REDIS_PORT"`
			PWD string `json:"REDIS_PWD"`
			DB string `json:"REDIS_DB"`
		} `json:"redis"`
	} `json:"config"`
}

func GetCommonConf() map[string]string {
	configJson, _ := os.Open("config/config.json")
	defer configJson.Close()
	var commonConf CommonConf
	decoder := json.NewDecoder(configJson)
	err := decoder.Decode(&commonConf)
	if err != nil {
		log.Println("'config.json' framework：Decoder Error = ", err.Error(), "common config加载失败")
		//os.Exit(200)
	}

	timezone := commonConf.Config.Timezone
	LogDebug := commonConf.Config.LogDebug
	host := commonConf.Config.Server.HOST
	port := commonConf.Config.Server.PORT
	env := commonConf.Config.Server.ENV

	if len(timezone) == 0 {
		timezone = "Asia/Shanghai"
	}

	if len(LogDebug) == 0 {
		LogDebug = "true"
	}

	if len(host) == 0 || host == "localhost" {
		host = "0.0.0.0"
	}

	if len(port) == 0 {
		port = "8080"
	}

	if len(env) == 0 {
		env = "release"
	}

	conf := make(map[string]string)
	conf["HOST"] = host // 服务地址
	conf["PORT"] = port // 服务端口
	conf["ENV"] = env // env
	conf["timezone"] = timezone // 时区
	conf["Log_debug"] = LogDebug // 时区

	return conf
}

func GetMysqlConf() map[string]string {
	configJson, _ := os.Open("config/config.json")
	defer configJson.Close()
	var mysqlConf MysqlConf
	decoder := json.NewDecoder(configJson)
	err := decoder.Decode(&mysqlConf)
	if err != nil {
		log.Println("'config.json' framework：Decoder Error = ", err.Error(), "mysql config加载失败")
		//os.Exit(200)
	}

	DB_HOST := mysqlConf.Config.Mysql.HOST
	DB_PORT := mysqlConf.Config.Mysql.PORT
	DB_NAME := mysqlConf.Config.Mysql.NAME
	DB_USER := mysqlConf.Config.Mysql.USER
	DB_PWD := mysqlConf.Config.Mysql.PWD
	DB_CHARSET := mysqlConf.Config.Mysql.CHARSET

	if len(DB_HOST) == 0 {
		DB_HOST = "127.0.0.1"
	}

	if len(DB_PORT) == 0 {
		DB_PORT = "3306"
	}

	if len(DB_NAME) == 0 {
		DB_NAME = "mydb"
	}

	if len(DB_USER) == 0 {
		DB_USER = "root"
	}

	if len(DB_PWD) == 0 {
		DB_PWD = "root"
	}

	if len(DB_CHARSET) == 0 {
		DB_CHARSET = "utf8"
	}

	conf := make(map[string]string)
	conf["DB_HOST"] = DB_HOST // host
	conf["DB_PORT"] = DB_PORT // 端口
	conf["DB_NAME"] = DB_NAME // 数据库名称
	conf["DB_USER"] = DB_USER // 用户
	conf["DB_PWD"] = DB_PWD // 密码
	conf["DB_CHARSET"] = DB_CHARSET // 字符集

	return conf
}


func GetRedisConf() map[string]string {
	configJson, _ := os.Open("config/config.json")
	defer configJson.Close()
	var redisConf RedisConf
	decoder := json.NewDecoder(configJson)
	err := decoder.Decode(&redisConf)
	if err != nil {
		log.Println("'config.json' framework：Decoder Error = ", err.Error(), "redis config加载失败")
		//os.Exit(200)
	}

	REDIS_HOST := redisConf.Config.Redis.HOST
	REDIS_PORT := redisConf.Config.Redis.PORT
	REDIS_PWD := redisConf.Config.Redis.PWD
	REDIS_DB := redisConf.Config.Redis.DB

	if len(REDIS_HOST) == 0 {
		REDIS_HOST = "127.0.0.1"
	}

	if len(REDIS_PORT) == 0 {
		REDIS_PORT = "6379"
	}

	if len(REDIS_PWD) == 0 {
		REDIS_PWD = ""
	}

	if len(REDIS_DB) == 0 {
		REDIS_DB = "0"
	}

	conf := make(map[string]string)
	conf["REDIS_HOST"] = REDIS_HOST // host
	conf["REDIS_PORT"] = REDIS_PORT // 端口
	conf["REDIS_PWD"] = REDIS_PWD // redis密码，无密码填空
	conf["REDIS_DB"] = REDIS_DB // 库

	return conf
}
