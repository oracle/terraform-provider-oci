// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"
	"fmt"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"
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
			"managed_instance_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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
	managedInstanceGroupId := s.Res.Id

	if managedInstanceIds, ok := s.D.GetOkExists("managed_instance_ids"); ok {
		interfaces := managedInstanceIds.([]interface{})
		for i := range interfaces {
			if interfaces[i] != nil {
				managedInstanceId := interfaces[i].(string)
				attachManagedInstanceToGroupRequest := oci_osmanagement.AttachManagedInstanceToManagedInstanceGroupRequest{}
				attachManagedInstanceToGroupRequest.ManagedInstanceId = &managedInstanceId
				attachManagedInstanceToGroupRequest.ManagedInstanceGroupId = managedInstanceGroupId

				_, attachErr := s.Client.AttachManagedInstanceToManagedInstanceGroup(context.Background(), attachManagedInstanceToGroupRequest)
				if attachErr != nil {
					return fmt.Errorf("failed to attach managed instance to  managed instance Group, error: %v", attachErr)
				}
			}
		}
	}

	request2 := oci_osmanagement.GetManagedInstanceGroupRequest{}
	request2.ManagedInstanceGroupId = managedInstanceGroupId
	request2.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")
	response2, err := s.Client.GetManagedInstanceGroup(context.Background(), request2)
	if err != nil {
		return err
	}

	s.Res = &response2.ManagedInstanceGroup

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ManagedInstanceGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.UpdateManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstanceGroup
	managedInstanceGroupId := s.Res.Id

	if s.D.HasChange("managed_instance_ids") {
		o, n := s.D.GetChange("managed_instance_ids")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := schema.NewSet(tfresource.LiteralTypeHashCodeForSets, o.([]interface{}))
		ns := schema.NewSet(tfresource.LiteralTypeHashCodeForSets, n.([]interface{}))

		newManagedInstancesToAttach := ns.Difference(os).List()
		oldManagedInstancesToDetach := os.Difference(ns).List()

		for _, oldManagedInstanceId := range oldManagedInstancesToDetach {
			managedInstanceId := oldManagedInstanceId.(string)
			detachManagedInstanceFromGroupRequest := oci_osmanagement.DetachManagedInstanceFromManagedInstanceGroupRequest{}
			detachManagedInstanceFromGroupRequest.ManagedInstanceId = &managedInstanceId
			detachManagedInstanceFromGroupRequest.ManagedInstanceGroupId = managedInstanceGroupId

			_, detachErr := s.Client.DetachManagedInstanceFromManagedInstanceGroup(context.Background(), detachManagedInstanceFromGroupRequest)
			if detachErr != nil {
				return fmt.Errorf("failed to detach managed instance from managed instance Group request, error: %v", detachErr)
			}
		}

		for _, newManagedInstance := range newManagedInstancesToAttach {
			managedInstanceId := newManagedInstance.(string)
			attachManagedInstanceToGroupRequest := oci_osmanagement.AttachManagedInstanceToManagedInstanceGroupRequest{}
			attachManagedInstanceToGroupRequest.ManagedInstanceId = &managedInstanceId
			attachManagedInstanceToGroupRequest.ManagedInstanceGroupId = managedInstanceGroupId

			_, attachErr := s.Client.AttachManagedInstanceToManagedInstanceGroup(context.Background(), attachManagedInstanceToGroupRequest)
			if attachErr != nil {
				return fmt.Errorf("failed to attach parent software source, error: %v", attachErr)
			}
		}
	}

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
	managedInstanceIds := []string{}
	for _, item := range s.Res.ManagedInstances {
		managedInstances = append(managedInstances, IdToMap(item))
		managedInstanceIds = append(managedInstanceIds, *item.Id)
	}
	s.D.Set("managed_instances", managedInstances)
	s.D.Set("managed_instance_ids", managedInstanceIds)

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
