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

func DatabaseVmClusterPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseVmClusterPatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"patches": {
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

func readDatabaseVmClusterPatches(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListVmClusterPatchesResponse
}

func (s *DatabaseVmClusterPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterPatchesDataSourceCrud) Get() error {
	request := oci_database.ListVmClusterPatchesRequest{}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListVmClusterPatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVmClusterPatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseVmClusterPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseVmClusterPatchesDataSource-", DatabaseVmClusterPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vmClusterPatch := map[string]interface{}{}

		vmClusterPatch["available_actions"] = r.AvailableActions

		if r.Description != nil {
			vmClusterPatch["description"] = *r.Description
		}

		if r.Id != nil {
			vmClusterPatch["id"] = *r.Id
		}

		vmClusterPatch["last_action"] = r.LastAction

		if r.LifecycleDetails != nil {
			vmClusterPatch["lifecycle_details"] = *r.LifecycleDetails
		}

		vmClusterPatch["state"] = r.LifecycleState

		if r.TimeReleased != nil {
			vmClusterPatch["time_released"] = r.TimeReleased.String()
		}

		if r.Version != nil {
			vmClusterPatch["version"] = *r.Version
		}

		resources = append(resources, vmClusterPatch)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseVmClusterPatchesDataSource().Schema["patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("patches", resources); err != nil {
		return err
	}

	return nil
}
