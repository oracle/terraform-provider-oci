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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_marketplace "github.com/oracle/oci-go-sdk/v58/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	PublicationResourceConfig = PublicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Optional, acctest.Update, publicationRepresentation)

	publicationSingularDataSourceRepresentation = map[string]interface{}{
		"publication_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_marketplace_publication.test_publication.id}`},
	}

	publicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"listing_type":      acctest.Representation{RepType: acctest.Required, Create: `COMMUNITY`},
		"name":              acctest.Representation{RepType: acctest.Optional, Create: []string{`name`}, Update: []string{`name2`}},
		"operating_systems": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_image.test_image.operating_system}`}},
		"publication_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_marketplace_publication.test_publication.id}`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: publicationDataSourceFilterRepresentation}}
	publicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_marketplace_publication.test_publication.id}`}},
	}

	publicationRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_agreement_acknowledged": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"listing_type":              acctest.Representation{RepType: acctest.Required, Create: `COMMUNITY`},
		"name":                      acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"package_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: publicationPackageDetailsRepresentation},
		"short_description":         acctest.Representation{RepType: acctest.Required, Create: `shortDescription`, Update: `shortDescription2`},
		"support_contacts":          acctest.RepresentationGroup{RepType: acctest.Required, Group: publicationSupportContactsRepresentation},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"long_description":          acctest.Representation{RepType: acctest.Optional, Create: `longDescription`, Update: `longDescription2`},
	}
	publicationPackageDetailsRepresentation = map[string]interface{}{
		"eula":             acctest.RepresentationGroup{RepType: acctest.Required, Group: publicationPackageDetailsEulaRepresentation},
		"operating_system": acctest.RepresentationGroup{RepType: acctest.Required, Group: publicationPackageDetailsOperatingSystemRepresentation},
		"package_type":     acctest.Representation{RepType: acctest.Required, Create: `IMAGE`},
		"package_version":  acctest.Representation{RepType: acctest.Required, Create: `packageVersion`},
		"image_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_image.test_image.id}`},
	}
	publicationSupportContactsRepresentation = map[string]interface{}{
		"email":   acctest.Representation{RepType: acctest.Required, Create: `email`, Update: `email2`},
		"name":    acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"phone":   acctest.Representation{RepType: acctest.Optional, Create: `phone`, Update: `phone2`},
		"subject": acctest.Representation{RepType: acctest.Optional, Create: `subject`, Update: `subject2`},
	}
	publicationPackageDetailsEulaRepresentation = map[string]interface{}{
		"eula_type":    acctest.Representation{RepType: acctest.Required, Create: `TEXT`},
		"license_text": acctest.Representation{RepType: acctest.Required, Create: `licenseText`},
	}
	publicationPackageDetailsOperatingSystemRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_image.test_image.operating_system}`},
	}

	PublicationResourceDependencies = ImageRequiredOnlyResource
)

// issue-routing-tag: marketplace/default
func TestMarketplacePublicationResource_basic(t *testing.T) {
	t.Skip("Skip this test till Marketplace automates background processes and reduces the turnaround time.")

	httpreplay.SetScenario("TestMarketplacePublicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_marketplace_publication.test_publication"
	datasourceName := "data.oci_marketplace_publications.test_publications"
	singularDatasourceName := "data.oci_marketplace_publication.test_publication"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckMarketplacePublicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PublicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Required, acctest.Create, publicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
				resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
					"eula_type": "TEXT",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_version", "packageVersion"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
				resource.TestCheckResourceAttr(resourceName, "support_contacts.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{},
					[]string{}),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PublicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PublicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Optional, acctest.Create, publicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
				resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
				resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
					"eula_type":    "TEXT",
					"license_text": "licenseText",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "package_details.0.image_id"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.0.name", "Oracle Linux"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_version", "packageVersion"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "support_contacts.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{
					"email":   "email",
					"name":    "name",
					"phone":   "phone",
					"subject": "subject",
				},
					[]string{}),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PublicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(publicationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
				resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
				resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
					"eula_type":    "TEXT",
					"license_text": "licenseText",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "package_details.0.image_id"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.0.name", "Oracle Linux"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_version", "packageVersion"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "support_contacts.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{
					"email":   "email",
					"name":    "name",
					"phone":   "phone",
					"subject": "subject",
				},
					[]string{}),

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
			Config: config + compartmentIdVariableStr + PublicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Optional, acctest.Update, publicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
				resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
				resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription2"),
				resource.TestCheckResourceAttr(resourceName, "name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
					"eula_type":    "TEXT",
					"license_text": "licenseText",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "package_details.0.image_id"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.0.name", "Oracle Linux"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_version", "packageVersion"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "support_contacts.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{
					"email":   "email2",
					"name":    "name2",
					"phone":   "phone2",
					"subject": "subject2",
				},
					[]string{}),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_publications", "test_publications", acctest.Optional, acctest.Update, publicationDataSourceRepresentation) +
				compartmentIdVariableStr + PublicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Optional, acctest.Update, publicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "listing_type", "COMMUNITY"),
				resource.TestCheckResourceAttr(datasourceName, "name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "operating_systems.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "publication_id"),

				resource.TestCheckResourceAttr(datasourceName, "publications.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "publications.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "publications.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "publications.0.listing_type", "COMMUNITY"),
				resource.TestCheckResourceAttr(datasourceName, "publications.0.name", "name2"),
				resource.TestCheckResourceAttrSet(datasourceName, "publications.0.package_type"),
				resource.TestCheckResourceAttr(datasourceName, "publications.0.short_description", "shortDescription2"),
				resource.TestCheckResourceAttrSet(datasourceName, "publications.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "publications.0.supported_operating_systems.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "publications.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Required, acctest.Create, publicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PublicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "publication_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "listing_type", "COMMUNITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_description", "longDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_description", "shortDescription2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "support_contacts.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "support_contacts", map[string]string{
					"email":   "email2",
					"name":    "name2",
					"phone":   "phone2",
					"subject": "subject2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "supported_operating_systems.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PublicationResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_agreement_acknowledged",
				"package_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMarketplacePublicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MarketplaceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_marketplace_publication" {
			noResourceFound = false
			request := oci_marketplace.GetPublicationRequest{}

			tmp := rs.Primary.ID
			request.PublicationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "marketplace")

			response, err := client.GetPublication(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_marketplace.PublicationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MarketplacePublication") {
		resource.AddTestSweepers("MarketplacePublication", &resource.Sweeper{
			Name:         "MarketplacePublication",
			Dependencies: acctest.DependencyGraph["publication"],
			F:            sweepMarketplacePublicationResource,
		})
	}
}

func sweepMarketplacePublicationResource(compartment string) error {
	marketplaceClient := acctest.GetTestClients(&schema.ResourceData{}).MarketplaceClient()
	publicationIds, err := getPublicationIds(compartment)
	if err != nil {
		return err
	}
	for _, publicationId := range publicationIds {
		if ok := acctest.SweeperDefaultResourceId[publicationId]; !ok {
			deletePublicationRequest := oci_marketplace.DeletePublicationRequest{}

			deletePublicationRequest.PublicationId = &publicationId

			deletePublicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "marketplace")
			_, error := marketplaceClient.DeletePublication(context.Background(), deletePublicationRequest)
			if error != nil {
				fmt.Printf("Error deleting Publication %s %s, It is possible that the resource is already deleted. Please verify manually \n", publicationId, error)
				continue
			}
			//acctest.WaitTillCondition(acctest.TestAccProvider, &publicationId, publicationSweepWaitCondition, time.Duration(3*time.Minute),
			//	publicationSweepResponseFetchOperation, "marketplace", true)
		}
	}
	return nil
}

func getPublicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PublicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	marketplaceClient := acctest.GetTestClients(&schema.ResourceData{}).MarketplaceClient()

	listPublicationsRequest := oci_marketplace.ListPublicationsRequest{}
	listPublicationsRequest.CompartmentId = &compartmentId

	listingTypes := oci_marketplace.GetListPublicationsListingTypeEnumValues()
	for _, listingType := range listingTypes {
		listPublicationsRequest.ListingType = listingType

		listPublicationsResponse, err := marketplaceClient.ListPublications(context.Background(), listPublicationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Publication list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, publication := range listPublicationsResponse.Items {
			id := *publication.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PublicationId", id)
		}

	}
	return resourceIds, nil
}

func publicationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if publicationResponse, ok := response.Response.(oci_marketplace.GetPublicationResponse); ok {
		return publicationResponse.LifecycleState != oci_marketplace.PublicationLifecycleStateDeleted
	}
	return false
}

func publicationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MarketplaceClient().GetPublication(context.Background(), oci_marketplace.GetPublicationRequest{
		PublicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
