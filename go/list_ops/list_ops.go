package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	accumulator := initial

	for _, item := range s {
		accumulator = fn(accumulator, item)
	}

	return accumulator
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial

	for i := s.Length() - 1; i >= 0; i-- {
		accumulator = fn(s[i], accumulator)
	}

	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filtered := make(IntList, 0)

	for _, item := range s {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func (s IntList) Length() int {
	length := 0

	for _, _ = range s {
		length += 1
	}

	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	mapped := make(IntList, s.Length())

	for i, item := range s {
		mapped[i] = fn(item)
	}

	return mapped
}

func (s IntList) Reverse() IntList {
	index := s.Length() - 1
	reversed := make(IntList, 0)

	for i := range s {
		reversed = append(reversed, s[index-i])
	}

	return reversed
}

func (s IntList) Append(lst IntList) IntList {
	list := s

	for _, item := range lst {
		list = append(list, item)
	}

	return list
}

func (s IntList) Concat(lists []IntList) IntList {
	flattened := s

	for _, list := range lists {
		flattened = append(flattened, list...)
	}

	return flattened
}
