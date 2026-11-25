// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkterraform "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"
	ddb "github.com/oracle/terraform-provider-oci/internal/service/distributed_database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DistributedDatabaseDistributedDatabaseRequiredOnlyResource = DistributedDatabaseDistributedDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation)

	DistributedDatabaseDistributedDatabaseResourceConfig = DistributedDatabaseDistributedDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabaseRepresentation)

	DistributedDatabaseDistributedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"distributed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_database.test_distributed_database.id}`},
		"metadata":                acctest.Representation{RepType: acctest.Optional, Create: `metadata`},
	}
	DistributedDatabaseDistributedDatabaseMetadataSingularDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"db_deployment_type":                  acctest.Representation{RepType: acctest.Optional, Create: `dbDeploymentType`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `gdn-name`, Update: `displayName2`},
		"is_auto_resource_management_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: `metadata`},
		"private_endpoint_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                              acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseDataSourceFilterRepresentation}}
	DistributedDatabaseDistributedDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_distributed_database_distributed_database.test_distributed_database.id}`}},
	}
	DistributedDatabaseDistributedDatabaseMetadataDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedDatabaseRepresentation = map[string]interface{}{
		"catalog_details":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsRepresentation},
		"character_set":                              acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"compartment_id":                             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_version":                           acctest.Representation{RepType: acctest.Required, Create: `26ai`},
		"display_name":                               acctest.Representation{RepType: acctest.Required, Create: `gdn-name`, Update: `displayName2`},
		"listener_port":                              acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ncharacter_set":                             acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"ons_port_local":                             acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ons_port_remote":                            acctest.Representation{RepType: acctest.Required, Create: `10`},
		"prefix":                                     acctest.Representation{RepType: acctest.Required, Create: `aaa`},
		"private_endpoint_ids":                       acctest.Representation{RepType: acctest.Required, Create: []string{`privateEndpointIds`}},
		"composite_data_guard_shard_spaces":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesRepresentation},
		"composite_raft_shard_spaces":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRepresentation},
		"db_backup_config":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseDbBackupConfigRepresentation},
		"defined_tags":                               acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"gsm_ssh_public_key":                         acctest.Representation{RepType: acctest.Optional, Create: `gsmSshPublicKey`},
		"listener_port_tls":                          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"notification_topic_ids":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`notificationTopicIds`}, Update: []string{`notificationTopicIds2`}},
		"patch_operations":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabasePatchOperationsRepresentation},
		"scan_listener_port":                         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"system_chunk_count":                         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"system_data_guard_databases":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesRepresentation},
		"system_raft_clusters":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemRaftClustersRepresentation},
		"system_raft_replication_unit_count":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"user_shard_spaces":                          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesRepresentation},
		"vcn_nsg_ids":                                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseVcnNsgIdsRepresentation},
		"add_gds_control_node_trigger":               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"change_db_backup_config_trigger":            acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"configure_auto_resource_management_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"configure_sharding_trigger":                 acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"generate_wallet_trigger":                    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"move_replication_unit_trigger":              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"recreate_failed_resource_trigger":           acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"scale_gsms_trigger":                         acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"start_database_trigger":                     acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"stop_database_trigger":                      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"validate_network_trigger":                   acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsRepresentation = map[string]interface{}{
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesRepresentation = map[string]interface{}{
		"chunk_count":         acctest.Representation{RepType: acctest.Required, Create: `10`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `west`},
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRepresentation = map[string]interface{}{
		"chunk_count":            acctest.Representation{RepType: acctest.Required, Create: `10`},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `west`},
		"raft_clusters":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersRepresentation},
		"replication_unit_count": acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_window":            acctest.Representation{RepType: acctest.Optional, Create: `SLOT_ONE`},
		"auto_full_backup_day":          acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`},
		"auto_full_backup_window":       acctest.Representation{RepType: acctest.Optional, Create: `SLOT_ONE`},
		"backup_deletion_policy":        acctest.Representation{RepType: acctest.Optional, Create: `DELETE_IMMEDIATELY`},
		"backup_destination_details":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseDbBackupConfigBackupDestinationDetailsRepresentation},
		"can_run_immediate_full_backup": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_auto_backup_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_remote_backup_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"recovery_window_in_days":       acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"remote_region":                 acctest.Representation{RepType: acctest.Optional, Create: `remoteRegion`},
	}
	DistributedDatabaseDistributedDatabasePatchOperationsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"value":     acctest.Representation{RepType: acctest.Optional, Create: `{}`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesRepresentation = map[string]interface{}{
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedDatabaseSystemRaftClustersRepresentation = map[string]interface{}{
		"databases":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseSystemRaftClustersDatabasesRepresentation},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `west`},
		"replication_factor": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ru_mode":            acctest.Representation{RepType: acctest.Required, Create: `READ_WRITE`},
		"witness_count":      acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesRepresentation = map[string]interface{}{
		"name":                acctest.Representation{RepType: acctest.Required, Create: `west`},
		"original_replica":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaRepresentation},
		"data_guard_replicas": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasRepresentation},
	}
	DistributedDatabaseDistributedDatabaseVcnNsgIdsRepresentation = map[string]interface{}{
		"nsg_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`nsgIds`}},
		"vcn_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersRepresentation = map[string]interface{}{
		"databases":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersDatabasesRepresentation},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `west`},
		"replication_factor": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ru_mode":            acctest.Representation{RepType: acctest.Required, Create: `READ_WRITE`},
		"witness_count":      acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"dbrs_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_policy.test_policy.id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `id`},
		"internet_proxy": acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"is_remote":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"remote_region":  acctest.Representation{RepType: acctest.Optional, Create: `remoteRegion`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasRepresentation = map[string]interface{}{
		"databases": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesRepresentation},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseSystemRaftClustersDatabasesRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemRaftClustersDatabasesDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemRaftClustersDatabasesVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasRepresentation = map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasDatabaseRepresentation},
		"name":     acctest.Representation{RepType: acctest.Required, Create: `west`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaDatabaseRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaDatabaseDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaDatabaseVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasDatabaseRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasDatabaseDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasDatabaseVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersDatabasesRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersDatabasesDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersDatabasesVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseSystemRaftClustersDatabasesDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseSystemRaftClustersDatabasesVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaDatabaseRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaDatabaseDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaDatabaseVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasDatabaseRepresentation = map[string]interface{}{
		"db_home_id":                                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_homes.t.db_homes.0.db_home_id}`},
		"source":                                    acctest.Representation{RepType: acctest.Required, Create: `XS_EXISTING_CLUSTER`},
		"admin_password":                            acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"admin_password_secret_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"admin_password_secret_version_number":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"availability_domain":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"db_storage_vault_details":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasDatabaseDbStorageVaultDetailsRepresentation},
		"exascale_db_storage_vault_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id}`},
		"kms_key_id":                                acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"protection_mode":                           acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"tde_wallet_password":                       acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},
		"tde_wallet_password_secret_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"tde_wallet_password_secret_version_number": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"transport_type":                            acctest.Representation{RepType: acctest.Optional, Create: `SYNC`},
		"vault_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		"vm_cluster_details":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasDatabaseVmClusterDetailsRepresentation},
		"vm_cluster_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaDatabaseDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsOriginalReplicaDatabaseVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasDatabaseDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsDataGuardReplicasDatabaseVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesOriginalReplicaDatabasesVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseCompositeDataGuardShardSpacesDataGuardReplicasDatabasesVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersDatabasesDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseCompositeRaftShardSpacesRaftClustersDatabasesVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesOriginalReplicaDatabasesVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseSystemDataGuardDatabasesDataGuardReplicasDatabasesVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaDatabaseDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesOriginalReplicaDatabaseVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasDatabaseDbStorageVaultDetailsRepresentation = map[string]interface{}{
		"additional_flash_cache_in_percent": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"autoscale_limit_in_gbs":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"compartment_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"high_capacity_database_storage":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_autoscale_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subscription_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
	}
	DistributedDatabaseDistributedDatabaseUserShardSpacesDataGuardReplicasDatabaseVmClusterDetailsRepresentation = map[string]interface{}{
		"backup_network_nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`backupNetworkNsgIds`}},
		"backup_subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"domain":                        acctest.Representation{RepType: acctest.Optional, Create: `domain`},
		"enabled_ecpu_count":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"is_diagnostics_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_health_monitoring_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_incident_logs_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":                 acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"nsg_ids":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}},
		"private_zone_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_zone.id}`},
		"ssh_public_keys":               acctest.Representation{RepType: acctest.Optional, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"subnet_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"subscription_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_onesubscription_subscription.test_subscription.id}`},
		"total_ecpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"vm_file_system_storage_size":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	DistributedDatabaseDistributedDatabaseResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, DatabaseDbHomeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Required, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_zone", "test_zone", acctest.Required, acctest.Create, DnsDnsZoneRepresentationPrimary) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtime_versions", "test_functions_runtime_versions", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtimes", "test_functions_runtimes", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Required, acctest.Create, FunctionsFunctionsPbfListingDataSourceRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
		KeyResourceDependencyConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_onesubscription_subscriptions", "test_subscriptions", acctest.Required, acctest.Create, OnesubscriptionOnesubscriptionSubscriptionDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_distributed_database_distributed_database.test_distributed_database"
	datasourceName := "data.oci_distributed_database_distributed_databases.test_distributed_databases"
	singularDatasourceName := "data.oci_distributed_database_distributed_database.test_distributed_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DistributedDatabaseDistributedDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation), "distributeddatabase", "distributedDatabase", t)

	acctest.ResourceTest(t, testAccCheckDistributedDatabaseDistributedDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
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
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_auto_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttrSet(resourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "gdn-name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gsm_ssh_public_key", "gsmSshPublicKey"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "notification_topic_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DistributedDatabaseDistributedDatabaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_auto_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttrSet(resourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "gdn-name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gsm_ssh_public_key", "gsmSshPublicKey"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "notification_topic_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_id"),
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
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_auto_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttrSet(resourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "gsm_ssh_public_key", "gsmSshPublicKey"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "notification_topic_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "aaa"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "system_raft_clusters.0.databases.0.vm_cluster_id"),
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
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.admin_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.exascale_db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password", "tdeWalletPassword"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.tde_wallet_password_secret_version_number", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.backup_subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.private_zone_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_databases", "test_distributed_databases", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "db_deployment_type", "dbDeploymentType"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_auto_resource_management_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "distributed_database_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "distributed_database_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auto_resource_management_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_data_guard_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.chunk_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.replication_factor", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.ru_mode", "READ_WRITE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.raft_clusters.0.witness_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "composite_raft_shard_spaces.0.replication_unit_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.is_auto_backup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_deployment_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gds_control_node_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gsm_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gsm_ssh_public_key", "gsmSshPublicKey"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "latest_gsm_image_details.#", "1"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_port", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sharding_method"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_chunk_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_data_guard_databases.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_raft_clusters.0.databases.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
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
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.data_guard_replicas.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.additional_flash_cache_in_percent", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.autoscale_limit_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.db_storage_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.high_capacity_database_storage", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.db_storage_vault_details.0.is_autoscale_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.source", "XS_EXISTING_CLUSTER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.transport_type", "SYNC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.domain", "domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.enabled_ecpu_count", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_diagnostics_events_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_health_monitoring_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.is_incident_logs_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.total_ecpu_count", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.vm_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.database.0.vm_cluster_details.0.vm_file_system_storage_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_shard_spaces.0.original_replica.0.name", "west"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vcn_nsg_ids.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + DistributedDatabaseDistributedDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"patch_operations",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDistributedDatabaseDistributedDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DistributedDbServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_distributed_database_distributed_database" {
			noResourceFound = false
			request := oci_distributed_database.GetDistributedDatabaseRequest{}

			tmp := rs.Primary.ID
			request.DistributedDatabaseId = &tmp

			if value, ok := rs.Primary.Attributes["metadata"]; ok {
				request.Metadata = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "distributed_database")

			response, err := client.GetDistributedDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_distributed_database.DistributedDatabaseLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DistributedDatabaseDistributedDatabase") {
		resource.AddTestSweepers("DistributedDatabaseDistributedDatabase", &resource.Sweeper{
			Name:         "DistributedDatabaseDistributedDatabase",
			Dependencies: acctest.DependencyGraph["distributedDatabase"],
			F:            sweepDistributedDatabaseDistributedDatabaseResource,
		})
	}
}

func sweepDistributedDatabaseDistributedDatabaseResource(compartment string) error {
	distributedDbServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DistributedDbServiceClient()
	distributedDatabaseIds, err := getDistributedDatabaseDistributedDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, distributedDatabaseId := range distributedDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[distributedDatabaseId]; !ok {
			deleteDistributedDatabaseRequest := oci_distributed_database.DeleteDistributedDatabaseRequest{}

			deleteDistributedDatabaseRequest.DistributedDatabaseId = &distributedDatabaseId

			deleteDistributedDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "distributed_database")
			_, error := distributedDbServiceClient.DeleteDistributedDatabase(context.Background(), deleteDistributedDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting DistributedDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", distributedDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &distributedDatabaseId, DistributedDatabaseDistributedDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DistributedDatabaseDistributedDatabaseSweepResponseFetchOperation, "distributed_database", true)
		}
	}
	return nil
}

func getDistributedDatabaseDistributedDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DistributedDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	distributedDbServiceClient := acctest.GetTestClients(&schema.ResourceData{}).DistributedDbServiceClient()

	listDistributedDatabasesRequest := oci_distributed_database.ListDistributedDatabasesRequest{}
	listDistributedDatabasesRequest.CompartmentId = &compartmentId
	listDistributedDatabasesRequest.LifecycleState = oci_distributed_database.DistributedDatabaseLifecycleStateActive
	listDistributedDatabasesResponse, err := distributedDbServiceClient.ListDistributedDatabases(context.Background(), listDistributedDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DistributedDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, distributedDatabase := range listDistributedDatabasesResponse.Items {
		id := *distributedDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DistributedDatabaseId", id)
	}
	return resourceIds, nil
}

func DistributedDatabaseDistributedDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if distributedDatabaseResponse, ok := response.Response.(oci_distributed_database.GetDistributedDatabaseResponse); ok {
		return distributedDatabaseResponse.LifecycleState != oci_distributed_database.DistributedDatabaseLifecycleStateDeleted
	}
	return false
}

func DistributedDatabaseDistributedDatabaseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DistributedDbServiceClient().GetDistributedDatabase(context.Background(), oci_distributed_database.GetDistributedDatabaseRequest{
		DistributedDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func TestDistributedDatabaseDistributedDatabase_SourceSpecificDatabaseHome(t *testing.T) {
	resource := ddb.DistributedDatabaseDistributedDatabaseResource()
	planner := *resource
	planner.Schema = map[string]*schema.Schema{"system_raft_clusters": resource.Schema["system_raft_clusters"]}
	for _, scenario := range []struct {
		name, source, home string
		invalid            bool
	}{
		{"Exascale cluster without database home", "XS_EXISTING_CLUSTER", "", false},
		{"Dedicated cluster without database home", "XD_EXISTING_CLUSTER", "", false},
		{"Database home source without home", "EXISTING_DB_HOME", "", true},
		{"Database home source with home", "EXISTING_DB_HOME", "home", false},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			database := map[string]interface{}{"source": scenario.source, "vm_cluster_id": "cluster"}
			if scenario.home != "" {
				database["db_home_id"] = scenario.home
			}
			config := map[string]interface{}{"system_raft_clusters": []interface{}{map[string]interface{}{"name": "rw", "replication_factor": 3, "witness_count": 0, "ru_mode": "READ_WRITE", "databases": []interface{}{database}}}}
			diagnostics := planner.Validate(sdkterraform.NewResourceConfigRaw(config))
			rawDatabase := map[string]cty.Value{"source": cty.StringVal(scenario.source)}
			if scenario.home != "" {
				rawDatabase["db_home_id"] = cty.StringVal(scenario.home)
			}
			// The SDK's direct Diff helper takes raw configuration from state;
			// supply it explicitly to exercise the protocol-level validation.
			raw := cty.ObjectVal(map[string]cty.Value{"system_raft_clusters": cty.TupleVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{"databases": cty.TupleVal([]cty.Value{cty.ObjectVal(rawDatabase)})}),
			})})
			_, err := planner.Diff(context.Background(), &sdkterraform.InstanceState{RawConfig: raw}, sdkterraform.NewResourceConfigRaw(config), nil)
			if scenario.invalid {
				if !diagnostics.HasError() && err == nil {
					t.Fatal("EXISTING_DB_HOME accepted without db_home_id")
				}
			} else if diagnostics.HasError() || err != nil {
				t.Fatalf("valid source rejected: %v %v", diagnostics, err)
			}
		})
	}
}

func TestDistributedDatabaseCredentialSchema(t *testing.T) {
	for _, auto := range []bool{false, true} {
		r := ddb.DistributedDatabaseDistributedDatabaseResource()
		if auto {
			r = ddb.DistributedDatabaseDistributedAutonomousDatabaseResource()
		}
		t.Run(fmt.Sprintf("autonomous=%t", auto), func(t *testing.T) {
			var check func(map[string]*schema.Schema, string)
			check = func(fields map[string]*schema.Schema, path string) {
				for name, field := range fields {
					if strings.HasPrefix(name, "observed_") {
						continue
					}
					key := path + name
					switch name {
					case "admin_password":
						if !field.Optional || field.Computed || !field.Sensitive || !field.ForceNew {
							t.Errorf("%s must be optional, sensitive, creation-only input", key)
						}
					case "admin_password_secret_id", "tde_wallet_password_secret_id":
						if field.Sensitive || !field.Optional || !field.Computed || !field.ForceNew {
							t.Errorf("%s must be a non-sensitive optional/computed creation-only reference", key)
						}
					}
					if nested, ok := field.Elem.(*schema.Resource); ok {
						check(nested.Schema, key+".")
					}
				}
			}
			check(r.Schema, "")
		})
	}
	// INSERTDG clones the database schema and must not re-mask the secret OCID.
	action := ddb.DistributedDatabaseDistributedDatabasesInsertDgResource()
	fields := action.Schema["insert_dg"].Elem.(*schema.Resource).Schema["database"].Elem.(*schema.Resource).Schema
	if fields["admin_password_secret_id"].Sensitive || !fields["admin_password"].Sensitive {
		t.Fatal("INSERTDG must distinguish secret references from plaintext credentials")
	}
}

func credentialReadbackFixture(auto bool) map[string]interface{} {
	source, placement, entity := "XS_EXISTING_CLUSTER", "vmClusterId", "EXADB"
	if auto {
		source, placement, entity = "ADBD_EXISTING_CLUSTER", "cloudAutonomousVmClusterId", "ADB_D"
	}
	node := func(name string) map[string]interface{} {
		result := map[string]interface{}{"source": source, placement: "cluster", "name": name, "status": "CREATED"}
		if auto {
			result["computeCount"] = 2
			result["dataStorageSizeInGbs"] = 32
		}
		return result
	}
	return map[string]interface{}{"id": "ocid1.gdd.test", "compartmentId": "compartment", "displayName": "credential-test", "databaseVersion": "23ai", "lifecycleState": "INACTIVE", "dbDeploymentType": entity, "shardingMethod": "SYSTEM", "replicationMethod": "RAFT", "prefix": "creds", "listenerPort": 1521, "onsPortLocal": 6200, "onsPortRemote": 6230, "privateEndpointIds": []string{"endpoint"}, "catalogDetails": []interface{}{map[string]interface{}{"originalReplica": map[string]interface{}{"name": "primary", "database": node("catalog")}, "dataGuardReplicas": []interface{}{map[string]interface{}{"name": "standby", "database": node("catalog-standby")}}}}, "systemRaftClusters": []interface{}{map[string]interface{}{"name": "rw", "replicationFactor": 3, "witnessCount": 0, "ruMode": "READ_WRITE", "databases": []interface{}{node("shard")}}}}
}

func credentialCreateConfig(auto, secret bool) map[string]interface{} {
	source, placement, deployment := "XS_EXISTING_CLUSTER", "vm_cluster_id", "EXADB"
	if auto {
		source, placement, deployment = "ADBD_EXISTING_CLUSTER", "cloud_autonomous_vm_cluster_id", "ADB_D"
	}
	node := func() map[string]interface{} {
		result := map[string]interface{}{"source": source, placement: "cluster"}
		if auto {
			result["compute_count"] = 2
			result["data_storage_size_in_gbs"] = 32
		}
		if secret {
			result["admin_password_secret_id"] = "ocid1.vaultsecret.test.admin"
		} else {
			result["admin_password"] = "TestOnlyPassword123_"
		}
		if !auto {
			if secret {
				result["tde_wallet_password_secret_id"] = "ocid1.vaultsecret.test.tde"
			} else {
				result["tde_wallet_password"] = "TestOnlyWallet123_"
			}
		}
		return result
	}
	return map[string]interface{}{"compartment_id": "compartment", "display_name": "credential-test", "database_version": "23ai", "db_deployment_type": deployment, "sharding_method": "SYSTEM", "replication_method": "RAFT", "prefix": "creds", "listener_port": 1521, "ons_port_local": 6200, "ons_port_remote": 6230, "private_endpoint_ids": []interface{}{"endpoint"}, "catalog_details": []interface{}{map[string]interface{}{"original_replica": []interface{}{map[string]interface{}{"name": "primary", "database": []interface{}{node()}}}, "data_guard_replicas": []interface{}{map[string]interface{}{"name": "standby", "database": []interface{}{node()}}}}}, "system_raft_clusters": []interface{}{map[string]interface{}{"name": "rw", "replication_factor": 3, "witness_count": 0, "ru_mode": "READ_WRITE", "databases": []interface{}{node()}}}}
}

func credentialConfigNodes(config map[string]interface{}) []map[string]interface{} {
	catalog := actionFirst(config["catalog_details"])
	return []map[string]interface{}{actionFirst(actionFirst(catalog["original_replica"])["database"]), actionFirst(actionFirst(catalog["data_guard_replicas"])["database"]), actionFirst(actionFirst(config["system_raft_clusters"])["databases"])}
}

func TestDistributedDatabaseCredentialLifecycle(t *testing.T) {
	paths := []string{"catalog_details.0.original_replica.0.database.0.", "catalog_details.0.data_guard_replicas.0.database.0.", "system_raft_clusters.0.databases.0."}
	for _, auto := range []bool{false, true} {
		for _, secret := range []bool{false, true} {
			t.Run(fmt.Sprintf("autonomous=%t/secret=%t", auto, secret), func(t *testing.T) {
				r := ddb.DistributedDatabaseDistributedDatabaseResource()
				entity := "distributeddatabase"
				if auto {
					r = ddb.DistributedDatabaseDistributedAutonomousDatabaseResource()
					entity = "distributedautonomousdatabase"
				}
				config := credentialCreateConfig(auto, secret)
				data := schema.TestResourceDataRaw(t, r.Schema, config)
				var requestBody map[string]interface{}
				creates := 0
				clients := operationClients(func(req *http.Request) (*http.Response, error) {
					switch {
					case req.Method == http.MethodPost:
						creates++
						if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
							t.Fatal(err)
						}
						return operationResponse(req, 202, credentialReadbackFixture(auto), "wr-credentials"), nil
					case req.Method == http.MethodGet && strings.Contains(req.URL.Path, "workRequests/"):
						return operationResponse(req, 200, map[string]interface{}{"status": "SUCCEEDED", "resources": []interface{}{map[string]interface{}{"entityType": entity, "actionType": "CREATED", "identifier": "ocid1.gdd.test"}}}, ""), nil
					case req.Method == http.MethodGet:
						return operationResponse(req, 200, credentialReadbackFixture(auto), ""), nil
					default:
						t.Fatalf("unexpected mutation: %s", req.Method)
						return nil, fmt.Errorf("unexpected mutation")
					}
				})
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if diagnostics := r.CreateContext(ctx, data, clients); diagnostics.HasError() {
					t.Fatal(diagnostics)
				}
				if creates != 1 || data.Id() != "ocid1.gdd.test" {
					t.Fatal("create did not complete exactly once")
				}
				catalog := requestBody["catalogDetails"].([]interface{})[0].(map[string]interface{})
				wireNodes := []map[string]interface{}{catalog["originalReplica"].(map[string]interface{})["database"].(map[string]interface{}), catalog["dataGuardReplicas"].([]interface{})[0].(map[string]interface{})["database"].(map[string]interface{}), requestBody["systemRaftClusters"].([]interface{})[0].(map[string]interface{})["databases"].([]interface{})[0].(map[string]interface{})}
				input, wire, other, want := "admin_password", "adminPassword", "adminPasswordSecretId", "TestOnlyPassword123_"
				if secret {
					input, wire, other, want = "admin_password_secret_id", "adminPasswordSecretId", "adminPassword", "ocid1.vaultsecret.test.admin"
				}
				for _, node := range wireNodes {
					if node[wire] != want {
						t.Fatal("create request lost selected admin credential")
					}
					if node[other] != nil {
						t.Fatal("create request populated the unused admin credential alternative")
					}
					if !auto && secret && node["tdeWalletPasswordSecretId"] != "ocid1.vaultsecret.test.tde" {
						t.Fatal("create request lost TDE secret reference")
					}
				}
				for refresh := 0; refresh < 2; refresh++ {
					if diagnostics := r.ReadContext(ctx, data, clients); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					for _, path := range paths {
						if data.Get(path+input) != want {
							t.Fatal("refresh lost configured admin credential")
						}
						if !auto && secret && data.Get(path+"tde_wallet_password_secret_id") != "ocid1.vaultsecret.test.tde" {
							t.Fatal("refresh lost configured TDE reference")
						}
					}
					diff, err := r.Diff(ctx, data.State(), sdkterraform.NewResourceConfigRaw(config), nil)
					if err != nil {
						t.Fatal(err)
					}
					if diff != nil && diff.RequiresNew() {
						for key, attribute := range diff.Attributes {
							if attribute.RequiresNew {
								t.Logf("replacement field: %s", key)
							}
						}
						t.Fatal("unchanged credentials plan replacement")
					}
				}
				for nodeIndex, path := range paths {
					fields := []string{input}
					if !auto && secret {
						fields = append(fields, "tde_wallet_password_secret_id")
					}
					for _, field := range fields {
						raw, err := json.Marshal(config)
						if err != nil {
							t.Fatal(err)
						}
						var edited map[string]interface{}
						if err = json.Unmarshal(raw, &edited); err != nil {
							t.Fatal(err)
						}
						credentialConfigNodes(edited)[nodeIndex][field] = "ChangedTestCredential123_"
						diff, err := r.Diff(ctx, data.State(), sdkterraform.NewResourceConfigRaw(edited), nil)
						if err != nil {
							t.Fatal(err)
						}
						if diff == nil || diff.Attributes[path+field] == nil || !diff.Attributes[path+field].RequiresNew {
							t.Fatalf("intentional %s edit did not require replacement", field)
						}
					}
				}
				if creates != 1 {
					t.Fatal("refresh or plan repeated resource creation")
				}
			})
		}
	}
}

func TestDistributedDatabaseUnknownPasswordDoesNotHideChange(t *testing.T) {
	for _, auto := range []bool{false, true} {
		t.Run(fmt.Sprintf("autonomous=%t", auto), func(t *testing.T) {
			r := ddb.DistributedDatabaseDistributedDatabaseResource()
			if auto {
				r = ddb.DistributedDatabaseDistributedAutonomousDatabaseResource()
			}
			config := credentialCreateConfig(auto, false)
			raw, err := json.Marshal(config)
			if err != nil {
				t.Fatal(err)
			}
			var missing map[string]interface{}
			if err = json.Unmarshal(raw, &missing); err != nil {
				t.Fatal(err)
			}
			for _, node := range credentialConfigNodes(missing) {
				delete(node, "admin_password")
			}
			data := schema.TestResourceDataRaw(t, r.Schema, missing)
			data.SetId("ocid1.gdd.test")
			clients := operationClients(func(req *http.Request) (*http.Response, error) {
				if req.Method != http.MethodGet {
					t.Fatalf("unexpected mutation: %s", req.Method)
				}
				return operationResponse(req, 200, credentialReadbackFixture(auto), ""), nil
			})
			if diagnostics := r.ReadContext(context.Background(), data, clients); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			diff, err := r.Diff(context.Background(), data.State(), sdkterraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			for _, path := range []string{"catalog_details.0.original_replica.0.database.0.", "catalog_details.0.data_guard_replicas.0.database.0.", "system_raft_clusters.0.databases.0."} {
				if diff == nil || diff.Attributes[path+"admin_password"] == nil || !diff.Attributes[path+"admin_password"].RequiresNew {
					t.Fatal("setting a password absent from state must remain an explicit change")
				}
			}
		})
	}
}

// Shared helpers for the existing DDB readback tests; kept here after removing the standalone unit test files.
type operationSigner struct{}

func (operationSigner) Sign(*http.Request) error { return nil }

type operationTransport func(*http.Request) (*http.Response, error)

func (f operationTransport) Do(r *http.Request) (*http.Response, error) { return f(r) }
func operationClients(handler operationTransport) *tf_client.OracleClients {
	base := common.DefaultBaseClientWithSigner(operationSigner{})
	base.Host = "https://ddb-test.invalid"
	base.BasePath = "20260101"
	base.HTTPClient = handler
	return &tf_client.OracleClients{SdkClientMap: map[string]interface{}{
		"oci_distributed_database.DistributedDbServiceClient":            &oci_distributed_database.DistributedDbServiceClient{BaseClient: base},
		"oci_distributed_database.DistributedAutonomousDbServiceClient":  &oci_distributed_database.DistributedAutonomousDbServiceClient{BaseClient: base},
		"oci_distributed_database.DistributedDbWorkRequestServiceClient": &oci_distributed_database.DistributedDbWorkRequestServiceClient{BaseClient: base},
	}}
}
func operationResponse(r *http.Request, status int, body interface{}, work string) *http.Response {
	raw, _ := json.Marshal(body)
	h := http.Header{"Content-Type": []string{"application/json"}}
	if work != "" {
		h.Set("opc-work-request-id", work)
	}
	return &http.Response{StatusCode: status, Header: h, Body: io.NopCloser(strings.NewReader(string(raw))), Request: r}
}

func requestConfigValue(value interface{}, field *schema.Schema) interface{} {
	nested, ok := field.Elem.(*schema.Resource)
	if !ok {
		return value
	}
	result := []interface{}{}
	for _, raw := range actionList(value) {
		item := actionMap(raw)
		filtered := map[string]interface{}{}
		for name, child := range nested.Schema {
			if child.Required || child.Optional {
				if v, ok := item[name]; ok {
					filtered[name] = requestConfigValue(v, child)
				}
			}
		}
		result = append(result, filtered)
	}
	return result
}

func actionList(value interface{}) []interface{} { items, _ := value.([]interface{}); return items }
func actionMap(value interface{}) map[string]interface{} {
	item, _ := value.(map[string]interface{})
	return item
}
func actionFirst(value interface{}) map[string]interface{} {
	items := actionList(value)
	if len(items) == 0 {
		return nil
	}
	return actionMap(items[0])
}
