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
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineClusterNamespaceProfileRequiredOnlyResource = ContainerengineClusterNamespaceProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileRepresentation)

	ContainerengineClusterNamespaceProfileResourceConfig = ContainerengineClusterNamespaceProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileRepresentation)

	ContainerengineClusterNamespaceProfileSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_namespace_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id}`},
	}

	ContainerengineClusterNamespaceProfileDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterNamespaceProfileDataSourceFilterRepresentation}}
	ContainerengineClusterNamespaceProfileDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id}`}},
	}

	ContainerengineClusterNamespaceProfileRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"namespace_suffix": acctest.Representation{RepType: acctest.Optional, Create: `namespacesuffix`},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreContainerengineClusterNamespaceProfileDefinedTagsChangesRepresentation},
	}

	ContainerengineClusterNamespaceProfileResourceDependencies = ""

	ignoreContainerengineClusterNamespaceProfileDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterNamespaceProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterNamespaceProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile"
	datasourceName := "data.oci_containerengine_cluster_namespace_profiles.test_cluster_namespace_profiles"
	singularDatasourceName := "data.oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterNamespaceProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Optional, acctest.Create, ContainerengineClusterNamespaceProfileRepresentation), "containerengine", "clusterNamespaceProfile", t)

	acctest.ResourceTest(t, testAccCheckContainerengineClusterNamespaceProfileDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Optional, acctest.Create, ContainerengineClusterNamespaceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "namespace_suffix", "namespacesuffix"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ContainerengineClusterNamespaceProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerengineClusterNamespaceProfileRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "namespace_suffix", "namespacesuffix"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "namespace_suffix", "namespacesuffix"),
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_namespace_profiles", "test_cluster_namespace_profiles", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterNamespaceProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "cluster_namespace_profile_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_namespace_profile_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile", "test_cluster_namespace_profile", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterNamespaceProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_namespace_profile_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace_suffix", "namespacesuffix"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerengineClusterNamespaceProfileRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineClusterNamespaceProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster_namespace_profile" {
			noResourceFound = false
			request := oci_containerengine.GetClusterNamespaceProfileRequest{}

			tmp := rs.Primary.ID
			request.ClusterNamespaceProfileId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetClusterNamespaceProfile(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.ClusterNamespaceProfileLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ContainerengineClusterNamespaceProfile") {
		resource.AddTestSweepers("ContainerengineClusterNamespaceProfile", &resource.Sweeper{
			Name:         "ContainerengineClusterNamespaceProfile",
			Dependencies: acctest.DependencyGraph["clusterNamespaceProfile"],
			F:            sweepContainerengineClusterNamespaceProfileResource,
		})
	}
}

func sweepContainerengineClusterNamespaceProfileResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	clusterNamespaceProfileIds, err := getContainerengineClusterNamespaceProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterNamespaceProfileId := range clusterNamespaceProfileIds {
		if ok := acctest.SweeperDefaultResourceId[clusterNamespaceProfileId]; !ok {
			deleteClusterNamespaceProfileRequest := oci_containerengine.DeleteClusterNamespaceProfileRequest{}

			deleteClusterNamespaceProfileRequest.ClusterNamespaceProfileId = &clusterNamespaceProfileId

			deleteClusterNamespaceProfileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteClusterNamespaceProfile(context.Background(), deleteClusterNamespaceProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterNamespaceProfile %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterNamespaceProfileId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterNamespaceProfileId, ContainerengineClusterNamespaceProfileSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerengineClusterNamespaceProfileSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineClusterNamespaceProfileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterNamespaceProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listClusterNamespaceProfilesRequest := oci_containerengine.ListClusterNamespaceProfilesRequest{}
	listClusterNamespaceProfilesRequest.CompartmentId = &compartmentId
	listClusterNamespaceProfilesRequest.LifecycleState = oci_containerengine.ClusterNamespaceProfileLifecycleStateActive
	listClusterNamespaceProfilesResponse, err := containerEngineClient.ListClusterNamespaceProfiles(context.Background(), listClusterNamespaceProfilesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ClusterNamespaceProfile list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, clusterNamespaceProfile := range listClusterNamespaceProfilesResponse.Items {
		id := *clusterNamespaceProfile.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterNamespaceProfileId", id)
	}
	return resourceIds, nil
}

func ContainerengineClusterNamespaceProfileSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterNamespaceProfileResponse, ok := response.Response.(oci_containerengine.GetClusterNamespaceProfileResponse); ok {
		return clusterNamespaceProfileResponse.LifecycleState != oci_containerengine.ClusterNamespaceProfileLifecycleStateDeleted
	}
	return false
}

func ContainerengineClusterNamespaceProfileSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetClusterNamespaceProfile(context.Background(), oci_containerengine.GetClusterNamespaceProfileRequest{
		ClusterNamespaceProfileId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
