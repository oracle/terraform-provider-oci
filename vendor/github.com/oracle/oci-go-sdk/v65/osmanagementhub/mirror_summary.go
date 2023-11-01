// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MirrorSummary Summary of a Mirror
type MirrorSummary struct {

	// OCID of a software source
	Id *string `mandatory:"true" json:"id"`

	// Current state of the mirror
	State MirrorStateEnum `mandatory:"true" json:"state"`

	// A decimal number representing the completness percentage
	Percentage *int `mandatory:"true" json:"percentage"`

	// Timestamp of the last time the mirror was sync
	TimeLastSynced *common.SDKTime `mandatory:"true" json:"timeLastSynced"`

	// The current log from the management station plugin.
	Log *string `mandatory:"true" json:"log"`

	// Display name of the mirror
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of the mirror
	Type MirrorTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The OS family the Software Source belongs to
	OsFamily OsFamilyEnum `mandatory:"false" json:"osFamily,omitempty"`

	// The architecture type supported by the Software Source
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
