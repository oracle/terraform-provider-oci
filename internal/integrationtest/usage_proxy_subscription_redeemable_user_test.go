// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v65/usage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	UsageProxySubscriptionRedeemableUserRequiredOnlyResource = UsageProxySubscriptionRedeemableUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", acctest.Required, acctest.Create, UsageProxySubscriptionRedeemableUserRepresentation)

	UsageProxySubscriptionRedeemableUserResourceConfig = UsageProxySubscriptionRedeemableUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", acctest.Optional, acctest.Update, UsageProxySubscriptionRedeemableUserRepresentation)

	UsageProxyUsageProxySubscriptionRedeemableUserSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	UsageProxyUsageProxySubscriptionRedeemableUserDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: UsageProxySubscriptionRedeemableUserDataSourceFilterRepresentation}}
	UsageProxySubscriptionRedeemableUserDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_usage_proxy_subscription_redeemable_user.test_subscription_redeemable_user.id}`}},
	}

	UsageProxySubscriptionRedeemableUserRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subscription_id}`},
		"tenancy_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: UsageProxySubscriptionRedeemableUserItemsRepresentation},
		"user_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.user_id}`},
	}
	UsageProxySubscriptionRedeemableUserItemsRepresentation = map[string]interface{}{
		"email_id": acctest.Representation{RepType: acctest.Required, Create: `${var.email_id}`},
	}

	UsageProxySubscriptionRedeemableUserResourceDependencies = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxySubscriptionRedeemableUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxySubscriptionRedeemableUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := utils.GetEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	userId := utils.GetEnvSettingWithBlankDefault("user_id")
	userIdVariableStr := fmt.Sprintf("variable \"user_id\" { default = \"%s\" }\n", userId)

	emailId := utils.GetEnvSettingWithBlankDefault("email_id")
	emailIdVariableStr := fmt.Sprintf("variable \"email_id\" { default = \"%s\" }\n", emailId)

	resourceName := "oci_usage_proxy_subscription_redeemable_user.test_subscription_redeemable_user"
	datasourceName := "data.oci_usage_proxy_subscription_redeemable_users.test_subscription_redeemable_users"
	singularDatasourceName := "data.oci_usage_proxy_subscription_redeemable_user.test_subscription_redeemable_user"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+UsageProxySubscriptionRedeemableUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", acctest.Required, acctest.Create, UsageProxySubscriptionRedeemableUserRepresentation), "usage", "subscriptionRedeemableUser", t)

	acctest.ResourceTest(t, testAccCheckUsageProxySubscriptionRedeemableUserDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + UsageProxySubscriptionRedeemableUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", acctest.Required, acctest.Create, UsageProxySubscriptionRedeemableUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + UsageProxySubscriptionRedeemableUserResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + UsageProxySubscriptionRedeemableUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", acctest.Required, acctest.Create, UsageProxySubscriptionRedeemableUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "items.0.email_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportTenant, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportTenant {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_users", "test_subscription_redeemable_users", acctest.Optional, acctest.Update, UsageProxyUsageProxySubscriptionRedeemableUserDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + UsageProxySubscriptionRedeemableUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", acctest.Optional, acctest.Update, UsageProxySubscriptionRedeemableUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),

				resource.TestCheckResourceAttr(datasourceName, "redeemable_user_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "redeemable_user_collection.0.items.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "redeemable_user_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", acctest.Required, acctest.Create, UsageProxyUsageProxySubscriptionRedeemableUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + UsageProxySubscriptionRedeemableUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + UsageProxySubscriptionRedeemableUserRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"tenancy_id",
				"user_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckUsageProxySubscriptionRedeemableUserDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).RewardsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_usage_proxy_subscription_redeemable_user" {
			noResourceFound = false
			request := oci_usage_proxy.ListRedeemableUsersRequest{}

			if value, ok := rs.Primary.Attributes["subscription_id"]; ok {
				request.SubscriptionId = &value
			}

			if value, ok := rs.Primary.Attributes["tenancy_id"]; ok {
				request.TenancyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "usage_proxy")

			_, err := client.ListRedeemableUsers(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("UsageProxySubscriptionRedeemableUser") {
		resource.AddTestSweepers("UsageProxySubscriptionRedeemableUser", &resource.Sweeper{
			Name:         "UsageProxySubscriptionRedeemableUser",
			Dependencies: acctest.DependencyGraph["subscriptionRedeemableUser"],
			F:            sweepUsageProxySubscriptionRedeemableUserResource,
		})
	}
}

func sweepUsageProxySubscriptionRedeemableUserResource(compartment string) error {
	rewardsClient := acctest.GetTestClients(&schema.ResourceData{}).RewardsClient()
	subscriptionRedeemableUserIds, err := getUsageProxySubscriptionRedeemableUserIds(compartment)
	if err != nil {
		return err
	}
	for _, subscriptionRedeemableUserId := range subscriptionRedeemableUserIds {
		if ok := acctest.SweeperDefaultResourceId[subscriptionRedeemableUserId]; !ok {
			deleteRedeemableUserRequest := oci_usage_proxy.DeleteRedeemableUserRequest{}

			deleteRedeemableUserRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "usage_proxy")
			_, error := rewardsClient.DeleteRedeemableUser(context.Background(), deleteRedeemableUserRequest)
			if error != nil {
				fmt.Printf("Error deleting SubscriptionRedeemableUser %s %s, It is possible that the resource is already deleted. Please verify manually \n", subscriptionRedeemableUserId, error)
				continue
			}
		}
	}
	return nil
}

func getUsageProxySubscriptionRedeemableUserIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubscriptionRedeemableUserId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	rewardsClient := acctest.GetTestClients(&schema.ResourceData{}).RewardsClient()

	listRedeemableUsersRequest := oci_usage_proxy.ListRedeemableUsersRequest{}
	listRedeemableUsersRequest.TenancyId = &compartmentId

	subscriptionIds, error := getOnsSubscriptionIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting subscriptionId required for SubscriptionRedeemableUser resource requests \n")
	}
	for _, subscriptionId := range subscriptionIds {
		listRedeemableUsersRequest.SubscriptionId = &subscriptionId
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting tenancyId required for SubscriptionRedeemableUser resource requests \n")
		}

		listRedeemableUsersResponse, err := rewardsClient.ListRedeemableUsers(context.Background(), listRedeemableUsersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SubscriptionRedeemableUser list for compartment id : %s , %s \n", compartmentId, err)
		}

		userCollection := listRedeemableUsersResponse.RedeemableUserCollection

		for _, redeemableUserSummary := range userCollection.Items {
			id := *redeemableUserSummary.EmailId
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubscriptionRedeemableUserId", id)
		}
	}
	return resourceIds, nil
}
