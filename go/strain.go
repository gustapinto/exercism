package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}

	keeps := make(Ints, 0)

	for _, item := range i {
		if filter(item) {
			keeps = append(keeps, item)
		}
	}

	return keeps
}

func (i Ints) Discard(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}

	discards := make(Ints, 0)

	for _, item := range i {
		if !filter(item) {
			discards = append(discards, item)
		}
	}

	return discards
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return nil
	}

	keeps := make(Lists, 0)

	for _, list := range l {
		if filter(list) {
			keeps = append(keeps, list)
		}
	}

	return keeps
}

func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return nil
	}

	keeps := make(Strings, 0)

	for _, item := range s {
		if filter(item) {
			keeps = append(keeps, item)
		}
	}

	return keeps
}
