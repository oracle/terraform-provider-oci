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

// RootSquashConfiguration An administrative feature that allows you to restrict root level access from clients that try to access your Lustre file system as root.
type RootSquashConfiguration struct {

	// Used when clients accessing the Lustre file system have their UID and GID remapped to
	// `squashUid` and `squashGid`. If `ROOT`, only the root user and group (UID/GID 0) are remapped;
	// if `NONE`, no remapping is done. If unspecified, defaults to `NONE`.
	IdentitySquash RootSquashConfigurationIdentitySquashEnum `mandatory:"false" json:"identitySquash,omitempty"`

	// The UID value to remap to when squashing a client UID. See
	// `identitySquash` for more details. If unspecified, defaults
	// to `65534`.
	SquashUid *int64 `mandatory:"false" json:"squashUid"`

	// The GID value to remap to when squashing a client GID. See
	// `identitySquash` for more details. If unspecified, defaults
	// to `65534`.
	SquashGid *int64 `mandatory:"false" json:"squashGid"`

	// A list of NIDs allowed with this lustre file system not to be squashed.
	// A maximum of 10 is allowed.
	ClientExceptions []string `mandatory:"false" json:"clientExceptions"`
}

func (m RootSquashConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RootSquashConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRootSquashConfigurationIdentitySquashEnum(string(m.IdentitySquash)); !ok && m.IdentitySquash != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdentitySquash: %s. Supported values are: %s.", m.IdentitySquash, strings.Join(GetRootSquashConfigurationIdentitySquashEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RootSquashConfigurationIdentitySquashEnum Enum with underlying type: string
type RootSquashConfigurationIdentitySquashEnum string

// Set of constants representing the allowable values for RootSquashConfigurationIdentitySquashEnum
const (
	RootSquashConfigurationIdentitySquashNone RootSquashConfigurationIdentitySquashEnum = "NONE"
	RootSquashConfigurationIdentitySquashRoot RootSquashConfigurationIdentitySquashEnum = "ROOT"
)

var mappingRootSquashConfigurationIdentitySquashEnum = map[string]RootSquashConfigurationIdentitySquashEnum{
	"NONE": RootSquashConfigurationIdentitySquashNone,
	"ROOT": RootSquashConfigurationIdentitySquashRoot,
}

var mappingRootSquashConfigurationIdentitySquashEnumLowerCase = map[string]RootSquashConfigurationIdentitySquashEnum{
	"none": RootSquashConfigurationIdentitySquashNone,
	"root": RootSquashConfigurationIdentitySquashRoot,
}

// GetRootSquashConfigurationIdentitySquashEnumValues Enumerates the set of values for RootSquashConfigurationIdentitySquashEnum
func GetRootSquashConfigurationIdentitySquashEnumValues() []RootSquashConfigurationIdentitySquashEnum {
	values := make([]RootSquashConfigurationIdentitySquashEnum, 0)
	for _, v := range mappingRootSquashConfigurationIdentitySquashEnum {
		values = append(values, v)
	}
	return values
}

// GetRootSquashConfigurationIdentitySquashEnumStringValues Enumerates the set of values in String for RootSquashConfigurationIdentitySquashEnum
func GetRootSquashConfigurationIdentitySquashEnumStringValues() []string {
	return []string{
		"NONE",
		"ROOT",
	}
}

// GetMappingRootSquashConfigurationIdentitySquashEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRootSquashConfigurationIdentitySquashEnum(val string) (RootSquashConfigurationIdentitySquashEnum, bool) {
	enum, ok := mappingRootSquashConfigurationIdentitySquashEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
