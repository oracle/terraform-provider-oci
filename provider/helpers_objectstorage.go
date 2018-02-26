// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"strings"
)

func resourceObjectStorageMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

// The SDK will return all 'metadata' header keys as lowercase, regardless of how they're specified in the config.
//
// To avoid unnecessary diffs and updates, we need to ensure all config keys for 'metadata' are lowercase.
// This avoids a conflict where our config has a non-lowercase key (e.g. MyKey) while the state file has a lowercase
// key (e.g. mykey) from the SDK.
//
// If we ran a 'terraform plan' on this without any config changes, Terraform will detect a diff between state and
// config; even though nothing changed in the state file.
func validateLowerCaseKeysInMetadata(raw interface{}, fieldName string) ([]string, []error) {
	metadataMap, ok := raw.(map[string]interface{})
	if !ok {
		return nil, []error{fmt.Errorf("Could not convert '%s' to map during validation.", fieldName)}
	}

	errors := []error{}
	for key := range metadataMap {
		lowercaseKey := strings.ToLower(key)
		if key != lowercaseKey {
			errors = append(errors, fmt.Errorf("All '%s' keys must be lowercase. Please specify '%s' as '%s'", fieldName, key, lowercaseKey))
		}
	}

	return nil, errors
}
