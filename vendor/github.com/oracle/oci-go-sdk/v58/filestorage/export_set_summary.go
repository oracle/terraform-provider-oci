// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ExportSetSummary Summary information for an export set.
type ExportSetSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the export set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My export set`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the export set.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the export set.
	LifecycleState ExportSetSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the export set was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the virtual cloud network (VCN) the export set is in.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The availability domain the export set is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
}

func (m ExportSetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportSetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExportSetSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExportSetSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportSetSummaryLifecycleStateEnum Enum with underlying type: string
type ExportSetSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExportSetSummaryLifecycleStateEnum
const (
	ExportSetSummaryLifecycleStateCreating ExportSetSummaryLifecycleStateEnum = "CREATING"
	ExportSetSummaryLifecycleStateActive   ExportSetSummaryLifecycleStateEnum = "ACTIVE"
	ExportSetSummaryLifecycleStateDeleting ExportSetSummaryLifecycleStateEnum = "DELETING"
	ExportSetSummaryLifecycleStateDeleted  ExportSetSummaryLifecycleStateEnum = "DELETED"
)

var mappingExportSetSummaryLifecycleStateEnum = map[string]ExportSetSummaryLifecycleStateEnum{
	"CREATING": ExportSetSummaryLifecycleStateCreating,
	"ACTIVE":   ExportSetSummaryLifecycleStateActive,
	"DELETING": ExportSetSummaryLifecycleStateDeleting,
	"DELETED":  ExportSetSummaryLifecycleStateDeleted,
}

// GetExportSetSummaryLifecycleStateEnumValues Enumerates the set of values for ExportSetSummaryLifecycleStateEnum
func GetExportSetSummaryLifecycleStateEnumValues() []ExportSetSummaryLifecycleStateEnum {
	values := make([]ExportSetSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExportSetSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExportSetSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExportSetSummaryLifecycleStateEnum
func GetExportSetSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingExportSetSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportSetSummaryLifecycleStateEnum(val string) (ExportSetSummaryLifecycleStateEnum, bool) {
	mappingExportSetSummaryLifecycleStateEnumIgnoreCase := make(map[string]ExportSetSummaryLifecycleStateEnum)
	for k, v := range mappingExportSetSummaryLifecycleStateEnum {
		mappingExportSetSummaryLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExportSetSummaryLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
