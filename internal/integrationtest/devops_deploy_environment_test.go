// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsDeployEnvironmentRequiredOnlyResource = DevopsDeployEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, DevopsdeployEnvironmentRepresentation)

	DevopsDeployEnvironmentResourceConfig = DevopsDeployEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, DevopsdeployEnvironmentRepresentation)

	DevopsDevopsDeployEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
	}

	DevopsDevopsDeployEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsDeployEnvironmentDataSourceFilterRepresentation}}
	DevopsDeployEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_deploy_environment.test_deploy_environment.id}`}},
	}

	cluster_fake_id                       = "ocid1.cluster.oc1.us-ashburn-1.aaaaaaaaafqtkm3fg4zwgnlggmywkzdemi2dcyzymfrdqojygcstocluster1"
	DevopsdeployEnvironmentRepresentation = map[string]interface{}{
		"deploy_environment_type": acctest.Representation{RepType: acctest.Required, Create: `OKE_CLUSTER`},
		"project_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"cluster_id":              acctest.Representation{RepType: acctest.Required, Create: cluster_fake_id},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	DevopsDeployEnvironmentResourceDependencies = AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", acctest.Required, acctest.Create, DevopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: devops/default
func TestDevopsDeployEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_environment.test_deploy_environment"
	datasourceName := "data.oci_devops_deploy_environments.test_deploy_environments"
	singularDatasourceName := "data.oci_devops_deploy_environment.test_deploy_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsDeployEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, DevopsdeployEnvironmentRepresentation), "devops", "deployEnvironment", t)

	acctest.ResourceTest(t, testAccCheckDevopsDeployEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, DevopsdeployEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Create, DevopsdeployEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
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
			Config: config + compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, DevopsdeployEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environments", "test_deploy_environments", acctest.Optional, acctest.Update, DevopsDevopsDeployEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeployEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Optional, acctest.Update, DevopsdeployEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_environment_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", acctest.Required, acctest.Create, DevopsDevopsDeployEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsDeployEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_environment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_environment_type", "OKE_CLUSTER"),
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
		// verify resource import
		{
			Config:                  config + DevopsDeployEnvironmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDevopsDeployEnvironmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DevopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_deploy_environment" {
			noResourceFound = false
			request := oci_devops.GetDeployEnvironmentRequest{}

			tmp := rs.Primary.ID
			request.DeployEnvironmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")

			response, err := client.GetDeployEnvironment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.DeployEnvironmentLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("DevopsDeployEnvironment") {
		resource.AddTestSweepers("DevopsDeployEnvironment", &resource.Sweeper{
			Name:         "DevopsDeployEnvironment",
			Dependencies: acctest.DependencyGraph["deployEnvironment"],
			F:            sweepDevopsDeployEnvironmentResource,
		})
	}
}

func sweepDevopsDeployEnvironmentResource(compartment string) error {
	deployEnvironmentClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()
	deployEnvironmentIds, err := getDevopsDeployEnvironmentIds(compartment)
	if err != nil {
		return err
	}
	for _, deployEnvironmentId := range deployEnvironmentIds {
		if ok := acctest.SweeperDefaultResourceId[deployEnvironmentId]; !ok {
			deleteDeployEnvironmentRequest := oci_devops.DeleteDeployEnvironmentRequest{}

			deleteDeployEnvironmentRequest.DeployEnvironmentId = &deployEnvironmentId

			deleteDeployEnvironmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			_, error := deployEnvironmentClient.DeleteDeployEnvironment(context.Background(), deleteDeployEnvironmentRequest)
			if error != nil {
				fmt.Printf("Error deleting DeployEnvironment %s %s, It is possible that the resource is already deleted. Please verify manually \n", deployEnvironmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &deployEnvironmentId, DevopsDeployEnvironmentSweepWaitCondition, time.Duration(3*time.Minute),
				DevopsDeployEnvironmentSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDevopsDeployEnvironmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DeployEnvironmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	deployEnvironmentClient := acctest.GetTestClients(&schema.ResourceData{}).DevopsClient()

	listDeployEnvironmentsRequest := oci_devops.ListDeployEnvironmentsRequest{}
	listDeployEnvironmentsRequest.CompartmentId = &compartmentId
	listDeployEnvironmentsRequest.LifecycleState = oci_devops.DeployEnvironmentLifecycleStateActive
	listDeployEnvironmentsResponse, err := deployEnvironmentClient.ListDeployEnvironments(context.Background(), listDeployEnvironmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DeployEnvironment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployEnvironment := range listDeployEnvironmentsResponse.Items {
		id := *deployEnvironment.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DeployEnvironmentId", id)
	}
	return resourceIds, nil
}

func DevopsDeployEnvironmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deployEnvironmentResponse, ok := response.Response.(oci_devops.GetDeployEnvironmentResponse); ok {
		return deployEnvironmentResponse.GetLifecycleState() != oci_devops.DeployEnvironmentLifecycleStateDeleted
	}
	return false
}

func DevopsDeployEnvironmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DevopsClient().GetDeployEnvironment(context.Background(), oci_devops.GetDeployEnvironmentRequest{
		DeployEnvironmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
