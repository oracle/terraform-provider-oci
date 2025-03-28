// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateSddc                          OperationTypesEnum = "CREATE_SDDC"
	OperationTypesDeleteSddc                          OperationTypesEnum = "DELETE_SDDC"
	OperationTypesCreateCluster                       OperationTypesEnum = "CREATE_CLUSTER"
	OperationTypesDeleteCluster                       OperationTypesEnum = "DELETE_CLUSTER"
	OperationTypesCreateEsxiHost                      OperationTypesEnum = "CREATE_ESXI_HOST"
	OperationTypesDeleteEsxiHost                      OperationTypesEnum = "DELETE_ESXI_HOST"
	OperationTypesUpgradeHcx                          OperationTypesEnum = "UPGRADE_HCX"
	OperationTypesDowngradeHcx                        OperationTypesEnum = "DOWNGRADE_HCX"
	OperationTypesCancelDowngradeHcx                  OperationTypesEnum = "CANCEL_DOWNGRADE_HCX"
	OperationTypesRefreshHcxLicenseStatus             OperationTypesEnum = "REFRESH_HCX_LICENSE_STATUS"
	OperationTypesSwapBilling                         OperationTypesEnum = "SWAP_BILLING"
	OperationTypesReplaceHost                         OperationTypesEnum = "REPLACE_HOST"
	OperationTypesInPlaceUpgrade                      OperationTypesEnum = "IN_PLACE_UPGRADE"
	OperationTypesCreateDatastore                     OperationTypesEnum = "CREATE_DATASTORE"
	OperationTypesUpdateDatastore                     OperationTypesEnum = "UPDATE_DATASTORE"
	OperationTypesAddBlockVolumeToDatastore           OperationTypesEnum = "ADD_BLOCK_VOLUME_TO_DATASTORE"
	OperationTypesDeleteDatastore                     OperationTypesEnum = "DELETE_DATASTORE"
	OperationTypesCreateDatastoreCluster              OperationTypesEnum = "CREATE_DATASTORE_CLUSTER"
	OperationTypesUpdateDatastoreCluster              OperationTypesEnum = "UPDATE_DATASTORE_CLUSTER"
	OperationTypesAttachDatastoreClusterToEsxiHost    OperationTypesEnum = "ATTACH_DATASTORE_CLUSTER_TO_ESXI_HOST"
	OperationTypesAttachDatastoreClusterToCluster     OperationTypesEnum = "ATTACH_DATASTORE_CLUSTER_TO_CLUSTER"
	OperationTypesDetachDatastoreClusterFromEsxiHost  OperationTypesEnum = "DETACH_DATASTORE_CLUSTER_FROM_ESXI_HOST"
	OperationTypesDetachDatastoreClusterFromCluster   OperationTypesEnum = "DETACH_DATASTORE_CLUSTER_FROM_CLUSTER"
	OperationTypesDeleteDatastoreCluster              OperationTypesEnum = "DELETE_DATASTORE_CLUSTER"
	OperationTypesAddDatastoreToDatastoreCluster      OperationTypesEnum = "ADD_DATASTORE_TO_DATASTORE_CLUSTER"
	OperationTypesRemoveDatastoreFromDatastoreCluster OperationTypesEnum = "REMOVE_DATASTORE_FROM_DATASTORE_CLUSTER"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_SDDC":                             OperationTypesCreateSddc,
	"DELETE_SDDC":                             OperationTypesDeleteSddc,
	"CREATE_CLUSTER":                          OperationTypesCreateCluster,
	"DELETE_CLUSTER":                          OperationTypesDeleteCluster,
	"CREATE_ESXI_HOST":                        OperationTypesCreateEsxiHost,
	"DELETE_ESXI_HOST":                        OperationTypesDeleteEsxiHost,
	"UPGRADE_HCX":                             OperationTypesUpgradeHcx,
	"DOWNGRADE_HCX":                           OperationTypesDowngradeHcx,
	"CANCEL_DOWNGRADE_HCX":                    OperationTypesCancelDowngradeHcx,
	"REFRESH_HCX_LICENSE_STATUS":              OperationTypesRefreshHcxLicenseStatus,
	"SWAP_BILLING":                            OperationTypesSwapBilling,
	"REPLACE_HOST":                            OperationTypesReplaceHost,
	"IN_PLACE_UPGRADE":                        OperationTypesInPlaceUpgrade,
	"CREATE_DATASTORE":                        OperationTypesCreateDatastore,
	"UPDATE_DATASTORE":                        OperationTypesUpdateDatastore,
	"ADD_BLOCK_VOLUME_TO_DATASTORE":           OperationTypesAddBlockVolumeToDatastore,
	"DELETE_DATASTORE":                        OperationTypesDeleteDatastore,
	"CREATE_DATASTORE_CLUSTER":                OperationTypesCreateDatastoreCluster,
	"UPDATE_DATASTORE_CLUSTER":                OperationTypesUpdateDatastoreCluster,
	"ATTACH_DATASTORE_CLUSTER_TO_ESXI_HOST":   OperationTypesAttachDatastoreClusterToEsxiHost,
	"ATTACH_DATASTORE_CLUSTER_TO_CLUSTER":     OperationTypesAttachDatastoreClusterToCluster,
	"DETACH_DATASTORE_CLUSTER_FROM_ESXI_HOST": OperationTypesDetachDatastoreClusterFromEsxiHost,
	"DETACH_DATASTORE_CLUSTER_FROM_CLUSTER":   OperationTypesDetachDatastoreClusterFromCluster,
	"DELETE_DATASTORE_CLUSTER":                OperationTypesDeleteDatastoreCluster,
	"ADD_DATASTORE_TO_DATASTORE_CLUSTER":      OperationTypesAddDatastoreToDatastoreCluster,
	"REMOVE_DATASTORE_FROM_DATASTORE_CLUSTER": OperationTypesRemoveDatastoreFromDatastoreCluster,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"create_sddc":                             OperationTypesCreateSddc,
	"delete_sddc":                             OperationTypesDeleteSddc,
	"create_cluster":                          OperationTypesCreateCluster,
	"delete_cluster":                          OperationTypesDeleteCluster,
	"create_esxi_host":                        OperationTypesCreateEsxiHost,
	"delete_esxi_host":                        OperationTypesDeleteEsxiHost,
	"upgrade_hcx":                             OperationTypesUpgradeHcx,
	"downgrade_hcx":                           OperationTypesDowngradeHcx,
	"cancel_downgrade_hcx":                    OperationTypesCancelDowngradeHcx,
	"refresh_hcx_license_status":              OperationTypesRefreshHcxLicenseStatus,
	"swap_billing":                            OperationTypesSwapBilling,
	"replace_host":                            OperationTypesReplaceHost,
	"in_place_upgrade":                        OperationTypesInPlaceUpgrade,
	"create_datastore":                        OperationTypesCreateDatastore,
	"update_datastore":                        OperationTypesUpdateDatastore,
	"add_block_volume_to_datastore":           OperationTypesAddBlockVolumeToDatastore,
	"delete_datastore":                        OperationTypesDeleteDatastore,
	"create_datastore_cluster":                OperationTypesCreateDatastoreCluster,
	"update_datastore_cluster":                OperationTypesUpdateDatastoreCluster,
	"attach_datastore_cluster_to_esxi_host":   OperationTypesAttachDatastoreClusterToEsxiHost,
	"attach_datastore_cluster_to_cluster":     OperationTypesAttachDatastoreClusterToCluster,
	"detach_datastore_cluster_from_esxi_host": OperationTypesDetachDatastoreClusterFromEsxiHost,
	"detach_datastore_cluster_from_cluster":   OperationTypesDetachDatastoreClusterFromCluster,
	"delete_datastore_cluster":                OperationTypesDeleteDatastoreCluster,
	"add_datastore_to_datastore_cluster":      OperationTypesAddDatastoreToDatastoreCluster,
	"remove_datastore_from_datastore_cluster": OperationTypesRemoveDatastoreFromDatastoreCluster,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"CREATE_SDDC",
		"DELETE_SDDC",
		"CREATE_CLUSTER",
		"DELETE_CLUSTER",
		"CREATE_ESXI_HOST",
		"DELETE_ESXI_HOST",
		"UPGRADE_HCX",
		"DOWNGRADE_HCX",
		"CANCEL_DOWNGRADE_HCX",
		"REFRESH_HCX_LICENSE_STATUS",
		"SWAP_BILLING",
		"REPLACE_HOST",
		"IN_PLACE_UPGRADE",
		"CREATE_DATASTORE",
		"UPDATE_DATASTORE",
		"ADD_BLOCK_VOLUME_TO_DATASTORE",
		"DELETE_DATASTORE",
		"CREATE_DATASTORE_CLUSTER",
		"UPDATE_DATASTORE_CLUSTER",
		"ATTACH_DATASTORE_CLUSTER_TO_ESXI_HOST",
		"ATTACH_DATASTORE_CLUSTER_TO_CLUSTER",
		"DETACH_DATASTORE_CLUSTER_FROM_ESXI_HOST",
		"DETACH_DATASTORE_CLUSTER_FROM_CLUSTER",
		"DELETE_DATASTORE_CLUSTER",
		"ADD_DATASTORE_TO_DATASTORE_CLUSTER",
		"REMOVE_DATASTORE_FROM_DATASTORE_CLUSTER",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
