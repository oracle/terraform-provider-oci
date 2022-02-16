// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v58/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	opctlName = `opctl-tf-` + utils.RandomString(5, utils.CharsetWithoutDigits)

	approverGroupOCID = getGroupOCID()

	OperatorControlRequiredOnlyResource = OperatorControlResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Required, acctest.Create, operatorControlRepresentation)

	OperatorControlResourceConfig = OperatorControlResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Optional, acctest.Update, operatorControlRepresentation)

	operatorControlSingularDataSourceRepresentation = map[string]interface{}{
		"operator_control_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_operator_access_control_operator_control.test_operator_control.id}`},
	}

	operatorControlDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"resource_type":  acctest.Representation{RepType: acctest.Required, Create: `EXADATAINFRASTRUCTURE`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `CREATED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: operatorControlDataSourceFilterRepresentation}}
	operatorControlDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_operator_access_control_operator_control.test_operator_control.id}`}},
	}

	operatorControlRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operator_control_name":       acctest.Representation{RepType: acctest.Required, Create: opctlName},
		"approver_groups_list":        acctest.Representation{RepType: acctest.Required, Create: []string{approverGroupOCID}, Update: []string{approverGroupOCID}},
		"approvers_list":              acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                 acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"email_id_list":               acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"service": "opctl"}, Update: map[string]string{"service": "opctl_2"}},
		"is_fully_pre_approved":       acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"pre_approved_op_action_list": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"resource_type":               acctest.Representation{RepType: acctest.Required, Create: `EXADATAINFRASTRUCTURE`},
		"system_message":              acctest.Representation{RepType: acctest.Optional, Create: `systemMessage`, Update: `systemMessage2`},
	}

	OperatorControlResourceDependencies = DefinedTagsDependencies
)

func getGroupOCID() string {
	// get the admin group ocid from identity service
	provider := common.DefaultConfigProvider()
	c, _ := identity.NewIdentityClientWithConfigurationProvider(provider)
	// override only for r1 region
	regionStr := utils.GetEnvSettingWithBlankDefault("region")
	if regionStr == "r1" {
		c.Host = "https://identity.r1.oracleiaas.com"
	}

	// The OCID of the tenancy containing the compartment.
	tenancyID, _ := provider.TenancyOCID()
	request := identity.ListGroupsRequest{
		CompartmentId: common.String(tenancyID),
	}
	r, _ := c.ListGroups(context.Background(), request)
	// find the admin group and get the OCID
	for _, g := range r.Items {
		if *g.Name == "Administrators" {
			return *g.Id
		}
	}
	return ""
}

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlOperatorControlResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOperatorAccessControlOperatorControlResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_operator_access_control_operator_control.test_operator_control"
	datasourceName := "data.oci_operator_access_control_operator_controls.test_operator_controls"
	singularDatasourceName := "data.oci_operator_access_control_operator_control.test_operator_control"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OperatorControlResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Optional, acctest.Create, operatorControlRepresentation), "operatoraccesscontrol", "operatorControl", t)

	acctest.ResourceTest(t, testAccCheckOperatorAccessControlOperatorControlDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OperatorControlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Required, acctest.Create, operatorControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approver_groups_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_fully_pre_approved", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_name"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OperatorControlResourceDependencies,
		},
		// verify Create with optionals
		{
			ExpectNonEmptyPlan: true,
			Config: config + compartmentIdVariableStr + OperatorControlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Optional, acctest.Create, operatorControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approver_groups_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approvers_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "email_id_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_fully_pre_approved", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_name"),
				resource.TestCheckResourceAttr(resourceName, "pre_approved_op_action_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(resourceName, "system_message", "systemMessage"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OperatorControlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(operatorControlRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approver_groups_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approvers_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "email_id_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_fully_pre_approved", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_name"),
				resource.TestCheckResourceAttr(resourceName, "pre_approved_op_action_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(resourceName, "system_message", "systemMessage"),

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
			Config: config + compartmentIdVariableStr + OperatorControlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Optional, acctest.Update, operatorControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approver_groups_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approvers_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "email_id_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_fully_pre_approved", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_name"),
				resource.TestCheckResourceAttr(resourceName, "pre_approved_op_action_list.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(resourceName, "system_message", "systemMessage2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_controls", "test_operator_controls", acctest.Optional, acctest.Update, operatorControlDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorControlResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Optional, acctest.Update, operatorControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "CREATED"),

				resource.TestCheckResourceAttr(datasourceName, "operator_control_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "operator_control_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Required, acctest.Create, operatorControlSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorControlResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operator_control_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "approval_required_op_action_list.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approver_groups_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approvers_list.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_id_list.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_fully_pre_approved", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_modified_info"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pre_approved_op_action_list.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_message", "systemMessage2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_creation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_modification"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + OperatorControlResourceConfig,
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

func testAccCheckOperatorAccessControlOperatorControlDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperatorControlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_operator_access_control_operator_control" {
			noResourceFound = false
			request := oci_operator_access_control.GetOperatorControlRequest{}

			tmp := rs.Primary.ID
			request.OperatorControlId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "operator_access_control")

			response, err := client.GetOperatorControl(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_operator_access_control.OperatorControlLifecycleStatesUnassigned): true, string(oci_operator_access_control.OperatorControlLifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
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
	if !acctest.InSweeperExcludeList("OperatorAccessControlOperatorControl") {
		resource.AddTestSweepers("OperatorAccessControlOperatorControl", &resource.Sweeper{
			Name:         "OperatorAccessControlOperatorControl",
			Dependencies: acctest.DependencyGraph["operatorControl"],
			F:            sweepOperatorAccessControlOperatorControlResource,
		})
	}
}

func sweepOperatorAccessControlOperatorControlResource(compartment string) error {
	operatorControlClient := acctest.GetTestClients(&schema.ResourceData{}).OperatorControlClient()
	operatorControlIds, err := getOperatorControlIds(compartment)
	if err != nil {
		return err
	}
	for _, operatorControlId := range operatorControlIds {
		if ok := acctest.SweeperDefaultResourceId[operatorControlId]; !ok {
			deleteOperatorControlRequest := oci_operator_access_control.DeleteOperatorControlRequest{}

			deleteOperatorControlRequest.OperatorControlId = &operatorControlId

			deleteOperatorControlRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "operator_access_control")
			_, error := operatorControlClient.DeleteOperatorControl(context.Background(), deleteOperatorControlRequest)
			if error != nil {
				fmt.Printf("Error deleting OperatorControl %s %s, It is possible that the resource is already deleted. Please verify manually \n", operatorControlId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &operatorControlId, operatorControlSweepWaitCondition, time.Duration(3*time.Minute),
				operatorControlSweepResponseFetchOperation, "operator_access_control", true)
		}
	}
	return nil
}

func getOperatorControlIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OperatorControlId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operatorControlClient := acctest.GetTestClients(&schema.ResourceData{}).OperatorControlClient()

	listOperatorControlsRequest := oci_operator_access_control.ListOperatorControlsRequest{}
	listOperatorControlsRequest.CompartmentId = &compartmentId
	listOperatorControlsRequest.LifecycleState = oci_operator_access_control.ListOperatorControlsLifecycleStateCreated
	listOperatorControlsResponse, err := operatorControlClient.ListOperatorControls(context.Background(), listOperatorControlsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OperatorControl list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, operatorControl := range listOperatorControlsResponse.Items {
		id := *operatorControl.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OperatorControlId", id)
	}
	return resourceIds, nil
}

func operatorControlSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if operatorControlResponse, ok := response.Response.(oci_operator_access_control.GetOperatorControlResponse); ok {
		return operatorControlResponse.LifecycleState != oci_operator_access_control.OperatorControlLifecycleStatesDeleted
	}
	return false
}

func operatorControlSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperatorControlClient().GetOperatorControl(context.Background(), oci_operator_access_control.GetOperatorControlRequest{
		OperatorControlId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
