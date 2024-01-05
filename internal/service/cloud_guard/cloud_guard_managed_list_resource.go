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

func CloudGuardManagedListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardManagedList,
		Read:     readCloudGuardManagedList,
		Update:   updateCloudGuardManagedList,
		Delete:   deleteCloudGuardManagedList,
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
			"list_items": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"list_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source_managed_list_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"feed_provider": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_editable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecyle_details": {
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

func createCloudGuardManagedList(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardManagedListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardManagedList(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardManagedListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardManagedList(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardManagedListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardManagedList(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardManagedListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardManagedListResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.ManagedList
	DisableNotFoundRetries bool
}

func (s *CloudGuardManagedListResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardManagedListResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardManagedListResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardManagedListResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardManagedListResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardManagedListResourceCrud) Create() error {
	request := oci_cloud_guard.CreateManagedListRequest{}

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

	if listItems, ok := s.D.GetOkExists("list_items"); ok {
		interfaces := listItems.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("list_items") {
			request.ListItems = tmp
		}
	}

	if listType, ok := s.D.GetOkExists("list_type"); ok {
		request.ListType = oci_cloud_guard.ManagedListTypeEnum(listType.(string))
	}

	if sourceManagedListId, ok := s.D.GetOkExists("source_managed_list_id"); ok {
		tmp := sourceManagedListId.(string)
		request.SourceManagedListId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateManagedList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedList
	return nil
}

func (s *CloudGuardManagedListResourceCrud) Get() error {
	request := oci_cloud_guard.GetManagedListRequest{}

	tmp := s.D.Id()
	request.ManagedListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetManagedList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedList
	return nil
}

func (s *CloudGuardManagedListResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_guard.UpdateManagedListRequest{}

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

	if listItems, ok := s.D.GetOkExists("list_items"); ok {
		interfaces := listItems.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("list_items") {
			request.ListItems = tmp
		}
	}

	tmp := s.D.Id()
	request.ManagedListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateManagedList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedList
	return nil
}

func (s *CloudGuardManagedListResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteManagedListRequest{}

	tmp := s.D.Id()
	request.ManagedListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteManagedList(context.Background(), request)
	return err
}

func (s *CloudGuardManagedListResourceCrud) SetData() error {
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

	s.D.Set("feed_provider", s.Res.FeedProvider)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEditable != nil {
		s.D.Set("is_editable", *s.Res.IsEditable)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	s.D.Set("list_items", s.Res.ListItems)

	s.D.Set("list_type", s.Res.ListType)

	if s.Res.SourceManagedListId != nil {
		s.D.Set("source_managed_list_id", *s.Res.SourceManagedListId)
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

func ManagedListSummaryToMap(obj oci_cloud_guard.ManagedListSummary) map[string]interface{} {
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

	result["feed_provider"] = string(obj.FeedProvider)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEditable != nil {
		result["is_editable"] = bool(*obj.IsEditable)
	}

	if obj.LifecyleDetails != nil {
		result["lifecyle_details"] = string(*obj.LifecyleDetails)
	}

	result["list_items"] = obj.ListItems

	result["list_type"] = string(obj.ListType)

	if obj.SourceManagedListId != nil {
		result["source_managed_list_id"] = string(*obj.SourceManagedListId)
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

func (s *CloudGuardManagedListResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_guard.ChangeManagedListCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ManagedListId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.ChangeManagedListCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
