// Package analysis
// @Title  analysis
// @Description
// @Reference https://www.codeleading.com/article/56096064668/
package analysis

//TODO：直接把代码抄过来
import (
	"fmt"
	"github.com/iridium-soda/massive-coderunner/pkg/model"
	"github.com/iridium-soda/massive-coderunner/pkg/utils"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

func Analysis(filePath string) string {
	// @title    Analysis
	// @description   analysis source code and generate structured data file
	// @param     filename:string file to be analysis
	// @return    path:string where data file saved
	fileSet := token.NewFileSet()

	fast, err := parser.ParseFile(fileSet, filePath, nil, 0) //NOTE: typeof(ast) is *ast.File
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	log.Printf("Package name is %s\n", fast.Name.Name)

	if !hasFuncDecl(fast) {
		log.Printf("%s has no function declaration", filePath)
		return ""
	}
	//Got Valid AST now
	var funcDeclsInfo = make([]*model.FunctionInfo, 0) //To save functions with struct functionInfo
	funcDeclsInfo = CreateASTFunctionsFromFile(filePath)
	printFunctionsInfo(funcDeclsInfo) //Only for debug
	jsonByte := utils.Jsonify(funcDeclsInfo)

	log.Printf("Info Jsonfied:\n%s", jsonByte)
	//Write file

	srcFileName := path.Base(filePath)                                           //srcFilename=src.go
	pureSrcFileName := srcFileName[0 : len(srcFileName)-len(path.Ext(filePath))] //pureFileName=src
	log.Printf("filePath is %s\nsrcFileName is %s\npureSrcFilename is %s", filePath, srcFileName, pureSrcFileName)
	dataFilePath := utils.WriteFile(pureSrcFileName, fast.Name.Name, jsonByte)

	if dataFilePath != "" {
		log.Println("Files generate successfully and name is:\t", dataFilePath)
	} else {
		log.Fatalln("File not generated.")
	}
	return dataFilePath
}

func createASTFunctionFromASTNode(node ast.Node, pack string) *model.FunctionInfo {
	//Check if it is function declaration and return function's info with struct
	fn, ok := node.(*ast.FuncDecl)
	if ok {
		astFunction := model.FunctionInfo{
			Package:  pack,
			Name:     fn.Name.Name,
			Exported: fn.Name.IsExported(),
		}
		astFunction.Params = parseFieldList(fn.Type.Params)
		astFunction.Results = parseFieldList(fn.Type.Results)
		astFunction.Receiver = parseFieldList(fn.Recv)
		return &astFunction
	}
	return nil
}
func parseFieldList(fList *ast.FieldList) []map[string]string {
	//Analysis field list(include para, results and receivers)
	dst := make([]map[string]string, 0)
	if fList != nil {
		list := fList.List
		for i := 0; i < len(list); i++ { //Traverse field list and append list
			names := list[i].Names
			typeStr := exprToTypeStringRecursively(list[i].Type)
			for j := 0; j < len(names); j++ {
				dst = append(dst, map[string]string{
					"Name": names[j].Name,
					"Type": typeStr,
				})
			}
			if len(names) == 0 {
				dst = append(dst, map[string]string{
					"Name": "",
					"Type": typeStr,
				})
			}
		}
	}
	return dst
}
func exprToTypeStringRecursively(expr ast.Expr) string {

	if arr, ok := expr.(*ast.ArrayType); ok {
		if arr.Len == nil {
			return "[]" + exprToTypeStringRecursively(arr.Elt)
		} else if lit, ok := arr.Len.(*ast.BasicLit); ok {
			return "[" + lit.Value + "]" + exprToTypeStringRecursively(arr.Elt)
		} else {
			// TODO 完备性检查

			panic(1)
		}
	}
	if _, ok := expr.(*ast.InterfaceType); ok {
		return "interface{}"
	}
	if indent, ok := expr.(*ast.Ident); ok {
		return indent.Name
	} else if selExpr, ok := expr.(*ast.SelectorExpr); ok {
		return exprToTypeStringRecursively(selExpr.X) + "." + exprToTypeStringRecursively(selExpr.Sel)
	} else if star, ok := expr.(*ast.StarExpr); ok {
		return "*" + exprToTypeStringRecursively(star.X)
	} else if mapType, ok := expr.(*ast.MapType); ok {
		return "map[" + exprToTypeStringRecursively(mapType.Key) + "]" + exprToTypeStringRecursively(mapType.Value)
	} else if funcType, ok := expr.(*ast.FuncType); ok {
		params := parseFieldList(funcType.Params)
		results := parseFieldList(funcType.Results)
		tf := func(data []map[string]string) string {
			ts := make([]string, 0)
			for _, v := range data {
				ts = append(ts, v["Name"]+" "+v["Type"])
			}
			return strings.Join(ts, ",")
		}
		return "func(" + tf(params) + ")" + " (" + tf(results) + ")"
	} else if chanType, ok := expr.(*ast.ChanType); ok {
		if chanType.Dir == ast.SEND {
			return "chan <- " + exprToTypeStringRecursively(chanType.Value)
		} else if chanType.Dir == ast.RECV {
			return "<- chan " + exprToTypeStringRecursively(chanType.Value)
		} else {
			return "chan " + exprToTypeStringRecursively(chanType.Value)
		}
	}
	//ast.StructType	不考虑这个类型

	//fmt.Println(expr)
	panic(1)
}
func hasFuncDecl(f *ast.File) bool {
	// @title    hasFuncDecl
	// @description   Check and return if the ast tree has any function declaration
	// @param     filename:string file to be analysis
	// @return    bool result
	if len(f.Decls) == 0 {
		return false
	}

	for _, decl := range f.Decls {
		_, ok := decl.(*ast.FuncDecl)
		if ok {
			return true
		}
	}

	return false
}
func printFunctionsInfo(funcDeclsInfo []*model.FunctionInfo) {
	//To print analysis result for type []*model.FunctionInfo, only for debug.
	/*type model.FunctionInfo struct {
		Package  string
		Name     string
		Exported bool //If open to the external package
		Receiver []map[string]string
		Params   []map[string]string
		Results  []map[string]string // Return values
	}
	*/
	log.Printf("Functions Count:%d\n", len(funcDeclsInfo))
	for i := 0; i < len(funcDeclsInfo); i++ {
		fmt.Printf(`-------------------
Package:	%s
Name:		%s
IfExported:	%v
`, funcDeclsInfo[i].Package, funcDeclsInfo[i].Name, funcDeclsInfo[i].Exported)
		fmt.Println("Receiver:", funcDeclsInfo[i].Receiver)
		fmt.Println("Paras:", funcDeclsInfo[i].Params)
		fmt.Println("Returns:", funcDeclsInfo[i].Results)
		fmt.Println("-------------------")
	}
}
func CreateASTFunctionsFromFile(target string) []*model.FunctionInfo {
	_, err := ioutil.ReadFile(target)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, target, nil, 0)
	if err != nil {
		log.Fatalln(err)

	}
	pack := ""
	//ast.Print(fileSet, file)
	functions := make([]*model.FunctionInfo, 0)
	ast.Inspect(file, func(node ast.Node) bool {
		pk, ok := node.(*ast.Ident)
		if ok {
			if pack == "" {
				pack = pk.Name
			}
		}
		fn := createASTFunctionFromASTNode(node, pack)
		if fn != nil {
			functions = append(functions, fn)
		}
		return true
	})
	return functions
}
