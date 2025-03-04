// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LustreFileStorageLustreFileSystemRequiredOnlyResource = LustreFileStorageLustreFileSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Required, acctest.Create, LustreFileStorageLustreFileSystemRepresentation)

	LustreFileStorageLustreFileSystemResourceConfig = LustreFileStorageLustreFileSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Optional, acctest.Update, LustreFileStorageLustreFileSystemRepresentation)

	LustreFileStorageLustreFileSystemSingularDataSourceRepresentation = map[string]interface{}{
		"lustre_file_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id}`},
	}

	LustreFileStorageLustreFileSystemDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: LustreFileStorageLustreFileSystemDataSourceFilterRepresentation}}
	LustreFileStorageLustreFileSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id}`}},
	}

	LustreFileStorageLustreFileSystemRepresentation = map[string]interface{}{
		"availability_domain":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"capacity_in_gbs":            acctest.Representation{RepType: acctest.Required, Create: `31200`, Update: nil},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"file_system_name":           acctest.Representation{RepType: acctest.Required, Create: `lustre`},
		"performance_tier":           acctest.Representation{RepType: acctest.Required, Create: `MBPS_PER_TB_125`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"root_squash_configuration":  acctest.RepresentationGroup{RepType: acctest.Required, Group: LustreFileStorageLustreFileSystemRootSquashConfigurationRepresentation},
		"subnet_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"cluster_placement_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"file_system_description":    acctest.Representation{RepType: acctest.Optional, Create: `fileSystemDescription`, Update: `fileSystemDescription2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"nsg_ids":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIds`}, Update: []string{`nsgIds2`}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreLfstSystemTagsChangesRep},
	}
	IgnoreLfstSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	LustreFileStorageLustreFileSystemRootSquashConfigurationRepresentation = map[string]interface{}{
		"client_exceptions": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"identity_squash":   acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `NONE`},
		"squash_gid":        acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"squash_uid":        acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
	}

	LustreFileStorageLustreFileSystemResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cluster_placement_groups_cluster_placement_group", "test_cluster_placement_group", acctest.Required, acctest.Create, ClusterPlacementGroupsClusterPlacementGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Required, acctest.Create, FileStorageFileSystemRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig
)

// issue-routing-tag: lustre_file_storage/default
func TestLustreFileStorageLustreFileSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLustreFileStorageLustreFileSystemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_lustre_file_storage_lustre_file_system.test_lustre_file_system"
	datasourceName := "data.oci_lustre_file_storage_lustre_file_systems.test_lustre_file_systems"
	singularDatasourceName := "data.oci_lustre_file_storage_lustre_file_system.test_lustre_file_system"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LustreFileStorageLustreFileSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Optional, acctest.Create, LustreFileStorageLustreFileSystemRepresentation), "lustrefilestorage", "lustreFileSystem", t)

	acctest.ResourceTest(t, testAccCheckLustreFileStorageLustreFileSystemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LustreFileStorageLustreFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Required, acctest.Create, LustreFileStorageLustreFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_in_gbs", "31200"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_name"),
				resource.TestCheckResourceAttr(resourceName, "performance_tier", "MBPS_PER_TB_125"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LustreFileStorageLustreFileSystemResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LustreFileStorageLustreFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Optional, acctest.Create, LustreFileStorageLustreFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_in_gbs", "31200"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_placement_group_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "file_system_description", "fileSystemDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "lnet"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "major_version"),
				resource.TestCheckResourceAttrSet(resourceName, "management_service_address"),
				resource.TestCheckResourceAttr(resourceName, "performance_tier", "MBPS_PER_TB_125"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.client_exceptions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.identity_squash", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.squash_gid", ""),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.squash_uid", ""),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_tags"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LustreFileStorageLustreFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LustreFileStorageLustreFileSystemRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_in_gbs", "31200"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_placement_group_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "file_system_description", "fileSystemDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "lnet"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "major_version"),
				resource.TestCheckResourceAttrSet(resourceName, "management_service_address"),
				resource.TestCheckResourceAttr(resourceName, "performance_tier", "MBPS_PER_TB_125"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.client_exceptions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.identity_squash", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.squash_gid", ""),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.squash_uid", ""),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_tags"),
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
			Config: config + compartmentIdVariableStr + LustreFileStorageLustreFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Optional, acctest.Update, LustreFileStorageLustreFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "capacity_in_gbs", "31200"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_placement_group_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "file_system_description", "fileSystemDescription2"),
				resource.TestCheckResourceAttrSet(resourceName, "file_system_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "lnet"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "major_version"),
				resource.TestCheckResourceAttrSet(resourceName, "management_service_address"),
				resource.TestCheckResourceAttr(resourceName, "performance_tier", "MBPS_PER_TB_125"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.client_exceptions.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.identity_squash", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.squash_gid", ""),
				resource.TestCheckResourceAttr(resourceName, "root_squash_configuration.0.squash_uid", ""),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "system_tags"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_systems", "test_lustre_file_systems", acctest.Optional, acctest.Update, LustreFileStorageLustreFileSystemDataSourceRepresentation) +
				compartmentIdVariableStr + LustreFileStorageLustreFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Optional, acctest.Update, LustreFileStorageLustreFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "lustre_file_system_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "lustre_file_system_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_lustre_file_storage_lustre_file_system", "test_lustre_file_system", acctest.Required, acctest.Create, LustreFileStorageLustreFileSystemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LustreFileStorageLustreFileSystemResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lustre_file_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capacity_in_gbs", "31200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "file_system_description", "fileSystemDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lnet"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "major_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_service_address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "performance_tier", "MBPS_PER_TB_125"),
				resource.TestCheckResourceAttr(singularDatasourceName, "root_squash_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "root_squash_configuration.0.client_exceptions.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "root_squash_configuration.0.identity_squash", "NONE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "root_squash_configuration.0.squash_gid", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "root_squash_configuration.0.squash_uid", ""),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_billing_cycle_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + LustreFileStorageLustreFileSystemRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLustreFileStorageLustreFileSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LustreFileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_lustre_file_storage_lustre_file_system" {
			noResourceFound = false
			request := oci_lustre_file_storage.GetLustreFileSystemRequest{}

			tmp := rs.Primary.ID
			request.LustreFileSystemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "lustre_file_storage")

			response, err := client.GetLustreFileSystem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_lustre_file_storage.LustreFileSystemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("LustreFileStorageLustreFileSystem") {
		resource.AddTestSweepers("LustreFileStorageLustreFileSystem", &resource.Sweeper{
			Name:         "LustreFileStorageLustreFileSystem",
			Dependencies: acctest.DependencyGraph["lustreFileSystem"],
			F:            sweepLustreFileStorageLustreFileSystemResource,
		})
	}
}

func sweepLustreFileStorageLustreFileSystemResource(compartment string) error {
	lustreFileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).LustreFileStorageClient()
	lustreFileSystemIds, err := getLustreFileStorageLustreFileSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, lustreFileSystemId := range lustreFileSystemIds {
		if ok := acctest.SweeperDefaultResourceId[lustreFileSystemId]; !ok {
			deleteLustreFileSystemRequest := oci_lustre_file_storage.DeleteLustreFileSystemRequest{}

			deleteLustreFileSystemRequest.LustreFileSystemId = &lustreFileSystemId

			deleteLustreFileSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "lustre_file_storage")
			_, error := lustreFileStorageClient.DeleteLustreFileSystem(context.Background(), deleteLustreFileSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting LustreFileSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", lustreFileSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &lustreFileSystemId, LustreFileStorageLustreFileSystemSweepWaitCondition, time.Duration(3*time.Minute),
				LustreFileStorageLustreFileSystemSweepResponseFetchOperation, "lustre_file_storage", true)
		}
	}
	return nil
}

func getLustreFileStorageLustreFileSystemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LustreFileSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	lustreFileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).LustreFileStorageClient()

	listLustreFileSystemsRequest := oci_lustre_file_storage.ListLustreFileSystemsRequest{}
	listLustreFileSystemsRequest.CompartmentId = &compartmentId
	listLustreFileSystemsRequest.LifecycleState = oci_lustre_file_storage.LustreFileSystemLifecycleStateActive
	listLustreFileSystemsResponse, err := lustreFileStorageClient.ListLustreFileSystems(context.Background(), listLustreFileSystemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LustreFileSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, lustreFileSystem := range listLustreFileSystemsResponse.Items {
		id := *lustreFileSystem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LustreFileSystemId", id)
	}
	return resourceIds, nil
}

func LustreFileStorageLustreFileSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if lustreFileSystemResponse, ok := response.Response.(oci_lustre_file_storage.GetLustreFileSystemResponse); ok {
		return lustreFileSystemResponse.LifecycleState != oci_lustre_file_storage.LustreFileSystemLifecycleStateDeleted
	}
	return false
}

func LustreFileStorageLustreFileSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LustreFileStorageClient().GetLustreFileSystem(context.Background(), oci_lustre_file_storage.GetLustreFileSystemRequest{
		LustreFileSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
