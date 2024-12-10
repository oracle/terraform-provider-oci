// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceMlApplicationRequiredOnlyResource = DatascienceMlApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Required, acctest.Create, DatascienceMlApplicationRepresentation)

	DatascienceMlApplicationResourceConfig = DatascienceMlApplicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Optional, acctest.Update, DatascienceMlApplicationRepresentation)

	DatascienceMlApplicationSingularDataSourceRepresentation = map[string]interface{}{
		"ml_application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
	}

	DatascienceMlApplicationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"ml_application_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `ml-app-name-test`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMlApplicationDataSourceFilterRepresentation},
	}

	DatascienceMlApplicationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_ml_application.test_ml_application.id}`}},
	}

	DatascienceMlApplicationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `ml-app-name-test`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description-update`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatascienceMlApplicationDefinedTagsChangesRepresentation},
	}

	ignoreDatascienceMlApplicationDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	DatascienceMlApplicationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceMlApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceMlApplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_ml_application.test_ml_application"
	datasourceName := "data.oci_datascience_ml_applications.test_ml_applications"
	singularDatasourceName := "data.oci_datascience_ml_application.test_ml_application"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceMlApplicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Optional, acctest.Create, DatascienceMlApplicationRepresentation), "datascience", "mlApplication", t)

	acctest.ResourceTest(t, testAccCheckDatascienceMlApplicationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Required, acctest.Create, DatascienceMlApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "ml-app-name-test"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Optional, acctest.Create, DatascienceMlApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "ml-app-name-test"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceMlApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceMlApplicationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "ml-app-name-test"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Optional, acctest.Update, DatascienceMlApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description-update"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "ml-app-name-test"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_applications", "test_ml_applications", acctest.Optional, acctest.Update, DatascienceMlApplicationDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlApplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Optional, acctest.Update, DatascienceMlApplicationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "ml_application_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "ml-app-name-test"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ml_application_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ml_application_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Required, acctest.Create, DatascienceMlApplicationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlApplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ml_application_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description-update"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "ml-app-name-test"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceMlApplicationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceMlApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_ml_application" {
			noResourceFound = false
			request := oci_datascience.GetMlApplicationRequest{}

			tmp := rs.Primary.ID
			request.MlApplicationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			_, err := client.GetMlApplication(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DatascienceMlApplication") {
		resource.AddTestSweepers("DatascienceMlApplication", &resource.Sweeper{
			Name:         "DatascienceMlApplication",
			Dependencies: acctest.DependencyGraph["mlApplication"],
			F:            sweepDatascienceMlApplicationResource,
		})
	}
}

func sweepDatascienceMlApplicationResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	mlApplicationIds, err := getDatascienceMlApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, mlApplicationId := range mlApplicationIds {
		if ok := acctest.SweeperDefaultResourceId[mlApplicationId]; !ok {
			deleteMlApplicationRequest := oci_datascience.DeleteMlApplicationRequest{}

			deleteMlApplicationRequest.MlApplicationId = &mlApplicationId

			deleteMlApplicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteMlApplication(context.Background(), deleteMlApplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting MlApplication %s %s, It is possible that the resource is already deleted. Please verify manually \n", mlApplicationId, error)
				continue
			}
		}
	}
	return nil
}

func getDatascienceMlApplicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MlApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listMlApplicationsRequest := oci_datascience.ListMlApplicationsRequest{}
	listMlApplicationsRequest.CompartmentId = &compartmentId
	listMlApplicationsResponse, err := dataScienceClient.ListMlApplications(context.Background(), listMlApplicationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MlApplication list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mlApplication := range listMlApplicationsResponse.Items {
		id := *mlApplication.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MlApplicationId", id)
	}
	return resourceIds, nil
}
