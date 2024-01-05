// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// fake
var (
	CloudMigrationsKmsKeyId        = `${var.vaultId}`
	CloudMigrationsImageId         = `${var.imageId}`
	CloudMigrationsMigrationPlanId = `${var.migrationPlanId}`
	CloudMigrationsSubnetId        = `${var.subnetId}`
	CloudMigrationsVlanId          = "fakevlan"
	CloudMigrationsBootVolumeId    = `${var.bootVolumeId}`
	CloudMigrationsSourceType      = "image"
	CloudMigrationsShape           = "VM.Standard.E4.Flex"
	CloudMigrationsAD              = ""

	CloudMigrationsTargetAssetRequiredOnlyResource = CloudMigrationsTargetAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Required, acctest.Create, CloudMigrationsTargetAssetRepresentation)

	CloudMigrationsTargetAssetResourceConfig = CloudMigrationsTargetAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Optional, acctest.Update, CloudMigrationsTargetAssetRepresentation)

	CloudMigrationsCloudMigrationsTargetAssetSingularDataSourceRepresentation = map[string]interface{}{
		"target_asset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_target_asset.test_target_asset.id}`},
	}

	CloudMigrationsCloudMigrationsTargetAssetDataSourceRepresentation = map[string]interface{}{
		"migration_plan_id": acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsMigrationPlanId},
	}

	CloudMigrationsTargetAssetRepresentation = map[string]interface{}{
		"is_excluded_from_execution": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"migration_plan_id":          acctest.Representation{RepType: acctest.Required, Create: CloudMigrationsMigrationPlanId},
		"preferred_shape_type":       acctest.Representation{RepType: acctest.Required, Create: `VM`, Update: `VM_INTEL`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `INSTANCE`},
		"user_spec":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudMigrationsTargetAssetUserSpecRepresentation},
		"block_volumes_performance":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"ms_license":                 acctest.Representation{RepType: acctest.Optional, Create: `msLicense`, Update: `msLicense2`},
	}
	CloudMigrationsTargetAssetUserSpecRepresentation = map[string]interface{}{
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsTargetAssetUserSpecAgentConfigRepresentation},
		"availability_domain":                 acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsAD, Update: CloudMigrationsAD},
		"compartment_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsTargetAssetUserSpecCreateVnicDetailsRepresentation},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `faultDomain`, Update: `faultDomain2`},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`, Update: `hostnameLabel2`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsTargetAssetUserSpecInstanceOptionsRepresentation},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`, Update: `ipxeScript2`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"preemptible_instance_config":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsTargetAssetUserSpecPreemptibleInstanceConfigRepresentation},
		"shape":                               acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsShape, Update: CloudMigrationsShape},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsTargetAssetUserSpecShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsTargetAssetUserSpecSourceDetailsRepresentation},
	}
	CloudMigrationsTargetAssetUserSpecAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_management_disabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_monitoring_disabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"plugins_config":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudMigrationsTargetAssetUserSpecAgentConfigPluginsConfigRepresentation},
	}
	CloudMigrationsTargetAssetUserSpecCreateVnicDetailsRepresentation = map[string]interface{}{
		"assign_private_dns_record": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"assign_public_ip":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"hostname_label":            acctest.Representation{RepType: acctest.Optional, Create: `hostnameLabel`, Update: `hostnameLabel2`},
		//"nsg_ids":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}, Update: []string{`nsgIds2`}},
		"private_ip":             acctest.Representation{RepType: acctest.Optional, Create: `privateIp`, Update: `privateIp2`},
		"skip_source_dest_check": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"subnet_id":              acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsSubnetId},
		"vlan_id":                acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsVlanId},
	}
	CloudMigrationsTargetAssetUserSpecInstanceOptionsRepresentation = map[string]interface{}{
		"are_legacy_imds_endpoints_disabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	CloudMigrationsTargetAssetUserSpecPreemptibleInstanceConfigRepresentation = map[string]interface{}{
		"preemption_action": acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudMigrationsTargetAssetUserSpecPreemptibleInstanceConfigPreemptionActionRepresentation},
	}
	CloudMigrationsTargetAssetUserSpecShapeConfigRepresentation = map[string]interface{}{
		"baseline_ocpu_utilization": acctest.Representation{RepType: acctest.Optional, Create: `BASELINE_1_8`, Update: `BASELINE_1_2`},
		"memory_in_gbs":             acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"ocpus":                     acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	CloudMigrationsTargetAssetUserSpecSourceDetailsRepresentation = map[string]interface{}{
		"source_type":             acctest.Representation{RepType: acctest.Required, Create: CloudMigrationsSourceType, Update: CloudMigrationsSourceType},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"boot_volume_vpus_per_gb": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"image_id":                acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsImageId},
		"kms_key_id":              acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsKmsKeyId},
	}
	CloudMigrationsTargetAssetUserSpecAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
		"name":          acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}
	CloudMigrationsTargetAssetUserSpecPreemptibleInstanceConfigPreemptionActionRepresentation = map[string]interface{}{
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `TERMINATE`},
		"preserve_boot_volume": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	CloudMigrationsTargetAssetResourceDependencies = ""
)

// issue-routing-tag: cloud_migrations/default
func TestCloudMigrationsTargetAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudMigrationsTargetAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	vaultId := utils.GetEnvSettingWithBlankDefault("vaultId")
	vaultIdVariableStr := fmt.Sprintf("variable \"vaultId\" { default = \"%s\" }\n", vaultId)
	imageId := utils.GetEnvSettingWithBlankDefault("imageId")
	imageIdVariableStr := fmt.Sprintf("variable \"imageId\" { default = \"%s\" }\n", imageId)
	migrationPlanId := utils.GetEnvSettingWithBlankDefault("migrationPlanId")
	migrationPlanIdVariableStr := fmt.Sprintf("variable \"migrationPlanId\" { default = \"%s\" }\n", migrationPlanId)
	subnetId := utils.GetEnvSettingWithBlankDefault("subnetId")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnetId\" { default = \"%s\" }\n", subnetId)
	bootVolumeId := utils.GetEnvSettingWithBlankDefault("bootVolumeId")
	bootVolumeIdVariableStr := fmt.Sprintf("variable \"bootVolumeId\" { default = \"%s\" }\n", bootVolumeId)

	variableStr := compartmentIdVariableStr + vaultIdVariableStr + imageIdVariableStr + migrationPlanIdVariableStr + subnetIdVariableStr + bootVolumeIdVariableStr

	resourceName := "oci_cloud_migrations_target_asset.test_target_asset"
	datasourceName := "data.oci_cloud_migrations_target_assets.test_target_assets"
	singularDatasourceName := "data.oci_cloud_migrations_target_asset.test_target_asset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudMigrationsTargetAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Optional, acctest.Create, CloudMigrationsTargetAssetRepresentation), "cloudmigrations", "targetAsset", t)

	acctest.ResourceTest(t, testAccCheckCloudMigrationsTargetAssetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + CloudMigrationsTargetAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Required, acctest.Create, CloudMigrationsTargetAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_excluded_from_execution", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_plan_id"),
				resource.TestCheckResourceAttr(resourceName, "preferred_shape_type", "VM"),
				resource.TestCheckResourceAttr(resourceName, "type", "INSTANCE"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variableStr + CloudMigrationsTargetAssetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variableStr + CloudMigrationsTargetAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Optional, acctest.Create, CloudMigrationsTargetAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "block_volumes_performance", "10"),
				resource.TestCheckResourceAttr(resourceName, "estimated_cost.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_excluded_from_execution", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_plan_id"),
				resource.TestCheckResourceAttr(resourceName, "ms_license", "msLicense"),
				resource.TestCheckResourceAttr(resourceName, "preferred_shape_type", "VM"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_assessed"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "INSTANCE"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.are_all_plugins_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.is_management_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.is_monitoring_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.plugins_config.0.desired_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.plugins_config.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.assign_private_dns_record", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.assign_public_ip", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.private_ip", "privateIp"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.skip_source_dest_check", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "user_spec.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_spec.0.create_vnic_details.0.vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.fault_domain", "faultDomain"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.hostname_label", "hostnameLabel"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.instance_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.instance_options.0.are_legacy_imds_endpoints_disabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.ipxe_script", "ipxeScript"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.is_pv_encryption_in_transit_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.0.preserve_boot_volume", "false"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.0.type", "TERMINATE"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape", CloudMigrationsShape),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.0.baseline_ocpu_utilization", "BASELINE_1_8"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.0.memory_in_gbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.source_details.0.source_type", CloudMigrationsSourceType),

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

		// verify updates to updatable parameters
		{
			Config: config + variableStr + CloudMigrationsTargetAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Optional, acctest.Update, CloudMigrationsTargetAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "block_volumes_performance", "11"),
				resource.TestCheckResourceAttr(resourceName, "estimated_cost.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_excluded_from_execution", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_plan_id"),
				resource.TestCheckResourceAttr(resourceName, "ms_license", "msLicense2"),
				resource.TestCheckResourceAttr(resourceName, "preferred_shape_type", "VM_INTEL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_assessed"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "INSTANCE"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.are_all_plugins_disabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.is_management_disabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.is_monitoring_disabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.plugins_config.0.desired_state", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.agent_config.0.plugins_config.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.hostname_label", "hostnameLabel2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.private_ip", "privateIp2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.create_vnic_details.0.skip_source_dest_check", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "user_spec.0.create_vnic_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_spec.0.create_vnic_details.0.vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.fault_domain", "faultDomain2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.hostname_label", "hostnameLabel2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.instance_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.ipxe_script", "ipxeScript2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.is_pv_encryption_in_transit_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.0.preserve_boot_volume", "true"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.0.type", "TERMINATE"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape", CloudMigrationsShape),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.0.baseline_ocpu_utilization", "BASELINE_1_2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.0.memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user_spec.0.source_details.0.source_type", CloudMigrationsSourceType),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_target_assets", "test_target_assets", acctest.Optional, acctest.Update, CloudMigrationsCloudMigrationsTargetAssetDataSourceRepresentation) +
				variableStr + CloudMigrationsTargetAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Optional, acctest.Update, CloudMigrationsTargetAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "target_asset_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_target_asset", "test_target_asset", acctest.Required, acctest.Create, CloudMigrationsCloudMigrationsTargetAssetSingularDataSourceRepresentation) +
				variableStr + CloudMigrationsTargetAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_asset_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "block_volumes_performance", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compatibility_messages.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "estimated_cost.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_excluded_from_execution", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ms_license", "msLicense2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "preferred_shape_type", "VM_INTEL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_assessed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "INSTANCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.agent_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.agent_config.0.are_all_plugins_disabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.agent_config.0.is_management_disabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.agent_config.0.is_monitoring_disabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.agent_config.0.plugins_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.agent_config.0.plugins_config.0.desired_state", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.agent_config.0.plugins_config.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.create_vnic_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.create_vnic_details.0.assign_private_dns_record", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.create_vnic_details.0.assign_public_ip", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.create_vnic_details.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.create_vnic_details.0.hostname_label", "hostnameLabel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.create_vnic_details.0.private_ip", "privateIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.create_vnic_details.0.skip_source_dest_check", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.fault_domain", "faultDomain2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.hostname_label", "hostnameLabel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.instance_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.instance_options.0.are_legacy_imds_endpoints_disabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.ipxe_script", "ipxeScript2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.is_pv_encryption_in_transit_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.preemptible_instance_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.0.preserve_boot_volume", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.preemptible_instance_config.0.preemption_action.0.type", "TERMINATE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.shape", CloudMigrationsShape),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.shape_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.shape_config.0.baseline_ocpu_utilization", "BASELINE_1_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.shape_config.0.memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.shape_config.0.ocpus", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.source_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_spec.0.source_details.0.source_type", CloudMigrationsSourceType),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudMigrationsTargetAssetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudMigrationsTargetAssetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_migrations_target_asset" {
			noResourceFound = false
			request := oci_cloud_migrations.GetTargetAssetRequest{}

			tmp := rs.Primary.ID
			request.TargetAssetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")

			response, err := client.GetTargetAsset(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_migrations.TargetAssetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudMigrationsTargetAsset") {
		resource.AddTestSweepers("CloudMigrationsTargetAsset", &resource.Sweeper{
			Name:         "CloudMigrationsTargetAsset",
			Dependencies: acctest.DependencyGraph["targetAsset"],
			F:            sweepCloudMigrationsTargetAssetResource,
		})
	}
}

func sweepCloudMigrationsTargetAssetResource(compartment string) error {
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()
	targetAssetIds, err := getCloudMigrationsTargetAssetIds(compartment)
	if err != nil {
		return err
	}
	for _, targetAssetId := range targetAssetIds {
		if ok := acctest.SweeperDefaultResourceId[targetAssetId]; !ok {
			deleteTargetAssetRequest := oci_cloud_migrations.DeleteTargetAssetRequest{}

			deleteTargetAssetRequest.TargetAssetId = &targetAssetId

			deleteTargetAssetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")
			_, error := migrationClient.DeleteTargetAsset(context.Background(), deleteTargetAssetRequest)
			if error != nil {
				fmt.Printf("Error deleting TargetAsset %s %s, It is possible that the resource is already deleted. Please verify manually \n", targetAssetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &targetAssetId, CloudMigrationsTargetAssetSweepWaitCondition, time.Duration(3*time.Minute),
				CloudMigrationsTargetAssetSweepResponseFetchOperation, "cloud_migrations", true)
		}
	}
	return nil
}

func getCloudMigrationsTargetAssetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TargetAssetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()

	listTargetAssetsRequest := oci_cloud_migrations.ListTargetAssetsRequest{}
	listTargetAssetsRequest.MigrationPlanId = &CloudMigrationsMigrationPlanId //some migrationPlanId.
	listTargetAssetsRequest.LifecycleState = oci_cloud_migrations.TargetAssetLifecycleStateActive
	listTargetAssetsResponse, err := migrationClient.ListTargetAssets(context.Background(), listTargetAssetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting TargetAsset list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, targetAsset := range listTargetAssetsResponse.Items {
		id := *targetAsset.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TargetAssetId", id)
	}
	return resourceIds, nil
}

func CloudMigrationsTargetAssetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if targetAssetResponse, ok := response.Response.(oci_cloud_migrations.GetTargetAssetResponse); ok {
		return targetAssetResponse.GetLifecycleState() != oci_cloud_migrations.TargetAssetLifecycleStateDeleted
	}
	return false
}

func CloudMigrationsTargetAssetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MigrationClient().GetTargetAsset(context.Background(), oci_cloud_migrations.GetTargetAssetRequest{
		TargetAssetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
