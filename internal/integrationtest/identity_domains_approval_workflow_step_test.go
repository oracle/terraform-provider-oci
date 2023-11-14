// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"regexp"
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
	IdentityDomainsApprovalWorkflowStepRequiredOnlyResource = IdentityDomainsApprovalWorkflowStepResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowStepRepresentation)

	IdentityDomainsApprovalWorkflowStepResourceConfig = IdentityDomainsApprovalWorkflowStepResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowStepRepresentation)

	IdentityDomainsApprovalWorkflowStepSingularDataSourceRepresentation = map[string]interface{}{
		"approval_workflow_step_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_approval_workflow_step.test_approval_workflow_step.id}`},
		"idcs_endpoint":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":            acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsApprovalWorkflowStepDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"approval_workflow_step_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"approval_workflow_step_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":                acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                   acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsApprovalWorkflowStepRepresentation = map[string]interface{}{
		"idcs_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"order":                 acctest.Representation{RepType: acctest.Required, Create: `1`},
		"schemas":               acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflowStep`}},
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `escalation`},
		"approvers":             acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsApprovalWorkflowStepApproversRepresentation},
		"approvers_expressions": acctest.Representation{RepType: acctest.Optional, Create: []string{`User Manager`}},
		"attribute_sets":        acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"minimum_approvals":     acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"tags":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsApprovalWorkflowStepTagsRepresentation},
	}
	IdentityDomainsApprovalWorkflowStepApproversRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `User`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_user.test_user.id}`},
	}
	IdentityDomainsApprovalWorkflowStepTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	IdentityDomainsApprovalWorkflowStepResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsApprovalWorkflowStepResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsApprovalWorkflowStepResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_approval_workflow_step.test_approval_workflow_step"
	datasourceName := "data.oci_identity_domains_approval_workflow_steps.test_approval_workflow_steps"
	singularDatasourceName := "data.oci_identity_domains_approval_workflow_step.test_approval_workflow_step"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsApprovalWorkflowStepResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Optional, acctest.Create, IdentityDomainsApprovalWorkflowStepRepresentation), "identitydomains", "approvalWorkflowStep", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsApprovalWorkflowStepDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowStepResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowStepRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "order", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "escalation"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowStepResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowStepResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Optional, acctest.Create, IdentityDomainsApprovalWorkflowStepRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approvers.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "approvers.0.type", "User"),
				resource.TestCheckResourceAttrSet(resourceName, "approvers.0.value"),
				resource.TestCheckResourceAttr(resourceName, "approvers_expressions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "minimum_approvals", "10"),
				resource.TestCheckResourceAttr(resourceName, "order", "1"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "type", "escalation"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "approvalWorkflowSteps", resId)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_approval_workflow_steps", "test_approval_workflow_steps", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowStepDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApprovalWorkflowStepResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowStepRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_step_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "approval_workflow_steps.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_steps.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowStepSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApprovalWorkflowStepResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approval_workflow_step_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "approvers.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approvers.0.type", "User"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approvers.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "approvers_expressions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "minimum_approvals", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "order", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "escalation"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsApprovalWorkflowStepRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_approval_workflow_step", "approvalWorkflowSteps"),
			ImportStateVerifyIgnore: []string{
				"attribute_sets",
				"attributes",
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				"tags",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsApprovalWorkflowStepDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_approval_workflow_step" {
			noResourceFound = false
			request := oci_identity_domains.GetApprovalWorkflowStepRequest{}

			tmp := rs.Primary.ID
			request.ApprovalWorkflowStepId = &tmp

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

			_, err := client.GetApprovalWorkflowStep(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsApprovalWorkflowStep") {
		resource.AddTestSweepers("IdentityDomainsApprovalWorkflowStep", &resource.Sweeper{
			Name:         "IdentityDomainsApprovalWorkflowStep",
			Dependencies: acctest.DependencyGraph["approvalWorkflowStep"],
			F:            sweepIdentityDomainsApprovalWorkflowStepResource,
		})
	}
}

func sweepIdentityDomainsApprovalWorkflowStepResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	approvalWorkflowStepIds, err := getIdentityDomainsApprovalWorkflowStepIds(compartment)
	if err != nil {
		return err
	}
	for _, approvalWorkflowStepId := range approvalWorkflowStepIds {
		if ok := acctest.SweeperDefaultResourceId[approvalWorkflowStepId]; !ok {
			deleteApprovalWorkflowStepRequest := oci_identity_domains.DeleteApprovalWorkflowStepRequest{}

			deleteApprovalWorkflowStepRequest.ApprovalWorkflowStepId = &approvalWorkflowStepId

			deleteApprovalWorkflowStepRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteApprovalWorkflowStep(context.Background(), deleteApprovalWorkflowStepRequest)
			if error != nil {
				fmt.Printf("Error deleting ApprovalWorkflowStep %s %s, It is possible that the resource is already deleted. Please verify manually \n", approvalWorkflowStepId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsApprovalWorkflowStepIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApprovalWorkflowStepId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listApprovalWorkflowStepsRequest := oci_identity_domains.ListApprovalWorkflowStepsRequest{}
	listApprovalWorkflowStepsResponse, err := identityDomainsClient.ListApprovalWorkflowSteps(context.Background(), listApprovalWorkflowStepsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ApprovalWorkflowStep list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, approvalWorkflowStep := range listApprovalWorkflowStepsResponse.Resources {
		id := *approvalWorkflowStep.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApprovalWorkflowStepId", id)
	}
	return resourceIds, nil
}
