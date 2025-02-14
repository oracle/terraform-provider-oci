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

func OsManagementHubProfileAttachManagementStationManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubProfileAttachManagementStationManagement,
		Read:     readOsManagementHubProfileAttachManagementStationManagement,
		Delete:   deleteOsManagementHubProfileAttachManagementStationManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"management_station_id": {
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

func createOsManagementHubProfileAttachManagementStationManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileAttachManagementStationManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubProfileAttachManagementStationManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubProfileAttachManagementStationManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubProfileAttachManagementStationManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.OnboardingClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OsManagementHubProfileAttachManagementStationManagementResourceCrud) ID() string {
	return *s.Res
}

func (s *OsManagementHubProfileAttachManagementStationManagementResourceCrud) Create() error {
	request := oci_os_management_hub.AttachManagementStationToProfileRequest{}

	if managementStationId, ok := s.D.GetOkExists("management_station_id"); ok {
		tmp := managementStationId.(string)
		request.ManagementStationId = &tmp
	}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.AttachManagementStationToProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = request.ProfileId
	return nil
}

func (s *OsManagementHubProfileAttachManagementStationManagementResourceCrud) SetData() error {
	return nil
}
