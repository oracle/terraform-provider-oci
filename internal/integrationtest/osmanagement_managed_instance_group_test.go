// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v56/osmanagement"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedInstanceGroupRequiredOnlyResource = ManagedInstanceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, managedInstanceGroupRepresentation)

	ManagedInstanceGroupResourceConfig = ManagedInstanceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, managedInstanceGroupRepresentation)

	managedInstanceGroupSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`},
	}

	managedGroupDisplayName                      = utils.RandomStringOrHttpReplayValue(10, utils.CharsetWithoutDigits, "displayName")
	managedGroupUpdateDisplayName                = utils.RandomStringOrHttpReplayValue(10, utils.CharsetWithoutDigits, "displayName2")
	managedInstanceGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: managedGroupDisplayName, Update: managedGroupUpdateDisplayName},
		"os_family":      acctest.Representation{RepType: acctest.Optional, Create: `WINDOWS`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: managedInstanceGroupDataSourceFilterRepresentation}}
	managedInstanceGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`}},
	}

	managedInstanceGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: managedGroupDisplayName, Update: managedGroupUpdateDisplayName},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"os_family":      acctest.Representation{RepType: acctest.Optional, Create: `WINDOWS`},
	}

	ManagedInstanceGroupResourceDependencies = DefinedTagsDependencies
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagedInstanceGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create, managedInstanceGroupRepresentation), "osmanagement", "managedInstanceGroup", t)

	acctest.ResourceTest(t, testAccCheckOsmanagementManagedInstanceGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, managedInstanceGroupRepresentation),
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
			Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create, managedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "WINDOWS"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(managedInstanceGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "WINDOWS"),

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
			Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, managedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "WINDOWS"),

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
				compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Update, managedInstanceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "WINDOWS"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_groups.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.os_family", "WINDOWS"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_groups.0.state"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, managedInstanceGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedInstanceGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", managedGroupUpdateDisplayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "managed_instances.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family", "WINDOWS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
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
	managedInstanceGroupIds, err := getManagedInstanceGroupIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &managedInstanceGroupId, managedInstanceGroupSweepWaitCondition, time.Duration(3*time.Minute),
				managedInstanceGroupSweepResponseFetchOperation, "osmanagement", true)
		}
	}
	return nil
}

func getManagedInstanceGroupIds(compartment string) ([]string, error) {
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

func managedInstanceGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managedInstanceGroupResponse, ok := response.Response.(oci_osmanagement.GetManagedInstanceGroupResponse); ok {
		return managedInstanceGroupResponse.LifecycleState != oci_osmanagement.LifecycleStatesDeleted
	}
	return false
}

func managedInstanceGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OsManagementClient().GetManagedInstanceGroup(context.Background(), oci_osmanagement.GetManagedInstanceGroupRequest{
		ManagedInstanceGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
