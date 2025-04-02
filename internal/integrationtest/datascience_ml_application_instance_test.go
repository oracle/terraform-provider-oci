// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

type RepresentationGroupArray struct {
	RepType acctest.RepresentationType
	Group   []interface{}
}

var (
	DatascienceMlApplicationInstanceRequiredOnlyResource = DatascienceMlApplicationInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Required, acctest.Create, DatascienceMlApplicationInstanceRepresentation)

	DatascienceMlApplicationInstanceResourceConfig = DatascienceMlApplicationInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Optional, acctest.Update, DatascienceMlApplicationInstanceRepresentation)

	DatascienceMlApplicationInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"ml_application_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application_instance.test_ml_application_instance.id}`},
	}

	DatascienceMlApplicationInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"ml_application_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMlApplicationInstanceDataSourceFilterRepresentation}}
	DatascienceMlApplicationInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_ml_application_instance.test_ml_application_instance.id}`}},
	}

	DatascienceMlApplicationInstanceRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ml_application_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application.test_ml_application.id}`},
		"ml_application_implementation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_ml_application_implementation.test_ml_application_implementation.id}`},
		"auth_configuration":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceMlApplicationInstanceAuthConfigurationRepresentation},
		"configuration":                    RepresentationGroupArray{RepType: acctest.Optional, Group: DatascienceMlApplicationInstanceConfigurationRepresentation},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":                       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatascienceMlApplicationInstanceDefinedTagsChangesRepresentation},
	}

	ignoreDatascienceMlApplicationInstanceDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	DatascienceMlApplicationInstanceAuthConfigurationRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `IAM`},
	}
	DatascienceMlApplicationInstanceConfigurationRepresentation []interface{}

	DatascienceMlApplicationInstanceResourceDependencies = //acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Required, acctest.Create, DataflowApplicationRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_implementation", "test_ml_application_implementation", acctest.Optional, acctest.Create, DatascienceMlApplicationImplementationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application", "test_ml_application", acctest.Required, acctest.Create, DatascienceMlApplicationRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceMlApplicationInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceMlApplicationInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_ml_application_instance.test_ml_application_instance"
	datasourceName := "data.oci_datascience_ml_application_instances.test_ml_application_instances"
	singularDatasourceName := "data.oci_datascience_ml_application_instance.test_ml_application_instance"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceMlApplicationInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Optional, acctest.Create, DatascienceMlApplicationInstanceRepresentation), "datascience", "mlApplicationInstance", t)

	acctest.ResourceTest(t, testAccCheckDatascienceMlApplicationInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			PreConfig: func() {
				generateMLAppPackage("ml-app-package-v1", "1.1")
			},
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Required, acctest.Create, DatascienceMlApplicationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationInstanceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Optional, acctest.Create, DatascienceMlApplicationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auth_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auth_configuration.0.type", "IAM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_substate"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_implementation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_implementation_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_name"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceMlApplicationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceMlApplicationInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auth_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auth_configuration.0.type", "IAM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_substate"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_implementation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_implementation_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_name"),
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
			Config: config + compartmentIdVariableStr + DatascienceMlApplicationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Optional, acctest.Update, DatascienceMlApplicationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auth_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "auth_configuration.0.type", "IAM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_substate"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_implementation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_implementation_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ml_application_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application_instances", "test_ml_application_instances", acctest.Optional, acctest.Update, DatascienceMlApplicationInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlApplicationInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Optional, acctest.Update, DatascienceMlApplicationInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "ml_application_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ml_application_instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ml_application_instance_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_ml_application_instance", "test_ml_application_instance", acctest.Required, acctest.Create, DatascienceMlApplicationInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMlApplicationInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ml_application_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auth_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auth_configuration.0.type", "IAM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_substate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ml_application_implementation_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ml_application_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "prediction_endpoint_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceMlApplicationInstanceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceMlApplicationInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_ml_application_instance" {
			noResourceFound = false
			request := oci_datascience.GetMlApplicationInstanceRequest{}

			tmp := rs.Primary.ID
			request.MlApplicationInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetMlApplicationInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.MlApplicationInstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceMlApplicationInstance") {
		resource.AddTestSweepers("DatascienceMlApplicationInstance", &resource.Sweeper{
			Name:         "DatascienceMlApplicationInstance",
			Dependencies: acctest.DependencyGraph["mlApplicationInstance"],
			F:            sweepDatascienceMlApplicationInstanceResource,
		})
	}
}

func sweepDatascienceMlApplicationInstanceResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	mlApplicationInstanceIds, err := getDatascienceMlApplicationInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, mlApplicationInstanceId := range mlApplicationInstanceIds {
		if ok := acctest.SweeperDefaultResourceId[mlApplicationInstanceId]; !ok {
			deleteMlApplicationInstanceRequest := oci_datascience.DeleteMlApplicationInstanceRequest{}

			deleteMlApplicationInstanceRequest.MlApplicationInstanceId = &mlApplicationInstanceId

			deleteMlApplicationInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteMlApplicationInstance(context.Background(), deleteMlApplicationInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting MlApplicationInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", mlApplicationInstanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mlApplicationInstanceId, DatascienceMlApplicationInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceMlApplicationInstanceSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceMlApplicationInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MlApplicationInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listMlApplicationInstancesRequest := oci_datascience.ListMlApplicationInstancesRequest{}
	listMlApplicationInstancesRequest.CompartmentId = &compartmentId
	listMlApplicationInstancesRequest.LifecycleState = oci_datascience.MlApplicationInstanceLifecycleStateActive
	listMlApplicationInstancesResponse, err := dataScienceClient.ListMlApplicationInstances(context.Background(), listMlApplicationInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MlApplicationInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mlApplicationInstance := range listMlApplicationInstancesResponse.Items {
		id := *mlApplicationInstance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MlApplicationInstanceId", id)
	}
	return resourceIds, nil
}

func DatascienceMlApplicationInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mlApplicationInstanceResponse, ok := response.Response.(oci_datascience.GetMlApplicationInstanceResponse); ok {
		return mlApplicationInstanceResponse.LifecycleState != oci_datascience.MlApplicationInstanceLifecycleStateDeleted
	}
	return false
}

func DatascienceMlApplicationInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetMlApplicationInstance(context.Background(), oci_datascience.GetMlApplicationInstanceRequest{
		MlApplicationInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
