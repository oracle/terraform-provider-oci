// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	CrossConnectGroupRequiredOnlyResource = CrossConnectGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Required, acctest.Create, crossConnectGroupRepresentation)

	CrossConnectGroupResourceConfig = CrossConnectGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Update, crossConnectGroupRepresentation)

	CrossConnectGroupResourceConfigCopyForVC = CrossConnectGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Update, crossConnectGroupRepresentationCopyForVC)

	crossConnectGroupSingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
	}

	crossConnectGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: crossConnectGroupDataSourceFilterRepresentation}}
	crossConnectGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_cross_connect_group.test_cross_connect_group.id}`}},
	}

	crossConnectGroupRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"customer_reference_name": acctest.Representation{RepType: acctest.Optional, Create: `customerReferenceName`, Update: `customerReferenceName2`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"macsec_properties":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: crossConnectGroupMacsecPropertiesRepresentation},
	}

	crossConnectGroupRepresentationCopyForVC = acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{"macsec_properties"}, crossConnectGroupRepresentation)

	crossConnectGroupMacsecPropertiesRepresentation = map[string]interface{}{
		"state":             acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `ENABLED`},
		"encryption_cipher": acctest.Representation{RepType: acctest.Optional, Create: `AES256_GCM`, Update: `AES256_GCM_XPN`},
		"primary_key":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: crossConnectGroupMacsecPropertiesPrimaryKeyRepresentation},
	}
	crossConnectGroupMacsecPropertiesPrimaryKeyRepresentation = map[string]interface{}{
		"connectivity_association_key_secret_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_ckn}`},
		"connectivity_association_name_secret_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.secret_ocid_cak}`},
		"connectivity_association_key_secret_version":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_version_cak}`},
		"connectivity_association_name_secret_version": acctest.Representation{RepType: acctest.Required, Create: `${var.secret_version_ckn}`},
	}

	CrossConnectGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestCoreCrossConnectGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	secretIdCKN := utils.GetEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := utils.GetEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	secretVersionCAK := utils.GetEnvSettingWithBlankDefault("secret_version_cak")
	secretVersionStrCAK := fmt.Sprintf("variable \"secret_version_cak\" { default = \"%s\" }\n", secretVersionCAK)

	secretVersionCKN := utils.GetEnvSettingWithBlankDefault("secret_version_ckn")
	secretVersionStrCKN := fmt.Sprintf("variable \"secret_version_ckn\" { default = \"%s\" }\n", secretVersionCKN)

	resourceName := "oci_core_cross_connect_group.test_cross_connect_group"
	datasourceName := "data.oci_core_cross_connect_groups.test_cross_connect_groups"
	singularDatasourceName := "data.oci_core_cross_connect_group.test_cross_connect_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CrossConnectGroupResourceDependencies+secretIdVariableStrCKN+secretIdVariableStrCAK+secretVersionStrCKN+secretVersionStrCAK+
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Create, crossConnectGroupRepresentation), "core", "crossConnectGroup", t)

	acctest.ResourceTest(t, testAccCheckCoreCrossConnectGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Required, acctest.Create, crossConnectGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK +
				secretVersionStrCKN + secretVersionStrCAK +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Create, crossConnectGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_version"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_version"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK +
				secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(crossConnectGroupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),

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
			Config: config + compartmentIdVariableStr + CrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Update, crossConnectGroupRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customer_reference_name", "customerReferenceName2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "macsec_properties.0.state", "ENABLED"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_groups", "test_cross_connect_groups", acctest.Optional, acctest.Update, crossConnectGroupDataSourceRepresentation) +
				compartmentIdVariableStr + CrossConnectGroupResourceDependencies + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Optional, acctest.Update, crossConnectGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.customer_reference_name", "customerReferenceName2"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.macsec_properties.0.primary_key.0.connectivity_association_key_secret_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.macsec_properties.0.primary_key.0.connectivity_association_key_secret_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.macsec_properties.0.primary_key.0.connectivity_association_name_secret_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.macsec_properties.0.primary_key.0.connectivity_association_name_secret_version"),
				resource.TestCheckResourceAttr(datasourceName, "cross_connect_groups.0.macsec_properties.0.state", "ENABLED"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_groups.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_group", "test_cross_connect_group", acctest.Required, acctest.Create, crossConnectGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CrossConnectGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_group_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_reference_name", "customerReferenceName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.0.encryption_cipher", "AES256_GCM_XPN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.0.primary_key.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.primary_key.0.connectivity_association_key_secret_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "macsec_properties.0.primary_key.0.connectivity_association_name_secret_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "macsec_properties.0.state", "ENABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + CrossConnectGroupResourceConfig + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
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

func testAccCheckCoreCrossConnectGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cross_connect_group" {
			noResourceFound = false
			request := oci_core.GetCrossConnectGroupRequest{}

			tmp := rs.Primary.ID
			request.CrossConnectGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetCrossConnectGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.CrossConnectGroupLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreCrossConnectGroup") {
		resource.AddTestSweepers("CoreCrossConnectGroup", &resource.Sweeper{
			Name:         "CoreCrossConnectGroup",
			Dependencies: acctest.DependencyGraph["crossConnectGroup"],
			F:            sweepCoreCrossConnectGroupResource,
		})
	}
}

func sweepCoreCrossConnectGroupResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	crossConnectGroupIds, err := getCrossConnectGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, crossConnectGroupId := range crossConnectGroupIds {
		if ok := acctest.SweeperDefaultResourceId[crossConnectGroupId]; !ok {
			deleteCrossConnectGroupRequest := oci_core.DeleteCrossConnectGroupRequest{}

			deleteCrossConnectGroupRequest.CrossConnectGroupId = &crossConnectGroupId

			deleteCrossConnectGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteCrossConnectGroup(context.Background(), deleteCrossConnectGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting CrossConnectGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", crossConnectGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &crossConnectGroupId, crossConnectGroupSweepWaitCondition, time.Duration(3*time.Minute),
				crossConnectGroupSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCrossConnectGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CrossConnectGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listCrossConnectGroupsRequest := oci_core.ListCrossConnectGroupsRequest{}
	listCrossConnectGroupsRequest.CompartmentId = &compartmentId
	listCrossConnectGroupsRequest.LifecycleState = oci_core.CrossConnectGroupLifecycleStateProvisioned
	listCrossConnectGroupsResponse, err := virtualNetworkClient.ListCrossConnectGroups(context.Background(), listCrossConnectGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CrossConnectGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, crossConnectGroup := range listCrossConnectGroupsResponse.Items {
		id := *crossConnectGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CrossConnectGroupId", id)
	}
	return resourceIds, nil
}

func crossConnectGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if crossConnectGroupResponse, ok := response.Response.(oci_core.GetCrossConnectGroupResponse); ok {
		return crossConnectGroupResponse.LifecycleState != oci_core.CrossConnectGroupLifecycleStateTerminated
	}
	return false
}

func crossConnectGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetCrossConnectGroup(context.Background(), oci_core.GetCrossConnectGroupRequest{
		CrossConnectGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
