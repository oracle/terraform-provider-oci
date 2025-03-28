// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"io"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwareSourceManifestDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["software_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubSoftwareSourceManifestResource(), fieldMap, readSingularOsManagementHubSoftwareSourceManifest)
}

func readSingularOsManagementHubSoftwareSourceManifest(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceManifestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceManifestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.GetSoftwareSourceManifestResponse
	Res2   *oci_os_management_hub.GetSoftwareSourceResponse
}

func (s *OsManagementHubSoftwareSourceManifestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceManifestDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetSoftwareSourceManifestRequest{}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetSoftwareSourceManifest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	request2 := oci_os_management_hub.GetSoftwareSourceRequest{}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request2.SoftwareSourceId = &tmp
	}

	request2.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response2, err := s.Client.GetSoftwareSource(context.Background(), request2)
	if err != nil {
		return err
	}

	s.Res2 = &response2

	return nil
}

func (s *OsManagementHubSoftwareSourceManifestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if s.Res2 == nil {
		return nil
	}

	s.D.SetId(*s.Res2.GetId())
	s.D.Set("software_source_id", *s.Res2.GetId())

	content, err := io.ReadAll(s.Res.Content)
	if err != nil {
		return nil
	}
	s.D.Set("content", string(content))

	return nil
}
