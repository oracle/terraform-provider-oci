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

func BdsBdsInstanceNodeBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBdsBdsInstanceNodeBackups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
						"id": {
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
				},
			},
		},
	}
}

func readBdsBdsInstanceNodeBackups(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceNodeBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceNodeBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListNodeBackupsResponse
}

func (s *BdsBdsInstanceNodeBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceNodeBackupsDataSourceCrud) Get() error {
	request := oci_bds.ListNodeBackupsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if nodeHostName, ok := s.D.GetOkExists("node_host_name"); ok {
		tmp := nodeHostName.(string)
		request.NodeHostName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.NodeBackupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListNodeBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNodeBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceNodeBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceNodeBackupsDataSource-", BdsBdsInstanceNodeBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bdsInstanceNodeBackup := map[string]interface{}{}

		bdsInstanceNodeBackup["backup_trigger_type"] = r.BackupTriggerType

		bdsInstanceNodeBackup["backup_type"] = r.BackupType

		if r.DisplayName != nil {
			bdsInstanceNodeBackup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			bdsInstanceNodeBackup["id"] = *r.Id
		}

		if r.NodeHostName != nil {
			bdsInstanceNodeBackup["node_host_name"] = *r.NodeHostName
		}

		if r.NodeInstanceId != nil {
			bdsInstanceNodeBackup["node_instance_id"] = *r.NodeInstanceId
		}

		bdsInstanceNodeBackup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bdsInstanceNodeBackup["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, bdsInstanceNodeBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BdsBdsInstanceNodeBackupsDataSource().Schema["node_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("node_backups", resources); err != nil {
		return err
	}

	return nil
}
