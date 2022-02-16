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
	oci_service_catalog "github.com/oracle/oci-go-sdk/v58/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ServiceCatalogAssociationRequiredOnlyResource = ServiceCatalogAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", acctest.Required, acctest.Create, serviceCatalogAssociationRepresentation)

	ServiceCatalogAssociationResourceConfig = ServiceCatalogAssociationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", acctest.Optional, acctest.Update, serviceCatalogAssociationRepresentation)

	serviceCatalogAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"service_catalog_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`},
	}

	serviceCatalogAssociationDataSourceRepresentation = map[string]interface{}{
		"entity_id":                      acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.entity_id}`},
		"entity_type":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.entity_type}`},
		"service_catalog_association_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`},
		"service_catalog_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceCatalogAssociationDataSourceFilterRepresentation}}
	serviceCatalogAssociationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_catalog_service_catalog_association.test_service_catalog_association.id}`}},
	}

	serviceCatalogAssociationRepresentation = map[string]interface{}{
		"entity_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_service_catalog_private_application.test_private_application.id}`},
		"service_catalog_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"entity_type":        acctest.Representation{RepType: acctest.Optional, Create: `privateapplication`},
	}

	ServiceCatalogAssociationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", acctest.Required, acctest.Create, privateApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Required, acctest.Create, serviceCatalogRepresentation)
)

// issue-routing-tag: service_catalog/default
func TestServiceCatalogServiceCatalogAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogServiceCatalogAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_service_catalog_service_catalog_association.test_service_catalog_association"
	datasourceName := "data.oci_service_catalog_service_catalog_associations.test_service_catalog_associations"
	singularDatasourceName := "data.oci_service_catalog_service_catalog_association.test_service_catalog_association"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ServiceCatalogAssociationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", acctest.Optional, acctest.Create, serviceCatalogAssociationRepresentation), "servicecatalog", "serviceCatalogAssociation", t)

	acctest.ResourceTest(t, testAccCheckServiceCatalogServiceCatalogAssociationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", acctest.Required, acctest.Create, serviceCatalogAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", acctest.Optional, acctest.Create, serviceCatalogAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "entity_type", "privateapplication"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "service_catalog_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog_associations", "test_service_catalog_associations", acctest.Optional, acctest.Update, serviceCatalogAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogAssociationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", acctest.Optional, acctest.Update, serviceCatalogAssociationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog_association", "test_service_catalog_association", acctest.Required, acctest.Create, serviceCatalogAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_catalog_service_catalog_association" {
			noResourceFound = false
			request := oci_service_catalog.GetServiceCatalogAssociationRequest{}

			tmp := rs.Primary.ID
			request.ServiceCatalogAssociationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_catalog")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ServiceCatalogServiceCatalogAssociation") {
		resource.AddTestSweepers("ServiceCatalogServiceCatalogAssociation", &resource.Sweeper{
			Name:         "ServiceCatalogServiceCatalogAssociation",
			Dependencies: acctest.DependencyGraph["serviceCatalogAssociation"],
			F:            sweepServiceCatalogServiceCatalogAssociationResource,
		})
	}
}

func sweepServiceCatalogServiceCatalogAssociationResource(compartment string) error {
	serviceCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceCatalogClient()
	serviceCatalogAssociationIds, err := getServiceCatalogAssociationIds(compartment)
	if err != nil {
		return err
	}
	for _, serviceCatalogAssociationId := range serviceCatalogAssociationIds {
		if ok := acctest.SweeperDefaultResourceId[serviceCatalogAssociationId]; !ok {
			deleteServiceCatalogAssociationRequest := oci_service_catalog.DeleteServiceCatalogAssociationRequest{}

			deleteServiceCatalogAssociationRequest.ServiceCatalogAssociationId = &serviceCatalogAssociationId

			deleteServiceCatalogAssociationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_catalog")
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
	ids := acctest.GetResourceIdsToSweep(compartment, "ServiceCatalogAssociationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceCatalogClient()

	listServiceCatalogAssociationsRequest := oci_service_catalog.ListServiceCatalogAssociationsRequest{}
	//listServiceCatalogAssociationsRequest.CompartmentId = &compartmentId
	listServiceCatalogAssociationsResponse, err := serviceCatalogClient.ListServiceCatalogAssociations(context.Background(), listServiceCatalogAssociationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ServiceCatalogAssociation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, serviceCatalogAssociation := range listServiceCatalogAssociationsResponse.Items {
		id := *serviceCatalogAssociation.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ServiceCatalogAssociationId", id)
	}
	return resourceIds, nil
}
