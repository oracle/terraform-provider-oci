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

func OsManagementHubProfileDetachSoftwareSourcesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubProfileDetachSoftwareSourcesManagement,
		Read:     readOsManagementHubProfileDetachSoftwareSourcesManagement,
		Delete:   deleteOsManagementHubProfileDetachSoftwareSourcesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"profile_id": {
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

			// Computed
		},
	}
}

func createOsManagementHubProfileDetachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileDetachSoftwareSourcesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubProfileDetachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubProfileDetachSoftwareSourcesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubProfileDetachSoftwareSourcesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.OnboardingClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OsManagementHubProfileDetachSoftwareSourcesManagementResourceCrud) ID() string {
	return *s.Res
}

func (s *OsManagementHubProfileDetachSoftwareSourcesManagementResourceCrud) Create() error {
	request := oci_os_management_hub.DetachSoftwareSourcesFromProfileRequest{}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DetachSoftwareSourcesFromProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = request.ProfileId
	return nil
}

func (s *OsManagementHubProfileDetachSoftwareSourcesManagementResourceCrud) SetData() error {
	return nil
}
