// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_logging "github.com/oracle/oci-go-sdk/v56/logging"
)

func LoggingLogResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoggingLog,
		Read:     readLoggingLog,
		Update:   updateLoggingLog,
		Delete:   deleteLoggingLog,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"category": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"resource": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"service": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"source_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"OCISERVICE",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
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
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"retention_duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoggingLog(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readLoggingLog(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

func updateLoggingLog(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoggingLog(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoggingLogResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_logging.LoggingManagementClient
	Res                    *oci_logging.Log
	DisableNotFoundRetries bool
}

func (s *LoggingLogResourceCrud) ID() string {
	return GetLogCompositeId(*s.Res.LogGroupId, *s.Res.Id)
}

func (s *LoggingLogResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_logging.LogLifecycleStateCreating),
	}
}

func (s *LoggingLogResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_logging.LogLifecycleStateActive),
	}
}

func (s *LoggingLogResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_logging.LogLifecycleStateDeleting),
	}
}

func (s *LoggingLogResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *LoggingLogResourceCrud) Create() error {
	request := oci_logging.CreateLogRequest{}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", 0)
			tmp, err := s.mapToLogConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Configuration = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	if logType, ok := s.D.GetOkExists("log_type"); ok {
		request.LogType = oci_logging.CreateLogDetailsLogTypeEnum(logType.(string))
	}

	if retentionDuration, ok := s.D.GetOkExists("retention_duration"); ok {
		tmp := retentionDuration.(int)
		request.RetentionDuration = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.CreateLog(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLogFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))

}

func (s *LoggingLogResourceCrud) getLogFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_logging.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	logId, err := logWaitForWorkRequest(workId, "log",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, logId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_logging.DeleteWorkRequestRequest{
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
	s.D.SetId(*logId)

	return s.Get()
}

func logWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "logging", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		/*if workRequestResponse, ok := response.Response.(oci_logging.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}*/
		return false
	}
}

func logWaitForWorkRequest(wId *string, entityType string, action oci_logging.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_logging.LoggingManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "logging")
	retryPolicy.ShouldRetryOperation = logWorkRequestShouldRetryFunc(timeout)

	response := oci_logging.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_logging.OperationStatusInProgress),
			string(oci_logging.OperationStatusAccepted),
			string(oci_logging.OperationStatusCancelling),
		},
		Target: []string{
			string(oci_logging.OperationStatusSucceeded),
			string(oci_logging.OperationStatusFailed),
			string(oci_logging.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_logging.GetWorkRequestRequest{
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
			if res.ActionType == action || res.ActionType == oci_logging.ActionTypesInProgress {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_logging.OperationStatusFailed || response.Status == oci_logging.OperationStatusCanceled {
		return nil, getErrorFromLoggingLogWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromLoggingLogWorkRequest(client *oci_logging.LoggingManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_logging.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_logging.ListWorkRequestErrorsRequest{
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

func (s *LoggingLogResourceCrud) Get() error {
	request := oci_logging.GetLogRequest{}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	tmp := s.D.Id()
	request.LogId = &tmp

	logGroupId, logId, err := parseLogsCompositeId(s.D.Id())
	if err == nil {
		request.LogGroupId = &logGroupId
		request.LogId = &logId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.GetLog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Log
	return nil
}

func (s *LoggingLogResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("log_group_id"); ok && s.D.HasChange("log_group_id") {
		oldRaw, newRaw := s.D.GetChange("log_group_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateLogGroup(oldRaw, newRaw)
			if err != nil {
				return err
			}
		}
	}

	request := oci_logging.UpdateLogRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	tmp := s.D.Id()
	request.LogId = &tmp

	if retentionDuration, ok := s.D.GetOkExists("retention_duration"); ok {
		tmp := retentionDuration.(int)
		request.RetentionDuration = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.UpdateLog(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLogFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *LoggingLogResourceCrud) Delete() error {
	request := oci_logging.DeleteLogRequest{}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	tmp := s.D.Id()
	request.LogId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	_, err := s.Client.DeleteLog(context.Background(), request)
	return err
}

func (s *LoggingLogResourceCrud) SetData() error {
	logGroupId, logId, err := parseLogsCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("log_group_id", &logGroupId)
		s.D.SetId(logId)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Configuration != nil {
		s.D.Set("configuration", []interface{}{LogConfigurationToMap(s.Res.Configuration)})
	} else {
		s.D.Set("configuration", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LogGroupId != nil {
		s.D.Set("log_group_id", *s.Res.LogGroupId)
	}

	s.D.Set("log_type", s.Res.LogType)

	if s.Res.RetentionDuration != nil {
		s.D.Set("retention_duration", *s.Res.RetentionDuration)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func (s *LoggingLogResourceCrud) mapToLogConfiguration(fieldKeyFormat string) (oci_logging.Configuration, error) {
	result := oci_logging.Configuration{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source"), 0)
			tmp, err := s.mapToSource(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source, encountered error: %v", err)
			}
			result.Source = tmp
		}
	}

	return result, nil
}

func LogConfigurationToMap(obj *oci_logging.Configuration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := SourceToMap(&obj.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		result["source"] = sourceArray
	}

	return result
}

func (s *LoggingLogResourceCrud) mapToSource(fieldKeyFormat string) (oci_logging.Source, error) {
	var baseObject oci_logging.Source
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("OCISERVICE"):
		details := oci_logging.OciService{}
		if category, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "category")); ok {
			tmp := category.(string)
			details.Category = &tmp
		}
		if parameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameters")); ok {
			details.Parameters = utils.ObjectMapToStringMap(parameters.(map[string]interface{}))
		}
		if resource_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource")); ok {
			tmp := resource_.(string)
			details.Resource = &tmp
		}
		if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
			tmp := service.(string)
			details.Service = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func SourceToMap(obj *oci_logging.Source) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_logging.OciService:
		result["source_type"] = "OCISERVICE"

		if v.Category != nil {
			result["category"] = string(*v.Category)
		}

		if v.Resource != nil {
			result["resource"] = string(*v.Resource)
		}

		if v.Service != nil {
			result["service"] = string(*v.Service)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *LoggingLogResourceCrud) updateLogGroup(oldLogGroupId interface{}, newLogGroupId interface{}) error {
	updateLogGroupRequest := oci_logging.ChangeLogLogGroupRequest{}

	oldLogGroupIdtmp := oldLogGroupId.(string)
	updateLogGroupRequest.LogGroupId = &oldLogGroupIdtmp

	tmpLogId := s.D.Id()
	updateLogGroupRequest.LogId = &tmpLogId

	updateLogGroupDetails := oci_logging.ChangeLogLogGroupDetails{}
	newLogGroupIdtmp := newLogGroupId.(string)
	updateLogGroupDetails.TargetLogGroupId = &newLogGroupIdtmp

	updateLogGroupRequest.ChangeLogLogGroupDetails = updateLogGroupDetails

	response, err := s.Client.ChangeLogLogGroup(context.Background(), updateLogGroupRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getLogFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesRelated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	return err
}

func GetLogCompositeId(logGroupId string, logId string) string {
	logGroupId = url.PathEscape(logGroupId)
	logId = url.PathEscape(logId)
	compositeId := "logGroupId/" + logGroupId + "/logId/" + logId
	return compositeId
}

func parseLogsCompositeId(compositeId string) (logGroupId string, logId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("logGroupId/.*/logId/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	logGroupId, _ = url.PathUnescape(parts[1])
	logId, _ = url.PathUnescape(parts[3])

	return
}
