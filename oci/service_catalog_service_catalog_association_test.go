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
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v52/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ServiceCatalogAssociationRequiredOnlyResource = ServiceCatalogAssociationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Required, Create, serviceCatalogAssociationRepresentation)

	ServiceCatalogAssociationResourceConfig = ServiceCatalogAssociationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Update, serviceCatalogAssociationRepresentation)

	serviceCatalogAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"service_catalog_association_id": Representation{RepType: Required, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`},
	}

	serviceCatalogAssociationDataSourceRepresentation = map[string]interface{}{
		"entity_id":                      Representation{RepType: Optional, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.entity_id}`},
		"entity_type":                    Representation{RepType: Optional, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.entity_type}`},
		"service_catalog_association_id": Representation{RepType: Optional, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`},
		"service_catalog_id":             Representation{RepType: Optional, Create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"filter":                         RepresentationGroup{Required, serviceCatalogAssociationDataSourceFilterRepresentation}}
	serviceCatalogAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`}},
	}

	serviceCatalogAssociationRepresentation = map[string]interface{}{
		"entity_id":          Representation{RepType: Required, Create: `${oci_service_catalog_private_application.test_private_application.id}`},
		"service_catalog_id": Representation{RepType: Required, Create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"entity_type":        Representation{RepType: Optional, Create: `privateapplication`},
	}

	ServiceCatalogAssociationResourceDependencies = GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Required, Create, privateApplicationRepresentation) +
		GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Required, Create, serviceCatalogRepresentation)
)

// issue-routing-tag: service_catalog/default
func TestServiceCatalogServiceCatalogAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogServiceCatalogAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_service_catalog_service_catalog_association.test_service_catalog_association"
	datasourceName := "data.oci_service_catalog_service_catalog_associations.test_service_catalog_associations"
	singularDatasourceName := "data.oci_service_catalog_service_catalog_association.test_service_catalog_association"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ServiceCatalogAssociationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Create, serviceCatalogAssociationRepresentation), "servicecatalog", "serviceCatalogAssociation", t)

	ResourceTest(t, testAccCheckServiceCatalogServiceCatalogAssociationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Required, Create, serviceCatalogAssociationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_catalog_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Create, serviceCatalogAssociationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "entity_type", "privateapplication"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_catalog_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog_associations", "test_service_catalog_associations", Optional, Update, serviceCatalogAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Optional, Update, serviceCatalogAssociationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", Required, Create, serviceCatalogAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogAssociationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "service_catalog")

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
	if !InSweeperExcludeList("ServiceCatalogServiceCatalogAssociation") {
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

			deleteServiceCatalogAssociationRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "service_catalog")
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
	ids := GetResourceIdsToSweep(compartment, "ServiceCatalogAssociationId")
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
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ServiceCatalogAssociationId", id)
	}
	return resourceIds, nil
}
