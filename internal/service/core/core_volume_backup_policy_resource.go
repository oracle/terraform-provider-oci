// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreVolumeBackupPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVolumeBackupPolicy,
		Read:     readCoreVolumeBackupPolicy,
		Update:   updateCoreVolumeBackupPolicy,
		Delete:   deleteCoreVolumeBackupPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"destination_region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"schedules": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schedulesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"backup_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"period": {
							Type:     schema.TypeString,
							Required: true,
						},
						"retention_seconds": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"day_of_month": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},
						"day_of_week": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"hour_of_day": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},
						"month": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"offset_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"offset_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"time_zone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreVolumeBackupPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreVolumeBackupPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

func updateCoreVolumeBackupPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreVolumeBackupPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVolumeBackupPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeBackupPolicy
	DisableNotFoundRetries bool
}

func (s *CoreVolumeBackupPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVolumeBackupPolicyResourceCrud) Create() error {
	request := oci_core.CreateVolumeBackupPolicyRequest{}

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

	if destinationRegion, ok := s.D.GetOkExists("destination_region"); ok {
		tmp := destinationRegion.(string)
		request.DestinationRegion = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if schedules, ok := s.D.GetOkExists("schedules"); ok {
		set := schedules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.VolumeBackupSchedule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := schedulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", stateDataIndex)
			converted, err := s.mapToVolumeBackupSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("schedules") {
			request.Schedules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeBackupPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackupPolicy
	return nil
}

func (s *CoreVolumeBackupPolicyResourceCrud) Get() error {
	request := oci_core.GetVolumeBackupPolicyRequest{}

	tmp := s.D.Id()
	request.PolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeBackupPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackupPolicy
	return nil
}

func (s *CoreVolumeBackupPolicyResourceCrud) Update() error {
	request := oci_core.UpdateVolumeBackupPolicyRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if destinationRegion, ok := s.D.GetOkExists("destination_region"); ok {
		tmp := destinationRegion.(string)
		request.DestinationRegion = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.PolicyId = &tmp

	if schedules, ok := s.D.GetOkExists("schedules"); ok {
		set := schedules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.VolumeBackupSchedule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := schedulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", stateDataIndex)
			converted, err := s.mapToVolumeBackupSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("schedules") {
			request.Schedules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolumeBackupPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackupPolicy
	return nil
}

func (s *CoreVolumeBackupPolicyResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeBackupPolicyRequest{}

	tmp := s.D.Id()
	request.PolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeBackupPolicy(context.Background(), request)
	return err
}

func (s *CoreVolumeBackupPolicyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DestinationRegion != nil {
		s.D.Set("destination_region", *s.Res.DestinationRegion)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	schedules := []interface{}{}
	for _, item := range s.Res.Schedules {
		schedules = append(schedules, VolumeBackupScheduleToMap(item))
	}
	s.D.Set("schedules", schema.NewSet(schedulesHashCodeForSets, schedules))

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreVolumeBackupPolicyResourceCrud) mapToVolumeBackupSchedule(fieldKeyFormat string) (oci_core.VolumeBackupSchedule, error) {
	result := oci_core.VolumeBackupSchedule{}

	if backupType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_type")); ok {
		result.BackupType = oci_core.VolumeBackupScheduleBackupTypeEnum(backupType.(string))
	}

	if dayOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_month")); ok {
		tmp := dayOfMonth.(int)
		if tmp != -1 {
			result.DayOfMonth = &tmp
		}
	}

	if dayOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_week")); ok {
		result.DayOfWeek = oci_core.VolumeBackupScheduleDayOfWeekEnum(dayOfWeek.(string))
	}

	if hourOfDay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hour_of_day")); ok {
		tmp := hourOfDay.(int)
		if tmp != -1 {
			result.HourOfDay = &tmp
		}
	}

	if month, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "month")); ok {
		result.Month = oci_core.VolumeBackupScheduleMonthEnum(month.(string))
	}

	if offsetSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "offset_seconds")); ok {
		tmp := offsetSeconds.(int)
		result.OffsetSeconds = &tmp
	}

	if offsetType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "offset_type")); ok {
		result.OffsetType = oci_core.VolumeBackupScheduleOffsetTypeEnum(offsetType.(string))
	}

	if period, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "period")); ok {
		result.Period = oci_core.VolumeBackupSchedulePeriodEnum(period.(string))
	}

	if retentionSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_seconds")); ok {
		tmp := retentionSeconds.(int)
		result.RetentionSeconds = &tmp
	}

	if timeZone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_zone")); ok {
		result.TimeZone = oci_core.VolumeBackupScheduleTimeZoneEnum(timeZone.(string))
	}

	return result, nil
}

func VolumeBackupScheduleToMap(obj oci_core.VolumeBackupSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	result["backup_type"] = string(obj.BackupType)

	if obj.DayOfMonth != nil {
		result["day_of_month"] = int(*obj.DayOfMonth)
	} else {
		result["day_of_month"] = -1
	}

	result["day_of_week"] = string(obj.DayOfWeek)

	if obj.HourOfDay != nil {
		result["hour_of_day"] = int(*obj.HourOfDay)
	} else {
		result["hour_of_day"] = -1
	}

	result["month"] = string(obj.Month)

	if obj.OffsetSeconds != nil {
		result["offset_seconds"] = int(*obj.OffsetSeconds)
	}

	result["offset_type"] = string(obj.OffsetType)

	result["period"] = string(obj.Period)

	if obj.RetentionSeconds != nil {
		result["retention_seconds"] = int(*obj.RetentionSeconds)
	}

	result["time_zone"] = string(obj.TimeZone)

	return result
}

func schedulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if backupType, ok := m["backup_type"]; ok && backupType != "" {
		buf.WriteString(fmt.Sprintf("%v-", backupType))
	}
	if dayOfMonth, ok := m["day_of_month"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", dayOfMonth))
	}
	if dayOfWeek, ok := m["day_of_week"]; ok && dayOfWeek != "" {
		buf.WriteString(fmt.Sprintf("%v-", dayOfWeek))
	}
	if hourOfDay, ok := m["hour_of_day"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", hourOfDay))
	}
	if month, ok := m["month"]; ok && month != "" {
		buf.WriteString(fmt.Sprintf("%v-", month))
	}
	if offsetSeconds, ok := m["offset_seconds"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", offsetSeconds))
	}
	if offsetType, ok := m["offset_type"]; ok && offsetType != "" {
		buf.WriteString(fmt.Sprintf("%v-", offsetType))
	}
	if period, ok := m["period"]; ok && period != "" {
		buf.WriteString(fmt.Sprintf("%v-", period))
	}
	if retentionSeconds, ok := m["retention_seconds"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", retentionSeconds))
	}
	if timeZone, ok := m["time_zone"]; ok && timeZone != "" {
		buf.WriteString(fmt.Sprintf("%v-", timeZone))
	} else {
		buf.WriteString(fmt.Sprintf("%v-", "UTC"))
	}
	return hashcode.String(buf.String())
}
