// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_oda "github.com/oracle/oci-go-sdk/v56/oda"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	OdaInstanceRequiredOnlyResource = OdaInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Required, acctest.Create, odaInstanceRepresentation)

	OdaInstanceResourceConfig = OdaInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Optional, acctest.Update, odaInstanceRepresentation)

	odaInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"oda_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_instance.test_oda_instance.id}`},
	}

	odaInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: odaInstanceDataSourceFilterRepresentation}}
	odaInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_oda_oda_instance.test_oda_instance.id}`}},
	}

	odaInstanceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":     acctest.Representation{RepType: acctest.Required, Create: `DEVELOPMENT`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}

	OdaInstanceResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: oda/default
func TestOdaOdaInstanceResource_basic(t *testing.T) {
	if httpreplay.ShouldRetryImmediately() {
		t.Skip("TestOdaOdaInstanceResource_basic test environment is not ready, skip this test for checkin test.")
	}

	httpreplay.SetScenario("TestOdaOdaInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_oda_oda_instance.test_oda_instance"
	datasourceName := "data.oci_oda_oda_instances.test_oda_instances"
	singularDatasourceName := "data.oci_oda_oda_instance.test_oda_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OdaInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Optional, acctest.Create, odaInstanceRepresentation), "oda", "odaInstance", t)

	acctest.ResourceTest(t, testAccCheckOdaOdaInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Required, acctest.Create, odaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Optional, acctest.Create, odaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Optional, acctest.Update, odaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OdaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(odaInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify switch back
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Optional, acctest.Update, odaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_instances", "test_oda_instances", acctest.Optional, acctest.Update, odaInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + OdaInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Optional, acctest.Update, odaInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "oda_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oda_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "oda_instances.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "oda_instances.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "oda_instances.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "oda_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oda_instances.0.shape_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "oda_instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "oda_instances.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "oda_instances.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", acctest.Required, acctest.Create, odaInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OdaInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oda_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connector_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "web_app_url"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceConfig,
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

func testAccCheckOdaOdaInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OdaClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oda_oda_instance" {
			noResourceFound = false
			request := oci_oda.GetOdaInstanceRequest{}

			tmp := rs.Primary.ID
			request.OdaInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")

			response, err := client.GetOdaInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_oda.OdaInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OdaOdaInstance") {
		resource.AddTestSweepers("OdaOdaInstance", &resource.Sweeper{
			Name:         "OdaOdaInstance",
			Dependencies: acctest.DependencyGraph["odaInstance"],
			F:            sweepOdaOdaInstanceResource,
		})
	}
}

func sweepOdaOdaInstanceResource(compartment string) error {
	odaClient := acctest.GetTestClients(&schema.ResourceData{}).OdaClient()
	odaInstanceIds, err := getOdaInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, odaInstanceId := range odaInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[odaInstanceId]; !ok {
			deleteOdaInstanceRequest := oci_oda.DeleteOdaInstanceRequest{}

			deleteOdaInstanceRequest.OdaInstanceId = &odaInstanceId

			deleteOdaInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")
			_, error := odaClient.DeleteOdaInstance(context.Background(), deleteOdaInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting OdaInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", odaInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &odaInstanceId, odaInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				odaInstanceSweepResponseFetchOperation, "oda", true)
		}
	}
	return nil
}

func getOdaInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OdaInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	odaClient := acctest.GetTestClients(&schema.ResourceData{}).OdaClient()

	listOdaInstancesRequest := oci_oda.ListOdaInstancesRequest{}
	listOdaInstancesRequest.CompartmentId = &compartmentId
	listOdaInstancesRequest.LifecycleState = oci_oda.ListOdaInstancesLifecycleStateActive
	listOdaInstancesResponse, err := odaClient.ListOdaInstances(context.Background(), listOdaInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OdaInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, odaInstance := range listOdaInstancesResponse.Items {
		id := *odaInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OdaInstanceId", id)
	}
	return resourceIds, nil
}

func odaInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if odaInstanceResponse, ok := response.Response.(oci_oda.GetOdaInstanceResponse); ok {
		return odaInstanceResponse.LifecycleState != oci_oda.OdaInstanceLifecycleStateDeleted
	}
	return false
}

func odaInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OdaClient().GetOdaInstance(context.Background(), oci_oda.GetOdaInstanceRequest{
		OdaInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
