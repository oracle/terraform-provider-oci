// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BigDataBmToVmInstanceMigration BigDataBmToVmInstanceMigration Details
type BigDataBmToVmInstanceMigration struct {

	// The OCID of the compartment that bm instance that customer is migrating.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the BigDataBmToVmInstanceMigration event that was created to migrate a BM Instance.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the BM instance that BigDataBmToVmInstanceMigration event migrating from.
	SourceInstanceId *string `mandatory:"true" json:"sourceInstanceId"`

	// The current BigDataBmToVmInstanceMigration state.
	LifecycleState BigDataBmToVmInstanceMigrationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the VM instance that BigDataBmToVmInstanceMigration event migrated to.
	TargetInstanceId *string `mandatory:"false" json:"targetInstanceId"`

	// The image OCID.
	TargetImageId *string `mandatory:"false" json:"targetImageId"`

	// The shape name.
	TargetShape *string `mandatory:"false" json:"targetShape"`

	// Custom metadata key/value string pairs that you provide. Any set of key/value pairs
	// provided here will completely replace the current set of key/value pairs in the `metadata`
	// field on the instance.
	// The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of
	// 32,000 bytes.
	TargetMetadata map[string]string `mandatory:"false" json:"targetMetadata"`

	// Additional metadata key/value pairs that you provide. They serve the same purpose and
	// functionality as fields in the `metadata` object.
	// They are distinguished from `metadata` fields in that these can be nested JSON objects
	// (whereas `metadata` fields are string/string maps only).
	// The "user_data" field and the "ssh_authorized_keys" field cannot be changed after an instance
	// has launched. Any request that updates, removes, or adds either of these fields will be
	// rejected. You must provide the same values for "user_data" and "ssh_authorized_keys" that
	// already exist on the instance.
	// The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of
	// 32,000 bytes.
	TargetExtendedMetadata map[string]interface{} `mandatory:"false" json:"targetExtendedMetadata"`
}

func (m BigDataBmToVmInstanceMigration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BigDataBmToVmInstanceMigration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBigDataBmToVmInstanceMigrationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBigDataBmToVmInstanceMigrationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BigDataBmToVmInstanceMigrationLifecycleStateEnum Enum with underlying type: string
type BigDataBmToVmInstanceMigrationLifecycleStateEnum string

// Set of constants representing the allowable values for BigDataBmToVmInstanceMigrationLifecycleStateEnum
const (
	BigDataBmToVmInstanceMigrationLifecycleStateAccepted       BigDataBmToVmInstanceMigrationLifecycleStateEnum = "ACCEPTED"
	BigDataBmToVmInstanceMigrationLifecycleStateInProgress     BigDataBmToVmInstanceMigrationLifecycleStateEnum = "IN_PROGRESS"
	BigDataBmToVmInstanceMigrationLifecycleStateSucceeded      BigDataBmToVmInstanceMigrationLifecycleStateEnum = "SUCCEEDED"
	BigDataBmToVmInstanceMigrationLifecycleStateFailed         BigDataBmToVmInstanceMigrationLifecycleStateEnum = "FAILED"
	BigDataBmToVmInstanceMigrationLifecycleStateNeedsAttention BigDataBmToVmInstanceMigrationLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingBigDataBmToVmInstanceMigrationLifecycleStateEnum = map[string]BigDataBmToVmInstanceMigrationLifecycleStateEnum{
	"ACCEPTED":        BigDataBmToVmInstanceMigrationLifecycleStateAccepted,
	"IN_PROGRESS":     BigDataBmToVmInstanceMigrationLifecycleStateInProgress,
	"SUCCEEDED":       BigDataBmToVmInstanceMigrationLifecycleStateSucceeded,
	"FAILED":          BigDataBmToVmInstanceMigrationLifecycleStateFailed,
	"NEEDS_ATTENTION": BigDataBmToVmInstanceMigrationLifecycleStateNeedsAttention,
}

var mappingBigDataBmToVmInstanceMigrationLifecycleStateEnumLowerCase = map[string]BigDataBmToVmInstanceMigrationLifecycleStateEnum{
	"accepted":        BigDataBmToVmInstanceMigrationLifecycleStateAccepted,
	"in_progress":     BigDataBmToVmInstanceMigrationLifecycleStateInProgress,
	"succeeded":       BigDataBmToVmInstanceMigrationLifecycleStateSucceeded,
	"failed":          BigDataBmToVmInstanceMigrationLifecycleStateFailed,
	"needs_attention": BigDataBmToVmInstanceMigrationLifecycleStateNeedsAttention,
}

// GetBigDataBmToVmInstanceMigrationLifecycleStateEnumValues Enumerates the set of values for BigDataBmToVmInstanceMigrationLifecycleStateEnum
func GetBigDataBmToVmInstanceMigrationLifecycleStateEnumValues() []BigDataBmToVmInstanceMigrationLifecycleStateEnum {
	values := make([]BigDataBmToVmInstanceMigrationLifecycleStateEnum, 0)
	for _, v := range mappingBigDataBmToVmInstanceMigrationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBigDataBmToVmInstanceMigrationLifecycleStateEnumStringValues Enumerates the set of values in String for BigDataBmToVmInstanceMigrationLifecycleStateEnum
func GetBigDataBmToVmInstanceMigrationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingBigDataBmToVmInstanceMigrationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBigDataBmToVmInstanceMigrationLifecycleStateEnum(val string) (BigDataBmToVmInstanceMigrationLifecycleStateEnum, bool) {
	enum, ok := mappingBigDataBmToVmInstanceMigrationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
