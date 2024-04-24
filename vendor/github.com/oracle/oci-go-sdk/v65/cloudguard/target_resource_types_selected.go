// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetResourceTypesSelected Target selection on basis of target resource types.
type TargetResourceTypesSelected struct {

	// Types of targets
	Values []TargetResourceTypeEnum `mandatory:"false" json:"values,omitempty"`
}

func (m TargetResourceTypesSelected) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetResourceTypesSelected) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Values {
		if _, ok := GetMappingTargetResourceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Values: %s. Supported values are: %s.", val, strings.Join(GetTargetResourceTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TargetResourceTypesSelected) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTargetResourceTypesSelected TargetResourceTypesSelected
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeTargetResourceTypesSelected
	}{
		"TARGETTYPES",
		(MarshalTypeTargetResourceTypesSelected)(m),
	}

	return json.Marshal(&s)
}
