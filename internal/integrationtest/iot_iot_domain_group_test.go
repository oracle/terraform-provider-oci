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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IotIotDomainGroupRequiredOnlyResource = IotIotDomainGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Required, acctest.Create, IotIotDomainGroupRepresentation)

	IotIotDomainGroupResourceConfig = IotIotDomainGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Optional, acctest.Update, IotIotDomainGroupRepresentation)

	IotIotDomainGroupSingularDataSourceRepresentation = map[string]interface{}{
		"iot_domain_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_iot_domain_group.test_iot_domain_group.id}`},
	}

	ignoreIotDomainGroupDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	IotIotDomainGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"iot_domain_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_iot_domain_group.test_iot_domain_group.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: IotIotDomainGroupDataSourceFilterRepresentation}}
	IotIotDomainGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_iot_iot_domain_group.test_iot_domain_group.id}`}},
	}

	IotIotDomainGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Protocol": "Mqtt"}, Update: map[string]string{"Protocol": "MQTT"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreIotDomainGroupDefinedTagsChangesRepresentation},
	}

	IotIotDomainGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: iot/default
func TestIotIotDomainGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotIotDomainGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_iot_iot_domain_group.test_iot_domain_group"
	datasourceName := "data.oci_iot_iot_domain_groups.test_iot_domain_groups"
	singularDatasourceName := "data.oci_iot_iot_domain_group.test_iot_domain_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IotIotDomainGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Optional, acctest.Create, IotIotDomainGroupRepresentation), "iot", "iotDomainGroup", t)

	acctest.ResourceTest(t, testAccCheckIotIotDomainGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IotIotDomainGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Required, acctest.Create, IotIotDomainGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IotIotDomainGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IotIotDomainGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Optional, acctest.Create, IotIotDomainGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + IotIotDomainGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IotIotDomainGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + IotIotDomainGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Optional, acctest.Update, IotIotDomainGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_iot_domain_groups", "test_iot_domain_groups", acctest.Optional, acctest.Update, IotIotDomainGroupDataSourceRepresentation) +
				compartmentIdVariableStr + IotIotDomainGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Optional, acctest.Update, IotIotDomainGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "iot_domain_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "iot_domain_group_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "iot_domain_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_iot_domain_group", "test_iot_domain_group", acctest.Required, acctest.Create, IotIotDomainGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IotIotDomainGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "iot_domain_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_host"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + IotIotDomainGroupRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIotIotDomainGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IotClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_iot_iot_domain_group" {
			noResourceFound = false
			request := oci_iot.GetIotDomainGroupRequest{}

			tmp := rs.Primary.ID
			request.IotDomainGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")

			response, err := client.GetIotDomainGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_iot.IotDomainGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("IotIotDomainGroup") {
		resource.AddTestSweepers("IotIotDomainGroup", &resource.Sweeper{
			Name:         "IotIotDomainGroup",
			Dependencies: acctest.DependencyGraph["iotDomainGroup"],
			F:            sweepIotIotDomainGroupResource,
		})
	}
}

func sweepIotIotDomainGroupResource(compartment string) error {
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()
	iotDomainGroupIds, err := getIotIotDomainGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, iotDomainGroupId := range iotDomainGroupIds {
		if ok := acctest.SweeperDefaultResourceId[iotDomainGroupId]; !ok {
			deleteIotDomainGroupRequest := oci_iot.DeleteIotDomainGroupRequest{}

			deleteIotDomainGroupRequest.IotDomainGroupId = &iotDomainGroupId

			deleteIotDomainGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")
			_, error := iotClient.DeleteIotDomainGroup(context.Background(), deleteIotDomainGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting IotDomainGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", iotDomainGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &iotDomainGroupId, IotIotDomainGroupSweepWaitCondition, time.Duration(3*time.Minute),
				IotIotDomainGroupSweepResponseFetchOperation, "iot", true)
		}
	}
	return nil
}

func getIotIotDomainGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IotDomainGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()

	listIotDomainGroupsRequest := oci_iot.ListIotDomainGroupsRequest{}
	listIotDomainGroupsRequest.CompartmentId = &compartmentId
	listIotDomainGroupsRequest.LifecycleState = oci_iot.IotDomainGroupLifecycleStateActive
	listIotDomainGroupsResponse, err := iotClient.ListIotDomainGroups(context.Background(), listIotDomainGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IotDomainGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, iotDomainGroup := range listIotDomainGroupsResponse.Items {
		id := *iotDomainGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IotDomainGroupId", id)
	}
	return resourceIds, nil
}

func IotIotDomainGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if iotDomainGroupResponse, ok := response.Response.(oci_iot.GetIotDomainGroupResponse); ok {
		return iotDomainGroupResponse.LifecycleState != oci_iot.IotDomainGroupLifecycleStateDeleted
	}
	return false
}

func IotIotDomainGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IotClient().GetIotDomainGroup(context.Background(), oci_iot.GetIotDomainGroupRequest{
		IotDomainGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
