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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ResourcemanagerPrivateEndpointRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Required, acctest.Create, ResourceManagerprivateEndpointRepresentation)

	ResourcemanagerPrivateEndpointResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Optional, acctest.Update, ResourceManagerprivateEndpointRepresentation)

	ResourcemanagerResourcemanagerPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resourcemanager_private_endpoint.test_rms_private_endpoint.id}`},
	}

	ResourcemanagerResourcemanagerPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `My Private Endpoint`, Update: `displayName2`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_resourcemanager_private_endpoint.test_rms_private_endpoint.id}`},
		"vcn_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourcemanagerPrivateEndpointDataSourceFilterRepresentation}}
	ResourcemanagerPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_resourcemanager_private_endpoint.test_rms_private_endpoint.id}`}},
	}

	ResourceManagerprivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My Private Endpoint`, Update: `displayName2`},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Example Private Endpoint`, Update: `description2`},
		"dns_zones":      acctest.Representation{RepType: acctest.Optional, Create: []string{`dnsZones`}, Update: []string{`dnsZones2`}},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_used_with_configuration_source_provider": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"nsg_id_list": acctest.Representation{RepType: acctest.Optional, Create: []string{`nsgIdList`}, Update: []string{`nsgIdList2`}},
	}

	ResourcemanagerPrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Required, acctest.Create, ResourceManagerprivateEndpointRepresentation)
)

// issue-routing-tag: resourcemanager/default
func TestResourcemanagerPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourcemanagerPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_resourcemanager_private_endpoint.test_rms_private_endpoint"
	datasourceName := "data.oci_resourcemanager_private_endpoints.test_private_endpoints"
	singularDatasourceName := "data.oci_resourcemanager_private_endpoint.test_rms_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Optional, acctest.Create, ResourceManagerprivateEndpointRepresentation), "resourcemanager", "privateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckResourcemanagerPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Required, acctest.Create, ResourceManagerprivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My Private Endpoint"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ResourcemanagerPrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Optional, acctest.Create, ResourceManagerprivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Example Private Endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My Private Endpoint"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_used_with_configuration_source_provider", "false"),
				resource.TestCheckResourceAttr(resourceName, "nsg_id_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ResourceManagerprivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Example Private Endpoint"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My Private Endpoint"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_used_with_configuration_source_provider", "false"),
				resource.TestCheckResourceAttr(resourceName, "nsg_id_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Optional, acctest.Update, ResourceManagerprivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_used_with_configuration_source_provider", "true"),
				resource.TestCheckResourceAttr(resourceName, "nsg_id_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_private_endpoints", "test_private_endpoints", acctest.Optional, acctest.Update, ResourcemanagerResourcemanagerPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Optional, acctest.Update, ResourceManagerprivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resourcemanager_private_endpoint", "test_rms_private_endpoint", acctest.Required, acctest.Create, ResourcemanagerResourcemanagerPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourcemanagerPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_used_with_configuration_source_provider", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_id_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_ips.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + ResourcemanagerPrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckResourcemanagerPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ResourceManagerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_resourcemanager_private_endpoint" {
			noResourceFound = false
			request := oci_resourcemanager.GetPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.PrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resourcemanager")

			response, err := client.GetPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_resourcemanager.PrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ResourcemanagerPrivateEndpoint") {
		resource.AddTestSweepers("ResourcemanagerPrivateEndpoint", &resource.Sweeper{
			Name:         "ResourcemanagerPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["privateEndpoint"],
			F:            sweepResourcemanagerPrivateEndpointResource,
		})
	}
}

func sweepResourcemanagerPrivateEndpointResource(compartment string) error {
	resourceManagerClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()
	privateEndpointIds, err := getResourcemanagerPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, privateEndpointId := range privateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[privateEndpointId]; !ok {
			deletePrivateEndpointRequest := oci_resourcemanager.DeletePrivateEndpointRequest{}

			deletePrivateEndpointRequest.PrivateEndpointId = &privateEndpointId

			deletePrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resourcemanager")
			_, error := resourceManagerClient.DeletePrivateEndpoint(context.Background(), deletePrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &privateEndpointId, ResourcemanagerPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				ResourcemanagerPrivateEndpointSweepResponseFetchOperation, "resourcemanager", true)
		}
	}
	return nil
}

func getResourcemanagerPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	resourceManagerClient := acctest.GetTestClients(&schema.ResourceData{}).ResourceManagerClient()

	listPrivateEndpointsRequest := oci_resourcemanager.ListPrivateEndpointsRequest{}
	listPrivateEndpointsRequest.CompartmentId = &compartmentId
	listPrivateEndpointsResponse, err := resourceManagerClient.ListPrivateEndpoints(context.Background(), listPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, privateEndpoint := range listPrivateEndpointsResponse.Items {
		id := *privateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PrivateEndpointId", id)
	}
	return resourceIds, nil
}

func ResourcemanagerPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if privateEndpointResponse, ok := response.Response.(oci_resourcemanager.GetPrivateEndpointResponse); ok {
		return privateEndpointResponse.LifecycleState != oci_resourcemanager.PrivateEndpointLifecycleStateDeleted
	}
	return false
}

func ResourcemanagerPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ResourceManagerClient().GetPrivateEndpoint(context.Background(), oci_resourcemanager.GetPrivateEndpointRequest{
		PrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
