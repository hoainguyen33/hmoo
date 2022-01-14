package validate

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/hoainguyen33/Go/bson"
	"github.com/hoainguyen33/Go/common/errors"
	"github.com/hoainguyen33/Go/e2e"
)

// format query []string to int
func QueryAtoi(query []string) (int, error) {
	if len(query) > 0 {
		i, err := strconv.Atoi(query[0])
		return i, err
	}
	return 0, errors.ErrQueryInvalid
}

// format query request to struct
func QueryModel(query url.Values, model interface{}) error {
	var req = map[string]string{}
	for i, s := range query {
		req[i] = s[0]
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		return errors.ErrQueryInvalid
	}
	json.Unmarshal(marshal, model)
	return nil
}

// only return bsonD
func QueryFilter(query url.Values, filter map[string]string) bson.D {
	props := bson.M{}
	for k, t := range filter {
		if query[k] != nil {
			if t == "regex" {
				props[k] = bson.Regex(NewPatternVNCode(query[k][0]), "ig")
			} else {
				props[k] = bson.M{"$in": FormatArrString(strings.Split(query[k][0], ","), t)}
			}
		}
	}
	return bson.D{bson.E("$match", props)}
}

// format sorter return nil or bsonD
func QuerySort(query url.Values, sort []string, sorter map[string]int) bson.D {
	props := bson.M{}
	for _, k := range sort {
		if query[k] != nil && sorter[query[k][0]] != 0 {
			props[k] = sorter[query[k][0]]
		}
	}
	if len(props) == 0 {
		return nil
	}
	return bson.D{bson.E("$sort", props)}
}

// format []string with type(string)
func FormatArrString(strs []string, Type string) interface{} {
	switch Type {
	case "bool":
		var result = []bool{}
		for _, v := range strs {
			result = append(result, StringToBool(v))
		}
		return result
	case "int":
		var result = []int{}
		for _, v := range strs {
			result = append(result, StringToInt(v))
		}
		return result
	default:
		return strs
	}
}

// change strRegex to strRegexVn
func NewPatternVNCode(strRegex string) string {
	pattern := strRegex
	for i, s := range e2e.VN_code {
		pattern = strings.ReplaceAll(pattern, i, s)
	}
	return pattern
}

// format string to bool
func StringToBool(str string) bool {
	if str == "true" {
		return true
	}
	return false
}

// format string to int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
