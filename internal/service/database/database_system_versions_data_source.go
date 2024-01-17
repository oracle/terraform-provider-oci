// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseSystemVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseSystemVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"gi_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
			},
			"system_version_collection": {
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
									"gi_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"shape": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_versions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func readDatabaseSystemVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSystemVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSystemVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListSystemVersionsResponse
}

func (s *DatabaseSystemVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSystemVersionsDataSourceCrud) Get() error {
	request := oci_database.ListSystemVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if giVersion, ok := s.D.GetOkExists("gi_version"); ok {
		tmp := giVersion.(string)
		request.GiVersion = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListSystemVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSystemVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseSystemVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseSystemVersionsDataSource-", DatabaseSystemVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	systemVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SystemVersionSummaryToMap(item))
	}
	systemVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseSystemVersionsDataSource().Schema["system_version_collection"].Elem.(*schema.Resource).Schema)
		systemVersion["items"] = items
	}

	resources = append(resources, systemVersion)
	if err := s.D.Set("system_version_collection", resources); err != nil {
		return err
	}

	return nil
}

func SystemVersionSummaryToMap(obj oci_database.SystemVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GiVersion != nil {
		result["gi_version"] = string(*obj.GiVersion)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	result["system_versions"] = obj.SystemVersions
	result["system_versions"] = obj.SystemVersions

	return result
}
