// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v28/common"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v28/osmanagement"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagedInstanceGroupRequiredOnlyResource = ManagedInstanceGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, managedInstanceGroupRepresentation)

	ManagedInstanceGroupResourceConfig = ManagedInstanceGroupResourceDependencies +
		generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Update, managedInstanceGroupRepresentation)

	managedInstanceGroupSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": Representation{repType: Required, create: `${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`},
	}

	managedGroupDisplayName                      = randomStringOrHttpReplayValue(10, charsetWithoutDigits, "displayName")
	managedGroupUpdateDisplayName                = randomStringOrHttpReplayValue(10, charsetWithoutDigits, "displayName2")
	managedInstanceGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: managedGroupDisplayName, update: managedGroupUpdateDisplayName},
		"os_family":      Representation{repType: Optional, create: `WINDOWS`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, managedInstanceGroupDataSourceFilterRepresentation}}
	managedInstanceGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_osmanagement_managed_instance_group.test_managed_instance_group.id}`}},
	}

	managedInstanceGroupRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: managedGroupDisplayName, update: managedGroupUpdateDisplayName},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"os_family":      Representation{repType: Optional, create: `WINDOWS`},
	}

	ManagedInstanceGroupResourceDependencies = DefinedTagsDependencies
)

func TestOsmanagementManagedInstanceGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceGroupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_osmanagement_managed_instance_group.test_managed_instance_group"
	datasourceName := "data.oci_osmanagement_managed_instance_groups.test_managed_instance_groups"
	singularDatasourceName := "data.oci_osmanagement_managed_instance_group.test_managed_instance_group"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOsmanagementManagedInstanceGroupDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, managedInstanceGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Create, managedInstanceGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "os_family", "WINDOWS"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ManagedInstanceGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Create,
						representationCopyWithNewProperties(managedInstanceGroupRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupDisplayName),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "os_family", "WINDOWS"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Update, managedInstanceGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", managedGroupUpdateDisplayName),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "os_family", "WINDOWS"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_groups", "test_managed_instance_groups", Optional, Update, managedInstanceGroupDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceGroupResourceDependencies +
					generateResourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Optional, Update, managedInstanceGroupRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", managedGroupUpdateDisplayName),
					resource.TestCheckResourceAttr(datasourceName, "os_family", "WINDOWS"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "managed_instance_groups.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance_group", "test_managed_instance_group", Required, Create, managedInstanceGroupSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceGroupResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}

func testAccCheckOsmanagementManagedInstanceGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).osManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_osmanagement_managed_instance_group" {
			noResourceFound = false
			request := oci_osmanagement.GetManagedInstanceGroupRequest{}

			tmp := rs.Primary.ID
			request.ManagedInstanceGroupId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "osmanagement")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("OsmanagementManagedInstanceGroup") {
		resource.AddTestSweepers("OsmanagementManagedInstanceGroup", &resource.Sweeper{
			Name:         "OsmanagementManagedInstanceGroup",
			Dependencies: DependencyGraph["managedInstanceGroup"],
			F:            sweepOsmanagementManagedInstanceGroupResource,
		})
	}
}

func sweepOsmanagementManagedInstanceGroupResource(compartment string) error {
	osManagementClient := GetTestClients(&schema.ResourceData{}).osManagementClient()
	managedInstanceGroupIds, err := getManagedInstanceGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, managedInstanceGroupId := range managedInstanceGroupIds {
		if ok := SweeperDefaultResourceId[managedInstanceGroupId]; !ok {
			deleteManagedInstanceGroupRequest := oci_osmanagement.DeleteManagedInstanceGroupRequest{}

			deleteManagedInstanceGroupRequest.ManagedInstanceGroupId = &managedInstanceGroupId

			deleteManagedInstanceGroupRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "osmanagement")
			_, error := osManagementClient.DeleteManagedInstanceGroup(context.Background(), deleteManagedInstanceGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagedInstanceGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", managedInstanceGroupId, error)
				continue
			}
			waitTillCondition(testAccProvider, &managedInstanceGroupId, managedInstanceGroupSweepWaitCondition, time.Duration(3*time.Minute),
				managedInstanceGroupSweepResponseFetchOperation, "osmanagement", true)
		}
	}
	return nil
}

func getManagedInstanceGroupIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ManagedInstanceGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	osManagementClient := GetTestClients(&schema.ResourceData{}).osManagementClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ManagedInstanceGroupId", id)
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

func managedInstanceGroupSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.osManagementClient().GetManagedInstanceGroup(context.Background(), oci_osmanagement.GetManagedInstanceGroupRequest{
		ManagedInstanceGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
