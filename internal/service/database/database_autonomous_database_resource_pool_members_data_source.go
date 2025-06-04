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

func DatabaseAutonomousDatabaseResourcePoolMembersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabaseResourcePoolMembers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_pool_member_collection": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDatabaseResourcePoolMembers(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseResourcePoolMembersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseResourcePoolMembersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListResourcePoolMembersResponse
}

func (s *DatabaseAutonomousDatabaseResourcePoolMembersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseResourcePoolMembersDataSourceCrud) Get() error {
	request := oci_database.ListResourcePoolMembersRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListResourcePoolMembers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourcePoolMembers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseResourcePoolMembersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseResourcePoolMembersDataSource-", DatabaseAutonomousDatabaseResourcePoolMembersDataSource(), s.D))
	resources := []map[string]interface{}{}
	autonomousDatabaseResourcePoolMember := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourcePoolMemberSummaryToMap(item))
	}
	autonomousDatabaseResourcePoolMember["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseAutonomousDatabaseResourcePoolMembersDataSource().Schema["resource_pool_member_collection"].Elem.(*schema.Resource).Schema)
		autonomousDatabaseResourcePoolMember["items"] = items
	}

	resources = append(resources, autonomousDatabaseResourcePoolMember)
	if err := s.D.Set("resource_pool_member_collection", resources); err != nil {
		return err
	}

	return nil
}

func ResourcePoolMemberSummaryToMap(obj oci_database.ResourcePoolMemberSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
