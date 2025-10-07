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
	temperatureSpecMap = `${jsonencode(
            {
              "@id": "dtmi:example:TemperatureDevice;1",
              "@type": "Interface",
              "@context": "dtmi:dtdl:context;3",
              "displayName": "Temperature Device",
              "contents": [
                {
                  "@type": "Property",
                  "name": "temperature",
                  "schema": "double",
                  "description": "Temperature in Celsius."
                },
                {
                  "@type": "Relationship",
                  "name": "connectedHumidity",
                  "target": "dtmi:example:HumidityDevice;1",
                  "properties": [
                    {
                      "@type": "Property",
                      "name": "connectionStrength",
                      "schema": "integer",
                      "description": "Strength of the connection between devices (0–100)."
                    }
                  ]
                }
              ]
            })}`

	humiditySpecMap = `${jsonencode(
            {
              "@id": "dtmi:example:HumidityDevice;1",
              "@type": "Interface",
              "@context": "dtmi:dtdl:context;3",
              "displayName": "Humidity Device",
              "contents": [
                {
                  "@type": "Property",
                  "name": "humidity",
                  "schema": "double",
                  "description": "Humidity in percentage."
                }
              ]
            })}`

	ignoreDigitalTwinRelationshipDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	iotTemperatureModelRepresentation = map[string]interface{}{
		"iot_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"spec":          acctest.Representation{RepType: acctest.Required, Create: temperatureSpecMap},
	}

	iotHumidityModelRepresentation = map[string]interface{}{
		"iot_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"spec":          acctest.Representation{RepType: acctest.Required, Create: humiditySpecMap},
	}

	iotTemperatureAdapterRepresentation = map[string]interface{}{
		"iot_domain_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"digital_twin_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_model.test_digital_twin_temperature_model.id}`},
	}

	iotHumidityAdapterRepresentation = map[string]interface{}{
		"iot_domain_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"digital_twin_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_model.test_digital_twin_humidity_model.id}`},
	}

	iotTemperatureInstanceRepresentation = map[string]interface{}{
		"auth_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.auth_id}`},
		"iot_domain_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"digital_twin_adapter_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_adapter.test_digital_twin_temperature_adapter.id}`},
	}

	iotHumidityInstanceRepresentation = map[string]interface{}{
		"auth_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.auth_id}`},
		"iot_domain_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"digital_twin_adapter_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_adapter.test_digital_twin_humidity_adapter.id}`},
	}

	IotDigitalTwinRelationshipRequiredOnlyResource = IotDigitalTwinRelationshipResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Required, acctest.Create, IotDigitalTwinRelationshipRepresentation)

	IotDigitalTwinRelationshipResourceConfig = IotDigitalTwinRelationshipResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Optional, acctest.Update, IotDigitalTwinRelationshipRepresentation)

	IotDigitalTwinRelationshipSingularDataSourceRepresentation = map[string]interface{}{
		"digital_twin_relationship_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_relationship.test_digital_twin_relationship.id}`},
	}

	IotDigitalTwinRelationshipDataSourceRepresentation = map[string]interface{}{
		"iot_domain_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"content_path":                    acctest.Representation{RepType: acctest.Optional, Create: `connectedHumidity`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                              acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_relationship.test_digital_twin_relationship.id}`},
		"source_digital_twin_instance_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_instance.test_digital_twin_temperature_instance.id}`},
		"state":                           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_digital_twin_instance_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_instance.test_digital_twin_humidity_instance.id}`},
		"filter":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: IotDigitalTwinRelationshipDataSourceFilterRepresentation}}
	IotDigitalTwinRelationshipDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_iot_digital_twin_relationship.test_digital_twin_relationship.id}`}},
	}

	contentMap = `${jsonencode({
        "connectionStrength" = 98
     })}`

	IotDigitalTwinRelationshipRepresentation = map[string]interface{}{
		"content_path":                    acctest.Representation{RepType: acctest.Required, Create: `connectedHumidity`},
		"iot_domain_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"source_digital_twin_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_instance.test_digital_twin_temperature_instance.id}`},
		"target_digital_twin_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_instance.test_digital_twin_humidity_instance.id}`},
		"content":                         acctest.Representation{RepType: acctest.Required, Create: contentMap, Update: contentMap},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Protocol": "Mqtt"}, Update: map[string]string{"Protocol": "MQTT"}},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDigitalTwinRelationshipDefinedTagsChangesRepresentation},
	}

	IotDigitalTwinRelationshipResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_temperature_model", acctest.Required, acctest.Create, iotTemperatureModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_humidity_model", acctest.Required, acctest.Create, iotHumidityModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_temperature_adapter", acctest.Required, acctest.Create, iotTemperatureAdapterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_humidity_adapter", acctest.Required, acctest.Create, iotHumidityAdapterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_temperature_instance", acctest.Required, acctest.Create, iotTemperatureInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_humidity_instance", acctest.Required, acctest.Create, iotHumidityInstanceRepresentation)
)

// issue-routing-tag: iot/default
func TestIotDigitalTwinRelationshipResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotDigitalTwinRelationshipResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	authId := utils.GetEnvSettingWithBlankDefault("auth_ocid")
	authIdVariableStr := fmt.Sprintf("variable \"auth_id\" { default = \"%s\" }\n", authId)

	resourceName := "oci_iot_digital_twin_relationship.test_digital_twin_relationship"
	datasourceName := "data.oci_iot_digital_twin_relationships.test_digital_twin_relationships"
	singularDatasourceName := "data.oci_iot_digital_twin_relationship.test_digital_twin_relationship"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+iotDomainIdVariableStr+authIdVariableStr+compartmentIdVariableStr+IotDigitalTwinRelationshipResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Optional, acctest.Create, IotDigitalTwinRelationshipRepresentation), "iot", "digitalTwinRelationship", t)

	acctest.ResourceTest(t, testAccCheckIotDigitalTwinRelationshipDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + authIdVariableStr + IotDigitalTwinRelationshipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Required, acctest.Create, IotDigitalTwinRelationshipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(resourceName, "content_path", "connectedHumidity"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_digital_twin_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_digital_twin_instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + authIdVariableStr + IotDigitalTwinRelationshipResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + authIdVariableStr + IotDigitalTwinRelationshipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Optional, acctest.Create, IotDigitalTwinRelationshipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(resourceName, "content_path", "connectedHumidity"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_digital_twin_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_digital_twin_instance_id"),
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
			Config: config + compartmentIdVariableStr + iotDomainIdVariableStr + authIdVariableStr + IotDigitalTwinRelationshipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Optional, acctest.Update, IotDigitalTwinRelationshipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(resourceName, "content_path", "connectedHumidity"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_digital_twin_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_digital_twin_instance_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_relationships", "test_digital_twin_relationships", acctest.Optional, acctest.Update, IotDigitalTwinRelationshipDataSourceRepresentation) +
				compartmentIdVariableStr + iotDomainIdVariableStr + authIdVariableStr + IotDigitalTwinRelationshipResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Optional, acctest.Update, IotDigitalTwinRelationshipRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "content_path", "connectedHumidity"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_digital_twin_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_digital_twin_instance_id"),

				resource.TestCheckResourceAttr(datasourceName, "digital_twin_relationship_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "digital_twin_relationship_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_relationship", "test_digital_twin_relationship", acctest.Required, acctest.Create, IotDigitalTwinRelationshipSingularDataSourceRepresentation) +
				compartmentIdVariableStr + iotDomainIdVariableStr + authIdVariableStr + IotDigitalTwinRelationshipResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_relationship_id"),

				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_path", "connectedHumidity"),
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
			Config:                  config + IotDigitalTwinRelationshipRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIotDigitalTwinRelationshipDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IotClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_iot_digital_twin_relationship" {
			noResourceFound = false
			request := oci_iot.GetDigitalTwinRelationshipRequest{}

			tmp := rs.Primary.ID
			request.DigitalTwinRelationshipId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")

			response, err := client.GetDigitalTwinRelationship(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IotDigitalTwinRelationship") {
		resource.AddTestSweepers("IotDigitalTwinRelationship", &resource.Sweeper{
			Name:         "IotDigitalTwinRelationship",
			Dependencies: acctest.DependencyGraph["digitalTwinRelationship"],
			F:            sweepIotDigitalTwinRelationshipResource,
		})
	}
}

func sweepIotDigitalTwinRelationshipResource(compartment string) error {
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()
	digitalTwinRelationshipIds, err := getIotDigitalTwinRelationshipIds(compartment)
	if err != nil {
		return err
	}
	for _, digitalTwinRelationshipId := range digitalTwinRelationshipIds {
		if ok := acctest.SweeperDefaultResourceId[digitalTwinRelationshipId]; !ok {
			deleteDigitalTwinRelationshipRequest := oci_iot.DeleteDigitalTwinRelationshipRequest{}

			deleteDigitalTwinRelationshipRequest.DigitalTwinRelationshipId = &digitalTwinRelationshipId

			deleteDigitalTwinRelationshipRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")
			_, error := iotClient.DeleteDigitalTwinRelationship(context.Background(), deleteDigitalTwinRelationshipRequest)
			if error != nil {
				fmt.Printf("Error deleting DigitalTwinRelationship %s %s, It is possible that the resource is already deleted. Please verify manually \n", digitalTwinRelationshipId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &digitalTwinRelationshipId, IotDigitalTwinRelationshipSweepWaitCondition, time.Duration(3*time.Minute),
				IotDigitalTwinRelationshipSweepResponseFetchOperation, "iot", true)
		}
	}
	return nil
}

func getIotDigitalTwinRelationshipIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DigitalTwinRelationshipId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()

	listDigitalTwinRelationshipsRequest := oci_iot.ListDigitalTwinRelationshipsRequest{}
	//listDigitalTwinRelationshipsRequest.CompartmentId = &compartmentId

	iotDomainIds, error := getIotIotDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting iotDomainId required for DigitalTwinRelationship resource requests \n")
	}
	for _, iotDomainId := range iotDomainIds {
		listDigitalTwinRelationshipsRequest.IotDomainId = &iotDomainId

		listDigitalTwinRelationshipsRequest.LifecycleState = oci_iot.ListDigitalTwinRelationshipsLifecycleStateActive
		listDigitalTwinRelationshipsResponse, err := iotClient.ListDigitalTwinRelationships(context.Background(), listDigitalTwinRelationshipsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DigitalTwinRelationship list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, digitalTwinRelationship := range listDigitalTwinRelationshipsResponse.Items {
			id := *digitalTwinRelationship.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DigitalTwinRelationshipId", id)
		}

	}
	return resourceIds, nil
}

func IotDigitalTwinRelationshipSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if digitalTwinRelationshipResponse, ok := response.Response.(oci_iot.GetDigitalTwinRelationshipResponse); ok {
		return digitalTwinRelationshipResponse.LifecycleState != oci_iot.LifecycleStateDeleted
	}
	return false
}

func IotDigitalTwinRelationshipSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IotClient().GetDigitalTwinRelationship(context.Background(), oci_iot.GetDigitalTwinRelationshipRequest{
		DigitalTwinRelationshipId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
