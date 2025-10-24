// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dif

import (
	oci_dif "github.com/oracle/oci-go-sdk/v65/dif"
)

func deduplicateServices(services []oci_dif.ServiceEnum) []oci_dif.ServiceEnum {
	seen := make(map[oci_dif.ServiceEnum]bool)
	var result []oci_dif.ServiceEnum
	for _, s := range services {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}

func deduplicateTemplates(templates []oci_dif.StackTemplateEnum) []oci_dif.StackTemplateEnum {
	seen := make(map[oci_dif.StackTemplateEnum]bool)
	var result []oci_dif.StackTemplateEnum
	for _, t := range templates {
		if !seen[t] {
			seen[t] = true
			result = append(result, t)
		}
	}
	return result
}

func hasAdbArtifactDeployment(item map[string]interface{}) bool {
	if val, ok := item["artifact_object_storage_path"]; ok && val != nil {
		if s, ok := val.(string); ok && s != "" {
			return true
		}
	}
	if val, ok := item["db_credentials"]; ok && val != nil {
		if creds, ok := val.([]interface{}); ok && len(creds) > 0 {
			return true
		}
	}
	return false
}

func hasDataflowArtifactDeployment(item map[string]interface{}) bool {
	if val, ok := item["execute"]; ok && val != nil {
		if s, ok := val.(string); ok && s != "" {
			return true
		}
	}
	if val, ok := item["archive_uri"]; ok && val != nil {
		if _, ok := val.(string); ok {
			return true
		}
	}
	return false
}

func hasGgcsArtifactDeployment(item map[string]interface{}) bool {
	if val, ok := item["artifact_object_storage_path"]; ok && val != nil {
		if s, ok := val.(string); ok && s != "" {
			return true
		}
	}
	if val, ok := item["users"]; ok && val != nil {
		if l, ok := val.([]interface{}); ok && len(l) > 0 {
			return true
		}
	}
	if val, ok := item["sources"]; ok && val != nil {
		if l, ok := val.([]interface{}); ok && len(l) > 0 {
			return true
		}
	}
	if val, ok := item["targets"]; ok && val != nil {
		if l, ok := val.([]interface{}); ok && len(l) > 0 {
			return true
		}
	}
	return false
}
