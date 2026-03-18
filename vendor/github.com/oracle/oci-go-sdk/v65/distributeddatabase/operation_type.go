// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDistributedDb                            OperationTypeEnum = "CREATE_DISTRIBUTED_DB"
	OperationTypePatchDistributedDbInsertShards                 OperationTypeEnum = "PATCH_DISTRIBUTED_DB_INSERT_SHARDS"
	OperationTypePatchDistributedDbRemoveShards                 OperationTypeEnum = "PATCH_DISTRIBUTED_DB_REMOVE_SHARDS"
	OperationTypePatchDistributedDbMergeCatalogShards           OperationTypeEnum = "PATCH_DISTRIBUTED_DB_MERGE_CATALOG_SHARDS"
	OperationTypeDeleteDistributedDb                            OperationTypeEnum = "DELETE_DISTRIBUTED_DB"
	OperationTypeChangeDistributedDbCompartment                 OperationTypeEnum = "CHANGE_DISTRIBUTED_DB_COMPARTMENT"
	OperationTypeConfigureDistributedDbGsms                     OperationTypeEnum = "CONFIGURE_DISTRIBUTED_DB_GSMS"
	OperationTypeConfigureDistributedDbSharding                 OperationTypeEnum = "CONFIGURE_DISTRIBUTED_DB_SHARDING"
	OperationTypeGenerateDistributedDbGsmCertSigningReq         OperationTypeEnum = "GENERATE_DISTRIBUTED_DB_GSM_CERT_SIGNING_REQ"
	OperationTypeRotateDistributedDbPasswords                   OperationTypeEnum = "ROTATE_DISTRIBUTED_DB_PASSWORDS"
	OperationTypeStartDistributedDb                             OperationTypeEnum = "START_DISTRIBUTED_DB"
	OperationTypeStopDistributedDb                              OperationTypeEnum = "STOP_DISTRIBUTED_DB"
	OperationTypeStartDistributedDbShard                        OperationTypeEnum = "START_DISTRIBUTED_DB_SHARD"
	OperationTypeStartDistributedDbCatalog                      OperationTypeEnum = "START_DISTRIBUTED_DB_CATALOG"
	OperationTypeStopDistributedDbShard                         OperationTypeEnum = "STOP_DISTRIBUTED_DB_SHARD"
	OperationTypeStopDistributedDbCatalog                       OperationTypeEnum = "STOP_DISTRIBUTED_DB_CATALOG"
	OperationTypeUploadDistributedDbSignedCertAndGenerateWallet OperationTypeEnum = "UPLOAD_DISTRIBUTED_DB_SIGNED_CERT_AND_GENERATE_WALLET"
	OperationTypeCreateDistributedDbShard                       OperationTypeEnum = "CREATE_DISTRIBUTED_DB_SHARD"
	OperationTypeCreateDistributedDbCatalog                     OperationTypeEnum = "CREATE_DISTRIBUTED_DB_CATALOG"
	OperationTypeCreateDistributedDbGsm                         OperationTypeEnum = "CREATE_DISTRIBUTED_DB_GSM"
	OperationTypeUpdateDistributedDbCatalogShards               OperationTypeEnum = "UPDATE_DISTRIBUTED_DB_CATALOG_SHARDS"
	OperationTypeValidateNetwork                                OperationTypeEnum = "VALIDATE_NETWORK"
	OperationTypeExecuteValidateNetworkTests                    OperationTypeEnum = "EXECUTE_VALIDATE_NETWORK_TESTS"
	OperationTypeCreatePrivateEndpoint                          OperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	OperationTypeDeletePrivateEndpoint                          OperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	OperationTypeChangePrivateEndpointCompartment               OperationTypeEnum = "CHANGE_PRIVATE_ENDPOINT_COMPARTMENT"
	OperationTypeReinstateProxyInstance                         OperationTypeEnum = "REINSTATE_PROXY_INSTANCE"
	OperationTypeDeleteDistributedDbShard                       OperationTypeEnum = "DELETE_DISTRIBUTED_DB_SHARD"
	OperationTypeDeleteDistributedDbCatalog                     OperationTypeEnum = "DELETE_DISTRIBUTED_DB_CATALOG"
	OperationTypeDeleteDistributedDbGsm                         OperationTypeEnum = "DELETE_DISTRIBUTED_DB_GSM"
	OperationTypeAddGdscontrolNode                              OperationTypeEnum = "ADD_GDSCONTROL_NODE"
	OperationTypeChangeDistributedDbBackupConfig                OperationTypeEnum = "CHANGE_DISTRIBUTED_DB_BACKUP_CONFIG"
	OperationTypeAddDistributedDbDg                             OperationTypeEnum = "ADD_DISTRIBUTED_DB_DG"
	OperationTypeRemoveDistributedDbDg                          OperationTypeEnum = "REMOVE_DISTRIBUTED_DB_DG"
	OperationTypeConfigureDistributedDbGsmsWallet               OperationTypeEnum = "CONFIGURE_DISTRIBUTED_DB_GSMS_WALLET"
	OperationTypeValidateCaBundle                               OperationTypeEnum = "VALIDATE_CA_BUNDLE"
	OperationTypeRecreateFailedResource                         OperationTypeEnum = "RECREATE_FAILED_RESOURCE"
	OperationTypeMoveReplicationUnits                           OperationTypeEnum = "MOVE_REPLICATION_UNITS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DISTRIBUTED_DB":                                 OperationTypeCreateDistributedDb,
	"PATCH_DISTRIBUTED_DB_INSERT_SHARDS":                    OperationTypePatchDistributedDbInsertShards,
	"PATCH_DISTRIBUTED_DB_REMOVE_SHARDS":                    OperationTypePatchDistributedDbRemoveShards,
	"PATCH_DISTRIBUTED_DB_MERGE_CATALOG_SHARDS":             OperationTypePatchDistributedDbMergeCatalogShards,
	"DELETE_DISTRIBUTED_DB":                                 OperationTypeDeleteDistributedDb,
	"CHANGE_DISTRIBUTED_DB_COMPARTMENT":                     OperationTypeChangeDistributedDbCompartment,
	"CONFIGURE_DISTRIBUTED_DB_GSMS":                         OperationTypeConfigureDistributedDbGsms,
	"CONFIGURE_DISTRIBUTED_DB_SHARDING":                     OperationTypeConfigureDistributedDbSharding,
	"GENERATE_DISTRIBUTED_DB_GSM_CERT_SIGNING_REQ":          OperationTypeGenerateDistributedDbGsmCertSigningReq,
	"ROTATE_DISTRIBUTED_DB_PASSWORDS":                       OperationTypeRotateDistributedDbPasswords,
	"START_DISTRIBUTED_DB":                                  OperationTypeStartDistributedDb,
	"STOP_DISTRIBUTED_DB":                                   OperationTypeStopDistributedDb,
	"START_DISTRIBUTED_DB_SHARD":                            OperationTypeStartDistributedDbShard,
	"START_DISTRIBUTED_DB_CATALOG":                          OperationTypeStartDistributedDbCatalog,
	"STOP_DISTRIBUTED_DB_SHARD":                             OperationTypeStopDistributedDbShard,
	"STOP_DISTRIBUTED_DB_CATALOG":                           OperationTypeStopDistributedDbCatalog,
	"UPLOAD_DISTRIBUTED_DB_SIGNED_CERT_AND_GENERATE_WALLET": OperationTypeUploadDistributedDbSignedCertAndGenerateWallet,
	"CREATE_DISTRIBUTED_DB_SHARD":                           OperationTypeCreateDistributedDbShard,
	"CREATE_DISTRIBUTED_DB_CATALOG":                         OperationTypeCreateDistributedDbCatalog,
	"CREATE_DISTRIBUTED_DB_GSM":                             OperationTypeCreateDistributedDbGsm,
	"UPDATE_DISTRIBUTED_DB_CATALOG_SHARDS":                  OperationTypeUpdateDistributedDbCatalogShards,
	"VALIDATE_NETWORK":                                      OperationTypeValidateNetwork,
	"EXECUTE_VALIDATE_NETWORK_TESTS":                        OperationTypeExecuteValidateNetworkTests,
	"CREATE_PRIVATE_ENDPOINT":                               OperationTypeCreatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT":                               OperationTypeDeletePrivateEndpoint,
	"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT":                   OperationTypeChangePrivateEndpointCompartment,
	"REINSTATE_PROXY_INSTANCE":                              OperationTypeReinstateProxyInstance,
	"DELETE_DISTRIBUTED_DB_SHARD":                           OperationTypeDeleteDistributedDbShard,
	"DELETE_DISTRIBUTED_DB_CATALOG":                         OperationTypeDeleteDistributedDbCatalog,
	"DELETE_DISTRIBUTED_DB_GSM":                             OperationTypeDeleteDistributedDbGsm,
	"ADD_GDSCONTROL_NODE":                                   OperationTypeAddGdscontrolNode,
	"CHANGE_DISTRIBUTED_DB_BACKUP_CONFIG":                   OperationTypeChangeDistributedDbBackupConfig,
	"ADD_DISTRIBUTED_DB_DG":                                 OperationTypeAddDistributedDbDg,
	"REMOVE_DISTRIBUTED_DB_DG":                              OperationTypeRemoveDistributedDbDg,
	"CONFIGURE_DISTRIBUTED_DB_GSMS_WALLET":                  OperationTypeConfigureDistributedDbGsmsWallet,
	"VALIDATE_CA_BUNDLE":                                    OperationTypeValidateCaBundle,
	"RECREATE_FAILED_RESOURCE":                              OperationTypeRecreateFailedResource,
	"MOVE_REPLICATION_UNITS":                                OperationTypeMoveReplicationUnits,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_distributed_db":                                 OperationTypeCreateDistributedDb,
	"patch_distributed_db_insert_shards":                    OperationTypePatchDistributedDbInsertShards,
	"patch_distributed_db_remove_shards":                    OperationTypePatchDistributedDbRemoveShards,
	"patch_distributed_db_merge_catalog_shards":             OperationTypePatchDistributedDbMergeCatalogShards,
	"delete_distributed_db":                                 OperationTypeDeleteDistributedDb,
	"change_distributed_db_compartment":                     OperationTypeChangeDistributedDbCompartment,
	"configure_distributed_db_gsms":                         OperationTypeConfigureDistributedDbGsms,
	"configure_distributed_db_sharding":                     OperationTypeConfigureDistributedDbSharding,
	"generate_distributed_db_gsm_cert_signing_req":          OperationTypeGenerateDistributedDbGsmCertSigningReq,
	"rotate_distributed_db_passwords":                       OperationTypeRotateDistributedDbPasswords,
	"start_distributed_db":                                  OperationTypeStartDistributedDb,
	"stop_distributed_db":                                   OperationTypeStopDistributedDb,
	"start_distributed_db_shard":                            OperationTypeStartDistributedDbShard,
	"start_distributed_db_catalog":                          OperationTypeStartDistributedDbCatalog,
	"stop_distributed_db_shard":                             OperationTypeStopDistributedDbShard,
	"stop_distributed_db_catalog":                           OperationTypeStopDistributedDbCatalog,
	"upload_distributed_db_signed_cert_and_generate_wallet": OperationTypeUploadDistributedDbSignedCertAndGenerateWallet,
	"create_distributed_db_shard":                           OperationTypeCreateDistributedDbShard,
	"create_distributed_db_catalog":                         OperationTypeCreateDistributedDbCatalog,
	"create_distributed_db_gsm":                             OperationTypeCreateDistributedDbGsm,
	"update_distributed_db_catalog_shards":                  OperationTypeUpdateDistributedDbCatalogShards,
	"validate_network":                                      OperationTypeValidateNetwork,
	"execute_validate_network_tests":                        OperationTypeExecuteValidateNetworkTests,
	"create_private_endpoint":                               OperationTypeCreatePrivateEndpoint,
	"delete_private_endpoint":                               OperationTypeDeletePrivateEndpoint,
	"change_private_endpoint_compartment":                   OperationTypeChangePrivateEndpointCompartment,
	"reinstate_proxy_instance":                              OperationTypeReinstateProxyInstance,
	"delete_distributed_db_shard":                           OperationTypeDeleteDistributedDbShard,
	"delete_distributed_db_catalog":                         OperationTypeDeleteDistributedDbCatalog,
	"delete_distributed_db_gsm":                             OperationTypeDeleteDistributedDbGsm,
	"add_gdscontrol_node":                                   OperationTypeAddGdscontrolNode,
	"change_distributed_db_backup_config":                   OperationTypeChangeDistributedDbBackupConfig,
	"add_distributed_db_dg":                                 OperationTypeAddDistributedDbDg,
	"remove_distributed_db_dg":                              OperationTypeRemoveDistributedDbDg,
	"configure_distributed_db_gsms_wallet":                  OperationTypeConfigureDistributedDbGsmsWallet,
	"validate_ca_bundle":                                    OperationTypeValidateCaBundle,
	"recreate_failed_resource":                              OperationTypeRecreateFailedResource,
	"move_replication_units":                                OperationTypeMoveReplicationUnits,
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
		"CREATE_DISTRIBUTED_DB",
		"PATCH_DISTRIBUTED_DB_INSERT_SHARDS",
		"PATCH_DISTRIBUTED_DB_REMOVE_SHARDS",
		"PATCH_DISTRIBUTED_DB_MERGE_CATALOG_SHARDS",
		"DELETE_DISTRIBUTED_DB",
		"CHANGE_DISTRIBUTED_DB_COMPARTMENT",
		"CONFIGURE_DISTRIBUTED_DB_GSMS",
		"CONFIGURE_DISTRIBUTED_DB_SHARDING",
		"GENERATE_DISTRIBUTED_DB_GSM_CERT_SIGNING_REQ",
		"ROTATE_DISTRIBUTED_DB_PASSWORDS",
		"START_DISTRIBUTED_DB",
		"STOP_DISTRIBUTED_DB",
		"START_DISTRIBUTED_DB_SHARD",
		"START_DISTRIBUTED_DB_CATALOG",
		"STOP_DISTRIBUTED_DB_SHARD",
		"STOP_DISTRIBUTED_DB_CATALOG",
		"UPLOAD_DISTRIBUTED_DB_SIGNED_CERT_AND_GENERATE_WALLET",
		"CREATE_DISTRIBUTED_DB_SHARD",
		"CREATE_DISTRIBUTED_DB_CATALOG",
		"CREATE_DISTRIBUTED_DB_GSM",
		"UPDATE_DISTRIBUTED_DB_CATALOG_SHARDS",
		"VALIDATE_NETWORK",
		"EXECUTE_VALIDATE_NETWORK_TESTS",
		"CREATE_PRIVATE_ENDPOINT",
		"DELETE_PRIVATE_ENDPOINT",
		"CHANGE_PRIVATE_ENDPOINT_COMPARTMENT",
		"REINSTATE_PROXY_INSTANCE",
		"DELETE_DISTRIBUTED_DB_SHARD",
		"DELETE_DISTRIBUTED_DB_CATALOG",
		"DELETE_DISTRIBUTED_DB_GSM",
		"ADD_GDSCONTROL_NODE",
		"CHANGE_DISTRIBUTED_DB_BACKUP_CONFIG",
		"ADD_DISTRIBUTED_DB_DG",
		"REMOVE_DISTRIBUTED_DB_DG",
		"CONFIGURE_DISTRIBUTED_DB_GSMS_WALLET",
		"VALIDATE_CA_BUNDLE",
		"RECREATE_FAILED_RESOURCE",
		"MOVE_REPLICATION_UNITS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
