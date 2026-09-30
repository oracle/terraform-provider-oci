// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDataSafeSubsettingPolicyWithContext,
		ReadContext:   readDataSafeSubsettingPolicyWithContext,
		UpdateContext: updateDataSafeSubsettingPolicyWithContext,
		DeleteContext: deleteDataSafeSubsettingPolicyWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schema_source": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"schema_source": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SENSITIVE_DATA_MODEL",
								"TARGET",
							}, true),
						},

						// Optional
						"schemas_for_subsetting": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringLenBetween(1, 128),
							},
						},
						"sensitive_data_model_id": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringLenBetween(1, 255),
						},
						"target_id": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringLenBetween(1, 255),
						},

						// Computed
						"derived_schemas": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_redo_logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_refresh_stats_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parallel_degree": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateSubsettingParallelDegree,
			},
			"post_subsetting_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_subsetting_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recompile": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"SERIAL", "PARALLEL", "NONE"}, true),
			},
			"unrelated_tables_action": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"TRUNCATE", "KEEP"}, true),
			},
			"check_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"ALL", "TABLESPACE_CHECK"}, false),
			},
			"tablespace": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 30),
			},
			"target_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
			},
			"target_credentials": {
				Type: schema.TypeList, Optional: true, MinItems: 1, MaxItems: 1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"user_name": {Type: schema.TypeString, Required: true},
					"password":  {Type: schema.TypeString, Required: true, Sensitive: true},
				}},
			},
			"generate_health_report_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeSubsettingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	return nil

}

func readDataSafeSubsettingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDataSafeSubsettingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	shouldGenerateHealthReport := false
	if _, ok := sync.D.GetOkExists("generate_health_report_trigger"); ok && sync.D.HasChange("generate_health_report_trigger") {
		oldRaw, newRaw := sync.D.GetChange("generate_health_report_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			shouldGenerateHealthReport = true
		} else {
			sync.D.Set("generate_health_report_trigger", oldRaw)
			return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	if shouldGenerateHealthReport {
		if err := sync.GenerateSubsettingHealthReport(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	return nil
}

func deleteDataSafeSubsettingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DataSafeSubsettingPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.SubsettingPolicy
	DisableNotFoundRetries bool
}

func (s *DataSafeSubsettingPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeSubsettingPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.SubsettingPolicyLifecycleStateCreating),
	}
}

func (s *DataSafeSubsettingPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.SubsettingPolicyLifecycleStateActive),
		string(oci_data_safe.SubsettingPolicyLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeSubsettingPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.SubsettingPolicyLifecycleStateDeleting),
	}
}

func (s *DataSafeSubsettingPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.SubsettingPolicyLifecycleStateDeleted),
	}
}

func (s *DataSafeSubsettingPolicyResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_data_safe.CreateSubsettingPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isRedoLoggingEnabled, ok := s.D.GetOkExists("is_redo_logging_enabled"); ok {
		tmp := isRedoLoggingEnabled.(bool)
		request.IsRedoLoggingEnabled = &tmp
	}

	if isRefreshStatsEnabled, ok := s.D.GetOkExists("is_refresh_stats_enabled"); ok {
		tmp := isRefreshStatsEnabled.(bool)
		request.IsRefreshStatsEnabled = &tmp
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if parallelDegree, ok := s.D.GetOkExists("parallel_degree"); ok {
		tmp := parallelDegree.(string)
		request.ParallelDegree = &tmp
	}

	if postSubsettingScript, ok := s.D.GetOkExists("post_subsetting_script"); ok {
		tmp := postSubsettingScript.(string)
		request.PostSubsettingScript = &tmp
	}

	if preSubsettingScript, ok := s.D.GetOkExists("pre_subsetting_script"); ok {
		tmp := preSubsettingScript.(string)
		request.PreSubsettingScript = &tmp
	}

	if recompile, ok := s.D.GetOkExists("recompile"); ok {
		request.Recompile = oci_data_safe.SubsettingPolicyRecompileEnum(recompile.(string))
	}

	if schemaSource, ok := s.D.GetOkExists("schema_source"); ok {
		if tmpList := schemaSource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schema_source", 0)
			tmp, err := s.mapToCreateSchemaSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SchemaSource = tmp
		}
	}

	if unrelatedTablesAction, ok := s.D.GetOkExists("unrelated_tables_action"); ok {
		request.UnrelatedTablesAction = oci_data_safe.SubsettingPolicyUnrelatedTablesActionEnum(unrelatedTablesAction.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateSubsettingPolicy(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSubsettingPolicyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeSubsettingPolicyResourceCrud) getSubsettingPolicyFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	subsettingPolicyId, err := subsettingPolicyWaitForWorkRequest(ctx, workId, "subsettingpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, subsettingPolicyId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_data_safe.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*subsettingPolicyId)

	return s.GetWithContext(ctx)
}

func subsettingPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func subsettingPolicyWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = subsettingPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_data_safe.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if dataSafeWorkRequestEntityTypeMatches(res.EntityType, entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeSubsettingPolicyWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

// dataSafeWorkRequestEntityTypeMatches accounts for the underscore used by the
// Data Safe service in work-request entity types (for example,
// "subsetting_Policy").
func dataSafeWorkRequestEntityTypeMatches(actual *string, expected string) bool {
	if actual == nil {
		return false
	}

	normalizedActual := strings.ReplaceAll(strings.ToLower(*actual), "_", "")
	normalizedExpected := strings.ReplaceAll(strings.ToLower(expected), "_", "")
	return strings.Contains(normalizedActual, normalizedExpected)
}

func getErrorFromDataSafeSubsettingPolicyWorkRequest(ctx context.Context, client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_data_safe.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DataSafeSubsettingPolicyResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingPolicyRequest{}

	tmp := s.D.Id()
	request.SubsettingPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetSubsettingPolicy(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.SubsettingPolicy
	return nil
}

func (s *DataSafeSubsettingPolicyResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateSubsettingPolicyRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isRedoLoggingEnabled, ok := s.D.GetOkExists("is_redo_logging_enabled"); ok {
		tmp := isRedoLoggingEnabled.(bool)
		request.IsRedoLoggingEnabled = &tmp
	}

	if isRefreshStatsEnabled, ok := s.D.GetOkExists("is_refresh_stats_enabled"); ok {
		tmp := isRefreshStatsEnabled.(bool)
		request.IsRefreshStatsEnabled = &tmp
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if parallelDegree, ok := s.D.GetOkExists("parallel_degree"); ok {
		tmp := parallelDegree.(string)
		request.ParallelDegree = &tmp
	}

	if postSubsettingScript, ok := s.D.GetOkExists("post_subsetting_script"); ok {
		tmp := postSubsettingScript.(string)
		request.PostSubsettingScript = &tmp
	}

	if preSubsettingScript, ok := s.D.GetOkExists("pre_subsetting_script"); ok {
		tmp := preSubsettingScript.(string)
		request.PreSubsettingScript = &tmp
	}

	if recompile, ok := s.D.GetOkExists("recompile"); ok {
		request.Recompile = oci_data_safe.SubsettingPolicyRecompileEnum(recompile.(string))
	}

	if schemaSource, ok := s.D.GetOkExists("schema_source"); ok {
		if tmpList := schemaSource.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schema_source", 0)
			tmp, err := s.mapToUpdateSchemaSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SchemaSource = tmp
		}
	}

	tmp := s.D.Id()
	request.SubsettingPolicyId = &tmp

	if unrelatedTablesAction, ok := s.D.GetOkExists("unrelated_tables_action"); ok {
		request.UnrelatedTablesAction = oci_data_safe.SubsettingPolicyUnrelatedTablesActionEnum(unrelatedTablesAction.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateSubsettingPolicy(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubsettingPolicyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeSubsettingPolicyResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_data_safe.DeleteSubsettingPolicyRequest{}

	tmp := s.D.Id()
	request.SubsettingPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteSubsettingPolicy(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := subsettingPolicyWaitForWorkRequest(ctx, workId, "subsettingpolicy",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeSubsettingPolicyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRedoLoggingEnabled != nil {
		s.D.Set("is_redo_logging_enabled", *s.Res.IsRedoLoggingEnabled)
	}

	if s.Res.IsRefreshStatsEnabled != nil {
		s.D.Set("is_refresh_stats_enabled", *s.Res.IsRefreshStatsEnabled)
	}

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.ParallelDegree != nil {
		s.D.Set("parallel_degree", *s.Res.ParallelDegree)
	}

	if s.Res.PostSubsettingScript != nil {
		s.D.Set("post_subsetting_script", *s.Res.PostSubsettingScript)
	}

	if s.Res.PreSubsettingScript != nil {
		s.D.Set("pre_subsetting_script", *s.Res.PreSubsettingScript)
	}

	s.D.Set("recompile", s.Res.Recompile)

	if s.Res.SchemaSource != nil {
		schemaSourceArray := []interface{}{}
		if schemaSourceMap := SchemaSourceDetailsToMap(&s.Res.SchemaSource); schemaSourceMap != nil {
			schemaSourceArray = append(schemaSourceArray, schemaSourceMap)
		}
		s.D.Set("schema_source", schemaSourceArray)
	} else {
		s.D.Set("schema_source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("unrelated_tables_action", s.Res.UnrelatedTablesAction)

	return nil
}

func (s *DataSafeSubsettingPolicyResourceCrud) GenerateSubsettingHealthReport(ctx context.Context) error {
	request := oci_data_safe.GenerateSubsettingHealthReportRequest{}

	if checkType, ok := s.D.GetOkExists("check_type"); ok {
		request.CheckType = oci_data_safe.GenerateSubsettingHealthReportDetailsCheckTypeEnum(checkType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	idTmp := s.D.Id()
	request.SubsettingPolicyId = &idTmp

	if tablespace, ok := s.D.GetOkExists("tablespace"); ok {
		tmp := tablespace.(string)
		request.Tablespace = &tmp
	}

	if targetCredentials, ok := s.D.GetOkExists("target_credentials"); ok {
		if tmpList := targetCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_credentials", 0)
			tmp, err := s.mapToCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetCredentials = &tmp
		}
	}
	if request.TargetCredentials == nil {
		return fmt.Errorf("target_credentials must be specified when generate_health_report_trigger is changed")
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GenerateSubsettingHealthReport(ctx, request)
	if err != nil {
		return err
	}

	if response.OpcWorkRequestId == nil {
		return fmt.Errorf("generate health report operation did not return a work request ID")
	}

	if _, err = generateSubsettingHealthReportWaitForWorkRequest(ctx, response.OpcWorkRequestId, "subsettingPolicyHealthReport",
		oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client); err != nil {
		return err
	}

	val := s.D.Get("generate_health_report_trigger")
	return s.D.Set("generate_health_report_trigger", val)
}

func (s *DataSafeSubsettingPolicyResourceCrud) mapToCreateSchemaSourceDetails(fieldKeyFormat string) (oci_data_safe.CreateSchemaSourceDetails, error) {
	var baseObject oci_data_safe.CreateSchemaSourceDetails
	//discriminator
	schemaSourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_source"))
	var schemaSource string
	if ok {
		schemaSource = schemaSourceRaw.(string)
	} else {
		schemaSource = "" // default value
	}
	switch strings.ToLower(schemaSource) {
	case strings.ToLower("SENSITIVE_DATA_MODEL"):
		details := oci_data_safe.CreateSchemaSourceFromSdmDetails{}
		sensitiveDataModelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sensitive_data_model_id"))
		if !ok || strings.TrimSpace(sensitiveDataModelId.(string)) == "" {
			return nil, fmt.Errorf("sensitive_data_model_id must be specified when schema_source is SENSITIVE_DATA_MODEL")
		}
		tmp := sensitiveDataModelId.(string)
		details.SensitiveDataModelId = &tmp
		baseObject = details
	case strings.ToLower("TARGET"):
		details := oci_data_safe.CreateSchemaSourceFromTargetDetails{}
		targetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_id"))
		if !ok || strings.TrimSpace(targetId.(string)) == "" {
			return nil, fmt.Errorf("target_id must be specified when schema_source is TARGET")
		}
		tmp := targetId.(string)
		details.TargetId = &tmp
		if schemasForSubsetting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schemas_for_subsetting")); ok {
			interfaces := schemasForSubsetting.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "schemas_for_subsetting")) {
				details.SchemasForSubsetting = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown schema_source '%v' was specified", schemaSource)
	}
	return baseObject, nil
}

func (s *DataSafeSubsettingPolicyResourceCrud) mapToUpdateSchemaSourceDetails(fieldKeyFormat string) (oci_data_safe.UpdateSchemaSourceDetails, error) {
	var baseObject oci_data_safe.UpdateSchemaSourceDetails
	//discriminator
	schemaSourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schema_source"))
	var schemaSource string
	if ok {
		schemaSource = schemaSourceRaw.(string)
	} else {
		schemaSource = "" // default value
	}
	switch strings.ToLower(schemaSource) {
	case strings.ToLower("SENSITIVE_DATA_MODEL"):
		details := oci_data_safe.UpdateSchemaSourceFromSdmDetails{}
		sensitiveDataModelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sensitive_data_model_id"))
		if !ok || strings.TrimSpace(sensitiveDataModelId.(string)) == "" {
			return nil, fmt.Errorf("sensitive_data_model_id must be specified when schema_source is SENSITIVE_DATA_MODEL")
		}
		tmp := sensitiveDataModelId.(string)
		details.SensitiveDataModelId = &tmp
		baseObject = details
	case strings.ToLower("TARGET"):
		details := oci_data_safe.UpdateSchemaSourceFromTargetDetails{}
		targetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_id"))
		if !ok || strings.TrimSpace(targetId.(string)) == "" {
			return nil, fmt.Errorf("target_id must be specified when schema_source is TARGET")
		}
		tmpTargetId := targetId.(string)
		details.TargetId = &tmpTargetId
		if schemasForSubsetting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schemas_for_subsetting")); ok {
			interfaces := schemasForSubsetting.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "schemas_for_subsetting")) {
				details.SchemasForSubsetting = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown schema_source '%v' was specified", schemaSource)
	}
	return baseObject, nil
}

func SchemaSourceDetailsToMap(obj *oci_data_safe.SchemaSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.SchemaSourceFromSdmDetails:
		result["schema_source"] = "SENSITIVE_DATA_MODEL"

		if v.SensitiveDataModelId != nil {
			result["sensitive_data_model_id"] = string(*v.SensitiveDataModelId)
		}

		result["derived_schemas"] = v.DerivedSchemas

		result["schemas_for_subsetting"] = v.SchemasForSubsetting
	case oci_data_safe.SchemaSourceFromTargetDetails:
		result["schema_source"] = "TARGET"

		if v.TargetId != nil {
			result["target_id"] = string(*v.TargetId)
		}

		result["derived_schemas"] = v.DerivedSchemas

		result["schemas_for_subsetting"] = v.SchemasForSubsetting
	default:
		log.Printf("[WARN] Received 'schema_source' of unknown type %v", *obj)
		return nil
	}

	return result
}

func SchemaSourceForSummaryToMap(obj *oci_data_safe.SchemaSourceForSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.SchemaSourceFromSdmForSummary:
		result["schema_source"] = "SENSITIVE_DATA_MODEL"

		if v.SensitiveDataModelId != nil {
			result["sensitive_data_model_id"] = string(*v.SensitiveDataModelId)
		}
	case oci_data_safe.SchemaSourceFromTargetForSummary:
		result["schema_source"] = "TARGET"

		if v.TargetId != nil {
			result["target_id"] = string(*v.TargetId)
		}
	default:
		log.Printf("[WARN] Received 'schema_source' of unknown type %v", *obj)
		return nil
	}

	return result
}

func SubsettingPolicySummaryToMap(obj oci_data_safe.SubsettingPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MaskingPolicyId != nil {
		result["masking_policy_id"] = string(*obj.MaskingPolicyId)
	}

	if obj.SchemaSource != nil {
		schemaSourceArray := []interface{}{}
		if schemaSourceMap := SchemaSourceForSummaryToMap(&obj.SchemaSource); schemaSourceMap != nil {
			schemaSourceArray = append(schemaSourceArray, schemaSourceMap)
		}
		result["schema_source"] = schemaSourceArray
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataSafeSubsettingPolicyResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeSubsettingPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SubsettingPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeSubsettingPolicyCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DataSafeSubsettingPolicyResourceCrud) mapToCredentials(fieldKeyFormat string) (oci_data_safe.Credentials, error) {
	result := oci_data_safe.Credentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
		tmp := userName.(string)
		result.UserName = &tmp
	}

	return result, nil
}
