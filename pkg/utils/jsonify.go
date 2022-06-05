package utils

//To jsonify and dejsonify struct functionInfo
import (
	"encoding/json"
	"github.com/iridium-soda/massive-coderunner/pkg/model"
	"io/ioutil"
	"log"
)

func Jsonify(funcInfo []*model.FunctionInfo) []byte {
	res, err := json.Marshal(funcInfo)
	if err != nil {
		log.Fatalln(err)
		return nil
	} else {
		return res
	}
}
func Dejsonify(js []byte) []model.FunctionInfo {
	//See https://blog.csdn.net/jiuyanjin5740/article/details/88877314
	var plainData []model.FunctionInfo
	if err := json.Unmarshal(js, &plainData); err != nil {
		log.Panicln(err)
	}
	return plainData
}

func Test(filepath string) []model.FunctionInfo {
	//To read file and verify the decoded part. Delete after verification.
	dataEncoded, _ := ioutil.ReadFile(filepath)
	return Dejsonify(dataEncoded)
}
