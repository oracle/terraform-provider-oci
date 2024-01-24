// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceModelVersionSetRequiredOnlyResource = DatascienceModelVersionSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_version_set", "test_model_version_set", acctest.Required, acctest.Create, DatascienceModelVersionSetRepresentation)

	DatascienceModelVersionSetResourceConfig = DatascienceModelVersionSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_version_set", "test_model_version_set", acctest.Optional, acctest.Update, DatascienceModelVersionSetRepresentation)

	DatascienceDatascienceModelVersionSetSingularDataSourceRepresentation = map[string]interface{}{
		"model_version_set_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model_version_set.test_model_version_set.id}`},
	}

	DatascienceDatascienceModelVersionSetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `createdBy`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_version_set.test_model_version_set.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name` + utils.RandomString(15, utils.CharsetWithoutDigits)},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelVersionSetDataSourceFilterRepresentation}}
	DatascienceModelVersionSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model_version_set.test_model_version_set.id}`}},
	}

	DatascienceModelVersionSetRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name` + utils.RandomString(15, utils.CharsetWithoutDigits)},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatascienceModelVersionSetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceModelVersionSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelVersionSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_version_set.test_model_version_set"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceModelVersionSetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_version_set", "test_model_version_set", acctest.Optional, acctest.Create, DatascienceModelVersionSetRepresentation), "datascience", "modelVersionSet", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelVersionSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelVersionSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_version_set", "test_model_version_set", acctest.Required, acctest.Create, DatascienceModelVersionSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelVersionSetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceModelVersionSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_version_set", "test_model_version_set", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceModelVersionSetRepresentation, map[string]interface{}{
						"name": acctest.Representation{RepType: acctest.Required, Create: `name` + utils.RandomString(15, utils.CharsetWithoutDigits)},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
		// verify resource import
		{
			Config:                  config + DatascienceModelVersionSetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceModelVersionSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model_version_set" {
			noResourceFound = false
			request := oci_datascience.GetModelVersionSetRequest{}

			tmp := rs.Primary.ID
			request.ModelVersionSetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetModelVersionSet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelVersionSetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceModelVersionSet") {
		resource.AddTestSweepers("DatascienceModelVersionSet", &resource.Sweeper{
			Name:         "DatascienceModelVersionSet",
			Dependencies: acctest.DependencyGraph["modelVersionSet"],
			F:            sweepDatascienceModelVersionSetResource,
		})
	}
}

func sweepDatascienceModelVersionSetResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	modelVersionSetIds, err := getDatascienceModelVersionSetIds(compartment)
	if err != nil {
		return err
	}
	for _, modelVersionSetId := range modelVersionSetIds {
		if ok := acctest.SweeperDefaultResourceId[modelVersionSetId]; !ok {
			deleteModelVersionSetRequest := oci_datascience.DeleteModelVersionSetRequest{}

			deleteModelVersionSetRequest.ModelVersionSetId = &modelVersionSetId

			deleteModelVersionSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModelVersionSet(context.Background(), deleteModelVersionSetRequest)
			if error != nil {
				fmt.Printf("Error deleting ModelVersionSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelVersionSetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelVersionSetId, DatascienceModelVersionSetSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceModelVersionSetSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceModelVersionSetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelVersionSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listModelVersionSetsRequest := oci_datascience.ListModelVersionSetsRequest{}
	listModelVersionSetsRequest.CompartmentId = &compartmentId
	listModelVersionSetsRequest.LifecycleState = oci_datascience.ListModelVersionSetsLifecycleStateActive
	listModelVersionSetsResponse, err := dataScienceClient.ListModelVersionSets(context.Background(), listModelVersionSetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ModelVersionSet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, modelVersionSet := range listModelVersionSetsResponse.Items {
		id := *modelVersionSet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelVersionSetId", id)
	}
	return resourceIds, nil
}

func DatascienceModelVersionSetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelVersionSetResponse, ok := response.Response.(oci_datascience.GetModelVersionSetResponse); ok {
		return modelVersionSetResponse.LifecycleState != oci_datascience.ModelVersionSetLifecycleStateDeleted
	}
	return false
}

func DatascienceModelVersionSetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetModelVersionSet(context.Background(), oci_datascience.GetModelVersionSetRequest{
		ModelVersionSetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
