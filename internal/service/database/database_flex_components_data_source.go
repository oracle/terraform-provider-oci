// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseFlexComponentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseFlexComponents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"flex_component_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"available_core_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"available_db_storage_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"minimum_core_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
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

func readDatabaseFlexComponents(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseFlexComponentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseFlexComponentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListFlexComponentsResponse
}

func (s *DatabaseFlexComponentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseFlexComponentsDataSourceCrud) Get() error {
	request := oci_database.ListFlexComponentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListFlexComponents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFlexComponents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseFlexComponentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseFlexComponentsDataSource-", DatabaseFlexComponentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	flexComponent := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FlexComponentSummaryToMap(item))
	}
	flexComponent["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseFlexComponentsDataSource().Schema["flex_component_collection"].Elem.(*schema.Resource).Schema)
		flexComponent["items"] = items
	}

	resources = append(resources, flexComponent)
	if err := s.D.Set("flex_component_collection", resources); err != nil {
		return err
	}

	return nil
}

func FlexComponentSummaryToMap(obj oci_database.FlexComponentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableCoreCount != nil {
		result["available_core_count"] = int(*obj.AvailableCoreCount)
	}

	if obj.AvailableDbStorageInGBs != nil {
		result["available_db_storage_in_gbs"] = int(*obj.AvailableDbStorageInGBs)
	}

	if obj.MinimumCoreCount != nil {
		result["minimum_core_count"] = int(*obj.MinimumCoreCount)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
