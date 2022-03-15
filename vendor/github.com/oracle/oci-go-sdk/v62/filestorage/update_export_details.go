// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// UpdateExportDetails Details for updating the export.
type UpdateExportDetails struct {

	// The export is modified to include a boolean to use ID mapping for Unix Groups rather than the group list provided within an NFS Request's RPC header. When this flag is true the Unix UID from the RPC header is used to retrieve the list of secondary groups from a the ID mapping subsystem. The primary GID is always taken from the RPC header. If ID mapping is not configured, incorrectly configured, unavailable or cannot be used to determine a list of secondary groups then the data path uses an empty secondary group list for authorization. If the number of groups exceeds the current limit of 256 groups the list retrieved from LDAP is truncated to the first 256 groups read.
	IsIdmapGroupsForSysAuth *bool `mandatory:"false" json:"isIdmapGroupsForSysAuth"`

	// Export can be set in 'ENABLED' or 'DISABLED' mode.
	// Attempt to mount the filesystem will fail if the export is in 'DISABLED' mode.
	ExportMode UpdateExportDetailsExportModeEnum `mandatory:"false" json:"exportMode,omitempty"`

	// New export options for the export.
	// **Setting to the empty array will make the export invisible to all clients.**
	// Leaving unset will leave the `exportOptions` unchanged.
	ExportOptions []ClientOptions `mandatory:"false" json:"exportOptions"`
}

func (m UpdateExportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateExportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateExportDetailsExportModeEnum(string(m.ExportMode)); !ok && m.ExportMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExportMode: %s. Supported values are: %s.", m.ExportMode, strings.Join(GetUpdateExportDetailsExportModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateExportDetailsExportModeEnum Enum with underlying type: string
type UpdateExportDetailsExportModeEnum string

// Set of constants representing the allowable values for UpdateExportDetailsExportModeEnum
const (
	UpdateExportDetailsExportModeEnabled  UpdateExportDetailsExportModeEnum = "ENABLED"
	UpdateExportDetailsExportModeDisabled UpdateExportDetailsExportModeEnum = "DISABLED"
)

var mappingUpdateExportDetailsExportModeEnum = map[string]UpdateExportDetailsExportModeEnum{
	"ENABLED":  UpdateExportDetailsExportModeEnabled,
	"DISABLED": UpdateExportDetailsExportModeDisabled,
}

var mappingUpdateExportDetailsExportModeEnumLowerCase = map[string]UpdateExportDetailsExportModeEnum{
	"enabled":  UpdateExportDetailsExportModeEnabled,
	"disabled": UpdateExportDetailsExportModeDisabled,
}

// GetUpdateExportDetailsExportModeEnumValues Enumerates the set of values for UpdateExportDetailsExportModeEnum
func GetUpdateExportDetailsExportModeEnumValues() []UpdateExportDetailsExportModeEnum {
	values := make([]UpdateExportDetailsExportModeEnum, 0)
	for _, v := range mappingUpdateExportDetailsExportModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExportDetailsExportModeEnumStringValues Enumerates the set of values in String for UpdateExportDetailsExportModeEnum
func GetUpdateExportDetailsExportModeEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateExportDetailsExportModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExportDetailsExportModeEnum(val string) (UpdateExportDetailsExportModeEnum, bool) {
	enum, ok := mappingUpdateExportDetailsExportModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
