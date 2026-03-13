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
	DistributedDatabaseDistributedDatabaseRequiredOnlyResource = DistributedDatabaseDistributedDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation)

	DistributedDatabaseDistributedDatabaseResourceConfig = DistributedDatabaseDistributedDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabaseRepresentation)

	DistributedDatabaseDistributedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"distributed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_database.test_distributed_database.id}`},
		// "metadata": acctest.Representation{RepType: acctest.Optional},
		// metadata is computed-only for this datasource; do not configure it.
	}
	DistributedDatabaseDistributedDatabaseMetadataSingularDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_deployment_type": acctest.Representation{RepType: acctest.Optional, Create: `EXADB_XS`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `st4_terraform_it`, Update: `st4_terraform_it`},
		// "metadata": acctest.Representation{RepType: acctest.Optional},
		// metadata is computed-only for this datasource; do not configure it.

		"state":  acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseDataSourceFilterRepresentation}}
	DistributedDatabaseDistributedDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_distributed_database_distributed_database.test_distributed_database.id}`}},
	}
	DistributedDatabaseDistributedDatabaseMetadataDataSourceRepresentation = map[string]interface{}{}

	DistributedDatabaseDistributedDatabaseRepresentation = map[string]interface{}{
		"catalog_details":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseCatalogDetailsRepresentation},
		"character_set":      acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_version":   acctest.Representation{RepType: acctest.Required, Create: `26ai`},
		"db_deployment_type": acctest.Representation{RepType: acctest.Required, Create: `EXADB_XS`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `st4_terraform_it`, Update: `st4_terraform_it`},
		// "distributed_database_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_database.test_distributed_database.id}`},
		"listener_port":        acctest.Representation{RepType: acctest.Required, Create: `10345`},
		"ncharacter_set":       acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"ons_port_local":       acctest.Representation{RepType: acctest.Required, Create: `10123`},
		"ons_port_remote":      acctest.Representation{RepType: acctest.Required, Create: `10234`},
		"prefix":               acctest.Representation{RepType: acctest.Required, Create: `st4`},
		"private_endpoint_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.private_endpoint_id}`}},
		"shard_details": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseShardDetailsRepresentation},
			{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseShardDetailsRepresentation},
			{RepType: acctest.Required, Group: DistributedDatabaseDistributedDatabaseShardDetailsRepresentation},
		},
		"sharding_method":  acctest.Representation{RepType: acctest.Required, Create: `SYSTEM`},
		"chunks":           acctest.Representation{RepType: acctest.Required, Create: `120`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseDbBackupConfigRepresentation},
		// "defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"test-namespace.test-tag": "value"}, Update: map[string]string{"test-namespace.test-tag": "updatedValue"}},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		// "listener_port_tls":               acctest.Representation{RepType: acctest.Optional, Create: `10`},
		// NOTE: exclude patch_operations from generated integration test configs.
		// The representation generator used by this test does not reliably emit
		// a valid required `value` string payload for this nested block.
		// Patch behavior is verified in dedicated update paths, not this flow.
		// "patch_operations":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabasePatchOperationsRepresentation},
		"replication_factor": acctest.Representation{RepType: acctest.Optional, Create: `3`},
		"replication_method": acctest.Representation{RepType: acctest.Optional, Create: `RAFT`},
		"replication_unit":   acctest.Representation{RepType: acctest.Optional, Create: `6`},
		// "change_db_backup_config_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "configure_sharding_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "download_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "generate_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "generate_wallet_trigger":                               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "upload_signed_certificate_and_generate_wallet_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		// "validate_network_trigger":                              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"state": acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `INACTIVE`},
	}
	DistributedDatabaseDistributedDatabaseCatalogDetailsRepresentation = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"source":         acctest.Representation{RepType: acctest.Required, Create: `EXADB_XS`},
		"vm_cluster_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_id}`},
		"kms_key_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
	}
	DistributedDatabaseDistributedDatabaseShardDetailsRepresentation = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"source":         acctest.Representation{RepType: acctest.Required, Create: `EXADB_XS`},
		"vm_cluster_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_id}`},
		"kms_key_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		// "shard_space":        acctest.Representation{RepType: acctest.Optional, Create: `shardSpace`},
		"vault_id": acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
	}
	DistributedDatabaseDistributedDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_window":            acctest.Representation{RepType: acctest.Optional, Create: `SLOT_ONE`},
		"auto_full_backup_day":          acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`},
		"auto_full_backup_window":       acctest.Representation{RepType: acctest.Optional, Create: `SLOT_ONE`},
		"backup_deletion_policy":        acctest.Representation{RepType: acctest.Optional, Create: `DELETE_IMMEDIATELY`},
		"backup_destination_details":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DistributedDatabaseDistributedDatabaseDbBackupConfigBackupDestinationDetailsRepresentation},
		"can_run_immediate_full_backup": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_auto_backup_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_remote_backup_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"recovery_window_in_days":       acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"remote_region":                 acctest.Representation{RepType: acctest.Optional, Create: `remoteRegion`},
	}
	DistributedDatabaseDistributedDatabasePatchOperationsRepresentation = map[string]interface{}{
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
	DistributedDatabaseDistributedDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORE`},
	}

	DistributedDatabaseDistributedDatabaseResourceDependencies = ""
	// acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
	// 	DefinedTagsDependencies +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
	// 	KeyResourceDependencyConfig +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation)
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	privateEndpointId := utils.GetEnvSettingWithBlankDefault("private_endpoint_id")
	privateEndpointIdVariableStr := fmt.Sprintf("variable \"private_endpoint_id\" { default = \"%s\" }\n", privateEndpointId)
	vmClusterId := utils.GetEnvSettingWithBlankDefault("vm_cluster_id")
	vmClusterIdVariableStr := fmt.Sprintf("variable \"vm_cluster_id\" { default = \"%s\" }\n", vmClusterId)
	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)
	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	resourceName := "oci_distributed_database_distributed_database.test_distributed_database"
	datasourceName := "data.oci_distributed_database_distributed_databases.test_distributed_databases"
	singularDatasourceName := "data.oci_distributed_database_distributed_database.test_distributed_database"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+privateEndpointIdVariableStr+vmClusterIdVariableStr+kmsKeyIdVariableStr+vaultIdVariableStr+DistributedDatabaseDistributedDatabaseResourceConfig, "distributeddatabase", "distributedDatabase", t)

	acctest.ResourceTest(t, testAccCheckDistributedDatabaseDistributedDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + privateEndpointIdVariableStr + vmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + DistributedDatabaseDistributedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "EXADB_XS"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.vm_cluster_id", vmClusterId),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "EXADB_XS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "st4_terraform_it"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "10345"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10123"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10234"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "st4"),
				resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "EXADB_XS"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.vm_cluster_id", vmClusterId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.source", "EXADB_XS"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.vm_cluster_id", vmClusterId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.1.kms_key_id", kmsKeyId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.source", "EXADB_XS"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.vm_cluster_id", vmClusterId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.vault_id", vaultId),
				resource.TestCheckResourceAttr(resourceName, "shard_details.2.kms_key_id", kmsKeyId),

				resource.TestCheckResourceAttr(resourceName, "sharding_method", "SYSTEM"),
			),
		},

		// // delete before next Create
		// {
		// 	Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies,
		// },
		// verify Create with optionals
		// {
		// 	Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.transport_type", "SYNC"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.shard_space", "shardSpace"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "EXADB_XS"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
		// 		resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		// 		resource.TestCheckResourceAttr(resourceName, "database_version", "23ai"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_auto_backup_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "EXADB_XS"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "CreateSdb567"),
		// 		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "gsm_ssh_public_key", "gsmSshPublicKey"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "prefix", "st4"),
		// 		resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.transport_type", "SYNC"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "shardSpace"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "EXADB_XS"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "sharding_method", "SYSTEM"),
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
		// 	Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Create,
		// 			acctest.RepresentationCopyWithNewProperties(DistributedDatabaseDistributedDatabaseRepresentation, map[string]interface{}{
		// 				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
		// 			})),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.transport_type", "SYNC"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.shard_space", "shardSpace"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "EXADB_XS"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
		// 		resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
		// 		resource.TestCheckResourceAttr(resourceName, "database_version", "23ai"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_auto_backup_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "EXADB_XS"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "CreateSdb567"),
		// 		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "gsm_ssh_public_key", "gsmSshPublicKey"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "prefix", "st4"),
		// 		resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.transport_type", "SYNC"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "shardSpace"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "EXADB_XS"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "sharding_method", "SYSTEM"),
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
		// 	Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabaseResourceDependencies +
		// 		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabaseRepresentation),
		// 	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_details.0.transport_type", "SYNC"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.peer_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.shard_space", "shardSpace"),
		// 		resource.TestCheckResourceAttr(resourceName, "catalog_details.0.source", "EXADB_XS"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vault_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
		// 		resource.TestCheckResourceAttr(resourceName, "chunks", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		// 		resource.TestCheckResourceAttr(resourceName, "database_version", "23ai"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "db_backup_config.0.backup_destination_details.0.dbrs_policy_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.id", "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.type", "NFS"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_auto_backup_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.remote_region", "remoteRegion"),
		// 		resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "EXADB_XS"),
		// 		resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
		// 		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "gsm_ssh_public_key", "gsmSshPublicKey"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "id"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_local", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "prefix", "st4"),
		// 		resource.TestCheckResourceAttr(resourceName, "private_endpoint_ids.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_factor", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_method", "RAFT"),
		// 		resource.TestCheckResourceAttr(resourceName, "replication_unit", "10"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "BEstrO0ng_#11"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.kms_key_version_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.#", "1"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.protection_mode", "MAXIMUM_AVAILABILITY"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.shard_group"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.time_updated"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_details.0.transport_type", "SYNC"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.peer_vm_cluster_ids.#", "1"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "shardSpace"),
		// 		resource.TestCheckResourceAttr(resourceName, "shard_details.0.source", "EXADB_XS"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vault_id"),
		// 		resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.vm_cluster_id"),
		// 		resource.TestCheckResourceAttr(resourceName, "sharding_method", "SYSTEM"),
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_databases", "test_distributed_databases", acctest.Optional, acctest.Update, DistributedDatabaseDistributedDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + privateEndpointIdVariableStr + vmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + DistributedDatabaseDistributedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_deployment_type", "EXADB_XS"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "st4_terraform_it"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "distributed_database_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "distributed_database_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + privateEndpointIdVariableStr + vmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + DistributedDatabaseDistributedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.container_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.db_home_id"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.metadata.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.shard_group"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.source", "EXADB_XS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "chunks", "120"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_version", "26ai"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.auto_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_deletion_policy", "DELETE_IMMEDIATELY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.id", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.internet_proxy", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.is_remote", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.remote_region", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.type", "OBJECT_STORE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.vpc_password", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.backup_destination_details.0.vpc_user", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.can_run_immediate_full_backup", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.is_auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.is_remote_backup_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.0.remote_region", "remoteRegion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_deployment_type", "EXADB_XS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "st4_terraform_it"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gsm_details.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "latest_gsm_image_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listener_port", "10345"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_port_tls"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_local", "10123"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_remote", "10234"),
				resource.TestCheckResourceAttr(singularDatasourceName, "prefix", "st4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_factor", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_method", "RAFT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_unit", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.container_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.db_home_id"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.metadata.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.shard_group"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.shard_space", "shardSpace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.source", "EXADB_XS"),
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
			Config:            config + compartmentIdVariableStr + privateEndpointIdVariableStr + vmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + DistributedDatabaseDistributedDatabaseResourceConfig,
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
	// WORKAROUND FOR GENERATED CODE ISSUE:
	// The integration test generator emitted a lifecycle state enum constant
	// `DistributedDatabaseLifecycleStateActiveNeedsAttention` which does
	// not exist in the vendored oci-go-sdk version (states are represented as
	// separate values like ACTIVE / NEEDS_ATTENTION).
	// Sweeper logic does not require lifecycle filtering, so omit this filter.
	// See JIRA: TOP-9429
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
