// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strings"
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
	DataSafeSubsettingReportSingularDataSourceRepresentation = map[string]interface{}{
		"subsetting_report_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_data_safe_subsetting_reports.test_subsetting_reports.subsetting_report_collection.0.items.0.id}`},
	}

	DataSafeSubsettingReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"subsetting_policy_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
	}

	DataSafeSubsettingReportResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "test_subsetting_policy", acctest.Required, acctest.Create, DataSafeSubsettingPolicyRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	ruleResource := acctest.GenerateResourceFromRepresentationMap(
		"oci_data_safe_subsetting_policy_subsetting_rule",
		"test_subsetting_policy_subsetting_rule",
		acctest.Required,
		acctest.Create,
		DataSafeSubsettingPolicySubsettingRuleRepresentation,
	)
	subsetResource := acctest.GenerateResourceFromRepresentationMap(
		"oci_data_safe_subset_data",
		"test_subset_data",
		acctest.Required,
		acctest.Create,
		dataSafeSubsetDataRepresentation,
	)
	subsetResource = strings.TrimSuffix(subsetResource, "}\n") + "depends_on = [\"oci_data_safe_subsetting_policy_subsetting_rule.test_subsetting_policy_subsetting_rule\"]\n}\n"
	jobDependencies := ruleResource + subsetResource
	reportsDataSource := acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_reports", "test_subsetting_reports", acctest.Required, acctest.Create, DataSafeSubsettingReportDataSourceRepresentation)
	reportsDataSource = strings.TrimSuffix(reportsDataSource, "}\n") + "depends_on = [\"oci_data_safe_subset_data.test_subset_data\"]\n}\n"
	reportDataSource := acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_report", "test_subsetting_report", acctest.Required, acctest.Create, DataSafeSubsettingReportSingularDataSourceRepresentation)
	reportDataSource = strings.TrimSuffix(reportDataSource, "}\n") + "depends_on = [\"oci_data_safe_subset_data.test_subset_data\"]\n}\n"

	datasourceName := "data.oci_data_safe_subsetting_reports.test_subsetting_reports"
	singularDatasourceName := "data.oci_data_safe_subsetting_report.test_subsetting_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				reportsDataSource + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingReportResourceConfig + jobDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_report_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_report_collection.0.items.0.subsetting_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_report_collection.0.items.0.target_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				reportsDataSource + reportDataSource + compartmentIdVariableStr + targetIdVariableStr + DataSafeSubsettingReportResourceConfig + jobDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_report_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_size_after_subsetting_in_kbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_size_before_subsetting_in_kbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_redo_logging_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_refresh_stats_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parallel_degree"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recompile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subsetting_work_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_subsetting_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_subsetting_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_post_subsetting_script_errors"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_pre_subsetting_script_errors"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_subsetted_objects"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_subsetted_rows"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_subsetted_schemas"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataSafeSubsettingReport") {
		resource.AddTestSweepers("DataSafeSubsettingReport", &resource.Sweeper{
			Name:         "DataSafeSubsettingReport",
			Dependencies: acctest.DependencyGraph["subsettingReport"],
			F:            sweepDataSafeSubsettingReportResource,
		})
	}
}

func sweepDataSafeSubsettingReportResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	subsettingReportIds, err := getDataSafeSubsettingReportIds(compartment)
	if err != nil {
		return err
	}
	for _, subsettingReportId := range subsettingReportIds {
		if ok := acctest.SweeperDefaultResourceId[subsettingReportId]; !ok {
			deleteSubsettingReportRequest := oci_data_safe.DeleteSubsettingReportRequest{}

			deleteSubsettingReportRequest.SubsettingReportId = &subsettingReportId

			deleteSubsettingReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSubsettingReport(context.Background(), deleteSubsettingReportRequest)
			if error != nil {
				fmt.Printf("Error deleting SubsettingReport %s %s, It is possible that the resource is already deleted. Please verify manually \n", subsettingReportId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subsettingReportId, DataSafeSubsettingReportSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSubsettingReportSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSubsettingReportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubsettingReportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSubsettingReportsRequest := oci_data_safe.ListSubsettingReportsRequest{}
	listSubsettingReportsRequest.CompartmentId = &compartmentId
	listSubsettingReportsResponse, err := dataSafeClient.ListSubsettingReports(context.Background(), listSubsettingReportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SubsettingReport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subsettingReport := range listSubsettingReportsResponse.Items {
		id := *subsettingReport.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubsettingReportId", id)
	}
	return resourceIds, nil
}

func DataSafeSubsettingReportSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subsettingReportResponse, ok := response.Response.(oci_data_safe.GetSubsettingReportResponse); ok {
		return subsettingReportResponse.LifecycleState != oci_data_safe.SubsettingReportLifecycleStateDeleted
	}
	return false
}

func DataSafeSubsettingReportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSubsettingReport(context.Background(), oci_data_safe.GetSubsettingReportRequest{
		SubsettingReportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
