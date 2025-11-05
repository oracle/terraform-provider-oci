// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PsaPsaWorkRequestSingularDataSourceRepresentation = map[string]interface{}{
		"work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_psa_psa_work_requests.test_psa_work_requests.work_request_summary_collection.0.items.0.id}`},
	}

	PsaPsaWorkRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"resource_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_psa_private_service_access.test_private_service_access.id}`},
	}
	PsaPsaWorkRequestResourcConfig = PsaPrivateServiceAccesResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psa_private_service_access", "test_private_service_access", acctest.Required, acctest.Create, PsaPrivateServiceAccesRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_work_requests", "test_psa_work_requests", acctest.Required, acctest.Create, PsaPsaWorkRequestDataSourceRepresentation)
)

// issue-routing-tag: psa/default
func TestPsaPsaWorkRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsaPsaWorkRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_psa_psa_work_requests.test_psa_work_requests"
	singularDatasourceName := "data.oci_psa_psa_work_request.test_psa_work_request"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + PsaPsaWorkRequestResourcConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_summary_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_summary_collection.0.items.#"),
				resource.TestCheckResourceAttr(datasourceName, "work_request_summary_collection.0.items.0.status", "SUCCEEDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_summary_collection.0.items.0.id"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + PsaPsaWorkRequestResourcConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_work_request", "test_psa_work_request", acctest.Required, acctest.Create, PsaPsaWorkRequestSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "percent_complete"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("PsaPsaWorkRequest") {
		resource.AddTestSweepers("PsaPsaWorkRequest", &resource.Sweeper{
			Name:         "PsaPsaWorkRequest",
			Dependencies: acctest.DependencyGraph["psaWorkRequest"],
			F:            sweepPsaPsaWorkRequestResource,
		})
	}
}

func sweepPsaPsaWorkRequestResource(compartment string) error {
	privateServiceAccessClient := acctest.GetTestClients(&schema.ResourceData{}).PrivateServiceAccessClient()
	psaWorkRequestIds, err := getPsaPsaWorkRequestIds(compartment)
	if err != nil {
		return err
	}
	for _, psaWorkRequestId := range psaWorkRequestIds {
		if ok := acctest.SweeperDefaultResourceId[psaWorkRequestId]; !ok {
			cancelPsaWorkRequestRequest := oci_psa.CancelPsaWorkRequestRequest{}

			cancelPsaWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psa")
			_, error := privateServiceAccessClient.CancelPsaWorkRequest(context.Background(), cancelPsaWorkRequestRequest)
			if error != nil {
				fmt.Printf("Error deleting PsaWorkRequest %s %s, It is possible that the resource is already deleted. Please verify manually \n", psaWorkRequestId, error)
				continue
			}
		}
	}
	return nil
}

func getPsaPsaWorkRequestIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PsaWorkRequestId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	privateServiceAccessClient := acctest.GetTestClients(&schema.ResourceData{}).PrivateServiceAccessClient()

	listPsaWorkRequestsRequest := oci_psa.ListPsaWorkRequestsRequest{}
	listPsaWorkRequestsRequest.CompartmentId = &compartmentId
	listPsaWorkRequestsResponse, err := privateServiceAccessClient.ListPsaWorkRequests(context.Background(), listPsaWorkRequestsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PsaWorkRequest list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, psaWorkRequest := range listPsaWorkRequestsResponse.Items {
		id := *psaWorkRequest.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PsaWorkRequestId", id)
	}
	return resourceIds, nil
}
