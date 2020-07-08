// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterDatasource("oci_database_db_system_shapes", DatabaseDbSystemShapesDataSource())
}

func DatabaseDbSystemShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbSystemShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_shapes": {
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
						"available_data_storage_in_tbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_db_node_storage_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_memory_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"core_count_increment": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_node_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_core_count_per_node": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_data_storage_in_tbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_db_node_storage_per_node_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_memory_per_node_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"minimum_core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"minimum_node_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:       schema.TypeString,
							Computed:   true,
							Deprecated: FieldDeprecatedForAnother("shape", "name"),
						},
						"shape_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbSystemShapes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseDbSystemShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemShapesResponse
}

func (s *DatabaseDbSystemShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemShapesDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystemShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystemShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbSystemShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemShape := map[string]interface{}{}

		if r.AvailableCoreCount != nil {
			dbSystemShape["available_core_count"] = *r.AvailableCoreCount
		}

		if r.AvailableDataStorageInTBs != nil {
			dbSystemShape["available_data_storage_in_tbs"] = *r.AvailableDataStorageInTBs
		}

		if r.AvailableDbNodeStorageInGBs != nil {
			dbSystemShape["available_db_node_storage_in_gbs"] = *r.AvailableDbNodeStorageInGBs
		}

		if r.AvailableMemoryInGBs != nil {
			dbSystemShape["available_memory_in_gbs"] = *r.AvailableMemoryInGBs
		}

		if r.CoreCountIncrement != nil {
			dbSystemShape["core_count_increment"] = *r.CoreCountIncrement
		}

		if r.MaximumNodeCount != nil {
			dbSystemShape["maximum_node_count"] = *r.MaximumNodeCount
		}

		if r.MinCoreCountPerNode != nil {
			dbSystemShape["min_core_count_per_node"] = *r.MinCoreCountPerNode
		}

		if r.MinDataStorageInTBs != nil {
			dbSystemShape["min_data_storage_in_tbs"] = *r.MinDataStorageInTBs
		}

		if r.MinDbNodeStoragePerNodeInGBs != nil {
			dbSystemShape["min_db_node_storage_per_node_in_gbs"] = *r.MinDbNodeStoragePerNodeInGBs
		}

		if r.MinMemoryPerNodeInGBs != nil {
			dbSystemShape["min_memory_per_node_in_gbs"] = *r.MinMemoryPerNodeInGBs
		}

		if r.MinimumCoreCount != nil {
			dbSystemShape["minimum_core_count"] = *r.MinimumCoreCount
		}

		if r.MinimumNodeCount != nil {
			dbSystemShape["minimum_node_count"] = *r.MinimumNodeCount
		}

		if r.Name != nil {
			dbSystemShape["name"] = *r.Name
		}

		if r.Shape != nil {
			dbSystemShape["shape"] = *r.Shape
		}

		if r.ShapeFamily != nil {
			dbSystemShape["shape_family"] = *r.ShapeFamily
		}

		resources = append(resources, dbSystemShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseDbSystemShapesDataSource().Schema["db_system_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_system_shapes", resources); err != nil {
		return err
	}

	return nil
}
