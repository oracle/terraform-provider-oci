// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseDbSystemShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbSystemShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
						"available_core_count_per_node": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_data_storage_in_tbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"available_data_storage_per_server_in_tbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"available_db_node_per_node_in_gbs": {
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
						"available_memory_per_node_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"core_count_increment": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_storage_count": {
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
						"min_storage_count": {
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
							Deprecated: tfresource.FieldDeprecatedForAnother("shape", "name"),
						},
						"shape_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_type": {
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
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbSystemShapesDataSource-", DatabaseDbSystemShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemShape := map[string]interface{}{}

		if r.AvailableCoreCount != nil {
			dbSystemShape["available_core_count"] = *r.AvailableCoreCount
		}

		if r.AvailableCoreCountPerNode != nil {
			dbSystemShape["available_core_count_per_node"] = *r.AvailableCoreCountPerNode
		}

		if r.AvailableDataStorageInTBs != nil {
			dbSystemShape["available_data_storage_in_tbs"] = *r.AvailableDataStorageInTBs
		}

		if r.AvailableDataStoragePerServerInTBs != nil {
			dbSystemShape["available_data_storage_per_server_in_tbs"] = *r.AvailableDataStoragePerServerInTBs
		}

		if r.AvailableDbNodePerNodeInGBs != nil {
			dbSystemShape["available_db_node_per_node_in_gbs"] = *r.AvailableDbNodePerNodeInGBs
		}

		if r.AvailableDbNodeStorageInGBs != nil {
			dbSystemShape["available_db_node_storage_in_gbs"] = *r.AvailableDbNodeStorageInGBs
		}

		if r.AvailableMemoryInGBs != nil {
			dbSystemShape["available_memory_in_gbs"] = *r.AvailableMemoryInGBs
		}

		if r.AvailableMemoryPerNodeInGBs != nil {
			dbSystemShape["available_memory_per_node_in_gbs"] = *r.AvailableMemoryPerNodeInGBs
		}

		if r.CoreCountIncrement != nil {
			dbSystemShape["core_count_increment"] = *r.CoreCountIncrement
		}

		if r.MaxStorageCount != nil {
			dbSystemShape["max_storage_count"] = *r.MaxStorageCount
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

		if r.MinStorageCount != nil {
			dbSystemShape["min_storage_count"] = *r.MinStorageCount
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

		dbSystemShape["shape_type"] = r.ShapeType

		resources = append(resources, dbSystemShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbSystemShapesDataSource().Schema["db_system_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_system_shapes", resources); err != nil {
		return err
	}

	return nil
}
