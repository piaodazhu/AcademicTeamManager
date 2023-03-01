package initializer

func Init() interface{} {
	loadConf()
	initRedis()
	initMysql()
	engine := initRouter()
	return engine
}
