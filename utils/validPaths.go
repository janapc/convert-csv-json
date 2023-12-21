package utils

import (
	"errors"
	"regexp"
)

func ValidPaths(paths map[string]string) error {
	if paths["json"] != "" {
		rgxJson, err := regexp.MatchString(`.*\.json$`, paths["json"])
		if err != nil || !rgxJson {
			return errors.New("format invalid")
		}
	}

	if paths["csv"] != "" {
		rgxJson, err := regexp.MatchString(`.*\.csv$`, paths["csv"])
		if err != nil || !rgxJson {
			return errors.New("format invalid")
		}
	}
	return nil
}
