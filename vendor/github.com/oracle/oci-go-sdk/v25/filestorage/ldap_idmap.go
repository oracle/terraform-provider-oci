// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// LdapIdmap Mount Target details about the LDAP ID Map configuration.
type LdapIdmap struct {

	// Schema type of LDAP account.
	SchemaType LdapIdmapSchemaTypeEnum `mandatory:"false" json:"schemaType,omitempty"`

	// Integer for how often the mount target should recheck LDAP for updates.
	CacheRefreshIntervalSeconds *int `mandatory:"false" json:"cacheRefreshIntervalSeconds"`

	// Integer for how long cached entries may be used.
	CacheLifetimeSeconds *int `mandatory:"false" json:"cacheLifetimeSeconds"`

	// Integer for how long to cache if idmap information is missing.
	NegativeCacheLifetimeSeconds *int `mandatory:"false" json:"negativeCacheLifetimeSeconds"`

	BaseDistinguishedUserName *string `mandatory:"false" json:"baseDistinguishedUserName"`

	BaseDistinguishedGroupName *string `mandatory:"false" json:"baseDistinguishedGroupName"`

	// OCID of the first LDAP Account
	OutboundConnector1Id *string `mandatory:"false" json:"outboundConnector1Id"`

	// OCID of the second LDAP Account
	OutboundConnector2Id *string `mandatory:"false" json:"outboundConnector2Id"`
}

func (m LdapIdmap) String() string {
	return common.PointerString(m)
}

// LdapIdmapSchemaTypeEnum Enum with underlying type: string
type LdapIdmapSchemaTypeEnum string

// Set of constants representing the allowable values for LdapIdmapSchemaTypeEnum
const (
	LdapIdmapSchemaTypeRfc2307 LdapIdmapSchemaTypeEnum = "RFC2307"
)

var mappingLdapIdmapSchemaType = map[string]LdapIdmapSchemaTypeEnum{
	"RFC2307": LdapIdmapSchemaTypeRfc2307,
}

// GetLdapIdmapSchemaTypeEnumValues Enumerates the set of values for LdapIdmapSchemaTypeEnum
func GetLdapIdmapSchemaTypeEnumValues() []LdapIdmapSchemaTypeEnum {
	values := make([]LdapIdmapSchemaTypeEnum, 0)
	for _, v := range mappingLdapIdmapSchemaType {
		values = append(values, v)
	}
	return values
}
