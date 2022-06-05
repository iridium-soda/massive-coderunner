package model

//To describe information about a single function

type FunctionInfo struct {
	Package  string
	Name     string
	Exported bool //If open to the external package
	Receiver []map[string]string
	Params   []map[string]string
	Results  []map[string]string // Return values
}
