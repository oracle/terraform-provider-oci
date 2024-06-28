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

func DatabaseExadbVmClusterUpdateHistoryEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExadbVmClusterUpdateHistoryEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"exadb_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"update_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exadb_vm_cluster_update_history_entries": {
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
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseExadbVmClusterUpdateHistoryEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterUpdateHistoryEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExadbVmClusterUpdateHistoryEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExadbVmClusterUpdateHistoryEntriesResponse
}

func (s *DatabaseExadbVmClusterUpdateHistoryEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadbVmClusterUpdateHistoryEntriesDataSourceCrud) Get() error {
	request := oci_database.ListExadbVmClusterUpdateHistoryEntriesRequest{}

	if exadbVmClusterId, ok := s.D.GetOkExists("exadb_vm_cluster_id"); ok {
		tmp := exadbVmClusterId.(string)
		request.ExadbVmClusterId = &tmp
	}

	if updateType, ok := s.D.GetOkExists("update_type"); ok {
		request.UpdateType = oci_database.ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum(updateType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExadbVmClusterUpdateHistoryEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExadbVmClusterUpdateHistoryEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExadbVmClusterUpdateHistoryEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExadbVmClusterUpdateHistoryEntriesDataSource-", DatabaseExadbVmClusterUpdateHistoryEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		exadbVmClusterUpdateHistoryEntry := map[string]interface{}{}

		if r.Id != nil {
			exadbVmClusterUpdateHistoryEntry["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			exadbVmClusterUpdateHistoryEntry["lifecycle_details"] = *r.LifecycleDetails
		}

		exadbVmClusterUpdateHistoryEntry["state"] = r.LifecycleState

		if r.TimeCompleted != nil {
			exadbVmClusterUpdateHistoryEntry["time_completed"] = r.TimeCompleted.String()
		}

		if r.TimeStarted != nil {
			exadbVmClusterUpdateHistoryEntry["time_started"] = r.TimeStarted.String()
		}

		exadbVmClusterUpdateHistoryEntry["update_action"] = r.UpdateAction

		if r.UpdateId != nil {
			exadbVmClusterUpdateHistoryEntry["update_id"] = *r.UpdateId
		}

		exadbVmClusterUpdateHistoryEntry["update_type"] = r.UpdateType

		if r.Version != nil {
			exadbVmClusterUpdateHistoryEntry["version"] = *r.Version
		}

		resources = append(resources, exadbVmClusterUpdateHistoryEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExadbVmClusterUpdateHistoryEntriesDataSource().Schema["exadb_vm_cluster_update_history_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("exadb_vm_cluster_update_history_entries", resources); err != nil {
		return err
	}

	return nil
}
