// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"
)

func OsmanagementManagedInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsmanagementManagedInstance,
		Read:     readOsmanagementManagedInstance,
		Update:   updateOsmanagementManagedInstance,
		Delete:   deleteOsmanagementManagedInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"is_data_collection_authorized": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"notification_topic_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"autonomous": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_auto_update_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"bug_updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"child_software_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enhancement_updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_reboot_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ksplice_effective_kernel_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_boot": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_checkin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_groups": {
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
			"os_family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_kernel_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"other_updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"parent_software_source": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"scheduled_job_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"security_updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updates_available": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"work_request_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createOsmanagementManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readOsmanagementManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateOsmanagementManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsmanagementManagedInstance(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsmanagementManagedInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_osmanagement.OsManagementClient
	Res                    *oci_osmanagement.ManagedInstance
	DisableNotFoundRetries bool
}

func (s *OsmanagementManagedInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsmanagementManagedInstanceResourceCrud) Create() error {
	request := oci_osmanagement.UpdateManagedInstanceRequest{}

	if isDataCollectionAuthorized, ok := s.D.GetOkExists("is_data_collection_authorized"); ok {
		tmp := isDataCollectionAuthorized.(bool)
		request.IsDataCollectionAuthorized = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.UpdateManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstance
	return nil
}

func (s *OsmanagementManagedInstanceResourceCrud) Get() error {
	request := oci_osmanagement.GetManagedInstanceRequest{}

	tmp := s.D.Id()
	request.ManagedInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstance
	return nil
}

func (s *OsmanagementManagedInstanceResourceCrud) Update() error {
	request := oci_osmanagement.UpdateManagedInstanceRequest{}

	if isDataCollectionAuthorized, ok := s.D.GetOkExists("is_data_collection_authorized"); ok {
		tmp := isDataCollectionAuthorized.(bool)
		request.IsDataCollectionAuthorized = &tmp
	}

	tmp := s.D.Id()
	request.ManagedInstanceId = &tmp

	if notificationTopicId, ok := s.D.GetOkExists("notification_topic_id"); ok {
		tmp := notificationTopicId.(string)
		request.NotificationTopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osmanagement")

	response, err := s.Client.UpdateManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagedInstance
	return nil
}

func (s *OsmanagementManagedInstanceResourceCrud) SetData() error {
	if s.Res.Autonomous != nil {
		s.D.Set("autonomous", []interface{}{AutonomousSettingsToMap(s.Res.Autonomous)})
	} else {
		s.D.Set("autonomous", nil)
	}

	if s.Res.BugUpdatesAvailable != nil {
		s.D.Set("bug_updates_available", *s.Res.BugUpdatesAvailable)
	}

	childSoftwareSources := []interface{}{}
	for _, item := range s.Res.ChildSoftwareSources {
		childSoftwareSources = append(childSoftwareSources, SoftwareSourceIdToMap(&item))
	}
	s.D.Set("child_software_sources", childSoftwareSources)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnhancementUpdatesAvailable != nil {
		s.D.Set("enhancement_updates_available", *s.Res.EnhancementUpdatesAvailable)
	}

	if s.Res.IsDataCollectionAuthorized != nil {
		s.D.Set("is_data_collection_authorized", *s.Res.IsDataCollectionAuthorized)
	}

	if s.Res.IsRebootRequired != nil {
		s.D.Set("is_reboot_required", *s.Res.IsRebootRequired)
	}

	if s.Res.KspliceEffectiveKernelVersion != nil {
		s.D.Set("ksplice_effective_kernel_version", *s.Res.KspliceEffectiveKernelVersion)
	}

	if s.Res.LastBoot != nil {
		s.D.Set("last_boot", *s.Res.LastBoot)
	}

	if s.Res.LastCheckin != nil {
		s.D.Set("last_checkin", *s.Res.LastCheckin)
	}

	managedInstanceGroups := []interface{}{}
	for _, item := range s.Res.ManagedInstanceGroups {
		managedInstanceGroups = append(managedInstanceGroups, IdToMap(item))
	}
	s.D.Set("managed_instance_groups", managedInstanceGroups)

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.OsKernelVersion != nil {
		s.D.Set("os_kernel_version", *s.Res.OsKernelVersion)
	}

	if s.Res.OsName != nil {
		s.D.Set("os_name", *s.Res.OsName)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	if s.Res.OtherUpdatesAvailable != nil {
		s.D.Set("other_updates_available", *s.Res.OtherUpdatesAvailable)
	}

	if s.Res.ParentSoftwareSource != nil {
		s.D.Set("parent_software_source", []interface{}{SoftwareSourceIdToMap(s.Res.ParentSoftwareSource)})
	} else {
		s.D.Set("parent_software_source", nil)
	}

	if s.Res.ScheduledJobCount != nil {
		s.D.Set("scheduled_job_count", *s.Res.ScheduledJobCount)
	}

	if s.Res.SecurityUpdatesAvailable != nil {
		s.D.Set("security_updates_available", *s.Res.SecurityUpdatesAvailable)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.UpdatesAvailable != nil {
		s.D.Set("updates_available", *s.Res.UpdatesAvailable)
	}

	if s.Res.WorkRequestCount != nil {
		s.D.Set("work_request_count", *s.Res.WorkRequestCount)
	}

	return nil
}

func AutonomousSettingsToMap(obj *oci_osmanagement.AutonomousSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsAutoUpdateEnabled != nil {
		result["is_auto_update_enabled"] = bool(*obj.IsAutoUpdateEnabled)
	}

	return result
}

func SoftwareSourceIdToMap(obj *oci_osmanagement.SoftwareSourceId) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
