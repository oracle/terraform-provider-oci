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

func DatabaseSystemVersionMinorVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseSystemVersionMinorVersions,
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
			"is_latest": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"major_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_version_minor_version_collection": {
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
									"version": {
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

func readDatabaseSystemVersionMinorVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSystemVersionMinorVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSystemVersionMinorVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListSystemVersionMinorVersionsResponse
}

func (s *DatabaseSystemVersionMinorVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSystemVersionMinorVersionsDataSourceCrud) Get() error {
	request := oci_database.ListSystemVersionMinorVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if giVersion, ok := s.D.GetOkExists("gi_version"); ok {
		tmp := giVersion.(string)
		request.GiVersion = &tmp
	}

	if isLatest, ok := s.D.GetOkExists("is_latest"); ok {
		tmp := isLatest.(bool)
		request.IsLatest = &tmp
	}

	if majorVersion, ok := s.D.GetOkExists("major_version"); ok {
		tmp := majorVersion.(string)
		request.MajorVersion = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListSystemVersionMinorVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSystemVersionMinorVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseSystemVersionMinorVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseSystemVersionMinorVersionsDataSource-", DatabaseSystemVersionMinorVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	systemVersionMinorVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SystemVersionMinorVersionSummaryToMap(item))
	}
	systemVersionMinorVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseSystemVersionMinorVersionsDataSource().Schema["system_version_minor_version_collection"].Elem.(*schema.Resource).Schema)
		systemVersionMinorVersion["items"] = items
	}

	resources = append(resources, systemVersionMinorVersion)
	if err := s.D.Set("system_version_minor_version_collection", resources); err != nil {
		return err
	}

	return nil
}

func SystemVersionMinorVersionSummaryToMap(obj oci_database.SystemVersionMinorVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
