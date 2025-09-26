// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateIotDomain                OperationTypeEnum = "CREATE_IOT_DOMAIN"
	OperationTypeUpdateIotDomain                OperationTypeEnum = "UPDATE_IOT_DOMAIN"
	OperationTypeDeleteIotDomain                OperationTypeEnum = "DELETE_IOT_DOMAIN"
	OperationTypeMoveIotDomain                  OperationTypeEnum = "MOVE_IOT_DOMAIN"
	OperationTypeMoveIotDomainGroup             OperationTypeEnum = "MOVE_IOT_DOMAIN_GROUP"
	OperationTypeCreateIotDomainGroup           OperationTypeEnum = "CREATE_IOT_DOMAIN_GROUP"
	OperationTypeUpdateIotDomainGroup           OperationTypeEnum = "UPDATE_IOT_DOMAIN_GROUP"
	OperationTypeDeleteIotDomainGroup           OperationTypeEnum = "DELETE_IOT_DOMAIN_GROUP"
	OperationTypeConfigureDomainDataAccess      OperationTypeEnum = "CONFIGURE_DOMAIN_DATA_ACCESS"
	OperationTypeConfigureDomainGroupDataAccess OperationTypeEnum = "CONFIGURE_DOMAIN_GROUP_DATA_ACCESS"
	OperationTypeChangeIotDomainDataRetention   OperationTypeEnum = "CHANGE_IOT_DOMAIN_DATA_RETENTION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_IOT_DOMAIN":                  OperationTypeCreateIotDomain,
	"UPDATE_IOT_DOMAIN":                  OperationTypeUpdateIotDomain,
	"DELETE_IOT_DOMAIN":                  OperationTypeDeleteIotDomain,
	"MOVE_IOT_DOMAIN":                    OperationTypeMoveIotDomain,
	"MOVE_IOT_DOMAIN_GROUP":              OperationTypeMoveIotDomainGroup,
	"CREATE_IOT_DOMAIN_GROUP":            OperationTypeCreateIotDomainGroup,
	"UPDATE_IOT_DOMAIN_GROUP":            OperationTypeUpdateIotDomainGroup,
	"DELETE_IOT_DOMAIN_GROUP":            OperationTypeDeleteIotDomainGroup,
	"CONFIGURE_DOMAIN_DATA_ACCESS":       OperationTypeConfigureDomainDataAccess,
	"CONFIGURE_DOMAIN_GROUP_DATA_ACCESS": OperationTypeConfigureDomainGroupDataAccess,
	"CHANGE_IOT_DOMAIN_DATA_RETENTION":   OperationTypeChangeIotDomainDataRetention,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_iot_domain":                  OperationTypeCreateIotDomain,
	"update_iot_domain":                  OperationTypeUpdateIotDomain,
	"delete_iot_domain":                  OperationTypeDeleteIotDomain,
	"move_iot_domain":                    OperationTypeMoveIotDomain,
	"move_iot_domain_group":              OperationTypeMoveIotDomainGroup,
	"create_iot_domain_group":            OperationTypeCreateIotDomainGroup,
	"update_iot_domain_group":            OperationTypeUpdateIotDomainGroup,
	"delete_iot_domain_group":            OperationTypeDeleteIotDomainGroup,
	"configure_domain_data_access":       OperationTypeConfigureDomainDataAccess,
	"configure_domain_group_data_access": OperationTypeConfigureDomainGroupDataAccess,
	"change_iot_domain_data_retention":   OperationTypeChangeIotDomainDataRetention,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_IOT_DOMAIN",
		"UPDATE_IOT_DOMAIN",
		"DELETE_IOT_DOMAIN",
		"MOVE_IOT_DOMAIN",
		"MOVE_IOT_DOMAIN_GROUP",
		"CREATE_IOT_DOMAIN_GROUP",
		"UPDATE_IOT_DOMAIN_GROUP",
		"DELETE_IOT_DOMAIN_GROUP",
		"CONFIGURE_DOMAIN_DATA_ACCESS",
		"CONFIGURE_DOMAIN_GROUP_DATA_ACCESS",
		"CHANGE_IOT_DOMAIN_DATA_RETENTION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
