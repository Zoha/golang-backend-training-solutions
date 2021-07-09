package formletter

import (
	"fmt"

	cfg "github.com/zoha/golang-backend-training-solutions/01.outputting/01.09.form-letter-package/cfg"
)

func Formletter() {
	fmt.Printf("Hello %s!\nThe weather today is %s\nToday's snack will be %s", cfg.Name, cfg.Weather, cfg.Snack)
}
