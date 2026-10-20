// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingReportManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeSubsettingReportManagement,
		ReadContext:   readDataSafeSubsettingReportManagement,
		DeleteContext: deleteDataSafeSubsettingReportManagement,
		Schema: map[string]*schema.Schema{
			"target_credentials": {
				Type: schema.TypeList, Required: true, ForceNew: true, MinItems: 1, MaxItems: 1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"user_name": {Type: schema.TypeString, Required: true},
					"password":  {Type: schema.TypeString, Required: true, Sensitive: true},
				}},
			},
			"target_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Computed: true,
				ForceNew: true,
			},
			"is_refresh_stats_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"parallel_degree": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validateSubsettingParallelDegree,
			},
			"recompile": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"SERIAL", "PARALLEL", "NONE"}, true),
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subsetting_work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"masking_report_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"masking_work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subsetting_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_subsetting_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_subsetting_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_subsetted_objects": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_subsetted_schemas": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_subsetted_rows": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_size_before_subsetting_in_kbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_size_after_subsetting_in_kbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_pre_subsetting_script_errors": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_post_subsetting_script_errors": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSubsettingReportManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingReportManagementResourceCurd{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.getSubsettingReportWorkReq(ctx)
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	err1 := sync.GetWithContext(ctx)
	if err1 != nil {
		return tfresource.HandleDiagError(m, err1)
	}

	err = sync.SetData()
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func readDataSafeSubsettingReportManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingReportManagementResourceCurd{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteDataSafeSubsettingReportManagement(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingReportManagementResourceCurd{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeSubsettingReportManagementResourceCurd struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SubsettingReport
	DisableNotFoundRetries bool
}

func (s *DataSafeSubsettingReportManagementResourceCurd) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSubsettingReportManagementResourceCurd) getSubsettingReportWorkReq(ctx context.Context) error {
	policyID := s.D.Get("subsetting_policy_id").(string)
	targetID, targetIDConfigured := s.D.GetOk("target_id")
	if !targetIDConfigured {
		_, resolvedTargetID, err := getSubsettingPolicyTargetAndCompartment(ctx, s.Client, policyID)
		if err != nil {
			return err
		}
		if resolvedTargetID == nil {
			return fmt.Errorf("target_id must be specified when it cannot be resolved from subsetting policy %s", policyID)
		}
		targetID = *resolvedTargetID
		if err := s.D.Set("target_id", targetID); err != nil {
			return err
		}
	}

	// Subsetting report will be in the same compartment as the target.
	getTargetDatabaseRequest := oci_data_safe.GetTargetDatabaseRequest{}
	targetIDValue := targetID.(string)
	getTargetDatabaseRequest.TargetDatabaseId = &targetIDValue

	getTargetDatabaseResponse, err := s.Client.GetTargetDatabase(ctx, getTargetDatabaseRequest)
	if err != nil {
		return err
	}

	compartmentID := getTargetDatabaseResponse.CompartmentId

	err = s.D.Set("compartment_id", compartmentID)
	if err != nil {
		return err
	}

	// List all subsetting reports for given target and subsetting policy ID
	err = s.GetSubsettingReportList(ctx)
	if err != nil {
		return err
	}

	// check if subsetting report id is set and subsetting report already exists
	if s.D.Id() != "" {
		return nil
	}

	// Mask target to generate subsetting report
	maskTargetDatabaseRequest := oci_data_safe.SubsetDataRequest{}
	maskTargetDatabaseRequest.SubsettingPolicyId = &policyID

	if targetID, ok := s.D.GetOk("target_id"); ok {
		tmp := targetID.(string)
		maskTargetDatabaseRequest.TargetId = &tmp
	}

	credentials := s.D.Get("target_credentials").([]interface{})[0].(map[string]interface{})
	userName := credentials["user_name"].(string)
	password := credentials["password"].(string)
	maskTargetDatabaseRequest.TargetCredentials = &oci_data_safe.Credentials{UserName: &userName, Password: &password}

	if masking, ok := s.D.GetOk("masking"); ok {
		maskTargetDatabaseRequest.Masking = oci_data_safe.SubsetDataDetailsMaskingEnum(masking.(string))
	}
	if isRerun, ok := s.D.GetOkExists("is_rerun"); ok {
		value := isRerun.(bool)
		maskTargetDatabaseRequest.IsRerun = &value
	}
	if reRunFromStep, ok := s.D.GetOk("re_run_from_step"); ok {
		maskTargetDatabaseRequest.ReRunFromStep = oci_data_safe.SubsetDataDetailsReRunFromStepEnum(reRunFromStep.(string))
	}
	if tablespace, ok := s.D.GetOk("tablespace"); ok {
		value := tablespace.(string)
		maskTargetDatabaseRequest.Tablespace = &value
	}
	if isRedoLoggingEnabled, ok := s.D.GetOkExists("is_redo_logging_enabled"); ok {
		value := isRedoLoggingEnabled.(bool)
		maskTargetDatabaseRequest.IsRedoLoggingEnabled = &value
	}
	if isRefreshStatsEnabled, ok := s.D.GetOkExists("is_refresh_stats_enabled"); ok {
		value := isRefreshStatsEnabled.(bool)
		maskTargetDatabaseRequest.IsRefreshStatsEnabled = &value
	}
	if parallelDegree, ok := s.D.GetOk("parallel_degree"); ok {
		value := parallelDegree.(string)
		maskTargetDatabaseRequest.ParallelDegree = &value
	}
	if recompile, ok := s.D.GetOk("recompile"); ok {
		maskTargetDatabaseRequest.Recompile = oci_data_safe.SubsettingPolicyRecompileEnum(recompile.(string))
	}

	maskTargetDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.SubsetData(ctx, maskTargetDatabaseRequest)
	if err != nil {
		return err
	}
	if response.OpcWorkRequestId == nil {
		return fmt.Errorf("subset operation did not return a work request ID")
	}
	workId := response.OpcWorkRequestId
	_, err = subsettingPolicyWaitForWorkRequest(ctx, workId, "subsettingReport",
		oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate),
		s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}

	return s.GetSubsettingReportList(ctx)
}

func (s *DataSafeSubsettingReportManagementResourceCurd) DeleteWithContext(ctx context.Context) error {
	reportID := s.D.Id()
	request := oci_data_safe.DeleteSubsettingReportRequest{SubsettingReportId: &reportID}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")
	response, err := s.Client.DeleteSubsettingReport(ctx, request)
	if err != nil {
		return err
	}
	if response.OpcWorkRequestId == nil {
		return fmt.Errorf("delete subsetting report operation did not return a work request ID")
	}

	_, err = subsettingPolicyWaitForWorkRequest(ctx, response.OpcWorkRequestId, "subsettingReport",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return err
}

func (s *DataSafeSubsettingReportManagementResourceCurd) GetSubsettingReportList(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingReportsRequest{}
	var subsettingReport = new(oci_data_safe.SubsettingReport)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
		request.SortOrder = oci_data_safe.ListSubsettingReportsSortOrderDesc
		request.SortBy = oci_data_safe.ListSubsettingReportsSortByTimesubsettingfinished
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListSubsettingReports(ctx, request)
	if err != nil {
		return err
	}
	if response.SubsettingReportCollection.Items != nil && len(response.SubsettingReportCollection.Items) > 0 {
		temp1 := response.SubsettingReportCollection.Items[0]
		subsettingReport.Id = temp1.Id
	}

	if subsettingReport.Id == nil {
		return nil
	}

	s.D.SetId(*subsettingReport.Id)
	return nil
}

func (s *DataSafeSubsettingReportManagementResourceCurd) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingReportRequest{}

	tmp := s.D.Id()
	request.SubsettingReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSubsettingReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SubsettingReport
	return nil
}

func (s *DataSafeSubsettingReportManagementResourceCurd) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.SubsettingPolicyId != nil {
		s.D.Set("subsetting_policy_id", *s.Res.SubsettingPolicyId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.MaskingReportId != nil {
		s.D.Set("masking_report_id", *s.Res.MaskingReportId)
	}

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.MaskingWorkRequestId != nil {
		s.D.Set("masking_work_request_id", *s.Res.MaskingWorkRequestId)
	}

	if s.Res.IsRedoLoggingEnabled != nil {
		s.D.Set("is_redo_logging_enabled", *s.Res.IsRedoLoggingEnabled)
	}

	if s.Res.IsRefreshStatsEnabled != nil {
		s.D.Set("is_refresh_stats_enabled", *s.Res.IsRefreshStatsEnabled)
	}

	if s.Res.SubsettingWorkRequestId != nil {
		s.D.Set("subsetting_work_request_id", *s.Res.SubsettingWorkRequestId)
	}

	if s.Res.ParallelDegree != nil {
		s.D.Set("parallel_degree", *s.Res.ParallelDegree)
	}

	if s.Res.Recompile != nil {
		s.D.Set("recompile", *s.Res.Recompile)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeSubsettingFinished != nil {
		s.D.Set("time_subsetting_finished", s.Res.TimeSubsettingFinished.String())
	}

	if s.Res.TimeSubsettingStarted != nil {
		s.D.Set("time_subsetting_started", s.Res.TimeSubsettingStarted.String())
	}

	s.D.Set("subsetting_status", s.Res.SubsettingStatus)

	if s.Res.TotalSubsettedObjects != nil {
		s.D.Set("total_subsetted_objects", strconv.FormatInt(*s.Res.TotalSubsettedObjects, 10))
	}

	if s.Res.TotalSubsettedSchemas != nil {
		s.D.Set("total_subsetted_schemas", strconv.FormatInt(*s.Res.TotalSubsettedSchemas, 10))
	}

	if s.Res.TotalSubsettedRows != nil {
		s.D.Set("total_subsetted_rows", strconv.FormatInt(*s.Res.TotalSubsettedRows, 10))
	}

	if s.Res.DatabaseSizeBeforeSubsettingInKBs != nil {
		s.D.Set("database_size_before_subsetting_in_kbs", *s.Res.DatabaseSizeBeforeSubsettingInKBs)
	}

	if s.Res.DatabaseSizeAfterSubsettingInKBs != nil {
		s.D.Set("database_size_after_subsetting_in_kbs", *s.Res.DatabaseSizeAfterSubsettingInKBs)
	}

	if s.Res.TotalPreSubsettingScriptErrors != nil {
		s.D.Set("total_pre_subsetting_script_errors", strconv.FormatInt(*s.Res.TotalPreSubsettingScriptErrors, 10))
	}

	if s.Res.TotalPostSubsettingScriptErrors != nil {
		s.D.Set("total_post_subsetting_script_errors", strconv.FormatInt(*s.Res.TotalPostSubsettingScriptErrors, 10))
	}

	return nil
}
