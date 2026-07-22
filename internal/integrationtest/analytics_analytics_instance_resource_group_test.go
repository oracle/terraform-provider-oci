// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AnalyticsAnalyticsInstanceResourceGroupRequiredOnlyResource = AnalyticsAnalyticsInstanceResourceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_resource_group", "test_analytics_instance_resource_group", acctest.Required, acctest.Create, AnalyticsAnalyticsInstanceResourceGroupRepresentation)

	AnalyticsAnalyticsInstanceResourceGroupResourceConfig = AnalyticsAnalyticsInstanceResourceGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_resource_group", "test_analytics_instance_resource_group", acctest.Optional, acctest.Update, AnalyticsAnalyticsInstanceResourceGroupRepresentation)

	AnalyticsAnalyticsInstanceResourceGroupSingularDataSourceRepresentation = map[string]interface{}{
		"analytics_instance_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"analytics_instance_resource_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance_resource_group.test_analytics_instance_resource_group.id}`},
	}

	AnalyticsAnalyticsInstanceResourceGroupDataSourceRepresentation = map[string]interface{}{
		"analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: AnalyticsAnalyticsInstanceResourceGroupDataSourceFilterRepresentation}}
	AnalyticsAnalyticsInstanceResourceGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_analytics_analytics_instance_resource_group.test_analytics_instance_resource_group.id}`}},
	}

	AnalyticsAnalyticsInstanceResourceGroupRepresentation = map[string]interface{}{
		"analytics_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
		"capacity":              acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"resource_name":         acctest.Representation{RepType: acctest.Required, Create: `test-resource-group`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	analyticsInstanceResourceGroupCapacityRepresentation = map[string]interface{}{
		"capacity_type":  acctest.Representation{RepType: acctest.Required, Create: `OLPU_COUNT`},
		"capacity_value": acctest.Representation{RepType: acctest.Required, Create: `2`},
	}
	analyticsInstanceResourceGroupIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`capacity`}},
	}

	AnalyticsAnalyticsInstanceResourceGroupResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_analytics_analytics_instance",
		"test_analytics_instance",
		acctest.Optional,
		acctest.Create,
		acctest.RepresentationCopyWithNewProperties(analyticsPublicInstanceRepresentation, map[string]interface{}{
			"capacity":  acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceResourceGroupCapacityRepresentation},
			"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: analyticsInstanceResourceGroupIgnoreChangesRepresentation},
		}),
	)
)

// issue-routing-tag: analytics/default
func TestAnalyticsAnalyticsInstanceResourceGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnalyticsAnalyticsInstanceResourceGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)
	oacMtShardId := utils.GetEnvSettingWithBlankDefault("oac_mt_shard_id")
	oacMtShardIdVariableStr := fmt.Sprintf("variable \"oac_mt_shard_id\" { default = \"%s\" }\n", oacMtShardId)
	identityMtShardId := utils.GetEnvSettingWithBlankDefault("identity_mt_shard_id")
	identityMtShardIdVariableStr := fmt.Sprintf("variable \"identity_mt_shard_id\" { default = \"%s\" }\n", identityMtShardId)

	resourceName := "oci_analytics_analytics_instance_resource_group.test_analytics_instance_resource_group"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+idcsAccessTokenVariableStr+oacMtShardIdVariableStr+identityMtShardIdVariableStr+AnalyticsAnalyticsInstanceResourceGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_resource_group", "test_analytics_instance_resource_group", acctest.Optional, acctest.Create, AnalyticsAnalyticsInstanceResourceGroupRepresentation), "analytics", "analyticsInstanceResourceGroup", t)

	acctest.ResourceTest(t, testAccCheckAnalyticsAnalyticsInstanceResourceGroupDestroy, []resource.TestStep{
		// verify Create with optionals for public instance
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + oacMtShardIdVariableStr + identityMtShardIdVariableStr + AnalyticsAnalyticsInstanceResourceGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_analytics_analytics_instance_resource_group", "test_analytics_instance_resource_group", acctest.Optional, acctest.Create, AnalyticsAnalyticsInstanceResourceGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "analytics_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "capacity", "10"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "resource_name", "test-resource-group"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
	})
}

func testAccCheckAnalyticsAnalyticsInstanceResourceGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_analytics_analytics_instance_resource_group" {
			noResourceFound = false
			request := oci_analytics.GetResourceGroupRequest{}

			if value, ok := rs.Primary.Attributes["analytics_instance_id"]; ok {
				request.AnalyticsInstanceId = &value
			}

			_, resourceGroupId, err := parseAnalyticsInstanceResourceGroupCompositeIdForTest(rs.Primary.ID)
			if err != nil {
				return err
			}

			tmp := resourceGroupId
			request.AnalyticsInstanceResourceGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "analytics")

			_, err = client.GetResourceGroup(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("AnalyticsAnalyticsInstanceResourceGroup") {
		resource.AddTestSweepers("AnalyticsAnalyticsInstanceResourceGroup", &resource.Sweeper{
			Name:         "AnalyticsAnalyticsInstanceResourceGroup",
			Dependencies: acctest.DependencyGraph["analyticsInstanceResourceGroup"],
			F:            sweepAnalyticsAnalyticsInstanceResourceGroupResource,
		})
	}
}

func sweepAnalyticsAnalyticsInstanceResourceGroupResource(compartment string) error {
	analyticsClient := acctest.GetTestClients(&schema.ResourceData{}).AnalyticsClient()
	analyticsInstanceResourceGroupIds, err := getAnalyticsAnalyticsInstanceResourceGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, analyticsInstanceResourceGroupId := range analyticsInstanceResourceGroupIds {
		if ok := acctest.SweeperDefaultResourceId[analyticsInstanceResourceGroupId]; !ok {
			analyticsInstanceId, resourceGroupId, err := parseAnalyticsInstanceResourceGroupCompositeIdForTest(analyticsInstanceResourceGroupId)
			if err != nil {
				fmt.Printf("Error parsing AnalyticsInstanceResourceGroup composite id %s, %s \n", analyticsInstanceResourceGroupId, err)
				continue
			}

			deleteResourceGroupRequest := oci_analytics.DeleteResourceGroupRequest{}
			deleteResourceGroupRequest.AnalyticsInstanceId = &analyticsInstanceId
			deleteResourceGroupRequest.AnalyticsInstanceResourceGroupId = &resourceGroupId

			deleteResourceGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "analytics")
			_, error := analyticsClient.DeleteResourceGroup(context.Background(), deleteResourceGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting AnalyticsInstanceResourceGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", analyticsInstanceResourceGroupId, error)
				continue
			}
		}
	}
	return nil
}

func getAnalyticsAnalyticsInstanceResourceGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AnalyticsInstanceResourceGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	analyticsClient := acctest.GetTestClients(&schema.ResourceData{}).AnalyticsClient()

	listResourceGroupsRequest := oci_analytics.ListResourceGroupsRequest{}

	analyticsInstanceIds, error := getAnalyticsAnalyticsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting analyticsInstanceId required for AnalyticsInstanceResourceGroup resource requests \n")
	}
	for _, analyticsInstanceId := range analyticsInstanceIds {
		listResourceGroupsRequest.AnalyticsInstanceId = &analyticsInstanceId

		listResourceGroupsResponse, err := analyticsClient.ListResourceGroups(context.Background(), listResourceGroupsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting AnalyticsInstanceResourceGroup list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, analyticsInstanceResourceGroup := range listResourceGroupsResponse.Items {
			id := fmt.Sprintf("analyticsInstances/%s/resourceGroups/%s", analyticsInstanceId, *analyticsInstanceResourceGroup.Id)
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AnalyticsInstanceResourceGroupId", id)
		}

	}
	return resourceIds, nil
}

func parseAnalyticsInstanceResourceGroupCompositeIdForTest(compositeId string) (string, string, error) {
	parts := strings.Split(compositeId, "/")
	if len(parts) != 4 || parts[0] != "analyticsInstances" || parts[2] != "resourceGroups" {
		return "", "", fmt.Errorf("illegal compositeId %s encountered", compositeId)
	}

	analyticsInstanceId, err := url.PathUnescape(parts[1])
	if err != nil {
		return "", "", err
	}

	resourceGroupId, err := url.PathUnescape(parts[3])
	if err != nil {
		return "", "", err
	}

	return analyticsInstanceId, resourceGroupId, nil
}
