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
	IotDigitalTwinModelRequiredOnlyResource = IotDigitalTwinModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Required, acctest.Create, IotDigitalTwinModelRepresentation)

	IotDigitalTwinModelResourceConfig = IotDigitalTwinModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Optional, acctest.Update, IotDigitalTwinModelRepresentation)

	IotDigitalTwinModelSingularDataSourceRepresentation = map[string]interface{}{
		"digital_twin_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.id}`},
	}

	IotDigitalTwinModelDataSourceRepresentation = map[string]interface{}{
		"iot_domain_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"spec_uri_starts_with": acctest.Representation{RepType: acctest.Optional, Create: `dtmi:com:oracle:example`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.id}`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: IotDigitalTwinModelDataSourceFilterRepresentation}}
	IotDigitalTwinModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_iot_digital_twin_model.test_digital_twin_model.id}`}},
	}

	ignoreDigitalTwinModelDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	specMap = `${jsonencode({
                             "@id": "dtmi:com:oracle:example:device;1",
                             "@type": "Interface",
                             "@context": "dtmi:dtdl:context;3",
                             "displayName": "IoT Device Model",
                             "description": "Represents a simple IoT device with temperature property.",
                             "contents": [
                               {
                                 "@type": "Property",
                                 "name": "temperature",
                                 "schema": "double",
                                 "displayName": "Temperature",
                                 "description": "The current temperature reading of the device."
                               }
                             ]
                           })}`

	IotDigitalTwinModelRepresentation = map[string]interface{}{
		"iot_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"spec":          acctest.Representation{RepType: acctest.Required, Create: specMap},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Protocol": "Mqtt"}, Update: map[string]string{"Protocol": "MQTT"}},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDigitalTwinModelDefinedTagsChangesRepresentation},
	}

	IotDigitalTwinModelResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: iot/default
func TestIotDigitalTwinModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotDigitalTwinModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	resourceName := "oci_iot_digital_twin_model.test_digital_twin_model"
	datasourceName := "data.oci_iot_digital_twin_models.test_digital_twin_models"
	singularDatasourceName := "data.oci_iot_digital_twin_model.test_digital_twin_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+iotDomainIdVariableStr+IotDigitalTwinModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Optional, acctest.Create, IotDigitalTwinModelRepresentation), "iot", "digitalTwinModel", t)

	acctest.ResourceTest(t, testAccCheckIotDigitalTwinModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + iotDomainIdVariableStr + IotDigitalTwinModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Required, acctest.Create, IotDigitalTwinModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "spec_uri"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + iotDomainIdVariableStr + IotDigitalTwinModelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + iotDomainIdVariableStr + IotDigitalTwinModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Optional, acctest.Create, IotDigitalTwinModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "spec_uri"),
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

		// verify updates to updatable parameters
		{
			Config: config + iotDomainIdVariableStr + IotDigitalTwinModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Optional, acctest.Update, IotDigitalTwinModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "spec_uri"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_models", "test_digital_twin_models", acctest.Optional, acctest.Update, IotDigitalTwinModelDataSourceRepresentation) +
				iotDomainIdVariableStr + IotDigitalTwinModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Optional, acctest.Update, IotDigitalTwinModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "spec_uri_starts_with", "dtmi:com:oracle:example"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "digital_twin_model_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "digital_twin_model_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Required, acctest.Create, IotDigitalTwinModelSingularDataSourceRepresentation) +
				iotDomainIdVariableStr + IotDigitalTwinModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_model_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "spec_uri"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + IotDigitalTwinModelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"spec",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIotDigitalTwinModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IotClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_iot_digital_twin_model" {
			noResourceFound = false
			request := oci_iot.GetDigitalTwinModelRequest{}

			tmp := rs.Primary.ID
			request.DigitalTwinModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")

			response, err := client.GetDigitalTwinModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_iot.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("IotDigitalTwinModel") {
		resource.AddTestSweepers("IotDigitalTwinModel", &resource.Sweeper{
			Name:         "IotDigitalTwinModel",
			Dependencies: acctest.DependencyGraph["digitalTwinModel"],
			F:            sweepIotDigitalTwinModelResource,
		})
	}
}

func sweepIotDigitalTwinModelResource(compartment string) error {
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()
	digitalTwinModelIds, err := getIotDigitalTwinModelIds(compartment)
	if err != nil {
		return err
	}
	for _, digitalTwinModelId := range digitalTwinModelIds {
		if ok := acctest.SweeperDefaultResourceId[digitalTwinModelId]; !ok {
			deleteDigitalTwinModelRequest := oci_iot.DeleteDigitalTwinModelRequest{}

			deleteDigitalTwinModelRequest.DigitalTwinModelId = &digitalTwinModelId

			deleteDigitalTwinModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")
			_, error := iotClient.DeleteDigitalTwinModel(context.Background(), deleteDigitalTwinModelRequest)
			if error != nil {
				fmt.Printf("Error deleting DigitalTwinModel %s %s, It is possible that the resource is already deleted. Please verify manually \n", digitalTwinModelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &digitalTwinModelId, IotDigitalTwinModelSweepWaitCondition, time.Duration(3*time.Minute),
				IotDigitalTwinModelSweepResponseFetchOperation, "iot", true)
		}
	}
	return nil
}

func getIotDigitalTwinModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DigitalTwinModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()

	listDigitalTwinModelsRequest := oci_iot.ListDigitalTwinModelsRequest{}
	//listDigitalTwinModelsRequest.CompartmentId = &compartmentId

	iotDomainIds, error := getIotIotDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting iotDomainId required for DigitalTwinModel resource requests \n")
	}
	for _, iotDomainId := range iotDomainIds {
		listDigitalTwinModelsRequest.IotDomainId = &iotDomainId

		listDigitalTwinModelsRequest.LifecycleState = oci_iot.ListDigitalTwinModelsLifecycleStateActive
		listDigitalTwinModelsResponse, err := iotClient.ListDigitalTwinModels(context.Background(), listDigitalTwinModelsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DigitalTwinModel list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, digitalTwinModel := range listDigitalTwinModelsResponse.Items {
			id := *digitalTwinModel.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DigitalTwinModelId", id)
		}

	}
	return resourceIds, nil
}

func IotDigitalTwinModelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if digitalTwinModelResponse, ok := response.Response.(oci_iot.GetDigitalTwinModelResponse); ok {
		return digitalTwinModelResponse.LifecycleState != oci_iot.LifecycleStateDeleted
	}
	return false
}

func IotDigitalTwinModelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IotClient().GetDigitalTwinModel(context.Background(), oci_iot.GetDigitalTwinModelRequest{
		DigitalTwinModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
