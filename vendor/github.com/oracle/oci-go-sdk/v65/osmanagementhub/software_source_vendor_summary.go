// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwareSourceVendorSummary Provides summary information about a software source vendor, including name, operating system family, and architecture type.
type SoftwareSourceVendorSummary struct {

	// Name of the vendor providing the software source.
	Name VendorNameEnum `mandatory:"true" json:"name"`

	// List of corresponding operating system families.
	OsFamilies []OsFamilyEnum `mandatory:"true" json:"osFamilies"`

	// List of corresponding architecture types.
	ArchTypes []ArchTypeEnum `mandatory:"true" json:"archTypes"`
}

func (m SoftwareSourceVendorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftwareSourceVendorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVendorNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetVendorNameEnumStringValues(), ",")))
	}
	for _, val := range m.OsFamilies {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamilies: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	for _, val := range m.ArchTypes {
		if _, ok := GetMappingArchTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchTypes: %s. Supported values are: %s.", val, strings.Join(GetArchTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
