package main

import (
	"flag"
	"fmt"
	"log"

	"uhe.com/go-best-practices/config"
	"uhe.com/go-best-practices/pkg/core"

	logger "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	mylogger "uhe.com/go-best-practices/pkg/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// config keys
const (
	KEY_SALES_DB_HOST      = "DB.salesHost"
	KEY_SALES_DB_USER      = "DB.salesUser"
	KEY_SALES_DB_PWD       = "DB.salesPwd"
	KEY_SALES_DB_MAX_CONNS = "DB.salesMaxConns"
	KEY_PORT               = "port"
)

// config related
var (
	configfile string
	loggerfile string
	port       int
	help       bool

	appConfig config.AppConfigViper
)

// app biz related
var (
	dbLogger  mylogger.DbLogger
	ginWriter mylogger.GinWriter

	salesDB     *gorm.DB
	coreService *core.CoreService
)

// Do framework level init
func init() {
	flag.StringVar(&configfile, "configfile", "/data/config/go/best-practices/config.yml", "config file full path")
	flag.StringVar(&loggerfile, "loggerfile", "", "seelog config file")
	flag.BoolVar(&help, "h", false, "show help")
	flag.IntVar(&port, "port", 0, "service port to listen")
	flag.Parse()

	if help {
		flag.Usage()
	}
	// init logger firstly!!!
	mylogger.Init(loggerfile)

	appConfig.GetConfig(configfile)

	logger.Infof("Init with config:%+v", appConfig)
}

// Do application leve init
func initMain() {
	// init db connection
	salesDB = openMySql(
		appConfig.GetString(KEY_SALES_DB_HOST),
		appConfig.GetString(KEY_SALES_DB_USER),
		appConfig.GetString(KEY_SALES_DB_PWD),
		appConfig.GetInt(KEY_SALES_DB_MAX_CONNS),
		"SALSES")

	coreService = core.New(
		salesDB,
	)

	coreService.Start()
}

func openMySql(host string, user string, pwd string, maxConns int, info string) *gorm.DB {
	url := fmt.Sprintf("%s:%s@(%s)/stool?charset=utf8&parseTime=True&loc=Local",
		user, pwd, host)

	db, err := gorm.Open("mysql", url)

	if err != nil {
		logger.Errorf("open database(%s) @%s failed , err=%+v\n", info, host, err)
		logger.Flush()
		log.Fatalf("open database(%s) @%s failed , err=%+v\n", info, host, err)
	} else {
		logger.Infof("open database(%s) @%s success!", info, host)
	}

	db.DB().SetMaxOpenConns(maxConns)
	db.SingularTable(true)
	db.LogMode(true)
	// replay gorm log to log file
	db.SetLogger(dbLogger)

	return db
}

func main() {
	initMain()

	// start restful framework
	gin.SetMode(gin.DebugMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = ginWriter
	gin.DefaultErrorWriter = ginWriter
	r := gin.Default()
	setupApi(r)

	// use startup parameter if there is
	p := port
	if p == 0 {
		p = appConfig.GetInt(KEY_PORT)
	}

	logger.Infof("start to listen port:%d", p)
	r.Run(fmt.Sprintf(":%d", p))
}
