// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceNodeBackupDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularBdsBdsInstanceNodeBackup,
		Schema: map[string]*schema.Schema{
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_backup_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"backup_trigger_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_backup_config_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularBdsBdsInstanceNodeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceNodeBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetNodeBackupResponse
}

func (s *BdsBdsInstanceNodeBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceNodeBackupDataSourceCrud) Get() error {
	request := oci_bds.GetNodeBackupRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if nodeBackupId, ok := s.D.GetOkExists("node_backup_id"); ok {
		tmp := nodeBackupId.(string)
		request.NodeBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetNodeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceNodeBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("backup_trigger_type", s.Res.BackupTriggerType)

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.NodeBackupConfigId != nil {
		s.D.Set("node_backup_config_id", *s.Res.NodeBackupConfigId)
	}

	if s.Res.NodeHostName != nil {
		s.D.Set("node_host_name", *s.Res.NodeHostName)
	}

	if s.Res.NodeInstanceId != nil {
		s.D.Set("node_instance_id", *s.Res.NodeInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
