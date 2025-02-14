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

func OsManagementHubProfileAttachLifecycleStageManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubProfileAttachLifecycleStageManagement,
		Read:     readOsManagementHubProfileAttachLifecycleStageManagement,
		Delete:   deleteOsManagementHubProfileAttachLifecycleStageManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"lifecycle_stage_id": {
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

func createOsManagementHubProfileAttachLifecycleStageManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileAttachLifecycleStageManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubProfileAttachLifecycleStageManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubProfileAttachLifecycleStageManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubProfileAttachLifecycleStageManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.OnboardingClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OsManagementHubProfileAttachLifecycleStageManagementResourceCrud) ID() string {
	return *s.Res
}

func (s *OsManagementHubProfileAttachLifecycleStageManagementResourceCrud) Create() error {
	request := oci_os_management_hub.AttachLifecycleStageToProfileRequest{}

	if lifecycleStageId, ok := s.D.GetOkExists("lifecycle_stage_id"); ok {
		tmp := lifecycleStageId.(string)
		request.LifecycleStageId = &tmp
	}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.AttachLifecycleStageToProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = request.ProfileId
	return nil
}

func (s *OsManagementHubProfileAttachLifecycleStageManagementResourceCrud) SetData() error {
	return nil
}
