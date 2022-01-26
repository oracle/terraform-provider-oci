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

func DatabaseVmClusterUpdatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseVmClusterUpdates,
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
			"vm_cluster_updates": {
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

func readDatabaseVmClusterUpdates(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterUpdatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterUpdatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListVmClusterUpdatesResponse
}

func (s *DatabaseVmClusterUpdatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterUpdatesDataSourceCrud) Get() error {
	request := oci_database.ListVmClusterUpdatesRequest{}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.VmClusterUpdateSummaryLifecycleStateEnum(state.(string))
	}

	if updateType, ok := s.D.GetOkExists("update_type"); ok {
		request.UpdateType = oci_database.ListVmClusterUpdatesUpdateTypeEnum(updateType.(string))
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListVmClusterUpdates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVmClusterUpdates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseVmClusterUpdatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseVmClusterUpdatesDataSource-", DatabaseVmClusterUpdatesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vmClusterUpdate := map[string]interface{}{}

		vmClusterUpdate["available_actions"] = r.AvailableActions

		if r.Description != nil {
			vmClusterUpdate["description"] = *r.Description
		}

		if r.Id != nil {
			vmClusterUpdate["id"] = *r.Id
		}

		vmClusterUpdate["last_action"] = r.LastAction

		if r.LifecycleDetails != nil {
			vmClusterUpdate["lifecycle_details"] = *r.LifecycleDetails
		}

		vmClusterUpdate["state"] = r.LifecycleState

		if r.TimeReleased != nil {
			vmClusterUpdate["time_released"] = r.TimeReleased.String()
		}

		vmClusterUpdate["update_type"] = r.UpdateType

		if r.Version != nil {
			vmClusterUpdate["version"] = *r.Version
		}

		resources = append(resources, vmClusterUpdate)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseVmClusterUpdatesDataSource().Schema["vm_cluster_updates"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vm_cluster_updates", resources); err != nil {
		return err
	}

	return nil
}
