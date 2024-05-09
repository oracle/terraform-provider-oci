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

func OsManagementHubManagedInstanceAttachProfileManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceAttachProfileManagement,
		Read:     readOsManagementHubManagedInstanceAttachProfileManagement,
		Delete:   deleteOsManagementHubManagedInstanceAttachProfileManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_id": {
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

func createOsManagementHubManagedInstanceAttachProfileManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceAttachProfileManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceAttachProfileManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubManagedInstanceAttachProfileManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubManagedInstanceAttachProfileManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceClient
	Res                    *oci_os_management_hub.GetManagedInstanceResponse
	DisableNotFoundRetries bool
}

func (s *OsManagementHubManagedInstanceAttachProfileManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceAttachProfileManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceAttachProfileManagementResourceCrud) Create() error {
	request := oci_os_management_hub.AttachProfileToManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.AttachProfileToManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *OsManagementHubManagedInstanceAttachProfileManagementResourceCrud) SetData() error {
	return nil
}
