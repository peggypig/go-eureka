/**
 * @Desc
 * @author zjhfyq 
 * @data 2018/4/12 17:43.
 */
package common


func Remove(s []string, re string) []string {
	i := 0
	for index, value := range s {
		if value == re {
			i = index
			break
		}
	}
	return append(s[:i], s[i+1:]...)
}
