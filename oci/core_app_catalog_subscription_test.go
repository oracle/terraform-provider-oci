// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
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
					resource.TestCheckResourceAttrSet(resourceName, "listing_resource_version"),
					resource.TestCheckResourceAttrSet(resourceName, "oracle_terms_of_use_link"),
					resource.TestCheckResourceAttrSet(resourceName, "signature"),
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
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_subscriptions.0.listing_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_subscriptions.0.listing_resource_version"),
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
