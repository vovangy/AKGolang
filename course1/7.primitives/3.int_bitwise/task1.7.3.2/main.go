package main

import "fmt"

const (
	Read    = 4
	Write   = 2
	Execute = 1
)

func getFilePermissions(flag int) string {
	otherRules := flag % 10
	flag /= 10
	groupRules := flag % 10
	flag /= 10
	ownerRules := flag % 10

	var rules map[string]int = make(map[string]int, 3)
	rules["Owner:"] = ownerRules
	rules["Group:"] = groupRules
	rules["Other:"] = otherRules

	result := ""
	end := " "
	for _, val := range []string{"Owner:", "Group:", "Other:"} {
		result += val
		if Read&rules[val] == Read {
			result += "Read,"
		} else {
			result += "-,"
		}

		if Write&rules[val] == Write {
			result += "Write,"
		} else {
			result += "-,"
		}

		if val == "Other:" {
			end = ""
		}

		if Execute&rules[val] == Execute {
			result += "Execute" + end
		} else {
			result += "-" + end
		}
	}

	return result
}

func main() {
	fmt.Println(getFilePermissions(777))
}
