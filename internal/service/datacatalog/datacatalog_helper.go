// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func propertiesDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old != "" && new != "" {
		return false
	}

	// Find the specific defined_tag key name (mainly if a resource supports tagging at multiple levels)
	// For example: "create_vnic_details.0.properties.mynamespace.mykey" => "create_vnic_details.0.properties"
	keyParts := strings.Split(key, ".")
	propertiesKeyParts := []string{}
	for _, keyPart := range keyParts {
		propertiesKeyParts = append(propertiesKeyParts, keyPart)
		if strings.EqualFold(keyPart, "properties") {
			break
		}
	}

	//Old value comes from refreshed state, while new value comes from config
	oldRaw, newRaw := d.GetChange(strings.Join(propertiesKeyParts, "."))
	if newRaw == nil || oldRaw == nil {
		return false
	}

	newValue, newValueOk := newRaw.(map[string]interface{})
	oldValue, oldValueOk := oldRaw.(map[string]interface{})
	if !newValueOk || !oldValueOk {
		return false
	}

	lowerCaseNewValueMap := tfresource.ToLowerCaseKeyMap(newValue)
	lowerCaseOldValueMap := tfresource.ToLowerCaseKeyMap(oldValue)

	if reflect.DeepEqual(lowerCaseOldValueMap, lowerCaseNewValueMap) {
		return true
	}
	return false
}

func encPropertiesDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old != "" && new != "" {
		return false
	}

	// Find the specific defined_tag key name (mainly if a resource supports tagging at multiple levels)
	// For example: "create_vnic_details.0.enc_properties.mynamespace.mykey" => "create_vnic_details.0.enc_properties"
	keyParts := strings.Split(key, ".")
	propertiesKeyParts := []string{}
	for _, keyPart := range keyParts {
		propertiesKeyParts = append(propertiesKeyParts, keyPart)
		if strings.EqualFold(keyPart, "enc_properties") {
			break
		}
	}

	//Old value comes from refreshed state, while new value comes from config
	oldRaw, newRaw := d.GetChange(strings.Join(propertiesKeyParts, "."))
	if newRaw == nil || oldRaw == nil {
		return false
	}

	newValue, newValueOk := newRaw.(map[string]interface{})
	oldValue, oldValueOk := oldRaw.(map[string]interface{})
	if !newValueOk || !oldValueOk {
		return false
	}

	lowerCaseNewValueMap := tfresource.ToLowerCaseKeyMap(newValue)
	lowerCaseOldValueMap := tfresource.ToLowerCaseKeyMap(oldValue)

	if reflect.DeepEqual(lowerCaseOldValueMap, lowerCaseNewValueMap) {
		return true
	}
	return false
}
