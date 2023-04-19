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

// ClientShareOptions SMB share options applied to a specified set of
// clients. Only governs access through the associated
// share. Access to the same file system through a different
// share will be governed
// by that share's share options.
type ClientShareOptions struct {

	// Clients these options should apply to. Must be a either
	// single IPv4 address or single IPv4 CIDR block.
	// **Note:** Access will also be limited by any applicable VCN
	// security rules and the ability to route IP packets to the
	// mount target. Mount targets do not have Internet-routable IP addresses.
	Source *string `mandatory:"true" json:"source"`

	// Type of access to grant clients using the file system
	// through this share. If unspecified defaults to `READ_ONLY`.
	Access ClientShareOptionsAccessEnum `mandatory:"false" json:"access,omitempty"`

	// UID value to remap to when squashing a client UID (see
	// identitySquash for more details.) If unspecified, defaults
	// to `0`.
	AnonymousUid *int64 `mandatory:"false" json:"anonymousUid"`

	// GID value to remap to when squashing a client GID (see
	// identitySquash for more details.) If unspecified defaults
	// to `0`.
	AnonymousGid *int64 `mandatory:"false" json:"anonymousGid"`
}

func (m ClientShareOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClientShareOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClientShareOptionsAccessEnum(string(m.Access)); !ok && m.Access != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Access: %s. Supported values are: %s.", m.Access, strings.Join(GetClientShareOptionsAccessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClientShareOptionsAccessEnum Enum with underlying type: string
type ClientShareOptionsAccessEnum string

// Set of constants representing the allowable values for ClientShareOptionsAccessEnum
const (
	ClientShareOptionsAccessWrite ClientShareOptionsAccessEnum = "READ_WRITE"
	ClientShareOptionsAccessOnly  ClientShareOptionsAccessEnum = "READ_ONLY"
)

var mappingClientShareOptionsAccessEnum = map[string]ClientShareOptionsAccessEnum{
	"READ_WRITE": ClientShareOptionsAccessWrite,
	"READ_ONLY":  ClientShareOptionsAccessOnly,
}

var mappingClientShareOptionsAccessEnumLowerCase = map[string]ClientShareOptionsAccessEnum{
	"read_write": ClientShareOptionsAccessWrite,
	"read_only":  ClientShareOptionsAccessOnly,
}

// GetClientShareOptionsAccessEnumValues Enumerates the set of values for ClientShareOptionsAccessEnum
func GetClientShareOptionsAccessEnumValues() []ClientShareOptionsAccessEnum {
	values := make([]ClientShareOptionsAccessEnum, 0)
	for _, v := range mappingClientShareOptionsAccessEnum {
		values = append(values, v)
	}
	return values
}

// GetClientShareOptionsAccessEnumStringValues Enumerates the set of values in String for ClientShareOptionsAccessEnum
func GetClientShareOptionsAccessEnumStringValues() []string {
	return []string{
		"READ_WRITE",
		"READ_ONLY",
	}
}

// GetMappingClientShareOptionsAccessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClientShareOptionsAccessEnum(val string) (ClientShareOptionsAccessEnum, bool) {
	enum, ok := mappingClientShareOptionsAccessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
