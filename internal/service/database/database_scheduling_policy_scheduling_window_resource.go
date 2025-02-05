// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseSchedulingPolicySchedulingWindowResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseSchedulingPolicySchedulingWindow,
		Read:     readDatabaseSchedulingPolicySchedulingWindow,
		Update:   updateDatabaseSchedulingPolicySchedulingWindow,
		Delete:   deleteDatabaseSchedulingPolicySchedulingWindow,
		Schema: map[string]*schema.Schema{
			// Required
			"scheduling_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"window_preference": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"days_of_week": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"duration": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"is_enforced_duration": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"start_time": {
							Type:     schema.TypeString,
							Required: true,
						},
						"weeks_of_month": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 4,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						// Optional
						"months": {
							Type: schema.TypeList,
							//Optional: true,
							Required: true,
							//Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
										//Computed: true,
									},

									// Optional

									// Computed
								},
							},
						},

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

			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_next_scheduling_window_starts": {
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

func createDatabaseSchedulingPolicySchedulingWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicySchedulingWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseSchedulingPolicySchedulingWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicySchedulingWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseSchedulingPolicySchedulingWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicySchedulingWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseSchedulingPolicySchedulingWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicySchedulingWindowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseSchedulingPolicySchedulingWindowResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.SchedulingWindow
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.SchedulingWindowLifecycleStateCreating),
	}
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.SchedulingWindowLifecycleStateAvailable),
	}
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.SchedulingWindowLifecycleStateDeleting),
	}
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.SchedulingWindowLifecycleStateDeleted),
	}
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) Create() error {
	request := oci_database.CreateSchedulingWindowRequest{}

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

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	if windowPreference, ok := s.D.GetOkExists("window_preference"); ok {
		if tmpList := windowPreference.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "window_preference", 0)
			tmp, err := s.mapToWindowPreferenceDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WindowPreference = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateSchedulingWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SchedulingWindow
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}

	return s.Get()
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) Get() error {
	request := oci_database.GetSchedulingWindowRequest{}

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	tmp := s.D.Id()
	request.SchedulingWindowId = &tmp

	schedulingPolicyId, schedulingWindowId, err := parseSchedulingPolicySchedulingWindowCompositeId(s.D.Id())
	if err == nil {
		request.SchedulingPolicyId = &schedulingPolicyId
		request.SchedulingWindowId = &schedulingWindowId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetSchedulingWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SchedulingWindow
	return nil
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) Update() error {
	request := oci_database.UpdateSchedulingWindowRequest{}

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

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	tmp := s.D.Id()
	request.SchedulingWindowId = &tmp

	if windowPreference, ok := s.D.GetOkExists("window_preference"); ok {
		if tmpList := windowPreference.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "window_preference", 0)
			tmp, err := s.mapToWindowPreferenceDetail(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WindowPreference = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateSchedulingWindow(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "schedulingwindow", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) Delete() error {
	request := oci_database.DeleteSchedulingWindowRequest{}

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	tmp := s.D.Id()
	request.SchedulingWindowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteSchedulingWindow(context.Background(), request)
	return err
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) SetData() error {
	log.Printf("[WARN] executing SetData()")
	schedulingPolicyId, schedulingWindowId, err := parseSchedulingPolicySchedulingWindowCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("scheduling_policy_id", &schedulingPolicyId)
		s.D.SetId(schedulingWindowId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SchedulingPolicyId != nil {
		s.D.Set("scheduling_policy_id", *s.Res.SchedulingPolicyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNextSchedulingWindowStarts != nil {
		s.D.Set("time_next_scheduling_window_starts", s.Res.TimeNextSchedulingWindowStarts.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WindowPreference != nil {
		s.D.Set("window_preference", []interface{}{WindowPreferenceDetailToMap(s.Res.WindowPreference)})
		log.Printf("[WARN] window_preference: %s", s.D.Get("window_preference"))
	} else {
		s.D.Set("window_preference", nil)
	}

	return nil
}

func GetSchedulingPolicySchedulingWindowCompositeId(schedulingPolicyId string, schedulingWindowId string) string {
	schedulingPolicyId = url.PathEscape(schedulingPolicyId)
	schedulingWindowId = url.PathEscape(schedulingWindowId)
	compositeId := "schedulingPolicies/" + schedulingPolicyId + "/schedulingWindows/" + schedulingWindowId
	return compositeId
}

func parseSchedulingPolicySchedulingWindowCompositeId(compositeId string) (schedulingPolicyId string, schedulingWindowId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("schedulingPolicies/.*/schedulingWindows/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	schedulingPolicyId, _ = url.PathUnescape(parts[1])
	schedulingWindowId, _ = url.PathUnescape(parts[3])

	return
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) mapToDayOfWeek(fieldKeyFormat string) (oci_database.DayOfWeek, error) {
	result := oci_database.DayOfWeek{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.DayOfWeekNameEnum(name.(string))
	}

	return result, nil
}

func DayOfWeekToMapWindow(obj oci_database.DayOfWeek) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) mapToMonth(fieldKeyFormat string) (oci_database.Month, error) {
	result := oci_database.Month{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_database.MonthNameEnum(name.(string))
	}

	return result, nil
}

func MonthToMapWindow(obj oci_database.Month) map[string]interface{} {
	result := map[string]interface{}{}

	result["name"] = string(obj.Name)

	return result
}

func (s *DatabaseSchedulingPolicySchedulingWindowResourceCrud) mapToWindowPreferenceDetail(fieldKeyFormat string) (oci_database.WindowPreferenceDetail, error) {
	result := oci_database.WindowPreferenceDetail{}

	if daysOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_of_week")); ok {
		interfaces := daysOfWeek.([]interface{})
		tmp := make([]oci_database.DayOfWeek, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "days_of_week"), stateDataIndex)
			converted, err := s.mapToDayOfWeek(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days_of_week")) {
			result.DaysOfWeek = tmp
		}
	}

	if duration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duration")); ok {
		tmp := duration.(int)
		result.Duration = &tmp
	}

	if isEnforcedDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enforced_duration")); ok {
		tmp := isEnforcedDuration.(bool)
		result.IsEnforcedDuration = &tmp
	}

	if months, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "months")); ok {
		interfaces := months.([]interface{})
		tmp := make([]oci_database.Month, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "months"), stateDataIndex)
			converted, err := s.mapToMonth(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "months")) {
			result.Months = tmp
		}
	}

	if startTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_time")); ok {
		tmp := startTime.(string)
		result.StartTime = &tmp
	}

	if weeksOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")); ok {
		interfaces := weeksOfMonth.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "weeks_of_month")) {
			result.WeeksOfMonth = tmp
		}
	}

	return result, nil
}

func WindowPreferenceDetailToMap(obj *oci_database.WindowPreferenceDetail) map[string]interface{} {
	result := map[string]interface{}{}

	var daysOfWeek []interface{}
	for _, item := range obj.DaysOfWeek {
		daysOfWeek = append(daysOfWeek, DayOfWeekToMapWindow(item))
	}
	result["days_of_week"] = daysOfWeek

	if obj.Duration != nil {
		result["duration"] = int(*obj.Duration)
	}

	if obj.IsEnforcedDuration != nil {
		result["is_enforced_duration"] = bool(*obj.IsEnforcedDuration)
	}

	var months []interface{}
	for _, item := range obj.Months {
		months = append(months, MonthToMapWindow(item))
	}
	result["months"] = months

	if obj.StartTime != nil {
		result["start_time"] = string(*obj.StartTime)
	}

	result["weeks_of_month"] = obj.WeeksOfMonth

	return result
}
