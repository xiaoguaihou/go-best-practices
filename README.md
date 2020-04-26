# go-best-practices

Try to show a best practices of a restful go project, including:
- project layout
- http framework
- application configuration
- database orm solution
- cron job solution
- json marshal/unarshal
- logger

## project layout
```
|-config : application configuration, we use yml format
|
|-pkg : we use domain based layout. one package in pkg folder implement a independent function
| |
| |- core : default service name
| |   |- main_core.go  : main file of core service
| |   |- api_dto.go : core service api definition
| |   |- ext_api_dto.go : external service api definition
| |   |- db_dto.go : database orm object definition
| |   |- corn_job.go : a sample of sub-function implementation
| |- logger :
|     |- gin_writer.go : gin writer, save gin log to log file
|     |- gorm_logger.go : gorm log to log file
|- main.go : main file of the project
|- routers.go : all restful api register in here
|- version.go : project version
 
```

## http framework

[gin](https://github.com/gin-gonic/gin)

## application configuration

use yaml configure file, if need configuration center, I prefer [Apollo](https://github.com/ctripcorp/apollo) (although there is no official go client)

## database orm solution

[gorm](https://github.com/jinzhu/gorm) seem good enough

[gen](https://github.com/smallnest/gen) : The gen tool produces golang structs from a given database for use in a .go file

## cron job solution

```
"github.com/robfig/cron/v3"
```

## json marshal/unarshal

use native library

## logger

Maybe [seelog](https://github.com/cihub/seelog) seems not cool as [logrus](https://github.com/sirupsen/logrus) or [zap](https://github.com/uber-go/zap). But in the matter of a logger, seelog is the best one:

- async logging
- file rotation&zip
- logging format configuration
