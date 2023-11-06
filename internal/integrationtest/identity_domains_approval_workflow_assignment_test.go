// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsApprovalWorkflowAssignmentRequiredOnlyResource = IdentityDomainsApprovalWorkflowAssignmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignment", "test_approval_workflow_assignment", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowAssignmentRepresentation)

	IdentityDomainsApprovalWorkflowAssignmentResourceConfig = IdentityDomainsApprovalWorkflowAssignmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignment", "test_approval_workflow_assignment", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowAssignmentRepresentation)

	IdentityDomainsApprovalWorkflowAssignmentSingularDataSourceRepresentation = map[string]interface{}{
		"approval_workflow_assignment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_approval_workflow_assignment.test_approval_workflow_assignment.id}`},
		"idcs_endpoint":                   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsApprovalWorkflowAssignmentDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"approval_workflow_assignment_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"approval_workflow_assignment_filter": acctest.Representation{RepType: acctest.Optional, Create: `id eq \"${oci_identity_domains_approval_workflow_assignment.test_approval_workflow_assignment.id}\"`},
		"attribute_sets":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                         acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsApprovalWorkflowAssignmentRepresentation = map[string]interface{}{
		"approval_workflow": acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsApprovalWorkflowAssignmentApprovalWorkflowRepresentation},
		"assigned_to":       acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsApprovalWorkflowAssignmentAssignedToRepresentation},
		"assignment_type":   acctest.Representation{RepType: acctest.Required, Create: `MEMBERSHIP`},
		"idcs_endpoint":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":           acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflowAssignment`}},
		"attribute_sets":    acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"tags":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsApprovalWorkflowAssignmentTagsRepresentation},
	}
	IdentityDomainsApprovalWorkflowAssignmentApprovalWorkflowRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `ApprovalWorkflow`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_approval_workflow.test_approval_workflow.id}`},
	}
	IdentityDomainsApprovalWorkflowAssignmentAssignedToRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `Group`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}
	IdentityDomainsApprovalWorkflowAssignmentTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsApprovalWorkflowAssignmentResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsApprovalWorkflowAssignmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsApprovalWorkflowAssignmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_approval_workflow_assignment.test_approval_workflow_assignment"
	datasourceName := "data.oci_identity_domains_approval_workflow_assignments.test_approval_workflow_assignments"
	singularDatasourceName := "data.oci_identity_domains_approval_workflow_assignment.test_approval_workflow_assignment"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsApprovalWorkflowAssignmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignment", "test_approval_workflow_assignment", acctest.Optional, acctest.Create, IdentityDomainsApprovalWorkflowAssignmentRepresentation), "identitydomains", "approvalWorkflowAssignment", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsApprovalWorkflowAssignmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignment", "test_approval_workflow_assignment", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approval_workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow.0.value"),
				resource.TestCheckResourceAttr(resourceName, "assigned_to.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "assigned_to.0.type", "Group"),
				resource.TestCheckResourceAttr(resourceName, "assigned_to.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "assignment_type", "MEMBERSHIP"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowAssignmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignment", "test_approval_workflow_assignment", acctest.Optional, acctest.Create, IdentityDomainsApprovalWorkflowAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approval_workflow.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow.0.value"),
				resource.TestCheckResourceAttr(resourceName, "assigned_to.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "assigned_to.0.type", "Group"),
				resource.TestCheckResourceAttr(resourceName, "assigned_to.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "assignment_type", "MEMBERSHIP"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "approvalWorkflowAssignments", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignments", "test_approval_workflow_assignments", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowAssignmentDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApprovalWorkflowAssignmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignment", "test_approval_workflow_assignment", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowAssignmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_assignment_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_assignments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_assignments.0.schemas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "approval_workflow_assignments.0.approval_workflow.0.value"),
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_assignments.0.assigned_to.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_assignments.0.assigned_to.0.type", "Group"),
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_assignments.0.assigned_to.0.value", "value"),
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_assignments.0.assignment_type", "MEMBERSHIP"),
				resource.TestCheckResourceAttrSet(datasourceName, "approval_workflow_assignments.0.id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_approval_workflow_assignment", "test_approval_workflow_assignment", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowAssignmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApprovalWorkflowAssignmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approval_workflow_assignment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "approval_workflow.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approval_workflow.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assigned_to.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assigned_to.0.type", "Group"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assigned_to.0.value", "value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assignment_type", "MEMBERSHIP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsApprovalWorkflowAssignmentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_approval_workflow_assignment", "approvalWorkflowAssignments"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
				"approval_workflow.0.type",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsApprovalWorkflowAssignmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_approval_workflow_assignment" {
			noResourceFound = false
			request := oci_identity_domains.GetApprovalWorkflowAssignmentRequest{}

			tmp := rs.Primary.ID
			request.ApprovalWorkflowAssignmentId = &tmp

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			_, err := client.GetApprovalWorkflowAssignment(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsApprovalWorkflowAssignment") {
		resource.AddTestSweepers("IdentityDomainsApprovalWorkflowAssignment", &resource.Sweeper{
			Name:         "IdentityDomainsApprovalWorkflowAssignment",
			Dependencies: acctest.DependencyGraph["approvalWorkflowAssignment"],
			F:            sweepIdentityDomainsApprovalWorkflowAssignmentResource,
		})
	}
}

func sweepIdentityDomainsApprovalWorkflowAssignmentResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	approvalWorkflowAssignmentIds, err := getIdentityDomainsApprovalWorkflowAssignmentIds(compartment)
	if err != nil {
		return err
	}
	for _, approvalWorkflowAssignmentId := range approvalWorkflowAssignmentIds {
		if ok := acctest.SweeperDefaultResourceId[approvalWorkflowAssignmentId]; !ok {
			deleteApprovalWorkflowAssignmentRequest := oci_identity_domains.DeleteApprovalWorkflowAssignmentRequest{}

			deleteApprovalWorkflowAssignmentRequest.ApprovalWorkflowAssignmentId = &approvalWorkflowAssignmentId

			deleteApprovalWorkflowAssignmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteApprovalWorkflowAssignment(context.Background(), deleteApprovalWorkflowAssignmentRequest)
			if error != nil {
				fmt.Printf("Error deleting ApprovalWorkflowAssignment %s %s, It is possible that the resource is already deleted. Please verify manually \n", approvalWorkflowAssignmentId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsApprovalWorkflowAssignmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApprovalWorkflowAssignmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listApprovalWorkflowAssignmentsRequest := oci_identity_domains.ListApprovalWorkflowAssignmentsRequest{}
	listApprovalWorkflowAssignmentsResponse, err := identityDomainsClient.ListApprovalWorkflowAssignments(context.Background(), listApprovalWorkflowAssignmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ApprovalWorkflowAssignment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, approvalWorkflowAssignment := range listApprovalWorkflowAssignmentsResponse.Resources {
		id := *approvalWorkflowAssignment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApprovalWorkflowAssignmentId", id)
	}
	return resourceIds, nil
}
