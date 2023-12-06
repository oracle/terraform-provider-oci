// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OsmanagementManagedInstanceGroupRequiredOnlyResource = OsmanagementManagedInstanceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsmanagementManagedInstanceGroupRepresentation)

	OsmanagementManagedInstanceGroupResourceConfig = OsmanagementManagedInstanceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, OsmanagementManagedInstanceGroupRepresentation)

	OsmanagementOsmanagementManagedInstanceGroupSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`},
	}

	managedGroupDisplayName                      = utils.RandomStringOrHttpReplayValue(10, utils.CharsetWithoutDigits, "displayName")
	managedGroupUpdateDisplayName                = utils.RandomStringOrHttpReplayValue(10, utils.CharsetWithoutDigits, "displayName2")
	managedInstanceGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: managedGroupDisplayName, Update: managedGroupUpdateDisplayName},
		"os_family":      acctest.Representation{RepType: acctest.Optional, Create: `LINUX`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OsmanagementManagedInstanceGroupDataSourceFilterRepresentation}}
	OsmanagementManagedInstanceGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`}},
	}

	OsmanagementManagedInstanceGroupRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: managedGroupDisplayName, Update: managedGroupUpdateDisplayName},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"os_family":            acctest.Representation{RepType: acctest.Optional, Create: `LINUX`},
		"managed_instance_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{managedInstanceOCID}},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: OsmanagementIgnoreDefinedTagsChangesRepresentation},
	}

	OsmanagementManagedInstanceGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: osmanagement/default
func TestOsmanagementManagedInstanceGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_osmanagement_managed_instance_group.test_managed_instance_group"
	datasourceName := "data.oci_osmanagement_managed_instance_groups.test_managed_instance_groups"
	singularDatasourceName := "data.oci_osmanagement_managed_instance_group.test_managed_instance_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsmanagementManagedInstanceGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create, OsmanagementManagedInstanceGroupRepresentation), "osmanagement", "managedInstanceGroup", t)

	acctest.ResourceTest(t, testAccCheckOsmanagementManagedInstanceGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsmanagementManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create, OsmanagementManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OsmanagementManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OsmanagementManagedInstanceGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),

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
			Config: config + compartmentIdVariableStr + OsmanagementManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, OsmanagementManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "LINUX"),
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_ids.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_groups", "test_managed_instance_groups", acctest.Optional, acctest.Update, managedInstanceGroupDataSourceRepresentation) +
				compartmentIdVariableStr + OsmanagementManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, OsmanagementManagedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "LINUX"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_groups.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.os_family", "LINUX"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_groups.0.state"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsmanagementOsmanagementManagedInstanceGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsmanagementManagedInstanceGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_instances.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family", "LINUX"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:                  config + OsmanagementManagedInstanceGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOsmanagementManagedInstanceGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_osmanagement_managed_instance_group" {
			noResourceFound = false
			request := oci_osmanagement.GetManagedInstanceGroupRequest{}

			tmp := rs.Primary.ID
			request.ManagedInstanceGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "osmanagement")

			response, err := client.GetManagedInstanceGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_osmanagement.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("OsmanagementManagedInstanceGroup") {
		resource.AddTestSweepers("OsmanagementManagedInstanceGroup", &resource.Sweeper{
			Name:         "OsmanagementManagedInstanceGroup",
			Dependencies: acctest.DependencyGraph["managedInstanceGroup"],
			F:            sweepOsmanagementManagedInstanceGroupResource,
		})
	}
}

func sweepOsmanagementManagedInstanceGroupResource(compartment string) error {
	osManagementClient := acctest.GetTestClients(&schema.ResourceData{}).OsManagementClient()
	managedInstanceGroupIds, err := getOsmanagementManagedInstanceGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, managedInstanceGroupId := range managedInstanceGroupIds {
		if ok := acctest.SweeperDefaultResourceId[managedInstanceGroupId]; !ok {
			deleteManagedInstanceGroupRequest := oci_osmanagement.DeleteManagedInstanceGroupRequest{}

			deleteManagedInstanceGroupRequest.ManagedInstanceGroupId = &managedInstanceGroupId

			deleteManagedInstanceGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "osmanagement")
			_, error := osManagementClient.DeleteManagedInstanceGroup(context.Background(), deleteManagedInstanceGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagedInstanceGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", managedInstanceGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managedInstanceGroupId, OsmanagementManagedInstanceGroupSweepWaitCondition, time.Duration(3*time.Minute),
				OsmanagementManagedInstanceGroupSweepResponseFetchOperation, "osmanagement", true)
		}
	}
	return nil
}

func getOsmanagementManagedInstanceGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagedInstanceGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	osManagementClient := acctest.GetTestClients(&schema.ResourceData{}).OsManagementClient()

	listManagedInstanceGroupsRequest := oci_osmanagement.ListManagedInstanceGroupsRequest{}
	listManagedInstanceGroupsRequest.CompartmentId = &compartmentId
	listManagedInstanceGroupsRequest.LifecycleState = oci_osmanagement.ListManagedInstanceGroupsLifecycleStateActive
	listManagedInstanceGroupsResponse, err := osManagementClient.ListManagedInstanceGroups(context.Background(), listManagedInstanceGroupsRequest)

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

func OsmanagementManagedInstanceGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managedInstanceGroupResponse, ok := response.Response.(oci_osmanagement.GetManagedInstanceGroupResponse); ok {
		return managedInstanceGroupResponse.LifecycleState != oci_osmanagement.LifecycleStatesDeleted
	}
	return false
}

func OsmanagementManagedInstanceGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OsManagementClient().GetManagedInstanceGroup(context.Background(), oci_osmanagement.GetManagedInstanceGroupRequest{
		ManagedInstanceGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
