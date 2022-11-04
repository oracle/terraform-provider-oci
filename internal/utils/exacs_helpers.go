package utils

import "strings"

func IsMultiVm(shape string, maxDataStorageInTBs *float64) bool {
	return (strings.EqualFold(shape, "Exadata.X8M") || strings.EqualFold(shape, "Exadata.X9M")) && maxDataStorageInTBs != nil
}
