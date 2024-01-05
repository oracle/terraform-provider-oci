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

// ClientOptions NFS export options applied to a specified set of
// clients. Only governs access through the associated
// export. Access to the same file system through a different
// export (on the same or different mount target) will be governed
// by that export's export options.
type ClientOptions struct {

	// Clients these options should apply to. Must be a either
	// single IPv4 address or single IPv4 CIDR block.
	// **Note:** Access will also be limited by any applicable VCN
	// security rules and the ability to route IP packets to the
	// mount target. Mount targets do not have Internet-routable IP addresses.
	Source *string `mandatory:"true" json:"source"`

	// If `true`, clients accessing the file system through this
	// export must connect from a privileged source port. If
	// unspecified, defaults to `true`.
	RequirePrivilegedSourcePort *bool `mandatory:"false" json:"requirePrivilegedSourcePort"`

	// Type of access to grant clients using the file system
	// through this export. If unspecified defaults to `READ_WRITE`.
	Access ClientOptionsAccessEnum `mandatory:"false" json:"access,omitempty"`

	// Used when clients accessing the file system through this export
	// have their UID and GID remapped to 'anonymousUid' and
	// 'anonymousGid'. If `ALL`, all users and groups are remapped;
	// if `ROOT`, only the root user and group (UID/GID 0) are
	// remapped; if `NONE`, no remapping is done. If unspecified,
	// defaults to `ROOT`.
	IdentitySquash ClientOptionsIdentitySquashEnum `mandatory:"false" json:"identitySquash,omitempty"`

	// UID value to remap to when squashing a client UID (see
	// identitySquash for more details.) If unspecified, defaults
	// to `65534`.
	AnonymousUid *int64 `mandatory:"false" json:"anonymousUid"`

	// GID value to remap to when squashing a client GID (see
	// identitySquash for more details.) If unspecified defaults
	// to `65534`.
	AnonymousGid *int64 `mandatory:"false" json:"anonymousGid"`

	// Whether or not to enable anonymous access to the file system through this export in cases where a user isn't found in the LDAP server used for ID mapping.
	// If true, and the user is not found in the LDAP directory, the operation uses the Squash UID and Squash GID.
	IsAnonymousAccessAllowed *bool `mandatory:"false" json:"isAnonymousAccessAllowed"`

	// Array of allowed NFS authentication types.
	AllowedAuth []ClientOptionsAllowedAuthEnum `mandatory:"false" json:"allowedAuth,omitempty"`
}

func (m ClientOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClientOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClientOptionsAccessEnum(string(m.Access)); !ok && m.Access != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Access: %s. Supported values are: %s.", m.Access, strings.Join(GetClientOptionsAccessEnumStringValues(), ",")))
	}
	if _, ok := GetMappingClientOptionsIdentitySquashEnum(string(m.IdentitySquash)); !ok && m.IdentitySquash != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdentitySquash: %s. Supported values are: %s.", m.IdentitySquash, strings.Join(GetClientOptionsIdentitySquashEnumStringValues(), ",")))
	}
	for _, val := range m.AllowedAuth {
		if _, ok := GetMappingClientOptionsAllowedAuthEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AllowedAuth: %s. Supported values are: %s.", val, strings.Join(GetClientOptionsAllowedAuthEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClientOptionsAccessEnum Enum with underlying type: string
type ClientOptionsAccessEnum string

// Set of constants representing the allowable values for ClientOptionsAccessEnum
const (
	ClientOptionsAccessWrite ClientOptionsAccessEnum = "READ_WRITE"
	ClientOptionsAccessOnly  ClientOptionsAccessEnum = "READ_ONLY"
)

var mappingClientOptionsAccessEnum = map[string]ClientOptionsAccessEnum{
	"READ_WRITE": ClientOptionsAccessWrite,
	"READ_ONLY":  ClientOptionsAccessOnly,
}

var mappingClientOptionsAccessEnumLowerCase = map[string]ClientOptionsAccessEnum{
	"read_write": ClientOptionsAccessWrite,
	"read_only":  ClientOptionsAccessOnly,
}

// GetClientOptionsAccessEnumValues Enumerates the set of values for ClientOptionsAccessEnum
func GetClientOptionsAccessEnumValues() []ClientOptionsAccessEnum {
	values := make([]ClientOptionsAccessEnum, 0)
	for _, v := range mappingClientOptionsAccessEnum {
		values = append(values, v)
	}
	return values
}

// GetClientOptionsAccessEnumStringValues Enumerates the set of values in String for ClientOptionsAccessEnum
func GetClientOptionsAccessEnumStringValues() []string {
	return []string{
		"READ_WRITE",
		"READ_ONLY",
	}
}

// GetMappingClientOptionsAccessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClientOptionsAccessEnum(val string) (ClientOptionsAccessEnum, bool) {
	enum, ok := mappingClientOptionsAccessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ClientOptionsIdentitySquashEnum Enum with underlying type: string
type ClientOptionsIdentitySquashEnum string

// Set of constants representing the allowable values for ClientOptionsIdentitySquashEnum
const (
	ClientOptionsIdentitySquashNone ClientOptionsIdentitySquashEnum = "NONE"
	ClientOptionsIdentitySquashRoot ClientOptionsIdentitySquashEnum = "ROOT"
	ClientOptionsIdentitySquashAll  ClientOptionsIdentitySquashEnum = "ALL"
)

var mappingClientOptionsIdentitySquashEnum = map[string]ClientOptionsIdentitySquashEnum{
	"NONE": ClientOptionsIdentitySquashNone,
	"ROOT": ClientOptionsIdentitySquashRoot,
	"ALL":  ClientOptionsIdentitySquashAll,
}

var mappingClientOptionsIdentitySquashEnumLowerCase = map[string]ClientOptionsIdentitySquashEnum{
	"none": ClientOptionsIdentitySquashNone,
	"root": ClientOptionsIdentitySquashRoot,
	"all":  ClientOptionsIdentitySquashAll,
}

// GetClientOptionsIdentitySquashEnumValues Enumerates the set of values for ClientOptionsIdentitySquashEnum
func GetClientOptionsIdentitySquashEnumValues() []ClientOptionsIdentitySquashEnum {
	values := make([]ClientOptionsIdentitySquashEnum, 0)
	for _, v := range mappingClientOptionsIdentitySquashEnum {
		values = append(values, v)
	}
	return values
}

// GetClientOptionsIdentitySquashEnumStringValues Enumerates the set of values in String for ClientOptionsIdentitySquashEnum
func GetClientOptionsIdentitySquashEnumStringValues() []string {
	return []string{
		"NONE",
		"ROOT",
		"ALL",
	}
}

// GetMappingClientOptionsIdentitySquashEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClientOptionsIdentitySquashEnum(val string) (ClientOptionsIdentitySquashEnum, bool) {
	enum, ok := mappingClientOptionsIdentitySquashEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ClientOptionsAllowedAuthEnum Enum with underlying type: string
type ClientOptionsAllowedAuthEnum string

// Set of constants representing the allowable values for ClientOptionsAllowedAuthEnum
const (
	ClientOptionsAllowedAuthSys   ClientOptionsAllowedAuthEnum = "SYS"
	ClientOptionsAllowedAuthKrb5  ClientOptionsAllowedAuthEnum = "KRB5"
	ClientOptionsAllowedAuthKrb5i ClientOptionsAllowedAuthEnum = "KRB5I"
	ClientOptionsAllowedAuthKrb5p ClientOptionsAllowedAuthEnum = "KRB5P"
)

var mappingClientOptionsAllowedAuthEnum = map[string]ClientOptionsAllowedAuthEnum{
	"SYS":   ClientOptionsAllowedAuthSys,
	"KRB5":  ClientOptionsAllowedAuthKrb5,
	"KRB5I": ClientOptionsAllowedAuthKrb5i,
	"KRB5P": ClientOptionsAllowedAuthKrb5p,
}

var mappingClientOptionsAllowedAuthEnumLowerCase = map[string]ClientOptionsAllowedAuthEnum{
	"sys":   ClientOptionsAllowedAuthSys,
	"krb5":  ClientOptionsAllowedAuthKrb5,
	"krb5i": ClientOptionsAllowedAuthKrb5i,
	"krb5p": ClientOptionsAllowedAuthKrb5p,
}

// GetClientOptionsAllowedAuthEnumValues Enumerates the set of values for ClientOptionsAllowedAuthEnum
func GetClientOptionsAllowedAuthEnumValues() []ClientOptionsAllowedAuthEnum {
	values := make([]ClientOptionsAllowedAuthEnum, 0)
	for _, v := range mappingClientOptionsAllowedAuthEnum {
		values = append(values, v)
	}
	return values
}

// GetClientOptionsAllowedAuthEnumStringValues Enumerates the set of values in String for ClientOptionsAllowedAuthEnum
func GetClientOptionsAllowedAuthEnumStringValues() []string {
	return []string{
		"SYS",
		"KRB5",
		"KRB5I",
		"KRB5P",
	}
}

// GetMappingClientOptionsAllowedAuthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClientOptionsAllowedAuthEnum(val string) (ClientOptionsAllowedAuthEnum, bool) {
	enum, ok := mappingClientOptionsAllowedAuthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
