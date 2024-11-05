package v1

import (
	"fmt"
	"path"
	"strings"

	"github.com/google/uuid"
)

const (
	ImageTemplate  = "%s/%s%s"
	DefaultSection = "user-upload"
)

func PrepareKeyForImageUpload(section string, filename string) string {
	if section == "" {
		section = DefaultSection
	}
	return fmt.Sprintf(ImageTemplate, section, uuid.New(), path.Ext(strings.ToLower(filename)))
}
