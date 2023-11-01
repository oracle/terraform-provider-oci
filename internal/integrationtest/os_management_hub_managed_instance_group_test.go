// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubManagedInstanceGroupRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation)

	OsManagementHubManagedInstanceGroupResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceGroupRepresentation)

	OsManagementHubManagedInstanceGroupSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
	}

	OsManagementHubManagedInstanceGroupDataSourceRepresentation = map[string]interface{}{
		"arch_type":                 acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}, Update: []string{`displayName2`}},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `display`},
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"os_family":                 acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"software_source_id":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubManagedInstanceGroupDataSourceFilterRepresentation}}
	OsManagementHubManagedInstanceGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`}},
	}

	OsManagementHubManagedInstanceGroupRepresentation = map[string]interface{}{
		"arch_type":           acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"os_family":           acctest.Representation{RepType: acctest.Required, Create: `ORACLE_LINUX_8`},
		"software_source_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}},
		"vendor_name":         acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubManagedInstanceGroupIgnoreDefinedTagsRepresentation},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OsManagementHubManagedInstanceGroupIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance_group.test_managed_instance_group"
	datasourceName := "data.oci_os_management_hub_managed_instance_groups.test_managed_instance_groups"
	singularDatasourceName := "data.oci_os_management_hub_managed_instance_group.test_managed_instance_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation), "osmanagementhub", "managedInstanceGroup", t)

	acctest.ResourceTest(t, testAccCheckOsManagementHubManagedInstanceGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "software_source_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "software_sources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "software_source_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "software_sources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_groups", "test_managed_instance_groups", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceGroupDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, OsManagementHubManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "display"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubManagedInstanceGroupResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pending_job_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_source_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_sources.0.description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_sources.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_sources.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_sources.0.software_source_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_modified"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vendor_name", "ORACLE"),
			),
		},
		// verify resource import
		{
			Config:            config + OsManagementHubManagedInstanceGroupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"software_source_ids",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOsManagementHubManagedInstanceGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagedInstanceGroupClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_managed_instance_group" {
			noResourceFound = false
			request := oci_os_management_hub.GetManagedInstanceGroupRequest{}

			tmp := rs.Primary.ID
			request.ManagedInstanceGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetManagedInstanceGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.ManagedInstanceGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OsManagementHubManagedInstanceGroup") {
		resource.AddTestSweepers("OsManagementHubManagedInstanceGroup", &resource.Sweeper{
			Name:         "OsManagementHubManagedInstanceGroup",
			Dependencies: acctest.DependencyGraph["managedInstanceGroup"],
			F:            sweepOsManagementHubManagedInstanceGroupResource,
		})
	}
}

func sweepOsManagementHubManagedInstanceGroupResource(compartment string) error {
	managedInstanceGroupClient := acctest.GetTestClients(&schema.ResourceData{}).ManagedInstanceGroupClient()
	managedInstanceGroupIds, err := getOsManagementHubManagedInstanceGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, managedInstanceGroupId := range managedInstanceGroupIds {
		if ok := acctest.SweeperDefaultResourceId[managedInstanceGroupId]; !ok {
			deleteManagedInstanceGroupRequest := oci_os_management_hub.DeleteManagedInstanceGroupRequest{}

			deleteManagedInstanceGroupRequest.ManagedInstanceGroupId = &managedInstanceGroupId

			deleteManagedInstanceGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := managedInstanceGroupClient.DeleteManagedInstanceGroup(context.Background(), deleteManagedInstanceGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagedInstanceGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", managedInstanceGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managedInstanceGroupId, OsManagementHubManagedInstanceGroupSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubManagedInstanceGroupSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubManagedInstanceGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagedInstanceGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managedInstanceGroupClient := acctest.GetTestClients(&schema.ResourceData{}).ManagedInstanceGroupClient()

	listManagedInstanceGroupsRequest := oci_os_management_hub.ListManagedInstanceGroupsRequest{}
	listManagedInstanceGroupsRequest.CompartmentId = &compartmentId
	listManagedInstanceGroupsRequest.LifecycleState = oci_os_management_hub.ManagedInstanceGroupLifecycleStateActive
	listManagedInstanceGroupsResponse, err := managedInstanceGroupClient.ListManagedInstanceGroups(context.Background(), listManagedInstanceGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagedInstanceGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managedInstanceGroup := range listManagedInstanceGroupsResponse.Items {
		id := *managedInstanceGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagedInstanceGroupId", id)
	}
	return resourceIds, nil
}

func OsManagementHubManagedInstanceGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managedInstanceGroupResponse, ok := response.Response.(oci_os_management_hub.GetManagedInstanceGroupResponse); ok {
		return managedInstanceGroupResponse.LifecycleState != oci_os_management_hub.ManagedInstanceGroupLifecycleStateDeleted
	}
	return false
}

func OsManagementHubManagedInstanceGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagedInstanceGroupClient().GetManagedInstanceGroup(context.Background(), oci_os_management_hub.GetManagedInstanceGroupRequest{
		ManagedInstanceGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
