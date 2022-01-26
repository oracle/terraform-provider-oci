// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_marketplace "github.com/oracle/oci-go-sdk/v56/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	AcceptedAgreementRequiredOnlyResource = AcceptedAgreementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Required, acctest.Create, acceptedAgreementRepresentation)

	AcceptedAgreementResourceConfig = AcceptedAgreementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Optional, acctest.Update, acceptedAgreementRepresentation)

	acceptedAgreementSingularDataSourceRepresentation = map[string]interface{}{
		"accepted_agreement_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_marketplace_accepted_agreement.test_accepted_agreement.id}`},
	}

	acceptedAgreementDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"accepted_agreement_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_marketplace_accepted_agreement.test_accepted_agreement.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"listing_id":            acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_marketplace_listing.test_listing.id}`},
		"package_version":       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_marketplace_listing.test_listing.default_package_version}`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: acceptedAgreementDataSourceFilterRepresentation}}
	acceptedAgreementDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_marketplace_accepted_agreement.test_accepted_agreement.id}`}},
	}

	acceptedAgreementRepresentation = map[string]interface{}{
		"agreement_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_marketplace_listing_package_agreement.test_listing_package_agreement.agreement_id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"listing_id":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.id}`},
		"package_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.default_package_version}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"signature":       acctest.Representation{RepType: acctest.Required, Create: `${oci_marketplace_listing_package_agreement.test_listing_package_agreement.signature}`},
	}

	AcceptedAgreementResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", acctest.Required, acctest.Create, listingDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing", "test_listing", acctest.Required, acctest.Create, listingSingularDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_package_agreements", "test_listing_package_agreements", acctest.Required, acctest.Create, listingPackageAgreementDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_marketplace_listing_package_agreement", "test_listing_package_agreement", acctest.Required, acctest.Create, listingPackageAgreementManagementRepresentation)
)

// issue-routing-tag: marketplace/default
func TestMarketplaceAcceptedAgreementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceAcceptedAgreementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_marketplace_accepted_agreement.test_accepted_agreement"
	datasourceName := "data.oci_marketplace_accepted_agreements.test_accepted_agreements"
	singularDatasourceName := "data.oci_marketplace_accepted_agreement.test_accepted_agreement"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AcceptedAgreementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Optional, acctest.Create, acceptedAgreementRepresentation), "marketplace", "acceptedAgreement", t)

	acctest.ResourceTest(t, testAccCheckMarketplaceAcceptedAgreementDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AcceptedAgreementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Required, acctest.Create, acceptedAgreementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agreement_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_version"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AcceptedAgreementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AcceptedAgreementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Optional, acctest.Create, acceptedAgreementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agreement_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_version"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),

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
			Config: config + compartmentIdVariableStr + AcceptedAgreementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Optional, acctest.Update, acceptedAgreementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agreement_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_version"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_accepted_agreements", "test_accepted_agreements", acctest.Optional, acctest.Update, acceptedAgreementDataSourceRepresentation) +
				compartmentIdVariableStr + AcceptedAgreementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Optional, acctest.Update, acceptedAgreementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "accepted_agreement_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "package_version"),

				resource.TestCheckResourceAttr(datasourceName, "accepted_agreements.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "accepted_agreements.0.agreement_id"),
				resource.TestCheckResourceAttr(datasourceName, "accepted_agreements.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "accepted_agreements.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "accepted_agreements.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "accepted_agreements.0.listing_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "accepted_agreements.0.package_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "accepted_agreements.0.time_accepted"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_accepted_agreement", "test_accepted_agreement", acctest.Required, acctest.Create, acceptedAgreementSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AcceptedAgreementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "accepted_agreement_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AcceptedAgreementResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"signature",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMarketplaceAcceptedAgreementDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MarketplaceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_marketplace_accepted_agreement" {
			noResourceFound = false
			request := oci_marketplace.GetAcceptedAgreementRequest{}

			tmp := rs.Primary.ID
			request.AcceptedAgreementId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "marketplace")

			_, err := client.GetAcceptedAgreement(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("MarketplaceAcceptedAgreement") {
		resource.AddTestSweepers("MarketplaceAcceptedAgreement", &resource.Sweeper{
			Name:         "MarketplaceAcceptedAgreement",
			Dependencies: acctest.DependencyGraph["acceptedAgreement"],
			F:            sweepMarketplaceAcceptedAgreementResource,
		})
	}
}

func sweepMarketplaceAcceptedAgreementResource(compartment string) error {
	marketplaceClient := acctest.GetTestClients(&schema.ResourceData{}).MarketplaceClient()
	acceptedAgreementIds, err := getAcceptedAgreementIds(compartment)
	if err != nil {
		return err
	}
	for _, acceptedAgreementId := range acceptedAgreementIds {
		if ok := acctest.SweeperDefaultResourceId[acceptedAgreementId]; !ok {
			deleteAcceptedAgreementRequest := oci_marketplace.DeleteAcceptedAgreementRequest{}

			deleteAcceptedAgreementRequest.AcceptedAgreementId = &acceptedAgreementId

			deleteAcceptedAgreementRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "marketplace")
			_, error := marketplaceClient.DeleteAcceptedAgreement(context.Background(), deleteAcceptedAgreementRequest)
			if error != nil {
				fmt.Printf("Error deleting AcceptedAgreement %s %s, It is possible that the resource is already deleted. Please verify manually \n", acceptedAgreementId, error)
				continue
			}
		}
	}
	return nil
}

func getAcceptedAgreementIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AcceptedAgreementId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	marketplaceClient := acctest.GetTestClients(&schema.ResourceData{}).MarketplaceClient()

	listAcceptedAgreementsRequest := oci_marketplace.ListAcceptedAgreementsRequest{}
	listAcceptedAgreementsRequest.CompartmentId = &compartmentId
	listAcceptedAgreementsResponse, err := marketplaceClient.ListAcceptedAgreements(context.Background(), listAcceptedAgreementsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AcceptedAgreement list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, acceptedAgreement := range listAcceptedAgreementsResponse.Items {
		id := *acceptedAgreement.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AcceptedAgreementId", id)
	}
	return resourceIds, nil
}
