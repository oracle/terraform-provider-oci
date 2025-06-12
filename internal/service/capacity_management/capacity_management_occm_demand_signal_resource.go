// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccmDemandSignalResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementOccmDemandSignal,
		Read:     readCapacityManagementOccmDemandSignal,
		Update:   updateCapacityManagementOccmDemandSignal,
		Delete:   deleteCapacityManagementOccmDemandSignal,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
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
			"lifecycle_details": {
				Type:     schema.TypeString,
				Optional: true,
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

func createCapacityManagementOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

func updateCapacityManagementOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementOccmDemandSignal(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CapacityManagementOccmDemandSignalResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.DemandSignalClient
	Res                    *oci_capacity_management.OccmDemandSignal
	DisableNotFoundRetries bool
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalLifecycleStateCreating),
	}
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalLifecycleStateActive),
	}
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalLifecycleStateDeleting),
	}
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_capacity_management.OccmDemandSignalLifecycleStateDeleted),
	}
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) Create() error {
	request := oci_capacity_management.CreateOccmDemandSignalRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.CreateOccmDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccmDemandSignal
	return nil
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) Get() error {
	request := oci_capacity_management.GetOccmDemandSignalRequest{}

	tmp := s.D.Id()
	request.OccmDemandSignalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.GetOccmDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccmDemandSignal
	return nil
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) Update() error {
	request := oci_capacity_management.UpdateOccmDemandSignalRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		request.LifecycleDetails = oci_capacity_management.UpdateOccmDemandSignalDetailsLifecycleDetailsEnum(lifecycleDetails.(string))
	}

	tmp := s.D.Id()
	request.OccmDemandSignalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateOccmDemandSignal(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccmDemandSignal
	return nil
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) Delete() error {
	request := oci_capacity_management.DeleteOccmDemandSignalRequest{}

	tmp := s.D.Id()
	request.OccmDemandSignalId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	_, err := s.Client.DeleteOccmDemandSignal(context.Background(), request)
	return err
}

func (s *CapacityManagementOccmDemandSignalResourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

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

func OccmDemandSignalSummaryToMap(obj oci_capacity_management.OccmDemandSignalSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

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
