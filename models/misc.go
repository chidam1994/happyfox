package models

import "fmt"

type Filter string

const (
	NameFilter  Filter = "name"
	EmailFilter Filter = "email"
	PhoneFilter Filter = "phnum"
)

func GetFilter(filterStr string) (Filter, error) {
	filterMap := map[string]Filter{
		"name":  NameFilter,
		"email": EmailFilter,
		"phnum": PhoneFilter,
	}
	filter, ok := filterMap[filterStr]
	if !ok {
		return Filter(""), fmt.Errorf("error converting string %s to filter", filterStr)
	}
	return filter, nil
}
