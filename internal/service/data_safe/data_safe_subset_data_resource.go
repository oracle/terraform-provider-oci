// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License 2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// DataSafeSubsetDataResource starts a subsetting operation. The operation is
// asynchronous and the resource is intentionally non-deletable.
func DataSafeSubsetDataResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			policyID := d.Get("subsetting_policy_id").(string)
			credentials := d.Get("target_credentials").([]interface{})
			if len(credentials) == 0 {
				return diag.Errorf("target_credentials must be specified")
			}
			values := credentials[0].(map[string]interface{})
			userName := values["user_name"].(string)
			password := values["password"].(string)
			request := oci_data_safe.SubsetDataRequest{
				SubsettingPolicyId: &policyID,
				SubsetDataDetails: oci_data_safe.SubsetDataDetails{
					TargetCredentials: &oci_data_safe.Credentials{
						UserName: &userName,
						Password: &password,
					},
				},
			}
			if masking, ok := d.GetOkExists("masking"); ok {
				request.SubsetDataDetails.Masking = oci_data_safe.SubsetDataDetailsMaskingEnum(masking.(string))
			}
			if targetID, ok := d.GetOk("target_id"); ok {
				targetIDValue := targetID.(string)
				request.SubsetDataDetails.TargetId = &targetIDValue
			}
			if isRerun, ok := d.GetOkExists("is_rerun"); ok {
				isRerunValue := isRerun.(bool)
				request.SubsetDataDetails.IsRerun = &isRerunValue
			}
			if reRunFromStep, ok := d.GetOk("re_run_from_step"); ok {
				request.SubsetDataDetails.ReRunFromStep = oci_data_safe.SubsetDataDetailsReRunFromStepEnum(reRunFromStep.(string))
			}
			if tablespace, ok := d.GetOk("tablespace"); ok {
				tablespaceValue := tablespace.(string)
				request.SubsetDataDetails.Tablespace = &tablespaceValue
			}
			if isRedoLoggingEnabled, ok := d.GetOkExists("is_redo_logging_enabled"); ok {
				isRedoLoggingEnabledValue := isRedoLoggingEnabled.(bool)
				request.SubsetDataDetails.IsRedoLoggingEnabled = &isRedoLoggingEnabledValue
			}
			if isRefreshStatsEnabled, ok := d.GetOkExists("is_refresh_stats_enabled"); ok {
				isRefreshStatsEnabledValue := isRefreshStatsEnabled.(bool)
				request.SubsetDataDetails.IsRefreshStatsEnabled = &isRefreshStatsEnabledValue
			}
			if parallelDegree, ok := d.GetOk("parallel_degree"); ok {
				parallelDegreeValue := parallelDegree.(string)
				request.SubsetDataDetails.ParallelDegree = &parallelDegreeValue
			}
			if recompile, ok := d.GetOk("recompile"); ok {
				request.SubsetDataDetails.Recompile = oci_data_safe.SubsettingPolicyRecompileEnum(recompile.(string))
			}
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")
			dataSafeClient := m.(*client.OracleClients).DataSafeClient()
			response, err := dataSafeClient.SubsetData(ctx, request)
			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
			if response.OpcWorkRequestId == nil {
				return diag.Errorf("subset operation did not return a work request ID")
			}
			if _, err = subsettingPolicyWaitForWorkRequest(ctx, response.OpcWorkRequestId, "subsettingReport", oci_data_safe.WorkRequestResourceActionTypeCreated, d.Timeout(schema.TimeoutCreate), false, dataSafeClient); err != nil {
				return tfresource.HandleDiagError(m, err)
			}
			d.SetId(policyID)
			return nil
		},
		ReadContext:   func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },
		DeleteContext: func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },
		Schema: map[string]*schema.Schema{
			"subsetting_policy_id": {Type: schema.TypeString, Required: true, ForceNew: true},
			"target_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"masking": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"ENABLED", "DISABLED"}, true),
			},
			"is_rerun": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"re_run_from_step": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PRE_SUBSETTING_SCRIPT", "POST_SUBSETTING_SCRIPT"}, true),
			},
			"tablespace": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 30),
			},
			"is_redo_logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"is_refresh_stats_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"parallel_degree": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateSubsettingParallelDegree,
			},
			"recompile": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"SERIAL", "PARALLEL", "NONE"}, true),
			},
			"target_credentials": {
				Type: schema.TypeList, Required: true, ForceNew: true, MinItems: 1, MaxItems: 1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"user_name": {Type: schema.TypeString, Required: true},
					"password":  {Type: schema.TypeString, Required: true, Sensitive: true},
				}},
			},
		},
	}
}
