// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSubsettingPolicyHealthReportSingularDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_health_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subsetting_policy_health_report_id}`},
	}

	DataSafeSubsettingPolicyHealthReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                       acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                              acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"subsetting_policy_health_report_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.subsetting_policy_health_report_id}`},
		"subsetting_policy_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.subsetting_policy_id}`},
		"target_id":                          acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicyHealthReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicyHealthReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	healthReportId := utils.GetEnvSettingWithBlankDefault("subsetting_health_report_id")
	healthReportIdVariableStr := fmt.Sprintf("variable \"subsetting_policy_health_report_id\" { default = \"%s\" }\n", healthReportId)
	policyId := utils.GetEnvSettingWithBlankDefault("subsetting_policy_id")
	policyIdVariableStr := fmt.Sprintf("variable \"subsetting_policy_id\" { default = \"%s\" }\n", policyId)

	datasourceName := "data.oci_data_safe_subsetting_policy_health_reports.test_subsetting_policy_health_reports"
	singularDatasourceName := "data.oci_data_safe_subsetting_policy_health_report.test_subsetting_policy_health_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + healthReportIdVariableStr + policyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_health_reports", "test_subsetting_policy_health_reports", acctest.Required, acctest.Create, DataSafeSubsettingPolicyHealthReportDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_health_report_collection.0.items.0.subsetting_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_health_report_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_health_report_collection.0.items.0.error_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_health_report_collection.0.items.0.warning_count"),
			),
		},
		// verify singular datasource
		{
			Config: config + healthReportIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_health_report", "test_subsetting_policy_health_report", acctest.Required, acctest.Create, DataSafeSubsettingPolicyHealthReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_policy_health_report_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "error_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "warning_count"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataSafeSubsettingPolicyHealthReport") {
		resource.AddTestSweepers("DataSafeSubsettingPolicyHealthReport", &resource.Sweeper{
			Name:         "DataSafeSubsettingPolicyHealthReport",
			Dependencies: acctest.DependencyGraph["subsettingPolicyHealthReport"],
			F:            sweepDataSafeSubsettingPolicyHealthReportResource,
		})
	}
}

func sweepDataSafeSubsettingPolicyHealthReportResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	subsettingPolicyHealthReportIds, err := getDataSafeSubsettingPolicyHealthReportIds(compartment)
	if err != nil {
		return err
	}
	for _, subsettingPolicyHealthReportId := range subsettingPolicyHealthReportIds {
		if ok := acctest.SweeperDefaultResourceId[subsettingPolicyHealthReportId]; !ok {
			deleteSubsettingPolicyHealthReportRequest := oci_data_safe.DeleteSubsettingPolicyHealthReportRequest{}

			deleteSubsettingPolicyHealthReportRequest.SubsettingPolicyHealthReportId = &subsettingPolicyHealthReportId

			deleteSubsettingPolicyHealthReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSubsettingPolicyHealthReport(context.Background(), deleteSubsettingPolicyHealthReportRequest)
			if error != nil {
				fmt.Printf("Error deleting SubsettingPolicyHealthReport %s %s, It is possible that the resource is already deleted. Please verify manually \n", subsettingPolicyHealthReportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subsettingPolicyHealthReportId, DataSafeSubsettingPolicyHealthReportSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSubsettingPolicyHealthReportSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSubsettingPolicyHealthReportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubsettingPolicyHealthReportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSubsettingPolicyHealthReportsRequest := oci_data_safe.ListSubsettingPolicyHealthReportsRequest{}
	listSubsettingPolicyHealthReportsRequest.CompartmentId = &compartmentId
	listSubsettingPolicyHealthReportsRequest.LifecycleState = oci_data_safe.ListSubsettingPolicyHealthReportsLifecycleStateActive
	listSubsettingPolicyHealthReportsResponse, err := dataSafeClient.ListSubsettingPolicyHealthReports(context.Background(), listSubsettingPolicyHealthReportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SubsettingPolicyHealthReport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subsettingPolicyHealthReport := range listSubsettingPolicyHealthReportsResponse.Items {
		id := *subsettingPolicyHealthReport.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubsettingPolicyHealthReportId", id)
	}
	return resourceIds, nil
}

func DataSafeSubsettingPolicyHealthReportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subsettingPolicyHealthReportResponse, ok := response.Response.(oci_data_safe.GetSubsettingPolicyHealthReportResponse); ok {
		return subsettingPolicyHealthReportResponse.LifecycleState != oci_data_safe.SubsettingPolicyHealthReportLifecycleStateDeleting
	}
	return false
}

func DataSafeSubsettingPolicyHealthReportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSubsettingPolicyHealthReport(context.Background(), oci_data_safe.GetSubsettingPolicyHealthReportRequest{
		SubsettingPolicyHealthReportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
