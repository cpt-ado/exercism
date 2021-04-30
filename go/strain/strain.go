package strain

// Required types
type Ints []int
type Lists [][]int
type Strings []string

// Required methods
func (object Ints) Keep(condition func(int) bool) Ints {
	if object == nil {
		return nil
	}
	result := make(Ints, 0)
	for _, element := range object {
		if condition(element) {
			result = append(result, element)
		}
	}
	return result
}

func (object Ints) Discard(condition func(int) bool) Ints {
	if object == nil {
		return nil
	}
	result := make(Ints, 0)
	for _, element := range object {
		if !condition(element) {
			result = append(result, element)
		}
	}
	return result
}

func (object Lists) Keep(condition func([]int) bool) Lists {
	if object == nil {
		return nil
	}
	result := make(Lists, 0)
	for _, element := range object {
		if condition(element) {
			result = append(result, element)
		}
	}
	return result
}

func (object Strings) Keep(condition func(string) bool) Strings {
	if object == nil {
		return nil
	}
	result := make(Strings, 0)
	for _, element := range object {
		if condition(element) {
			result = append(result, element)
		}
	}
	return result
}

//    (Strings) Keep(func(string) bool) Strings
