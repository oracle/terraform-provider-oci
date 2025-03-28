// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiNewsReportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiNewsReport,
		Read:     readOpsiNewsReport,
		Update:   updateOpsiNewsReport,
		Delete:   deleteOpsiNewsReport,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_types": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Optional
						"actionable_insights_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"capacity_planning_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sql_insights_fleet_analysis_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sql_insights_performance_degradation_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sql_insights_plan_changes_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sql_insights_top_databases_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sql_insights_top_sql_by_insights_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sql_insights_top_sql_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						// Optional
						// Computed
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"news_frequency": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ons_topic_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"are_child_compartments_included": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"day_of_week": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"match_rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag_filters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"lifecycle_details": {
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

func createOpsiNewsReport(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiNewsReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiNewsReport(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiNewsReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiNewsReport(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiNewsReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiNewsReport(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiNewsReportResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiNewsReportResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.NewsReport
	DisableNotFoundRetries bool
}

func (s *OpsiNewsReportResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpsiNewsReportResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateCreating),
	}
}

func (s *OpsiNewsReportResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateActive),
		string(oci_opsi.LifecycleStateNeedsAttention),
	}
}

func (s *OpsiNewsReportResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleting),
	}
}

func (s *OpsiNewsReportResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleted),
	}
}

func (s *OpsiNewsReportResourceCrud) Create() error {
	request := oci_opsi.CreateNewsReportRequest{}

	if areChildCompartmentsIncluded, ok := s.D.GetOkExists("are_child_compartments_included"); ok {
		tmp := areChildCompartmentsIncluded.(bool)
		request.AreChildCompartmentsIncluded = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if contentTypes, ok := s.D.GetOkExists("content_types"); ok {
		if tmpList := contentTypes.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "content_types", 0)
			tmp, err := s.mapToNewsContentTypes(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ContentTypes = &tmp
		}
	}

	if dayOfWeek, ok := s.D.GetOkExists("day_of_week"); ok {
		request.DayOfWeek = oci_opsi.DayOfWeekEnum(dayOfWeek.(string))
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if locale, ok := s.D.GetOkExists("locale"); ok {
		request.Locale = oci_opsi.NewsLocaleEnum(locale.(string))
	}

	if matchRule, ok := s.D.GetOkExists("match_rule"); ok {
		request.MatchRule = oci_opsi.MatchRuleEnum(matchRule.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if newsFrequency, ok := s.D.GetOkExists("news_frequency"); ok {
		request.NewsFrequency = oci_opsi.NewsFrequencyEnum(newsFrequency.(string))
	}

	if onsTopicId, ok := s.D.GetOkExists("ons_topic_id"); ok {
		tmp := onsTopicId.(string)
		request.OnsTopicId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_opsi.ResourceStatusEnum(status.(string))
	}

	if tagFilters, ok := s.D.GetOkExists("tag_filters"); ok {
		interfaces := tagFilters.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("tag_filters") {
			request.TagFilters = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateNewsReport(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNewsReportFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpsiNewsReportResourceCrud) getNewsReportFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	newsReportId, err := newsReportWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*newsReportId)

	return s.Get()
}

func newsReportWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func newsReportWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = newsReportWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiNewsReportWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiNewsReportWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiNewsReportResourceCrud) Get() error {
	request := oci_opsi.GetNewsReportRequest{}

	tmp := s.D.Id()
	request.NewsReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetNewsReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NewsReport
	return nil
}

func (s *OpsiNewsReportResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateNewsReportRequest{}

	if areChildCompartmentsIncluded, ok := s.D.GetOkExists("are_child_compartments_included"); ok {
		tmp := areChildCompartmentsIncluded.(bool)
		request.AreChildCompartmentsIncluded = &tmp
	}

	if contentTypes, ok := s.D.GetOkExists("content_types"); ok {
		if tmpList := contentTypes.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "content_types", 0)
			tmp, err := s.mapToNewsContentTypes(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ContentTypes = &tmp
		}
	}

	if dayOfWeek, ok := s.D.GetOkExists("day_of_week"); ok {
		request.DayOfWeek = oci_opsi.DayOfWeekEnum(dayOfWeek.(string))
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if locale, ok := s.D.GetOkExists("locale"); ok {
		request.Locale = oci_opsi.NewsLocaleEnum(locale.(string))
	}

	if matchRule, ok := s.D.GetOkExists("match_rule"); ok {
		request.MatchRule = oci_opsi.MatchRuleEnum(matchRule.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if newsFrequency, ok := s.D.GetOkExists("news_frequency"); ok {
		request.NewsFrequency = oci_opsi.NewsFrequencyEnum(newsFrequency.(string))
	}

	tmp := s.D.Id()
	request.NewsReportId = &tmp

	if onsTopicId, ok := s.D.GetOkExists("ons_topic_id"); ok {
		tmp := onsTopicId.(string)
		request.OnsTopicId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_opsi.ResourceStatusEnum(status.(string))
	}

	if tagFilters, ok := s.D.GetOkExists("tag_filters"); ok {
		interfaces := tagFilters.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("tag_filters") {
			request.TagFilters = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateNewsReport(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNewsReportFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiNewsReportResourceCrud) Delete() error {
	request := oci_opsi.DeleteNewsReportRequest{}

	tmp := s.D.Id()
	request.NewsReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteNewsReport(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := newsReportWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiNewsReportResourceCrud) SetData() error {
	if s.Res.AreChildCompartmentsIncluded != nil {
		s.D.Set("are_child_compartments_included", *s.Res.AreChildCompartmentsIncluded)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContentTypes != nil {
		s.D.Set("content_types", []interface{}{NewsContentTypesToMap(s.Res.ContentTypes)})
	} else {
		s.D.Set("content_types", nil)
	}

	s.D.Set("day_of_week", s.Res.DayOfWeek)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("locale", s.Res.Locale)

	s.D.Set("match_rule", s.Res.MatchRule)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("news_frequency", s.Res.NewsFrequency)

	if s.Res.OnsTopicId != nil {
		s.D.Set("ons_topic_id", *s.Res.OnsTopicId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("tag_filters", s.Res.TagFilters)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *OpsiNewsReportResourceCrud) mapToNewsContentTypes(fieldKeyFormat string) (oci_opsi.NewsContentTypes, error) {
	result := oci_opsi.NewsContentTypes{}

	if actionableInsightsResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "actionable_insights_resources")); ok {
		interfaces := actionableInsightsResources.([]interface{})
		tmp := make([]oci_opsi.ActionableInsightsContentTypesResourceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.ActionableInsightsContentTypesResourceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "actionable_insights_resources")) {
			result.ActionableInsightsResources = tmp
		}
	}

	if capacityPlanningResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity_planning_resources")); ok {
		strArray := capacityPlanningResources.([]interface{})
		tmp := make([]oci_opsi.NewsContentTypesResourceEnum, len(strArray))
		for i := range strArray {
			switch strArray[i].(string) {
			case "HOST":
				tmp[i] = oci_opsi.NewsContentTypesResourceHost
			case "DATABASE":
				tmp[i] = oci_opsi.NewsContentTypesResourceDatabase
			case "EXADATA":
				tmp[i] = oci_opsi.NewsContentTypesResourceExadata
			default:
				return result, fmt.Errorf("Not a valid capacity planning resource was provided")
			}
		}
		result.CapacityPlanningResources = tmp
	}

	if sqlInsightsFleetAnalysisResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_insights_fleet_analysis_resources")); ok {
		interfaces := sqlInsightsFleetAnalysisResources.([]interface{})
		tmp := make([]oci_opsi.NewsSqlInsightsContentTypesResourceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.NewsSqlInsightsContentTypesResourceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sql_insights_fleet_analysis_resources")) {
			result.SqlInsightsFleetAnalysisResources = tmp
		}
	}

	if sqlInsightsPerformanceDegradationResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_insights_performance_degradation_resources")); ok {
		interfaces := sqlInsightsPerformanceDegradationResources.([]interface{})
		tmp := make([]oci_opsi.NewsSqlInsightsContentTypesResourceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.NewsSqlInsightsContentTypesResourceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sql_insights_performance_degradation_resources")) {
			result.SqlInsightsPerformanceDegradationResources = tmp
		}
	}

	if sqlInsightsPlanChangesResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_insights_plan_changes_resources")); ok {
		interfaces := sqlInsightsPlanChangesResources.([]interface{})
		tmp := make([]oci_opsi.NewsSqlInsightsContentTypesResourceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.NewsSqlInsightsContentTypesResourceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sql_insights_plan_changes_resources")) {
			result.SqlInsightsPlanChangesResources = tmp
		}
	}

	if sqlInsightsTopDatabasesResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_insights_top_databases_resources")); ok {
		interfaces := sqlInsightsTopDatabasesResources.([]interface{})
		tmp := make([]oci_opsi.NewsSqlInsightsContentTypesResourceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.NewsSqlInsightsContentTypesResourceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sql_insights_top_databases_resources")) {
			result.SqlInsightsTopDatabasesResources = tmp
		}
	}

	if sqlInsightsTopSqlByInsightsResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_insights_top_sql_by_insights_resources")); ok {
		interfaces := sqlInsightsTopSqlByInsightsResources.([]interface{})
		tmp := make([]oci_opsi.NewsSqlInsightsContentTypesResourceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.NewsSqlInsightsContentTypesResourceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sql_insights_top_sql_by_insights_resources")) {
			result.SqlInsightsTopSqlByInsightsResources = tmp
		}
	}

	if sqlInsightsTopSqlResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_insights_top_sql_resources")); ok {
		interfaces := sqlInsightsTopSqlResources.([]interface{})
		tmp := make([]oci_opsi.NewsSqlInsightsContentTypesResourceEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.NewsSqlInsightsContentTypesResourceEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sql_insights_top_sql_resources")) {
			result.SqlInsightsTopSqlResources = tmp
		}
	}

	return result, nil
}

func NewsContentTypesToMap(obj *oci_opsi.NewsContentTypes) map[string]interface{} {
	result := map[string]interface{}{}
	if obj.CapacityPlanningResources != nil && len(obj.CapacityPlanningResources) != 0 {
		capacityPlanningResources := []interface{}{}

		for _, item := range obj.CapacityPlanningResources {
			capacityPlanningResources = append(capacityPlanningResources, NewsContentTypesResourceToMap(item))
		}
		result["capacity_planning_resources"] = capacityPlanningResources
		return result
	} else if obj.SqlInsightsFleetAnalysisResources != nil && len(obj.SqlInsightsFleetAnalysisResources) != 0 {
		sqlInsightsFleetAnalysisResources := []interface{}{}

		for _, item := range obj.SqlInsightsFleetAnalysisResources {
			sqlInsightsFleetAnalysisResources = append(sqlInsightsFleetAnalysisResources, NewsSqlInsightsContentTypesResourceToMap(item))
		}
		result["sql_insights_fleet_analysis_resources"] = sqlInsightsFleetAnalysisResources
		return result
	} else if obj.SqlInsightsPerformanceDegradationResources != nil && len(obj.SqlInsightsPerformanceDegradationResources) != 0 {
		sqlInsightsPerformanceDegradationResources := []interface{}{}

		for _, item := range obj.SqlInsightsPerformanceDegradationResources {
			sqlInsightsPerformanceDegradationResources = append(sqlInsightsPerformanceDegradationResources, NewsSqlInsightsContentTypesResourceToMap(item))
		}
		result["sql_insights_performance_degradation_resources"] = sqlInsightsPerformanceDegradationResources
		return result
	} else if obj.SqlInsightsPlanChangesResources != nil && len(obj.SqlInsightsPlanChangesResources) != 0 {
		sqlInsightsPlanChangesResources := []interface{}{}

		for _, item := range obj.SqlInsightsPlanChangesResources {
			sqlInsightsPlanChangesResources = append(sqlInsightsPlanChangesResources, NewsSqlInsightsContentTypesResourceToMap(item))
		}
		result["sql_insights_plan_changes_resources"] = sqlInsightsPlanChangesResources
		return result
	} else if obj.SqlInsightsTopDatabasesResources != nil && len(obj.SqlInsightsTopDatabasesResources) != 0 {
		sqlInsightsTopDatabasesResources := []interface{}{}

		for _, item := range obj.SqlInsightsTopDatabasesResources {
			sqlInsightsTopDatabasesResources = append(sqlInsightsTopDatabasesResources, NewsSqlInsightsContentTypesResourceToMap(item))
		}
		result["sql_insights_top_databases_resources"] = sqlInsightsTopDatabasesResources
		return result
	} else if obj.SqlInsightsTopSqlByInsightsResources != nil && len(obj.SqlInsightsTopSqlByInsightsResources) != 0 {
		sqlInsightsTopSqlByInsightsResources := []interface{}{}

		for _, item := range obj.SqlInsightsTopSqlByInsightsResources {
			sqlInsightsTopSqlByInsightsResources = append(sqlInsightsTopSqlByInsightsResources, NewsSqlInsightsContentTypesResourceToMap(item))
		}
		result["sql_insights_top_sql_by_insights_resources"] = sqlInsightsTopSqlByInsightsResources
		return result
	} else if obj.SqlInsightsTopSqlResources != nil && len(obj.SqlInsightsTopSqlResources) != 0 {
		sqlInsightsTopSqlResources := []interface{}{}

		for _, item := range obj.SqlInsightsTopSqlResources {
			sqlInsightsTopSqlResources = append(sqlInsightsTopSqlResources, NewsSqlInsightsContentTypesResourceToMap(item))
		}
		result["sql_insights_top_sql_resources"] = sqlInsightsTopSqlResources
		return result
	} else {
		actionableInsightsResources := []interface{}{}

		for _, item := range obj.ActionableInsightsResources {
			actionableInsightsResources = append(actionableInsightsResources, NewsActionableInsightsContentTypesResourceToMap(item))
		}
		result["actionable_insights_resources"] = obj.ActionableInsightsResources
		return result
	}
}

func NewsContentTypesResourceToMap(obj oci_opsi.NewsContentTypesResourceEnum) string {
	var result string

	switch obj {
	case oci_opsi.NewsContentTypesResourceHost:
		result = "HOST"
	case oci_opsi.NewsContentTypesResourceDatabase:
		result = "DATABASE"
	case oci_opsi.NewsContentTypesResourceExadata:
		result = "EXADATA"
	default:
		fmt.Println("ERROR, Nota a valid resource")
	}
	return result
}

func NewsSqlInsightsContentTypesResourceToMap(obj oci_opsi.NewsSqlInsightsContentTypesResourceEnum) string {
	var result string

	switch obj {
	case oci_opsi.NewsSqlInsightsContentTypesResourceDatabase:
		result = "DATABASE"
	case oci_opsi.NewsSqlInsightsContentTypesResourceExadata:
		result = "EXADATA"
	default:
		fmt.Println("ERROR, Nota a valid resource")
	}
	return result
}

func NewsActionableInsightsContentTypesResourceToMap(obj oci_opsi.ActionableInsightsContentTypesResourceEnum) string {
	var result string

	switch obj {
	case oci_opsi.ActionableInsightsContentTypesResourceNewHighs:
		result = "NEW_HIGHS"
	case oci_opsi.ActionableInsightsContentTypesResourceBigChanges:
		result = "BIG_CHANGES"
	case oci_opsi.ActionableInsightsContentTypesResourceCurrentInventory:
		result = "CURRENT_INVENTORY"
	case oci_opsi.ActionableInsightsContentTypesResourceInventoryChanges:
		result = "INVENTORY_CHANGES"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetStatistics:
		result = "FLEET_STATISTICS"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisSummaryDbCount:
		result = "FLEET_ANALYSIS_SUMMARY_DB_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisSummarySqlAnalyzedCount:
		result = "FLEET_ANALYSIS_SUMMARY_SQL_ANALYZED_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisSummaryNewSqlCount:
		result = "FLEET_ANALYSIS_SUMMARY_NEW_SQL_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisSummaryBusiestDb:
		result = "FLEET_ANALYSIS_SUMMARY_BUSIEST_DB"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlCount:
		result = "FLEET_ANALYSIS_DEGRADING_SQL_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlByDb:
		result = "FLEET_ANALYSIS_DEGRADING_SQL_BY_DB"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisDegradingSqlBySqlId:
		result = "FLEET_ANALYSIS_DEGRADING_SQL_BY_SQL_ID"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesCount:
		result = "FLEET_ANALYSIS_PLAN_CHANGES_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesDbMostChanges:
		result = "FLEET_ANALYSIS_PLAN_CHANGES_DB_MOST_CHANGES"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdImproved:
		result = "FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_IMPROVED"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisPlanChangesBySqlIdDegraded:
		result = "FLEET_ANALYSIS_PLAN_CHANGES_BY_SQL_ID_DEGRADED"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsCount:
		result = "FLEET_ANALYSIS_INVALIDATION_STORMS_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisInvalidationStormsHighest:
		result = "FLEET_ANALYSIS_INVALIDATION_STORMS_HIGHEST"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesCount:
		result = "FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesByDb:
		result = "FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_DB"
	case oci_opsi.ActionableInsightsContentTypesResourceFleetAnalysisCursorSharingIssuesBySql:
		result = "FLEET_ANALYSIS_CURSOR_SHARING_ISSUES_BY_SQL"
	case oci_opsi.ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDbCount:
		result = "PERFORMANCE_DEGRADATION_SUMMARY_DB_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlAnalyzedCount:
		result = "PERFORMANCE_DEGRADATION_SUMMARY_SQL_ANALYZED_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePerformanceDegradationSummarySqlPerformanceTrendsCount:
		result = "PERFORMANCE_DEGRADATION_SUMMARY_SQL_PERFORMANCE_TRENDS_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePerformanceDegradationSummaryDegradedSqlCount:
		result = "PERFORMANCE_DEGRADATION_SUMMARY_DEGRADED_SQL_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePerformanceDegradationSummaryImprovedSqlCount:
		result = "PERFORMANCE_DEGRADATION_SUMMARY_IMPROVED_SQL_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePerformanceDegradationDbDegradedCount:
		result = "PERFORMANCE_DEGRADATION_DB_DEGRADED_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePerformanceDegradationSqlDegradedTable:
		result = "PERFORMANCE_DEGRADATION_SQL_DEGRADED_TABLE"
	case oci_opsi.ActionableInsightsContentTypesResourcePlanChangesSummaryDbCount:
		result = "PLAN_CHANGES_SUMMARY_DB_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePlanChangesSummarySqlAnalyzedCount:
		result = "PLAN_CHANGES_SUMMARY_SQL_ANALYZED_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePlanChangesSummaryPlanChangesCount:
		result = "PLAN_CHANGES_SUMMARY_PLAN_CHANGES_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePlanChangesSummaryImprovementsCount:
		result = "PLAN_CHANGES_SUMMARY_IMPROVEMENTS_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePlanChangesSummaryDegradationCount:
		result = "PLAN_CHANGES_SUMMARY_DEGRADATION_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourcePlanChangesTopPlanChangesTable:
		result = "PLAN_CHANGES_TOP_PLAN_CHANGES_TABLE"
	case oci_opsi.ActionableInsightsContentTypesResourceTopDbSummaryDbCount:
		result = "TOP_DB_SUMMARY_DB_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceTopDbSummarySqlAnalyzedCount:
		result = "TOP_DB_SUMMARY_SQL_ANALYZED_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceTopDbSummaryBusiestDb:
		result = "TOP_DB_SUMMARY_BUSIEST_DB"
	case oci_opsi.ActionableInsightsContentTypesResourceTopTable:
		result = "TOP_TABLE"
	case oci_opsi.ActionableInsightsContentTypesResourceCollectionDelayCount:
		result = "COLLECTION_DELAY_COUNT"
	case oci_opsi.ActionableInsightsContentTypesResourceCollectionDelayPreviousWeekCount:
		result = "COLLECTION_DELAY_PREVIOUS_WEEK_COUNT"
	default:
		fmt.Println("ERROR, Nota a valid resource")
	}
	return result
}

func NewsReportSummaryToMap(obj oci_opsi.NewsReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreChildCompartmentsIncluded != nil {
		result["are_child_compartments_included"] = bool(*obj.AreChildCompartmentsIncluded)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ContentTypes != nil {
		result["content_types"] = []interface{}{NewsContentTypesToMap(obj.ContentTypes)}
	}

	result["day_of_week"] = string(obj.DayOfWeek)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["locale"] = string(obj.Locale)

	result["match_rule"] = string(obj.MatchRule)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["news_frequency"] = string(obj.NewsFrequency)

	if obj.OnsTopicId != nil {
		result["ons_topic_id"] = string(*obj.OnsTopicId)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	result["tag_filters"] = obj.TagFilters

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *OpsiNewsReportResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeNewsReportCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NewsReportId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeNewsReportCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNewsReportFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
