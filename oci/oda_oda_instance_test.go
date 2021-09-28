// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_oda "github.com/oracle/oci-go-sdk/v48/oda"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	OdaInstanceRequiredOnlyResource = OdaInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Required, Create, odaInstanceRepresentation)

	OdaInstanceResourceConfig = OdaInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Optional, Update, odaInstanceRepresentation)

	odaInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"oda_instance_id": Representation{repType: Required, create: `${oci_oda_oda_instance.test_oda_instance.id}`},
	}

	odaInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, odaInstanceDataSourceFilterRepresentation}}
	odaInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_oda_oda_instance.test_oda_instance.id}`}},
	}

	odaInstanceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"shape_name":     Representation{repType: Required, create: `DEVELOPMENT`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"state":          Representation{repType: Optional, create: `INACTIVE`, update: `ACTIVE`},
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

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_oda_oda_instance.test_oda_instance"
	datasourceName := "data.oci_oda_oda_instances.test_oda_instances"
	singularDatasourceName := "data.oci_oda_oda_instance.test_oda_instance"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+OdaInstanceResourceDependencies+
		generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Optional, Create, odaInstanceRepresentation), "oda", "odaInstance", t)

	ResourceTest(t, testAccCheckOdaOdaInstanceDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies +
				generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Required, Create, odaInstanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies +
				generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Optional, Create, odaInstanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OdaInstanceResourceDependencies +
				generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Optional, Update, odaInstanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OdaInstanceResourceDependencies +
				generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Optional, Update,
					representationCopyWithNewProperties(odaInstanceRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
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
				generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Optional, Update, odaInstanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				generateDataSourceFromRepresentationMap("oci_oda_oda_instances", "test_oda_instances", Optional, Update, odaInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + OdaInstanceResourceDependencies +
				generateResourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Optional, Update, odaInstanceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "oda_instances.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oda_instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "oda_instances.0.defined_tags.%", "1"),
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
				generateDataSourceFromRepresentationMap("oci_oda_oda_instance", "test_oda_instance", Required, Create, odaInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OdaInstanceResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oda_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connector_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
	client := testAccProvider.Meta().(*OracleClients).odaClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oda_oda_instance" {
			noResourceFound = false
			request := oci_oda.GetOdaInstanceRequest{}

			tmp := rs.Primary.ID
			request.OdaInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "oda")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("OdaOdaInstance") {
		resource.AddTestSweepers("OdaOdaInstance", &resource.Sweeper{
			Name:         "OdaOdaInstance",
			Dependencies: DependencyGraph["odaInstance"],
			F:            sweepOdaOdaInstanceResource,
		})
	}
}

func sweepOdaOdaInstanceResource(compartment string) error {
	odaClient := GetTestClients(&schema.ResourceData{}).odaClient()
	odaInstanceIds, err := getOdaInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, odaInstanceId := range odaInstanceIds {
		if ok := SweeperDefaultResourceId[odaInstanceId]; !ok {
			deleteOdaInstanceRequest := oci_oda.DeleteOdaInstanceRequest{}

			deleteOdaInstanceRequest.OdaInstanceId = &odaInstanceId

			deleteOdaInstanceRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "oda")
			_, error := odaClient.DeleteOdaInstance(context.Background(), deleteOdaInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting OdaInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", odaInstanceId, error)
				continue
			}
			waitTillCondition(testAccProvider, &odaInstanceId, odaInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				odaInstanceSweepResponseFetchOperation, "oda", true)
		}
	}
	return nil
}

func getOdaInstanceIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "OdaInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	odaClient := GetTestClients(&schema.ResourceData{}).odaClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "OdaInstanceId", id)
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

func odaInstanceSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.odaClient().GetOdaInstance(context.Background(), oci_oda.GetOdaInstanceRequest{
		OdaInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
