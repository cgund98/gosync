package sync

import (
	"fmt"

	c "github.com/cgund98/gosync/internal/config"
)

type SyncFunc func(source, destination string, excludePatterns []string)

func MapSyncTypeToFunc(t string) (SyncFunc, error) {

	if t == c.DestinationLocal {
		return SyncToLocal, nil
	}

	if t == c.DestinationS3 {
		return SyncToS3, nil
	}

	return nil, fmt.Errorf("no sync function found for type: '%s'", t)
}
