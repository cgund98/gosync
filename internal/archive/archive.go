package archive

import (
	"fmt"

	c "github.com/cgund98/gosync/internal/config"
)

type ArchiveFunc = func(source, destination string)

func MapArchiveTypeToFunc(t string) (ArchiveFunc, error) {

	if t == c.DestinationS3Glacier {
		return ArchiveS3Glacier, nil
	}

	return nil, fmt.Errorf("no archive function found for type: '%s'", t)
}
