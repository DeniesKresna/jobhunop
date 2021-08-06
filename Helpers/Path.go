package Helpers

import (
	"os"
	"strings"
)

func GetAppDomain() string {
	return os.Getenv("DOMAIN")
}

func DeleteFile(filepath string) error {
	err := os.Remove("./" + filepath)
	if err != nil {
		return err
	}
	return nil
}

func AcademyPath(additional string) string {
	return strings.Join([]string{"Assets/Academies/", additional}, "")
}
