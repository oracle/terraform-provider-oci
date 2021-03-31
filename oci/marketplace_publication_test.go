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
	"github.com/oracle/oci-go-sdk/v38/common"
	oci_marketplace "github.com/oracle/oci-go-sdk/v38/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PublicationResourceConfig = PublicationResourceDependencies +
		generateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", Optional, Update, publicationRepresentation)

	publicationSingularDataSourceRepresentation = map[string]interface{}{
		"publication_id": Representation{repType: Required, create: `${oci_marketplace_publication.test_publication.id}`},
	}

	publicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    Representation{repType: Required, create: `${var.compartment_id}`},
		"listing_type":      Representation{repType: Required, create: `COMMUNITY`},
		"name":              Representation{repType: Optional, create: []string{`name`}, update: []string{`name2`}},
		"operating_systems": Representation{repType: Optional, create: []string{`${oci_core_image.test_image.operating_system}`}},
		"publication_id":    Representation{repType: Optional, create: `${oci_marketplace_publication.test_publication.id}`},
		"filter":            RepresentationGroup{Required, publicationDataSourceFilterRepresentation}}
	publicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_marketplace_publication.test_publication.id}`}},
	}

	publicationRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"is_agreement_acknowledged": Representation{repType: Required, create: `true`},
		"listing_type":              Representation{repType: Required, create: `COMMUNITY`},
		"name":                      Representation{repType: Required, create: `name`, update: `name2`},
		"package_details":           RepresentationGroup{Required, publicationPackageDetailsRepresentation},
		"short_description":         Representation{repType: Required, create: `shortDescription`, update: `shortDescription2`},
		"support_contacts":          RepresentationGroup{Required, publicationSupportContactsRepresentation},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"long_description":          Representation{repType: Optional, create: `longDescription`, update: `longDescription2`},
	}
	publicationPackageDetailsRepresentation = map[string]interface{}{
		"eula":             RepresentationGroup{Required, publicationPackageDetailsEulaRepresentation},
		"operating_system": RepresentationGroup{Required, publicationPackageDetailsOperatingSystemRepresentation},
		"package_type":     Representation{repType: Required, create: `IMAGE`},
		"package_version":  Representation{repType: Required, create: `packageVersion`},
		"image_id":         Representation{repType: Required, create: `${oci_core_image.test_image.id}`},
	}
	publicationSupportContactsRepresentation = map[string]interface{}{
		"email":   Representation{repType: Required, create: `email`, update: `email2`},
		"name":    Representation{repType: Required, create: `name`, update: `name2`},
		"phone":   Representation{repType: Optional, create: `phone`, update: `phone2`},
		"subject": Representation{repType: Optional, create: `subject`, update: `subject2`},
	}
	publicationPackageDetailsEulaRepresentation = map[string]interface{}{
		"eula_type":    Representation{repType: Required, create: `TEXT`},
		"license_text": Representation{repType: Required, create: `licenseText`},
	}
	publicationPackageDetailsOperatingSystemRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `${oci_core_image.test_image.operating_system}`},
	}

	PublicationResourceDependencies = ImageRequiredOnlyResource
)

func TestMarketplacePublicationResource_basic(t *testing.T) {
	t.Skip("Skip this test till Marketplace automates background processes and reduces the turnaround time.")

	httpreplay.SetScenario("TestMarketplacePublicationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_marketplace_publication.test_publication"
	datasourceName := "data.oci_marketplace_publications.test_publications"
	singularDatasourceName := "data.oci_marketplace_publication.test_publication"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckMarketplacePublicationDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PublicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", Required, Create, publicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
					resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
						"eula_type": "TEXT",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "package_details.0.operating_system.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "IMAGE"),
					resource.TestCheckResourceAttr(resourceName, "package_details.0.package_version", "packageVersion"),
					resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
					resource.TestCheckResourceAttr(resourceName, "support_contacts.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PublicationResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PublicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", Optional, Create, publicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
					resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
					resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
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
					CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{
						"email":   "email",
						"name":    "name",
						"phone":   "phone",
						"subject": "subject",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PublicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", Optional, Create,
						representationCopyWithNewProperties(publicationRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
					resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
					resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription"),
					resource.TestCheckResourceAttr(resourceName, "name", "name"),
					resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
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
					CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{
						"email":   "email",
						"name":    "name",
						"phone":   "phone",
						"subject": "subject",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", Optional, Update, publicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_agreement_acknowledged", "true"),
					resource.TestCheckResourceAttr(resourceName, "listing_type", "COMMUNITY"),
					resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription2"),
					resource.TestCheckResourceAttr(resourceName, "name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "package_details.0.eula.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "package_details.0.eula", map[string]string{
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
					CheckResourceSetContainsElementWithProperties(resourceName, "support_contacts", map[string]string{
						"email":   "email2",
						"name":    "name2",
						"phone":   "phone2",
						"subject": "subject2",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_marketplace_publications", "test_publications", Optional, Update, publicationDataSourceRepresentation) +
					compartmentIdVariableStr + PublicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", Optional, Update, publicationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_marketplace_publication", "test_publication", Required, Create, publicationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + PublicationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "publication_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "listing_type", "COMMUNITY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "long_description", "longDescription2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "name2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "package_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "short_description", "shortDescription2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "support_contacts.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "support_contacts", map[string]string{
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
		},
	})
}

func testAccCheckMarketplacePublicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).marketplaceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_marketplace_publication" {
			noResourceFound = false
			request := oci_marketplace.GetPublicationRequest{}

			tmp := rs.Primary.ID
			request.PublicationId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "marketplace")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("MarketplacePublication") {
		resource.AddTestSweepers("MarketplacePublication", &resource.Sweeper{
			Name:         "MarketplacePublication",
			Dependencies: DependencyGraph["publication"],
			F:            sweepMarketplacePublicationResource,
		})
	}
}

func sweepMarketplacePublicationResource(compartment string) error {
	marketplaceClient := GetTestClients(&schema.ResourceData{}).marketplaceClient()
	publicationIds, err := getPublicationIds(compartment)
	if err != nil {
		return err
	}
	for _, publicationId := range publicationIds {
		if ok := SweeperDefaultResourceId[publicationId]; !ok {
			deletePublicationRequest := oci_marketplace.DeletePublicationRequest{}

			deletePublicationRequest.PublicationId = &publicationId

			deletePublicationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "marketplace")
			_, error := marketplaceClient.DeletePublication(context.Background(), deletePublicationRequest)
			if error != nil {
				fmt.Printf("Error deleting Publication %s %s, It is possible that the resource is already deleted. Please verify manually \n", publicationId, error)
				continue
			}
			//waitTillCondition(testAccProvider, &publicationId, publicationSweepWaitCondition, time.Duration(3*time.Minute),
			//	publicationSweepResponseFetchOperation, "marketplace", true)
		}
	}
	return nil
}

func getPublicationIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PublicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	marketplaceClient := GetTestClients(&schema.ResourceData{}).marketplaceClient()

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "PublicationId", id)
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

func publicationSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.marketplaceClient().GetPublication(context.Background(), oci_marketplace.GetPublicationRequest{
		PublicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
