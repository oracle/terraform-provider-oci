// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	/*
		this addon tests require existence of KubernetsDashboard and its configuration. re-evaluate this when condition change.
		alternative way would be to parse response of addon_option data source, find a non-essential add-on that support replicas configuration
	*/
	addonName              = "KubernetesDashboard"
	addonConfigKey         = "numOfReplicas"
	addonConfigValue       = "1"
	addonConfigValueUpdate = "2"

	essentialAddonName        = "CoreDNS"
	essentialAddonConfigKey   = "minReplica"
	essentialAddonConfigValue = "4"

	ContainerengineAddonSingularDataSourceRepresentation = map[string]interface{}{
		"addon_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_addon.test_addon.addon_name}`},
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}

	ContainerengineAddonDataSourceRepresentation = map[string]interface{}{
		"cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}

	ContainerengineAddonDataSource = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_addons", "test_addons", acctest.Required, acctest.Create, ContainerengineAddonDataSourceRepresentation)

	ContainerengineAddonSingularDataSource = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_addon", "test_addon", acctest.Required, acctest.Create, ContainerengineAddonSingularDataSourceRepresentation)

	ContainerengineAddonRepresentation = map[string]interface{}{
		"cluster_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"addon_name":                       acctest.Representation{RepType: acctest.Required, Create: addonName},
		"remove_addon_resources_on_delete": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"configurations":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineAddonConfigurationsRepresentation},
		"version":                          acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `${data.oci_containerengine_addon_options.adddon_options_dashboard.addon_options[0].versions[0].version_number}`},
	}

	ContainerengineEssentialAddonRepresentation = map[string]interface{}{
		"cluster_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
		"addon_name":                       acctest.Representation{RepType: acctest.Required, Create: essentialAddonName},
		"remove_addon_resources_on_delete": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"override_existing":                acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"configurations":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineEssentialAddonConfigurationsRepresentation},
	}

	ContainerengineAddonConfigurationsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: addonConfigKey, Update: addonConfigKey},
		"value": acctest.Representation{RepType: acctest.Optional, Create: addonConfigValue, Update: addonConfigValueUpdate},
	}

	ContainerengineEssentialAddonConfigurationsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: essentialAddonConfigKey},
		"value": acctest.Representation{RepType: acctest.Optional, Create: essentialAddonConfigValue},
	}

	ContainerengineAddonRequiredOnlyResourceCreate = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_addon", "test_addon", acctest.Required, acctest.Create, ContainerengineAddonRepresentation)

	ContainerengineAddonOptionalResourceCreate = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_addon", "test_addon", acctest.Optional, acctest.Create, ContainerengineAddonRepresentation)

	ContainerengineAddonOptionalResourceConfigUpdate = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_addon", "test_addon", acctest.Optional, acctest.Update, ContainerengineAddonRepresentation)

	ContainerengineEssentialAddonResourceCreate = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_addon", "test_essential_addon", acctest.Optional, acctest.Create, ContainerengineEssentialAddonRepresentation)

	AddonOptionDashboardDataSourceRepresentation = map[string]interface{}{
		"kubernetes_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions[length(data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions)-2]}`},
		"addon_name":         acctest.Representation{RepType: acctest.Optional, Create: addonName},
	}

	ContainerengineAddonResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(ContainerengineClusterRepresentation, map[string]interface{}{
			"type": acctest.Representation{RepType: acctest.Required, Create: `ENHANCED_CLUSTER`, Update: `ENHANCED_CLUSTER`},
			//"cluster_pod_network_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: clusterClusterPodNetworkOptionsRepresentation},
			"endpoint_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterEndpointConfigRepresentation},
		})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_option", "test_cluster_option", acctest.Required, acctest.Create, ContainerengineContainerengineClusterOptionSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_addon_options", "adddon_options_dashboard", acctest.Optional, acctest.Create, AddonOptionDashboardDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: containerengine/default
func TestContainerengineAddonResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineAddonResource_basic")
	defer httpreplay.SaveScenario()

	fmt.Printf("ContainerengineEssentialAddonResourceCreate: %v", ContainerengineEssentialAddonResourceCreate)

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_addon.test_addon"
	essentialAddonResourceName := "oci_containerengine_addon.test_essential_addon"
	datasourceName := "data.oci_containerengine_addons.test_addons"
	singularDatasourceName := "data.oci_containerengine_addon.test_addon"

	baseConfig := config + compartmentIdVariableStr + ContainerengineAddonResourceDependencies
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(baseConfig+ContainerengineAddonOptionalResourceCreate,
		"containerengine", "addon", t)

	acctest.ResourceTest(t, testAccCheckContainerengineAddonDestroy, []resource.TestStep{
		// verify Create
		{
			Config: baseConfig + ContainerengineAddonRequiredOnlyResourceCreate,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "addon_name", addonName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: baseConfig,
		},
		// verify Create with optionals
		{
			Config: baseConfig + ContainerengineAddonOptionalResourceCreate,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.key", addonConfigKey),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.value", addonConfigValue),
				resource.TestCheckResourceAttrSet(resourceName, "current_installed_version"),
				resource.TestCheckResourceAttr(resourceName, "addon_name", addonName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "current_installed_version"),

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
			Config: baseConfig + ContainerengineAddonOptionalResourceConfigUpdate,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.key", addonConfigKey),
				resource.TestCheckResourceAttr(resourceName, "configurations.0.value", addonConfigValueUpdate),
				resource.TestCheckResourceAttrSet(resourceName, "current_installed_version"),
				resource.TestCheckResourceAttr(resourceName, "addon_name", addonName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),
				resource.TestCheckResourceAttrSet(resourceName, "current_installed_version"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated. %s, %s", resId, resId2)
					}
					return err
				},
			),
		},
		// verify update-on-install of an essential addon
		{
			Config: baseConfig + ContainerengineAddonOptionalResourceConfigUpdate + ContainerengineEssentialAddonResourceCreate,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(essentialAddonResourceName, "cluster_id"),
				resource.TestCheckResourceAttr(essentialAddonResourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(essentialAddonResourceName, "configurations.0.key", essentialAddonConfigKey),
				resource.TestCheckResourceAttr(essentialAddonResourceName, "configurations.0.value", essentialAddonConfigValue),
				resource.TestCheckResourceAttrSet(essentialAddonResourceName, "current_installed_version"),
				resource.TestCheckResourceAttr(essentialAddonResourceName, "addon_name", essentialAddonName),
				resource.TestCheckResourceAttrSet(essentialAddonResourceName, "state"),
			),
		},
		// verify datasource
		{
			Config: baseConfig + ContainerengineAddonDataSource + ContainerengineAddonOptionalResourceConfigUpdate,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "addons.0.addon_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "addons.0.current_installed_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "addons.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "addons.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: baseConfig + ContainerengineAddonSingularDataSource + ContainerengineAddonOptionalResourceConfigUpdate,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.key", addonConfigKey),
				resource.TestCheckResourceAttr(singularDatasourceName, "configurations.0.value", addonConfigValueUpdate),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_installed_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addon_name", addonName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  baseConfig + ContainerengineAddonRequiredOnlyResourceCreate,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"remove_addon_resources_on_delete", "override_existing"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineAddonDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_addon" && rs.Primary.Attributes["addon_name"] != essentialAddonName {
			noResourceFound = false
			request := oci_containerengine.GetAddonRequest{}

			if value, ok := rs.Primary.Attributes["addon_name"]; ok {
				request.AddonName = &value
			}

			if value, ok := rs.Primary.Attributes["cluster_id"]; ok {
				request.ClusterId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetAddon(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.AddonLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ContainerengineAddon") {
		resource.AddTestSweepers("ContainerengineAddon", &resource.Sweeper{
			Name:         "ContainerengineAddon",
			Dependencies: acctest.DependencyGraph["addon"],
			F:            sweepContainerengineAddonResource,
		})
	}
}

func sweepContainerengineAddonResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	addonIds, err := getContainerengineAddonIds(compartment)
	if err != nil {
		return err
	}
	for _, addonId := range addonIds {
		if ok := acctest.SweeperDefaultResourceId[addonId]; !ok {
			idArr := strings.Split(addonId, ".")
			if len(idArr) != 2 {
				return fmt.Errorf("invalid addonId %s", addonId)
			}

			disableAddonRequest := oci_containerengine.DisableAddonRequest{}
			disableAddonRequest.ClusterId = &idArr[0]
			disableAddonRequest.AddonName = &idArr[1]
			disableAddonRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DisableAddon(context.Background(), disableAddonRequest)
			if error != nil {
				fmt.Printf("Error deleting Addon %s %s, It is possible that the resource is already deleted. Please verify manually \n", addonId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &addonId, ContainerengineAddonSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerengineAddonSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineAddonIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AddonId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listAddonsRequest := oci_containerengine.ListAddonsRequest{}

	clusterIds, error := getClusterIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting clusterId required for Addon resource requests \n")
	}
	for _, clusterId := range clusterIds {
		listAddonsRequest.ClusterId = &clusterId

		listAddonsResponse, err := containerEngineClient.ListAddons(context.Background(), listAddonsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Addon list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, addon := range listAddonsResponse.Items {
			id := *addon.Name
			resourceIds = append(resourceIds, fmt.Sprintf("%s.%s", clusterId, id))
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AddonId", id)
		}

	}
	return resourceIds, nil
}

func getClusterIds(compartment string) ([]string, error) {
	var resourceIds []string
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listClusterRequest := oci_containerengine.ListClustersRequest{}
	listClusterRequest.CompartmentId = &compartment

	listClusterReponse, err := containerEngineClient.ListClusters(context.Background(), listClusterRequest)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting cluster list for compartment id: %s, %s \n", compartment, err)
	}

	for _, cluster := range listClusterReponse.Items {
		resourceIds = append(resourceIds, *cluster.Id)
	}

	return resourceIds, nil
}

func ContainerengineAddonSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if addonResponse, ok := response.Response.(oci_containerengine.GetAddonResponse); ok {
		return addonResponse.LifecycleState != oci_containerengine.AddonLifecycleStateDeleted
	}
	return false
}

func ContainerengineAddonSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetAddon(context.Background(), oci_containerengine.GetAddonRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
