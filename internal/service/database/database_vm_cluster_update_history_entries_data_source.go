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

func DatabaseVmClusterUpdateHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseVmClusterUpdateHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_cluster_update_history_entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
				},
			},
		},
	}
}

func readDatabaseVmClusterUpdateHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterUpdateHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterUpdateHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListVmClusterUpdateHistoryEntriesResponse
}

func (s *DatabaseVmClusterUpdateHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterUpdateHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListVmClusterUpdateHistoryEntriesRequest{}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum(state.(string))
	}

	if updateType, ok := s.D.GetOkExists("update_type"); ok {
		request.UpdateType = oci_database.ListVmClusterUpdateHistoryEntriesUpdateTypeEnum(updateType.(string))
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListVmClusterUpdateHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVmClusterUpdateHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseVmClusterUpdateHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseVmClusterUpdateHistoryEntriesDataSource-", DatabaseVmClusterUpdateHistoryEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vmClusterUpdateHistoryEntry := map[string]interface{}{}

		if r.Id != nil {
			vmClusterUpdateHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			vmClusterUpdateHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		vmClusterUpdateHistoryEntry["state"] = r.LifecycleState

		if r.TimeCompleted != nil {
			vmClusterUpdateHistoryEntry["time_completed"] = r.TimeCompleted.String()
		}

		if r.TimeStarted != nil {
			vmClusterUpdateHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		vmClusterUpdateHistoryEntry["update_action"] = r.UpdateAction

		if r.UpdateId != nil {
			vmClusterUpdateHistoryEntry["update_id"] = *r.UpdateId
		}

		vmClusterUpdateHistoryEntry["update_type"] = r.UpdateType

		resources = append(resources, vmClusterUpdateHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseVmClusterUpdateHistoryEntriesDataSource().Schema["vm_cluster_update_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vm_cluster_update_history_entries", resources); err != nil {
		return err
	}

	return nil
}
