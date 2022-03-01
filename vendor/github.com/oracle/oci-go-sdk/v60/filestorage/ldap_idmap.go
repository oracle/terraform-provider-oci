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
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// LdapIdmap Mount Target details about the LDAP ID Map configuration.
type LdapIdmap struct {

	// All LDAP searches are recursive starting at this Base Distinguished user name.
	BaseDistinguishedUserName *string `mandatory:"true" json:"baseDistinguishedUserName"`

	// All LDAP searches are recursive starting at this Base Distinguished group name.
	BaseDistinguishedGroupName *string `mandatory:"true" json:"baseDistinguishedGroupName"`

	// Schema type of LDAP account.
	SchemaType LdapIdmapSchemaTypeEnum `mandatory:"false" json:"schemaType,omitempty"`

	// Integer for how often the mount target should recheck LDAP for updates.
	CacheRefreshIntervalSeconds *int `mandatory:"false" json:"cacheRefreshIntervalSeconds"`

	// Integer for how long cached entries may be used.
	CacheLifetimeSeconds *int `mandatory:"false" json:"cacheLifetimeSeconds"`

	// Integer for how long to cache if idmap information is missing.
	NegativeCacheLifetimeSeconds *int `mandatory:"false" json:"negativeCacheLifetimeSeconds"`

	// OCID of the first LDAP Account
	OutboundConnector1Id *string `mandatory:"false" json:"outboundConnector1Id"`

	// OCID of the second LDAP Account
	OutboundConnector2Id *string `mandatory:"false" json:"outboundConnector2Id"`
}

func (m LdapIdmap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LdapIdmap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLdapIdmapSchemaTypeEnum(string(m.SchemaType)); !ok && m.SchemaType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchemaType: %s. Supported values are: %s.", m.SchemaType, strings.Join(GetLdapIdmapSchemaTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LdapIdmapSchemaTypeEnum Enum with underlying type: string
type LdapIdmapSchemaTypeEnum string

// Set of constants representing the allowable values for LdapIdmapSchemaTypeEnum
const (
	LdapIdmapSchemaTypeRfc2307 LdapIdmapSchemaTypeEnum = "RFC2307"
)

var mappingLdapIdmapSchemaTypeEnum = map[string]LdapIdmapSchemaTypeEnum{
	"RFC2307": LdapIdmapSchemaTypeRfc2307,
}

var mappingLdapIdmapSchemaTypeEnumLowerCase = map[string]LdapIdmapSchemaTypeEnum{
	"rfc2307": LdapIdmapSchemaTypeRfc2307,
}

// GetLdapIdmapSchemaTypeEnumValues Enumerates the set of values for LdapIdmapSchemaTypeEnum
func GetLdapIdmapSchemaTypeEnumValues() []LdapIdmapSchemaTypeEnum {
	values := make([]LdapIdmapSchemaTypeEnum, 0)
	for _, v := range mappingLdapIdmapSchemaTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLdapIdmapSchemaTypeEnumStringValues Enumerates the set of values in String for LdapIdmapSchemaTypeEnum
func GetLdapIdmapSchemaTypeEnumStringValues() []string {
	return []string{
		"RFC2307",
	}
}

// GetMappingLdapIdmapSchemaTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLdapIdmapSchemaTypeEnum(val string) (LdapIdmapSchemaTypeEnum, bool) {
	enum, ok := mappingLdapIdmapSchemaTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
