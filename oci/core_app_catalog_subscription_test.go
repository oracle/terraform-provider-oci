// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AppCatalogSubscriptionRequiredOnlyResource = AppCatalogSubscriptionResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_app_catalog_subscription", "test_app_catalog_subscription", Required, Create, appCatalogSubscriptionRepresentation)

	appCatalogSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"listing_id":     Representation{repType: Optional, create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_id}`},
		"filter":         RepresentationGroup{Required, appCatalogSubscriptionDataSourceFilterRepresentation}}
	appCatalogSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `listing_resource_version`},
		"values": Representation{repType: Required, create: []string{`${oci_core_app_catalog_subscription.test_app_catalog_subscription.listing_resource_version}`}},
	}

	appCatalogSubscriptionRepresentation = map[string]interface{}{
		"compartment_id":           Representation{repType: Optional, create: `${var.compartment_id}`},
		"eula_link":                Representation{repType: Optional, create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.eula_link}`},
		"listing_id":               Representation{repType: Optional, create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_id}`},
		"listing_resource_version": Representation{repType: Optional, create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_resource_version}`},
		"oracle_terms_of_use_link": Representation{repType: Optional, create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.oracle_terms_of_use_link}`},
		"signature":                Representation{repType: Optional, create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.signature}`},
		"time_retrieved":           Representation{repType: Optional, create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.time_retrieved}`},
	}

	AppCatalogSubscriptionResourceDependencies = AppCatalogListingResourceVersionAgreementResourceConfig
)

func TestCoreAppCatalogSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreAppCatalogSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_app_catalog_subscription.test_app_catalog_subscription"
	datasourceName := "data.oci_core_app_catalog_subscriptions.test_app_catalog_subscriptions"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreAppCatalogSubscriptionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AppCatalogSubscriptionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_app_catalog_subscription", "test_app_catalog_subscription", Optional, Create, appCatalogSubscriptionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_app_catalog_subscriptions", "test_app_catalog_subscriptions", Optional, Create, appCatalogSubscriptionDataSourceRepresentation) +
					compartmentIdVariableStr + AppCatalogSubscriptionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_app_catalog_subscription", "test_app_catalog_subscription", Optional, Create, appCatalogSubscriptionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),

					resource.TestCheckResourceAttr(datasourceName, "app_catalog_subscriptions.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "app_catalog_subscriptions.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_subscriptions.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_subscriptions.0.listing_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_subscriptions.0.listing_resource_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_subscriptions.0.summary"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_subscriptions.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"eula_link",
					"oracle_terms_of_use_link",
					"signature",
					"time_retrieved",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCoreAppCatalogSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_app_catalog_subscription" {
			noResourceFound = false
			request := oci_core.ListAppCatalogSubscriptionsRequest{}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			if value, ok := rs.Primary.Attributes["listing_id"]; ok {
				request.ListingId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			listingResourceVersion, _ := rs.Primary.Attributes["listing_resource_version"]

			response, err := client.ListAppCatalogSubscriptions(context.Background(), request)

			if err != nil {
				for _, item := range response.Items {
					if *item.ListingResourceVersion == listingResourceVersion {
						return fmt.Errorf("deletion failed")
					}
				}

				for response.OpcNextPage != nil {
					request.Page = response.OpcNextPage
					response, err := client.ListAppCatalogSubscriptions(context.Background(), request)
					if err != nil {
						return err
					}
					for _, item := range response.Items {
						if *item.ListingResourceVersion == listingResourceVersion {
							return fmt.Errorf("deletion failed")
						}
					}
				}
				return nil
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	resource.AddTestSweepers("CoreAppCatalogSubscription", &resource.Sweeper{
		Name:         "CoreAppCatalogSubscription",
		Dependencies: DependencyGraph["appCatalogSubscription"],
		F:            sweepCoreAppCatalogSubscriptionResource,
	})
}

func sweepCoreAppCatalogSubscriptionResource(compartment string) error {
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient
	appCatalogSubscriptionIds, err := getAppCatalogSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, appCatalogSubscriptionId := range appCatalogSubscriptionIds {
		if ok := SweeperDefaultResourceId[appCatalogSubscriptionId]; !ok {
			deleteAppCatalogSubscriptionRequest := oci_core.DeleteAppCatalogSubscriptionRequest{}

			deleteAppCatalogSubscriptionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := computeClient.DeleteAppCatalogSubscription(context.Background(), deleteAppCatalogSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting AppCatalogSubscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", appCatalogSubscriptionId, error)
				continue
			}
		}
	}
	return nil
}

func getAppCatalogSubscriptionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "AppCatalogSubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := GetTestClients(&schema.ResourceData{}).computeClient

	listAppCatalogSubscriptionsRequest := oci_core.ListAppCatalogSubscriptionsRequest{}
	listAppCatalogSubscriptionsRequest.CompartmentId = &compartmentId
	listAppCatalogSubscriptionsResponse, err := computeClient.ListAppCatalogSubscriptions(context.Background(), listAppCatalogSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AppCatalogSubscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, appCatalogSubscription := range listAppCatalogSubscriptionsResponse.Items {
		id := *appCatalogSubscription.ListingId
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "AppCatalogSubscriptionId", id)
	}
	return resourceIds, nil
}
