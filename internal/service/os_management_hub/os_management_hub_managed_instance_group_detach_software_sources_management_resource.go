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

func OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagement,
		Read:     readOsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagement,
		Delete:   deleteOsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"software_sources": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"work_request_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
		},
	}
}

func createOsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceGroupClient
	Res                    *oci_os_management_hub.GetManagedInstanceGroupResponse
	DisableNotFoundRetries bool
}

func (s *OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceGroupRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResourceCrud) Create() error {
	request := oci_os_management_hub.DetachSoftwareSourcesFromManagedInstanceGroupRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	if softwareSources, ok := s.D.GetOkExists("software_sources"); ok {
		interfaces := softwareSources.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("software_sources") {
			request.SoftwareSources = tmp
		}
	}

	if workRequestDetails, ok := s.D.GetOkExists("work_request_details"); ok {
		if tmpList := workRequestDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "work_request_details", 0)
			tmp, err := s.mapToWorkRequestDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.WorkRequestDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DetachSoftwareSourcesFromManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResourceCrud) SetData() error {
	return nil
}

func (s *OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResourceCrud) mapToWorkRequestDetails(fieldKeyFormat string) (oci_os_management_hub.WorkRequestDetails, error) {
	result := oci_os_management_hub.WorkRequestDetails{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}
