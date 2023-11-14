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
	IdentityDomainsApprovalWorkflowRequiredOnlyResource = IdentityDomainsApprovalWorkflowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowRepresentation)

	IdentityDomainsApprovalWorkflowResourceConfig = IdentityDomainsApprovalWorkflowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowRepresentation)

	IdentityDomainsApprovalWorkflowSingularDataSourceRepresentation = map[string]interface{}{
		"approval_workflow_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_approval_workflow.test_approval_workflow.id}`},
		"idcs_endpoint":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":       acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsApprovalWorkflowDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"approval_workflow_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"approval_workflow_filter": acctest.Representation{RepType: acctest.Optional, Create: `id eq \"${oci_identity_domains_approval_workflow.test_approval_workflow.id}\"`},
		"attribute_sets":           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsApprovalWorkflowRepresentation = map[string]interface{}{
		"idcs_endpoint":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"max_duration":            acctest.RepresentationGroup{RepType: acctest.Required, Group: IdentityDomainsApprovalWorkflowMaxDurationRepresentation},
		"name":                    acctest.Representation{RepType: acctest.Required, Create: `name`},
		"schemas":                 acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:ApprovalWorkflow`}},
		"approval_workflow_steps": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsApprovalWorkflowApprovalWorkflowStepsRepresentation},
		"attribute_sets":          acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"tags":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityDomainsApprovalWorkflowTagsRepresentation},
	}
	IdentityDomainsApprovalWorkflowMaxDurationRepresentation = map[string]interface{}{
		"unit":  acctest.Representation{RepType: acctest.Required, Create: `MONTH`, Update: `WEEK`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
	}

	IdentityDomainsApprovalWorkflowApprovalWorkflowStepsRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_approval_workflow_step.test_approval_workflow_step.type}`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_approval_workflow_step.test_approval_workflow_step.id}`},
	}
	IdentityDomainsApprovalWorkflowTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	IdentityDomainsApprovalWorkflowResourceDependencies = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow_step", "test_approval_workflow_step", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowStepRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_user", "test_user", acctest.Required, acctest.Create, IdentityDomainsUserRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsApprovalWorkflowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsApprovalWorkflowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_approval_workflow.test_approval_workflow"
	datasourceName := "data.oci_identity_domains_approval_workflows.test_approval_workflows"
	singularDatasourceName := "data.oci_identity_domains_approval_workflow.test_approval_workflow"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsApprovalWorkflowResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Optional, acctest.Create, IdentityDomainsApprovalWorkflowRepresentation), "identitydomains", "approvalWorkflow", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsApprovalWorkflowDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.0.unit", "MONTH"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.0.value", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Optional, acctest.Create, IdentityDomainsApprovalWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approval_workflow_steps.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow_steps.0.type"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow_steps.0.value"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.0.unit", "MONTH"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.0.value", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
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

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "approvalWorkflows", resId)
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsApprovalWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// TODO: implement ApprovalWorkflowSteps dependency
				resource.TestCheckResourceAttr(resourceName, "approval_workflow_steps.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow_steps.0.type"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow_steps.0.value"),
				resource.TestCheckResourceAttr(resourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.0.unit", "WEEK"),
				resource.TestCheckResourceAttr(resourceName, "max_duration.0.value", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "tags.0.value", "value2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_approval_workflows", "test_approval_workflows", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApprovalWorkflowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Optional, acctest.Update, IdentityDomainsApprovalWorkflowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "approval_workflow_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "approval_workflows.#", regexp.MustCompile("[1-9][0-9]*")),
				resource.TestCheckResourceAttr(datasourceName, "approval_workflows.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_approval_workflow", "test_approval_workflow", acctest.Required, acctest.Create, IdentityDomainsApprovalWorkflowSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsApprovalWorkflowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approval_workflow_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(resourceName, "approval_workflow_steps.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow_steps.0.type"),
				resource.TestCheckResourceAttrSet(resourceName, "approval_workflow_steps.0.value"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_duration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_duration.0.unit", "WEEK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_duration.0.value", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsApprovalWorkflowRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_approval_workflow", "approvalWorkflows"),
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

func testAccCheckIdentityDomainsApprovalWorkflowDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_approval_workflow" {
			noResourceFound = false
			request := oci_identity_domains.GetApprovalWorkflowRequest{}

			tmp := rs.Primary.ID
			request.ApprovalWorkflowId = &tmp

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

			_, err := client.GetApprovalWorkflow(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("IdentityDomainsApprovalWorkflow") {
		resource.AddTestSweepers("IdentityDomainsApprovalWorkflow", &resource.Sweeper{
			Name:         "IdentityDomainsApprovalWorkflow",
			Dependencies: acctest.DependencyGraph["approvalWorkflow"],
			F:            sweepIdentityDomainsApprovalWorkflowResource,
		})
	}
}

func sweepIdentityDomainsApprovalWorkflowResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	approvalWorkflowIds, err := getIdentityDomainsApprovalWorkflowIds(compartment)
	if err != nil {
		return err
	}
	for _, approvalWorkflowId := range approvalWorkflowIds {
		if ok := acctest.SweeperDefaultResourceId[approvalWorkflowId]; !ok {
			deleteApprovalWorkflowRequest := oci_identity_domains.DeleteApprovalWorkflowRequest{}

			deleteApprovalWorkflowRequest.ApprovalWorkflowId = &approvalWorkflowId

			deleteApprovalWorkflowRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteApprovalWorkflow(context.Background(), deleteApprovalWorkflowRequest)
			if error != nil {
				fmt.Printf("Error deleting ApprovalWorkflow %s %s, It is possible that the resource is already deleted. Please verify manually \n", approvalWorkflowId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsApprovalWorkflowIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ApprovalWorkflowId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listApprovalWorkflowsRequest := oci_identity_domains.ListApprovalWorkflowsRequest{}
	listApprovalWorkflowsResponse, err := identityDomainsClient.ListApprovalWorkflows(context.Background(), listApprovalWorkflowsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ApprovalWorkflow list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, approvalWorkflow := range listApprovalWorkflowsResponse.Resources {
		id := *approvalWorkflow.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ApprovalWorkflowId", id)
	}
	return resourceIds, nil
}
