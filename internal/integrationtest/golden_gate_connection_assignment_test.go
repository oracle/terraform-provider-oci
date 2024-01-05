// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

/*
Note:

	This tests creates a connection assignment, and requires an existing deployment and connection which is assignable to the deployment.
	Required environment variables:
	 - TF_VAR_compartment_id - defines where to create the assignment
	 - TF_VAR_deployment_id - deployment which we would like to assign connection to
	 - TF_VAR_connection_id - connection for assign
*/
var (
	GoldenGateConnectionAssignmentRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection_assignment", "test_connection_assignment", acctest.Optional, acctest.Create, GoldenGateConnectionAssignmentRepresentation)

	GoldenGateConnectionAssignmentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection_assignment", "test_connection_assignment", acctest.Optional, acctest.Update, GoldenGateConnectionAssignmentRepresentation)

	GoldenGateGoldenGateConnectionAssignmentSingularDataSourceRepresentation = map[string]interface{}{
		"connection_assignment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_connection_assignment.test_connection_assignment.id}`},
	}

	GoldenGateGoldenGateConnectionAssignmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.connection_id}`},
		"deployment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.deployment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GoldenGateConnectionAssignmentDataSourceFilterRepresentation}}
	GoldenGateConnectionAssignmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_golden_gate_connection_assignment.test_connection_assignment.id}`}},
	}

	GoldenGateConnectionAssignmentRepresentation = map[string]interface{}{
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connection_id}`},
		"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.deployment_id}`},
	}
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateConnectionAssignmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateConnectionAssignmentResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID = "compartment_id"
		DEPLOYMENT_ID  = "deployment_id"
		CONNECTION_ID  = "connection_id"
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(DEPLOYMENT_ID, t) +
		makeVariableStr(CONNECTION_ID, t)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id")
	connectionId := utils.GetEnvSettingWithBlankDefault("connection_id")
	deploymentId := utils.GetEnvSettingWithBlankDefault("deployment_ocid")

	resourceName := "oci_golden_gate_connection_assignment.test_connection_assignment"
	datasourceName := "data.oci_golden_gate_connection_assignments.test_connection_assignments"
	singularDatasourceName := "data.oci_golden_gate_connection_assignment.test_connection_assignment"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	/*connectionResourceDependencies := acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", "test_connection", acctest.Required, acctest.Create, GoldenGateConnectionRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "depl_test_ggs_deployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation)
	*/
	acctest.SaveConfigContent(config, "goldengate", "connectionAssignment", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateConnectionAssignmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection_assignment", "test_connection_assignment", acctest.Required, acctest.Create, GoldenGateConnectionAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_connection_assignments", "test_connection_assignments", acctest.Optional, acctest.Update, GoldenGateGoldenGateConnectionAssignmentDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection_assignment", "test_connection_assignment", acctest.Optional, acctest.Update, GoldenGateConnectionAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "connection_id", connectionId),
				resource.TestCheckResourceAttr(datasourceName, "deployment_id", deploymentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "connection_assignment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "connection_assignment_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_connection_assignment", "test_connection_assignment", acctest.Required, acctest.Create, GoldenGateGoldenGateConnectionAssignmentSingularDataSourceRepresentation) +
				GoldenGateConnectionAssignmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_assignment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alias_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GoldenGateConnectionAssignmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGoldenGateConnectionAssignmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_connection_assignment" {
			noResourceFound = false
			request := oci_golden_gate.GetConnectionAssignmentRequest{}

			tmp := rs.Primary.ID
			request.ConnectionAssignmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			_, err := client.GetConnectionAssignment(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("GoldenGateConnectionAssignment") {
		resource.AddTestSweepers("GoldenGateConnectionAssignment", &resource.Sweeper{
			Name:         "GoldenGateConnectionAssignment",
			Dependencies: acctest.DependencyGraph["connectionAssignment"],
			F:            sweepGoldenGateConnectionAssignmentResource,
		})
	}
}

func sweepGoldenGateConnectionAssignmentResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	connectionAssignmentIds, err := getGoldenGateConnectionAssignmentIds(compartment)
	if err != nil {
		return err
	}
	for _, connectionAssignmentId := range connectionAssignmentIds {
		if ok := acctest.SweeperDefaultResourceId[connectionAssignmentId]; !ok {
			deleteConnectionAssignmentRequest := oci_golden_gate.DeleteConnectionAssignmentRequest{}

			deleteConnectionAssignmentRequest.ConnectionAssignmentId = &connectionAssignmentId

			deleteConnectionAssignmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteConnectionAssignment(context.Background(), deleteConnectionAssignmentRequest)
			if error != nil {
				fmt.Printf("Error deleting ConnectionAssignment %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectionAssignmentId, error)
				continue
			}
		}
	}
	return nil
}

func getGoldenGateConnectionAssignmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConnectionAssignmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listConnectionAssignmentsRequest := oci_golden_gate.ListConnectionAssignmentsRequest{}
	listConnectionAssignmentsRequest.CompartmentId = &compartmentId
	listConnectionAssignmentsResponse, err := goldenGateClient.ListConnectionAssignments(context.Background(), listConnectionAssignmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ConnectionAssignment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, connectionAssignment := range listConnectionAssignmentsResponse.Items {
		id := *connectionAssignment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionAssignmentId", id)
	}
	return resourceIds, nil
}
