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

// UpdateLdapIdmapDetails Mount target details about the LDAP ID mapping configuration.
type UpdateLdapIdmapDetails struct {

	// Schema type of the LDAP account.
	SchemaType UpdateLdapIdmapDetailsSchemaTypeEnum `mandatory:"false" json:"schemaType,omitempty"`

	// The amount of time that the mount target should allow an entry to persist in its cache before attempting to refresh the entry.
	CacheRefreshIntervalSeconds *int `mandatory:"false" json:"cacheRefreshIntervalSeconds"`

	// The maximum amount of time the mount target is allowed to use a cached entry.
	CacheLifetimeSeconds *int `mandatory:"false" json:"cacheLifetimeSeconds"`

	// The amount of time that a mount target will maintain information that a user is not found in the ID mapping configuration.
	NegativeCacheLifetimeSeconds *int `mandatory:"false" json:"negativeCacheLifetimeSeconds"`

	// All LDAP searches are recursive starting at this user.
	// Example: `CN=User,DC=domain,DC=com`
	UserSearchBase *string `mandatory:"false" json:"userSearchBase"`

	// All LDAP searches are recursive starting at this group.
	// Example: `CN=Group,DC=domain,DC=com`
	GroupSearchBase *string `mandatory:"false" json:"groupSearchBase"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the first connector to use to communicate with the LDAP server.
	OutboundConnector1Id *string `mandatory:"false" json:"outboundConnector1Id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the second connector to use to communicate with the LDAP server.
	OutboundConnector2Id *string `mandatory:"false" json:"outboundConnector2Id"`
}

func (m UpdateLdapIdmapDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLdapIdmapDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateLdapIdmapDetailsSchemaTypeEnum(string(m.SchemaType)); !ok && m.SchemaType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchemaType: %s. Supported values are: %s.", m.SchemaType, strings.Join(GetUpdateLdapIdmapDetailsSchemaTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateLdapIdmapDetailsSchemaTypeEnum Enum with underlying type: string
type UpdateLdapIdmapDetailsSchemaTypeEnum string

// Set of constants representing the allowable values for UpdateLdapIdmapDetailsSchemaTypeEnum
const (
	UpdateLdapIdmapDetailsSchemaTypeRfc2307 UpdateLdapIdmapDetailsSchemaTypeEnum = "RFC2307"
)

var mappingUpdateLdapIdmapDetailsSchemaTypeEnum = map[string]UpdateLdapIdmapDetailsSchemaTypeEnum{
	"RFC2307": UpdateLdapIdmapDetailsSchemaTypeRfc2307,
}

var mappingUpdateLdapIdmapDetailsSchemaTypeEnumLowerCase = map[string]UpdateLdapIdmapDetailsSchemaTypeEnum{
	"rfc2307": UpdateLdapIdmapDetailsSchemaTypeRfc2307,
}

// GetUpdateLdapIdmapDetailsSchemaTypeEnumValues Enumerates the set of values for UpdateLdapIdmapDetailsSchemaTypeEnum
func GetUpdateLdapIdmapDetailsSchemaTypeEnumValues() []UpdateLdapIdmapDetailsSchemaTypeEnum {
	values := make([]UpdateLdapIdmapDetailsSchemaTypeEnum, 0)
	for _, v := range mappingUpdateLdapIdmapDetailsSchemaTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateLdapIdmapDetailsSchemaTypeEnumStringValues Enumerates the set of values in String for UpdateLdapIdmapDetailsSchemaTypeEnum
func GetUpdateLdapIdmapDetailsSchemaTypeEnumStringValues() []string {
	return []string{
		"RFC2307",
	}
}

// GetMappingUpdateLdapIdmapDetailsSchemaTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateLdapIdmapDetailsSchemaTypeEnum(val string) (UpdateLdapIdmapDetailsSchemaTypeEnum, bool) {
	enum, ok := mappingUpdateLdapIdmapDetailsSchemaTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
