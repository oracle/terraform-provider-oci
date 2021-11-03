// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v50/common"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v50/usage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SubscriptionRedeemableUserRequiredOnlyResource = SubscriptionRedeemableUserResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", Required, Create, subscriptionRedeemableUserRepresentation)

	SubscriptionRedeemableUserResourceConfig = SubscriptionRedeemableUserResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", Optional, Update, subscriptionRedeemableUserRepresentation)

	subscriptionRedeemableUserSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": Representation{RepType: Required, Create: `${var.subscription_id}`},
		"tenancy_id":      Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	subscriptionRedeemableUserDataSourceRepresentation = map[string]interface{}{
		"subscription_id": Representation{RepType: Required, Create: `${var.subscription_id}`},
		"tenancy_id":      Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"filter":          RepresentationGroup{Required, subscriptionRedeemableUserDataSourceFilterRepresentation}}
	subscriptionRedeemableUserDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_usage_proxy_subscription_redeemable_user.test_subscription_redeemable_user.id}`}},
	}

	subscriptionRedeemableUserRepresentation = map[string]interface{}{
		"subscription_id": Representation{RepType: Required, Create: `${var.subscription_id}`},
		"tenancy_id":      Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"items":           RepresentationGroup{Required, subscriptionRedeemableUserItemsRepresentation},
		"user_id":         Representation{RepType: Optional, Create: `${var.user_id}`},
	}
	subscriptionRedeemableUserItemsRepresentation = map[string]interface{}{
		"email_id": Representation{RepType: Required, Create: `${var.email_id}`},
	}

	SubscriptionRedeemableUserResourceDependencies = ""
)

// issue-routing-tag: usage_proxy/default
func TestUsageProxySubscriptionRedeemableUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestUsageProxySubscriptionRedeemableUserResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subscriptionId := getEnvSettingWithBlankDefault("subscription_id")
	subscriptionIdVariableStr := fmt.Sprintf("variable \"subscription_id\" { default = \"%s\" }\n", subscriptionId)

	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	userId := getEnvSettingWithBlankDefault("user_id")
	userIdVariableStr := fmt.Sprintf("variable \"user_id\" { default = \"%s\" }\n", userId)

	emailId := getEnvSettingWithBlankDefault("email_id")
	emailIdVariableStr := fmt.Sprintf("variable \"email_id\" { default = \"%s\" }\n", emailId)

	resourceName := "oci_usage_proxy_subscription_redeemable_user.test_subscription_redeemable_user"
	datasourceName := "data.oci_usage_proxy_subscription_redeemable_users.test_subscription_redeemable_users"
	singularDatasourceName := "data.oci_usage_proxy_subscription_redeemable_user.test_subscription_redeemable_user"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+SubscriptionRedeemableUserResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", Required, Create, subscriptionRedeemableUserRepresentation), "usage", "subscriptionRedeemableUser", t)

	ResourceTest(t, testAccCheckUsageProxySubscriptionRedeemableUserDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + SubscriptionRedeemableUserResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", Required, Create, subscriptionRedeemableUserRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + SubscriptionRedeemableUserResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + SubscriptionRedeemableUserResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", Required, Create, subscriptionRedeemableUserRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "items.0.email_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportTenant, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportTenant {
						if errExport := TestExportCompartmentWithResourceName(&resId, &tenancyId, resourceName); errExport != nil {
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
				GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_users", "test_subscription_redeemable_users", Optional, Update, subscriptionRedeemableUserDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + SubscriptionRedeemableUserResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", Optional, Update, subscriptionRedeemableUserRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_usage_proxy_subscription_redeemable_user", "test_subscription_redeemable_user", Required, Create, subscriptionRedeemableUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + SubscriptionRedeemableUserResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + subscriptionIdVariableStr + emailIdVariableStr + userIdVariableStr + SubscriptionRedeemableUserResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
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
	client := testAccProvider.Meta().(*OracleClients).rewardsClient()
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

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "usage_proxy")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("UsageProxySubscriptionRedeemableUser") {
		resource.AddTestSweepers("UsageProxySubscriptionRedeemableUser", &resource.Sweeper{
			Name:         "UsageProxySubscriptionRedeemableUser",
			Dependencies: DependencyGraph["subscriptionRedeemableUser"],
			F:            sweepUsageProxySubscriptionRedeemableUserResource,
		})
	}
}

func sweepUsageProxySubscriptionRedeemableUserResource(compartment string) error {
	rewardsClient := GetTestClients(&schema.ResourceData{}).rewardsClient()
	subscriptionRedeemableUserIds, err := getSubscriptionRedeemableUserIds(compartment)
	if err != nil {
		return err
	}
	for _, subscriptionRedeemableUserId := range subscriptionRedeemableUserIds {
		if ok := SweeperDefaultResourceId[subscriptionRedeemableUserId]; !ok {
			deleteRedeemableUserRequest := oci_usage_proxy.DeleteRedeemableUserRequest{}

			deleteRedeemableUserRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "usage_proxy")
			_, error := rewardsClient.DeleteRedeemableUser(context.Background(), deleteRedeemableUserRequest)
			if error != nil {
				fmt.Printf("Error deleting SubscriptionRedeemableUser %s %s, It is possible that the resource is already deleted. Please verify manually \n", subscriptionRedeemableUserId, error)
				continue
			}
		}
	}
	return nil
}

func getSubscriptionRedeemableUserIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "SubscriptionRedeemableUserId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	rewardsClient := GetTestClients(&schema.ResourceData{}).rewardsClient()

	listRedeemableUsersRequest := oci_usage_proxy.ListRedeemableUsersRequest{}
	listRedeemableUsersRequest.TenancyId = &compartmentId

	subscriptionIds, error := getSubscriptionIds(compartment)
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
			AddResourceIdToSweeperResourceIdMap(compartmentId, "SubscriptionRedeemableUserId", id)
		}
	}
	return resourceIds, nil
}
