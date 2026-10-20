// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingMaskingIntegrationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingMaskingIntegrationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentID := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	targetID := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	databaseUserName := utils.GetEnvSettingWithBlankDefault("target_database_user_name")
	databasePassword := utils.GetEnvSettingWithBlankDefault("target_database_password")

	variables := fmt.Sprintf(`
variable "compartment_id" {
  default = %q
}

variable "target_id" {
  default = %q
}

variable "target_database_user_name" {
  default = %q
}

variable "target_database_password" {
  default = %q
}
`, compartmentID, targetID, databaseUserName, databasePassword)

	maskingPolicyRepresentation := map[string]interface{}{
		"column_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"column_source": acctest.Representation{RepType: acctest.Required, Create: `TARGET`},
			"target_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		}},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
	maskingColumnRepresentation := map[string]interface{}{
		"column_name":       acctest.Representation{RepType: acctest.Required, Create: `EMAIL`},
		"masking_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.subsetting_masking_policy.id}`},
		"object":            acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		"schema_name":       acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`},
		"masking_formats": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"format_entries": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
				"type":         acctest.Representation{RepType: acctest.Required, Create: `FIXED_STRING`},
				"fixed_string": acctest.Representation{RepType: acctest.Required, Create: `MASKED_EMAIL`},
			}},
		}},
	}
	subsettingPolicyRepresentation := map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"masking_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.subsetting_masking_policy.id}`},
		"schema_source": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"schema_source":          acctest.Representation{RepType: acctest.Required, Create: `TARGET`},
			"schemas_for_subsetting": acctest.Representation{RepType: acctest.Required, Create: []string{`HR_TEST`}},
			"target_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		}},
	}
	ruleRepresentation := map[string]interface{}{
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.subsetting_masking_policy.id}`},
		"scope": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"schema_name": acctest.Representation{RepType: acctest.Required, Create: `HR_TEST`},
			"scope_type":  acctest.Representation{RepType: acctest.Required, Create: `SPECIFIC`},
			"object":      acctest.Representation{RepType: acctest.Required, Create: `EMPLOYEES`},
		}},
		"subset_rule_entry": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"rule_type": acctest.Representation{RepType: acctest.Required, Create: `PERCENT`},
			"percent":   acctest.Representation{RepType: acctest.Required, Create: `10`},
		}},
	}
	subsetDataRepresentation := map[string]interface{}{
		"masking":              acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.subsetting_masking_policy.id}`},
		"target_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"target_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"user_name": acctest.Representation{RepType: acctest.Required, Create: `${var.target_database_user_name}`},
			"password":  acctest.Representation{RepType: acctest.Required, Create: `${var.target_database_password}`},
		}},
	}

	maskingPolicy := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "subsetting_masking_policy", acctest.Required, acctest.Create, maskingPolicyRepresentation)
	maskingColumn := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_masking_column", "subsetting_masking_column", acctest.Required, acctest.Create, maskingColumnRepresentation)
	maskingColumn = withIntegrationDependsOn(maskingColumn, "oci_data_safe_masking_policy.subsetting_masking_policy")
	subsettingPolicy := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy", "subsetting_masking_policy", acctest.Required, acctest.Create, subsettingPolicyRepresentation)
	subsettingPolicy = withIntegrationDependsOn(subsettingPolicy, "oci_data_safe_masking_policies_masking_column.subsetting_masking_column")
	rule := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subsetting_policy_subsetting_rule", "subsetting_masking_rule", acctest.Required, acctest.Create, ruleRepresentation)
	subsetData := acctest.GenerateResourceFromRepresentationMap("oci_data_safe_subset_data", "subsetting_masking_subset_data", acctest.Required, acctest.Create, subsetDataRepresentation)
	subsetData = withIntegrationDependsOn(subsetData, "oci_data_safe_subsetting_policy_subsetting_rule.subsetting_masking_rule")

	reports := acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_reports", "subsetting_masking_reports", acctest.Required, acctest.Create, map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subsetting_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_subsetting_policy.subsetting_masking_policy.id}`},
		"target_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	})
	reports = withIntegrationDependsOn(reports, "oci_data_safe_subset_data.subsetting_masking_subset_data")
	subsettingReport := acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_report", "subsetting_masking_report", acctest.Required, acctest.Create, map[string]interface{}{
		"subsetting_report_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_data_safe_subsetting_reports.subsetting_masking_reports.subsetting_report_collection.0.items.0.id}`},
	})
	subsettingReport = withIntegrationDependsOn(subsettingReport, "data.oci_data_safe_subsetting_reports.subsetting_masking_reports")

	resourceName := "oci_data_safe_subset_data.subsetting_masking_subset_data"
	subsettingReportName := "data.oci_data_safe_subsetting_report.subsetting_masking_report"

	acctest.ResourceTest(t, nil, []resource.TestStep{{
		Config: config + variables + maskingPolicy + maskingColumn + subsettingPolicy + rule + subsetData + reports + subsettingReport,
		Check: acctest.ComposeAggregateTestCheckFuncWrapper(
			resource.TestCheckResourceAttrSet(resourceName, "id"),
			resource.TestCheckResourceAttr(subsettingReportName, "subsetting_status", "SUCCESS"),
			resource.TestCheckResourceAttrSet(subsettingReportName, "masking_policy_id"),
			resource.TestCheckResourceAttrPair(subsettingReportName, "masking_policy_id", "oci_data_safe_masking_policy.subsetting_masking_policy", "id"),
			resource.TestCheckResourceAttrSet(subsettingReportName, "masking_status"),
			resource.TestCheckResourceAttrSet(subsettingReportName, "masking_work_request_id"),
			checkDataSafeSubsettingMaskingJob,
		),
	}})
}

func checkDataSafeSubsettingMaskingJob(s *terraform.State) error {
	subsettingReportName := "data.oci_data_safe_subsetting_report.subsetting_masking_report"
	maskingWorkRequestID, err := acctest.FromInstanceState(s, subsettingReportName, "masking_work_request_id")
	if err != nil {
		return err
	}

	dataSafeClient := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	var workRequestResponse oci_data_safe.GetWorkRequestResponse
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			workRequestResponse, err = dataSafeClient.GetWorkRequest(context.Background(), oci_data_safe.GetWorkRequestRequest{
				WorkRequestId: &maskingWorkRequestID,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: tfresource.GetRetryPolicy(false, "data_safe"),
				},
			})
			if err != nil {
				return nil, "", err
			}
			return workRequestResponse, string(workRequestResponse.Status), nil
		},
		Timeout: 30 * time.Minute,
	}
	if _, err = stateConf.WaitForStateContext(context.Background()); err != nil {
		return err
	}
	if workRequestResponse.Status != oci_data_safe.WorkRequestStatusSucceeded {
		return fmt.Errorf("masking work request %s did not succeed: %s", maskingWorkRequestID, workRequestResponse.Status)
	}

	var maskingReportID string
	for _, workRequestResource := range workRequestResponse.Resources {
		if workRequestResource.EntityType == nil || workRequestResource.Identifier == nil {
			continue
		}
		entityType := strings.ReplaceAll(strings.ToLower(*workRequestResource.EntityType), "_", "")
		if strings.Contains(entityType, "maskingreport") &&
			(workRequestResource.ActionType == oci_data_safe.WorkRequestResourceActionTypeCreated ||
				workRequestResource.ActionType == oci_data_safe.WorkRequestResourceActionTypeSucceeded) {
			maskingReportID = *workRequestResource.Identifier
			break
		}
	}
	if maskingReportID == "" {
		return fmt.Errorf("masking work request %s succeeded but did not return a masking report ID", maskingWorkRequestID)
	}

	maskingReportResponse, err := dataSafeClient.GetMaskingReport(context.Background(), oci_data_safe.GetMaskingReportRequest{
		MaskingReportId: &maskingReportID,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "data_safe"),
		},
	})
	if err != nil {
		return err
	}
	if maskingReportResponse.Id == nil || *maskingReportResponse.Id != maskingReportID {
		return fmt.Errorf("masking report lookup returned an unexpected report ID for masking work request %s", maskingWorkRequestID)
	}
	if maskingReportResponse.MaskingStatus != oci_data_safe.MaskingReportMaskingStatusSuccess {
		return fmt.Errorf("masking report %s did not succeed: %s", maskingReportID, maskingReportResponse.MaskingStatus)
	}

	subsettingReportID, err := acctest.FromInstanceState(s, subsettingReportName, "id")
	if err != nil {
		return err
	}
	subsettingReportResponse, err := dataSafeClient.GetSubsettingReport(context.Background(), oci_data_safe.GetSubsettingReportRequest{
		SubsettingReportId: &subsettingReportID,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "data_safe"),
		},
	})
	if err != nil {
		return err
	}
	if subsettingReportResponse.MaskingReportId == nil || *subsettingReportResponse.MaskingReportId == "" {
		return fmt.Errorf("subsetting report %s does not contain masking_report_id after masking work request %s succeeded", subsettingReportID, maskingWorkRequestID)
	}
	if *subsettingReportResponse.MaskingReportId != maskingReportID {
		return fmt.Errorf("subsetting report %s references masking report %s, expected %s", subsettingReportID, *subsettingReportResponse.MaskingReportId, maskingReportID)
	}

	return nil
}

func withIntegrationDependsOn(block string, dependencies ...string) string {
	return strings.TrimSuffix(block, "}\n") + fmt.Sprintf("depends_on = [%s]\n}\n", quoteTerraformStrings(dependencies...))
}

func quoteTerraformStrings(values ...string) string {
	quoted := make([]string, len(values))
	for i, value := range values {
		quoted[i] = strconv.Quote(value)
	}
	return strings.Join(quoted, ", ")
}
