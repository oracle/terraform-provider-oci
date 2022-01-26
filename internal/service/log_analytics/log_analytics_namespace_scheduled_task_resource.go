// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v56/loganalytics"
)

func LogAnalyticsNamespaceScheduledTaskResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsNamespaceScheduledTask,
		Read:     readLogAnalyticsNamespaceScheduledTask,
		Update:   updateLogAnalyticsNamespaceScheduledTask,
		Delete:   deleteLogAnalyticsNamespaceScheduledTask,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kind": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ACCELERATION",
					"STANDARD",
				}, true),
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"task_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PURGE",
								"STREAM",
							}, true),
						},

						// Optional
						"compartment_id_in_subtree": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"data_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"purge_compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"purge_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"query_string": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"saved_search_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"schedules": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"schedule": {
							Type:     schema.TypeSet,
							Required: true,
							Set:      scheduleHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CRON",
											"FIXED_FREQUENCY",
										}, true),
									},

									// Optional
									"expression": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"misfire_policy": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"recurring_interval": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"repeat_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"time_zone": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
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
			"saved_search_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"scheduled_task_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"num_occurrences": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_status": {
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
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsNamespaceScheduledTask(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceScheduledTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceScheduledTask(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceScheduledTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsNamespaceScheduledTask(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceScheduledTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsNamespaceScheduledTask(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceScheduledTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsNamespaceScheduledTaskResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.ScheduledTask
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) ID() string {
	var namespace, scheduledTaskId string

	if tmp, ok := s.D.GetOkExists("namespace"); ok {
		namespace = tmp.(string)
	}

	if tmp, ok := s.D.GetOkExists("scheduled_task_id"); ok {
		scheduledTaskId = tmp.(string)
	} else if s.Res != nil {
		scheduledTaskId = *((*(s.Res)).GetId())
		s.D.Set("scheduled_task_id", scheduledTaskId)
	}

	return GetNamespaceScheduledTaskCompositeId(namespace, scheduledTaskId)
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_log_analytics.ScheduledTaskLifecycleStateActive),
	}
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_log_analytics.ScheduledTaskLifecycleStateDeleted),
	}
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) Create() error {
	request := oci_log_analytics.CreateScheduledTaskRequest{}
	err := s.populateTopLevelPolymorphicCreateScheduledTaskRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.CreateScheduledTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledTask
	return nil
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) Get() error {
	request := oci_log_analytics.GetScheduledTaskRequest{}

	namespace, scheduledTaskId, err := parseNamespaceScheduledTaskCompositeId(s.D.Id())
	if err == nil {
		request.NamespaceName = &namespace
		request.ScheduledTaskId = &scheduledTaskId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetScheduledTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledTask
	return nil
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	request := oci_log_analytics.UpdateScheduledTaskRequest{}

	// TODO: fix this when ACCELERATED tasks are supported
	// var details oci_log_analytics.UpdateScheduledTaskDetails
	var details oci_log_analytics.UpdateStandardTaskDetails

	if kind, ok := s.D.GetOkExists("kind"); ok {
		tmp := kind.(string)
		if strings.ToLower(tmp) == strings.ToLower("STANDARD") {
			details = oci_log_analytics.UpdateStandardTaskDetails{}
		} else {
			return fmt.Errorf("unsupported kind %s found in UpdateScheduledTaskRequest", kind.(string))
		}
	} else {
		return fmt.Errorf("unknown kind %s found in UpdateScheduledTaskRequest", kind.(string))
	}

	namespace, scheduledTaskId, err := parseNamespaceScheduledTaskCompositeId(s.D.Id())
	if err == nil {
		request.NamespaceName = &namespace
		request.ScheduledTaskId = &scheduledTaskId
	} else {
		log.Printf("[WARN] Update() unable to parse current ID: %s", s.D.Id())
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		details.DisplayName = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		details.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if schedules, ok := s.D.GetOkExists("schedules"); ok {
		if tmpList := schedules.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", 0)
			tmp, err := s.mapToSchedules(fieldKeyFormat)
			if err != nil {
				return err
			}
			if len(tmp) != 0 || s.D.HasChange("schedules") {
				details.Schedules = tmp
			}
		}
	}

	request.UpdateScheduledTaskDetails = details
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateScheduledTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledTask
	return nil
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) Delete() error {
	request := oci_log_analytics.DeleteScheduledTaskRequest{}

	namespace, scheduledTaskId, err := parseNamespaceScheduledTaskCompositeId(s.D.Id())
	if err == nil {
		request.NamespaceName = &namespace
		request.ScheduledTaskId = &scheduledTaskId
	} else {
		log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err = s.Client.DeleteScheduledTask(context.Background(), request)
	return err
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) SetData() error {
	if s.Res == nil || *(s.Res) == nil {
		return nil
	}

	result := *(s.Res)

	if result.GetAction() != nil {
		actionArray := []interface{}{}
		if actionMap := LAActionToMap(result.GetAction()); actionMap != nil {
			actionArray = append(actionArray, actionMap)
		}
		s.D.Set("action", actionArray)
	} else {
		s.D.Set("action", nil)
	}

	if result.GetCompartmentId() != nil {
		s.D.Set("compartment_id", result.GetCompartmentId())
	}

	if result.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(result.GetDefinedTags()))
	}

	if result.GetDisplayName() != nil {
		s.D.Set("display_name", result.GetDisplayName())
	}

	s.D.Set("freeform_tags", result.GetFreeformTags())

	if result.GetNumOccurrences() != nil {
		s.D.Set("num_occurrences", strconv.FormatInt(*result.GetNumOccurrences(), 10))
	}

	if result.GetSchedules() != nil {
		s.D.Set("schedules", []interface{}{ScheduleListToMap(result.GetSchedules(), false)})
	} else {
		s.D.Set("schedules", nil)
	}

	s.D.Set("state", result.GetLifecycleState())

	s.D.Set("task_status", result.GetTaskStatus())

	s.D.Set("task_type", result.GetTaskType())

	s.D.Set("scheduled_task_id", result.GetId())

	if result.GetTimeCreated() != nil {
		s.D.Set("time_created", result.GetTimeCreated().String())
	}

	if result.GetTimeUpdated() != nil {
		s.D.Set("time_updated", result.GetTimeUpdated().String())
	}

	if result.GetWorkRequestId() != nil {
		s.D.Set("work_request_id", result.GetWorkRequestId())
	}

	return nil
}

func GetNamespaceScheduledTaskCompositeId(namespace string, scheduledTaskId string) string {
	namespace = url.PathEscape(namespace)
	scheduledTaskId = url.PathEscape(scheduledTaskId)
	compositeId := "namespaces/" + namespace + "/scheduledTasks/" + scheduledTaskId
	return compositeId
}

func parseNamespaceScheduledTaskCompositeId(compositeId string) (namespace string, scheduledTaskId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/scheduledTasks/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	scheduledTaskId, _ = url.PathUnescape(parts[3])

	return
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) mapToAction(fieldKeyFormat string) (oci_log_analytics.Action, error) {
	var baseObject oci_log_analytics.Action
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("PURGE"):
		details := oci_log_analytics.PurgeAction{}
		if compartmentIdInSubtree, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id_in_subtree")); ok {
			tmp := compartmentIdInSubtree.(bool)
			details.CompartmentIdInSubtree = &tmp
		}
		if dataType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_type")); ok {
			details.DataType = oci_log_analytics.StorageDataTypeEnum(dataType.(string))
		}
		if purgeCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "purge_compartment_id")); ok {
			tmp := purgeCompartmentId.(string)
			details.PurgeCompartmentId = &tmp
		}
		if purgeDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "purge_duration")); ok {
			tmp := purgeDuration.(string)
			details.PurgeDuration = &tmp
		}
		if queryString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_string")); ok {
			tmp := queryString.(string)
			details.QueryString = &tmp
		}
		baseObject = details
	case strings.ToLower("STREAM"):
		details := oci_log_analytics.StreamAction{}
		if savedSearchId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "saved_search_id")); ok {
			tmp := savedSearchId.(string)
			details.SavedSearchId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) mapToSchedules(fieldKeyFormat string) ([]oci_log_analytics.Schedule, error) {
	var result []oci_log_analytics.Schedule
	if schedules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule")); ok {
		set := schedules.(*schema.Set)
		interfaces := set.List()
		result = make([]oci_log_analytics.Schedule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := scheduleHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "schedule"), stateDataIndex)
			converted, err := s.mapToSchedule(fieldKeyFormat)
			if err != nil {
				return result, err
			}
			result[i] = converted
		}
	}
	return result, nil
}

func LAActionToMap(obj oci_log_analytics.Action) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_log_analytics.PurgeAction:
		result["type"] = "PURGE"

		if v.CompartmentIdInSubtree != nil {
			result["compartment_id_in_subtree"] = bool(*v.CompartmentIdInSubtree)
		}

		result["data_type"] = string(v.DataType)

		if v.PurgeCompartmentId != nil {
			result["purge_compartment_id"] = string(*v.PurgeCompartmentId)
		}

		if v.PurgeDuration != nil {
			result["purge_duration"] = string(*v.PurgeDuration)
		}

		if v.QueryString != nil {
			result["query_string"] = string(*v.QueryString)
		}
	case oci_log_analytics.StreamAction:
		result["type"] = "STREAM"

		if v.SavedSearchId != nil {
			result["saved_search_id"] = string(*v.SavedSearchId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) mapToSchedule(fieldKeyFormat string) (oci_log_analytics.Schedule, error) {
	var baseObject oci_log_analytics.Schedule
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("CRON"):
		details := oci_log_analytics.CronSchedule{}
		if expression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expression")); ok {
			tmp := expression.(string)
			details.Expression = &tmp
		}
		if timeZone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_zone")); ok {
			tmp := timeZone.(string)
			details.TimeZone = &tmp
		}
		if misfirePolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "misfire_policy")); ok {
			details.MisfirePolicy = oci_log_analytics.ScheduleMisfirePolicyEnum(misfirePolicy.(string))
		}
		baseObject = details
	case strings.ToLower("FIXED_FREQUENCY"):
		details := oci_log_analytics.FixedFrequencySchedule{}
		if recurringInterval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recurring_interval")); ok {
			tmp := recurringInterval.(string)
			details.RecurringInterval = &tmp
		}
		if repeatCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repeat_count")); ok {
			tmp := repeatCount.(int)
			details.RepeatCount = &tmp
		}
		if misfirePolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "misfire_policy")); ok {
			details.MisfirePolicy = oci_log_analytics.ScheduleMisfirePolicyEnum(misfirePolicy.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ScheduleListToMap(obj []oci_log_analytics.Schedule, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	schedules := []interface{}{}
	for _, item := range obj {
		schedules = append(schedules, LoganScheduleToMap(item))
	}
	if datasource {
		result["schedule"] = schedules
	} else {
		result["schedule"] = schema.NewSet(scheduleHashCodeForSets, schedules)
	}

	return result
}

func LoganScheduleToMap(obj oci_log_analytics.Schedule) map[string]interface{} {
	result := map[string]interface{}{}
	result["misfire_policy"] = string(obj.GetMisfirePolicy())

	switch v := (obj).(type) {
	case oci_log_analytics.CronSchedule:
		result["type"] = "CRON"

		if v.Expression != nil {
			result["expression"] = string(*v.Expression)
		}

		if v.TimeZone != nil {
			result["time_zone"] = string(*v.TimeZone)
		}
	case oci_log_analytics.FixedFrequencySchedule:
		result["type"] = "FIXED_FREQUENCY"

		if v.RecurringInterval != nil {
			result["recurring_interval"] = string(*v.RecurringInterval)
		}

		if v.RepeatCount != nil {
			result["repeat_count"] = int(*v.RepeatCount)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func ScheduledTaskSummaryToMap(obj oci_log_analytics.ScheduledTaskSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

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

	result["task_status"] = string(obj.TaskStatus)

	result["task_type"] = string(obj.TaskType)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.WorkRequestId != nil {
		result["work_request_id"] = string(*obj.WorkRequestId)
	}

	return result
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) populateTopLevelPolymorphicCreateScheduledTaskRequest(request *oci_log_analytics.CreateScheduledTaskRequest) error {
	//discriminator
	kindRaw, ok := s.D.GetOkExists("kind")
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("ACCELERATION"):
		// TODO: fix this after ACCELERATION is supported
		return fmt.Errorf("ACCELERATION scheduled tasks are not supported yet")

	case strings.ToLower("STANDARD"):
		details := oci_log_analytics.CreateStandardTaskDetails{}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			request.NamespaceName = &tmp
		}
		if taskType, ok := s.D.GetOkExists("task_type"); ok {
			details.TaskType = oci_log_analytics.TaskTypeEnum(taskType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if action, ok := s.D.GetOkExists("action"); ok {
			if tmpList := action.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action", 0)
				tmp, err := s.mapToAction(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Action = tmp
			}
		}
		if schedules, ok := s.D.GetOkExists("schedules"); ok {
			if tmpList := schedules.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", 0)
				tmp, err := s.mapToSchedules(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Schedules = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateScheduledTaskDetails = details
	default:
		return fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return nil
}

func (s *LogAnalyticsNamespaceScheduledTaskResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_log_analytics.ChangeScheduledTaskCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		changeCompartmentRequest.NamespaceName = &tmp
	} else {
		return fmt.Errorf("namespace is required to update the compartment of a scheduled task")
	}

	if scheduledTaskId, ok := s.D.GetOkExists("scheduled_task_id"); ok {
		tmp := scheduledTaskId.(string)
		changeCompartmentRequest.ScheduledTaskId = &tmp
	} else {
		return fmt.Errorf("scheduled_task_id is required to update the compartment of a scheduled task")
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.ChangeScheduledTaskCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func scheduleHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if schedule_type, ok := m["type"]; ok && schedule_type != "" {
		buf.WriteString(fmt.Sprintf("%v-", schedule_type))
	}
	if expr, ok := m["expression"]; ok && expr != "" {
		buf.WriteString(fmt.Sprintf("%v-", expr))
	}
	if misfirePolicy, ok := m["misfire_policy"]; ok && misfirePolicy != "" {
		buf.WriteString(fmt.Sprintf("%v-", misfirePolicy))
	}
	if recurringInterval, ok := m["recurring_interval"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", recurringInterval))
	}
	if repeatCount, ok := m["repeat_count"]; ok && repeatCount != "" {
		buf.WriteString(fmt.Sprintf("%v-", repeatCount))
	}
	if timeZone, ok := m["time_zone"]; ok && timeZone != "" {
		buf.WriteString(fmt.Sprintf("%v-", timeZone))
	}
	return hashcode.String(buf.String())
}
