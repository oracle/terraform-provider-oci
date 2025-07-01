// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesScheduledQueryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmTracesScheduledQuery,
		Read:     readApmTracesScheduledQuery,
		Update:   updateApmTracesScheduledQuery,
		Delete:   deleteApmTracesScheduledQuery,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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
			"opc_dry_run": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_query_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_query_maximum_runtime_in_seconds": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"scheduled_query_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_query_processing_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"custom_metric": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"compartment": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_anomaly_detection_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"is_metric_published": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"resource_group": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unit": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"object_storage": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name_space": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"object_name_prefix": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"streaming": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"stream_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"scheduled_query_processing_sub_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_query_processing_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_query_retention_criteria": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_query_retention_period_in_ms": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"scheduled_query_schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_query_text": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"scheduled_query_instances": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_query_next_run_in_ms": {
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
		},
	}
}

func createApmTracesScheduledQuery(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesScheduledQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledQueryClient()

	return tfresource.CreateResource(d, sync)
}

func readApmTracesScheduledQuery(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesScheduledQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledQueryClient()

	return tfresource.ReadResource(sync)
}

func updateApmTracesScheduledQuery(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesScheduledQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledQueryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmTracesScheduledQuery(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesScheduledQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledQueryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmTracesScheduledQueryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_traces.ScheduledQueryClient
	Res                    *oci_apm_traces.ScheduledQuery
	DisableNotFoundRetries bool
}

func (s *ApmTracesScheduledQueryResourceCrud) ID() string {
	return GetScheduledQueryCompositeId(s.D.Get("apm_domain_id").(string), *s.Res.Id)
}

func (s *ApmTracesScheduledQueryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apm_traces.LifecycleStatesCreating),
	}
}

func (s *ApmTracesScheduledQueryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apm_traces.LifecycleStatesActive),
	}
}

func (s *ApmTracesScheduledQueryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apm_traces.LifecycleStatesDeleting),
	}
}

func (s *ApmTracesScheduledQueryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apm_traces.LifecycleStatesDeleted),
	}
}

func (s *ApmTracesScheduledQueryResourceCrud) Create() error {
	request := oci_apm_traces.CreateScheduledQueryRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
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

	if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
		tmp := opcDryRun.(string)
		request.OpcDryRun = &tmp
	}

	if scheduledQueryDescription, ok := s.D.GetOkExists("scheduled_query_description"); ok {
		tmp := scheduledQueryDescription.(string)
		request.ScheduledQueryDescription = &tmp
	}

	if scheduledQueryMaximumRuntimeInSeconds, ok := s.D.GetOkExists("scheduled_query_maximum_runtime_in_seconds"); ok {
		tmp := scheduledQueryMaximumRuntimeInSeconds.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert scheduledQueryMaximumRuntimeInSeconds string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ScheduledQueryMaximumRuntimeInSeconds = &tmpInt64
	}

	if scheduledQueryName, ok := s.D.GetOkExists("scheduled_query_name"); ok {
		tmp := scheduledQueryName.(string)
		request.ScheduledQueryName = &tmp
	}

	if scheduledQueryProcessingConfiguration, ok := s.D.GetOkExists("scheduled_query_processing_configuration"); ok {
		if tmpList := scheduledQueryProcessingConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_query_processing_configuration", 0)
			tmp, err := s.mapToScheduledQueryProcessingConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScheduledQueryProcessingConfiguration = &tmp
		}
	}

	if scheduledQueryProcessingSubType, ok := s.D.GetOkExists("scheduled_query_processing_sub_type"); ok {
		request.ScheduledQueryProcessingSubType = oci_apm_traces.ScheduledQueryProcessingSubTypeEnum(scheduledQueryProcessingSubType.(string))
	}

	if scheduledQueryProcessingType, ok := s.D.GetOkExists("scheduled_query_processing_type"); ok {
		request.ScheduledQueryProcessingType = oci_apm_traces.ScheduledQueryProcessingTypeEnum(scheduledQueryProcessingType.(string))
	}

	if scheduledQueryRetentionCriteria, ok := s.D.GetOkExists("scheduled_query_retention_criteria"); ok {
		request.ScheduledQueryRetentionCriteria = oci_apm_traces.ScheduledQueryRetentionCriteriaEnum(scheduledQueryRetentionCriteria.(string))
	}

	if scheduledQueryRetentionPeriodInMs, ok := s.D.GetOkExists("scheduled_query_retention_period_in_ms"); ok {
		tmp := scheduledQueryRetentionPeriodInMs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert scheduledQueryRetentionPeriodInMs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ScheduledQueryRetentionPeriodInMs = &tmpInt64
	}

	if scheduledQuerySchedule, ok := s.D.GetOkExists("scheduled_query_schedule"); ok {
		tmp := scheduledQuerySchedule.(string)
		request.ScheduledQuerySchedule = &tmp
	}

	if scheduledQueryText, ok := s.D.GetOkExists("scheduled_query_text"); ok {
		tmp := scheduledQueryText.(string)
		request.ScheduledQueryText = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_traces")

	response, err := s.Client.CreateScheduledQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledQuery
	return nil
}

func (s *ApmTracesScheduledQueryResourceCrud) Get() error {

	request := oci_apm_traces.GetScheduledQueryRequest{}

	apmDomainId, scheduledQueryId, err := parseScheduledQueryCompositeId(s.D.Id())
	if err == nil {
		request.ScheduledQueryId = &scheduledQueryId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_traces")

	response, err := s.Client.GetScheduledQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledQuery
	return nil
}

func (s *ApmTracesScheduledQueryResourceCrud) Update() error {
	request := oci_apm_traces.UpdateScheduledQueryRequest{}

	apmDomainId, scheduledQueryId, err := parseScheduledQueryCompositeId(s.D.Id())
	if err == nil {
		request.ScheduledQueryId = &scheduledQueryId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
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

	if opcDryRun, ok := s.D.GetOkExists("opc_dry_run"); ok {
		tmp := opcDryRun.(string)
		request.OpcDryRun = &tmp
	}

	if scheduledQueryDescription, ok := s.D.GetOkExists("scheduled_query_description"); ok {
		tmp := scheduledQueryDescription.(string)
		request.ScheduledQueryDescription = &tmp
	}

	/*tmp := s.D.Id()
	request.ScheduledQueryId = &tmp*/

	if scheduledQueryMaximumRuntimeInSeconds, ok := s.D.GetOkExists("scheduled_query_maximum_runtime_in_seconds"); ok {
		tmp := scheduledQueryMaximumRuntimeInSeconds.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert scheduledQueryMaximumRuntimeInSeconds string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ScheduledQueryMaximumRuntimeInSeconds = &tmpInt64
	}

	if scheduledQueryName, ok := s.D.GetOkExists("scheduled_query_name"); ok {
		tmp := scheduledQueryName.(string)
		request.ScheduledQueryName = &tmp
	}

	/*if scheduledQueryProcessingConfiguration, ok := s.D.GetOkExists("scheduled_query_processing_configuration"); ok {
		if tmpList := scheduledQueryProcessingConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scheduled_query_processing_configuration", 0)
			tmp, err := s.mapToScheduledQueryProcessingConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScheduledQueryProcessingConfiguration = &tmp
		}
	}

	if scheduledQueryProcessingSubType, ok := s.D.GetOkExists("scheduled_query_processing_sub_type"); ok {
		request.ScheduledQueryProcessingSubType = oci_apm_traces.ScheduledQueryProcessingSubTypeEnum(scheduledQueryProcessingSubType.(string))
	}

	if scheduledQueryProcessingType, ok := s.D.GetOkExists("scheduled_query_processing_type"); ok {
		request.ScheduledQueryProcessingType = oci_apm_traces.ScheduledQueryProcessingTypeEnum(scheduledQueryProcessingType.(string))
	}

	if scheduledQueryRetentionCriteria, ok := s.D.GetOkExists("scheduled_query_retention_criteria"); ok {
		request.ScheduledQueryRetentionCriteria = oci_apm_traces.ScheduledQueryRetentionCriteriaEnum(scheduledQueryRetentionCriteria.(string))
	}

	if scheduledQueryRetentionPeriodInMs, ok := s.D.GetOkExists("scheduled_query_retention_period_in_ms"); ok {
		tmp := scheduledQueryRetentionPeriodInMs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert scheduledQueryRetentionPeriodInMs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ScheduledQueryRetentionPeriodInMs = &tmpInt64
	}*/

	/*if scheduledQuerySchedule, ok := s.D.GetOkExists("scheduled_query_schedule"); ok {
		tmp := scheduledQuerySchedule.(string)
		request.ScheduledQuerySchedule = &tmp
	}*/

	/*if scheduledQueryText, ok := s.D.GetOkExists("scheduled_query_text"); ok {
		tmp := scheduledQueryText.(string)
		request.ScheduledQueryText = &tmp
	}*/

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_traces")

	response, err := s.Client.UpdateScheduledQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledQuery
	return nil
}

func (s *ApmTracesScheduledQueryResourceCrud) Delete() error {
	request := oci_apm_traces.DeleteScheduledQueryRequest{}

	if tmp := s.D.Id(); tmp != "" {
		apmDomainId, scheduledQueryId, err := parseScheduledQueryCompositeId(s.D.Id())
		if err == nil {
			request.ScheduledQueryId = &scheduledQueryId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_traces")

	_, err := s.Client.DeleteScheduledQuery(context.Background(), request)
	return err
}

func (s *ApmTracesScheduledQueryResourceCrud) SetData() error {

	apmDomainId, scheduledQueryId, err := parseScheduledQueryCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s and apmDomain ID: %s", scheduledQueryId, apmDomainId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ScheduledQueryDescription != nil {
		s.D.Set("scheduled_query_description", *s.Res.ScheduledQueryDescription)
	}

	if s.Res.ScheduledQueryInstances != nil {
		s.D.Set("scheduled_query_instances", *s.Res.ScheduledQueryInstances)
	}

	if s.Res.ScheduledQueryMaximumRuntimeInSeconds != nil {
		s.D.Set("scheduled_query_maximum_runtime_in_seconds", strconv.FormatInt(*s.Res.ScheduledQueryMaximumRuntimeInSeconds, 10))
	}

	if s.Res.ScheduledQueryName != nil {
		s.D.Set("scheduled_query_name", *s.Res.ScheduledQueryName)
	}

	if s.Res.ScheduledQueryNextRunInMs != nil {
		s.D.Set("scheduled_query_next_run_in_ms", strconv.FormatInt(*s.Res.ScheduledQueryNextRunInMs, 10))
	}

	if s.Res.ScheduledQueryProcessingConfiguration != nil {
		s.D.Set("scheduled_query_processing_configuration", []interface{}{ScheduledQueryProcessingConfigToMap(s.Res.ScheduledQueryProcessingConfiguration)})
	} else {
		s.D.Set("scheduled_query_processing_configuration", nil)
	}

	s.D.Set("scheduled_query_processing_sub_type", s.Res.ScheduledQueryProcessingSubType)

	s.D.Set("scheduled_query_processing_type", s.Res.ScheduledQueryProcessingType)

	s.D.Set("scheduled_query_retention_criteria", s.Res.ScheduledQueryRetentionCriteria)

	if s.Res.ScheduledQueryRetentionPeriodInMs != nil {
		s.D.Set("scheduled_query_retention_period_in_ms", strconv.FormatInt(*s.Res.ScheduledQueryRetentionPeriodInMs, 10))
	}

	if s.Res.ScheduledQuerySchedule != nil {
		s.D.Set("scheduled_query_schedule", *s.Res.ScheduledQuerySchedule)
	}

	if s.Res.ScheduledQueryText != nil {
		s.D.Set("scheduled_query_text", *s.Res.ScheduledQueryText)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	return nil
}

func GetScheduledQueryCompositeId(apmDomainId string, scheduledQueryId string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	scheduledQueryId = url.PathEscape(scheduledQueryId)
	compositeId := "scheduledQueries/" + scheduledQueryId + "/apmDomainId/" + apmDomainId
	return compositeId
}

func parseScheduledQueryCompositeId(compositeId string) (apmDomainId string, scheduledQueryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("scheduledQueries/.*/apmDomainId/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	scheduledQueryId, _ = url.PathUnescape(parts[1])
	apmDomainId, _ = url.PathUnescape(parts[3])

	return
}

func (s *ApmTracesScheduledQueryResourceCrud) mapToCustomMetric(fieldKeyFormat string) (oci_apm_traces.CustomMetric, error) {
	result := oci_apm_traces.CustomMetric{}

	if compartment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment")); ok {
		tmp := compartment.(string)
		result.Compartment = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if isAnomalyDetectionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_anomaly_detection_enabled")); ok {
		tmp := isAnomalyDetectionEnabled.(bool)
		result.IsAnomalyDetectionEnabled = &tmp
	}

	if isMetricPublished, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_metric_published")); ok {
		tmp := isMetricPublished.(bool)
		result.IsMetricPublished = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if resourceGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_group")); ok {
		tmp := resourceGroup.(string)
		result.ResourceGroup = &tmp
	}

	if unit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unit")); ok {
		tmp := unit.(string)
		result.Unit = &tmp
	}

	return result, nil
}

func CustomMetricToMap(obj *oci_apm_traces.CustomMetric) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Compartment != nil {
		result["compartment"] = string(*obj.Compartment)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsAnomalyDetectionEnabled != nil {
		result["is_anomaly_detection_enabled"] = bool(*obj.IsAnomalyDetectionEnabled)
	}

	if obj.IsMetricPublished != nil {
		result["is_metric_published"] = bool(*obj.IsMetricPublished)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ResourceGroup != nil {
		result["resource_group"] = string(*obj.ResourceGroup)
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	return result
}

func (s *ApmTracesScheduledQueryResourceCrud) mapToObjectStorage(fieldKeyFormat string) (oci_apm_traces.ObjectStorage, error) {
	result := oci_apm_traces.ObjectStorage{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if nameSpace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name_space")); ok {
		tmp := nameSpace.(string)
		result.NameSpace = &tmp
	}

	if objectNamePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_name_prefix")); ok {
		tmp := objectNamePrefix.(string)
		result.ObjectNamePrefix = &tmp
	}

	return result, nil
}

func ObjectStorageToMap(obj *oci_apm_traces.ObjectStorage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.NameSpace != nil {
		result["name_space"] = string(*obj.NameSpace)
	}

	if obj.ObjectNamePrefix != nil {
		result["object_name_prefix"] = string(*obj.ObjectNamePrefix)
	}

	return result
}

func (s *ApmTracesScheduledQueryResourceCrud) mapToScheduledQueryProcessingConfig(fieldKeyFormat string) (oci_apm_traces.ScheduledQueryProcessingConfig, error) {
	result := oci_apm_traces.ScheduledQueryProcessingConfig{}

	if customMetric, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_metric")); ok {
		if tmpList := customMetric.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "custom_metric"), 0)
			tmp, err := s.mapToCustomMetric(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert custom_metric, encountered error: %v", err)
			}
			result.CustomMetric = &tmp
		}
	}

	if objectStorage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_storage")); ok {
		if tmpList := objectStorage.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_storage"), 0)
			tmp, err := s.mapToObjectStorage(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert object_storage, encountered error: %v", err)
			}
			result.ObjectStorage = &tmp
		}
	}

	if streaming, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "streaming")); ok {
		if tmpList := streaming.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "streaming"), 0)
			tmp, err := s.mapToStreaming(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert streaming, encountered error: %v", err)
			}
			result.Streaming = &tmp
		}
	}

	return result, nil
}

func ScheduledQueryProcessingConfigToMap(obj *oci_apm_traces.ScheduledQueryProcessingConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomMetric != nil {
		result["custom_metric"] = []interface{}{CustomMetricToMap(obj.CustomMetric)}
	}

	if obj.ObjectStorage != nil {
		result["object_storage"] = []interface{}{ObjectStorageToMap(obj.ObjectStorage)}
	}

	if obj.Streaming != nil {
		result["streaming"] = []interface{}{StreamingToMap(obj.Streaming)}
	}

	return result
}

func ScheduledQuerySummaryToMap(obj oci_apm_traces.ScheduledQuerySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ScheduledQueryInstances != nil {
		result["scheduled_query_instances"] = string(*obj.ScheduledQueryInstances)
	}

	if obj.ScheduledQueryName != nil {
		result["scheduled_query_name"] = string(*obj.ScheduledQueryName)
	}

	if obj.ScheduledQueryNextRunInMs != nil {
		result["scheduled_query_next_run_in_ms"] = strconv.FormatInt(*obj.ScheduledQueryNextRunInMs, 10)
	}

	if obj.ScheduledQueryProcessingConfiguration != nil {
		result["scheduled_query_processing_configuration"] = []interface{}{ScheduledQueryProcessingConfigToMap(obj.ScheduledQueryProcessingConfiguration)}
	}

	result["scheduled_query_processing_sub_type"] = string(obj.ScheduledQueryProcessingSubType)

	result["scheduled_query_processing_type"] = string(obj.ScheduledQueryProcessingType)

	result["scheduled_query_retention_criteria"] = string(obj.ScheduledQueryRetentionCriteria)

	if obj.ScheduledQuerySchedule != nil {
		result["scheduled_query_schedule"] = string(*obj.ScheduledQuerySchedule)
	}

	if obj.ScheduledQueryText != nil {
		result["scheduled_query_text"] = string(*obj.ScheduledQueryText)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}

func (s *ApmTracesScheduledQueryResourceCrud) mapToStreaming(fieldKeyFormat string) (oci_apm_traces.Streaming, error) {
	result := oci_apm_traces.Streaming{}

	if streamId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_id")); ok {
		tmp := streamId.(string)
		result.StreamId = &tmp
	}

	return result, nil
}

func StreamingToMap(obj *oci_apm_traces.Streaming) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.StreamId != nil {
		result["stream_id"] = string(*obj.StreamId)
	}

	return result
}
