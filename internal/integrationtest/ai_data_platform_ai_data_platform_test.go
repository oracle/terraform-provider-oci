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
	oci_ai_data_platform "github.com/oracle/oci-go-sdk/v65/aidataplatform"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// Developer Notes
// Disabled the defined tag dependency since it can only be created from the PHX home region, which is currently unstable,
// Defined-tags usage are tested separately
// Created ticket https://jira.oci.oraclecorp.com/browse/DATAHUB-23129 for investigation for ai_data_platform_type issue, commented it out for testing currently
var (
	AiDataPlatformAiDataPlatformRequiredOnlyResource = AiDataPlatformAiDataPlatformResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Required, acctest.Create, AiDataPlatformAiDataPlatformRepresentation)

	AiDataPlatformAiDataPlatformResourceConfig = AiDataPlatformAiDataPlatformResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Optional, acctest.Update, AiDataPlatformAiDataPlatformRepresentation)

	AiDataPlatformAiDataPlatformSingularDataSourceRepresentation = map[string]interface{}{
		"ai_data_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_data_platform_ai_data_platform.test_ai_data_platform.id}`},
	}

	AiDataPlatformAiDataPlatformDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"exclude_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `CREATING`},
		"id":                      acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_data_platform_ai_data_platform.test_ai_data_platform.id}`},
		"include_legacy":          acctest.Representation{RepType: acctest.Optional, Create: `includeLegacy`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDataPlatformAiDataPlatformDataSourceFilterRepresentation}}
	AiDataPlatformAiDataPlatformDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_data_platform_ai_data_platform.test_ai_data_platform.id}`}},
	}

	AiDataPlatformAiDataPlatformRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//"ai_data_platform_type":  acctest.Representation{RepType: acctest.Optional, Create: `aiDataPlatformType`},
		"default_workspace_name": acctest.Representation{RepType: acctest.Optional, Create: `test_workspace`},
		//"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	//AiDataPlatformAiDataPlatformResourceDependencies = DefinedTagsDependencies
	AiDataPlatformAiDataPlatformResourceDependencies = ""
)

// test
// issue-routing-tag: ai_data_platform/default
func TestAiDataPlatformAiDataPlatformResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiDataPlatformAiDataPlatformResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_id_for_create")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_data_platform_ai_data_platform.test_ai_data_platform"
	datasourceName := "data.oci_ai_data_platform_ai_data_platforms.test_ai_data_platforms"
	singularDatasourceName := "data.oci_ai_data_platform_ai_data_platform.test_ai_data_platform"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiDataPlatformAiDataPlatformResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Optional, acctest.Create, AiDataPlatformAiDataPlatformRepresentation), "aidataplatform", "aiDataPlatform", t)

	acctest.ResourceTest(t, testAccCheckAiDataPlatformAiDataPlatformDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + AiDataPlatformAiDataPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Required, acctest.Create, AiDataPlatformAiDataPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiDataPlatformAiDataPlatformResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiDataPlatformAiDataPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Optional, acctest.Create, AiDataPlatformAiDataPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "ai_data_platform_type", "aiDataPlatformType"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "default_workspace_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiDataPlatformAiDataPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiDataPlatformAiDataPlatformRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "ai_data_platform_type", "aiDataPlatformType"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "default_workspace_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + AiDataPlatformAiDataPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Optional, acctest.Update, AiDataPlatformAiDataPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "ai_data_platform_type", "aiDataPlatformType2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "default_workspace_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_data_platform_ai_data_platforms", "test_ai_data_platforms", acctest.Optional, acctest.Update, AiDataPlatformAiDataPlatformDataSourceRepresentation) +
				compartmentIdVariableStr + AiDataPlatformAiDataPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Optional, acctest.Update, AiDataPlatformAiDataPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "exclude_lifecycle_state", "CREATING"),
				resource.TestCheckResourceAttr(datasourceName, "include_legacy", "includeLegacy"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "ai_data_platform_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ai_data_platform_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_data_platform_ai_data_platform", "test_ai_data_platform", acctest.Required, acctest.Create, AiDataPlatformAiDataPlatformSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiDataPlatformAiDataPlatformResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ai_data_platform_id"),

				//resource.TestCheckResourceAttr(singularDatasourceName, "ai_data_platform_type", "aiDataPlatformType2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alias_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + AiDataPlatformAiDataPlatformRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"default_workspace_name",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckAiDataPlatformAiDataPlatformDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiDataPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_data_platform_ai_data_platform" {
			noResourceFound = false
			request := oci_ai_data_platform.GetAiDataPlatformRequest{}

			tmp := rs.Primary.ID
			request.AiDataPlatformId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_data_platform")

			response, err := client.GetAiDataPlatform(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_data_platform.AiDataPlatformLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiDataPlatformAiDataPlatform") {
		resource.AddTestSweepers("AiDataPlatformAiDataPlatform", &resource.Sweeper{
			Name:         "AiDataPlatformAiDataPlatform",
			Dependencies: acctest.DependencyGraph["aiDataPlatform"],
			F:            sweepAiDataPlatformAiDataPlatformResource,
		})
	}
}

func sweepAiDataPlatformAiDataPlatformResource(compartment string) error {
	aiDataPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).AiDataPlatformClient()
	aiDataPlatformIds, err := getAiDataPlatformAiDataPlatformIds(compartment)
	if err != nil {
		return err
	}
	for _, aiDataPlatformId := range aiDataPlatformIds {
		if ok := acctest.SweeperDefaultResourceId[aiDataPlatformId]; !ok {
			deleteAiDataPlatformRequest := oci_ai_data_platform.DeleteAiDataPlatformRequest{}

			deleteAiDataPlatformRequest.AiDataPlatformId = &aiDataPlatformId

			deleteAiDataPlatformRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_data_platform")
			_, error := aiDataPlatformClient.DeleteAiDataPlatform(context.Background(), deleteAiDataPlatformRequest)
			if error != nil {
				fmt.Printf("Error deleting AiDataPlatform %s %s, It is possible that the resource is already deleted. Please verify manually \n", aiDataPlatformId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &aiDataPlatformId, AiDataPlatformAiDataPlatformSweepWaitCondition, time.Duration(3*time.Minute),
				AiDataPlatformAiDataPlatformSweepResponseFetchOperation, "ai_data_platform", true)
		}
	}
	return nil
}

func getAiDataPlatformAiDataPlatformIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AiDataPlatformId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiDataPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).AiDataPlatformClient()

	listAiDataPlatformsRequest := oci_ai_data_platform.ListAiDataPlatformsRequest{}
	listAiDataPlatformsRequest.CompartmentId = &compartmentId
	listAiDataPlatformsRequest.LifecycleState = oci_ai_data_platform.AiDataPlatformLifecycleStateActive
	listAiDataPlatformsResponse, err := aiDataPlatformClient.ListAiDataPlatforms(context.Background(), listAiDataPlatformsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AiDataPlatform list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, aiDataPlatform := range listAiDataPlatformsResponse.Items {
		id := *aiDataPlatform.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AiDataPlatformId", id)

	}
	return resourceIds, nil
}

func AiDataPlatformAiDataPlatformSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if aiDataPlatformResponse, ok := response.Response.(oci_ai_data_platform.GetAiDataPlatformResponse); ok {
		return aiDataPlatformResponse.LifecycleState != oci_ai_data_platform.AiDataPlatformLifecycleStateDeleted
	}
	return false
}

func AiDataPlatformAiDataPlatformSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiDataPlatformClient().GetAiDataPlatform(context.Background(), oci_ai_data_platform.GetAiDataPlatformRequest{
		AiDataPlatformId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
