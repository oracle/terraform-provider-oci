// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeMaskingPolicyHealthReportSingularDataSourceRepresentation = map[string]interface{}{
		"masking_policy_health_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.masking_policy_health_report_id}`},
	}
	DataSafeMaskingPolicyHealthReportsDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycleState":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"masking_policy_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	DataSafeMaskingPolicyHealthReportResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingPolicyHealthReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingPolicyHealthReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	healthReportId := utils.GetEnvSettingWithBlankDefault("masking_health_report_id")
	healthReportIdVariableStr := fmt.Sprintf("variable \"masking_policy_health_report_id\" { default = \"%s\" }\n", healthReportId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	//compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	//compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	datasourceName := "data.oci_data_safe_masking_policy_health_reports.test_masking_policy_health_reports"
	singularDatasourceName := "data.oci_data_safe_masking_policy_health_report.test_masking_policy_health_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policy_health_reports", "test_masking_policy_health_reports", acctest.Required, acctest.Create, DataSafeMaskingPolicyHealthReportsDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
			),
		},
		// verify singular datasource
		{
			Config: config + healthReportIdVariableStr + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policy_health_report", "test_masking_policy_health_report", acctest.Required, acctest.Create, DataSafeMaskingPolicyHealthReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_policy_health_report_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataSafeMaskingPolicyHealthReport") {
		resource.AddTestSweepers("DataSafeMaskingPolicyHealthReport", &resource.Sweeper{
			Name:         "DataSafeMaskingPolicyHealthReport",
			Dependencies: acctest.DependencyGraph["maskingPolicyHealthReport"],
			F:            sweepDataSafeMaskingPolicyHealthReportResource,
		})
	}
}

func sweepDataSafeMaskingPolicyHealthReportResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	maskingPolicyHealthReportIds, err := getDataSafeMaskingPolicyHealthReportIds(compartment)
	if err != nil {
		return err
	}
	for _, maskingPolicyHealthReportId := range maskingPolicyHealthReportIds {
		if ok := acctest.SweeperDefaultResourceId[maskingPolicyHealthReportId]; !ok {
			deleteMaskingPolicyHealthReportRequest := oci_data_safe.DeleteMaskingPolicyHealthReportRequest{}

			deleteMaskingPolicyHealthReportRequest.MaskingPolicyHealthReportId = &maskingPolicyHealthReportId

			deleteMaskingPolicyHealthReportRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteMaskingPolicyHealthReport(context.Background(), deleteMaskingPolicyHealthReportRequest)
			if error != nil {
				fmt.Printf("Error deleting MaskingPolicyHealthReport %s %s, It is possible that the resource is already deleted. Please verify manually \n", maskingPolicyHealthReportId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeMaskingPolicyHealthReportIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MaskingPolicyHealthReportId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listMaskingPolicyHealthReportsRequest := oci_data_safe.ListMaskingPolicyHealthReportsRequest{}
	listMaskingPolicyHealthReportsRequest.CompartmentId = &compartmentId
	listMaskingPolicyHealthReportsRequest.LifecycleState = oci_data_safe.ListMaskingPolicyHealthReportsLifecycleStateActive
	listMaskingPolicyHealthReportsResponse, err := dataSafeClient.ListMaskingPolicyHealthReports(context.Background(), listMaskingPolicyHealthReportsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MaskingPolicyHealthReport list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, maskingPolicyHealthReport := range listMaskingPolicyHealthReportsResponse.Items {
		id := *maskingPolicyHealthReport.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MaskingPolicyHealthReportId", id)
	}
	return resourceIds, nil
}

func DataSafeMaskingPolicyHealthReportSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetMaskingPolicyHealthReport(context.Background(), oci_data_safe.GetMaskingPolicyHealthReportRequest{
		MaskingPolicyHealthReportId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
