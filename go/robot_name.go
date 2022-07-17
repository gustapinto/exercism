package robotname

import (
	"errors"
	"math/rand"
	"time"
)

var Names = map[string]int{}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name := RandomStr("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2) + RandomStr("1234567890", 3)

	if _, exists := Names[name]; exists {
		return r.Name()
	}

	Names[name] += 1
	if Names[name] > 1 {
		return "", errors.New("exceeded maximum possible names for the robot")
	}

	r.name = name
	return name, nil
}

func (r *Robot) Reset() {
	r.name = ""
	r.name, _ = r.Name()
}

func RandomStr(chars string, length int) (str string) {
	for i := 0; i < length; i++ {
		str += Choice(chars)
	}

	return str
}

func Choice(str string) string {
	rand.Seed(time.Now().UnixMicro())

	randomNumber := rand.Intn(len(str) - 1)

	return string(str[randomNumber])
}
