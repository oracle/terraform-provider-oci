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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v54/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	resourceId  = `ocid1.exadatainfrastructure.test..` + RandomString(50, charsetWithoutDigits)
	resourceIdO = `ocid1.exadatainfrastructure.test..` + RandomString(50, charsetWithoutDigits)

	OperatorControlAssignmentRequiredOnlyResource = OperatorControlAssignmentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", Required, Create, operatorControlAssignmentRepresentation)

	OperatorControlAssignmentResourceConfig = OperatorControlAssignmentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", Optional, Update, operatorControlAssignmentRepresentation)

	operatorControlAssignmentSingularDataSourceRepresentation = map[string]interface{}{
		"operator_control_assignment_id": Representation{RepType: Required, Create: `${oci_operator_access_control_operator_control_assignment.test_operator_control_assignment.id}`},
	}

	operatorControlAssignmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        Representation{RepType: Required, Create: `${var.compartment_id}`},
		"operator_control_name": Representation{RepType: Optional, Create: `${oci_operator_access_control_operator_control.test_operator_control.operator_control_name}`},
		"resource_name":         Representation{RepType: Optional, Create: `resourceName`},
		"resource_type":         Representation{RepType: Required, Create: `EXADATAINFRASTRUCTURE`},
		"state":                 Representation{RepType: Optional, Create: `CREATED`},
		"filter":                RepresentationGroup{Required, operatorControlAssignmentDataSourceFilterRepresentation}}
	operatorControlAssignmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_operator_access_control_operator_control_assignment.test_operator_control_assignment.id}`}},
	}

	operatorControlAssignmentRepresentation = map[string]interface{}{
		"compartment_id":                     Representation{RepType: Required, Create: `${var.compartment_id}`},
		"is_enforced_always":                 Representation{RepType: Required, Create: `true`, Update: `true`},
		"operator_control_id":                Representation{RepType: Required, Create: `${oci_operator_access_control_operator_control.test_operator_control.id}`},
		"resource_compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"resource_id":                        Representation{RepType: Required, Create: resourceId},
		"resource_name":                      Representation{RepType: Required, Create: `resourceName`},
		"resource_type":                      Representation{RepType: Required, Create: `EXADATAINFRASTRUCTURE`},
		"comment":                            Representation{RepType: Optional, Create: `comment`, Update: `comment2`},
		"defined_tags":                       Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"is_auto_approve_during_maintenance": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_log_forwarded":                   Representation{RepType: Optional, Create: `true`, Update: `true`},
		"remote_syslog_server_address":       Representation{RepType: Optional, Create: `remoteSyslogServerAddress`, Update: `remoteSyslogServerAddress2`},
		"remote_syslog_server_ca_cert":       Representation{RepType: Optional, Create: `remoteSyslogServerCACert`, Update: `remoteSyslogServerCACert2`},
		"remote_syslog_server_port":          Representation{RepType: Optional, Create: `10`, Update: `11`},
		"time_assignment_from":               Representation{RepType: Optional, Create: nil, Update: nil},
		"time_assignment_to":                 Representation{RepType: Optional, Create: nil, Update: nil},
	}

	OperatorControlAssignmentResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", Required, Create, operatorControlRepresentation)
)

func getTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	return t
}

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlOperatorControlAssignmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOperatorAccessControlOperatorControlAssignmentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_operator_access_control_operator_control_assignment.test_operator_control_assignment"
	datasourceName := "data.oci_operator_access_control_operator_control_assignments.test_operator_control_assignments"
	singularDatasourceName := "data.oci_operator_access_control_operator_control_assignment.test_operator_control_assignment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+OperatorControlAssignmentResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignmentO", Optional, Create, operatorControlAssignmentRepresentation), "operatoraccesscontrol", "operatorControlAssignment", t)

	ResourceTest(t, testAccCheckOperatorAccessControlOperatorControlAssignmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OperatorControlAssignmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", Required, Create, operatorControlAssignmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			ExpectNonEmptyPlan: true,
			Config:             config + compartmentIdVariableStr + OperatorControlAssignmentResourceDependencies,
		},
		// verify Create with optionals
		{
			ExpectNonEmptyPlan: true,
			Config: config + compartmentIdVariableStr + OperatorControlAssignmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", Optional, Create, operatorControlAssignmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_log_forwarded", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_id"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_address", "remoteSyslogServerAddress"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_ca_cert", "remoteSyslogServerCACert"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_from"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_to"),

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

		// verify updates to updatable parameters
		{
			ExpectNonEmptyPlan: true,
			Config: config + compartmentIdVariableStr + OperatorControlAssignmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", Optional, Update, operatorControlAssignmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment", "comment2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_log_forwarded", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_id"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_address", "remoteSyslogServerAddress2"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_ca_cert", "remoteSyslogServerCACert2"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_port", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_from"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_to"),

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
			ExpectNonEmptyPlan: true,
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_control_assignments", "test_operator_control_assignments", Optional, Update, operatorControlAssignmentDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorControlAssignmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", Optional, Update, operatorControlAssignmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "operator_control_name"),
				resource.TestCheckResourceAttr(datasourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "CREATED"),
			),
		},
		// verify singular datasource
		{
			ExpectNonEmptyPlan: true,
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", Required, Create, operatorControlAssignmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorControlAssignmentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operator_control_assignment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "assigner_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "comment", "comment2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "error_code"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "error_message"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_approve_during_maintenance", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_log_forwarded", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "remote_syslog_server_address", "remoteSyslogServerAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "remote_syslog_server_ca_cert", "remoteSyslogServerCACert2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "remote_syslog_server_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "time_assignment_from"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "time_assignment_to"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_assignment"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "time_of_deletion"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			ExpectNonEmptyPlan: true,
			Config:             config + compartmentIdVariableStr + OperatorControlAssignmentResourceConfig,
		},
		// verify resource import
		{
			ExpectNonEmptyPlan:      true,
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOperatorAccessControlOperatorControlAssignmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).operatorControlAssignmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_operator_access_control_operator_control_assignment" {
			noResourceFound = false
			request := oci_operator_access_control.GetOperatorControlAssignmentRequest{}

			tmp := rs.Primary.ID
			request.OperatorControlAssignmentId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "operator_access_control")

			response, err := client.GetOperatorControlAssignment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_operator_access_control.OperatorControlAssignmentLifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
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
	if !InSweeperExcludeList("OperatorAccessControlOperatorControlAssignment") {
		resource.AddTestSweepers("OperatorAccessControlOperatorControlAssignment", &resource.Sweeper{
			Name:         "OperatorAccessControlOperatorControlAssignment",
			Dependencies: DependencyGraph["operatorControlAssignment"],
			F:            sweepOperatorAccessControlOperatorControlAssignmentResource,
		})
	}
}

func sweepOperatorAccessControlOperatorControlAssignmentResource(compartment string) error {
	operatorControlAssignmentClient := GetTestClients(&schema.ResourceData{}).operatorControlAssignmentClient()
	operatorControlAssignmentIds, err := getOperatorControlAssignmentIds(compartment)
	if err != nil {
		return err
	}
	for _, operatorControlAssignmentId := range operatorControlAssignmentIds {
		if ok := SweeperDefaultResourceId[operatorControlAssignmentId]; !ok {
			deleteOperatorControlAssignmentRequest := oci_operator_access_control.DeleteOperatorControlAssignmentRequest{}

			deleteOperatorControlAssignmentRequest.OperatorControlAssignmentId = &operatorControlAssignmentId

			deleteOperatorControlAssignmentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "operator_access_control")
			_, error := operatorControlAssignmentClient.DeleteOperatorControlAssignment(context.Background(), deleteOperatorControlAssignmentRequest)
			if error != nil {
				fmt.Printf("Error deleting OperatorControlAssignment %s %s, It is possible that the resource is already deleted. Please verify manually \n", operatorControlAssignmentId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &operatorControlAssignmentId, operatorControlAssignmentSweepWaitCondition, time.Duration(3*time.Minute),
				operatorControlAssignmentSweepResponseFetchOperation, "operator_access_control", true)
		}
	}
	return nil
}

func getOperatorControlAssignmentIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "OperatorControlAssignmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operatorControlAssignmentClient := GetTestClients(&schema.ResourceData{}).operatorControlAssignmentClient()

	listOperatorControlAssignmentsRequest := oci_operator_access_control.ListOperatorControlAssignmentsRequest{}
	listOperatorControlAssignmentsRequest.CompartmentId = &compartmentId
	listOperatorControlAssignmentsRequest.LifecycleState = oci_operator_access_control.ListOperatorControlAssignmentsLifecycleStateCreated
	listOperatorControlAssignmentsResponse, err := operatorControlAssignmentClient.ListOperatorControlAssignments(context.Background(), listOperatorControlAssignmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OperatorControlAssignment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, operatorControlAssignment := range listOperatorControlAssignmentsResponse.Items {
		id := *operatorControlAssignment.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "OperatorControlAssignmentId", id)
	}
	return resourceIds, nil
}

func operatorControlAssignmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if operatorControlAssignmentResponse, ok := response.Response.(oci_operator_access_control.GetOperatorControlAssignmentResponse); ok {
		return operatorControlAssignmentResponse.LifecycleState != oci_operator_access_control.OperatorControlAssignmentLifecycleStatesDeleted
	}
	return false
}

func operatorControlAssignmentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.operatorControlAssignmentClient().GetOperatorControlAssignment(context.Background(), oci_operator_access_control.GetOperatorControlAssignmentRequest{
		OperatorControlAssignmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
