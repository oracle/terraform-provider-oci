// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMeteringComputationSchedule,
		Read:     readMeteringComputationSchedule,
		Update:   updateMeteringComputationSchedule,
		Delete:   deleteMeteringComputationSchedule,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"result_location": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
						},
						"location_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OBJECT_STORAGE",
							}, true),
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"region": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"schedule_recurrences": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_scheduled": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
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
			"output_file_format": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_properties": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"date_range": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"date_range_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DYNAMIC",
											"STATIC",
										}, true),
									},

									// Optional
									"dynamic_date_range_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"time_usage_ended": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"time_usage_started": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},

									// Computed
								},
							},
						},
						"granularity": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"compartment_depth": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"filter": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"group_by": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"group_by_tag": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"is_aggregate_by_time": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"query_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"saved_report_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
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
			"time_next_run": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMeteringComputationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.CreateResource(d, sync)
}

func readMeteringComputationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

func updateMeteringComputationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMeteringComputationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MeteringComputationScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_metering_computation.UsageapiClient
	Res                    *oci_metering_computation.Schedule
	DisableNotFoundRetries bool
}

func (s *MeteringComputationScheduleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MeteringComputationScheduleResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MeteringComputationScheduleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_metering_computation.ScheduleLifecycleStateActive),
	}
}

func (s *MeteringComputationScheduleResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MeteringComputationScheduleResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *MeteringComputationScheduleResourceCrud) Create() error {
	request := oci_metering_computation.CreateScheduleRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if outputFileFormat, ok := s.D.GetOkExists("output_file_format"); ok {
		request.OutputFileFormat = oci_metering_computation.CreateScheduleDetailsOutputFileFormatEnum(outputFileFormat.(string))
	}

	if queryProperties, ok := s.D.GetOkExists("query_properties"); ok {
		if tmpList := queryProperties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_properties", 0)
			tmp, err := s.mapToQueryProperties(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryProperties = &tmp
		}
	}

	if resultLocation, ok := s.D.GetOkExists("result_location"); ok {
		if tmpList := resultLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "result_location", 0)
			tmp, err := s.mapToResultLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResultLocation = tmp
		}
	}

	if savedReportId, ok := s.D.GetOkExists("saved_report_id"); ok {
		tmp := savedReportId.(string)
		request.SavedReportId = &tmp
	}

	if scheduleRecurrences, ok := s.D.GetOkExists("schedule_recurrences"); ok {
		tmp := scheduleRecurrences.(string)
		request.ScheduleRecurrences = &tmp
	}

	if timeScheduled, ok := s.D.GetOkExists("time_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeScheduled = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.CreateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *MeteringComputationScheduleResourceCrud) Get() error {
	request := oci_metering_computation.GetScheduleRequest{}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.GetSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *MeteringComputationScheduleResourceCrud) Update() error {
	request := oci_metering_computation.UpdateScheduleRequest{}

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

	if outputFileFormat, ok := s.D.GetOkExists("output_file_format"); ok {
		request.OutputFileFormat = oci_metering_computation.UpdateScheduleDetailsOutputFileFormatEnum(outputFileFormat.(string))
	}

	if resultLocation, ok := s.D.GetOkExists("result_location"); ok {
		if tmpList := resultLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "result_location", 0)
			tmp, err := s.mapToResultLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResultLocation = tmp
		}
	}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.UpdateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *MeteringComputationScheduleResourceCrud) Delete() error {
	request := oci_metering_computation.DeleteScheduleRequest{}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	_, err := s.Client.DeleteSchedule(context.Background(), request)
	return err
}

func (s *MeteringComputationScheduleResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("output_file_format", s.Res.OutputFileFormat)

	if s.Res.QueryProperties != nil {
		s.D.Set("query_properties", []interface{}{QueryPropertiesToMap(s.Res.QueryProperties)})
	} else {
		s.D.Set("query_properties", nil)
	}

	if s.Res.ResultLocation != nil {
		resultLocationArray := []interface{}{}
		if resultLocationMap := ResultLocationToMap(&s.Res.ResultLocation); resultLocationMap != nil {
			resultLocationArray = append(resultLocationArray, resultLocationMap)
		}
		s.D.Set("result_location", resultLocationArray)
	} else {
		s.D.Set("result_location", nil)
	}

	if s.Res.SavedReportId != nil {
		s.D.Set("saved_report_id", *s.Res.SavedReportId)
	}

	if s.Res.ScheduleRecurrences != nil {
		s.D.Set("schedule_recurrences", *s.Res.ScheduleRecurrences)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNextRun != nil {
		s.D.Set("time_next_run", s.Res.TimeNextRun.String())
	}

	if s.Res.TimeScheduled != nil {
		s.D.Set("time_scheduled", s.Res.TimeScheduled.Format(time.RFC3339Nano))
	}

	return nil
}

func (s *MeteringComputationScheduleResourceCrud) mapToDateRange(fieldKeyFormat string) (oci_metering_computation.DateRange, error) {
	var baseObject oci_metering_computation.DateRange
	//discriminator
	dateRangeTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date_range_type"))
	var dateRangeType string
	if ok {
		dateRangeType = dateRangeTypeRaw.(string)
	} else {
		dateRangeType = "" // default value
	}
	switch strings.ToLower(dateRangeType) {
	case strings.ToLower("DYNAMIC"):
		details := oci_metering_computation.DynamicDateRange{}
		if dynamicDateRangeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dynamic_date_range_type")); ok {
			details.DynamicDateRangeType = oci_metering_computation.DynamicDateRangeDynamicDateRangeTypeEnum(dynamicDateRangeType.(string))
		}
		baseObject = details
	case strings.ToLower("STATIC"):
		details := oci_metering_computation.StaticDateRange{}
		if timeUsageEnded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_usage_ended")); ok {
			tmp, err := time.Parse(time.RFC3339, timeUsageEnded.(string))
			if err != nil {
				return details, err
			}
			details.TimeUsageEnded = &oci_common.SDKTime{Time: tmp}
		}
		if timeUsageStarted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_usage_started")); ok {
			tmp, err := time.Parse(time.RFC3339, timeUsageStarted.(string))
			if err != nil {
				return details, err
			}
			details.TimeUsageStarted = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown date_range_type '%v' was specified", dateRangeType)
	}
	return baseObject, nil
}

func DateRangeToMap(obj *oci_metering_computation.DateRange) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_metering_computation.DynamicDateRange:
		result["date_range_type"] = "DYNAMIC"

		result["dynamic_date_range_type"] = string(v.DynamicDateRangeType)
	case oci_metering_computation.StaticDateRange:
		result["date_range_type"] = "STATIC"

		if v.TimeUsageEnded != nil {
			result["time_usage_ended"] = v.TimeUsageEnded.Format(time.RFC3339Nano)
		}

		if v.TimeUsageStarted != nil {
			result["time_usage_started"] = v.TimeUsageStarted.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'date_range_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MeteringComputationScheduleResourceCrud) mapToQueryProperties(fieldKeyFormat string) (oci_metering_computation.QueryProperties, error) {
	result := oci_metering_computation.QueryProperties{}

	if compartmentDepth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_depth")); ok {
		tmp := compartmentDepth.(float32)
		result.CompartmentDepth = &tmp
	}

	if dateRange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date_range")); ok {
		if tmpList := dateRange.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "date_range"), 0)
			tmp, err := s.mapToDateRange(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert date_range, encountered error: %v", err)
			}
			result.DateRange = tmp
		}
	}

	if filter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter")); ok {
		tmp := filter.(string)
		var filterObj oci_metering_computation.Filter
		err := json.Unmarshal([]byte(tmp), &filterObj)
		if err != nil {
			return result, fmt.Errorf("[ERROR ]encountered error: %v", err.Error())
		}
		result.Filter = &filterObj
	}

	if granularity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "granularity")); ok {
		result.Granularity = oci_metering_computation.QueryPropertiesGranularityEnum(granularity.(string))
	}

	if groupBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_by")); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_by")) {
			result.GroupBy = tmp
		}
	}

	if groupByTag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_by_tag")); ok {
		interfaces := groupByTag.([]interface{})
		tmp := make([]oci_metering_computation.Tag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "group_by_tag"), stateDataIndex)
			converted, err := s.mapToTag(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_by_tag")) {
			result.GroupByTag = tmp
		}
	}

	if isAggregateByTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_aggregate_by_time")); ok {
		tmp := isAggregateByTime.(bool)
		result.IsAggregateByTime = &tmp
	}

	if queryType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_type")); ok {
		result.QueryType = oci_metering_computation.QueryPropertiesQueryTypeEnum(queryType.(string))
	}

	return result, nil
}

func QueryPropertiesToMap(obj *oci_metering_computation.QueryProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentDepth != nil {
		result["compartment_depth"] = float32(*obj.CompartmentDepth)
	}

	if obj.DateRange != nil {
		dateRangeArray := []interface{}{}
		if dateRangeMap := DateRangeToMap(&obj.DateRange); dateRangeMap != nil {
			dateRangeArray = append(dateRangeArray, dateRangeMap)
		}
		result["date_range"] = dateRangeArray
	}

	if obj.Filter != nil {
		result["filter"] = obj.Filter
	}

	result["granularity"] = string(obj.Granularity)

	result["group_by"] = obj.GroupBy

	groupByTag := []interface{}{}
	for _, item := range obj.GroupByTag {
		groupByTag = append(groupByTag, TagToMap(item))
	}
	result["group_by_tag"] = groupByTag

	if obj.IsAggregateByTime != nil {
		result["is_aggregate_by_time"] = bool(*obj.IsAggregateByTime)
	}

	result["query_type"] = string(obj.QueryType)

	return result
}

func (s *MeteringComputationScheduleResourceCrud) mapToResultLocation(fieldKeyFormat string) (oci_metering_computation.ResultLocation, error) {
	var baseObject oci_metering_computation.ResultLocation
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_metering_computation.ObjectStorageLocation{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func ResultLocationToMap(obj *oci_metering_computation.ResultLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_metering_computation.ObjectStorageLocation:
		result["location_type"] = "OBJECT_STORAGE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ScheduleSummaryToMap(obj oci_metering_computation.ScheduleSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ScheduleRecurrences != nil {
		result["schedule_recurrences"] = string(*obj.ScheduleRecurrences)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeNextRun != nil {
		result["time_next_run"] = obj.TimeNextRun.String()
	}

	if obj.TimeScheduled != nil {
		result["time_scheduled"] = obj.TimeScheduled.String()
	}

	return result
}

func (s *MeteringComputationScheduleResourceCrud) mapToTag(fieldKeyFormat string) (oci_metering_computation.Tag, error) {
	result := oci_metering_computation.Tag{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func TagToMap(obj oci_metering_computation.Tag) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
