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
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSciencePrivateEndpointRequiredOnlyResource = DataSciencePrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Required, acctest.Create, DataSciencePrivateEndpointRepresentation)

	DataSciencePrivateEndpointResourceConfig = DataSciencePrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Optional, acctest.Update, DataSciencePrivateEndpointRepresentation)

	DataSciencePrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"data_science_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_private_endpoint.test_data_science_private_endpoint.id}`},
	}

	DataSciencePrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.user_id}`},
		"data_science_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `NOTEBOOK_SESSION`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `pe_1234`, Update: `displayName2`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSciencePrivateEndpointDataSourceFilterRepresentation}}

	DataSciencePrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_private_endpoint.test_data_science_private_endpoint.id}`}},
	}

	DataSciencePrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_science_resource_type": acctest.Representation{RepType: acctest.Required, Create: `NOTEBOOK_SESSION`},
		"subnet_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `pe_1234`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"sub_domain":                 acctest.Representation{RepType: acctest.Optional, Create: `subDomain`},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: dataSciencePrivateEndpointIgnoreRepresentation},
	}

	dataSciencePrivateEndpointIgnoreRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DataSciencePrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceDataSciencePrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceDataSciencePrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	userId := utils.GetEnvSettingWithBlankDefault("user_ocid")
	userIdVariableStr := fmt.Sprintf("variable \"user_id\" { default = \"%s\" }\n", userId)

	resourceName := "oci_datascience_private_endpoint.test_data_science_private_endpoint"
	datasourceName := "data.oci_datascience_private_endpoints.test_data_science_private_endpoint"
	singularDatasourceName := "data.oci_datascience_private_endpoint.test_data_science_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSciencePrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Optional, acctest.Create, DataSciencePrivateEndpointRepresentation), "datascience", "dataSciencePrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckDatascienceDataSciencePrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSciencePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Required, acctest.Create, DataSciencePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_science_resource_type", "NOTEBOOK_SESSION"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSciencePrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSciencePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Optional, acctest.Create, DataSciencePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "data_science_resource_type", "NOTEBOOK_SESSION"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_1234"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "sub_domain", "subDomain"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSciencePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSciencePrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "data_science_resource_type", "NOTEBOOK_SESSION"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_1234"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "sub_domain", "subDomain"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + DataSciencePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Optional, acctest.Update, DataSciencePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "data_science_resource_type", "NOTEBOOK_SESSION"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "fqdn"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "sub_domain", "subDomain"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_private_endpoints", "test_data_science_private_endpoint", acctest.Optional, acctest.Update, DataSciencePrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + userIdVariableStr + DataSciencePrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Optional, acctest.Update, DataSciencePrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "data_science_resource_type", "NOTEBOOK_SESSION"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "data_science_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_science_private_endpoints.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "data_science_private_endpoints.0.data_science_resource_type", "NOTEBOOK_SESSION"),
				resource.TestCheckResourceAttr(datasourceName, "data_science_private_endpoints.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_science_private_endpoints.0.fqdn"),
				resource.TestCheckResourceAttr(datasourceName, "data_science_private_endpoints.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_science_private_endpoints.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_science_private_endpoints.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_science_private_endpoints.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_science_private_endpoints.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_science_private_endpoints.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Required, acctest.Create, DataSciencePrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSciencePrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_science_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_science_resource_type", "NOTEBOOK_SESSION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DataSciencePrivateEndpointRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"sub_domain",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatascienceDataSciencePrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_private_endpoint" {
			noResourceFound = false
			request := oci_datascience.GetDataSciencePrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.DataSciencePrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetDataSciencePrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.DataSciencePrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceDataSciencePrivateEndpoint") {
		resource.AddTestSweepers("DatascienceDataSciencePrivateEndpoint", &resource.Sweeper{
			Name:         "DatascienceDataSciencePrivateEndpoint",
			Dependencies: acctest.DependencyGraph["dataSciencePrivateEndpoint"],
			F:            sweepDatascienceDataSciencePrivateEndpointResource,
		})
	}
}

func sweepDatascienceDataSciencePrivateEndpointResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	dataSciencePrivateEndpointIds, err := getDatascienceDataSciencePrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, dataSciencePrivateEndpointId := range dataSciencePrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[dataSciencePrivateEndpointId]; !ok {
			deleteDataSciencePrivateEndpointRequest := oci_datascience.DeleteDataSciencePrivateEndpointRequest{}

			deleteDataSciencePrivateEndpointRequest.DataSciencePrivateEndpointId = &dataSciencePrivateEndpointId

			deleteDataSciencePrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteDataSciencePrivateEndpoint(context.Background(), deleteDataSciencePrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting DataSciencePrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataSciencePrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dataSciencePrivateEndpointId, DatascienceDataSciencePrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceDataSciencePrivateEndpointSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceDataSciencePrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DataSciencePrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listDataSciencePrivateEndpointsRequest := oci_datascience.ListDataSciencePrivateEndpointsRequest{}
	listDataSciencePrivateEndpointsRequest.CompartmentId = &compartmentId
	listDataSciencePrivateEndpointsRequest.LifecycleState = oci_datascience.ListDataSciencePrivateEndpointsLifecycleStateActive
	listDataSciencePrivateEndpointsResponse, err := dataScienceClient.ListDataSciencePrivateEndpoints(context.Background(), listDataSciencePrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DataSciencePrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataSciencePrivateEndpoint := range listDataSciencePrivateEndpointsResponse.Items {
		id := *dataSciencePrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DataSciencePrivateEndpointId", id)
	}
	return resourceIds, nil
}

func DatascienceDataSciencePrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dataSciencePrivateEndpointResponse, ok := response.Response.(oci_datascience.GetDataSciencePrivateEndpointResponse); ok {
		return dataSciencePrivateEndpointResponse.LifecycleState != oci_datascience.DataSciencePrivateEndpointLifecycleStateDeleted
	}
	return false
}

func DatascienceDataSciencePrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetDataSciencePrivateEndpoint(context.Background(), oci_datascience.GetDataSciencePrivateEndpointRequest{
		DataSciencePrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
