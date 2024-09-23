// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateFleet             OperationTypeEnum = "CREATE_FLEET"
	OperationTypeUpdateFleet             OperationTypeEnum = "UPDATE_FLEET"
	OperationTypeDeleteFleet             OperationTypeEnum = "DELETE_FLEET"
	OperationTypeConfirmTarget           OperationTypeEnum = "CONFIRM_TARGET"
	OperationTypeGenerateCompliance      OperationTypeEnum = "GENERATE_COMPLIANCE"
	OperationTypeRequestTargetDiscovery  OperationTypeEnum = "REQUEST_TARGET_DISCOVERY"
	OperationTypeValidateResource        OperationTypeEnum = "VALIDATE_RESOURCE"
	OperationTypeCreateCredential        OperationTypeEnum = "CREATE_CREDENTIAL"
	OperationTypeUpdateCredential        OperationTypeEnum = "UPDATE_CREDENTIAL"
	OperationTypeDeleteCredential        OperationTypeEnum = "DELETE_CREDENTIAL"
	OperationTypeCreateSchedule          OperationTypeEnum = "CREATE_SCHEDULE"
	OperationTypeUpdateSchedule          OperationTypeEnum = "UPDATE_SCHEDULE"
	OperationTypeUpdateMaintenanceWindow OperationTypeEnum = "UPDATE_MAINTENANCE_WINDOW"
	OperationTypeDeleteMaintenanceWindow OperationTypeEnum = "DELETE_MAINTENANCE_WINDOW"
	OperationTypeCreateFleetResource     OperationTypeEnum = "CREATE_FLEET_RESOURCE"
	OperationTypeUpdateFleetResource     OperationTypeEnum = "UPDATE_FLEET_RESOURCE"
	OperationTypeDeleteFleetResource     OperationTypeEnum = "DELETE_FLEET_RESOURCE"
	OperationTypeCreateFamsOnboarding    OperationTypeEnum = "CREATE_FAMS_ONBOARDING"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_FLEET":              OperationTypeCreateFleet,
	"UPDATE_FLEET":              OperationTypeUpdateFleet,
	"DELETE_FLEET":              OperationTypeDeleteFleet,
	"CONFIRM_TARGET":            OperationTypeConfirmTarget,
	"GENERATE_COMPLIANCE":       OperationTypeGenerateCompliance,
	"REQUEST_TARGET_DISCOVERY":  OperationTypeRequestTargetDiscovery,
	"VALIDATE_RESOURCE":         OperationTypeValidateResource,
	"CREATE_CREDENTIAL":         OperationTypeCreateCredential,
	"UPDATE_CREDENTIAL":         OperationTypeUpdateCredential,
	"DELETE_CREDENTIAL":         OperationTypeDeleteCredential,
	"CREATE_SCHEDULE":           OperationTypeCreateSchedule,
	"UPDATE_SCHEDULE":           OperationTypeUpdateSchedule,
	"UPDATE_MAINTENANCE_WINDOW": OperationTypeUpdateMaintenanceWindow,
	"DELETE_MAINTENANCE_WINDOW": OperationTypeDeleteMaintenanceWindow,
	"CREATE_FLEET_RESOURCE":     OperationTypeCreateFleetResource,
	"UPDATE_FLEET_RESOURCE":     OperationTypeUpdateFleetResource,
	"DELETE_FLEET_RESOURCE":     OperationTypeDeleteFleetResource,
	"CREATE_FAMS_ONBOARDING":    OperationTypeCreateFamsOnboarding,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_fleet":              OperationTypeCreateFleet,
	"update_fleet":              OperationTypeUpdateFleet,
	"delete_fleet":              OperationTypeDeleteFleet,
	"confirm_target":            OperationTypeConfirmTarget,
	"generate_compliance":       OperationTypeGenerateCompliance,
	"request_target_discovery":  OperationTypeRequestTargetDiscovery,
	"validate_resource":         OperationTypeValidateResource,
	"create_credential":         OperationTypeCreateCredential,
	"update_credential":         OperationTypeUpdateCredential,
	"delete_credential":         OperationTypeDeleteCredential,
	"create_schedule":           OperationTypeCreateSchedule,
	"update_schedule":           OperationTypeUpdateSchedule,
	"update_maintenance_window": OperationTypeUpdateMaintenanceWindow,
	"delete_maintenance_window": OperationTypeDeleteMaintenanceWindow,
	"create_fleet_resource":     OperationTypeCreateFleetResource,
	"update_fleet_resource":     OperationTypeUpdateFleetResource,
	"delete_fleet_resource":     OperationTypeDeleteFleetResource,
	"create_fams_onboarding":    OperationTypeCreateFamsOnboarding,
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
		"CREATE_FLEET",
		"UPDATE_FLEET",
		"DELETE_FLEET",
		"CONFIRM_TARGET",
		"GENERATE_COMPLIANCE",
		"REQUEST_TARGET_DISCOVERY",
		"VALIDATE_RESOURCE",
		"CREATE_CREDENTIAL",
		"UPDATE_CREDENTIAL",
		"DELETE_CREDENTIAL",
		"CREATE_SCHEDULE",
		"UPDATE_SCHEDULE",
		"UPDATE_MAINTENANCE_WINDOW",
		"DELETE_MAINTENANCE_WINDOW",
		"CREATE_FLEET_RESOURCE",
		"UPDATE_FLEET_RESOURCE",
		"DELETE_FLEET_RESOURCE",
		"CREATE_FAMS_ONBOARDING",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
