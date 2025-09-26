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
	IotDigitalTwinAdapterRequiredOnlyResource = IotDigitalTwinAdapterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Required, acctest.Create, IotDigitalTwinAdapterRepresentation)

	IotDigitalTwinAdapterResourceConfig = IotDigitalTwinAdapterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Optional, acctest.Update, IotDigitalTwinAdapterRepresentation)

	IotDigitalTwinAdapterSingularDataSourceRepresentation = map[string]interface{}{
		"digital_twin_adapter_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_adapter.test_digital_twin_adapter.id}`},
	}

	ignoreDigitalTwinAdapterDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	IotDigitalTwinAdapterDataSourceRepresentation = map[string]interface{}{
		"iot_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"digital_twin_model_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.id}`},
		"digital_twin_model_spec_uri": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.spec_uri}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_adapter.test_digital_twin_adapter.id}`},
		"state":                       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: IotDigitalTwinAdapterDataSourceFilterRepresentation}}
	IotDigitalTwinAdapterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_iot_digital_twin_adapter.test_digital_twin_adapter.id}`}},
	}

	IotDigitalTwinAdapterRepresentation = map[string]interface{}{
		"iot_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"digital_twin_model_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.id}`},
		"digital_twin_model_spec_uri": acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.spec_uri}`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Protocol": "Mqtt"}, Update: map[string]string{"Protocol": "MQTT"}},
		"inbound_envelope":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: IotDigitalTwinAdapterInboundEnvelopeRepresentation},
		"inbound_routes":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IotDigitalTwinAdapterInboundRoutesRepresentation},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDigitalTwinAdapterDefinedTagsChangesRepresentation},
	}

	IotDigitalTwinAdapterInboundEnvelopeRepresentation = map[string]interface{}{
		"reference_endpoint": acctest.Representation{RepType: acctest.Required, Create: `/`, Update: `/test`},
		"envelope_mapping":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: IotDigitalTwinAdapterInboundEnvelopeEnvelopeMappingRepresentation},
		"reference_payload":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IotDigitalTwinAdapterInboundEnvelopeReferencePayloadRepresentation},
	}

	IotDigitalTwinAdapterInboundRoutesRepresentation = map[string]interface{}{
		"condition":         acctest.Representation{RepType: acctest.Required, Create: `*`, Update: `*`},
		"description":       acctest.Representation{RepType: acctest.Required, Create: `Default condition`, Update: `Default condition 1`},
		"payload_mapping":   acctest.Representation{RepType: acctest.Required, Create: map[string]string{"$.temperature": "$.temperature"}, Update: map[string]string{"$.temperature": "$.temperature"}},
		"reference_payload": acctest.RepresentationGroup{RepType: acctest.Required, Group: IotDigitalTwinAdapterInboundRoutesReferencePayloadRepresentation},
	}
	IotDigitalTwinAdapterInboundEnvelopeEnvelopeMappingRepresentation = map[string]interface{}{
		"time_observed": acctest.Representation{RepType: acctest.Required, Create: `$.time`, Update: `$.time`},
	}

	IotDigitalTwinAdapterInboundEnvelopeReferencePayloadRepresentation = map[string]interface{}{
		"data":        acctest.Representation{RepType: acctest.Required, Create: map[string]string{"time": "2025-08-26T05:47:13.842497Z", "temperature": "98"}, Update: map[string]string{"time": "2025-08-26T05:47:13.842557Z", "temperature": "98.6"}},
		"data_format": acctest.Representation{RepType: acctest.Required, Create: `JSON`},
	}

	IotDigitalTwinAdapterInboundRoutesReferencePayloadRepresentation = map[string]interface{}{
		"data":        acctest.Representation{RepType: acctest.Required, Create: map[string]string{"time": "2025-08-26T05:47:13.842497Z", "temperature": "98"}, Update: map[string]string{"time": "2025-08-26T05:47:13.842557Z", "temperature": "98.6"}},
		"data_format": acctest.Representation{RepType: acctest.Required, Create: `JSON`},
	}

	IotDigitalTwinAdapterResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Required, acctest.Create, IotDigitalTwinModelRepresentation)
)

// issue-routing-tag: iot/default
func TestIotDigitalTwinAdapterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotDigitalTwinAdapterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	resourceName := "oci_iot_digital_twin_adapter.test_digital_twin_adapter"
	datasourceName := "data.oci_iot_digital_twin_adapters.test_digital_twin_adapters"
	singularDatasourceName := "data.oci_iot_digital_twin_adapter.test_digital_twin_adapter"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+iotDomainIdVariableStr+IotDigitalTwinAdapterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Optional, acctest.Create, IotDigitalTwinAdapterRepresentation), "iot", "digitalTwinAdapter", t)

	acctest.ResourceTest(t, testAccCheckIotDigitalTwinAdapterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + iotDomainIdVariableStr + IotDigitalTwinAdapterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Required, acctest.Create, IotDigitalTwinAdapterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_model_spec_uri"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + iotDomainIdVariableStr + IotDigitalTwinAdapterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + iotDomainIdVariableStr + IotDigitalTwinAdapterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Optional, acctest.Create, IotDigitalTwinAdapterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_model_spec_uri"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.envelope_mapping.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.envelope_mapping.0.time_observed", "$.time"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_endpoint", "/"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_payload.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_payload.0.data.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_payload.0.data_format", "JSON"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.condition", "*"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.description", "Default condition"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.payload_mapping.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.reference_payload.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.reference_payload.0.data.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.reference_payload.0.data_format", "JSON"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
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
			Config: config + iotDomainIdVariableStr + IotDigitalTwinAdapterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Optional, acctest.Update, IotDigitalTwinAdapterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_model_spec_uri"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.envelope_mapping.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.envelope_mapping.0.time_observed", "$.time"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_endpoint", "/test"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_payload.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_payload.0.data.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "inbound_envelope.0.reference_payload.0.data_format", "JSON"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.condition", "*"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.description", "Default condition 1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.payload_mapping.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.reference_payload.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.reference_payload.0.data.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "inbound_routes.0.reference_payload.0.data_format", "JSON"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_adapters", "test_digital_twin_adapters", acctest.Optional, acctest.Update, IotDigitalTwinAdapterDataSourceRepresentation) +
				iotDomainIdVariableStr + IotDigitalTwinAdapterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Optional, acctest.Update, IotDigitalTwinAdapterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "digital_twin_model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "digital_twin_model_spec_uri"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "digital_twin_adapter_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "digital_twin_adapter_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Required, acctest.Create, IotDigitalTwinAdapterSingularDataSourceRepresentation) +
				iotDomainIdVariableStr + IotDigitalTwinAdapterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_adapter_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_model_spec_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_envelope.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_envelope.0.envelope_mapping.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_envelope.0.envelope_mapping.0.time_observed", "$.time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_envelope.0.reference_endpoint", "/test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_envelope.0.reference_payload.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_envelope.0.reference_payload.0.data.%", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_envelope.0.reference_payload.0.data_format", "JSON"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_routes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_routes.0.condition", "*"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_routes.0.description", "Default condition 1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_routes.0.payload_mapping.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_routes.0.reference_payload.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_routes.0.reference_payload.0.data.%", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inbound_routes.0.reference_payload.0.data_format", "JSON"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + IotDigitalTwinAdapterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIotDigitalTwinAdapterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IotClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_iot_digital_twin_adapter" {
			noResourceFound = false
			request := oci_iot.GetDigitalTwinAdapterRequest{}

			tmp := rs.Primary.ID
			request.DigitalTwinAdapterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")

			response, err := client.GetDigitalTwinAdapter(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IotDigitalTwinAdapter") {
		resource.AddTestSweepers("IotDigitalTwinAdapter", &resource.Sweeper{
			Name:         "IotDigitalTwinAdapter",
			Dependencies: acctest.DependencyGraph["digitalTwinAdapter"],
			F:            sweepIotDigitalTwinAdapterResource,
		})
	}
}

func sweepIotDigitalTwinAdapterResource(compartment string) error {
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()
	digitalTwinAdapterIds, err := getIotDigitalTwinAdapterIds(compartment)
	if err != nil {
		return err
	}
	for _, digitalTwinAdapterId := range digitalTwinAdapterIds {
		if ok := acctest.SweeperDefaultResourceId[digitalTwinAdapterId]; !ok {
			deleteDigitalTwinAdapterRequest := oci_iot.DeleteDigitalTwinAdapterRequest{}

			deleteDigitalTwinAdapterRequest.DigitalTwinAdapterId = &digitalTwinAdapterId

			deleteDigitalTwinAdapterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")
			_, error := iotClient.DeleteDigitalTwinAdapter(context.Background(), deleteDigitalTwinAdapterRequest)
			if error != nil {
				fmt.Printf("Error deleting DigitalTwinAdapter %s %s, It is possible that the resource is already deleted. Please verify manually \n", digitalTwinAdapterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &digitalTwinAdapterId, IotDigitalTwinAdapterSweepWaitCondition, time.Duration(3*time.Minute),
				IotDigitalTwinAdapterSweepResponseFetchOperation, "iot", true)
		}
	}
	return nil
}

func getIotDigitalTwinAdapterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DigitalTwinAdapterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()

	listDigitalTwinAdaptersRequest := oci_iot.ListDigitalTwinAdaptersRequest{}
	//listDigitalTwinAdaptersRequest.CompartmentId = &compartmentId

	iotDomainIds, error := getIotIotDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting iotDomainId required for DigitalTwinAdapter resource requests \n")
	}
	for _, iotDomainId := range iotDomainIds {
		listDigitalTwinAdaptersRequest.IotDomainId = &iotDomainId

		listDigitalTwinAdaptersRequest.LifecycleState = oci_iot.ListDigitalTwinAdaptersLifecycleStateActive
		listDigitalTwinAdaptersResponse, err := iotClient.ListDigitalTwinAdapters(context.Background(), listDigitalTwinAdaptersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DigitalTwinAdapter list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, digitalTwinAdapter := range listDigitalTwinAdaptersResponse.Items {
			id := *digitalTwinAdapter.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DigitalTwinAdapterId", id)
		}

	}
	return resourceIds, nil
}

func IotDigitalTwinAdapterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if digitalTwinAdapterResponse, ok := response.Response.(oci_iot.GetDigitalTwinAdapterResponse); ok {
		return digitalTwinAdapterResponse.LifecycleState != oci_iot.LifecycleStateDeleted
	}
	return false
}

func IotDigitalTwinAdapterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IotClient().GetDigitalTwinAdapter(context.Background(), oci_iot.GetDigitalTwinAdapterRequest{
		DigitalTwinAdapterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
