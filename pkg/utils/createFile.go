package utils

import (
	"io/ioutil"
	"log"
	"os"
)

// WriteFile
/*To create file with formats and return file handle*/
//filename specification:./data/[PackageName]-[SourcecodeFilename].json
//Example: examples-a.json
//Reference: https://laravelacademy.org/post/21943
//Reference: https://zhuanlan.zhihu.com/p/80403583
func WriteFile(sourceCodeName string, packageName string, jsonByte []byte) string { //NOTE: change params to pointer to improve the performance
	//Generate functionName
	if jsonByte == nil {
		log.Fatalln("Functions Info is empty and nothing to be written.")
		return ""
	}

	//Get Package name
	err := os.MkdirAll("data/"+packageName, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	filePath := "./data/" + packageName + "/" + sourceCodeName + ".json" //./data/[PackageName]-[SourcecodeFilename].json
	//Write file
	filePointer, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer filePointer.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = ioutil.WriteFile(filePath, jsonByte, 0666)
	if err != nil {
		log.Panicln(err)
	}
	return filePath

}
func ReadFile(filePath string) []byte {
	jsonByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return jsonByte
}
