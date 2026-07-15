// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateInstance                        OperationTypeEnum = "CREATE_INSTANCE"
	OperationTypeDeleteInstance                        OperationTypeEnum = "DELETE_INSTANCE"
	OperationTypeUpdateInstance                        OperationTypeEnum = "UPDATE_INSTANCE"
	OperationTypeChangeInstanceCompartment             OperationTypeEnum = "CHANGE_INSTANCE_COMPARTMENT"
	OperationTypeCreateBaseVmClusterNetwork            OperationTypeEnum = "CREATE_BASE_VM_CLUSTER_NETWORK"
	OperationTypeDeleteBaseVmClusterNetwork            OperationTypeEnum = "DELETE_BASE_VM_CLUSTER_NETWORK"
	OperationTypeValidateBaseVmClusterNetwork          OperationTypeEnum = "VALIDATE_BASE_VM_CLUSTER_NETWORK"
	OperationTypeUpdateBaseVmClusterNetwork            OperationTypeEnum = "UPDATE_BASE_VM_CLUSTER_NETWORK"
	OperationTypeChangeBaseVmClusterNetworkCompartment OperationTypeEnum = "CHANGE_BASE_VM_CLUSTER_NETWORK_COMPARTMENT"
	OperationTypeStartBaseVm                           OperationTypeEnum = "START_BASE_VM"
	OperationTypeStopBaseVm                            OperationTypeEnum = "STOP_BASE_VM"
	OperationTypeStartInstance                         OperationTypeEnum = "START_INSTANCE"
	OperationTypeStopInstance                          OperationTypeEnum = "STOP_INSTANCE"
	OperationTypeScaleInstance                         OperationTypeEnum = "SCALE_INSTANCE"
	OperationTypeRestartInstance                       OperationTypeEnum = "RESTART_INSTANCE"
	OperationTypeRestartBaseVm                         OperationTypeEnum = "RESTART_BASE_VM"
	OperationTypeMigrateInstance                       OperationTypeEnum = "MIGRATE_INSTANCE"
	OperationTypeScaleBaseVm                           OperationTypeEnum = "SCALE_BASE_VM"
	OperationTypeUpdateSoftwareImage                   OperationTypeEnum = "UPDATE_SOFTWARE_IMAGE"
	OperationTypeChangeSoftwareImageCompartment        OperationTypeEnum = "CHANGE_SOFTWARE_IMAGE_COMPARTMENT"
	OperationTypeValidateProxy                         OperationTypeEnum = "VALIDATE_PROXY"
	OperationTypeCreateBaseInfraCapacity               OperationTypeEnum = "CREATE_BASE_INFRA_CAPACITY"
	OperationTypeUpdateBaseInfraCapacity               OperationTypeEnum = "UPDATE_BASE_INFRA_CAPACITY"
	OperationTypeDeleteBaseInfraCapacity               OperationTypeEnum = "DELETE_BASE_INFRA_CAPACITY"
	OperationTypeActivateBaseInfrastructure            OperationTypeEnum = "ACTIVATE_BASE_INFRASTRUCTURE"
	OperationTypeCreateBaseInfrastructure              OperationTypeEnum = "CREATE_BASE_INFRASTRUCTURE"
	OperationTypeDeleteBaseInfrastructure              OperationTypeEnum = "DELETE_BASE_INFRASTRUCTURE"
	OperationTypeUpdateBaseInfrastructure              OperationTypeEnum = "UPDATE_BASE_INFRASTRUCTURE"
	OperationTypeUpdateBaseInfrastructureSoftware      OperationTypeEnum = "UPDATE_BASE_INFRASTRUCTURE_SOFTWARE"
	OperationTypeValidateBaseInfrastructure            OperationTypeEnum = "VALIDATE_BASE_INFRASTRUCTURE"
	OperationTypeChangeBaseInfrastructureCompartment   OperationTypeEnum = "CHANGE_BASE_INFRASTRUCTURE_COMPARTMENT"
	OperationTypeAnalyzeInfrastructure                 OperationTypeEnum = "ANALYZE_INFRASTRUCTURE"
	OperationTypeAnalyzeBaseInfraNetwork               OperationTypeEnum = "ANALYZE_BASE_INFRA_NETWORK"
	OperationTypeInfrastructureSreValidation           OperationTypeEnum = "INFRASTRUCTURE_SRE_VALIDATION"
	OperationTypeCreateMaintenanceRun                  OperationTypeEnum = "CREATE_MAINTENANCE_RUN"
	OperationTypeUpdateMaintenanceRun                  OperationTypeEnum = "UPDATE_MAINTENANCE_RUN"
	OperationTypeRescheduleMaintenanceRun              OperationTypeEnum = "RESCHEDULE_MAINTENANCE_RUN"
	OperationTypeInfrastructureScaleStorage            OperationTypeEnum = "INFRASTRUCTURE_SCALE_STORAGE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_INSTANCE":                            OperationTypeCreateInstance,
	"DELETE_INSTANCE":                            OperationTypeDeleteInstance,
	"UPDATE_INSTANCE":                            OperationTypeUpdateInstance,
	"CHANGE_INSTANCE_COMPARTMENT":                OperationTypeChangeInstanceCompartment,
	"CREATE_BASE_VM_CLUSTER_NETWORK":             OperationTypeCreateBaseVmClusterNetwork,
	"DELETE_BASE_VM_CLUSTER_NETWORK":             OperationTypeDeleteBaseVmClusterNetwork,
	"VALIDATE_BASE_VM_CLUSTER_NETWORK":           OperationTypeValidateBaseVmClusterNetwork,
	"UPDATE_BASE_VM_CLUSTER_NETWORK":             OperationTypeUpdateBaseVmClusterNetwork,
	"CHANGE_BASE_VM_CLUSTER_NETWORK_COMPARTMENT": OperationTypeChangeBaseVmClusterNetworkCompartment,
	"START_BASE_VM":                              OperationTypeStartBaseVm,
	"STOP_BASE_VM":                               OperationTypeStopBaseVm,
	"START_INSTANCE":                             OperationTypeStartInstance,
	"STOP_INSTANCE":                              OperationTypeStopInstance,
	"SCALE_INSTANCE":                             OperationTypeScaleInstance,
	"RESTART_INSTANCE":                           OperationTypeRestartInstance,
	"RESTART_BASE_VM":                            OperationTypeRestartBaseVm,
	"MIGRATE_INSTANCE":                           OperationTypeMigrateInstance,
	"SCALE_BASE_VM":                              OperationTypeScaleBaseVm,
	"UPDATE_SOFTWARE_IMAGE":                      OperationTypeUpdateSoftwareImage,
	"CHANGE_SOFTWARE_IMAGE_COMPARTMENT":          OperationTypeChangeSoftwareImageCompartment,
	"VALIDATE_PROXY":                             OperationTypeValidateProxy,
	"CREATE_BASE_INFRA_CAPACITY":                 OperationTypeCreateBaseInfraCapacity,
	"UPDATE_BASE_INFRA_CAPACITY":                 OperationTypeUpdateBaseInfraCapacity,
	"DELETE_BASE_INFRA_CAPACITY":                 OperationTypeDeleteBaseInfraCapacity,
	"ACTIVATE_BASE_INFRASTRUCTURE":               OperationTypeActivateBaseInfrastructure,
	"CREATE_BASE_INFRASTRUCTURE":                 OperationTypeCreateBaseInfrastructure,
	"DELETE_BASE_INFRASTRUCTURE":                 OperationTypeDeleteBaseInfrastructure,
	"UPDATE_BASE_INFRASTRUCTURE":                 OperationTypeUpdateBaseInfrastructure,
	"UPDATE_BASE_INFRASTRUCTURE_SOFTWARE":        OperationTypeUpdateBaseInfrastructureSoftware,
	"VALIDATE_BASE_INFRASTRUCTURE":               OperationTypeValidateBaseInfrastructure,
	"CHANGE_BASE_INFRASTRUCTURE_COMPARTMENT":     OperationTypeChangeBaseInfrastructureCompartment,
	"ANALYZE_INFRASTRUCTURE":                     OperationTypeAnalyzeInfrastructure,
	"ANALYZE_BASE_INFRA_NETWORK":                 OperationTypeAnalyzeBaseInfraNetwork,
	"INFRASTRUCTURE_SRE_VALIDATION":              OperationTypeInfrastructureSreValidation,
	"CREATE_MAINTENANCE_RUN":                     OperationTypeCreateMaintenanceRun,
	"UPDATE_MAINTENANCE_RUN":                     OperationTypeUpdateMaintenanceRun,
	"RESCHEDULE_MAINTENANCE_RUN":                 OperationTypeRescheduleMaintenanceRun,
	"INFRASTRUCTURE_SCALE_STORAGE":               OperationTypeInfrastructureScaleStorage,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_instance":                            OperationTypeCreateInstance,
	"delete_instance":                            OperationTypeDeleteInstance,
	"update_instance":                            OperationTypeUpdateInstance,
	"change_instance_compartment":                OperationTypeChangeInstanceCompartment,
	"create_base_vm_cluster_network":             OperationTypeCreateBaseVmClusterNetwork,
	"delete_base_vm_cluster_network":             OperationTypeDeleteBaseVmClusterNetwork,
	"validate_base_vm_cluster_network":           OperationTypeValidateBaseVmClusterNetwork,
	"update_base_vm_cluster_network":             OperationTypeUpdateBaseVmClusterNetwork,
	"change_base_vm_cluster_network_compartment": OperationTypeChangeBaseVmClusterNetworkCompartment,
	"start_base_vm":                              OperationTypeStartBaseVm,
	"stop_base_vm":                               OperationTypeStopBaseVm,
	"start_instance":                             OperationTypeStartInstance,
	"stop_instance":                              OperationTypeStopInstance,
	"scale_instance":                             OperationTypeScaleInstance,
	"restart_instance":                           OperationTypeRestartInstance,
	"restart_base_vm":                            OperationTypeRestartBaseVm,
	"migrate_instance":                           OperationTypeMigrateInstance,
	"scale_base_vm":                              OperationTypeScaleBaseVm,
	"update_software_image":                      OperationTypeUpdateSoftwareImage,
	"change_software_image_compartment":          OperationTypeChangeSoftwareImageCompartment,
	"validate_proxy":                             OperationTypeValidateProxy,
	"create_base_infra_capacity":                 OperationTypeCreateBaseInfraCapacity,
	"update_base_infra_capacity":                 OperationTypeUpdateBaseInfraCapacity,
	"delete_base_infra_capacity":                 OperationTypeDeleteBaseInfraCapacity,
	"activate_base_infrastructure":               OperationTypeActivateBaseInfrastructure,
	"create_base_infrastructure":                 OperationTypeCreateBaseInfrastructure,
	"delete_base_infrastructure":                 OperationTypeDeleteBaseInfrastructure,
	"update_base_infrastructure":                 OperationTypeUpdateBaseInfrastructure,
	"update_base_infrastructure_software":        OperationTypeUpdateBaseInfrastructureSoftware,
	"validate_base_infrastructure":               OperationTypeValidateBaseInfrastructure,
	"change_base_infrastructure_compartment":     OperationTypeChangeBaseInfrastructureCompartment,
	"analyze_infrastructure":                     OperationTypeAnalyzeInfrastructure,
	"analyze_base_infra_network":                 OperationTypeAnalyzeBaseInfraNetwork,
	"infrastructure_sre_validation":              OperationTypeInfrastructureSreValidation,
	"create_maintenance_run":                     OperationTypeCreateMaintenanceRun,
	"update_maintenance_run":                     OperationTypeUpdateMaintenanceRun,
	"reschedule_maintenance_run":                 OperationTypeRescheduleMaintenanceRun,
	"infrastructure_scale_storage":               OperationTypeInfrastructureScaleStorage,
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
		"CREATE_INSTANCE",
		"DELETE_INSTANCE",
		"UPDATE_INSTANCE",
		"CHANGE_INSTANCE_COMPARTMENT",
		"CREATE_BASE_VM_CLUSTER_NETWORK",
		"DELETE_BASE_VM_CLUSTER_NETWORK",
		"VALIDATE_BASE_VM_CLUSTER_NETWORK",
		"UPDATE_BASE_VM_CLUSTER_NETWORK",
		"CHANGE_BASE_VM_CLUSTER_NETWORK_COMPARTMENT",
		"START_BASE_VM",
		"STOP_BASE_VM",
		"START_INSTANCE",
		"STOP_INSTANCE",
		"SCALE_INSTANCE",
		"RESTART_INSTANCE",
		"RESTART_BASE_VM",
		"MIGRATE_INSTANCE",
		"SCALE_BASE_VM",
		"UPDATE_SOFTWARE_IMAGE",
		"CHANGE_SOFTWARE_IMAGE_COMPARTMENT",
		"VALIDATE_PROXY",
		"CREATE_BASE_INFRA_CAPACITY",
		"UPDATE_BASE_INFRA_CAPACITY",
		"DELETE_BASE_INFRA_CAPACITY",
		"ACTIVATE_BASE_INFRASTRUCTURE",
		"CREATE_BASE_INFRASTRUCTURE",
		"DELETE_BASE_INFRASTRUCTURE",
		"UPDATE_BASE_INFRASTRUCTURE",
		"UPDATE_BASE_INFRASTRUCTURE_SOFTWARE",
		"VALIDATE_BASE_INFRASTRUCTURE",
		"CHANGE_BASE_INFRASTRUCTURE_COMPARTMENT",
		"ANALYZE_INFRASTRUCTURE",
		"ANALYZE_BASE_INFRA_NETWORK",
		"INFRASTRUCTURE_SRE_VALIDATION",
		"CREATE_MAINTENANCE_RUN",
		"UPDATE_MAINTENANCE_RUN",
		"RESCHEDULE_MAINTENANCE_RUN",
		"INFRASTRUCTURE_SCALE_STORAGE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
