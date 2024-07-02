// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExadbVmClusterUpdateHistoryEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseExadbVmClusterUpdateHistoryEntry,
		Schema: map[string]*schema.Schema{
			"exadb_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"update_history_entry_id": {
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseExadbVmClusterUpdateHistoryEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterUpdateHistoryEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExadbVmClusterUpdateHistoryEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExadbVmClusterUpdateHistoryEntryResponse
}

func (s *DatabaseExadbVmClusterUpdateHistoryEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadbVmClusterUpdateHistoryEntryDataSourceCrud) Get() error {
	request := oci_database.GetExadbVmClusterUpdateHistoryEntryRequest{}

	if exadbVmClusterId, ok := s.D.GetOkExists("exadb_vm_cluster_id"); ok {
		tmp := exadbVmClusterId.(string)
		request.ExadbVmClusterId = &tmp
	}

	if updateHistoryEntryId, ok := s.D.GetOkExists("update_history_entry_id"); ok {
		tmp := updateHistoryEntryId.(string)
		request.UpdateHistoryEntryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExadbVmClusterUpdateHistoryEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExadbVmClusterUpdateHistoryEntryDataSourceCrud) SetData() error {
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

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
