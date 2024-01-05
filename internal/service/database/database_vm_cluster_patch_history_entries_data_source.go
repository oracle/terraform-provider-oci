// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseVmClusterPatchHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseVmClusterPatchHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patch_history_entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
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
				},
			},
		},
	}
}

func readDatabaseVmClusterPatchHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterPatchHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterPatchHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListVmClusterPatchHistoryEntriesResponse
}

func (s *DatabaseVmClusterPatchHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterPatchHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListVmClusterPatchHistoryEntriesRequest{}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListVmClusterPatchHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVmClusterPatchHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseVmClusterPatchHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseVmClusterPatchHistoryEntriesDataSource-", DatabaseVmClusterPatchHistoryEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vmClusterPatchHistoryEntry := map[string]interface{}{}

		vmClusterPatchHistoryEntry["action"] = r.Action

		if r.Id != nil {
			vmClusterPatchHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			vmClusterPatchHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PatchId != nil {
			vmClusterPatchHistoryEntry["patch_id"] = *r.PatchId
		}

		vmClusterPatchHistoryEntry["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			vmClusterPatchHistoryEntry["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			vmClusterPatchHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, vmClusterPatchHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseVmClusterPatchHistoryEntriesDataSource().Schema["patch_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patch_history_entries", resources); err != nil {
		return err
	}

	return nil
}
