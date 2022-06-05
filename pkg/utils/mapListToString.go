package utils

//To convert []map[string]string to printable string
func mapList2String(mps []map[string]string) string {
	res := ""
	for i := range mps {
		for k, v := range mps[i] {
			res += k + ":" + v
		}
	}
	return res
}
