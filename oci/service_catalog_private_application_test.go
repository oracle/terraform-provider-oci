// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v54/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PrivateApplicationResourceConfig = PrivateApplicationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Optional, Update, privateApplicationRepresentation)

	privateApplicationSingularDataSourceRepresentation = map[string]interface{}{
		"private_application_id": Representation{RepType: Required, Create: `${oci_service_catalog_private_application.test_private_application.id}`},
	}

	privateApplicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":           Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"private_application_id": Representation{RepType: Optional, Create: `${oci_service_catalog_private_application.test_private_application.id}`},
		"filter":                 RepresentationGroup{Required, privateApplicationDataSourceFilterRepresentation}}
	privateApplicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_service_catalog_private_application.test_private_application.id}`}},
	}

	privateApplicationRepresentation = map[string]interface{}{
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":            Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"package_details":         RepresentationGroup{Required, privateApplicationPackageDetailsRepresentation},
		"short_description":       Representation{RepType: Required, Create: `shortDescription`, Update: `shortDescription2`},
		"defined_tags":            Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"logo_file_base64encoded": Representation{RepType: Optional, Create: `data:image/jpeg;base64,SWNvbkZvclRlcnJhZm9ybVRlc3Rpbmc=`, Update: `data:image/jpeg;base64,VXBkYXRlZEljb25Gb3JUZXJyYWZvcm1UZXN0aW5n`},
		"long_description":        Representation{RepType: Optional, Create: `longDescription`, Update: `longDescription2`},
	}
	privateApplicationPackageDetailsRepresentation = map[string]interface{}{
		"package_type":           Representation{RepType: Required, Create: `STACK`},
		"version":                Representation{RepType: Required, Create: `version`},
		"zip_file_base64encoded": Representation{RepType: Required, Create: `data:application/zip;base64,VGVzdERhdGFGb3JUZXJyYWZvcm0=`},
	}

	PrivateApplicationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: service_catalog/default
func TestServiceCatalogPrivateApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogPrivateApplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_service_catalog_private_application.test_private_application"
	datasourceName := "data.oci_service_catalog_private_applications.test_private_applications"
	singularDatasourceName := "data.oci_service_catalog_private_application.test_private_application"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+
		GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Optional, Create, privateApplicationRepresentation), "servicecatalog", "privateApplication", t)

	ResourceTest(t, testAccCheckServiceCatalogPrivateApplicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PrivateApplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Required, Create, privateApplicationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "STACK"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.version", "version"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.zip_file_base64encoded", "data:application/zip;base64,VGVzdERhdGFGb3JUZXJyYWZvcm0="),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PrivateApplicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PrivateApplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Optional, Create, privateApplicationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "logo_file_base64encoded", "data:image/jpeg;base64,SWNvbkZvclRlcnJhZm9ybVRlc3Rpbmc="),
				resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "STACK"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.version", "version"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.zip_file_base64encoded", "data:application/zip;base64,VGVzdERhdGFGb3JUZXJyYWZvcm0="),
				resource.TestCheckResourceAttrSet(resourceName, "package_type"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PrivateApplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Optional, Create,
					RepresentationCopyWithNewProperties(privateApplicationRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "logo_file_base64encoded", "data:image/jpeg;base64,SWNvbkZvclRlcnJhZm9ybVRlc3Rpbmc="),
				resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "STACK"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.version", "version"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.zip_file_base64encoded", "data:application/zip;base64,VGVzdERhdGFGb3JUZXJyYWZvcm0="),
				resource.TestCheckResourceAttrSet(resourceName, "package_type"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PrivateApplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Optional, Update, privateApplicationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "logo_file_base64encoded", "data:image/jpeg;base64,VXBkYXRlZEljb25Gb3JUZXJyYWZvcm1UZXN0aW5n"),
				resource.TestCheckResourceAttr(resourceName, "long_description", "longDescription2"),
				resource.TestCheckResourceAttr(resourceName, "package_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.package_type", "STACK"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.version", "version"),
				resource.TestCheckResourceAttr(resourceName, "package_details.0.zip_file_base64encoded", "data:application/zip;base64,VGVzdERhdGFGb3JUZXJyYWZvcm0="),
				resource.TestCheckResourceAttrSet(resourceName, "package_type"),
				resource.TestCheckResourceAttr(resourceName, "short_description", "shortDescription2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_service_catalog_private_applications", "test_private_applications", Optional, Update, privateApplicationDataSourceRepresentation) +
				compartmentIdVariableStr + PrivateApplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Optional, Update, privateApplicationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_application_id"),

				resource.TestCheckResourceAttr(datasourceName, "private_application_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_application_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", Required, Create, privateApplicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PrivateApplicationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_application_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logo.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "long_description", "longDescription2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "short_description", "shortDescription2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PrivateApplicationResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"logo_file_base64encoded",
				"package_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckServiceCatalogPrivateApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).serviceCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_catalog_private_application" {
			noResourceFound = false
			request := oci_service_catalog.GetPrivateApplicationRequest{}

			tmp := rs.Primary.ID
			request.PrivateApplicationId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "service_catalog")

			response, err := client.GetPrivateApplication(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_catalog.PrivateApplicationLifecycleStateDeleted): true,
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
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("ServiceCatalogPrivateApplication") {
		resource.AddTestSweepers("ServiceCatalogPrivateApplication", &resource.Sweeper{
			Name:         "ServiceCatalogPrivateApplication",
			Dependencies: DependencyGraph["privateApplication"],
			F:            sweepServiceCatalogPrivateApplicationResource,
		})
	}
}

func sweepServiceCatalogPrivateApplicationResource(compartment string) error {
	serviceCatalogClient := GetTestClients(&schema.ResourceData{}).serviceCatalogClient()
	privateApplicationIds, err := getPrivateApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, privateApplicationId := range privateApplicationIds {
		if ok := SweeperDefaultResourceId[privateApplicationId]; !ok {
			deletePrivateApplicationRequest := oci_service_catalog.DeletePrivateApplicationRequest{}

			deletePrivateApplicationRequest.PrivateApplicationId = &privateApplicationId

			deletePrivateApplicationRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "service_catalog")
			_, err := serviceCatalogClient.DeletePrivateApplication(context.Background(), deletePrivateApplicationRequest)
			if err != nil {
				fmt.Printf("Error deleting PrivateApplication %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateApplicationId, err)
				continue
			}
			WaitTillCondition(TestAccProvider, &privateApplicationId, privateApplicationSweepWaitCondition, time.Duration(3*time.Minute),
				privateApplicationSweepResponseFetchOperation, "service_catalog", true)
		}
	}
	return nil
}

func getPrivateApplicationIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "PrivateApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceCatalogClient := GetTestClients(&schema.ResourceData{}).serviceCatalogClient()

	listPrivateApplicationsRequest := oci_service_catalog.ListPrivateApplicationsRequest{}
	listPrivateApplicationsRequest.CompartmentId = &compartmentId
	listPrivateApplicationsResponse, err := serviceCatalogClient.ListPrivateApplications(context.Background(), listPrivateApplicationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PrivateApplication list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, privateApplication := range listPrivateApplicationsResponse.Items {
		id := *privateApplication.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "PrivateApplicationId", id)
	}
	return resourceIds, nil
}

func privateApplicationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if privateApplicationResponse, ok := response.Response.(oci_service_catalog.GetPrivateApplicationResponse); ok {
		return privateApplicationResponse.LifecycleState != oci_service_catalog.PrivateApplicationLifecycleStateDeleted
	}
	return false
}

func privateApplicationSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.serviceCatalogClient().GetPrivateApplication(context.Background(), oci_service_catalog.GetPrivateApplicationRequest{
		PrivateApplicationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
