// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DatabaseVmClusterUpdateHistoryEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseVmClusterUpdateHistoryEntry,
		Schema: map[string]*schema.Schema{
			"update_history_entry_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_completed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_action": {
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

func readSingularDatabaseVmClusterUpdateHistoryEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterUpdateHistoryEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterUpdateHistoryEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetVmClusterUpdateHistoryEntryResponse
}

func (s *DatabaseVmClusterUpdateHistoryEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterUpdateHistoryEntryDataSourceCrud) Get() error {
	request := oci_database.GetVmClusterUpdateHistoryEntryRequest{}

	if updateHistoryEntryId, ok := s.D.GetOkExists("update_history_entry_id"); ok {
		tmp := updateHistoryEntryId.(string)
		request.UpdateHistoryEntryId = &tmp
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetVmClusterUpdateHistoryEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseVmClusterUpdateHistoryEntryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCompleted != nil {
		s.D.Set("time_completed", s.Res.TimeCompleted.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	s.D.Set("update_action", s.Res.UpdateAction)

	if s.Res.UpdateId != nil {
		s.D.Set("update_id", *s.Res.UpdateId)
	}

	s.D.Set("update_type", s.Res.UpdateType)

	return nil
}
