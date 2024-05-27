package deepmap

func DeepInsert(dest map[string]any, keys []string, value any) map[string]any {
	keys_len := len(keys)
	// return as-is if no keys are passed
	if keys_len == 0 {
		return dest
	}

	last_index := len(keys) - 1
	for index, key := range keys {
		// if current key is last index -> just set the value
		if index == last_index {
			dest[key] = value
		} else {
			existing, has := dest[key]
			if !has {
				// not the last key and current key does not exist yet
				// create new value and call itself
				dest[key] = DeepInsert(make(map[string]any), keys[index+1:], value)
			} else {
				// not the last key and current key already exists
				// use existing value if map
				asserted, assertion_ok := existing.(map[string]any)
				if assertion_ok {
					dest[key] = DeepInsert(asserted, keys[index+1:], value)
				}
			}
		}
	}

	return dest
}
