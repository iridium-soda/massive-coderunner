package main

import (
	"flag"
	"fmt"
	"github.com/iridium-soda/massive-coderunner/pkg/analysis"
	"github.com/iridium-soda/massive-coderunner/pkg/utils"
	"log"
	"os"
	"strings"
)

var Usage = `To analysis and save and query structured AST tree.

Usage:
astool analysis <file>			To analysis target file and gather useful information. 
astool locate <function>		To locate function by name. Be sure all files you need to search were resolved successfully before.
astool trace <function>:<arg>	To trace the statement inside the function that operates on this parameter.
astool -h
astool -v

`

/*Options:
locate <function>       :To locate function by name. Be sure all files you need to search were resolved successfully before.
trace <function>:<arg> :To trace the statement inside the function that operates on this parameter.
-h                  :Show this screen.
-v                  :Show version.*/
var (
	ifHelp           bool
	ifAnalysis       = false
	ifLocateFunction = false
	ifTraceArgs      = false
	ifVersion        bool
)
var version = "v0.0.0"
var (
	targetFilename    = "" //For analysis mode
	functionName      = "" //For locate mode
	traceFunctionName = "" //For trace mode
	traceArg          = "" //For trace mode
)

func usage() {
	_, err := fmt.Fprintf(os.Stderr, Usage)
	if err != nil {
		log.Fatalln(err)
		return
	}
	flag.PrintDefaults()
}

func init() {
	//Bind flags
	flag.BoolVar(&ifHelp, "h", false, "Show this screen.")
	flag.BoolVar(&ifVersion, "v", false, "Show version.")

	//Set logger
	utils.InitLogger()
}
func main() {
	flag.Parse()

	//Parse helps and version
	argsList, argNums := flag.Args(), flag.NArg()
	if ifHelp {
		usage()
	} else {
		if ifVersion {
			_, err := fmt.Fprintf(os.Stderr, version)
			if err != nil {
				log.Fatalln(err)
				return
			}
		}
	}
	//Wrong format handle
	if argNums != 2 {
		usage()
		return
	}
	//Parse the rest, trace and locate
	log.Printf("Args got,%s", utils.Slice2String(argsList[:]))

	if argNums != 2 && argsList[0] != "locate" && argsList[0] != "trace" && argsList[0] != "analysis" {
		usage()
	} //Check if argsList is valid
	if strings.ToLower(argsList[0]) == "analysis" {
		ifAnalysis = true
	} else {
		if strings.ToLower(argsList[0]) == "locate" {
			ifLocateFunction = true
		} else {
			if strings.ToLower(argsList[0]) == "trace" {
				ifTraceArgs = true
			}
		}
	} //Parse modes: Analysis, Locate, Trace

	//Parse <filename>, <function> or <function>:<arg>
	if ifAnalysis {
		targetFilename = argsList[1]
		log.Printf(`Args parsed:
Type:	Analysis
Target:	%s`, targetFilename)
		//NOTE: run analyzer here
		datapath := analysis.Analysis(targetFilename)
		log.Printf("Analysis successfully, data saved at %s", datapath)
		return

	} else {
		if ifLocateFunction {
			functionName = argsList[1]
			log.Printf(`Args parsed:
Type:	Locate
Target:	%s`, functionName)
			//Note: run Locator here
		} else {
			if ifTraceArgs {
				tmpSlice := strings.Split(argsList[1], ":") //argList[1] should be <function>:<arg>
				if len(tmpSlice) != 2 {
					log.Fatalf("Wrong args %s: should be <function>:<arg>", argsList[1])
				} else {
					traceFunctionName, traceArg = tmpSlice[0], tmpSlice[1]
					log.Printf(`Args parsed:
Type:				Trace
Target Function:	%s
Target Arg:			%s`, traceFunctionName, traceArg)
					//Note: run Tracer here
				}
			}

		}
	}
}
