// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateLustreFileSystemDetails The data to update a Lustre file system.
type UpdateLustreFileSystemDetails struct {

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My Lustre file system`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Short description of the Lustre file system.
	// Avoid entering confidential information.
	FileSystemDescription *string `mandatory:"false" json:"fileSystemDescription"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A list of Network Security Group OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this lustre file system.
	// A maximum of 5 is allowed.
	// Setting this to an empty array after the list is created removes the lustre file system from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key used to encrypt the encryption keys associated with this file system.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// Capacity of the Lustre file system in GB. You can increase capacity only in multiples of 5 TB.
	CapacityInGBs *int `mandatory:"false" json:"capacityInGBs"`

	RootSquashConfiguration *RootSquashConfiguration `mandatory:"false" json:"rootSquashConfiguration"`
}

func (m UpdateLustreFileSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLustreFileSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
