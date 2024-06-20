// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GloballyDistributedDatabaseShardedDatabaseRequiredOnlyResource = GloballyDistributedDatabaseShardedDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Required, acctest.Create, GloballyDistributedDatabaseShardedDatabaseRepresentation)

	GloballyDistributedDatabaseShardedDatabaseResourceConfig = GloballyDistributedDatabaseShardedDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Optional, acctest.Update, GloballyDistributedDatabaseShardedDatabaseRepresentation)

	GloballyDistributedDatabaseShardedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"sharded_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_globally_distributed_database_sharded_database.test_sharded_database.id}`},
		"metadata":            acctest.Representation{RepType: acctest.Optional, Create: `{}`},
	}

	GloballyDistributedDatabaseShardedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `Sdb0001`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GloballyDistributedDatabaseShardedDatabaseDataSourceFilterRepresentation}}
	GloballyDistributedDatabaseShardedDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_globally_distributed_database_sharded_database.test_sharded_database.id}`}},
	}

	GloballyDistributedDatabaseShardedDatabaseRepresentation = map[string]interface{}{
		"catalog_details":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: GloballyDistributedDatabaseShardedDatabaseCatalogDetailsRepresentation},
		"character_set":                   acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"sharding_method":                 acctest.Representation{RepType: acctest.Required, Create: `USER`},
		"cluster_certificate_common_name": acctest.Representation{RepType: acctest.Required, Create: `gdad_cert`},
		"chunks":                          acctest.Representation{RepType: acctest.Optional, Create: `120`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_deployment_type":              acctest.Representation{RepType: acctest.Required, Create: `DEDICATED`},
		"db_version":                      acctest.Representation{RepType: acctest.Required, Create: `19c`},
		"db_workload":                     acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `Sdb0001`, Update: `displayName2`},
		"listener_port":                   acctest.Representation{RepType: acctest.Required, Create: `40001`},
		"listener_port_tls":               acctest.Representation{RepType: acctest.Required, Create: `40002`},
		"ncharacter_set":                  acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"ons_port_local":                  acctest.Representation{RepType: acctest.Required, Create: `40003`},
		"ons_port_remote":                 acctest.Representation{RepType: acctest.Required, Create: `40004`},
		"prefix":                          acctest.Representation{RepType: acctest.Required, Create: `s01`},
		"shard_details":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: GloballyDistributedDatabaseShardedDatabaseShardDetailsRepresentation},
		//"sharded_database_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_globally_distributed_database_sharded_database.test_sharded_database.id}`},
		//"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		/*"patch_operations":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: GloballyDistributedDatabaseShardedDatabasePatchOperationsRepresentation},
		"configure_gsms_trigger":          acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"configure_sharding_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"download_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"generate_gsm_certificate_signing_request_trigger":      acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"generate_wallet_trigger":                               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"get_connection_string_trigger":                         acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"start_database_trigger":                                acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"stop_database_trigger":                                 acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"upload_signed_certificate_and_generate_wallet_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"validate_network_trigger":                              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},*/
	}
	GloballyDistributedDatabaseShardedDatabaseCatalogDetailsRepresentation = map[string]interface{}{
		"admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `WElcomeHome1234##`},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_autonomous_vm_cluster_id}`},
		"compute_count":                  acctest.Representation{RepType: acctest.Required, Create: `2`},
		"data_storage_size_in_gbs":       acctest.Representation{RepType: acctest.Required, Create: `32`},
		"is_auto_scaling_enabled":        acctest.Representation{RepType: acctest.Required, Create: `false`},
		"encryption_key_details":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: GloballyDistributedDatabaseShardedDatabaseCatalogDetailsEncryptionKeyDetailsRepresentation},
		//"peer_cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.cloud_autonomous_vm_cluster_id}`},
	}
	GloballyDistributedDatabaseShardedDatabaseShardDetailsRepresentation = map[string]interface{}{
		"admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `WElcomeHome1234##`},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_autonomous_vm_cluster_id}`},
		"compute_count":                  acctest.Representation{RepType: acctest.Required, Create: `2`},
		"data_storage_size_in_gbs":       acctest.Representation{RepType: acctest.Required, Create: `32`},
		"is_auto_scaling_enabled":        acctest.Representation{RepType: acctest.Required, Create: `false`},
		"encryption_key_details":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: GloballyDistributedDatabaseShardedDatabaseShardDetailsEncryptionKeyDetailsRepresentation},
		//"peer_cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.cloud_autonomous_vm_cluster_id}`},
		"shard_space": acctest.Representation{RepType: acctest.Required, Create: `xyz`},
	}
	/*GloballyDistributedDatabaseShardedDatabasePatchOperationsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"value":     acctest.Representation{RepType: acctest.Optional, Create: `{}`},
	}*/
	GloballyDistributedDatabaseShardedDatabaseCatalogDetailsEncryptionKeyDetailsRepresentation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
		//"kms_key_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
	}
	GloballyDistributedDatabaseShardedDatabaseShardDetailsEncryptionKeyDetailsRepresentation = map[string]interface{}{
		"kms_key_id": acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
		//"kms_key_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
	}

	GloballyDistributedDatabaseShardedDatabaseResourceDependencies = "" /*acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
	AvailabilityDomainConfig +*/
	/*DefinedTagsDependencies*/ /*+
	acctest.GenerateResourceFromRepresentationMap("oci_kms_key_version", "test_key_version", acctest.Required, acctest.Create, KmsKeyVersionRepresentation) +
	KeyResourceDependencyConfig +
	acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation)*/
)

// issue-routing-tag: globally_distributed_database/default
func TestGloballyDistributedDatabaseShardedDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGloballyDistributedDatabaseShardedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	/*vcnId := utils.GetEnvSettingWithBlankDefault("vcn_ocid")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)*/

	/*subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)*/

	cloudAutonomousVmClusterId := utils.GetEnvSettingWithBlankDefault("cloud_autonomous_vm_cluster_id")
	cloudAutonomousVmClusterIdVariableStr := fmt.Sprintf("variable \"cloud_autonomous_vm_cluster_id\" { default = \"%s\" }\n", cloudAutonomousVmClusterId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	resourceName := "oci_globally_distributed_database_sharded_database.test_sharded_database"
	datasourceName := "data.oci_globally_distributed_database_sharded_databases.test_sharded_databases"
	singularDatasourceName := "data.oci_globally_distributed_database_sharded_database.test_sharded_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+cloudAutonomousVmClusterIdVariableStr+kmsKeyIdVariableStr+vaultIdVariableStr+GloballyDistributedDatabaseShardedDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Optional, acctest.Create, GloballyDistributedDatabaseShardedDatabaseRepresentation), "globallydistributeddatabase", "shardedDatabase", t)

	acctest.ResourceTest(t, testAccCheckGloballyDistributedDatabaseShardedDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + GloballyDistributedDatabaseShardedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Required, acctest.Create, GloballyDistributedDatabaseShardedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Sdb0001"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "40001"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "40002"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "40003"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "40004"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "s01"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
				resource.TestCheckResourceAttr(resourceName, "cluster_certificate_common_name", "gdad_cert"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + GloballyDistributedDatabaseShardedDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + GloballyDistributedDatabaseShardedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Optional, acctest.Create, GloballyDistributedDatabaseShardedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
				//resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "chunks", "120"),
				resource.TestCheckResourceAttr(resourceName, "cluster_certificate_common_name", "gdad_cert"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Sdb0001"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "40001"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "40002"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "40003"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "40004"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "s01"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
				//resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "xyz"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + GloballyDistributedDatabaseShardedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GloballyDistributedDatabaseShardedDatabaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
				//resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "chunks", "120"),
				resource.TestCheckResourceAttr(resourceName, "cluster_certificate_common_name", "gdad_cert"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Sdb0001"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "40001"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "40002"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "40003"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "40004"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "s01"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
				//resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "xyz"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + GloballyDistributedDatabaseShardedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Optional, acctest.Update, GloballyDistributedDatabaseShardedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.encryption_key_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.name"),
				//resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.shard_group"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "catalog_details.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName, "chunks", "120"),
				resource.TestCheckResourceAttr(resourceName, "cluster_certificate_common_name", "gdad_cert"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_deployment_type", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "listener_port", "40001"),
				resource.TestCheckResourceAttr(resourceName, "listener_port_tls", "40002"),
				resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_local", "40003"),
				resource.TestCheckResourceAttr(resourceName, "ons_port_remote", "40004"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "s01"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.admin_password", "WElcomeHome1234##"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.compute_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.encryption_key_details.0.vault_id"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.name"),
				//resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.shard_group"),
				resource.TestCheckResourceAttr(resourceName, "shard_details.0.shard_space", "xyz"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.status"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "shard_details.0.time_updated"),
				resource.TestCheckResourceAttr(resourceName, "sharding_method", "USER"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_globally_distributed_database_sharded_databases", "test_sharded_databases", acctest.Optional, acctest.Update, GloballyDistributedDatabaseShardedDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + compartmentIdUVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + GloballyDistributedDatabaseShardedDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Optional, acctest.Update, GloballyDistributedDatabaseShardedDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),

				//resource.TestCheckResourceAttr(datasourceName, "sharded_database_collection.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "sharded_database_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_globally_distributed_database_sharded_database", "test_sharded_database", acctest.Required, acctest.Create, GloballyDistributedDatabaseShardedDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + compartmentIdUVariableStr + cloudAutonomousVmClusterIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + GloballyDistributedDatabaseShardedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(singularDatasourceName, "metadata", "metadata"),

				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.compute_count", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.container_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.container_database_parent_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "catalog_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.shard_group"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.time_ssl_certificate_expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_details.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "chunks", "120"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_certificate_common_name", "gdad_cert"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_strings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_deployment_type", "DEDICATED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_version", "19c"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gsms.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listener_port", "40001"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listener_port_tls", "40002"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_local", "40003"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ons_port_remote", "40004"),
				resource.TestCheckResourceAttr(singularDatasourceName, "prefix", "s01"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.compute_count", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.container_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.container_database_parent_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.data_storage_size_in_gbs", "32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.encryption_key_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.is_auto_scaling_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.shard_group"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shard_details.0.shard_space", "xyz"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.supporting_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.time_ssl_certificate_expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shard_details.0.time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sharding_method", "USER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_zone"),
			),
		},
		// verify resource import
		{
			Config:            config + GloballyDistributedDatabaseShardedDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"patch_operations", "catalog_details.0.admin_password", "shard_details.0.admin_password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckGloballyDistributedDatabaseShardedDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ShardedDatabaseServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_globally_distributed_database_sharded_database" {
			noResourceFound = false
			request := oci_globally_distributed_database.GetShardedDatabaseRequest{}

			if value, ok := rs.Primary.Attributes["metadata"]; ok {
				request.Metadata = &value
			}

			tmp := rs.Primary.ID
			request.ShardedDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "globally_distributed_database")

			response, err := client.GetShardedDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_globally_distributed_database.ShardedDatabaseLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("GloballyDistributedDatabaseShardedDatabase") {
		resource.AddTestSweepers("GloballyDistributedDatabaseShardedDatabase", &resource.Sweeper{
			Name:         "GloballyDistributedDatabaseShardedDatabase",
			Dependencies: acctest.DependencyGraph["shardedDatabase"],
			F:            sweepGloballyDistributedDatabaseShardedDatabaseResource,
		})
	}
}

func sweepGloballyDistributedDatabaseShardedDatabaseResource(compartment string) error {
	shardedDatabaseServiceClient := acctest.GetTestClients(&schema.ResourceData{}).ShardedDatabaseServiceClient()
	shardedDatabaseIds, err := getGloballyDistributedDatabaseShardedDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, shardedDatabaseId := range shardedDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[shardedDatabaseId]; !ok {
			deleteShardedDatabaseRequest := oci_globally_distributed_database.DeleteShardedDatabaseRequest{}

			deleteShardedDatabaseRequest.ShardedDatabaseId = &shardedDatabaseId

			deleteShardedDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "globally_distributed_database")
			_, error := shardedDatabaseServiceClient.DeleteShardedDatabase(context.Background(), deleteShardedDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ShardedDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", shardedDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &shardedDatabaseId, GloballyDistributedDatabaseShardedDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				GloballyDistributedDatabaseShardedDatabaseSweepResponseFetchOperation, "globally_distributed_database", true)
		}
	}
	return nil
}

func getGloballyDistributedDatabaseShardedDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ShardedDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	shardedDatabaseServiceClient := acctest.GetTestClients(&schema.ResourceData{}).ShardedDatabaseServiceClient()

	listShardedDatabasesRequest := oci_globally_distributed_database.ListShardedDatabasesRequest{}
	listShardedDatabasesRequest.CompartmentId = &compartmentId
	listShardedDatabasesRequest.LifecycleState = oci_globally_distributed_database.ShardedDatabaseLifecycleStateNeedsAttention
	listShardedDatabasesResponse, err := shardedDatabaseServiceClient.ListShardedDatabases(context.Background(), listShardedDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ShardedDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, shardedDatabase := range listShardedDatabasesResponse.Items {
		id := *shardedDatabase.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ShardedDatabaseId", id)
	}
	return resourceIds, nil
}

func GloballyDistributedDatabaseShardedDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if shardedDatabaseResponse, ok := response.Response.(oci_globally_distributed_database.GetShardedDatabaseResponse); ok {
		return shardedDatabaseResponse.GetLifecycleState() != oci_globally_distributed_database.ShardedDatabaseLifecycleStateDeleted
	}
	return false
}

func GloballyDistributedDatabaseShardedDatabaseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ShardedDatabaseServiceClient().GetShardedDatabase(context.Background(), oci_globally_distributed_database.GetShardedDatabaseRequest{
		ShardedDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
