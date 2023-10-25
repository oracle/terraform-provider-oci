// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreRemediationRunDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	AdmRemediationRunRequiredOnlyResource = AdmRemediationRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Required, acctest.Create, AdmRemediationRunRepresentation)
	AdmRemediationRunResourceConfig = AdmRemediationRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Optional, acctest.Update, AdmRemediationRunRepresentation)

	AdmRemediationRunSingularDataSourceRepresentation = map[string]interface{}{
		"remediation_run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_adm_remediation_run.test_remediation_run.id}`},
	}

	AdmRemediationRunDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_adm_remediation_run.test_remediation_run.id}`},
		"remediation_recipe_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_adm_remediation_recipe.test_remediation_recipe.id}`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: AdmRemediationRunDataSourceFilterRepresentation}}
	AdmRemediationRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_adm_remediation_run.test_remediation_run.id}`}},
	}

	AdmRemediationRunRepresentation = map[string]interface{}{
		"remediation_recipe_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_adm_remediation_recipe.test_remediation_recipe.id}`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRemediationRunDefinedTagsChangesRepresentation},
	}

	AdmRemediationRunResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, knowledgeBaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Required, acctest.Create, AdmRemediationRecipeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: adm/default
func TestAdmRemediationRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAdmRemediationRunResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	secretId := utils.GetEnvSettingWithBlankDefault("kms_secret_ocid")
	secretIdVariableStr := fmt.Sprintf("variable \"kms_secret_ocid\" { default = \"%s\" }\n", secretId)

	codeRepositoryId := utils.GetEnvSettingWithBlankDefault("devops_code_repository_ocid")
	codeRepositoryIdStr := fmt.Sprintf("variable \"devops_code_repository_ocid\" { default = \"%s\" }\n", codeRepositoryId)

	pipelineId := utils.GetEnvSettingWithBlankDefault("devops_build_pipeline_ocid")
	pipelineIdStr := fmt.Sprintf("variable \"devops_build_pipeline_ocid\" { default = \"%s\" }\n", pipelineId)

	resourceName := "oci_adm_remediation_run.test_remediation_run"
	datasourceName := "data.oci_adm_remediation_runs.test_remediation_runs"
	singularDatasourceName := "data.oci_adm_remediation_run.test_remediation_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+AdmRemediationRunResourceDependencies+secretIdVariableStr+codeRepositoryIdStr+pipelineIdStr+
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(AdmRemediationRunRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "adm", "remediationRun", t)

	acctest.ResourceTest(t, testAccCheckAdmRemediationRunDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + AdmRemediationRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Required, acctest.Create, AdmRemediationRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "remediation_recipe_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + AdmRemediationRunResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + AdmRemediationRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AdmRemediationRunRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remediation_recipe_id"),
				resource.TestCheckResourceAttrSet(resourceName, "remediation_run_source"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr + AdmRemediationRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(AdmRemediationRunRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remediation_recipe_id"),
				resource.TestCheckResourceAttrSet(resourceName, "remediation_run_source"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_adm_remediation_runs", "test_remediation_runs", acctest.Optional, acctest.Update, AdmRemediationRunDataSourceRepresentation) +
				compartmentIdVariableStr + AdmRemediationRunResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Optional, acctest.Update, AdmRemediationRunRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "remediation_recipe_id"),

				resource.TestCheckResourceAttr(datasourceName, "remediation_run_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "remediation_run_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Required, acctest.Create, AdmRemediationRunSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AdmRemediationRunResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remediation_run_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remediation_run_source"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stages.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AdmRemediationRunRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"state", "summary", "time_finished", "time_updated"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAdmRemediationRunDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApplicationDependencyManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_adm_remediation_run" {
			noResourceFound = false
			request := oci_adm.GetRemediationRunRequest{}

			tmp := rs.Primary.ID
			request.RemediationRunId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "adm")

			response, err := client.GetRemediationRun(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_adm.RemediationRunLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AdmRemediationRun") {
		resource.AddTestSweepers("AdmRemediationRun", &resource.Sweeper{
			Name:         "AdmRemediationRun",
			Dependencies: acctest.DependencyGraph["remediationRun"],
			F:            sweepAdmRemediationRunResource,
		})
	}
}

func sweepAdmRemediationRunResource(compartment string) error {
	applicationDependencyManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ApplicationDependencyManagementClient()
	remediationRunIds, err := getAdmRemediationRunIds(compartment)
	if err != nil {
		return err
	}
	for _, remediationRunId := range remediationRunIds {
		if ok := acctest.SweeperDefaultResourceId[remediationRunId]; !ok {
			deleteRemediationRunRequest := oci_adm.DeleteRemediationRunRequest{}

			deleteRemediationRunRequest.RemediationRunId = &remediationRunId

			deleteRemediationRunRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "adm")
			_, error := applicationDependencyManagementClient.DeleteRemediationRun(context.Background(), deleteRemediationRunRequest)
			if error != nil {
				fmt.Printf("Error deleting RemediationRun %s %s, It is possible that the resource is already deleted. Please verify manually \n", remediationRunId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &remediationRunId, AdmRemediationRunSweepWaitCondition, time.Duration(3*time.Minute),
				AdmRemediationRunSweepResponseFetchOperation, "adm", true)
		}
	}
	return nil
}

func getAdmRemediationRunIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RemediationRunId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	applicationDependencyManagementClient := acctest.GetTestClients(&schema.ResourceData{}).ApplicationDependencyManagementClient()

	listRemediationRunsRequest := oci_adm.ListRemediationRunsRequest{}
	listRemediationRunsRequest.CompartmentId = &compartmentId
	listRemediationRunsResponse, err := applicationDependencyManagementClient.ListRemediationRuns(context.Background(), listRemediationRunsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RemediationRun list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, remediationRun := range listRemediationRunsResponse.Items {
		id := *remediationRun.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RemediationRunId", id)
	}
	return resourceIds, nil
}

func AdmRemediationRunSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if remediationRunResponse, ok := response.Response.(oci_adm.GetRemediationRunResponse); ok {
		return remediationRunResponse.LifecycleState != oci_adm.RemediationRunLifecycleStateDeleted
	}
	return false
}

func AdmRemediationRunSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ApplicationDependencyManagementClient().GetRemediationRun(context.Background(), oci_adm.GetRemediationRunRequest{
		RemediationRunId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
