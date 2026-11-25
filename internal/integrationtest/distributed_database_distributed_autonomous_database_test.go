// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkterraform "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	ddb "github.com/oracle/terraform-provider-oci/internal/service/distributed_database"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DistributedDatabaseDistributedAutonomousDatabaseRequiredOnlyResource = DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation)

	DistributedDatabaseDistributedAutonomousDatabaseResourceConfig = DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedAutonomousDatabaseRepresentation)

	DistributedDatabaseDistributedAutonomousDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"distributed_autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`},
		"metadata":                           acctest.Representation{RepType: acctest.Optional, Create: `metadata`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseMetadataSingularDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedAutonomousDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"db_deployment_type":                  acctest.Representation{RepType: acctest.Optional, Create: `dbDeploymentType`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `gdn-name`, Update: `displayName2`},
		"is_auto_resource_management_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: `metadata`},
		"private_endpoint_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                              acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseDataSourceFilterRepresentation}}
	DistributedDatabaseDistributedAutonomousDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`}},
	}
	DistributedDatabaseDistributedAutonomousDatabaseMetadataDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedAutonomousDatabaseRepresentation = map[string]interface{}{
		"catalog_details":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsRepresentation},
		"character_set":                              acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"compartment_id":                             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_version":                           acctest.Representation{RepType: acctest.Required, Create: `26ai`},
		"db_workload_type":                           acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"display_name":                               acctest.Representation{RepType: acctest.Required, Create: `gdn-name`, Update: `displayName2`},
		"listener_port":                              acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ncharacter_set":                             acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"ons_port_local":                             acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ons_port_remote":                            acctest.Representation{RepType: acctest.Required, Create: `10`},
		"prefix":                                     acctest.Representation{RepType: acctest.Required, Create: `aaa`},
		"private_endpoint_ids":                       acctest.Representation{RepType: acctest.Required, Create: []string{`privateEndpointIds`}},
		"composite_data_guard_shard_spaces":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesRepresentation},
		"composite_raft_shard_spaces":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeRaftShardSpacesRepresentation},
		"db_backup_config":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigRepresentation},
		"defined_tags":                               acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"listener_port_tls":                          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"notification_topic_ids":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`notificationTopicIds`}, Update: []string{`notificationTopicIds2`}},
		"patch_operations":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabasePatchOperationsRepresentation},
		"system_chunk_count":                         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"system_data_guard_databases":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesRepresentation},
		"system_raft_clusters":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseSystemRaftClustersRepresentation},
		"system_raft_replication_unit_count":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"user_shard_spaces":                          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesRepresentation},
		"vcn_nsg_ids":                                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseVcnNsgIdsRepresentation},
		"add_gds_control_node_trigger":               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"change_db_backup_config_trigger":            acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"configure_auto_resource_management_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"configure_gsm_wallet_trigger":               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"configure_sharding_trigger":                 acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"generate_wallet_trigger":                    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"move_replication_unit_trigger":              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"recreate_failed_resource_trigger":           acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"scale_gsms_trigger":                         acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"start_database_trigger":                     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"stop_database_trigger":                      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"validate_ca_bundle_trigger":                 acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"validate_network_trigger":                   acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsRepresentation = map[string]interface{}{
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesRepresentation = map[string]interface{}{
		"chunk_count":         acctest.Representation{RepType: acctest.Required, Create: `10`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `west`},
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeRaftShardSpacesRepresentation = map[string]interface{}{
		"chunk_count":            acctest.Representation{RepType: acctest.Required, Create: `10`},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `west`},
		"raft_clusters":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeRaftShardSpacesRaftClustersRepresentation},
		"replication_unit_count": acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedAutonomousDatabasePatchOperationsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"value":     acctest.Representation{RepType: acctest.Optional, Create: `{}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesRepresentation = map[string]interface{}{
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedAutonomousDatabaseSystemRaftClustersRepresentation = map[string]interface{}{
		"databases":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseSystemRaftClustersDatabasesRepresentation},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `west`},
		"replication_factor": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ru_mode":            acctest.Representation{RepType: acctest.Required, Create: `READ_WRITE`},
		"witness_count":      acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesRepresentation = map[string]interface{}{
		"name":                acctest.Representation{RepType: acctest.Required, Create: `west`},
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedAutonomousDatabaseVcnNsgIdsRepresentation = map[string]interface{}{
		"nsg_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`nsgIds`}},
		"vcn_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsOriginalReplicaRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsOriginalReplicaDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsDataGuardReplicasRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsDataGuardReplicasDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesOriginalReplicaRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesDataGuardReplicasRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeRaftShardSpacesRaftClustersRepresentation = map[string]interface{}{
		"databases":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCompositeRaftShardSpacesRaftClustersDatabasesRepresentation},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `west`},
		"replication_factor": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ru_mode":            acctest.Representation{RepType: acctest.Required, Create: `READ_WRITE`},
		"witness_count":      acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"dbrs_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_policy.test_policy.id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `id`},
		"internet_proxy": acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"is_remote":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"remote_region":  acctest.Representation{RepType: acctest.Optional, Create: `remoteRegion`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesOriginalReplicaRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesDataGuardReplicasRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseSystemRaftClustersDatabasesRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesOriginalReplicaRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesOriginalReplicaDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesDataGuardReplicasRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesDataGuardReplicasDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsOriginalReplicaDatabaseRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsDataGuardReplicasDatabaseRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCompositeRaftShardSpacesRaftClustersDatabasesRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesOriginalReplicaDatabaseRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseUserShardSpacesDataGuardReplicasDatabaseRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"compute_count":                             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"data_storage_size_in_gbs":                  acctest.Representation{RepType: acctest.Required, Create: `1.0`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `ADBD_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"fast_start_fail_over_lag_limit_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_auto_scaling_enabled":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"okv_end_point_group_name":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.name}`},
		"okv_key_store_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"shard_space_name":                          acctest.Representation{RepType: acctest.Optional, Create: `shardSpaceName`},
		"standby_maintenance_buffer_in_days":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
	}

	DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Required, acctest.Create, DatabaseKeyStoreRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtime_versions", "test_functions_runtime_versions", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtimes", "test_functions_runtimes", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Required, acctest.Create, FunctionsFunctionsPbfListingDataSourceRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, IdentityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
		KeyResourceDependencyConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedAutonomousDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedAutonomousDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database"
	datasourceName := "data.oci_distributed_database_distributed_autonomous_databases.test_distributed_autonomous_databases"
	singularDatasourceName := "data.oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation), "distributeddatabase", "distributedAutonomousDatabase", t)

	acctest.ResourceTest(t, testAccCheckDistributedDatabaseDistributedAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_workload_type", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "gdn-name"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(resourceName, "db_workload_type", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "gdn-name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "notification_topic_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_replication_unit_count", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "vcn_nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_nsg_ids.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DistributedDatabaseDistributedAutonomousDatabaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(resourceName, "db_workload_type", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "gdn-name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "notification_topic_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_replication_unit_count", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "vcn_nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_nsg_ids.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedAutonomousDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(resourceName, "db_workload_type", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "notification_topic_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_replication_unit_count", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.okv_end_point_group_name"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.okv_key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.shard_space_name", "shardSpaceName"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "vcn_nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_nsg_ids.0.vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_databases", "test_distributed_autonomous_databases", acctest.Optional, acctest.Update, DistributedDatabaseDistributedAutonomousDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedAutonomousDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "db_deployment_type", "dbDeploymentType"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_auto_resource_management_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "distributed_autonomous_database_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "distributed_autonomous_database_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auto_resource_management_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload_type", "OLTP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gds_control_node_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gsm_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "latest_gsm_image.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listener_port_tls", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_topic_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_method"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_replication_unit_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.compute_count", "1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.data_storage_size_in_gbs", "1.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "ADBD_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.standby_maintenance_buffer_in_days", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vcn_nsg_ids.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + DistributedDatabaseDistributedAutonomousDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"patch_operations",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDistributedDatabaseDistributedAutonomousDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DistributedAutonomousDbServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_distributed_database_distributed_autonomous_database" {
			noResourceFound = false
			request := oci_distributed_database.GetDistributedAutonomousDatabaseRequest{}

			tmp := rs.Primary.ID
			request.DistributedAutonomousDatabaseId = &tmp

			if value, ok := rs.Primary.Attributes["metadata"]; ok {
				request.Metadata = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "distributed_database")

			response, err := client.GetDistributedAutonomousDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DistributedDatabaseDistributedAutonomousDatabase") {
		resource.AddTestSweepers("DistributedDatabaseDistributedAutonomousDatabase", &resource.Sweeper{
			Name:         "DistributedDatabaseDistributedAutonomousDatabase",
			Dependencies: acctest.DependencyGraph["distributedAutonomousDatabase"],
			F:            sweepDistributedDatabaseDistributedAutonomousDatabaseResource,
		})
	}
}

func sweepDistributedDatabaseDistributedAutonomousDatabaseResource(compartment string) error {
	distributedAutonomousDbServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DistributedAutonomousDbServiceClient()
	distributedAutonomousDatabaseIds, err := getDistributedDatabaseDistributedAutonomousDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, distributedAutonomousDatabaseId := range distributedAutonomousDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[distributedAutonomousDatabaseId]; !ok {
			deleteDistributedAutonomousDatabaseRequest := oci_distributed_database.DeleteDistributedAutonomousDatabaseRequest{}

			deleteDistributedAutonomousDatabaseRequest.DistributedAutonomousDatabaseId = &distributedAutonomousDatabaseId

			deleteDistributedAutonomousDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "distributed_database")
			_, error := distributedAutonomousDbServiceClient.DeleteDistributedAutonomousDatabase(context.Background(), deleteDistributedAutonomousDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting DistributedAutonomousDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", distributedAutonomousDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &distributedAutonomousDatabaseId, DistributedDatabaseDistributedAutonomousDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DistributedDatabaseDistributedAutonomousDatabaseSweepResponseFetchOperation, "distributed_database", true)
		}
	}
	return nil
}

func getDistributedDatabaseDistributedAutonomousDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DistributedAutonomousDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	distributedAutonomousDbServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DistributedAutonomousDbServiceClient()

	listDistributedAutonomousDatabasesRequest := oci_distributed_database.ListDistributedAutonomousDatabasesRequest{}
	listDistributedAutonomousDatabasesRequest.CompartmentId = &compartmentId
	listDistributedAutonomousDatabasesRequest.LifecycleState = oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateActive
	listDistributedAutonomousDatabasesResponse, err := distributedAutonomousDbServiceClient.ListDistributedAutonomousDatabases(context.Background(), listDistributedAutonomousDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DistributedAutonomousDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, distributedAutonomousDatabase := range listDistributedAutonomousDatabasesResponse.Items {
		id := *distributedAutonomousDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DistributedAutonomousDatabaseId", id)
	}
	return resourceIds, nil
}

func DistributedDatabaseDistributedAutonomousDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if distributedAutonomousDatabaseResponse, ok := response.Response.(oci_distributed_database.GetDistributedAutonomousDatabaseResponse); ok {
		return distributedAutonomousDatabaseResponse.LifecycleState != oci_distributed_database.DistributedAutonomousDatabaseLifecycleStateDeleted
	}
	return false
}

func DistributedDatabaseDistributedAutonomousDatabaseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DistributedAutonomousDbServiceClient().GetDistributedAutonomousDatabase(context.Background(), oci_distributed_database.GetDistributedAutonomousDatabaseRequest{
		DistributedAutonomousDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// Missing typed readback must not discard the ECPU count already returned in metadata.
func TestAutonomousDatabaseComputeReadback(t *testing.T) {
	for _, tc := range []struct {
		name     string
		typed    *float32
		metadata *oci_distributed_database.DistributedAutonomousDbMetadata
		want     float32
		present  bool
	}{
		{name: "typed value wins", typed: common.Float32(8), metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "32.0"}}, want: 8, present: true},
		{name: "explicit zero is not missing", typed: common.Float32(0), metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "32.0"}}, want: 0, present: true},
		{name: "catalog ECPU metadata", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "32.0"}}, want: 32, present: true},
		{name: "DG ECPU metadata", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "8.0"}}, want: 8, present: true},
		{name: "fractional ECPU", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "0.2"}}, want: 0.2, present: true},
		{name: "whitespace", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": " ecpu ", "cpu": " 8.0 "}}, want: 8, present: true},
		{name: "missing metadata"},
		{name: "missing map", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{}},
		{name: "missing model", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"cpu": "8.0"}}},
		{name: "different units", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "OCPU", "cpu": "8.0"}}},
		{name: "missing cpu", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU"}}},
		{name: "invalid cpu", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "unknown"}}},
		{name: "zero cpu", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "0"}}},
		{name: "negative cpu", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "-2"}}},
		{name: "NaN cpu", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "NaN"}}},
		{name: "infinite cpu", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "+Inf"}}},
		{name: "overflow cpu", metadata: &oci_distributed_database.DistributedAutonomousDbMetadata{PropertiesMap: map[string]string{"computeModel": "ECPU", "cpu": "1e100"}}},
	} {
		for _, catalog := range []bool{true, false} {
			t.Run(fmt.Sprintf("%s/catalog=%t", tc.name, catalog), func(t *testing.T) {
				var got map[string]interface{}
				if catalog {
					var node oci_distributed_database.DistributedAutonomousDatabaseCatalogDatabase = oci_distributed_database.DistributedAutonomousDatabaseCatalogWithDedicatedInfra{ComputeCount: tc.typed, Metadata: tc.metadata}
					got = ddb.DistributedAutonomousDatabaseCatalogDatabaseToMap(&node)
				} else {
					got = ddb.DistributedAutonomousDatabaseShardDatabaseToMap(oci_distributed_database.DistributedAutonomousDatabaseShardWithDedicatedInfra{ComputeCount: tc.typed, Metadata: tc.metadata})
				}
				value, present := got["compute_count"]
				if present != tc.present || (present && value != tc.want) {
					t.Fatalf("compute_count=%v present=%t, want %v present=%t", value, present, tc.want, tc.present)
				}
			})
		}
	}
}

func autonomousComputeReadbackFixture(raft, typed bool) map[string]interface{} {
	node := func(name string, cpu float32) map[string]interface{} {
		var count interface{}
		if typed {
			count = cpu
		}
		return map[string]interface{}{"source": "ADBD_EXISTING_CLUSTER", "name": name, "status": "CONFIGURED", "cloudAutonomousVmClusterId": "cluster", "computeCount": count, "dataStorageSizeInGbs": 256, "metadata": map[string]interface{}{"map": map[string]string{"computeModel": "ECPU", "cpu": fmt.Sprintf("%.1f", cpu), "supportingResourceLifecycleState": "AVAILABLE"}}}
	}
	cpu := float32(8)
	if raft {
		cpu = 32
	}
	body := map[string]interface{}{"id": "ocid1.gdd.test", "compartmentId": "compartment", "displayName": "compute-test", "lifecycleState": "ACTIVE", "dbDeploymentType": "ADB_D", "shardingMethod": "SYSTEM", "listenerPort": 1521, "onsPortLocal": 6200, "onsPortRemote": 6230, "catalogDetails": []interface{}{map[string]interface{}{"originalReplica": map[string]interface{}{"name": "primary", "database": node("catalog", cpu)}, "dataGuardReplicas": []interface{}{map[string]interface{}{"name": "standby", "database": node("catalog-standby", cpu)}}}}}
	if raft {
		body["replicationMethod"] = "RAFT"
		body["systemRaftClusters"] = []interface{}{map[string]interface{}{"name": "rw", "ruMode": "READ_WRITE", "replicationFactor": 3, "witnessCount": 0, "databases": []interface{}{node("shard", 2)}}}
	} else {
		body["replicationMethod"] = "DG"
		body["systemDataGuardDatabases"] = map[string]interface{}{"originalReplica": map[string]interface{}{"name": "primary", "databases": []interface{}{node("shard", 8)}}, "dataGuardReplicas": []interface{}{map[string]interface{}{"name": "standby", "databases": []interface{}{node("shard-standby", 8)}}}}
	}
	return body
}

// Replay GET through the SDK and public Read handler, then plan with known sizing.
// Without the fallback, imports incorrectly replace shards whose typed count is null.
func TestAutonomousDatabaseComputeImportRefresh(t *testing.T) {
	for _, raft := range []bool{true, false} {
		for _, imported := range []bool{true, false} {
			t.Run(fmt.Sprintf("raft=%t/imported=%t", raft, imported), func(t *testing.T) {
				r := ddb.DistributedDatabaseDistributedAutonomousDatabaseResource()
				clientsFor := func(typed bool) *tf_client.OracleClients {
					return operationClients(func(req *http.Request) (*http.Response, error) {
						if req.Method != http.MethodGet {
							t.Fatalf("read/plan attempted mutation: %s", req.Method)
						}
						return operationResponse(req, 200, autonomousComputeReadbackFixture(raft, typed), ""), nil
					})
				}
				baseline := schema.TestResourceDataRaw(t, r.Schema, nil)
				baseline.SetId("ocid1.gdd.test")
				if diagnostics := r.ReadContext(context.Background(), baseline, clientsFor(true)); diagnostics.HasError() {
					t.Fatal(diagnostics)
				}
				config := map[string]interface{}{}
				for name, field := range r.Schema {
					topology := name == "catalog_details" || name == "system_raft_clusters" || name == "system_data_guard_databases"
					if field.Required || (topology && len(actionList(baseline.Get(name))) > 0) {
						config[name] = requestConfigValue(baseline.Get(name), field)
					}
				}
				data := schema.TestResourceDataRaw(t, r.Schema, config)
				if imported {
					data = schema.TestResourceDataRaw(t, r.Schema, nil)
				}
				data.SetId("ocid1.gdd.test")
				for refresh := 0; refresh < 2; refresh++ {
					if diagnostics := r.ReadContext(context.Background(), data, clientsFor(false)); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					diff, err := r.Diff(context.Background(), data.State(), sdkterraform.NewResourceConfigRaw(config), nil)
					if err != nil {
						t.Fatal(err)
					}
					if diff != nil && diff.RequiresNew() {
						for key, attribute := range diff.Attributes {
							if attribute.RequiresNew {
								t.Logf("replacement field: %s", key)
							}
						}
						t.Fatal("unchanged sizing plans replacement after null compute readback")
					}
				}
				catalogCompute := float64(8)
				if raft {
					catalogCompute = 32
				}
				key := "catalog_details.0.original_replica.0.database.0.compute_count"
				if data.Get(key) != catalogCompute {
					t.Fatalf("catalog compute=%v, want %v", data.Get(key), catalogCompute)
				}
				observed := "observed_catalog_details.0.original_replica.0.database.0.compute_count"
				if data.Get(observed) != catalogCompute {
					t.Fatalf("observed catalog compute=%v, want %v", data.Get(observed), catalogCompute)
				}
				nodeKey := "system_data_guard_databases.0.original_replica.0.databases.0.compute_count"
				nodeCompute := float64(8)
				if raft {
					nodeKey = "system_raft_clusters.0.databases.0.compute_count"
					nodeCompute = 2
				}
				if data.Get(nodeKey) != nodeCompute {
					t.Fatalf("shard compute=%v, want %v", data.Get(nodeKey), nodeCompute)
				}
				for _, attribute := range []string{"compute_count", "data_storage_size_in_gbs"} {
					// Build independent edited config so each check changes only one input.
					edited := map[string]interface{}{}
					for name, field := range r.Schema {
						if _, ok := config[name]; ok {
							edited[name] = requestConfigValue(baseline.Get(name), field)
						}
					}
					database := actionFirst(actionFirst(actionFirst(edited["catalog_details"])["original_replica"])["database"])
					database[attribute] = float64(512)
					diff, err := r.Diff(context.Background(), data.State(), sdkterraform.NewResourceConfigRaw(edited), nil)
					if err != nil {
						t.Fatal(err)
					}
					changedKey := "catalog_details.0.original_replica.0.database.0." + attribute
					if diff == nil || diff.Attributes[changedKey] == nil || !diff.Attributes[changedKey].RequiresNew {
						t.Fatalf("intentional %s edit was hidden", attribute)
					}
				}
			})
		}
	}
}
