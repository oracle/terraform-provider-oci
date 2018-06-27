// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbSystemShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbSystemShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
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
						"core_count_increment": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_node_count": {
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
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDbSystemShapes(d *schema.ResourceData, m interface{}) error {
	sync := &DbSystemShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbSystemShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemShapesResponse
}

func (s *DbSystemShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbSystemShapesDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
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

func (s *DbSystemShapesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystemShape := map[string]interface{}{}

		if r.AvailableCoreCount != nil {
			dbSystemShape["available_core_count"] = *r.AvailableCoreCount
		}

		if r.CoreCountIncrement != nil {
			dbSystemShape["core_count_increment"] = *r.CoreCountIncrement
		}

		if r.MaximumNodeCount != nil {
			dbSystemShape["maximum_node_count"] = *r.MaximumNodeCount
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

		resources = append(resources, dbSystemShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DbSystemShapesDataSource().Schema["db_system_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_system_shapes", resources); err != nil {
		panic(err)
	}

	return
}
