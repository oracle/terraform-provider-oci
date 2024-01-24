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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreAppCatalogSubscriptionRequiredOnlyResource = CoreAppCatalogSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_app_catalog_subscription", "test_app_catalog_subscription", acctest.Required, acctest.Create, CoreAppCatalogSubscriptionRepresentation)

	CoreCoreAppCatalogSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"listing_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreAppCatalogSubscriptionDataSourceFilterRepresentation}}
	CoreAppCatalogSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `listing_resource_version`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_app_catalog_subscription.test_app_catalog_subscription.listing_resource_version}`}},
	}

	CoreAppCatalogSubscriptionRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"listing_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_id}`},
		"listing_resource_version": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.listing_resource_version}`},
		"oracle_terms_of_use_link": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.oracle_terms_of_use_link}`},
		"signature":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.signature}`},
		"time_retrieved":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.time_retrieved}`},
		"eula_link":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement.eula_link}`},
	}

	CoreAppCatalogSubscriptionResourceDependencies = AppCatalogListingResourceVersionAgreementResourceConfig
)

// issue-routing-tag: core/computeImaging
func TestCoreAppCatalogSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreAppCatalogSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_app_catalog_subscription.test_app_catalog_subscription"
	datasourceName := "data.oci_core_app_catalog_subscriptions.test_app_catalog_subscriptions"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreAppCatalogSubscriptionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_app_catalog_subscription", "test_app_catalog_subscription", acctest.Optional, acctest.Create, CoreAppCatalogSubscriptionRepresentation), "core", "appCatalogSubscription", t)

	acctest.ResourceTest(t, testAccCheckCoreAppCatalogSubscriptionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreAppCatalogSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_app_catalog_subscription", "test_app_catalog_subscription", acctest.Required, acctest.Create, CoreAppCatalogSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "eula_link"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_resource_version"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_terms_of_use_link"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestCheckResourceAttrSet(resourceName, "time_retrieved"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_app_catalog_subscriptions", "test_app_catalog_subscriptions", acctest.Optional, acctest.Create, CoreCoreAppCatalogSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + CoreAppCatalogSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_app_catalog_subscription", "test_app_catalog_subscription", acctest.Optional, acctest.Create, CoreAppCatalogSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config:            config + CoreAppCatalogSubscriptionRequiredOnlyResource,
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
	})
}

func testAccCheckCoreAppCatalogSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeClient()
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

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreAppCatalogSubscription") {
		resource.AddTestSweepers("CoreAppCatalogSubscription", &resource.Sweeper{
			Name:         "CoreAppCatalogSubscription",
			Dependencies: acctest.DependencyGraph["appCatalogSubscription"],
			F:            sweepCoreAppCatalogSubscriptionResource,
		})
	}
}

func sweepCoreAppCatalogSubscriptionResource(compartment string) error {
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()
	appCatalogSubscriptionIds, err := getCoreAppCatalogSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, appCatalogSubscriptionId := range appCatalogSubscriptionIds {
		if ok := acctest.SweeperDefaultResourceId[appCatalogSubscriptionId]; !ok {
			deleteAppCatalogSubscriptionRequest := oci_core.DeleteAppCatalogSubscriptionRequest{}

			deleteAppCatalogSubscriptionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := computeClient.DeleteAppCatalogSubscription(context.Background(), deleteAppCatalogSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting AppCatalogSubscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", appCatalogSubscriptionId, error)
				continue
			}
		}
	}
	return nil
}

func getCoreAppCatalogSubscriptionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AppCatalogSubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeClient()

	listAppCatalogSubscriptionsRequest := oci_core.ListAppCatalogSubscriptionsRequest{}
	listAppCatalogSubscriptionsRequest.CompartmentId = &compartmentId
	listAppCatalogSubscriptionsResponse, err := computeClient.ListAppCatalogSubscriptions(context.Background(), listAppCatalogSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AppCatalogSubscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, appCatalogSubscription := range listAppCatalogSubscriptionsResponse.Items {
		id := *appCatalogSubscription.ListingId
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AppCatalogSubscriptionId", id)
	}
	return resourceIds, nil
}
