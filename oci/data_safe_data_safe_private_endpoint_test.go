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
	oci_data_safe "github.com/oracle/oci-go-sdk/v48/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DataSafePrivateEndpointRequiredOnlyResource = DataSafePrivateEndpointResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Required, Create, dataSafePrivateEndpointRepresentation)

	DataSafePrivateEndpointResourceConfig = DataSafePrivateEndpointResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Optional, Update, dataSafePrivateEndpointRepresentation)

	dataSafePrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"data_safe_private_endpoint_id": Representation{RepType: Required, Create: `${oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint.id}`},
	}

	dataSafePrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"access_level":              Representation{RepType: Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `true`},
		"display_name":              Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":                     Representation{RepType: Optional, Create: `ACTIVE`},
		"vcn_id":                    Representation{RepType: Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":                    RepresentationGroup{Required, dataSafePrivateEndpointDataSourceFilterRepresentation}}
	dataSafePrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint.id}`}},
	}

	dataSafePrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"subnet_id":      Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"vcn_id":         Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":        Representation{RepType: Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
	}

	DataSafePrivateEndpointResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeDataSafePrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeDataSafePrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint"
	datasourceName := "data.oci_data_safe_data_safe_private_endpoints.test_data_safe_private_endpoints"
	singularDatasourceName := "data.oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DataSafePrivateEndpointResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Optional, Create, dataSafePrivateEndpointRepresentation), "datasafe", "dataSafePrivateEndpoint", t)

	ResourceTest(t, testAccCheckDataSafeDataSafePrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Required, Create, dataSafePrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafePrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Optional, Create, dataSafePrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Optional, Create,
					RepresentationCopyWithNewProperties(dataSafePrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataSafePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Optional, Update, dataSafePrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoints", "test_data_safe_private_endpoints", Optional, Update, dataSafePrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafePrivateEndpointResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Optional, Update, dataSafePrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "data_safe_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_safe_private_endpoints.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "data_safe_private_endpoints.0.defined_tags.%", "0"),
				resource.TestCheckResourceAttr(datasourceName, "data_safe_private_endpoints.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "data_safe_private_endpoints.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "data_safe_private_endpoints.0.freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_safe_private_endpoints.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_safe_private_endpoints.0.private_endpoint_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_safe_private_endpoints.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_safe_private_endpoints.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_safe_private_endpoints.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_safe_private_endpoints.0.vcn_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_data_safe_data_safe_private_endpoint", "test_data_safe_private_endpoint", Required, Create, dataSafePrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafePrivateEndpointResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_safe_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataSafePrivateEndpointResourceConfig,
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

func testAccCheckDataSafeDataSafePrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_data_safe_private_endpoint" {
			noResourceFound = false
			request := oci_data_safe.GetDataSafePrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.DataSafePrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "data_safe")

			response, err := client.GetDataSafePrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.LifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("DataSafeDataSafePrivateEndpoint") {
		resource.AddTestSweepers("DataSafeDataSafePrivateEndpoint", &resource.Sweeper{
			Name:         "DataSafeDataSafePrivateEndpoint",
			Dependencies: DependencyGraph["dataSafePrivateEndpoint"],
			F:            sweepDataSafeDataSafePrivateEndpointResource,
		})
	}
}

func sweepDataSafeDataSafePrivateEndpointResource(compartment string) error {
	dataSafeClient := GetTestClients(&schema.ResourceData{}).dataSafeClient()
	dataSafePrivateEndpointIds, err := getDataSafePrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, dataSafePrivateEndpointId := range dataSafePrivateEndpointIds {
		if ok := SweeperDefaultResourceId[dataSafePrivateEndpointId]; !ok {
			deleteDataSafePrivateEndpointRequest := oci_data_safe.DeleteDataSafePrivateEndpointRequest{}

			deleteDataSafePrivateEndpointRequest.DataSafePrivateEndpointId = &dataSafePrivateEndpointId

			deleteDataSafePrivateEndpointRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteDataSafePrivateEndpoint(context.Background(), deleteDataSafePrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting DataSafePrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataSafePrivateEndpointId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &dataSafePrivateEndpointId, dataSafePrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				dataSafePrivateEndpointSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafePrivateEndpointIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DataSafePrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := GetTestClients(&schema.ResourceData{}).dataSafeClient()

	listDataSafePrivateEndpointsRequest := oci_data_safe.ListDataSafePrivateEndpointsRequest{}
	listDataSafePrivateEndpointsRequest.CompartmentId = &compartmentId
	listDataSafePrivateEndpointsRequest.LifecycleState = oci_data_safe.ListDataSafePrivateEndpointsLifecycleStateActive
	listDataSafePrivateEndpointsResponse, err := dataSafeClient.ListDataSafePrivateEndpoints(context.Background(), listDataSafePrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DataSafePrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataSafePrivateEndpoint := range listDataSafePrivateEndpointsResponse.Items {
		id := *dataSafePrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DataSafePrivateEndpointId", id)
	}
	return resourceIds, nil
}

func dataSafePrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dataSafePrivateEndpointResponse, ok := response.Response.(oci_data_safe.GetDataSafePrivateEndpointResponse); ok {
		return dataSafePrivateEndpointResponse.LifecycleState != oci_data_safe.LifecycleStateDeleted
	}
	return false
}

func dataSafePrivateEndpointSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataSafeClient().GetDataSafePrivateEndpoint(context.Background(), oci_data_safe.GetDataSafePrivateEndpointRequest{
		DataSafePrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
