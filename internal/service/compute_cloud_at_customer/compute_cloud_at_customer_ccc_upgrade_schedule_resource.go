// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package compute_cloud_at_customer

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ComputeCloudAtCustomerCccUpgradeScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createComputeCloudAtCustomerCccUpgradeSchedule,
		Read:     readComputeCloudAtCustomerCccUpgradeSchedule,
		Update:   updateComputeCloudAtCustomerCccUpgradeSchedule,
		Delete:   deleteComputeCloudAtCustomerCccUpgradeSchedule,
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
			"events": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"schedule_event_duration": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time_start": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Optional
						"schedule_event_recurrences": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
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

			// Computed
			"infrastructure_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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

func createComputeCloudAtCustomerCccUpgradeSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.CreateResource(d, sync)
}

func readComputeCloudAtCustomerCccUpgradeSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.ReadResource(sync)
}

func updateComputeCloudAtCustomerCccUpgradeSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteComputeCloudAtCustomerCccUpgradeSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient
	Res                    *oci_compute_cloud_at_customer.CccUpgradeSchedule
	DisableNotFoundRetries bool
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateActive),
		string(oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateNeedsAttention),
	}
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateDeleted),
	}
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) Create() error {
	request := oci_compute_cloud_at_customer.CreateCccUpgradeScheduleRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if events, ok := s.D.GetOkExists("events"); ok {
		interfaces := events.([]interface{})
		tmp := make([]oci_compute_cloud_at_customer.CreateCccScheduleEvent, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "events", stateDataIndex)
			converted, err := s.mapToCreateCccScheduleEvent(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("events") {
			request.Events = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	response, err := s.Client.CreateCccUpgradeSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CccUpgradeSchedule
	return nil
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) Get() error {
	request := oci_compute_cloud_at_customer.GetCccUpgradeScheduleRequest{}

	tmp := s.D.Id()
	request.CccUpgradeScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	response, err := s.Client.GetCccUpgradeSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CccUpgradeSchedule
	return nil
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_compute_cloud_at_customer.UpdateCccUpgradeScheduleRequest{}

	tmp := s.D.Id()
	request.CccUpgradeScheduleId = &tmp

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if events, ok := s.D.GetOkExists("events"); ok {
		interfaces := events.([]interface{})
		tmp := make([]oci_compute_cloud_at_customer.UpdateCccScheduleEvent, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "events", stateDataIndex)
			converted, err := s.mapToUpdateCccScheduleEvent(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("events") {
			request.Events = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	response, err := s.Client.UpdateCccUpgradeSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CccUpgradeSchedule
	return nil
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) Delete() error {
	request := oci_compute_cloud_at_customer.DeleteCccUpgradeScheduleRequest{}

	tmp := s.D.Id()
	request.CccUpgradeScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	_, err := s.Client.DeleteCccUpgradeSchedule(context.Background(), request)
	return err
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	events := []interface{}{}
	for _, item := range s.Res.Events {
		events = append(events, CccScheduleEventToMap(item))
	}
	s.D.Set("events", events)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_ids", s.Res.InfrastructureIds)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

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

func CccUpgradeScheduleSummaryToMap(obj oci_compute_cloud_at_customer.CccUpgradeScheduleSummary) map[string]interface{} {
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

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) mapToCreateCccScheduleEvent(fieldKeyFormat string) (oci_compute_cloud_at_customer.CreateCccScheduleEvent, error) {
	result := oci_compute_cloud_at_customer.CreateCccScheduleEvent{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if scheduleEventDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_event_duration")); ok {
		tmp := scheduleEventDuration.(string)
		result.ScheduleEventDuration = &tmp
	}

	if scheduleEventRecurrences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_event_recurrences")); ok {
		tmp := scheduleEventRecurrences.(string)
		result.ScheduleEventRecurrences = &tmp
	}

	if timeStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return result, err
		}
		result.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) mapToUpdateCccScheduleEvent(fieldKeyFormat string) (oci_compute_cloud_at_customer.UpdateCccScheduleEvent, error) {
	result := oci_compute_cloud_at_customer.UpdateCccScheduleEvent{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if scheduleEventDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_event_duration")); ok {
		tmp := scheduleEventDuration.(string)
		result.ScheduleEventDuration = &tmp
	}

	if scheduleEventRecurrences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_event_recurrences")); ok {
		tmp := scheduleEventRecurrences.(string)
		result.ScheduleEventRecurrences = &tmp
	}

	if timeStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return result, err
		}
		result.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func CccScheduleEventToMap(obj oci_compute_cloud_at_customer.CccScheduleEvent) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ScheduleEventDuration != nil {
		result["schedule_event_duration"] = string(*obj.ScheduleEventDuration)
	}

	if obj.ScheduleEventRecurrences != nil {
		result["schedule_event_recurrences"] = string(*obj.ScheduleEventRecurrences)
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.Format(time.RFC3339Nano)
	}

	return result
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_compute_cloud_at_customer.ChangeCccUpgradeScheduleCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CccUpgradeScheduleId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "compute_cloud_at_customer")

	_, err := s.Client.ChangeCccUpgradeScheduleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
