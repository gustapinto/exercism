package flatten

func flatten(item interface{}, flattened []interface{}) []interface{} {
	slice, isSlice := item.([]interface{})

	if isSlice {
		for _, elem := range slice {
			flattened = flatten(elem, flattened)
		}
	} else {
		if item != nil {
			flattened = append(flattened, item)
		}
	}

	return flattened
}

func Flatten(nested interface{}) []interface{} {
	flattened := flatten(nested, make([]interface{}, 0))

	return flattened
}
