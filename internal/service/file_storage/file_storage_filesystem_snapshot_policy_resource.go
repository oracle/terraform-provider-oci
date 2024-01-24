// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageFilesystemSnapshotPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageFilesystemSnapshotPolicy,
		Read:     readFileStorageFilesystemSnapshotPolicy,
		Update:   updateFileStorageFilesystemSnapshotPolicy,
		Delete:   deleteFileStorageFilesystemSnapshotPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
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
			"policy_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"period": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time_zone": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"day_of_month": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"day_of_week": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"hour_of_day": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"month": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"retention_duration_in_seconds": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"schedule_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_schedule_start": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateInactive),
					string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateActive),
				}, true),
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFileStorageFilesystemSnapshotPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFilesystemSnapshotPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_file_storage.FilesystemSnapshotPolicyLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_file_storage.FilesystemSnapshotPolicyLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopFilesystemSnapshotPolicy(); err != nil {
			return err
		}
		sync.D.Set("state", oci_file_storage.FilesystemSnapshotPolicyLifecycleStateInactive)
	}
	return nil

}

func readFileStorageFilesystemSnapshotPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFilesystemSnapshotPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageFilesystemSnapshotPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFilesystemSnapshotPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_file_storage.FilesystemSnapshotPolicyLifecycleStateActive == oci_file_storage.FilesystemSnapshotPolicyLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_file_storage.FilesystemSnapshotPolicyLifecycleStateInactive == oci_file_storage.FilesystemSnapshotPolicyLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartFilesystemSnapshotPolicy(); err != nil {
			return err
		}
		sync.D.Set("state", oci_file_storage.FilesystemSnapshotPolicyLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopFilesystemSnapshotPolicy(); err != nil {
			return err
		}
		sync.D.Set("state", oci_file_storage.FilesystemSnapshotPolicyLifecycleStateInactive)
	}

	return nil
}

func deleteFileStorageFilesystemSnapshotPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFilesystemSnapshotPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageFilesystemSnapshotPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.FilesystemSnapshotPolicy
	DisableNotFoundRetries bool
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateCreating),
	}
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateActive),
	}
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateDeleting),
	}
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.FilesystemSnapshotPolicyLifecycleStateDeleted),
	}
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) Create() error {
	request := oci_file_storage.CreateFilesystemSnapshotPolicyRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if policyPrefix, ok := s.D.GetOkExists("policy_prefix"); ok {
		tmp := policyPrefix.(string)
		request.PolicyPrefix = &tmp
	}

	if schedules, ok := s.D.GetOkExists("schedules"); ok {
		interfaces := schedules.([]interface{})
		tmp := make([]oci_file_storage.SnapshotSchedule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", stateDataIndex)
			converted, err := s.mapToSnapshotSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("schedules") {
			request.Schedules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateFilesystemSnapshotPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FilesystemSnapshotPolicy
	return nil
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) Get() error {
	request := oci_file_storage.GetFilesystemSnapshotPolicyRequest{}

	tmp := s.D.Id()
	request.FilesystemSnapshotPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetFilesystemSnapshotPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FilesystemSnapshotPolicy
	return nil
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_file_storage.UpdateFilesystemSnapshotPolicyRequest{}

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

	tmp := s.D.Id()
	request.FilesystemSnapshotPolicyId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if policyPrefix, ok := s.D.GetOkExists("policy_prefix"); ok {
		tmp := policyPrefix.(string)
		request.PolicyPrefix = &tmp
	}

	if schedules, ok := s.D.GetOkExists("schedules"); ok {
		interfaces := schedules.([]interface{})
		tmp := make([]oci_file_storage.SnapshotSchedule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", stateDataIndex)
			converted, err := s.mapToSnapshotSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("schedules") {
			request.Schedules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateFilesystemSnapshotPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FilesystemSnapshotPolicy
	return nil
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) Delete() error {
	request := oci_file_storage.DeleteFilesystemSnapshotPolicyRequest{}

	tmp := s.D.Id()
	request.FilesystemSnapshotPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteFilesystemSnapshotPolicy(context.Background(), request)
	return err
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
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

	if s.Res.PolicyPrefix != nil {
		s.D.Set("policy_prefix", *s.Res.PolicyPrefix)
	}

	schedules := []interface{}{}
	for _, item := range s.Res.Schedules {
		schedules = append(schedules, SnapshotScheduleToMap(item))
	}
	s.D.Set("schedules", schedules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) StartFilesystemSnapshotPolicy() error {
	request := oci_file_storage.UnpauseFilesystemSnapshotPolicyRequest{}

	idTmp := s.D.Id()
	request.FilesystemSnapshotPolicyId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.UnpauseFilesystemSnapshotPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res.LifecycleState == oci_file_storage.FilesystemSnapshotPolicyLifecycleStateActive
	}
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) StopFilesystemSnapshotPolicy() error {
	request := oci_file_storage.PauseFilesystemSnapshotPolicyRequest{}

	idTmp := s.D.Id()
	request.FilesystemSnapshotPolicyId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.PauseFilesystemSnapshotPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool {
		return s.Res.LifecycleState == oci_file_storage.FilesystemSnapshotPolicyLifecycleStateInactive
	}
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) mapToSnapshotSchedule(fieldKeyFormat string) (oci_file_storage.SnapshotSchedule, error) {
	result := oci_file_storage.SnapshotSchedule{}

	if dayOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_month")); ok {
		tmp := dayOfMonth.(int)
		if tmp == 0 {
			result.DayOfMonth = nil
		} else {
			result.DayOfMonth = &tmp
		}
	}

	if dayOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_week")); ok {
		result.DayOfWeek = oci_file_storage.SnapshotScheduleDayOfWeekEnum(dayOfWeek.(string))
	}

	if period, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "period")); ok {
		result.Period = oci_file_storage.SnapshotSchedulePeriodEnum(period.(string))
	}

	if hourOfDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hour_of_day")); ok {
		tmp := hourOfDay.(int)
		if tmp == 0 && result.Period == oci_file_storage.SnapshotSchedulePeriodHourly {
			result.HourOfDay = nil
		} else {
			result.HourOfDay = &tmp
		}
	}

	if month, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "month")); ok {
		result.Month = oci_file_storage.SnapshotScheduleMonthEnum(month.(string))
	}

	if retentionDurationInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_duration_in_seconds")); ok {
		tmp := retentionDurationInSeconds.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert retentionDurationInSeconds string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.RetentionDurationInSeconds = &tmpInt64
	}

	if schedulePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_prefix")); ok {
		tmp := schedulePrefix.(string)
		result.SchedulePrefix = &tmp
	}

	if timeScheduleStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_schedule_start")); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduleStart.(string))
		if err != nil {
			return result, err
		}
		result.TimeScheduleStart = &oci_common.SDKTime{Time: tmp}
	}

	if timeZone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_zone")); ok {
		result.TimeZone = oci_file_storage.SnapshotScheduleTimeZoneEnum(timeZone.(string))
	}

	return result, nil
}

func SnapshotScheduleToMap(obj oci_file_storage.SnapshotSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DayOfMonth != nil {
		result["day_of_month"] = int(*obj.DayOfMonth)
	}

	result["day_of_week"] = string(obj.DayOfWeek)

	if obj.HourOfDay != nil {
		result["hour_of_day"] = int(*obj.HourOfDay)
	}

	result["month"] = string(obj.Month)

	result["period"] = string(obj.Period)

	if obj.RetentionDurationInSeconds != nil {
		result["retention_duration_in_seconds"] = strconv.FormatInt(*obj.RetentionDurationInSeconds, 10)
	}

	if obj.SchedulePrefix != nil {
		result["schedule_prefix"] = string(*obj.SchedulePrefix)
	}

	if obj.TimeScheduleStart != nil {
		result["time_schedule_start"] = obj.TimeScheduleStart.Format(time.RFC3339Nano)
	}

	result["time_zone"] = string(obj.TimeZone)

	return result
}

func (s *FileStorageFilesystemSnapshotPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_file_storage.ChangeFilesystemSnapshotPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FilesystemSnapshotPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.ChangeFilesystemSnapshotPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
