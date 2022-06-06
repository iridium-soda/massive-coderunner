# Massive Roadrunner
基于`Funtrace`的针对RunC底层代码的分析工具。模块更新中。
# Description
用于静态分析RunC的源代码。在`functrace`的基础上，实现两个定位模块:
- 函数级的搜索模块：给定一个函数名、或者函数参数的名称后，能够通过AST定位到`runc`调用链中满足条件的目标函数（或者包含给定参数的函数）
- 函数内部的搜索模块：给定了一个函数的参数，能够跟踪函数内部对这个参数进行操作的语句
# Install
# Sample
For `Examples/a.go`:
```go
package Examples

func a(q int, w int, e int) int {
	return b(q, w, e)
}
func b(q1 int, q2 int, q3 int) int {
	return q1 + q2 + q3 + 1
}
func main() {
	a(1, 2, 3)

}
```
Analysis result is the following and saved in `data/Examples/a.json`
```json
[
  {
    "Package": "Examples",
    "Name": "a",
    "Exported": false,
    "Receiver": [],
    "Params": [
      {
        "Name": "q",
        "Type": "int"
      },
      {
        "Name": "w",
        "Type": "int"
      },
      {
        "Name": "e",
        "Type": "int"
      }
    ],
    "Results": [
      {
        "Name": "",
        "Type": "int"
      }
    ]
  },
  {
    "Package": "Examples",
    "Name": "b",
    "Exported": false,
    "Receiver": [],
    "Params": [
      {
        "Name": "q1",
        "Type": "int"
      },
      {
        "Name": "q2",
        "Type": "int"
      },
      {
        "Name": "q3",
        "Type": "int"
      }
    ],
    "Results": [
      {
        "Name": "",
        "Type": "int"
      }
    ]
  },
  {
    "Package": "Examples",
    "Name": "main",
    "Exported": false,
    "Receiver": [],
    "Params": [],
    "Results": []
  }
]
```
# Usage
To analysis and save and query structured AST tree.
```
To analysis and save and query structured AST tree.

Usage:
astool analysis <file>
astool locate <function>
astool trace <function>:<arg>
astool -h
astool -v

Options:
analysis <file>         :To analysis target file and gather useful information. 
locate <function>       :To locate function by name. Be sure all files you need to search were resolved successfully before.
trace <function>:<arg>  :To trace the statement inside the function that operates on this parameter.
-h                      :Show this screen.
-v                      :Show version.
```
分析后的文件存放在`./data/<packname>/<filename>.json`.
# Self-Compile
# Attention
- While inputting filepath, **DON'T** use `\` as separators. Use `/` instead.
# TODO
- [ ] Batch process files
- [ ] Better performance
- [ ] Limited workspace?
- [ ] Database added for analysis data?
- [ ] Bad AST handle
- [ ] ...
