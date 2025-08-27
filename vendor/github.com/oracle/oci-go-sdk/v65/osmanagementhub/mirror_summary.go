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

// MirrorSummary Provides summary information for a software source mirror.
type MirrorSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	Id *string `mandatory:"true" json:"id"`

	// Current state of the software source mirror.
	State MirrorStateEnum `mandatory:"true" json:"state"`

	// A decimal number representing the percentage of the software source that has been synced.
	Percentage *int `mandatory:"true" json:"percentage"`

	// Time that the software source was last synced (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeLastSynced *common.SDKTime `mandatory:"true" json:"timeLastSynced"`

	// The current log from the management station plugin.
	Log *string `mandatory:"true" json:"log"`

	// The number of packages within the mirrored software source.
	PackageCount *int `mandatory:"true" json:"packageCount"`

	// The size the mirrored software source in bytes.
	Size *int64 `mandatory:"true" json:"size"`

	// Display name of the mirror.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of software source.
	Type MirrorTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The OS family of the software source.
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	// The architecture type supported by the software source.
	ArchType ArchTypeEnum `mandatory:"false" json:"archType,omitempty"`
}

func (m MirrorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MirrorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMirrorStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetMirrorStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMirrorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMirrorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamilyEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamilyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArchTypeEnum(string(m.ArchType)); !ok && m.ArchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", m.ArchType, strings.Join(GetArchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
