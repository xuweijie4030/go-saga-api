package global

import "encoding/json"

func JsonMarshalToStr(param interface{}) (result string, err error) {
	b, err := json.Marshal(param)
	if err != nil {
		return result, err
	}
	result = string(b)

	return result, nil
}

func StructToMap(arg interface{}) (result map[string]interface{}, err error) {
	tempResult, err := JsonMarshalToStr(arg)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal([]byte(tempResult), &result); err != nil {
		return result, err
	}

	return result, nil
}
