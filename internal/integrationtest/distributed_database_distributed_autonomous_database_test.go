// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
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
		// metadata is computed-only for this datasource; do not configure it.
	}
	DistributedDatabaseDistributedAutonomousDatabaseMetadataSingularDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedAutonomousDatabaseDataSourceRepresentation = map[string]interface{}{
		/*
			"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"db_deployment_type":  acctest.Representation{RepType: acctest.Optional, Create: `ADB_D`},
			"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `my-dadb`, Update: `displayName2`},
			"metadata":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseMetadataRepresentation},
			"private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_distributed_database_distributed_database_private_endpoint.test_distributed_database_private_endpoint.id}`},
			"state":               acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
			"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseDataSourceFilterRepresentation}}
		*/
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_deployment_type": acctest.Representation{RepType: acctest.Optional, Create: `ADB_D`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `ss2_tf_dadb_it`, Update: `ss2_tf_dadb_it`},
		// metadata is computed-only for this datasource; do not configure it.
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseDataSourceFilterRepresentation}}
	DistributedDatabaseDistributedAutonomousDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`}},
	}
	DistributedDatabaseDistributedAutonomousDatabaseMetadataDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedAutonomousDatabaseRepresentation = map[string]interface{}{
		/*
				"catalog_details":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsRepresentation},
				"character_set":                      acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
				"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
				"database_version":                   acctest.Representation{RepType: acctest.Required, Create: `19c`},
				"db_deployment_type":                 acctest.Representation{RepType: acctest.Required, Create: `ADB_D`},
				"db_workload":                        acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
				"display_name":                       acctest.Representation{RepType: acctest.Required, Create: `my-dadb`, Update: `displayName2`},
				"distributed_autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`},
				"listener_port":                      acctest.Representation{RepType: acctest.Required, Create: `10`},
				"ncharacter_set":                     acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
				"ons_port_local":                     acctest.Representation{RepType: acctest.Required, Create: `10`},
				"ons_port_remote":                    acctest.Representation{RepType: acctest.Required, Create: `10`},
				"prefix":                             acctest.Representation{RepType: acctest.Required, Create: `pre`},
				"private_endpoint_ids":               acctest.Representation{RepType: acctest.Required, Create: []string{`privateEndpointIds`}},
				"shard_details":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseShardDetailsRepresentation},
				"sharding_method":                    acctest.Representation{RepType: acctest.Required, Create: `USER`},
				"chunks":                             acctest.Representation{RepType: acctest.Optional, Create: `10`},
				"db_backup_config":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigRepresentation},
				"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
				"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
				"listener_port_tls":                  acctest.Representation{RepType: acctest.Optional, Create: `10`},
				"patch_operations":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabasePatchOperationsRepresentation},
				"replication_factor":                 acctest.Representation{RepType: acctest.Optional, Create: `10`},
				"replication_method":                 acctest.Representation{RepType: acctest.Optional, Create: `RAFT`},
				"replication_unit":                   acctest.Representation{RepType: acctest.Optional, Create: `10`},
				"change_db_backup_config_trigger":    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"configure_gsm_wallet_trigger":       acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"configure_sharding_trigger":         acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"download_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"generate_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"generate_wallet_trigger":                               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"move_replication_unit_trigger":                         acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"recreate_failed_resource_trigger":                      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"upload_signed_certificate_and_generate_wallet_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"validate_ca_bundle_trigger":                            acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"validate_network_trigger":                              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
				"state":                                                 acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
			}
			DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsRepresentation = map[string]interface{}{
				"admin_password":                       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
				"cloud_autonomous_vm_cluster_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
				"compute_count":                        acctest.Representation{RepType: acctest.Required, Create: `1.0`},
				"data_storage_size_in_gbs":             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
				"is_auto_scaling_enabled":              acctest.Representation{RepType: acctest.Required, Create: `false`},
				"source":                               acctest.Representation{RepType: acctest.Required, Create: `ADB_D`},
				"kms_key_id":                           acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
				"kms_key_version_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
				"okv_end_point_group":                  acctest.Representation{RepType: acctest.Optional, Create: `okvEndPointGroup`},
				"okv_key_store_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
				"peer_cloud_autonomous_vm_cluster_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`peerCloudAutonomousVmClusterIds`}},
				"peer_details":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsPeerDetailsRepresentation},
				"vault_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
			}
			DistributedDatabaseDistributedAutonomousDatabaseShardDetailsRepresentation = map[string]interface{}{
				"admin_password":                       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
				"cloud_autonomous_vm_cluster_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
				"compute_count":                        acctest.Representation{RepType: acctest.Required, Create: `1.0`},
				"data_storage_size_in_gbs":             acctest.Representation{RepType: acctest.Required, Create: `1.0`},
				"is_auto_scaling_enabled":              acctest.Representation{RepType: acctest.Required, Create: `false`},
				"source":                               acctest.Representation{RepType: acctest.Required, Create: `ADB_D`},
				"kms_key_id":                           acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
				"kms_key_version_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
				"okv_end_point_group":                  acctest.Representation{RepType: acctest.Optional, Create: `okvEndPointGroup`},
				"okv_key_store_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
				"peer_cloud_autonomous_vm_cluster_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`peerCloudAutonomousVmClusterIds`}},
				"peer_details":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseShardDetailsPeerDetailsRepresentation},
				"shard_space":                          acctest.Representation{RepType: acctest.Optional, Create: `ss2`},
				"vault_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_vault.test_vault.id}`},
		*/
		"catalog_details":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsRepresentation},
		"character_set":      acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_version":   acctest.Representation{RepType: acctest.Required, Create: `26ai`},
		"db_deployment_type": acctest.Representation{RepType: acctest.Required, Create: `ADB_D`},
		"db_workload":        acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `ss2_tf_dadb_it`, Update: `ss2_tf_dadb_it`},
		// "distributed_autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`},
		"listener_port":        acctest.Representation{RepType: acctest.Required, Create: `21211`},
		"ncharacter_set":       acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"ons_port_local":       acctest.Representation{RepType: acctest.Required, Create: `21212`},
		"ons_port_remote":      acctest.Representation{RepType: acctest.Required, Create: `21213`},
		"prefix":               acctest.Representation{RepType: acctest.Required, Create: `ss2`},
		"private_endpoint_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.private_endpoint_id}`}},
		"shard_details": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseShardDetailsRepresentation},
			{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseShardDetailsRepresentation},
			{RepType: acctest.Required, Group: DistributedDatabaseDistributedAutonomousDatabaseShardDetailsRepresentation},
		},
		"sharding_method":   acctest.Representation{RepType: acctest.Required, Create: `SYSTEM`},
		"chunks":            acctest.Representation{RepType: acctest.Required, Create: `120`},
		"db_backup_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigRepresentation},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"listener_port_tls": acctest.Representation{RepType: acctest.Required, Create: `21214`},
		// "patch_operations":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabasePatchOperationsRepresentation},
		"replication_factor": acctest.Representation{RepType: acctest.Optional, Create: `3`},
		"replication_method": acctest.Representation{RepType: acctest.Optional, Create: `RAFT`},
		"replication_unit":   acctest.Representation{RepType: acctest.Optional, Create: `6`},
		// "change_db_backup_config_trigger":    acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "configure_sharding_trigger":         acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "download_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "generate_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "generate_wallet_trigger":                               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "upload_signed_certificate_and_generate_wallet_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "validate_network_trigger":                              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"state": acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `INACTIVE`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseCatalogDetailsRepresentation = map[string]interface{}{
		"admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_autonomous_vm_cluster_id}`},
		"compute_count":                  acctest.Representation{RepType: acctest.Required, Create: `4`},
		"data_storage_size_in_gbs":       acctest.Representation{RepType: acctest.Required, Create: `128`},
		"is_auto_scaling_enabled":        acctest.Representation{RepType: acctest.Required, Create: `false`},
		"source":                         acctest.Representation{RepType: acctest.Required, Create: `ADB_D`},
		"kms_key_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseShardDetailsRepresentation = map[string]interface{}{
		"admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_autonomous_vm_cluster_id}`},
		"compute_count":                  acctest.Representation{RepType: acctest.Required, Create: `4`},
		"data_storage_size_in_gbs":       acctest.Representation{RepType: acctest.Required, Create: `128`},
		"is_auto_scaling_enabled":        acctest.Representation{RepType: acctest.Required, Create: `false`},
		"source":                         acctest.Representation{RepType: acctest.Required, Create: `ADB_D`},
		"kms_key_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
	}
	DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DistributedDatabaseDistributedAutonomousDatabasePatchOperationsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		// WORKAROUND FOR GENERATED CODE ISSUE:
		// The integration test code generator emits untyped composite literals (`{}`)
		// inside map[string]interface{} values. This is invalid Go and causes `go vet`
		// to fail with: "missing type in composite literal".
		//
		// Go requires explicitly typed empty literals (e.g. map[string]interface{}{})
		// in this context. Until the generator is fixed, ensure all empty Create/Update
		// representations use concrete typed literals.
		//
		// See JIRA: TOP-9427
		//"value":     acctest.Representation{RepType: acctest.Optional, Create: {}},
		"value": acctest.Representation{RepType: acctest.Optional, Create: map[string]interface{}{}},
	}
	DistributedDatabaseDistributedAutonomousDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORE`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `id`},
		"internet_proxy": acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"is_remote":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"remote_region":  acctest.Representation{RepType: acctest.Optional, Create: `remoteRegion`},
		"vpc_password":   acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`},
		"vpc_user":       acctest.Representation{RepType: acctest.Optional, Create: `vpcUser`},
	}
	/*
		DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Required, acctest.Create, DatabaseKeyStoreRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database_private_endpoint", "test_distributed_database_private_endpoint", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabasePrivateEndpointRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
			acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Required, acctest.Create, FunctionsPbfListingDataSourceRepresentation) +
			AvailabilityDomainConfig +
			DefinedTagsDependencies +
			acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
			KeyResourceDependencyConfig +
			acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
	*/
	DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies = ""
	// acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
	// 	AvailabilityDomainConfig +
	// 	DefinedTagsDependencies +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
	// 	KeyResourceDependencyConfig +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation)
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedAutonomousDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedAutonomousDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	privateEndpointId := utils.GetEnvSettingWithBlankDefault("private_endpoint_id")
	privateEndpointIdVariableStr := fmt.Sprintf("variable \"private_endpoint_id\" { default = \"%s\" }\n", privateEndpointId)
	cloudAutonomousVmClusterId := utils.GetEnvSettingWithBlankDefault("cloud_autonomous_vm_cluster_id")
	cloudAutonomousVmClusterIdVariableStr := fmt.Sprintf("variable \"cloud_autonomous_vm_cluster_id\" { default = \"%s\" }\n", cloudAutonomousVmClusterId)
	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)
	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)
	displayName := "ss2_tf_dadb_it"
	listenerPort := "21211"
	onsPortLocal := "21212"
	onsPortRemote := "21213"
	listenerPortTls := "21214"
	prefix := "ss2"

	resourceName := "oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database"
	datasourceName := "data.oci_distributed_database_distributed_autonomous_databases.test_distributed_autonomous_databases"
	singularDatasourceName := "data.oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database"
	resourceConfig := config + compartmentIdVariableStr + privateEndpointIdVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceConfig
	pluralDataSourceConfig := acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_databases", "test_distributed_autonomous_databases", acctest.Optional, acctest.Update, DistributedDatabaseDistributedAutonomousDatabaseDataSourceRepresentation) +
		resourceConfig
	singularDataSourceConfig := acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseSingularDataSourceRepresentation) +
		resourceConfig
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(resourceConfig, "distributeddatabase", "distributedAutonomousDatabase", t)

	acctest.ResourceTest(t, testAccCheckDistributedDatabaseDistributedAutonomousDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: resourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", ""),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id", cloudAutonomousVmClusterId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "ADB_D"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "ADB_D"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "listener_port", listenerPort),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", onsPortLocal),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", onsPortRemote),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", listenerPortTls),
				resource.TestCheckResourceAttr(resourceName, "prefix", prefix),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", ""),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id", cloudAutonomousVmClusterId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "ADB_D"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.admin_password", ""),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.cloud_autonomous_vm_cluster_id", cloudAutonomousVmClusterId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.data_storage_size_in_gbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.source", "ADB_D"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.admin_password", ""),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.cloud_autonomous_vm_cluster_id", cloudAutonomousVmClusterId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.data_storage_size_in_gbs", "128"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.source", "ADB_D"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "sharding_method", "SYSTEM"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
			),
		},

		// delete before next Create
		/*
			{
				Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.okv_key_store_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "ADB_D"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
					resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORAGE"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "ADB_D"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "my-dadb"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
					resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
					resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
					resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
					resource.TestCheckResourceAttr(resourceName, "prefix", "pre"),
					resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
					resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
					resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.okv_key_store_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "ss2"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "ADB_D"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
					resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.okv_key_store_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "ADB_D"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
					resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORAGE"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "ADB_D"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "my-dadb"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
					resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
					resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
					resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
					resource.TestCheckResourceAttr(resourceName, "prefix", "pre"),
					resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
					resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
					resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.okv_key_store_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "ss2"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "ADB_D"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
					resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.okv_key_store_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "ADB_D"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
					resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORAGE"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "ADB_D"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
					resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
					resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
					resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
					resource.TestCheckResourceAttr(resourceName, "prefix", "pre"),
					resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
					resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
					resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.okv_key_store_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "ss2"),
					resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "ADB_D"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
					resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
		*/
		// {
		// 	Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies,
		// },
		// // verify Create with optionals
		// {
		// 	Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "ADB_D"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
		// 		resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		// 		resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORAGE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "ADB_D"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "my-dadb"),
		// 		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "prefix", "pre"),
		// 		resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "ss2"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "ADB_D"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

		// 		func(s *terraform.State) (err error) {
		// 			resId, err = acctest.FromInstanceState(s, resourceName, "id")
		// 			if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
		// 				if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
		// 					return errExport
		// 				}
		// 			}
		// 			return err
		// 		},
		// 	),
		// },

		// // verify Update to the compartment (the compartment will be switched back in the next step)
		// {
		// 	Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr + compartmentIdUVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Create,
		// 			acctest.RepresentationCopyWithNewProperties(DistributedDatabaseDistributedAutonomousDatabaseRepresentation, map[string]interface{}{
		// 				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
		// 			})),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "ADB_D"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
		// 		resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
		// 		resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORAGE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "ADB_D"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "my-dadb"),
		// 		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "prefix", "pre"),
		// 		resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "ss2"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "ADB_D"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

		// 		func(s *terraform.State) (err error) {
		// 			resId2, err = acctest.FromInstanceState(s, resourceName, "id")
		// 			if resId != resId2 {
		// 				return fmt.Errorf("resource recreated when it was supposed to be updated")
		// 			}
		// 			return err
		// 		},
		// 	),
		// },

		// // verify updates to updatable parameters
		// {
		// 	Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr + DistributedDatabaseDistributedAutonomousDatabaseResourceDependencies +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedAutonomousDatabaseRepresentation),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "ADB_D"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
		// 		resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		// 		resource.TestCheckResourceAttr(resourceName, "database_version", "19c"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORAGE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "ADB_D"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
		// 		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "prefix", "pre"),
		// 		resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "1.0"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.cloud_autonomous_vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "ss2"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "ADB_D"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

		// 		func(s *terraform.State) (err error) {
		// 			resId2, err = acctest.FromInstanceState(s, resourceName, "id")
		// 			if resId != resId2 {
		// 				return fmt.Errorf("Resource recreated when it was supposed to be updated.")
		// 			}
		// 			return err
		// 		},
		// 	),
		// },
		// verify datasource
		{
			Config: pluralDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_deployment_type", "ADB_D"),
				/*
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "metadata.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				*/
				resource.TestCheckResourceAttr(datasourceName, "display_name", displayName),
				// resource.TestCheckResourceAttr(datasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "distributed_autonomous_database_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "distributed_autonomous_database_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: singularDataSourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),

				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.compute_count", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_storage_size_in_gbs", "128"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.name"),
				/*
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.peer_details.0.container_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.peer_details.0.metadata.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.peer_details.0.supporting_resource_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.peer_details.0.time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.shard_group"),
				*/
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.source", "ADB_D"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "chunks", "120"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_deployment_type", "ADB_D"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gsm_details.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "latest_gsm_image.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listener_port", listenerPort),
				resource.TestCheckResourceAttr(singularDatasourceName, "listener_port_tls", listenerPortTls),
				// resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_local", onsPortLocal),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_remote", onsPortRemote),
				resource.TestCheckResourceAttr(singularDatasourceName, "prefix", prefix),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_factor", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_method", "RAFT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_unit", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.compute_count", "4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.data_storage_size_in_gbs", "128"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.name"),
				/*
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.okv_end_point_group", "okvEndPointGroup"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_ids.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.peer_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.peer_details.0.container_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.peer_details.0.fast_start_fail_over_lag_limit_in_seconds", "10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.peer_details.0.is_automatic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.peer_details.0.metadata.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.peer_details.0.shard_group"),
					resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.peer_details.0.standby_maintenance_buffer_in_days", "10"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.peer_details.0.status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.peer_details.0.supporting_resource_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.peer_details.0.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.peer_details.0.time_updated"),
				*/
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.shard_group"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.source", "ADB_D"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sharding_method", "SYSTEM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            resourceConfig,
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
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The integration test generator emitted a lifecycle state enum constant
	// `DistributedAutonomousDatabaseLifecycleStateActiveNeedsAttention` which does
	// not exist in the vendored oci-go-sdk version (states are represented as
	// separate values like ACTIVE / NEEDS_ATTENTION).
	// Sweeper logic does not require lifecycle filtering, so omit this filter.
	// See JIRA: TOP-9429
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
