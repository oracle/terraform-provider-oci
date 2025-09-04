// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafemaskingReportSingularDataSourceRepresentation = map[string]interface{}{
		"masking_report_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.masking_report_id}`},
		"target_database_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.target_database_group_id}`},
	}

	DataSafemaskingReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingReportResource_basic(t *testing.T) {
	//t.Skip("Skipping this test as the report ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeMaskingReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	maskingReportId := utils.GetEnvSettingWithBlankDefault("data_safe_masking_report_id")
	maskingReportIdVariableStr := fmt.Sprintf("variable \"masking_report_id\" { default = \"%s\" }\n", maskingReportId)

	targetGroupId := utils.GetEnvSettingWithBlankDefault("data_safe_target_group_ocid")
	targetGroupIdVariableStr := fmt.Sprintf("variable \"target_database_group_id\" { default = \"%s\" }\n", targetGroupId)

	datasourceName := "data.oci_data_safe_masking_reports.test_masking_reports"
	singularDatasourceName := "data.oci_data_safe_masking_report.test_masking_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + maskingReportIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_reports", "test_masking_reports", acctest.Optional, acctest.Create, DataSafemaskingReportDataSourceRepresentation) +
				compartmentIdVariableStr + targetGroupIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "masking_report_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + maskingReportIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_report", "test_masking_report", acctest.Required, acctest.Create, DataSafemaskingReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_report_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_drop_temp_tables_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_redo_logging_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_refresh_stats_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_work_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parallel_degree"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recompile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_masking_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_masking_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_columns"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_objects"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_schemas"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_sensitive_types"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_values"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_post_masking_script_errors"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_pre_masking_script_errors"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataSafeMaskingReport") {
		resource.AddTestSweepers("DataSafeMaskingReport", &resource.Sweeper{
			Name:         "DataSafeMaskingReport",
			Dependencies: acctest.DependencyGraph["maskingReport"],
			F:            sweepDataSafeMaskingReportResource,
		})
	}
}

func sweepDataSafeMaskingReportResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	maskingReportIds, err := getDataSafeMaskingReportIds(compartment)
	if err != nil {
		return err
	}
	for _, maskingReportId := range maskingReportIds {
		if ok := acctest.SweeperDefaultResourceId[maskingReportId]; !ok {
			deleteMaskingReportRequest := oci_data_safe.DeleteMaskingReportRequest{}

			deleteMaskingReportRequest.MaskingReportId = &maskingReportId

			deleteMaskingReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteMaskingReport(context.Background(), deleteMaskingReportRequest)
			if error != nil {
				fmt.Printf("Error deleting MaskingReport %s %s, It is possible that the resource is already deleted. Please verify manually \n", maskingReportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &maskingReportId, DataSafeMaskingReportSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeMaskingReportSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeMaskingReportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MaskingReportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listMaskingReportsRequest := oci_data_safe.ListMaskingReportsRequest{}
	listMaskingReportsRequest.CompartmentId = &compartmentId
	listMaskingReportsResponse, err := dataSafeClient.ListMaskingReports(context.Background(), listMaskingReportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MaskingReport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, maskingReport := range listMaskingReportsResponse.Items {
		id := *maskingReport.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MaskingReportId", id)
	}
	return resourceIds, nil
}

func DataSafeMaskingReportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if maskingReportResponse, ok := response.Response.(oci_data_safe.GetMaskingReportResponse); ok {
		return maskingReportResponse.LifecycleState != oci_data_safe.MaskingLifecycleStateDeleted
	}
	return false
}

func DataSafeMaskingReportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetMaskingReport(context.Background(), oci_data_safe.GetMaskingReportRequest{
		MaskingReportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
