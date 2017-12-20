// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

func resourceObjectStorageMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}
