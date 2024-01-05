// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ExportSet A set of file systems to export through one or more mount
// targets. Composed of zero or more export resources.
type ExportSet struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the export set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My export set`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the export set.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the export set.
	LifecycleState ExportSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	// Controls the maximum `tbytes`, `fbytes`, and `abytes`,
	// values reported by `NFS FSSTAT` calls through any associated
	// mount targets. This is an advanced feature. For most
	// applications, use the default value. The
	// `tbytes` value reported by `FSSTAT` will be
	// `maxFsStatBytes`. The value of `fbytes` and `abytes` will be
	// `maxFsStatBytes` minus the metered size of the file
	// system. If the metered size is larger than `maxFsStatBytes`,
	// then `fbytes` and `abytes` will both be '0'.
	MaxFsStatBytes *int64 `mandatory:"false" json:"maxFsStatBytes"`

	// Controls the maximum `tfiles`, `ffiles`, and `afiles`
	// values reported by `NFS FSSTAT` calls through any associated
	// mount targets. This is an advanced feature. For most
	// applications, use the default value. The
	// `tfiles` value reported by `FSSTAT` will be
	// `maxFsStatFiles`. The value of `ffiles` and `afiles` will be
	// `maxFsStatFiles` minus the metered size of the file
	// system. If the metered size is larger than `maxFsStatFiles`,
	// then `ffiles` and `afiles` will both be '0'.
	MaxFsStatFiles *int64 `mandatory:"false" json:"maxFsStatFiles"`
}

func (m ExportSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExportSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExportSetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportSetLifecycleStateEnum Enum with underlying type: string
type ExportSetLifecycleStateEnum string

// Set of constants representing the allowable values for ExportSetLifecycleStateEnum
const (
	ExportSetLifecycleStateCreating ExportSetLifecycleStateEnum = "CREATING"
	ExportSetLifecycleStateActive   ExportSetLifecycleStateEnum = "ACTIVE"
	ExportSetLifecycleStateDeleting ExportSetLifecycleStateEnum = "DELETING"
	ExportSetLifecycleStateDeleted  ExportSetLifecycleStateEnum = "DELETED"
)

var mappingExportSetLifecycleStateEnum = map[string]ExportSetLifecycleStateEnum{
	"CREATING": ExportSetLifecycleStateCreating,
	"ACTIVE":   ExportSetLifecycleStateActive,
	"DELETING": ExportSetLifecycleStateDeleting,
	"DELETED":  ExportSetLifecycleStateDeleted,
}

var mappingExportSetLifecycleStateEnumLowerCase = map[string]ExportSetLifecycleStateEnum{
	"creating": ExportSetLifecycleStateCreating,
	"active":   ExportSetLifecycleStateActive,
	"deleting": ExportSetLifecycleStateDeleting,
	"deleted":  ExportSetLifecycleStateDeleted,
}

// GetExportSetLifecycleStateEnumValues Enumerates the set of values for ExportSetLifecycleStateEnum
func GetExportSetLifecycleStateEnumValues() []ExportSetLifecycleStateEnum {
	values := make([]ExportSetLifecycleStateEnum, 0)
	for _, v := range mappingExportSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExportSetLifecycleStateEnumStringValues Enumerates the set of values in String for ExportSetLifecycleStateEnum
func GetExportSetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingExportSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportSetLifecycleStateEnum(val string) (ExportSetLifecycleStateEnum, bool) {
	enum, ok := mappingExportSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
