// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	oci_opa "github.com/oracle/oci-go-sdk/v65/opa"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpaInstanceRequiredOnlyResource = OpaInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Required, acctest.Create, opaInstanceRepresentation)

	OpaInstanceResourceConfig = OpaInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Optional, acctest.Update, opaInstanceRepresentation)

	opaInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"opa_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opa_opa_instance.test_opa_instance.id}`},
	}

	opaInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_opa_opa_instance.test_opa_instance.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: opaInstanceDataSourceFilterRepresentation}}
	opaInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opa_opa_instance.test_opa_instance.id}`}},
	}

	opaInstanceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"shape_name":        acctest.Representation{RepType: acctest.Required, Create: `PRODUCTION`},
		"consumption_model": acctest.Representation{RepType: acctest.Required, Create: `UCM`},
		// "defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"idcs_at":               acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"is_breakglass_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metering_type":         acctest.Representation{RepType: acctest.Required, Create: `EXECUTION_PACK`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: opaIgnoreDefinedTagsDifferencesRepresentationAgain},
	}

	opaIgnoreDefinedTagsDifferencesRepresentationAgain = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	OpaInstanceResourceDependencies = "" /* DefinedTagsDependencies */
)

// issue-routing-tag: opa/default
func TestOpaOpaInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpaOpaInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opa_opa_instance.test_opa_instance"
	datasourceName := "data.oci_opa_opa_instances.test_opa_instances"
	singularDatasourceName := "data.oci_opa_opa_instance.test_opa_instance"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpaInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Optional, acctest.Create, opaInstanceRepresentation), "opa", "opaInstance", t)

	acctest.ResourceTest(t, testAccCheckOpaOpaInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr() + OpaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Required, acctest.Create, opaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OpaInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr() + OpaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Optional, acctest.Create, opaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"), //
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttr(resourceName, "is_breakglass_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "metering_type", "EXECUTION_PACK"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
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
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsAccessTokenVariableStr() + OpaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(opaInstanceRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: "ACTIVE"},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttr(resourceName, "is_breakglass_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "metering_type", "EXECUTION_PACK"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsAccessTokenVariableStr() + OpaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(opaInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttr(resourceName, "is_breakglass_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "metering_type", "EXECUTION_PACK"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr() + OpaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Optional, acctest.Update, opaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_at"),
				resource.TestCheckResourceAttr(resourceName, "is_breakglass_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "metering_type", "EXECUTION_PACK"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opa_opa_instances", "test_opa_instances", acctest.Optional, acctest.Update, opaInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr() + OpaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Optional, acctest.Update, opaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "opa_instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "opa_instance_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opa_opa_instance", "test_opa_instance", acctest.Required, acctest.Create, opaInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr() + OpaInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opa_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attachments.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumption_model", "UCM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_app_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_app_guid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_app_opc_service_instance_guid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_domain_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_breakglass_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metering_type", "EXECUTION_PACK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + OpaInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"idcs_at",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOpaOpaInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OpaInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opa_opa_instance" {
			noResourceFound = false
			request := oci_opa.GetOpaInstanceRequest{}

			tmp := rs.Primary.ID
			request.OpaInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opa")

			response, err := client.GetOpaInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opa.OpaInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpaOpaInstance") {
		resource.AddTestSweepers("OpaOpaInstance", &resource.Sweeper{
			Name:         "OpaOpaInstance",
			Dependencies: acctest.DependencyGraph["opaInstance"],
			F:            sweepOpaOpaInstanceResource,
		})
	}
}

func sweepOpaOpaInstanceResource(compartment string) error {
	opaInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).OpaInstanceClient()
	opaInstanceIds, err := getOpaInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, opaInstanceId := range opaInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[opaInstanceId]; !ok {
			deleteOpaInstanceRequest := oci_opa.DeleteOpaInstanceRequest{}

			deleteOpaInstanceRequest.OpaInstanceId = &opaInstanceId

			deleteOpaInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opa")
			_, error := opaInstanceClient.DeleteOpaInstance(context.Background(), deleteOpaInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting OpaInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", opaInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &opaInstanceId, opaInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				opaInstanceSweepResponseFetchOperation, "opa", true)
		}
	}
	return nil
}

func getOpaInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OpaInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	opaInstanceClient := acctest.GetTestClients(&schema.ResourceData{}).OpaInstanceClient()

	listOpaInstancesRequest := oci_opa.ListOpaInstancesRequest{}
	listOpaInstancesRequest.CompartmentId = &compartmentId
	listOpaInstancesRequest.LifecycleState = oci_opa.OpaInstanceLifecycleStateActive
	listOpaInstancesResponse, err := opaInstanceClient.ListOpaInstances(context.Background(), listOpaInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OpaInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, opaInstance := range listOpaInstancesResponse.Items {
		id := *opaInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OpaInstanceId", id)
	}
	return resourceIds, nil
}

func opaInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if opaInstanceResponse, ok := response.Response.(oci_opa.GetOpaInstanceResponse); ok {
		return opaInstanceResponse.LifecycleState != oci_opa.OpaInstanceLifecycleStateDeleted
	}
	return false
}

func opaInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OpaInstanceClient().GetOpaInstance(context.Background(), oci_opa.GetOpaInstanceRequest{
		OpaInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
