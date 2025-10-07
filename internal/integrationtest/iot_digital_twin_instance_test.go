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
	IotDigitalTwinInstanceRequiredOnlyResource = IotDigitalTwinInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Required, acctest.Create, IotDigitalTwinInstanceRepresentation)

	IotDigitalTwinInstanceResourceConfig = IotDigitalTwinInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Optional, acctest.Update, IotDigitalTwinInstanceRepresentation)

	IotDigitalTwinInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"digital_twin_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_iot_digital_twin_instance.test_digital_twin_instance.id}`},
	}

	ignoreDigitalTwinInstanceDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	IotDigitalTwinInstanceDataSourceRepresentation = map[string]interface{}{
		"iot_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"digital_twin_model_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.id}`},
		"digital_twin_model_spec_uri": acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_model.test_digital_twin_model.spec_uri}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_instance.test_digital_twin_instance.id}`},
		"state":                       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: IotDigitalTwinInstanceDataSourceFilterRepresentation}}
	IotDigitalTwinInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_iot_digital_twin_instance.test_digital_twin_instance.id}`}},
	}

	IotDigitalTwinInstanceRepresentation = map[string]interface{}{
		"auth_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.auth_id}`},
		"iot_domain_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.iot_domain_id}`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"digital_twin_adapter_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_iot_digital_twin_adapter.test_digital_twin_adapter.id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"external_key":            acctest.Representation{RepType: acctest.Optional, Create: `externalKey`, Update: `externalKey2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Protocol": "Mqtt"}, Update: map[string]string{"Protocol": "MQTT"}},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDigitalTwinInstanceDefinedTagsChangesRepresentation},
	}

	IotDigitalTwinInstanceResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_adapter", "test_digital_twin_adapter", acctest.Required, acctest.Create, IotDigitalTwinAdapterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_model", "test_digital_twin_model", acctest.Required, acctest.Create, IotDigitalTwinModelRepresentation)
)

// issue-routing-tag: iot/default
func TestIotDigitalTwinInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIotDigitalTwinInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	iotDomainId := utils.GetEnvSettingWithBlankDefault("iot_domain_ocid")
	iotDomainIdVariableStr := fmt.Sprintf("variable \"iot_domain_id\" { default = \"%s\" }\n", iotDomainId)

	authId := utils.GetEnvSettingWithBlankDefault("auth_ocid")
	authIdVariableStr := fmt.Sprintf("variable \"auth_id\" { default = \"%s\" }\n", authId)

	resourceName := "oci_iot_digital_twin_instance.test_digital_twin_instance"
	datasourceName := "data.oci_iot_digital_twin_instances.test_digital_twin_instances"
	singularDatasourceName := "data.oci_iot_digital_twin_instance.test_digital_twin_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+authIdVariableStr+iotDomainIdVariableStr+IotDigitalTwinInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Optional, acctest.Create, IotDigitalTwinInstanceRepresentation), "iot", "digitalTwinInstance", t)

	acctest.ResourceTest(t, testAccCheckIotDigitalTwinInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + authIdVariableStr + iotDomainIdVariableStr + IotDigitalTwinInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Required, acctest.Create, IotDigitalTwinInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "auth_id"),
				resource.TestCheckResourceAttrSet(resourceName, "iot_domain_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + authIdVariableStr + iotDomainIdVariableStr + IotDigitalTwinInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + authIdVariableStr + iotDomainIdVariableStr + IotDigitalTwinInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Optional, acctest.Create, IotDigitalTwinInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "auth_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_adapter_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "external_key", "externalKey"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + authIdVariableStr + iotDomainIdVariableStr + IotDigitalTwinInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Optional, acctest.Update, IotDigitalTwinInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "auth_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "digital_twin_adapter_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "external_key", "externalKey2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_instances", "test_digital_twin_instances", acctest.Optional, acctest.Update, IotDigitalTwinInstanceDataSourceRepresentation) +
				authIdVariableStr + iotDomainIdVariableStr + IotDigitalTwinInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Optional, acctest.Update, IotDigitalTwinInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "digital_twin_model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "digital_twin_model_spec_uri"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "iot_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "digital_twin_instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "digital_twin_instance_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_iot_digital_twin_instance", "test_digital_twin_instance", acctest.Required, acctest.Create, IotDigitalTwinInstanceSingularDataSourceRepresentation) +
				authIdVariableStr + iotDomainIdVariableStr + IotDigitalTwinInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "digital_twin_model_spec_uri"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "external_key", "externalKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + IotDigitalTwinInstanceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIotDigitalTwinInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IotClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_iot_digital_twin_instance" {
			noResourceFound = false
			request := oci_iot.GetDigitalTwinInstanceRequest{}

			tmp := rs.Primary.ID
			request.DigitalTwinInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")

			response, err := client.GetDigitalTwinInstance(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IotDigitalTwinInstance") {
		resource.AddTestSweepers("IotDigitalTwinInstance", &resource.Sweeper{
			Name:         "IotDigitalTwinInstance",
			Dependencies: acctest.DependencyGraph["digitalTwinInstance"],
			F:            sweepIotDigitalTwinInstanceResource,
		})
	}
}

func sweepIotDigitalTwinInstanceResource(compartment string) error {
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()
	digitalTwinInstanceIds, err := getIotDigitalTwinInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, digitalTwinInstanceId := range digitalTwinInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[digitalTwinInstanceId]; !ok {
			deleteDigitalTwinInstanceRequest := oci_iot.DeleteDigitalTwinInstanceRequest{}

			deleteDigitalTwinInstanceRequest.DigitalTwinInstanceId = &digitalTwinInstanceId

			deleteDigitalTwinInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "iot")
			_, error := iotClient.DeleteDigitalTwinInstance(context.Background(), deleteDigitalTwinInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting DigitalTwinInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", digitalTwinInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &digitalTwinInstanceId, IotDigitalTwinInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				IotDigitalTwinInstanceSweepResponseFetchOperation, "iot", true)
		}
	}
	return nil
}

func getIotDigitalTwinInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DigitalTwinInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	iotClient := acctest.GetTestClients(&schema.ResourceData{}).IotClient()

	listDigitalTwinInstancesRequest := oci_iot.ListDigitalTwinInstancesRequest{}
	//listDigitalTwinInstancesRequest.CompartmentId = &compartmentId

	iotDomainIds, error := getIotIotDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting iotDomainId required for DigitalTwinInstance resource requests \n")
	}
	for _, iotDomainId := range iotDomainIds {
		listDigitalTwinInstancesRequest.IotDomainId = &iotDomainId

		listDigitalTwinInstancesRequest.LifecycleState = oci_iot.ListDigitalTwinInstancesLifecycleStateActive
		listDigitalTwinInstancesResponse, err := iotClient.ListDigitalTwinInstances(context.Background(), listDigitalTwinInstancesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DigitalTwinInstance list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, digitalTwinInstance := range listDigitalTwinInstancesResponse.Items {
			id := *digitalTwinInstance.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DigitalTwinInstanceId", id)
		}

	}
	return resourceIds, nil
}

func IotDigitalTwinInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if digitalTwinInstanceResponse, ok := response.Response.(oci_iot.GetDigitalTwinInstanceResponse); ok {
		return digitalTwinInstanceResponse.LifecycleState != oci_iot.LifecycleStateDeleted
	}
	return false
}

func IotDigitalTwinInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.IotClient().GetDigitalTwinInstance(context.Background(), oci_iot.GetDigitalTwinInstanceRequest{
		DigitalTwinInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
