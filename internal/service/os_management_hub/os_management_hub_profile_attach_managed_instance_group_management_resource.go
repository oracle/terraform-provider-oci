// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubProfileAttachManagedInstanceGroupManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubProfileAttachManagedInstanceGroupManagement,
		Read:     readOsManagementHubProfileAttachManagedInstanceGroupManagement,
		Delete:   deleteOsManagementHubProfileAttachManagedInstanceGroupManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"profile_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createOsManagementHubProfileAttachManagedInstanceGroupManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileAttachManagedInstanceGroupManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubProfileAttachManagedInstanceGroupManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubProfileAttachManagedInstanceGroupManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubProfileAttachManagedInstanceGroupManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.OnboardingClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OsManagementHubProfileAttachManagedInstanceGroupManagementResourceCrud) ID() string {
	return *s.Res
}

func (s *OsManagementHubProfileAttachManagedInstanceGroupManagementResourceCrud) Create() error {
	request := oci_os_management_hub.AttachManagedInstanceGroupToProfileRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.AttachManagedInstanceGroupToProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = request.ProfileId
	return nil
}

func (s *OsManagementHubProfileAttachManagedInstanceGroupManagementResourceCrud) SetData() error {
	return nil
}
