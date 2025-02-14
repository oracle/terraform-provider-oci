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

func OsManagementHubSoftwareSourceGenerateMetadataManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubSoftwareSourceGenerateMetadataManagement,
		Read:     readOsManagementHubSoftwareSourceGenerateMetadataManagement,
		Delete:   deleteOsManagementHubSoftwareSourceGenerateMetadataManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"software_source_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createOsManagementHubSoftwareSourceGenerateMetadataManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceGenerateMetadataManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubSoftwareSourceGenerateMetadataManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubSoftwareSourceGenerateMetadataManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubSoftwareSourceGenerateMetadataManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.SoftwareSourceClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OsManagementHubSoftwareSourceGenerateMetadataManagementResourceCrud) ID() string {
	return *s.Res
}

func (s *OsManagementHubSoftwareSourceGenerateMetadataManagementResourceCrud) Create() error {
	request := oci_os_management_hub.SoftwareSourceGenerateMetadataRequest{}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.SoftwareSourceGenerateMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = request.SoftwareSourceId
	return nil
}

func (s *OsManagementHubSoftwareSourceGenerateMetadataManagementResourceCrud) SetData() error {
	return nil
}
