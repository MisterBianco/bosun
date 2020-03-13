package utils

import (
	"fmt"

	"github.com/blang/semver"
)

func GenerateTags(name string, baseTags []string) []string {
	result := []string{}

	for j := 0; j < len(baseTags); j++ {
		v, err := semver.Parse(baseTags[j])
		if err != nil {
			fmt.Println("Its likely your tags arent configured properly.")
			fmt.Println("Your tags should follow semver. https://semver.org/")
			fmt.Println("Offending tag:", baseTags[j])
			result = append(result, baseTags[j])
			continue
		}

		if v.Pre == nil {
			result = append(result, fmt.Sprintf("%s:%v", name,v.Major))
			result = append(result, fmt.Sprintf("%s:%v.%v", name,v.Major, v.Minor))
			result = append(result, fmt.Sprintf("%s:%v.%v.%v", name,v.Major, v.Minor, v.Patch))
		} else {
			result = append(result, fmt.Sprintf("%s:%v-%s", name, v.Major, v.Pre[0]))
			result = append(result, fmt.Sprintf("%s:%v.%v-%s", name,v.Major, v.Minor, v.Pre[0]))
			result = append(result, fmt.Sprintf("%s:%v.%v.%v-%s", name,v.Major, v.Minor, v.Patch, v.Pre[0]))
		}
	}
	result = append(result, fmt.Sprintf("%s:latest", name))

	return result
}