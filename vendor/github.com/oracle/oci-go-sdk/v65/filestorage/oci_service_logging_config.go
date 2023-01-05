// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciServiceLoggingConfig Log resource configuration
type OciServiceLoggingConfig struct {

	// Log resource OCID.
	LogId *string `mandatory:"true" json:"logId"`

	// Resource OCID. Log would be configured for this resource.
	Resource *string `mandatory:"true" json:"resource"`

	// Name of the category.
	Category *string `mandatory:"true" json:"category"`

	// The OCID of the tenancy this log belongs to.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// Parameters for the category.
	Parameters map[string]string `mandatory:"false" json:"parameters"`

	// The current state of the logging resource in the context of the target resource. This field must be present in the context of a response but must not be present in the context of a request.
	LifecycleState OciServiceLoggingConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m OciServiceLoggingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciServiceLoggingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOciServiceLoggingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciServiceLoggingConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciServiceLoggingConfigLifecycleStateEnum Enum with underlying type: string
type OciServiceLoggingConfigLifecycleStateEnum string

// Set of constants representing the allowable values for OciServiceLoggingConfigLifecycleStateEnum
const (
	OciServiceLoggingConfigLifecycleStateCreating OciServiceLoggingConfigLifecycleStateEnum = "CREATING"
	OciServiceLoggingConfigLifecycleStateActive   OciServiceLoggingConfigLifecycleStateEnum = "ACTIVE"
	OciServiceLoggingConfigLifecycleStateUpdating OciServiceLoggingConfigLifecycleStateEnum = "UPDATING"
	OciServiceLoggingConfigLifecycleStateDeleting OciServiceLoggingConfigLifecycleStateEnum = "DELETING"
	OciServiceLoggingConfigLifecycleStateDeleted  OciServiceLoggingConfigLifecycleStateEnum = "DELETED"
	OciServiceLoggingConfigLifecycleStateFailed   OciServiceLoggingConfigLifecycleStateEnum = "FAILED"
)

var mappingOciServiceLoggingConfigLifecycleStateEnum = map[string]OciServiceLoggingConfigLifecycleStateEnum{
	"CREATING": OciServiceLoggingConfigLifecycleStateCreating,
	"ACTIVE":   OciServiceLoggingConfigLifecycleStateActive,
	"UPDATING": OciServiceLoggingConfigLifecycleStateUpdating,
	"DELETING": OciServiceLoggingConfigLifecycleStateDeleting,
	"DELETED":  OciServiceLoggingConfigLifecycleStateDeleted,
	"FAILED":   OciServiceLoggingConfigLifecycleStateFailed,
}

var mappingOciServiceLoggingConfigLifecycleStateEnumLowerCase = map[string]OciServiceLoggingConfigLifecycleStateEnum{
	"creating": OciServiceLoggingConfigLifecycleStateCreating,
	"active":   OciServiceLoggingConfigLifecycleStateActive,
	"updating": OciServiceLoggingConfigLifecycleStateUpdating,
	"deleting": OciServiceLoggingConfigLifecycleStateDeleting,
	"deleted":  OciServiceLoggingConfigLifecycleStateDeleted,
	"failed":   OciServiceLoggingConfigLifecycleStateFailed,
}

// GetOciServiceLoggingConfigLifecycleStateEnumValues Enumerates the set of values for OciServiceLoggingConfigLifecycleStateEnum
func GetOciServiceLoggingConfigLifecycleStateEnumValues() []OciServiceLoggingConfigLifecycleStateEnum {
	values := make([]OciServiceLoggingConfigLifecycleStateEnum, 0)
	for _, v := range mappingOciServiceLoggingConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOciServiceLoggingConfigLifecycleStateEnumStringValues Enumerates the set of values in String for OciServiceLoggingConfigLifecycleStateEnum
func GetOciServiceLoggingConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOciServiceLoggingConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciServiceLoggingConfigLifecycleStateEnum(val string) (OciServiceLoggingConfigLifecycleStateEnum, bool) {
	enum, ok := mappingOciServiceLoggingConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
