// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v35/common"
	oci_sch "github.com/oracle/oci-go-sdk/v35/sch"
)

func init() {
	RegisterResource("oci_sch_service_connector", SchServiceConnectorResource())
}

func SchServiceConnectorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createSchServiceConnector,
		Read:     readSchServiceConnector,
		Update:   updateSchServiceConnector,
		Delete:   deleteSchServiceConnector,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"logging",
							}, true),
						},
						"log_sources": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"compartment_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"log_group_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"target": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"functions",
								"loggingAnalytics",
								"monitoring",
								"notifications",
								"objectStorage",
								"streaming",
							}, true),
						},

						// Optional
						"batch_rollover_size_in_mbs": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"batch_rollover_time_in_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"function_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric_namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_name_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stream_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"topic_id": {
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
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			"tasks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"condition": {
							Type:     schema.TypeString,
							Required: true,
						},
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"logRule",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_sch.LifecycleStateInactive),
					string(oci_sch.LifecycleStateActive),
				}, true),
			},

			// Computed
			"lifecyle_details": {
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

func createSchServiceConnector(d *schema.ResourceData, m interface{}) error {
	sync := &SchServiceConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceConnectorClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_sch.LifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_sch.LifecycleStateInactive {
			powerOff = true
		}
	}

	if e := CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopSchResource(); err != nil {
			return err
		}
		sync.D.Set("state", oci_sch.LifecycleStateInactive)
	}
	return nil
}

func readSchServiceConnector(d *schema.ResourceData, m interface{}) error {
	sync := &SchServiceConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceConnectorClient()

	return ReadResource(sync)
}

func updateSchServiceConnector(d *schema.ResourceData, m interface{}) error {
	sync := &SchServiceConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceConnectorClient()

	// switch to power on
	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_sch.LifecycleStateActive == oci_sch.LifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_sch.LifecycleStateInactive == oci_sch.LifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartSchResource(); err != nil {
			return err
		}
		sync.D.Set("state", oci_sch.LifecycleStateActive)
	}

	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	// switch to power off
	if powerOff {
		if err := sync.StopSchResource(); err != nil {
			return err
		}
		sync.D.Set("state", oci_sch.LifecycleStateInactive)
	}
	return nil
}

func deleteSchServiceConnector(d *schema.ResourceData, m interface{}) error {
	sync := &SchServiceConnectorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceConnectorClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type SchServiceConnectorResourceCrud struct {
	BaseCrud
	Client                 *oci_sch.ServiceConnectorClient
	Res                    *oci_sch.ServiceConnector
	DisableNotFoundRetries bool
}

func (s *SchServiceConnectorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SchServiceConnectorResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_sch.LifecycleStateCreating),
	}
}

func (s *SchServiceConnectorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_sch.LifecycleStateActive),
	}
}

func (s *SchServiceConnectorResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_sch.LifecycleStateDeleting),
	}
}

func (s *SchServiceConnectorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_sch.LifecycleStateDeleted),
	}
}

func (s *SchServiceConnectorResourceCrud) Create() error {
	request := oci_sch.CreateServiceConnectorRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
			tmp, err := s.mapToSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Source = tmp
		}
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target", 0)
			tmp, err := s.mapToTargetDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Target = tmp
		}
	}

	if tasks, ok := s.D.GetOkExists("tasks"); ok {
		interfaces := tasks.([]interface{})
		tmp := make([]oci_sch.TaskDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tasks", stateDataIndex)
			converted, err := s.mapToTaskDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tasks") {
			request.Tasks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "sch")

	response, err := s.Client.CreateServiceConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getServiceConnectorFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "sch"), oci_sch.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *SchServiceConnectorResourceCrud) getServiceConnectorFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_sch.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	serviceConnectorId, err := serviceConnectorWaitForWorkRequest(workId, "serviceConnector",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] operation failed: %v for identifier: %v\n", workId, serviceConnectorId)
		return err
	}
	s.D.SetId(*serviceConnectorId)

	return s.Get()
}

func serviceConnectorWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "sch", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_sch.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func serviceConnectorWaitForWorkRequest(wId *string, entityType string, action oci_sch.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_sch.ServiceConnectorClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "sch")
	retryPolicy.ShouldRetryOperation = serviceConnectorWorkRequestShouldRetryFunc(timeout)

	response := oci_sch.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_sch.OperationStatusInProgress),
			string(oci_sch.OperationStatusAccepted),
			string(oci_sch.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_sch.OperationStatusSucceeded),
			string(oci_sch.OperationStatusFailed),
			string(oci_sch.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_sch.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	//return identifier, workRequestErr
	if identifier == nil || response.Status == oci_sch.OperationStatusFailed || response.Status == oci_sch.OperationStatusCanceled {
		return nil, getErrorFromServiceConnectorWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromServiceConnectorWorkRequest(client *oci_sch.ServiceConnectorClient, wId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_sch.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_sch.ListWorkRequestErrorsRequest{
			WorkRequestId: wId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *SchServiceConnectorResourceCrud) Get() error {
	request := oci_sch.GetServiceConnectorRequest{}

	tmp := s.D.Id()
	request.ServiceConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "sch")

	response, err := s.Client.GetServiceConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceConnector
	return nil
}

func (s *SchServiceConnectorResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_sch.UpdateServiceConnectorRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ServiceConnectorId = &tmp

	if source, ok := s.D.GetOkExists("source"); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
			tmp, err := s.mapToSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Source = tmp
		}
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target", 0)
			tmp, err := s.mapToTargetDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Target = tmp
		}
	}

	if tasks, ok := s.D.GetOkExists("tasks"); ok {
		interfaces := tasks.([]interface{})
		tmp := make([]oci_sch.TaskDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tasks", stateDataIndex)
			converted, err := s.mapToTaskDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tasks") {
			request.Tasks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "sch")

	response, err := s.Client.UpdateServiceConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getServiceConnectorFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "sch"), oci_sch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *SchServiceConnectorResourceCrud) Delete() error {
	request := oci_sch.DeleteServiceConnectorRequest{}

	tmp := s.D.Id()
	request.ServiceConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "sch")

	response, err := s.Client.DeleteServiceConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := serviceConnectorWaitForWorkRequest(workId, "serviceConnector",
		oci_sch.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *SchServiceConnectorResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := SourceDetailsToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		s.D.Set("source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.Target != nil {
		targetArray := []interface{}{}
		if targetMap := TargetDetailsToMap(&s.Res.Target); targetMap != nil {
			targetArray = append(targetArray, targetMap)
		}
		s.D.Set("target", targetArray)
	} else {
		s.D.Set("target", nil)
	}

	tasks := []interface{}{}
	for _, item := range s.Res.Tasks {
		tasks = append(tasks, TaskDetailsToMap(item))
	}
	s.D.Set("tasks", tasks)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *SchServiceConnectorResourceCrud) mapToLogSource(fieldKeyFormat string) (oci_sch.LogSource, error) {
	result := oci_sch.LogSource{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func LogSourceToMap(obj oci_sch.LogSource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func ServiceConnectorSummaryToMap(obj oci_sch.ServiceConnectorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = systemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *SchServiceConnectorResourceCrud) mapToSourceDetails(fieldKeyFormat string) (oci_sch.SourceDetails, error) {
	var baseObject oci_sch.SourceDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("logging"):
		details := oci_sch.LoggingSourceDetails{}
		if logSources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_sources")); ok {
			interfaces := logSources.([]interface{})
			tmp := make([]oci_sch.LogSource, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "log_sources"), stateDataIndex)
				converted, err := s.mapToLogSource(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "log_sources")) {
				details.LogSources = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func SourceDetailsToMap(obj *oci_sch.SourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_sch.LoggingSourceDetails:
		result["kind"] = "logging"

		logSources := []interface{}{}
		for _, item := range v.LogSources {
			logSources = append(logSources, LogSourceToMap(item))
		}
		result["log_sources"] = logSources
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *SchServiceConnectorResourceCrud) mapToTargetDetails(fieldKeyFormat string) (oci_sch.TargetDetails, error) {
	var baseObject oci_sch.TargetDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("functions"):
		details := oci_sch.FunctionsTargetDetails{}
		if functionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "function_id")); ok {
			tmp := functionId.(string)
			details.FunctionId = &tmp
		}
		baseObject = details
	case strings.ToLower("loggingAnalytics"):
		details := oci_sch.LoggingAnalyticsTargetDetails{}
		if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
			tmp := logGroupId.(string)
			details.LogGroupId = &tmp
		}
		baseObject = details
	case strings.ToLower("monitoring"):
		details := oci_sch.MonitoringTargetDetails{}
		if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if metric, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric")); ok {
			tmp := metric.(string)
			details.Metric = &tmp
		}
		if metricNamespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_namespace")); ok {
			tmp := metricNamespace.(string)
			details.MetricNamespace = &tmp
		}
		baseObject = details
	case strings.ToLower("notifications"):
		details := oci_sch.NotificationsTargetDetails{}
		if topicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "topic_id")); ok {
			tmp := topicId.(string)
			details.TopicId = &tmp
		}
		baseObject = details
	case strings.ToLower("objectStorage"):
		details := oci_sch.ObjectStorageTargetDetails{}
		if batchRolloverSizeInMBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_rollover_size_in_mbs")); ok {
			tmp := batchRolloverSizeInMBs.(int)
			details.BatchRolloverSizeInMBs = &tmp
		}
		if batchRolloverTimeInMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_rollover_time_in_ms")); ok {
			tmp := batchRolloverTimeInMs.(int)
			details.BatchRolloverTimeInMs = &tmp
		}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if objectNamePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_name_prefix")); ok {
			tmp := objectNamePrefix.(string)
			details.ObjectNamePrefix = &tmp
		}
		baseObject = details
	case strings.ToLower("streaming"):
		details := oci_sch.StreamingTargetDetails{}
		if streamId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_id")); ok {
			tmp := streamId.(string)
			details.StreamId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func TargetDetailsToMap(obj *oci_sch.TargetDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_sch.FunctionsTargetDetails:
		result["kind"] = "functions"

		if v.FunctionId != nil {
			result["function_id"] = string(*v.FunctionId)
		}
	case oci_sch.LoggingAnalyticsTargetDetails:
		result["kind"] = "loggingAnalytics"

		if v.LogGroupId != nil {
			result["log_group_id"] = string(*v.LogGroupId)
		}
	case oci_sch.MonitoringTargetDetails:
		result["kind"] = "monitoring"

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.Metric != nil {
			result["metric"] = string(*v.Metric)
		}

		if v.MetricNamespace != nil {
			result["metric_namespace"] = string(*v.MetricNamespace)
		}
	case oci_sch.NotificationsTargetDetails:
		result["kind"] = "notifications"

		if v.TopicId != nil {
			result["topic_id"] = string(*v.TopicId)
		}
	case oci_sch.ObjectStorageTargetDetails:
		result["kind"] = "objectStorage"

		if v.BatchRolloverSizeInMBs != nil {
			result["batch_rollover_size_in_mbs"] = int(*v.BatchRolloverSizeInMBs)
		}

		if v.BatchRolloverTimeInMs != nil {
			result["batch_rollover_time_in_ms"] = int(*v.BatchRolloverTimeInMs)
		}

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ObjectNamePrefix != nil {
			result["object_name_prefix"] = string(*v.ObjectNamePrefix)
		}
	case oci_sch.StreamingTargetDetails:
		result["kind"] = "streaming"

		if v.StreamId != nil {
			result["stream_id"] = string(*v.StreamId)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *SchServiceConnectorResourceCrud) mapToTaskDetails(fieldKeyFormat string) (oci_sch.TaskDetails, error) {
	var baseObject oci_sch.TaskDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("logRule"):
		details := oci_sch.LogRuleTaskDetails{}
		if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
			tmp := condition.(string)
			details.Condition = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func TaskDetailsToMap(obj oci_sch.TaskDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_sch.LogRuleTaskDetails:
		result["kind"] = "logRule"

		if v.Condition != nil {
			result["condition"] = string(*v.Condition)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *SchServiceConnectorResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_sch.ChangeServiceConnectorCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ServiceConnectorId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "sch")

	response, err := s.Client.ChangeServiceConnectorCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, updateWorkRequestErr := serviceConnectorWaitForWorkRequest(workId, "serviceConnector",
		oci_sch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	return updateWorkRequestErr
}

func (s *SchServiceConnectorResourceCrud) StartSchResource() error {
	request := oci_sch.ActivateServiceConnectorRequest{}

	tmp := s.D.Id()
	request.ServiceConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "sch")

	response, err := s.Client.ActivateServiceConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getServiceConnectorFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "sch"), oci_sch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *SchServiceConnectorResourceCrud) StopSchResource() error {
	request := oci_sch.DeactivateServiceConnectorRequest{}

	tmp := s.D.Id()
	request.ServiceConnectorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "sch")

	response, err := s.Client.DeactivateServiceConnector(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getServiceConnectorFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "sch"), oci_sch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
