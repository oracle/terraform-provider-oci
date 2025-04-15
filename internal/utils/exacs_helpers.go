package utils

func IsMultiVm(activatedStorageCount *int) bool {
	return activatedStorageCount != nil && *activatedStorageCount > 0
}
