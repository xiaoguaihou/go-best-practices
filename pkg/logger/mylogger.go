package logger

import (
	"fmt"
	"log"

	seelog "github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

func Init(file string) {
	config := `
<seelog minlevel="info" type="asynctimer" asyncinterval="500000000">
	<outputs formatid="common">
		<console/>
		<rollingfile type="date" 
					 filename="/data/log/go/best-practices/best-practices.log" 
					 maxrolls="30"
					 datepattern="2006-01-02"
					 archivetype="zip"/>
        
    </outputs>
    <formats>
        <format id="common" format="%DateT%Time %File:%Line:%FuncShort %LEV %Msg%n" />
    </formats>
</seelog>
`
	var err error
	if len(file) == 0 {
		logger, err = seelog.LoggerFromConfigAsBytes([]byte(config))
	} else {
		logger, err = seelog.LoggerFromConfigAsFile(file)
	}

	if err != nil {
		log.Fatal(fmt.Sprintf("Init seelog failure: %+v\n", err))
	}

	logger.Flush()
	_ = seelog.ReplaceLogger(logger)
}
