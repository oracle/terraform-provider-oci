// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiHostedApplicationStorageRequiredOnlyResource = GenerativeAiHostedApplicationStorageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_storage", "test_hosted_application_storage", acctest.Required, acctest.Create, GenerativeAiHostedApplicationStorageRepresentation)

	GenerativeAiHostedApplicationStorageResourceConfig = GenerativeAiHostedApplicationStorageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_storage", "test_hosted_application_storage", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationStorageRepresentation)

	GenerativeAiHostedApplicationStorageSingularDataSourceRepresentation = map[string]interface{}{
		"hosted_application_storage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_hosted_application_storage.test_hosted_application_storage.id}`},
	}

	GenerativeAiHostedApplicationStorageDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"hosted_application_storage_type": acctest.Representation{RepType: acctest.Optional, Create: `CACHE`},
		"id":                              acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_hosted_application_storage.test_hosted_application_storage.id}`},
		"state":                           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiHostedApplicationStorageDataSourceFilterRepresentation}}
	GenerativeAiHostedApplicationStorageDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_hosted_application_storage.test_hosted_application_storage.id}`}},
	}

	GenerativeAiHostedApplicationStorageRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`},
		"storage_type":   acctest.Representation{RepType: acctest.Required, Create: `CACHE`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
	}

	GenerativeAiHostedApplicationStorageResourceDependencies = ""
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiHostedApplicationStorageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiHostedApplicationStorageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_generative_ai_hosted_application_storage.test_hosted_application_storage"
	datasourceName := "data.oci_generative_ai_hosted_application_storages.test_hosted_application_storages"
	singularDatasourceName := "data.oci_generative_ai_hosted_application_storage.test_hosted_application_storage"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiHostedApplicationStorageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_storage", "test_hosted_application_storage", acctest.Optional, acctest.Create, GenerativeAiHostedApplicationStorageRepresentation), "generativeai", "hostedApplicationStorage", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiHostedApplicationStorageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationStorageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_storage", "test_hosted_application_storage", acctest.Required, acctest.Create, GenerativeAiHostedApplicationStorageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "storage_type", "CACHE"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationStorageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiHostedApplicationStorageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_storage", "test_hosted_application_storage", acctest.Optional, acctest.Create, GenerativeAiHostedApplicationStorageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_ids"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_type", "CACHE"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_application_storages", "test_hosted_application_storages", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationStorageDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedApplicationStorageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_hosted_application_storage", "test_hosted_application_storage", acctest.Optional, acctest.Update, GenerativeAiHostedApplicationStorageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "hosted_application_storage_type", "CACHE"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "hosted_application_storage_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "hosted_application_storage_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_hosted_application_storage", "test_hosted_application_storage", acctest.Required, acctest.Create, GenerativeAiHostedApplicationStorageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiHostedApplicationStorageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hosted_application_storage_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "application_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_type", "CACHE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiHostedApplicationStorageRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiHostedApplicationStorageDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_hosted_application_storage" {
			noResourceFound = false
			request := oci_generative_ai.GetHostedApplicationStorageRequest{}

			tmp := rs.Primary.ID
			request.HostedApplicationStorageId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetHostedApplicationStorage(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.HostedApplicationStorageLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiHostedApplicationStorage") {
		resource.AddTestSweepers("GenerativeAiHostedApplicationStorage", &resource.Sweeper{
			Name:         "GenerativeAiHostedApplicationStorage",
			Dependencies: acctest.DependencyGraph["hostedApplicationStorage"],
			F:            sweepGenerativeAiHostedApplicationStorageResource,
		})
	}
}

func sweepGenerativeAiHostedApplicationStorageResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	hostedApplicationStorageIds, err := getGenerativeAiHostedApplicationStorageIds(compartment)
	if err != nil {
		return err
	}
	for _, hostedApplicationStorageId := range hostedApplicationStorageIds {
		if ok := acctest.SweeperDefaultResourceId[hostedApplicationStorageId]; !ok {
			deleteHostedApplicationStorageRequest := oci_generative_ai.DeleteHostedApplicationStorageRequest{}

			deleteHostedApplicationStorageRequest.HostedApplicationStorageId = &hostedApplicationStorageId

			deleteHostedApplicationStorageRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteHostedApplicationStorage(context.Background(), deleteHostedApplicationStorageRequest)
			if error != nil {
				fmt.Printf("Error deleting HostedApplicationStorage %s %s, It is possible that the resource is already deleted. Please verify manually \n", hostedApplicationStorageId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &hostedApplicationStorageId, GenerativeAiHostedApplicationStorageSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiHostedApplicationStorageSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiHostedApplicationStorageIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "HostedApplicationStorageId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listHostedApplicationStoragesRequest := oci_generative_ai.ListHostedApplicationStoragesRequest{}
	listHostedApplicationStoragesRequest.CompartmentId = &compartmentId
	listHostedApplicationStoragesRequest.LifecycleState = oci_generative_ai.HostedApplicationStorageLifecycleStateActive
	listHostedApplicationStoragesResponse, err := generativeAiClient.ListHostedApplicationStorages(context.Background(), listHostedApplicationStoragesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting HostedApplicationStorage list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, hostedApplicationStorage := range listHostedApplicationStoragesResponse.Items {
		id := *hostedApplicationStorage.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "HostedApplicationStorageId", id)
	}
	return resourceIds, nil
}

func GenerativeAiHostedApplicationStorageSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if hostedApplicationStorageResponse, ok := response.Response.(oci_generative_ai.GetHostedApplicationStorageResponse); ok {
		return hostedApplicationStorageResponse.LifecycleState != oci_generative_ai.HostedApplicationStorageLifecycleStateDeleted
	}
	return false
}

func GenerativeAiHostedApplicationStorageSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetHostedApplicationStorage(context.Background(), oci_generative_ai.GetHostedApplicationStorageRequest{
		HostedApplicationStorageId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
