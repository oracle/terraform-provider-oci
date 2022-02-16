// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_osmanagement "github.com/oracle/oci-go-sdk/v58/osmanagement"
)

func OsmanagementManagedInstanceGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsmanagementManagedInstanceGroup,
		Read:     readOsmanagementManagedInstanceGroup,
		Update:   updateOsmanagementManagedInstanceGroup,
		Delete:   deleteOsmanagementManagedInstanceGroup,
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
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"managed_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createOsmanagementManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readOsmanagementManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateOsmanagementManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsmanagementManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OsmanagementManagedInstanceGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_osmanagement.OsManagementClient
	Res                    *oci_osmanagement.ManagedInstanceGroup
	DisableNotFoundRetries bool
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesCreating),
	}
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesActive),
	}
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesDeleting),
	}
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_osmanagement.LifecycleStatesDeleted),
	}
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) Create() error {
	request := oci_osmanagement.CreateManagedInstanceGroupRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_osmanagement.OsFamiliesEnum(osFamily.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.CreateManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstanceGroup
	return nil
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) Get() error {
	request := oci_osmanagement.GetManagedInstanceGroupRequest{}

	tmp := s.D.Id()
	request.ManagedInstanceGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.GetManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstanceGroup
	return nil
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_osmanagement.UpdateManagedInstanceGroupRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ManagedInstanceGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.UpdateManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstanceGroup
	return nil
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) Delete() error {
	request := oci_osmanagement.DeleteManagedInstanceGroupRequest{}

	tmp := s.D.Id()
	request.ManagedInstanceGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	_, err := s.Client.DeleteManagedInstanceGroup(context.Background(), request)
	return err
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) SetData() error {
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

	managedInstances := []interface{}{}
	for _, item := range s.Res.ManagedInstances {
		managedInstances = append(managedInstances, IdToMap(item))
	}
	s.D.Set("managed_instances", managedInstances)

	s.D.Set("os_family", s.Res.OsFamily)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}

func IdToMap(obj oci_osmanagement.Id) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *OsmanagementManagedInstanceGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_osmanagement.ChangeManagedInstanceGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ManagedInstanceGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	_, err := s.Client.ChangeManagedInstanceGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
