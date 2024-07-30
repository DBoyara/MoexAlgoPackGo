package utils

func MapData(keys []string, values [][]interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	for _, v := range values {
		m := make(map[string]interface{})
		for i := 0; i < len(keys); i++ {
			m[keys[i]] = v[i]
		}
		result = append(result, m)
	}
	return result
}
