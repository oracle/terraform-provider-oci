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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSdmMaskingPolicyDifferenceRequiredOnlyResource = DataSafeSdmMaskingPolicyDifferenceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Required, acctest.Create, DataSafeSdmMaskingPolicyDifferenceRepresentation)

	DataSafeSdmMaskingPolicyDifferenceResourceConfig = DataSafeSdmMaskingPolicyDifferenceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Optional, acctest.Update, DataSafeSdmMaskingPolicyDifferenceRepresentation)

	DataSafeSdmMaskingPolicyDifferenceSingularDataSourceRepresentation = map[string]interface{}{
		"sdm_masking_policy_difference_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id}`},
	}

	DataSafeSdmMaskingPolicyDifferenceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeSdmMaskingPolicyDifferenceDataSourceFilterRepresentation}}
	DataSafeSdmMaskingPolicyDifferenceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id}`}},
	}

	DataSafeSdmMaskingPolicyDifferenceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"masking_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"difference_type":   acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDifferenceSystemTagsChangesRep},
	}

	DataSafeSdmMaskingPolicyDifferenceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Required, acctest.Create, maskingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model1", acctest.Required, acctest.Create, sensitiveDataModelRepresentation) +
		DefinedTagsDependencies

	ignoreDifferenceSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSdmMaskingPolicyDifferenceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSdmMaskingPolicyDifferenceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference"
	datasourceName := "data.oci_data_safe_sdm_masking_policy_differences.test_sdm_masking_policy_differences"
	singularDatasourceName := "data.oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSdmMaskingPolicyDifferenceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Optional, acctest.Create, DataSafeSdmMaskingPolicyDifferenceRepresentation), "datasafe", "sdmMaskingPolicyDifference", t)

	acctest.ResourceTest(t, testAccCheckDataSafeSdmMaskingPolicyDifferenceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Required, acctest.Create, DataSafeSdmMaskingPolicyDifferenceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceResourceDependencies + targetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Optional, acctest.Create, DataSafeSdmMaskingPolicyDifferenceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "difference_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_creation_started"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeSdmMaskingPolicyDifferenceResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSdmMaskingPolicyDifferenceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "difference_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_creation_started"),

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
			Config: config + compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Optional, acctest.Update, DataSafeSdmMaskingPolicyDifferenceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "difference_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "masking_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_creation_started"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_differences", "test_sdm_masking_policy_differences", acctest.Optional, acctest.Update, DataSafeSdmMaskingPolicyDifferenceDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Optional, acctest.Update, DataSafeSdmMaskingPolicyDifferenceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sdm_masking_policy_difference_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Required, acctest.Create, DataSafeSdmMaskingPolicyDifferenceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSdmMaskingPolicyDifferenceResourceConfig + targetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sdm_masking_policy_difference_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "difference_type", "ALL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_creation_started"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeSdmMaskingPolicyDifferenceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeSdmMaskingPolicyDifferenceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_sdm_masking_policy_difference" {
			noResourceFound = false
			request := oci_data_safe.GetSdmMaskingPolicyDifferenceRequest{}

			tmp := rs.Primary.ID
			request.SdmMaskingPolicyDifferenceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetSdmMaskingPolicyDifference(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeSdmMaskingPolicyDifference") {
		resource.AddTestSweepers("DataSafeSdmMaskingPolicyDifference", &resource.Sweeper{
			Name:         "DataSafeSdmMaskingPolicyDifference",
			Dependencies: acctest.DependencyGraph["sdmMaskingPolicyDifference"],
			F:            sweepDataSafeSdmMaskingPolicyDifferenceResource,
		})
	}
}

func sweepDataSafeSdmMaskingPolicyDifferenceResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sdmMaskingPolicyDifferenceIds, err := getDataSafeSdmMaskingPolicyDifferenceIds(compartment)
	if err != nil {
		return err
	}
	for _, sdmMaskingPolicyDifferenceId := range sdmMaskingPolicyDifferenceIds {
		if ok := acctest.SweeperDefaultResourceId[sdmMaskingPolicyDifferenceId]; !ok {
			deleteSdmMaskingPolicyDifferenceRequest := oci_data_safe.DeleteSdmMaskingPolicyDifferenceRequest{}

			deleteSdmMaskingPolicyDifferenceRequest.SdmMaskingPolicyDifferenceId = &sdmMaskingPolicyDifferenceId

			deleteSdmMaskingPolicyDifferenceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSdmMaskingPolicyDifference(context.Background(), deleteSdmMaskingPolicyDifferenceRequest)
			if error != nil {
				fmt.Printf("Error deleting SdmMaskingPolicyDifference %s %s, It is possible that the resource is already deleted. Please verify manually \n", sdmMaskingPolicyDifferenceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sdmMaskingPolicyDifferenceId, DataSafeSdmMaskingPolicyDifferenceSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSdmMaskingPolicyDifferenceSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSdmMaskingPolicyDifferenceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SdmMaskingPolicyDifferenceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSdmMaskingPolicyDifferencesRequest := oci_data_safe.ListSdmMaskingPolicyDifferencesRequest{}
	listSdmMaskingPolicyDifferencesRequest.CompartmentId = &compartmentId
	listSdmMaskingPolicyDifferencesRequest.LifecycleState = oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateActive
	listSdmMaskingPolicyDifferencesResponse, err := dataSafeClient.ListSdmMaskingPolicyDifferences(context.Background(), listSdmMaskingPolicyDifferencesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SdmMaskingPolicyDifference list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sdmMaskingPolicyDifference := range listSdmMaskingPolicyDifferencesResponse.Items {
		id := *sdmMaskingPolicyDifference.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SdmMaskingPolicyDifferenceId", id)
	}
	return resourceIds, nil
}

func DataSafeSdmMaskingPolicyDifferenceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sdmMaskingPolicyDifferenceResponse, ok := response.Response.(oci_data_safe.GetSdmMaskingPolicyDifferenceResponse); ok {
		return sdmMaskingPolicyDifferenceResponse.LifecycleState != oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateDeleted
	}
	return false
}

func DataSafeSdmMaskingPolicyDifferenceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSdmMaskingPolicyDifference(context.Background(), oci_data_safe.GetSdmMaskingPolicyDifferenceRequest{
		SdmMaskingPolicyDifferenceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
