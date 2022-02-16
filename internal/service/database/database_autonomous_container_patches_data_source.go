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

func DatabaseAutonomousContainerPatchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousContainerPatches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_patches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
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
						"patch_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quarter": {
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
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"year": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousContainerPatches(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerPatchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerPatchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListContainerDatabasePatchesResponse
}

func (s *DatabaseAutonomousContainerPatchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerPatchesDataSourceCrud) Get() error {
	request := oci_database.ListContainerDatabasePatchesRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListContainerDatabasePatches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListContainerDatabasePatches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousContainerPatchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousContainerPatchesDataSource-", DatabaseAutonomousContainerPatchesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousContainerPatch := map[string]interface{}{}

		if r.Description != nil {
			autonomousContainerPatch["description"] = *r.Description
		}

		if r.Id != nil {
			autonomousContainerPatch["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			autonomousContainerPatch["lifecycle_details"] = *r.LifecycleDetails
		}

		autonomousContainerPatch["patch_model"] = r.PatchModel

		if r.Quarter != nil {
			autonomousContainerPatch["quarter"] = *r.Quarter
		}

		autonomousContainerPatch["state"] = r.LifecycleState

		if r.TimeReleased != nil {
			autonomousContainerPatch["time_released"] = r.TimeReleased.String()
		}

		if r.Type != nil {
			autonomousContainerPatch["type"] = *r.Type
		}

		if r.Version != nil {
			autonomousContainerPatch["version"] = *r.Version
		}

		if r.Year != nil {
			autonomousContainerPatch["year"] = *r.Year
		}

		resources = append(resources, autonomousContainerPatch)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousContainerPatchesDataSource().Schema["autonomous_patches"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_patches", resources); err != nil {
		return err
	}

	return nil
}
