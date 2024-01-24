// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeDiscoveryScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudBridgeDiscoverySchedule,
		Read:     readCloudBridgeDiscoverySchedule,
		Update:   updateCloudBridgeDiscoverySchedule,
		Delete:   deleteCloudBridgeDiscoverySchedule,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"execution_recurrences": {
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

			// Computed
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

func createCloudBridgeDiscoverySchedule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeDiscoveryScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudBridgeDiscoverySchedule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeDiscoveryScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()

	return tfresource.ReadResource(sync)
}

func updateCloudBridgeDiscoverySchedule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeDiscoveryScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudBridgeDiscoverySchedule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeDiscoveryScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudBridgeDiscoveryScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_bridge.DiscoveryClient
	Res                    *oci_cloud_bridge.DiscoverySchedule
	DisableNotFoundRetries bool
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_bridge.DiscoveryScheduleLifecycleStateActive),
	}
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_bridge.DiscoveryScheduleLifecycleStateDeleted),
	}
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) Create() error {
	request := oci_cloud_bridge.CreateDiscoveryScheduleRequest{}

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

	if executionRecurrences, ok := s.D.GetOkExists("execution_recurrences"); ok {
		tmp := executionRecurrences.(string)
		request.ExecutionRecurrences = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.CreateDiscoverySchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DiscoverySchedule
	return nil
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) Get() error {
	request := oci_cloud_bridge.GetDiscoveryScheduleRequest{}

	tmp := s.D.Id()
	request.DiscoveryScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.GetDiscoverySchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DiscoverySchedule
	return nil
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_bridge.UpdateDiscoveryScheduleRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DiscoveryScheduleId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if executionRecurrences, ok := s.D.GetOkExists("execution_recurrences"); ok {
		tmp := executionRecurrences.(string)
		request.ExecutionRecurrences = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.UpdateDiscoverySchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DiscoverySchedule
	return nil
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) Delete() error {
	request := oci_cloud_bridge.DeleteDiscoveryScheduleRequest{}

	tmp := s.D.Id()
	request.DiscoveryScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.DeleteDiscoverySchedule(context.Background(), request)
	return err
}

func (s *CloudBridgeDiscoveryScheduleResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExecutionRecurrences != nil {
		s.D.Set("execution_recurrences", *s.Res.ExecutionRecurrences)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

func DiscoveryScheduleSummaryToMap(obj oci_cloud_bridge.DiscoveryScheduleSummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

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

func (s *CloudBridgeDiscoveryScheduleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_bridge.ChangeDiscoveryScheduleCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DiscoveryScheduleId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.ChangeDiscoveryScheduleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
