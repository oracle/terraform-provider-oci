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
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v65/operatoraccesscontrol"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	resourceId  = `ocid1.exadatainfrastructure.test..` + utils.RandomString(50, utils.CharsetWithoutDigits)
	resourceIdO = `ocid1.exadatainfrastructure.test..` + utils.RandomString(50, utils.CharsetWithoutDigits)

	OperatorAccessControlOperatorControlAssignmentRequiredOnlyResource = OperatorAccessControlOperatorControlAssignmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Required, acctest.Create, OperatorAccessControlOperatorControlAssignmentRepresentation)

	OperatorAccessControlOperatorControlAssignmentResourceConfig = OperatorAccessControlOperatorControlAssignmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Optional, acctest.Update, OperatorAccessControlOperatorControlAssignmentRepresentation)

	OperatorAccessControlOperatorAccessControlOperatorControlAssignmentSingularDataSourceRepresentation = map[string]interface{}{
		"operator_control_assignment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_operator_access_control_operator_control_assignment.test_operator_control_assignment.id}`},
	}

	OperatorAccessControlOperatorAccessControlOperatorControlAssignmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"operator_control_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_operator_access_control_operator_control.test_operator_control.operator_control_name}`},
		"resource_name":         acctest.Representation{RepType: acctest.Optional, Create: `resourceName`},
		"resource_type":         acctest.Representation{RepType: acctest.Required, Create: `EXADATAINFRASTRUCTURE`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `CREATED`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OperatorAccessControlOperatorControlAssignmentDataSourceFilterRepresentation}}
	OperatorAccessControlOperatorControlAssignmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_operator_access_control_operator_control_assignment.test_operator_control_assignment.id}`}},
	}

	OperatorAccessControlOperatorControlAssignmentRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_enforced_always":                 acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"operator_control_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_operator_access_control_operator_control.test_operator_control.id}`},
		"resource_compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"resource_id":                        acctest.Representation{RepType: acctest.Required, Create: resourceId},
		"resource_name":                      acctest.Representation{RepType: acctest.Required, Create: `resourceName`},
		"resource_type":                      acctest.Representation{RepType: acctest.Required, Create: `EXADATAINFRASTRUCTURE`},
		"comment":                            acctest.Representation{RepType: acctest.Optional, Create: `comment`, Update: `comment2`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"is_auto_approve_during_maintenance": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},

		"is_hypervisor_log_forwarded": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_log_forwarded":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},

		"remote_syslog_server_address": acctest.Representation{RepType: acctest.Optional, Create: `remoteSyslogServerAddress`, Update: `remoteSyslogServerAddress2`},
		"remote_syslog_server_ca_cert": acctest.Representation{RepType: acctest.Optional, Create: `cmVtb3RlU3lzbG9nU2VydmVyQ0FDZXJ0`, Update: `cmVtb3RlU3lzbG9nU2VydmVyQ0FDZXJ0Mg==`},
		"remote_syslog_server_port":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},

		/*"time_assignment_from":               acctest.Representation{RepType: acctest.Optional, Create: `timeAssignmentFrom`, Update: `timeAssignmentFrom2`},
		"time_assignment_to":                 acctest.Representation{RepType: acctest.Optional, Create: `timeAssignmentTo`, Update: `timeAssignmentTo2`},*/
		"validate_assignment_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},

		"time_assignment_from": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"time_assignment_to":   acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
	}

	OperatorAccessControlOperatorControlAssignmentResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control", "test_operator_control", acctest.Required, acctest.Create, OperatorAccessControlOperatorControlRepresentation)
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

	config := acctest.ProviderTestConfig()

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_operator_access_control_operator_control_assignment.test_operator_control_assignment"
	datasourceName := "data.oci_operator_access_control_operator_control_assignments.test_operator_control_assignments"
	singularDatasourceName := "data.oci_operator_access_control_operator_control_assignment.test_operator_control_assignment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OperatorAccessControlOperatorControlAssignmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignmentO", acctest.Optional, acctest.Create, OperatorAccessControlOperatorControlAssignmentRepresentation), "operatoraccesscontrol", "operatorControlAssignment", t)

	acctest.ResourceTest(t, testAccCheckOperatorAccessControlOperatorControlAssignmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OperatorAccessControlOperatorControlAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Required, acctest.Create, OperatorAccessControlOperatorControlAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_default_assignment", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			//ExpectNonEmptyPlan: true,
			Config: config + compartmentIdVariableStr + OperatorAccessControlOperatorControlAssignmentResourceDependencies,
		},
		// verify Create with optionals
		{
			ExpectNonEmptyPlan: true,
			Config: config + compartmentIdVariableStr + OperatorAccessControlOperatorControlAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Optional, acctest.Create, OperatorAccessControlOperatorControlAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_hypervisor_log_forwarded", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_log_forwarded", "false"),

				resource.TestCheckResourceAttr(resourceName, "is_default_assignment", "false"),

				resource.TestCheckResourceAttrSet(resourceName, "operator_control_id"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_address", "remoteSyslogServerAddress"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_ca_cert", "cmVtb3RlU3lzbG9nU2VydmVyQ0FDZXJ0"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_from"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_to"),

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
			ExpectNonEmptyPlan: true,
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OperatorAccessControlOperatorControlAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OperatorAccessControlOperatorControlAssignmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_hypervisor_log_forwarded", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_log_forwarded", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_id"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_address", "remoteSyslogServerAddress"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_ca_cert", "cmVtb3RlU3lzbG9nU2VydmVyQ0FDZXJ0"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_name"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),

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
			ExpectNonEmptyPlan: true,
			Config: config + compartmentIdVariableStr + OperatorAccessControlOperatorControlAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Optional, acctest.Update, OperatorAccessControlOperatorControlAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "comment", "comment2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_auto_approve_during_maintenance", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_hypervisor_log_forwarded", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_default_assignment", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_log_forwarded", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "operator_control_id"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_address", "remoteSyslogServerAddress2"),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_ca_cert", "cmVtb3RlU3lzbG9nU2VydmVyQ0FDZXJ0Mg=="),
				resource.TestCheckResourceAttr(resourceName, "remote_syslog_server_port", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_name", "resourceName"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_from"),
				resource.TestCheckNoResourceAttr(resourceName, "time_assignment_to"),

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
			ExpectNonEmptyPlan: true,
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_control_assignments", "test_operator_control_assignments", acctest.Optional, acctest.Update, OperatorAccessControlOperatorAccessControlOperatorControlAssignmentDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlOperatorControlAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Optional, acctest.Update, OperatorAccessControlOperatorControlAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_control_assignment", "test_operator_control_assignment", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlOperatorControlAssignmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlOperatorControlAssignmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operator_control_assignment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "assigner_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "comment", "comment2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "error_code"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "error_message"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_approve_during_maintenance", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_default_assignment", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enforced_always", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_hypervisor_log_forwarded", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_log_forwarded", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "op_control_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "remote_syslog_server_address", "remoteSyslogServerAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "remote_syslog_server_ca_cert", "cmVtb3RlU3lzbG9nU2VydmVyQ0FDZXJ0Mg=="),
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
		// verify resource import
		{
			ExpectNonEmptyPlan:      true,
			Config:                  config + OperatorAccessControlOperatorControlAssignmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"defined_tags", "validate_assignment_trigger"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOperatorAccessControlOperatorControlAssignmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperatorControlAssignmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_operator_access_control_operator_control_assignment" {
			noResourceFound = false
			request := oci_operator_access_control.GetOperatorControlAssignmentRequest{}

			tmp := rs.Primary.ID
			request.OperatorControlAssignmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "operator_access_control")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OperatorAccessControlOperatorControlAssignment") {
		resource.AddTestSweepers("OperatorAccessControlOperatorControlAssignment", &resource.Sweeper{
			Name:         "OperatorAccessControlOperatorControlAssignment",
			Dependencies: acctest.DependencyGraph["operatorControlAssignment"],
			F:            sweepOperatorAccessControlOperatorControlAssignmentResource,
		})
	}
}

func sweepOperatorAccessControlOperatorControlAssignmentResource(compartment string) error {
	operatorControlAssignmentClient := acctest.GetTestClients(&schema.ResourceData{}).OperatorControlAssignmentClient()
	operatorControlAssignmentIds, err := getOperatorAccessControlOperatorControlAssignmentIds(compartment)
	if err != nil {
		return err
	}
	for _, operatorControlAssignmentId := range operatorControlAssignmentIds {
		if ok := acctest.SweeperDefaultResourceId[operatorControlAssignmentId]; !ok {
			deleteOperatorControlAssignmentRequest := oci_operator_access_control.DeleteOperatorControlAssignmentRequest{}

			deleteOperatorControlAssignmentRequest.OperatorControlAssignmentId = &operatorControlAssignmentId

			deleteOperatorControlAssignmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "operator_access_control")
			_, error := operatorControlAssignmentClient.DeleteOperatorControlAssignment(context.Background(), deleteOperatorControlAssignmentRequest)
			if error != nil {
				fmt.Printf("Error deleting OperatorControlAssignment %s %s, It is possible that the resource is already deleted. Please verify manually \n", operatorControlAssignmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &operatorControlAssignmentId, OperatorAccessControlOperatorControlAssignmentSweepWaitCondition, time.Duration(3*time.Minute),
				OperatorAccessControlOperatorControlAssignmentSweepResponseFetchOperation, "operator_access_control", true)
		}
	}
	return nil
}

func getOperatorAccessControlOperatorControlAssignmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OperatorControlAssignmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operatorControlAssignmentClient := acctest.GetTestClients(&schema.ResourceData{}).OperatorControlAssignmentClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OperatorControlAssignmentId", id)
	}
	return resourceIds, nil
}

func OperatorAccessControlOperatorControlAssignmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if operatorControlAssignmentResponse, ok := response.Response.(oci_operator_access_control.GetOperatorControlAssignmentResponse); ok {
		return operatorControlAssignmentResponse.LifecycleState != oci_operator_access_control.OperatorControlAssignmentLifecycleStatesDeleted
	}
	return false
}

func OperatorAccessControlOperatorControlAssignmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperatorControlAssignmentClient().GetOperatorControlAssignment(context.Background(), oci_operator_access_control.GetOperatorControlAssignmentRequest{
		OperatorControlAssignmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
