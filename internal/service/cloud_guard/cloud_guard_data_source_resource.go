// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func CloudGuardDataSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardDataSource,
		Read:     readCloudGuardDataSource,
		Update:   updateCloudGuardDataSource,
		Delete:   deleteCloudGuardDataSource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_source_feed_provider": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"data_source_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"data_source_feed_provider": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"LOGGINGQUERY",
								"SCHEDULEDQUERY",
							}, true),
						},

						// Optional
						"additional_entities_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interval_in_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interval_in_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"logging_query_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"logging_query_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"INSIGHT",
										}, true),
									},

									// Optional
									"key_entities_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"logging_query_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_start_time": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"start_policy_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"ABSOLUTE_TIME_START_POLICY",
											"NO_DELAY_START_POLICY",
										}, true),
									},

									// Optional
									"query_start_time": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},

									// Computed
								},
							},
						},
						"regions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"scheduled_query_scope_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"region": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"resource_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"resource_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"data_source_detector_mapping_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"detector_recipe_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"detector_rule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"region_status_detail": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createCloudGuardDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardDataSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.DataSource
	DisableNotFoundRetries bool
}

func (s *CloudGuardDataSourceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardDataSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardDataSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardDataSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardDataSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardDataSourceResourceCrud) Create() error {
	request := oci_cloud_guard.CreateDataSourceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataSourceDetails, ok := s.D.GetOkExists("data_source_details"); ok {
		if tmpList := dataSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_source_details", 0)
			tmp, err := s.mapToDataSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataSourceDetails = tmp
		}
	}

	if dataSourceFeedProvider, ok := s.D.GetOkExists("data_source_feed_provider"); ok {
		request.DataSourceFeedProvider = oci_cloud_guard.DataSourceFeedProviderEnum(dataSourceFeedProvider.(string))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_cloud_guard.DataSourceStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_cloud_guard.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_cloud_guard.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getDataSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard"), oci_cloud_guard.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *CloudGuardDataSourceResourceCrud) getDataSourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_guard.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	dataSourceId, err := dataSourceWaitForWorkRequest(workId, "cloudGuardDataSource",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, dataSourceId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_cloud_guard.CancelWorkRequestRequest{
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
	s.D.SetId(*dataSourceId)

	return s.Get()
}

func dataSourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "cloud_guard", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_cloud_guard.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func dataSourceWaitForWorkRequest(wId *string, entityType string, action oci_cloud_guard.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_guard.CloudGuardClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_guard")
	retryPolicy.ShouldRetryOperation = dataSourceWorkRequestShouldRetryFunc(timeout)

	response := oci_cloud_guard.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_cloud_guard.OperationStatusInProgress),
			string(oci_cloud_guard.OperationStatusAccepted),
			string(oci_cloud_guard.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_cloud_guard.OperationStatusSucceeded),
			string(oci_cloud_guard.OperationStatusFailed),
			string(oci_cloud_guard.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_cloud_guard.GetWorkRequestRequest{
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
		identifier = res.Identifier
		break
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_cloud_guard.OperationStatusFailed || response.Status == oci_cloud_guard.OperationStatusCanceled {
		return nil, getErrorFromCloudGuardDataSourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudGuardDataSourceWorkRequest(client *oci_cloud_guard.CloudGuardClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_guard.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_cloud_guard.ListWorkRequestErrorsRequest{
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

func (s *CloudGuardDataSourceResourceCrud) Get() error {
	request := oci_cloud_guard.GetDataSourceRequest{}

	tmp := s.D.Id()
	request.DataSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataSource
	return nil
}

func (s *CloudGuardDataSourceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_guard.UpdateDataSourceRequest{}

	if dataSourceDetails, ok := s.D.GetOkExists("data_source_details"); ok {
		if tmpList := dataSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "data_source_details", 0)
			tmp, err := s.mapToDataSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DataSourceDetails = tmp
		}
	}

	tmp := s.D.Id()
	request.DataSourceId = &tmp

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_cloud_guard.DataSourceStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDataSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard"), oci_cloud_guard.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *CloudGuardDataSourceResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteDataSourceRequest{}

	tmp := s.D.Id()
	request.DataSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.DeleteDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := dataSourceWaitForWorkRequest(workId, "cloud_guard",
		oci_cloud_guard.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *CloudGuardDataSourceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSourceDetails != nil {
		dataSourceDetailsArray := []interface{}{}
		if dataSourceDetailsMap := DataSourceDetailsToMap(&s.Res.DataSourceDetails); dataSourceDetailsMap != nil {
			dataSourceDetailsArray = append(dataSourceDetailsArray, dataSourceDetailsMap)
		}
		s.D.Set("data_source_details", dataSourceDetailsArray)
	} else {
		s.D.Set("data_source_details", nil)
	}

	dataSourceDetectorMappingInfo := []interface{}{}
	for _, item := range s.Res.DataSourceDetectorMappingInfo {
		dataSourceDetectorMappingInfo = append(dataSourceDetectorMappingInfo, DataSourceMappingInfoToMap(item))
	}
	s.D.Set("data_source_detector_mapping_info", dataSourceDetectorMappingInfo)

	s.D.Set("data_source_feed_provider", s.Res.DataSourceFeedProvider)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	regionStatusDetail := []interface{}{}
	for _, item := range s.Res.RegionStatusDetail {
		regionStatusDetail = append(regionStatusDetail, RegionStatusDetailToMap(item))
	}
	s.D.Set("region_status_detail", regionStatusDetail)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

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

func (s *CloudGuardDataSourceResourceCrud) mapToContinuousQueryStartPolicy(fieldKeyFormat string) (oci_cloud_guard.ContinuousQueryStartPolicy, error) {
	var baseObject oci_cloud_guard.ContinuousQueryStartPolicy
	//discriminator
	startPolicyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_policy_type"))
	var startPolicyType string
	if ok {
		startPolicyType = startPolicyTypeRaw.(string)
	} else {
		startPolicyType = "" // default value
	}
	switch strings.ToLower(startPolicyType) {
	case strings.ToLower("ABSOLUTE_TIME_START_POLICY"):
		details := oci_cloud_guard.AbsoluteTimeStartPolicy{}
		if queryStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_start_time")); ok {
			tmp, err := time.Parse(time.RFC3339, queryStartTime.(string))
			if err != nil {
				return details, err
			}
			details.QueryStartTime = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	case strings.ToLower("NO_DELAY_START_POLICY"):
		details := oci_cloud_guard.NoDelayStartPolicy{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown start_policy_type '%v' was specified", startPolicyType)
	}
	return baseObject, nil
}

func ContinuousQueryStartPolicyToMap(obj *oci_cloud_guard.ContinuousQueryStartPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_guard.AbsoluteTimeStartPolicy:
		result["start_policy_type"] = "ABSOLUTE_TIME_START_POLICY"

		if v.QueryStartTime != nil {
			result["query_start_time"] = v.QueryStartTime.Format(time.RFC3339Nano)
		}
	case oci_cloud_guard.NoDelayStartPolicy:
		result["start_policy_type"] = "NO_DELAY_START_POLICY"
	default:
		log.Printf("[WARN] Received 'start_policy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CloudGuardDataSourceResourceCrud) mapToDataSourceDetails(fieldKeyFormat string) (oci_cloud_guard.DataSourceDetails, error) {
	var baseObject oci_cloud_guard.DataSourceDetails
	//discriminator
	dataSourceFeedProviderRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_source_feed_provider"))
	var dataSourceFeedProvider string
	if ok {
		dataSourceFeedProvider = dataSourceFeedProviderRaw.(string)
	} else {
		dataSourceFeedProvider = "" // default value
	}
	switch strings.ToLower(dataSourceFeedProvider) {
	case strings.ToLower("LOGGINGQUERY"):
		details := oci_cloud_guard.LoggingQueryDataSourceDetails{}
		if additionalEntitiesCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_entities_count")); ok {
			tmp := additionalEntitiesCount.(int)
			details.AdditionalEntitiesCount = &tmp
		}
		if intervalInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval_in_minutes")); ok {
			tmp := intervalInMinutes.(int)
			details.IntervalInMinutes = &tmp
		}
		if loggingQueryDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logging_query_details")); ok {
			if tmpList := loggingQueryDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "logging_query_details"), 0)
				tmp, err := s.mapToLoggingQueryDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert logging_query_details, encountered error: %v", err)
				}
				details.LoggingQueryDetails = tmp
			}
		}
		if loggingQueryType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logging_query_type")); ok {
			details.LoggingQueryType = oci_cloud_guard.LoggingQueryTypeEnum(loggingQueryType.(string))
		}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_cloud_guard.LoggingQueryOperatorTypeEnum(operator.(string))
		}
		if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
			tmp := query.(string)
			details.Query = &tmp
		}
		if queryStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_start_time")); ok {
			if tmpList := queryStartTime.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "query_start_time"), 0)
				tmp, err := s.mapToContinuousQueryStartPolicy(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert query_start_time, encountered error: %v", err)
				}
				details.QueryStartTime = tmp
			}
		}
		if regions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "regions")); ok {
			interfaces := regions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "regions")) {
				details.Regions = tmp
			}
		}
		if threshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold")); ok {
			tmp := threshold.(int)
			details.Threshold = &tmp
		}
		baseObject = details
	case strings.ToLower("SCHEDULEDQUERY"):
		details := oci_cloud_guard.ScheduledQueryDataSourceObjDetails{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if intervalInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval_in_seconds")); ok {
			tmp := intervalInSeconds.(int)
			details.IntervalInSeconds = &tmp
		}
		if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
			tmp := query.(string)
			details.Query = &tmp
		}
		if scheduledQueryScopeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scheduled_query_scope_details")); ok {
			interfaces := scheduledQueryScopeDetails.([]interface{})
			tmp := make([]oci_cloud_guard.ScheduledQueryScopeDetail, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scheduled_query_scope_details"), stateDataIndex)
				converted, err := s.mapToScheduledQueryScopeDetail(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "scheduled_query_scope_details")) {
				details.ScheduledQueryScopeDetails = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown data_source_feed_provider '%v' was specified", dataSourceFeedProvider)
	}
	return baseObject, nil
}

func DataSourceDetailsToMap(obj *oci_cloud_guard.DataSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_guard.LoggingQueryDataSourceDetails:
		result["data_source_feed_provider"] = "LOGGINGQUERY"

		if v.AdditionalEntitiesCount != nil {
			result["additional_entities_count"] = int(*v.AdditionalEntitiesCount)
		}

		if v.IntervalInMinutes != nil {
			result["interval_in_minutes"] = int(*v.IntervalInMinutes)
		}

		if v.LoggingQueryDetails != nil {
			loggingQueryDetailsArray := []interface{}{}
			if loggingQueryDetailsMap := LoggingQueryDetailsToMap(&v.LoggingQueryDetails); loggingQueryDetailsMap != nil {
				loggingQueryDetailsArray = append(loggingQueryDetailsArray, loggingQueryDetailsMap)
			}
			result["logging_query_details"] = loggingQueryDetailsArray
		}

		result["logging_query_type"] = string(v.LoggingQueryType)

		result["operator"] = string(v.Operator)

		if v.Query != nil {
			result["query"] = string(*v.Query)
		}

		if v.QueryStartTime != nil {
			queryStartTimeArray := []interface{}{}
			if queryStartTimeMap := ContinuousQueryStartPolicyToMap(&v.QueryStartTime); queryStartTimeMap != nil {
				queryStartTimeArray = append(queryStartTimeArray, queryStartTimeMap)
			}
			result["query_start_time"] = queryStartTimeArray
		}

		result["regions"] = v.Regions

		if v.Threshold != nil {
			result["threshold"] = int(*v.Threshold)
		}
	case oci_cloud_guard.ScheduledQueryDataSourceObjDetails:
		result["data_source_feed_provider"] = "SCHEDULEDQUERY"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.IntervalInSeconds != nil {
			result["interval_in_seconds"] = int(*v.IntervalInSeconds)
		}

		if v.Query != nil {
			result["query"] = string(*v.Query)
		}

		scheduledQueryScopeDetails := []interface{}{}
		for _, item := range v.ScheduledQueryScopeDetails {
			scheduledQueryScopeDetails = append(scheduledQueryScopeDetails, ScheduledQueryScopeDetailToMap(item))
		}
		result["scheduled_query_scope_details"] = scheduledQueryScopeDetails
	default:
		log.Printf("[WARN] Received 'data_source_feed_provider' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CloudGuardDataSourceResourceCrud) mapToDataSourceMappingInfo(fieldKeyFormat string) (oci_cloud_guard.DataSourceMappingInfo, error) {
	result := oci_cloud_guard.DataSourceMappingInfo{}

	if detectorRecipeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detector_recipe_id")); ok {
		tmp := detectorRecipeId.(string)
		result.DetectorRecipeId = &tmp
	}

	if detectorRuleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detector_rule_id")); ok {
		tmp := detectorRuleId.(string)
		result.DetectorRuleId = &tmp
	}

	return result, nil
}

func DataSourceMappingInfoToMap(obj oci_cloud_guard.DataSourceMappingInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DetectorRecipeId != nil {
		result["detector_recipe_id"] = string(*obj.DetectorRecipeId)
	}

	if obj.DetectorRuleId != nil {
		result["detector_rule_id"] = string(*obj.DetectorRuleId)
	}

	return result
}

func DataSourceSummaryToMap(obj oci_cloud_guard.DataSourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["data_source_feed_provider"] = string(obj.DataSourceFeedProvider)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

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

func DataSourceSummaryDetailsToMap(obj *oci_cloud_guard.DataSourceSummaryDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_guard.LoggingQueryDataSourceSummaryDetails:
		result["data_source_feed_provider"] = "LOGGINGQUERY"

		dataSourceDetectorMappingInfo := []interface{}{}
		for _, item := range v.DataSourceDetectorMappingInfo {
			dataSourceDetectorMappingInfo = append(dataSourceDetectorMappingInfo, DataSourceMappingInfoToMap(item))
		}
		result["data_source_detector_mapping_info"] = dataSourceDetectorMappingInfo

		regionStatusDetail := []interface{}{}
		for _, item := range v.RegionStatusDetail {
			regionStatusDetail = append(regionStatusDetail, RegionStatusDetailToMap(item))
		}
		result["region_status_detail"] = regionStatusDetail

		result["regions"] = v.Regions
	case oci_cloud_guard.ScheduledQueryDataSourceSummaryObjDetails:
		result["data_source_feed_provider"] = "SCHEDULEDQUERY"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.IntervalInSeconds != nil {
			result["interval_in_seconds"] = int(*v.IntervalInSeconds)
		}

		regionStatusDetail := []interface{}{}
		for _, item := range v.RegionStatusDetail {
			regionStatusDetail = append(regionStatusDetail, RegionStatusDetailToMap(item))
		}
		result["region_status_detail"] = regionStatusDetail

		scheduledQueryScopeDetails := []interface{}{}
		for _, item := range v.ScheduledQueryScopeDetails {
			scheduledQueryScopeDetails = append(scheduledQueryScopeDetails, ScheduledQueryScopeDetailToMap(item))
		}
		result["scheduled_query_scope_details"] = scheduledQueryScopeDetails
	default:
		log.Printf("[WARN] Received 'data_source_feed_provider' of unknown type %v", *obj)
		return nil
	}

	return result
}

/*
	func (s *CloudGuardDataSourceResourceCrud) mapToLoggingQueryDetails(fieldKeyFormat string) (oci_cloud_guard.LoggingQueryDetails, error) {
		var baseObject oci_cloud_guard.LoggingQueryDetails
		//discriminator
		loggingQueryTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logging_query_type"))
		var loggingQueryType string
		if ok {
			loggingQueryType = loggingQueryTypeRaw.(string)
		} else {
			loggingQueryType = "" // default value
		}
		switch strings.ToLower(loggingQueryType) {
		case strings.ToLower("INSIGHT"):
			details := oci_cloud_guard.InsightTypeLoggingQueryDetails{}
			if keyEntitiesCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_entities_count")); ok {
				tmp := keyEntitiesCount.(int)
				details.KeyEntitiesCount = &tmp
			}
			baseObject = details
		default:
			return nil, fmt.Errorf("unknown logging_query_type '%v' was specified", loggingQueryType)
		}
		return baseObject, nil
	}
*/
func LoggingQueryDetailsToMap(obj *oci_cloud_guard.LoggingQueryDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_guard.InsightTypeLoggingQueryDetails:
		result["logging_query_type"] = "INSIGHT"

		if v.KeyEntitiesCount != nil {
			result["key_entities_count"] = int(*v.KeyEntitiesCount)
		}
	default:
		log.Printf("[WARN] Received 'logging_query_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func RegionStatusDetailToMap(obj oci_cloud_guard.RegionStatusDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["status"] = string(obj.Status)

	return result
}

func (s *CloudGuardDataSourceResourceCrud) mapToScheduledQueryScopeDetail(fieldKeyFormat string) (oci_cloud_guard.ScheduledQueryScopeDetail, error) {
	result := oci_cloud_guard.ScheduledQueryScopeDetail{}

	if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
		tmp := region.(string)
		result.Region = &tmp
	}

	if resourceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_ids")); ok {
		interfaces := resourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "resource_ids")) {
			result.ResourceIds = tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		tmp := resourceType.(string)
		result.ResourceType = &tmp
	}

	return result, nil
}

func ScheduledQueryScopeDetailToMap(obj oci_cloud_guard.ScheduledQueryScopeDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["resource_ids"] = obj.ResourceIds

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	return result
}

func (s *CloudGuardDataSourceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_guard.ChangeDataSourceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DataSourceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.ChangeDataSourceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDataSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard"), oci_cloud_guard.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *CloudGuardDataSourceResourceCrud) mapToLoggingQueryDetails(fieldKeyFormat string) (oci_cloud_guard.LoggingQueryDetails, error) {
	var baseObject oci_cloud_guard.LoggingQueryDetails
	//discriminator
	loggingQueryTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logging_query_type"))
	var loggingQueryType string
	if ok {
		loggingQueryType = loggingQueryTypeRaw.(string)
	} else {
		loggingQueryType = "" // default value
	}
	switch strings.ToLower(loggingQueryType) {
	case strings.ToLower("INSIGHT"):
		details := oci_cloud_guard.InsightTypeLoggingQueryDetails{}
		if keyEntitiesCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_entities_count")); ok {
			tmp := keyEntitiesCount.(int)
			details.KeyEntitiesCount = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown logging_query_type '%v' was specified", loggingQueryType)
	}
	return baseObject, nil
}
