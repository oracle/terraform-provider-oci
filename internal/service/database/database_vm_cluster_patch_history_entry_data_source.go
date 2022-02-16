// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseVmClusterPatchHistoryEntryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseVmClusterPatchHistoryEntry,
		Schema: map[string]*schema.Schema{
			"patch_history_entry_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseVmClusterPatchHistoryEntry(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterPatchHistoryEntryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterPatchHistoryEntryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetVmClusterPatchHistoryEntryResponse
}

func (s *DatabaseVmClusterPatchHistoryEntryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterPatchHistoryEntryDataSourceCrud) Get() error {
	request := oci_database.GetVmClusterPatchHistoryEntryRequest{}

	if patchHistoryEntryId, ok := s.D.GetOkExists("patch_history_entry_id"); ok {
		tmp := patchHistoryEntryId.(string)
		request.PatchHistoryEntryId = &tmp
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetVmClusterPatchHistoryEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseVmClusterPatchHistoryEntryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("action", s.Res.Action)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PatchId != nil {
		s.D.Set("patch_id", *s.Res.PatchId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
