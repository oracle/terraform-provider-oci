// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDkim             OperationTypeEnum = "CREATE_DKIM"
	OperationTypeDeleteDkim             OperationTypeEnum = "DELETE_DKIM"
	OperationTypeMoveDkim               OperationTypeEnum = "MOVE_DKIM"
	OperationTypeUpdateDkim             OperationTypeEnum = "UPDATE_DKIM"
	OperationTypeCreateEmailDomain      OperationTypeEnum = "CREATE_EMAIL_DOMAIN"
	OperationTypeDeleteEmailDomain      OperationTypeEnum = "DELETE_EMAIL_DOMAIN"
	OperationTypeMoveEmailDomain        OperationTypeEnum = "MOVE_EMAIL_DOMAIN"
	OperationTypeUpdateEmailDomain      OperationTypeEnum = "UPDATE_EMAIL_DOMAIN"
	OperationTypeCreatePrivateEndpoint  OperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	OperationTypeDeletePrivateEndpoint  OperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	OperationTypeMovePrivateEndpoint    OperationTypeEnum = "MOVE_PRIVATE_ENDPOINT"
	OperationTypeUpdatePrivateEndpoint  OperationTypeEnum = "UPDATE_PRIVATE_ENDPOINT"
	OperationTypeCreateEmailTrackConfig OperationTypeEnum = "CREATE_EMAIL_TRACK_CONFIG"
	OperationTypeUpdateEmailTrackConfig OperationTypeEnum = "UPDATE_EMAIL_TRACK_CONFIG"
	OperationTypeDeleteEmailTrackConfig OperationTypeEnum = "DELETE_EMAIL_TRACK_CONFIG"
	OperationTypeMoveEmailTrackConfig   OperationTypeEnum = "MOVE_EMAIL_TRACK_CONFIG"
	OperationTypeCreateReturnPath       OperationTypeEnum = "CREATE_RETURN_PATH"
	OperationTypeDeleteReturnPath       OperationTypeEnum = "DELETE_RETURN_PATH"
	OperationTypeUpdateReturnPath       OperationTypeEnum = "UPDATE_RETURN_PATH"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DKIM":               OperationTypeCreateDkim,
	"DELETE_DKIM":               OperationTypeDeleteDkim,
	"MOVE_DKIM":                 OperationTypeMoveDkim,
	"UPDATE_DKIM":               OperationTypeUpdateDkim,
	"CREATE_EMAIL_DOMAIN":       OperationTypeCreateEmailDomain,
	"DELETE_EMAIL_DOMAIN":       OperationTypeDeleteEmailDomain,
	"MOVE_EMAIL_DOMAIN":         OperationTypeMoveEmailDomain,
	"UPDATE_EMAIL_DOMAIN":       OperationTypeUpdateEmailDomain,
	"CREATE_PRIVATE_ENDPOINT":   OperationTypeCreatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":   OperationTypeDeletePrivateEndpoint,
	"MOVE_PRIVATE_ENDPOINT":     OperationTypeMovePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT":   OperationTypeUpdatePrivateEndpoint,
	"CREATE_EMAIL_TRACK_CONFIG": OperationTypeCreateEmailTrackConfig,
	"UPDATE_EMAIL_TRACK_CONFIG": OperationTypeUpdateEmailTrackConfig,
	"DELETE_EMAIL_TRACK_CONFIG": OperationTypeDeleteEmailTrackConfig,
	"MOVE_EMAIL_TRACK_CONFIG":   OperationTypeMoveEmailTrackConfig,
	"CREATE_RETURN_PATH":        OperationTypeCreateReturnPath,
	"DELETE_RETURN_PATH":        OperationTypeDeleteReturnPath,
	"UPDATE_RETURN_PATH":        OperationTypeUpdateReturnPath,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_dkim":               OperationTypeCreateDkim,
	"delete_dkim":               OperationTypeDeleteDkim,
	"move_dkim":                 OperationTypeMoveDkim,
	"update_dkim":               OperationTypeUpdateDkim,
	"create_email_domain":       OperationTypeCreateEmailDomain,
	"delete_email_domain":       OperationTypeDeleteEmailDomain,
	"move_email_domain":         OperationTypeMoveEmailDomain,
	"update_email_domain":       OperationTypeUpdateEmailDomain,
	"create_private_endpoint":   OperationTypeCreatePrivateEndpoint,
	"delete_private_endpoint":   OperationTypeDeletePrivateEndpoint,
	"move_private_endpoint":     OperationTypeMovePrivateEndpoint,
	"update_private_endpoint":   OperationTypeUpdatePrivateEndpoint,
	"create_email_track_config": OperationTypeCreateEmailTrackConfig,
	"update_email_track_config": OperationTypeUpdateEmailTrackConfig,
	"delete_email_track_config": OperationTypeDeleteEmailTrackConfig,
	"move_email_track_config":   OperationTypeMoveEmailTrackConfig,
	"create_return_path":        OperationTypeCreateReturnPath,
	"delete_return_path":        OperationTypeDeleteReturnPath,
	"update_return_path":        OperationTypeUpdateReturnPath,
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
		"CREATE_DKIM",
		"DELETE_DKIM",
		"MOVE_DKIM",
		"UPDATE_DKIM",
		"CREATE_EMAIL_DOMAIN",
		"DELETE_EMAIL_DOMAIN",
		"MOVE_EMAIL_DOMAIN",
		"UPDATE_EMAIL_DOMAIN",
		"CREATE_PRIVATE_ENDPOINT",
		"DELETE_PRIVATE_ENDPOINT",
		"MOVE_PRIVATE_ENDPOINT",
		"UPDATE_PRIVATE_ENDPOINT",
		"CREATE_EMAIL_TRACK_CONFIG",
		"UPDATE_EMAIL_TRACK_CONFIG",
		"DELETE_EMAIL_TRACK_CONFIG",
		"MOVE_EMAIL_TRACK_CONFIG",
		"CREATE_RETURN_PATH",
		"DELETE_RETURN_PATH",
		"UPDATE_RETURN_PATH",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
