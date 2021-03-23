package main

import (
	"flag"
	"log"
	"os"
	"pef/suite"
	"strings"

	"github.com/feiyuw/boomer"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("no suite specified! \nUsage: \n\npef %s [suite options]\n(use --help to show all options)\n", strings.Join(suite.SuiteMap.List(), "|"))
	}

	suite, err := suite.SuiteMap.Get(args[0])
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := suite.Init(flag.CommandLine, args[1:]); err != nil {
		log.Fatalln(err.Error())	
	}

	boomer.Run(suite.GetTask())
}
