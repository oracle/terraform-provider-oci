// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsWlsDomainServerBackupDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWlmsWlsDomainServerBackup,
		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"backup_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularWlmsWlsDomainServerBackup(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainServerBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainServerBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.GetWlsDomainServerBackupResponse
}

func (s *WlmsWlsDomainServerBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainServerBackupDataSourceCrud) Get() error {
	request := oci_wlms.GetWlsDomainServerBackupRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	if serverId, ok := s.D.GetOkExists("server_id"); ok {
		tmp := serverId.(string)
		request.ServerId = &tmp
	}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.GetWlsDomainServerBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WlmsWlsDomainServerBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BackupLocation != nil {
		s.D.Set("backup_location", *s.Res.BackupLocation)
	}

	s.D.Set("content_type", s.Res.ContentType)

	if s.Res.ManagedInstanceId != nil {
		s.D.Set("managed_instance_id", *s.Res.ManagedInstanceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
