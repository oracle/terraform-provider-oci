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

func DatabaseAutonomousExadataInfrastructureShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousExadataInfrastructureShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_exadata_infrastructure_shapes": {
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
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousExadataInfrastructureShapes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousExadataInfrastructureShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousExadataInfrastructureShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousExadataInfrastructureShapesResponse
}

func (s *DatabaseAutonomousExadataInfrastructureShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousExadataInfrastructureShapesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousExadataInfrastructureShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousExadataInfrastructureShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousExadataInfrastructureShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousExadataInfrastructureShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousExadataInfrastructureShapesDataSource-", DatabaseAutonomousExadataInfrastructureShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousExadataInfrastructureShape := map[string]interface{}{}

		if r.AvailableCoreCount != nil {
			autonomousExadataInfrastructureShape["available_core_count"] = *r.AvailableCoreCount
		}

		if r.CoreCountIncrement != nil {
			autonomousExadataInfrastructureShape["core_count_increment"] = *r.CoreCountIncrement
		}

		if r.MaximumNodeCount != nil {
			autonomousExadataInfrastructureShape["maximum_node_count"] = *r.MaximumNodeCount
		}

		if r.MinimumCoreCount != nil {
			autonomousExadataInfrastructureShape["minimum_core_count"] = *r.MinimumCoreCount
		}

		if r.MinimumNodeCount != nil {
			autonomousExadataInfrastructureShape["minimum_node_count"] = *r.MinimumNodeCount
		}

		if r.Name != nil {
			autonomousExadataInfrastructureShape["name"] = *r.Name
		}

		resources = append(resources, autonomousExadataInfrastructureShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousExadataInfrastructureShapesDataSource().Schema["autonomous_exadata_infrastructure_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_exadata_infrastructure_shapes", resources); err != nil {
		return err
	}

	return nil
}
