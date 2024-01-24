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

func DatabaseAutonomousDatabaseRefreshableClonesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabaseRefreshableClones,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"refreshable_clone_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
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
									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDatabaseRefreshableClones(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseRefreshableClonesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseRefreshableClonesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseRefreshableClonesResponse
}

func (s *DatabaseAutonomousDatabaseRefreshableClonesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseRefreshableClonesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseRefreshableClonesRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseRefreshableClones(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabaseRefreshableClones(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseRefreshableClonesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseRefreshableClonesDataSource-", DatabaseAutonomousDatabaseRefreshableClonesDataSource(), s.D))
	resources := []map[string]interface{}{}
	autonomousDatabaseRefreshableClone := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RefreshableCloneSummaryToMap(item))
	}
	autonomousDatabaseRefreshableClone["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseAutonomousDatabaseRefreshableClonesDataSource().Schema["refreshable_clone_collection"].Elem.(*schema.Resource).Schema)
		autonomousDatabaseRefreshableClone["items"] = items
	}

	resources = append(resources, autonomousDatabaseRefreshableClone)
	if err := s.D.Set("refreshable_clone_collection", resources); err != nil {
		return err
	}

	return nil
}

func RefreshableCloneSummaryToMap(obj oci_database.RefreshableCloneSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	return result
}
