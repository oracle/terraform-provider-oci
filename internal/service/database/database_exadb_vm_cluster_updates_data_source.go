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

func DatabaseExadbVmClusterUpdatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExadbVmClusterUpdates,
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
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exadb_vm_cluster_updates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"available_actions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_action": {
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
						"time_released": {
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

func readDatabaseExadbVmClusterUpdates(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterUpdatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExadbVmClusterUpdatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExadbVmClusterUpdatesResponse
}

func (s *DatabaseExadbVmClusterUpdatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadbVmClusterUpdatesDataSourceCrud) Get() error {
	request := oci_database.ListExadbVmClusterUpdatesRequest{}

	if exadbVmClusterId, ok := s.D.GetOkExists("exadb_vm_cluster_id"); ok {
		tmp := exadbVmClusterId.(string)
		request.ExadbVmClusterId = &tmp
	}

	if updateType, ok := s.D.GetOkExists("update_type"); ok {
		request.UpdateType = oci_database.ListExadbVmClusterUpdatesUpdateTypeEnum(updateType.(string))
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExadbVmClusterUpdates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExadbVmClusterUpdates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExadbVmClusterUpdatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExadbVmClusterUpdatesDataSource-", DatabaseExadbVmClusterUpdatesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		exadbVmClusterUpdate := map[string]interface{}{}

		exadbVmClusterUpdate["available_actions"] = r.AvailableActions

		if r.Description != nil {
			exadbVmClusterUpdate["description"] = *r.Description
		}

		if r.Id != nil {
			exadbVmClusterUpdate["id"] = *r.Id
		}

		exadbVmClusterUpdate["last_action"] = r.LastAction

		if r.LifecycleDetails != nil {
			exadbVmClusterUpdate["lifecycle_details"] = *r.LifecycleDetails
		}

		exadbVmClusterUpdate["state"] = r.LifecycleState

		if r.TimeReleased != nil {
			exadbVmClusterUpdate["time_released"] = r.TimeReleased.String()
		}

		exadbVmClusterUpdate["update_type"] = r.UpdateType

		if r.Version != nil {
			exadbVmClusterUpdate["version"] = *r.Version
		}

		resources = append(resources, exadbVmClusterUpdate)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExadbVmClusterUpdatesDataSource().Schema["exadb_vm_cluster_updates"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("exadb_vm_cluster_updates", resources); err != nil {
		return err
	}

	return nil
}
