package main

import (
	"fmt"
)

func STDerrhandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*func safeIntConvert(r []string, n int) int {
	if len(r) > n {
		if r[n] == "" {
			return 0
		}
		querschnittClean := strings.Replace(r[n], ",", "", -1)
		querschnitt, _ := strconv.Atoi(querschnittClean)
		return querschnitt
	}
	return 0
}
*/
