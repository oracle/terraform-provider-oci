// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeployArtifactRequiredOnlyResource = DeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployArtifactRepresentation)

	DeployArtifactResourceConfig = DeployArtifactResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployArtifactRepresentation)

	deployArtifactSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_artifact_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
	}

	deployArtifactDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_artifact.test_deploy_artifact.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: deployArtifactDataSourceFilterRepresentation}}
	deployArtifactDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deploy_artifact.test_deploy_artifact.id}`}},
	}

	deployArtifactRepresentation = map[string]interface{}{
		"argument_substitution_mode": acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `SUBSTITUTE_PLACEHOLDERS`},
		"deploy_artifact_source":     acctest.RepresentationGroup{RepType: acctest.Required, Group: deployArtifactDeployArtifactSourceRepresentation},
		"deploy_artifact_type":       acctest.Representation{RepType: acctest.Required, Create: `KUBERNETES_MANIFEST`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	base64_encode                                    = "YXBpVmVyc2lvbjogYmF0Y2gvdjEKa2luZDogSm9iCm1ldGFkYXRhOgogIGdlbmVyYXRlTmFtZTogaGVsbG93b3JsZAogIGxhYmVsczoKICAgIGFwcDogaGVsbG93b3JsZApzcGVjOgogIHR0bFNlY29uZHNBZnRlckZpbmlzaGVkOiAxMjAKICB0ZW1wbGF0ZToKICAgIHNwZWM6CiAgICAgIGNvbnRhaW5lcnM6CiAgICAgICAgLSBuYW1lOiBoZWxsb3dvcmxkCiAgICAgICAgICBpbWFnZTogcGh4Lm9jaXIuaW8vYXgwMjJ3dmdtanBxL2hlbGxvd29ybGQtb2tlLXZlcmlmaWVyOmxhdGVzdAogICAgICAgICAgY29tbWFuZDoKICAgICAgICAgICAgLSAiL2Jpbi9iYXNoIgogICAgICAgICAgICAtICItYyIKICAgICAgICAgICAgLSAic2xlZXAgMjsgZWNobyBIZWxsbyBXb3JsZDsiCiAgICAgIHJlc3RhcnRQb2xpY3k6IE5ldmVy"
	base64_encode_update                             = "a2luZDogTmFtZXNwYWNlCmFwaVZlcnNpb246IHYxCm1ldGFkYXRhOgogIG5hbWU6IGhlbGxvd29ybGQtZGVtbwotLS0KYXBpVmVyc2lvbjogYXBwcy92MQpraW5kOiBEZXBsb3ltZW50Cm1ldGFkYXRhOgogIG5hbWU6IGhlbGxvd29ybGQtZGVwbG95bWVudAogIG5hbWVzcGFjZTogaGVsbG93b3JsZC1kZW1vCnNwZWM6CiAgc2VsZWN0b3I6CiAgICBtYXRjaExhYmVsczoKICAgICAgYXBwOiBoZWxsb3dvcmxkCiAgcmVwbGljYXM6IDMKICB0ZW1wbGF0ZToKICAgIG1ldGFkYXRhOgogICAgICBsYWJlbHM6CiAgICAgICAgYXBwOiBoZWxsb3dvcmxkCiAgICBzcGVjOgogICAgICBjb250YWluZXJzOgogICAgICAgIC0gbmFtZTogaGVsbG93b3JsZAogICAgICAgICAgIyBlbnRlciB0aGUgcGF0aCB0byB5b3VyIGltYWdlLCBiZSBzdXJlIHRvIGluY2x1ZGUgdGhlIGNvcnJlY3QgcmVnaW9uIHByZWZpeAogICAgICAgICAgaW1hZ2U6IGlhZC5vY2lyLmlvL2F4MDIyd3ZnbWpwcS9oZWxsb3dvcmxkOnYxCiAgICAgICAgICBwb3J0czoKICAgICAgICAgICAgLSBjb250YWluZXJQb3J0OiA4ODg4CiAgICAgICAgICAgICAgcHJvdG9jb2w6IFRDUAoKLS0tCmFwaVZlcnNpb246IHYxCmtpbmQ6IFNlcnZpY2UKbWV0YWRhdGE6CiAgbmFtZTogaGVsbG93b3JsZC1zZXJ2aWNlCiAgbmFtZXNwYWNlOiBoZWxsb3dvcmxkLWRlbW8KICBhbm5vdGF0aW9uczoKICAgIHNlcnZpY2UuYmV0YS5rdWJlcm5ldGVzLmlvL29jaS1sb2FkLWJhbGFuY2VyLXNoYXBlOiAiMTBNYnBzIgpzcGVjOgogIHR5cGU6IExvYWRCYWxhbmNlcgogIHBvcnRzOgogICAgLSBwb3J0OiA4MDgwCiAgICAgIHByb3RvY29sOiBUQ1AKICAgICAgdGFyZ2V0UG9ydDogODg4OAogIHNlbGVjdG9yOgogICAgYXBwOiBoZWxsb3dvcmxk"
	deployArtifactDeployArtifactSourceRepresentation = map[string]interface{}{
		"deploy_artifact_source_type": acctest.Representation{RepType: acctest.Required, Create: `INLINE`},
		"base64encoded_content":       acctest.Representation{RepType: acctest.Required, Create: base64_encode, Update: base64_encode_update},
	}

	DeployArtifactResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, devopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsDeployArtifactResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployArtifactResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_artifact.test_deploy_artifact"
	datasourceName := "data.oci_devops_deploy_artifacts.test_deploy_artifacts"
	singularDatasourceName := "data.oci_devops_deploy_artifact.test_deploy_artifact"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DeployArtifactResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, deployArtifactRepresentation), "devops", "deployArtifact", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployArtifactDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.base64encoded_content", base64_encode),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Create, deployArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.base64encoded_content", base64_encode),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.base64encoded_content", base64_encode_update),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(resourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifacts", "test_deploy_artifacts", acctest.Optional, acctest.Update, deployArtifactDataSourceRepresentation) +
				compartmentIdVariableStr + DeployArtifactResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Optional, acctest.Update, deployArtifactRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_artifact_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deploy_artifact_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_artifact", acctest.Required, acctest.Create, deployArtifactSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployArtifactResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_artifact_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "argument_substitution_mode", "SUBSTITUTE_PLACEHOLDERS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.base64encoded_content", base64_encode_update),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_source.0.deploy_artifact_source_type", "INLINE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_artifact_type", "KUBERNETES_MANIFEST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeployArtifactResourceConfig,
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

func testAccCheckDevopsDeployArtifactDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_deploy_artifact" {
			noResourceFound = false
			request := oci_devops.GetDeployArtifactRequest{}

			tmp := rs.Primary.ID
			request.DeployArtifactId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			response, err := client.GetDeployArtifact(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.DeployArtifactLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DevopsDeployArtifact") {
		resource.AddTestSweepers("DevopsDeployArtifact", &resource.Sweeper{
			Name:         "DevopsDeployArtifact",
			Dependencies: acctest.DependencyGraph["deployArtifact"],
			F:            sweepDevopsDeployArtifactResource,
		})
	}
}

func sweepDevopsDeployArtifactResource(compartment string) error {
	deployArtifactClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	deployArtifactIds, err := getDeployArtifactIds(compartment)
	if err != nil {
		return err
	}
	for _, deployArtifactId := range deployArtifactIds {
		if ok := acctest.SweeperDefaultResourceId[deployArtifactId]; !ok {
			deleteDeployArtifactRequest := oci_devops.DeleteDeployArtifactRequest{}

			deleteDeployArtifactRequest.DeployArtifactId = &deployArtifactId

			deleteDeployArtifactRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := deployArtifactClient.DeleteDeployArtifact(context.Background(), deleteDeployArtifactRequest)
			if error != nil {
				fmt.Printf("Error deleting DeployArtifact %s %s, It is possible that the resource is already deleted. Please verify manually \n", deployArtifactId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &deployArtifactId, deployArtifactSweepWaitCondition, time.Duration(3*time.Minute),
				deployArtifactSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDeployArtifactIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DeployArtifactId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	deployArtifactClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

	listDeployArtifactsRequest := oci_devops.ListDeployArtifactsRequest{}
	listDeployArtifactsRequest.CompartmentId = &compartmentId
	listDeployArtifactsRequest.LifecycleState = oci_devops.DeployArtifactLifecycleStateActive
	listDeployArtifactsResponse, err := deployArtifactClient.ListDeployArtifacts(context.Background(), listDeployArtifactsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DeployArtifact list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployArtifact := range listDeployArtifactsResponse.Items {
		id := *deployArtifact.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DeployArtifactId", id)
	}
	return resourceIds, nil
}

func deployArtifactSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deployArtifactResponse, ok := response.Response.(oci_devops.GetDeployArtifactResponse); ok {
		return deployArtifactResponse.LifecycleState != oci_devops.DeployArtifactLifecycleStateDeleted
	}
	return false
}

func deployArtifactSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DevopsClient().GetDeployArtifact(context.Background(), oci_devops.GetDeployArtifactRequest{
		DeployArtifactId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
