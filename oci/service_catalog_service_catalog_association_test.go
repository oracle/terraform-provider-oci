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
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v40/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ServiceCatalogAssociationRequiredOnlyResource = ServiceCatalogAssociationResourceDependencies +
		generateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Required, Create, serviceCatalogAssociationRepresentation)

	ServiceCatalogAssociationResourceConfig = ServiceCatalogAssociationResourceDependencies +
		generateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Update, serviceCatalogAssociationRepresentation)

	serviceCatalogAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"service_catalog_association_id": Representation{repType: Required, create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`},
	}

	serviceCatalogAssociationDataSourceRepresentation = map[string]interface{}{
		"entity_id":                      Representation{repType: Optional, create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.entity_id}`},
		"entity_type":                    Representation{repType: Optional, create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.entity_type}`},
		"service_catalog_association_id": Representation{repType: Optional, create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`},
		"service_catalog_id":             Representation{repType: Optional, create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"filter":                         RepresentationGroup{Required, serviceCatalogAssociationDataSourceFilterRepresentation}}
	serviceCatalogAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`}},
	}

	serviceCatalogAssociationRepresentation = map[string]interface{}{
		"entity_id":          Representation{repType: Required, create: `${oci_service_catalog_private_application.test_private_application.id}`},
		"service_catalog_id": Representation{repType: Required, create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"entity_type":        Representation{repType: Optional, create: `privateapplication`},
	}

	ServiceCatalogAssociationResourceDependencies = generateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Required, Create, privateApplicationRepresentation) +
		generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Required, Create, serviceCatalogRepresentation)
)

func TestServiceCatalogServiceCatalogAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogServiceCatalogAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_service_catalog_service_catalog_association.test_service_catalog_association"
	datasourceName := "data.oci_service_catalog_service_catalog_associations.test_service_catalog_associations"
	singularDatasourceName := "data.oci_service_catalog_service_catalog_association.test_service_catalog_association"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ServiceCatalogAssociationResourceDependencies+
		generateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Create, serviceCatalogAssociationRepresentation), "servicecatalog", "serviceCatalogAssociation", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckServiceCatalogServiceCatalogAssociationDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
					generateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Required, Create, serviceCatalogAssociationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "service_catalog_id"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
					generateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Create, serviceCatalogAssociationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
					resource.TestCheckResourceAttr(resourceName, "entity_type", "privateapplication"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "service_catalog_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog_associations", "test_service_catalog_associations", Optional, Update, serviceCatalogAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
					generateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Update, serviceCatalogAssociationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "entity_id"),
					resource.TestCheckResourceAttr(datasourceName, "entity_type", "privateapplication"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_catalog_association_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_catalog_id"),

					resource.TestCheckResourceAttr(datasourceName, "service_catalog_association_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "service_catalog_association_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Required, Create, serviceCatalogAssociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ServiceCatalogAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "service_catalog_association_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "entity_type", "privateapplication"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckServiceCatalogServiceCatalogAssociationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).serviceCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_catalog_service_catalog_association" {
			noResourceFound = false
			request := oci_service_catalog.GetServiceCatalogAssociationRequest{}

			tmp := rs.Primary.ID
			request.ServiceCatalogAssociationId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "service_catalog")

			_, err := client.GetServiceCatalogAssociation(context.Background(), request)

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ServiceCatalogServiceCatalogAssociation") {
		resource.AddTestSweepers("ServiceCatalogServiceCatalogAssociation", &resource.Sweeper{
			Name:         "ServiceCatalogServiceCatalogAssociation",
			Dependencies: DependencyGraph["serviceCatalogAssociation"],
			F:            sweepServiceCatalogServiceCatalogAssociationResource,
		})
	}
}

func sweepServiceCatalogServiceCatalogAssociationResource(compartment string) error {
	serviceCatalogClient := GetTestClients(&schema.ResourceData{}).serviceCatalogClient()
	serviceCatalogAssociationIds, err := getServiceCatalogAssociationIds(compartment)
	if err != nil {
		return err
	}
	for _, serviceCatalogAssociationId := range serviceCatalogAssociationIds {
		if ok := SweeperDefaultResourceId[serviceCatalogAssociationId]; !ok {
			deleteServiceCatalogAssociationRequest := oci_service_catalog.DeleteServiceCatalogAssociationRequest{}

			deleteServiceCatalogAssociationRequest.ServiceCatalogAssociationId = &serviceCatalogAssociationId

			deleteServiceCatalogAssociationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "service_catalog")
			_, error := serviceCatalogClient.DeleteServiceCatalogAssociation(context.Background(), deleteServiceCatalogAssociationRequest)
			if error != nil {
				fmt.Printf("Error deleting ServiceCatalogAssociation %s %s, It is possible that the resource is already deleted. Please verify manually \n", serviceCatalogAssociationId, error)
				continue
			}
		}
	}
	return nil
}

func getServiceCatalogAssociationIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ServiceCatalogAssociationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceCatalogClient := GetTestClients(&schema.ResourceData{}).serviceCatalogClient()

	listServiceCatalogAssociationsRequest := oci_service_catalog.ListServiceCatalogAssociationsRequest{}
	//listServiceCatalogAssociationsRequest.CompartmentId = &compartmentId
	listServiceCatalogAssociationsResponse, err := serviceCatalogClient.ListServiceCatalogAssociations(context.Background(), listServiceCatalogAssociationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ServiceCatalogAssociation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, serviceCatalogAssociation := range listServiceCatalogAssociationsResponse.Items {
		id := *serviceCatalogAssociation.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ServiceCatalogAssociationId", id)
	}
	return resourceIds, nil
}
