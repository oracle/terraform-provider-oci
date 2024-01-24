// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardSecurityZoneResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardSecurityZone,
		Read:     readCloudGuardSecurityZone,
		Update:   updateCloudGuardSecurityZone,
		Delete:   deleteCloudGuardSecurityZone,
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
			"security_zone_recipe_id": {
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
			"inherited_by_compartments": {
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
			"security_zone_target_id": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCloudGuardSecurityZone(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardSecurityZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardSecurityZone(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardSecurityZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardSecurityZone(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardSecurityZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardSecurityZone(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardSecurityZoneResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardSecurityZoneResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.SecurityZone
	DisableNotFoundRetries bool
}

func (s *CloudGuardSecurityZoneResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardSecurityZoneResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardSecurityZoneResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardSecurityZoneResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardSecurityZoneResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardSecurityZoneResourceCrud) Create() error {
	request := oci_cloud_guard.CreateSecurityZoneRequest{}

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

	if securityZoneRecipeId, ok := s.D.GetOkExists("security_zone_recipe_id"); ok {
		tmp := securityZoneRecipeId.(string)
		request.SecurityZoneRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateSecurityZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityZone
	return nil
}

func (s *CloudGuardSecurityZoneResourceCrud) Get() error {
	request := oci_cloud_guard.GetSecurityZoneRequest{}

	tmp := s.D.Id()
	request.SecurityZoneId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetSecurityZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityZone
	return nil
}

func (s *CloudGuardSecurityZoneResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_guard.UpdateSecurityZoneRequest{}

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

	tmp := s.D.Id()
	request.SecurityZoneId = &tmp

	if securityZoneRecipeId, ok := s.D.GetOkExists("security_zone_recipe_id"); ok {
		tmp := securityZoneRecipeId.(string)
		request.SecurityZoneRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateSecurityZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityZone
	return nil
}

func (s *CloudGuardSecurityZoneResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteSecurityZoneRequest{}

	tmp := s.D.Id()
	request.SecurityZoneId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteSecurityZone(context.Background(), request)
	return err
}

func (s *CloudGuardSecurityZoneResourceCrud) SetData() error {
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

	s.D.Set("inherited_by_compartments", s.Res.InheritedByCompartments)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecurityZoneRecipeId != nil {
		s.D.Set("security_zone_recipe_id", *s.Res.SecurityZoneRecipeId)
	}

	if s.Res.SecurityZoneTargetId != nil {
		s.D.Set("security_zone_target_id", *s.Res.SecurityZoneTargetId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func SecurityZoneSummaryToMap(obj oci_cloud_guard.SecurityZoneSummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.SecurityZoneRecipeId != nil {
		result["security_zone_recipe_id"] = string(*obj.SecurityZoneRecipeId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *CloudGuardSecurityZoneResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_guard.ChangeSecurityZoneCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SecurityZoneId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.ChangeSecurityZoneCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
