package utils

import (
	"github.com/anandvarma/namegen"
)

func GenerateUniqueTag() string {
	ngen := namegen.New()
	return ngen.Get()
}
