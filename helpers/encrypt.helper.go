package helpers

import (
	"time"

	utils "gitlab.finema.co/finema/idin-core/utils"
)

func GenerateHash() string {
	return utils.NewSha256(time.Now().String())
}
