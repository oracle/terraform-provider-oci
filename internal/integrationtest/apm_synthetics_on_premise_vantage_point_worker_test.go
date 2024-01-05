// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApmSyntheticsOnPremiseVantagePointWorkerRequiredOnlyResource = ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Required, acctest.Create, ApmSyntheticsOnPremiseVantagePointWorkerRepresentation)

	ApmSyntheticsOnPremiseVantagePointWorkerResourceConfig = ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointWorkerRepresentation)

	ApmSyntheticsOnPremiseVantagePointWorkerSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"on_premise_vantage_point_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.id}`},
		"worker_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_on_premise_vantage_point_worker.test_on_premise_vantage_point_worker.id}`},
	}

	ApmSyntheticsOnPremiseVantagePointWorkerDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"on_premise_vantage_point_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.id}`},
		"name":                        acctest.Representation{RepType: acctest.Optional, Create: `Test`},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmSyntheticsOnPremiseVantagePointWorkerDataSourceFilterRepresentation}}
	ApmSyntheticsOnPremiseVantagePointWorkerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_synthetics_on_premise_vantage_point_worker.test_on_premise_vantage_point_worker.name}`}},
	}

	ApmSyntheticsOnPremiseVantagePointWorkerRepresentation = map[string]interface{}{
		"apm_domain_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"name":                                acctest.Representation{RepType: acctest.Required, Create: `Test`},
		"on_premise_vantage_point_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.id}`},
		"resource_principal_token_public_key": acctest.Representation{RepType: acctest.Required, Create: `-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0GuJMVpuYE3l2pAe4mwrB73pExN42hG5OkbiJimrSbSHBffng6NacHz4hX8Ri9WhuQSq51nXsGjixnVvjnI3RzgenAaLVrf48a8RmS5D0pwrjshkf5Vt/hSXYL2lVUToGTUdOzXb5ZAH6BN9SE+LPEeBl6QnXn90teMXeVPnarg9WE1LMf8eNoD3PRaXEa9i3Q0Q2/3cfXVX1MhHk5wi/fUKsnbTjy67a49vvC3DKbakw76q4lrdtvp7M5EKN+paD0qg64wRKn8/bCYvI/tjM+LufvSLJJSj7KQs83t5xKBK60FVRUK/d3bRdilb8XnezBSGSdPDY9fL6yn0z8UORQIDAQAB\n-----END PUBLIC KEY-----`},
		"version":                             acctest.Representation{RepType: acctest.Required, Create: `1.2.4`},
		"worker_type":                         acctest.Representation{RepType: acctest.Required, Create: `DOCKER`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"priority":                            acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `9`},
		"status":                              acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
	}

	ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point", "test_on_premise_vantage_point", acctest.Required, acctest.Create, ApmSyntheticsOnPremiseVantagePointRepresentation)
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsOnPremiseVantagePointWorkerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsOnPremiseVantagePointWorkerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_synthetics_on_premise_vantage_point_worker.test_on_premise_vantage_point_worker"
	datasourceName := "data.oci_apm_synthetics_on_premise_vantage_point_workers.test_on_premise_vantage_point_workers"
	singularDatasourceName := "data.oci_apm_synthetics_on_premise_vantage_point_worker.test_on_premise_vantage_point_worker"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Optional, acctest.Create, ApmSyntheticsOnPremiseVantagePointWorkerRepresentation), "apmsynthetics", "onPremiseVantagePointWorker", t)

	acctest.ResourceTest(t, testAccCheckApmSyntheticsOnPremiseVantagePointWorkerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Required, acctest.Create, ApmSyntheticsOnPremiseVantagePointWorkerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "Test"),
				resource.TestCheckResourceAttrSet(resourceName, "on_premise_vantage_point_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_principal_token_public_key"),
				resource.TestCheckResourceAttr(resourceName, "version", "1.2.4"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Optional, acctest.Create, ApmSyntheticsOnPremiseVantagePointWorkerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "Test"),
				resource.TestCheckResourceAttrSet(resourceName, "on_premise_vantage_point_id"),
				resource.TestCheckResourceAttrSet(resourceName, "opvp_id"),
				resource.TestCheckResourceAttrSet(resourceName, "opvp_name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_principal_token_public_key"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "version", "1.2.4"),
				resource.TestCheckResourceAttr(resourceName, "version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "worker_type", "DOCKER"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointWorkerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "Test"),
				resource.TestCheckResourceAttrSet(resourceName, "on_premise_vantage_point_id"),
				resource.TestCheckResourceAttrSet(resourceName, "opvp_id"),
				resource.TestCheckResourceAttrSet(resourceName, "opvp_name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "9"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_principal_token_public_key"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_id"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "version", "1.2.4"),
				resource.TestCheckResourceAttr(resourceName, "version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "worker_type", "DOCKER"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_workers", "test_on_premise_vantage_point_workers", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointWorkerDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointWorkerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Optional, acctest.Update, ApmSyntheticsOnPremiseVantagePointWorkerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "Test"),
				resource.TestCheckResourceAttrSet(datasourceName, "on_premise_vantage_point_id"),

				resource.TestCheckResourceAttr(datasourceName, "worker_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "worker_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_on_premise_vantage_point_worker", "test_on_premise_vantage_point_worker", acctest.Required, acctest.Create, ApmSyntheticsOnPremiseVantagePointWorkerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsOnPremiseVantagePointWorkerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "on_premise_vantage_point_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "worker_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "geo_info"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identity_info.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "monitor_list.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "Test"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opvp_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opvp_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "priority", "9"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runtime_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_sync_up"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "worker_type", "DOCKER"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmSyntheticsOnPremiseVantagePointWorkerRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
				"resource_principal_token_public_key",
				"version",
				"configuration_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApmSyntheticsOnPremiseVantagePointWorkerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ApmSyntheticClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_synthetics_on_premise_vantage_point_worker" {
			noResourceFound = false
			request := oci_apm_synthetics.GetWorkerRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			if value, ok := rs.Primary.Attributes["on_premise_vantage_point_id"]; ok {
				request.OnPremiseVantagePointId = &value
			}

			tmp := rs.Primary.ID
			request.WorkerId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")

			_, err := client.GetWorker(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("ApmSyntheticsOnPremiseVantagePointWorker") {
		resource.AddTestSweepers("ApmSyntheticsOnPremiseVantagePointWorker", &resource.Sweeper{
			Name:         "ApmSyntheticsOnPremiseVantagePointWorker",
			Dependencies: acctest.DependencyGraph["onPremiseVantagePointWorker"],
			F:            sweepApmSyntheticsOnPremiseVantagePointWorkerResource,
		})
	}
}

func sweepApmSyntheticsOnPremiseVantagePointWorkerResource(compartment string) error {
	apmSyntheticClient := acctest.GetTestClients(&schema.ResourceData{}).ApmSyntheticClient()
	onPremiseVantagePointWorkerIds, err := getApmSyntheticsOnPremiseVantagePointWorkerIds(compartment)
	if err != nil {
		return err
	}
	for _, onPremiseVantagePointWorkerId := range onPremiseVantagePointWorkerIds {
		if ok := acctest.SweeperDefaultResourceId[onPremiseVantagePointWorkerId]; !ok {
			deleteWorkerRequest := oci_apm_synthetics.DeleteWorkerRequest{}

			deleteWorkerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_synthetics")
			_, error := apmSyntheticClient.DeleteWorker(context.Background(), deleteWorkerRequest)
			if error != nil {
				fmt.Printf("Error deleting OnPremiseVantagePointWorker %s %s, It is possible that the resource is already deleted. Please verify manually \n", onPremiseVantagePointWorkerId, error)
				continue
			}
		}
	}
	return nil
}

func getApmSyntheticsOnPremiseVantagePointWorkerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OnPremiseVantagePointWorkerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apmSyntheticClient := acctest.GetTestClients(&schema.ResourceData{}).ApmSyntheticClient()

	listWorkersRequest := oci_apm_synthetics.ListWorkersRequest{}
	//listWorkersRequest.CompartmentId = &compartmentId

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for OnPremiseVantagePointWorker resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listWorkersRequest.ApmDomainId = &apmDomainId

		onPremiseVantagePointIds, error := getApmSyntheticsOnPremiseVantagePointIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting onPremiseVantagePointId required for OnPremiseVantagePointWorker resource requests \n")
		}
		for _, onPremiseVantagePointId := range onPremiseVantagePointIds {
			listWorkersRequest.OnPremiseVantagePointId = &onPremiseVantagePointId

			listWorkersResponse, err := apmSyntheticClient.ListWorkers(context.Background(), listWorkersRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting OnPremiseVantagePointWorker list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, onPremiseVantagePointWorker := range listWorkersResponse.Items {
				id := *onPremiseVantagePointWorker.Id
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OnPremiseVantagePointWorkerId", id)
			}

		}
	}
	return resourceIds, nil
}
