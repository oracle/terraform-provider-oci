// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeReportDefinitionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeReportDefinition,
		Read:     readDataSafeReportDefinition,
		Update:   updateDataSafeReportDefinition,
		Delete:   deleteDataSafeReportDefinition,
		Schema: map[string]*schema.Schema{
			// Required
			"column_filters": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"expressions": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"field_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"is_hidden": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"column_info": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_order": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"field_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_hidden": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"data_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"column_sortings": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"field_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_ascending": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"sorting_order": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"summary": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"display_order": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"count_of": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_by_field_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_hidden": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"scim_filter": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compliance_standards": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_order": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_seeded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"record_time_span": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_report_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_report_mime_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_report_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_report_row_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scim_filter": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createDataSafeReportDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeReportDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeReportDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeReportDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeReportDefinitionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.ReportDefinition
	DisableNotFoundRetries bool
}

func (s *DataSafeReportDefinitionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeReportDefinitionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.ReportDefinitionLifecycleStateCreating),
	}
}

func (s *DataSafeReportDefinitionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.ReportDefinitionLifecycleStateActive),
	}
}

func (s *DataSafeReportDefinitionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.ReportDefinitionLifecycleStateDeleting),
	}
}

func (s *DataSafeReportDefinitionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.ReportDefinitionLifecycleStateDeleted),
	}
}

func (s *DataSafeReportDefinitionResourceCrud) Create() error {
	request := oci_data_safe.CreateReportDefinitionRequest{}

	if columnFilters, ok := s.D.GetOkExists("column_filters"); ok {
		interfaces := columnFilters.([]interface{})
		tmp := make([]oci_data_safe.ColumnFilter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_filters", stateDataIndex)
			converted, err := s.mapTocolumnFilter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("column_filters") {
			request.ColumnFilters = tmp
		}
	}

	if columnInfo, ok := s.D.GetOkExists("column_info"); ok {
		interfaces := columnInfo.([]interface{})
		tmp := make([]oci_data_safe.Column, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_info", stateDataIndex)
			converted, err := s.mapTocolumn(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("column_info") {
			request.ColumnInfo = tmp
		}
	}

	if columnSortings, ok := s.D.GetOkExists("column_sortings"); ok {
		interfaces := columnSortings.([]interface{})
		tmp := make([]oci_data_safe.ColumnSorting, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_sortings", stateDataIndex)
			converted, err := s.mapTocolumnSorting(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("column_sortings") {
			request.ColumnSortings = tmp
		}
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if parentId, ok := s.D.GetOkExists("parent_id"); ok {
		tmp := parentId.(string)
		request.ParentId = &tmp
	}

	if summary, ok := s.D.GetOkExists("summary"); ok {
		interfaces := summary.([]interface{})
		tmp := make([]oci_data_safe.Summary, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "summary", stateDataIndex)
			converted, err := s.mapTosummary(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("summary") {
			request.Summary = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateReportDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getReportDefinitionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeReportDefinitionResourceCrud) getReportDefinitionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	reportDefinitionId, err := reportDefinitionWaitForWorkRequest(workId, "reportdefinition",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*reportDefinitionId)

	return s.Get()
}

func reportDefinitionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func reportDefinitionWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = reportDefinitionWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
			string(oci_data_safe.WorkRequestStatusInProgress),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
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
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeReportDefinitionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	fmt.Printf("TestDataSafeReportDefinitionResource_basic  ***** CALLED5 ***** \n")
	return identifier, nil
}

func getErrorFromDataSafeReportDefinitionWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
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

func (s *DataSafeReportDefinitionResourceCrud) Get() error {
	request := oci_data_safe.GetReportDefinitionRequest{}

	tmp := s.D.Id()
	request.ReportDefinitionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetReportDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ReportDefinition
	return nil
}

func (s *DataSafeReportDefinitionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateReportDefinitionRequest{}

	if columnFilters, ok := s.D.GetOkExists("column_filters"); ok {
		interfaces := columnFilters.([]interface{})
		tmp := make([]oci_data_safe.ColumnFilter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_filters", stateDataIndex)
			converted, err := s.mapTocolumnFilter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("column_filters") {
			request.ColumnFilters = tmp
		}
	}

	if columnInfo, ok := s.D.GetOkExists("column_info"); ok {
		interfaces := columnInfo.([]interface{})
		tmp := make([]oci_data_safe.Column, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_info", stateDataIndex)
			converted, err := s.mapTocolumn(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("column_info") {
			request.ColumnInfo = tmp
		}
	}

	if columnSortings, ok := s.D.GetOkExists("column_sortings"); ok {
		interfaces := columnSortings.([]interface{})
		tmp := make([]oci_data_safe.ColumnSorting, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "column_sortings", stateDataIndex)
			converted, err := s.mapTocolumnSorting(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("column_sortings") {
			request.ColumnSortings = tmp
		}
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

	tmp := s.D.Id()
	request.ReportDefinitionId = &tmp

	if summary, ok := s.D.GetOkExists("summary"); ok {
		interfaces := summary.([]interface{})
		tmp := make([]oci_data_safe.Summary, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "summary", stateDataIndex)
			converted, err := s.mapTosummary(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("summary") {
			request.Summary = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateReportDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getReportDefinitionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeReportDefinitionResourceCrud) Delete() error {
	request := oci_data_safe.DeleteReportDefinitionRequest{}

	tmp := s.D.Id()
	request.ReportDefinitionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteReportDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := reportDefinitionWaitForWorkRequest(workId, "reportdefinition",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeReportDefinitionResourceCrud) SetData() error {
	s.D.Set("category", s.Res.Category)

	columnFilters := []interface{}{}
	for _, item := range s.Res.ColumnFilters {
		columnFilters = append(columnFilters, columnFilterToMap(item))
	}
	s.D.Set("column_filters", columnFilters)

	columnInfo := []interface{}{}
	for _, item := range s.Res.ColumnInfo {
		columnInfo = append(columnInfo, columnToMap(item))
	}
	s.D.Set("column_info", columnInfo)

	columnSortings := []interface{}{}
	for _, item := range s.Res.ColumnSortings {
		columnSortings = append(columnSortings, columnSortingToMap(item))
	}
	s.D.Set("column_sortings", columnSortings)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compliance_standards", s.Res.ComplianceStandards)
	s.D.Set("compliance_standards", s.Res.ComplianceStandards)

	s.D.Set("data_source", s.Res.DataSource)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DisplayOrder != nil {
		s.D.Set("display_order", *s.Res.DisplayOrder)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSeeded != nil {
		s.D.Set("is_seeded", *s.Res.IsSeeded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ParentId != nil {
		s.D.Set("parent_id", *s.Res.ParentId)
	}

	if s.Res.RecordTimeSpan != nil {
		s.D.Set("record_time_span", *s.Res.RecordTimeSpan)
	}

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	if s.Res.ScheduledReportCompartmentId != nil {
		s.D.Set("scheduled_report_compartment_id", *s.Res.ScheduledReportCompartmentId)
	}

	s.D.Set("scheduled_report_mime_type", s.Res.ScheduledReportMimeType)

	if s.Res.ScheduledReportName != nil {
		s.D.Set("scheduled_report_name", *s.Res.ScheduledReportName)
	}

	if s.Res.ScheduledReportRowLimit != nil {
		s.D.Set("scheduled_report_row_limit", *s.Res.ScheduledReportRowLimit)
	}

	if s.Res.ScimFilter != nil {
		s.D.Set("scim_filter", *s.Res.ScimFilter)
	}

	s.D.Set("state", s.Res.LifecycleState)

	summary := []interface{}{}
	for _, item := range s.Res.Summary {
		summary = append(summary, summaryToMap(item))
	}
	s.D.Set("summary", summary)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ReportDefinitionSummaryToMap(obj oci_data_safe.ReportDefinitionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["category"] = string(obj.Category)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["compliance_standards"] = obj.ComplianceStandards
	result["compliance_standards"] = obj.ComplianceStandards

	result["data_source"] = string(obj.DataSource)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DisplayOrder != nil {
		result["display_order"] = int(*obj.DisplayOrder)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsSeeded != nil {
		result["is_seeded"] = bool(*obj.IsSeeded)
	}

	if obj.Schedule != nil {
		result["schedule"] = string(*obj.Schedule)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataSafeReportDefinitionResourceCrud) mapTocolumn(fieldKeyFormat string) (oci_data_safe.Column, error) {
	result := oci_data_safe.Column{}

	if dataType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_type")); ok {
		tmp := dataType.(string)
		result.DataType = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if displayOrder, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_order")); ok {
		tmp := displayOrder.(int)
		result.DisplayOrder = &tmp
	}

	if fieldName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_name")); ok {
		tmp := fieldName.(string)
		result.FieldName = &tmp
	}

	if isHidden, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_hidden")); ok {
		tmp := isHidden.(bool)
		result.IsHidden = &tmp
	}

	return result, nil
}

func columnToMap(obj oci_data_safe.Column) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DisplayOrder != nil {
		result["display_order"] = int(*obj.DisplayOrder)
	}

	if obj.FieldName != nil {
		result["field_name"] = string(*obj.FieldName)
	}

	if obj.IsHidden != nil {
		result["is_hidden"] = bool(*obj.IsHidden)
	}

	return result
}

func (s *DataSafeReportDefinitionResourceCrud) mapTocolumnFilter(fieldKeyFormat string) (oci_data_safe.ColumnFilter, error) {
	result := oci_data_safe.ColumnFilter{}

	if expressions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expressions")); ok {
		interfaces := expressions.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "expressions")) {
			result.Expressions = tmp
		}
	}

	if fieldName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_name")); ok {
		tmp := fieldName.(string)
		result.FieldName = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if isHidden, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_hidden")); ok {
		tmp := isHidden.(bool)
		result.IsHidden = &tmp
	}

	if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
		result.Operator = oci_data_safe.ColumnFilterOperatorEnum(operator.(string))
	}

	return result, nil
}

func columnFilterToMap(obj oci_data_safe.ColumnFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["expressions"] = obj.Expressions

	if obj.FieldName != nil {
		result["field_name"] = string(*obj.FieldName)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.IsHidden != nil {
		result["is_hidden"] = bool(*obj.IsHidden)
	}

	result["operator"] = string(obj.Operator)

	return result
}

func (s *DataSafeReportDefinitionResourceCrud) mapTocolumnSorting(fieldKeyFormat string) (oci_data_safe.ColumnSorting, error) {
	result := oci_data_safe.ColumnSorting{}

	if fieldName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_name")); ok {
		tmp := fieldName.(string)
		result.FieldName = &tmp
	}

	if isAscending, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ascending")); ok {
		tmp := isAscending.(bool)
		result.IsAscending = &tmp
	}

	if sortingOrder, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sorting_order")); ok {
		tmp := sortingOrder.(int)
		result.SortingOrder = &tmp
	}

	return result, nil
}

func columnSortingToMap(obj oci_data_safe.ColumnSorting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FieldName != nil {
		result["field_name"] = string(*obj.FieldName)
	}

	if obj.IsAscending != nil {
		result["is_ascending"] = bool(*obj.IsAscending)
	}

	if obj.SortingOrder != nil {
		result["sorting_order"] = int(*obj.SortingOrder)
	}

	return result
}

func (s *DataSafeReportDefinitionResourceCrud) mapTosummary(fieldKeyFormat string) (oci_data_safe.Summary, error) {
	result := oci_data_safe.Summary{}

	if countOf, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "count_of")); ok {
		tmp := countOf.(string)
		result.CountOf = &tmp
	}

	if displayOrder, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_order")); ok {
		tmp := displayOrder.(int)
		result.DisplayOrder = &tmp
	}

	if groupByFieldName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_by_field_name")); ok {
		tmp := groupByFieldName.(string)
		result.GroupByFieldName = &tmp
	}

	if isHidden, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_hidden")); ok {
		tmp := isHidden.(bool)
		result.IsHidden = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if scimFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scim_filter")); ok {
		tmp := scimFilter.(string)
		if tmp != "" {
			result.ScimFilter = &tmp
		}
	}

	return result, nil
}

func summaryToMap(obj oci_data_safe.Summary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CountOf != nil {
		result["count_of"] = string(*obj.CountOf)
	}

	if obj.DisplayOrder != nil {
		result["display_order"] = int(*obj.DisplayOrder)
	}

	if obj.GroupByFieldName != nil {
		result["group_by_field_name"] = string(*obj.GroupByFieldName)
	}

	if obj.IsHidden != nil {
		result["is_hidden"] = bool(*obj.IsHidden)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ScimFilter != nil && *obj.ScimFilter != "" {
		result["scim_filter"] = string(*obj.ScimFilter)
	}

	return result
}

func (s *DataSafeReportDefinitionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeReportDefinitionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ReportDefinitionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeReportDefinitionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getReportDefinitionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
