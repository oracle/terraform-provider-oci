// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceGroup,
		Read:     readOsManagementHubManagedInstanceGroup,
		Update:   updateOsManagementHubManagedInstanceGroup,
		Delete:   deleteOsManagementHubManagedInstanceGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"arch_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"autonomous_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_data_collection_authorized": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
						"scheduled_job_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
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
			"location": {
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
			"notification_topic_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"software_source_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"is_managed_by_autonomous_linux": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"managed_instance_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pending_job_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_mandatory_for_autonomous_linux": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"software_source_type": {
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOsManagementHubManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsManagementHubManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OsManagementHubManagedInstanceGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceGroupClient
	Res                    *oci_os_management_hub.ManagedInstanceGroup
	DisableNotFoundRetries bool
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.ManagedInstanceGroupLifecycleStateCreating),
	}
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.ManagedInstanceGroupLifecycleStateActive),
	}
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.ManagedInstanceGroupLifecycleStateDeleting),
	}
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.ManagedInstanceGroupLifecycleStateDeleted),
	}
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) Create() error {
	request := oci_os_management_hub.CreateManagedInstanceGroupRequest{}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		request.ArchType = oci_os_management_hub.ArchTypeEnum(archType.(string))
	}

	if autonomousSettings, ok := s.D.GetOkExists("autonomous_settings"); ok {
		if tmpList := autonomousSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "autonomous_settings", 0)
			tmp, err := s.mapToUpdatableAutonomousSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AutonomousSettings = &tmp
		}
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

	if location, ok := s.D.GetOkExists("location"); ok {
		request.Location = oci_os_management_hub.ManagedInstanceLocationEnum(location.(string))
	}

	if managedInstanceIds, ok := s.D.GetOkExists("managed_instance_ids"); ok {
		interfaces := managedInstanceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("managed_instance_ids") {
			request.ManagedInstanceIds = tmp
		}
	}

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_os_management_hub.OsFamilyEnum(osFamily.(string))
	}

	if softwareSourceIds, ok := s.D.GetOkExists("software_source_ids"); ok {
		interfaces := softwareSourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("software_source_ids") {
			request.SoftwareSourceIds = tmp
		}
	}

	if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
		request.VendorName = oci_os_management_hub.VendorNameEnum(vendorName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.CreateManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstanceGroup
	return nil
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceGroupRequest{}

	tmp := s.D.Id()
	request.ManagedInstanceGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstanceGroup
	return nil
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeManagedInstanceGroupCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_os_management_hub.UpdateManagedInstanceGroupRequest{}

	isManagedByAutonomousLinux, _ := s.D.GetOkExists("is_managed_by_autonomous_linux")
	if autonomousSettings, ok := s.D.GetOkExists("autonomous_settings"); ok && isManagedByAutonomousLinux.(bool) {
		if tmpList := autonomousSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "autonomous_settings", 0)
			tmp, err := s.mapToUpdatableAutonomousSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AutonomousSettings = &tmp
		}
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

	tmp := s.D.Id()
	request.ManagedInstanceGroupId = &tmp

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstanceGroup
	return nil
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) Delete() error {
	request := oci_os_management_hub.DeleteManagedInstanceGroupRequest{}

	tmp := s.D.Id()
	request.ManagedInstanceGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DeleteManagedInstanceGroup(context.Background(), request)
	return err
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) SetData() error {
	s.D.Set("arch_type", s.Res.ArchType)

	if s.Res.AutonomousSettings != nil {
		s.D.Set("autonomous_settings", []interface{}{AutonomousSettingsToMap(s.Res.AutonomousSettings)})
	} else {
		s.D.Set("autonomous_settings", nil)
	}

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

	if s.Res.IsManagedByAutonomousLinux != nil {
		s.D.Set("is_managed_by_autonomous_linux", *s.Res.IsManagedByAutonomousLinux)
	}

	s.D.Set("location", s.Res.Location)

	if s.Res.ManagedInstanceCount != nil {
		s.D.Set("managed_instance_count", *s.Res.ManagedInstanceCount)
	}

	s.D.Set("managed_instance_ids", s.Res.ManagedInstanceIds)

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.PendingJobCount != nil {
		s.D.Set("pending_job_count", *s.Res.PendingJobCount)
	}

	softwareSources := []interface{}{}
	softwareSourceIds := []string{}
	for _, item := range s.Res.SoftwareSourceIds {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
		softwareSourceIds = append(softwareSourceIds, *item.Id)
	}
	s.D.Set("software_sources", softwareSources)

	s.D.Set("software_source_ids", softwareSourceIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	s.D.Set("vendor_name", s.Res.VendorName)

	return nil
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) ChangeManagedInstanceGroupCompartment() error {
	request := oci_os_management_hub.ChangeManagedInstanceGroupCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.ManagedInstanceGroupId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeManagedInstanceGroupCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func ManagedInstanceGroupSummaryToMap(obj oci_os_management_hub.ManagedInstanceGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["arch_type"] = string(obj.ArchType)

	if obj.AutonomousSettings != nil {
		result["autonomous_settings"] = []interface{}{AutonomousSettingsToMap(obj.AutonomousSettings)}
	}

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

	if obj.IsManagedByAutonomousLinux != nil {
		result["is_managed_by_autonomous_linux"] = bool(*obj.IsManagedByAutonomousLinux)
	}

	result["location"] = string(obj.Location)

	if obj.ManagedInstanceCount != nil {
		result["managed_instance_count"] = int(*obj.ManagedInstanceCount)
	}

	if obj.NotificationTopicId != nil {
		result["notification_topic_id"] = string(*obj.NotificationTopicId)
	}

	result["os_family"] = string(obj.OsFamily)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	result["vendor_name"] = string(obj.VendorName)

	return result
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) mapToUpdatableAutonomousSettings(fieldKeyFormat string) (oci_os_management_hub.UpdatableAutonomousSettings, error) {
	result := oci_os_management_hub.UpdatableAutonomousSettings{}

	if isDataCollectionAuthorized, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_data_collection_authorized")); ok {
		tmp := isDataCollectionAuthorized.(bool)
		result.IsDataCollectionAuthorized = &tmp
	}

	return result, nil
}

func (s *OsManagementHubManagedInstanceGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_os_management_hub.ChangeManagedInstanceGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ManagedInstanceGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeManagedInstanceGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
