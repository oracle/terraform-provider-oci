// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubWindowsUpdateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsManagementHubWindowsUpdate,
		Schema: map[string]*schema.Schema{
			"windows_update_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"installable": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"installation_requirements": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_reboot_required_for_installation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"kb_article_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOsManagementHubWindowsUpdate(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubWindowsUpdateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubWindowsUpdateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.GetWindowsUpdateResponse
}

func (s *OsManagementHubWindowsUpdateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubWindowsUpdateDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetWindowsUpdateRequest{}

	if windowsUpdateId, ok := s.D.GetOkExists("windows_update_id"); ok {
		tmp := windowsUpdateId.(string)
		request.WindowsUpdateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetWindowsUpdate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubWindowsUpdateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubWindowsUpdateDataSource-", OsManagementHubWindowsUpdateDataSource(), s.D))

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("installable", s.Res.Installable)

	s.D.Set("installation_requirements", s.Res.InstallationRequirements)

	if s.Res.IsRebootRequiredForInstallation != nil {
		s.D.Set("is_reboot_required_for_installation", *s.Res.IsRebootRequiredForInstallation)
	}

	s.D.Set("kb_article_ids", s.Res.KbArticleIds)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SizeInBytes != nil {
		s.D.Set("size_in_bytes", strconv.FormatInt(*s.Res.SizeInBytes, 10))
	}

	if s.Res.UpdateId != nil {
		s.D.Set("update_id", *s.Res.UpdateId)
	}

	s.D.Set("update_type", s.Res.UpdateType)

	return nil
}
