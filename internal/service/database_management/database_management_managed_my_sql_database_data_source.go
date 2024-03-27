// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedMySqlDatabaseDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedMySqlDatabase,
		Schema: map[string]*schema.Schema{
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"heat_wave_cluster_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"heat_wave_memory_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"heat_wave_node_shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"heat_wave_nodes": {
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
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_heat_wave_active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_heat_wave_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_lakehouse_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created_heat_wave": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedMySqlDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.GetManagedMySqlDatabaseResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseDataSourceCrud) Get() error {
	request := oci_database_management.GetManagedMySqlDatabaseRequest{}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetManagedMySqlDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.HeatWaveClusterDisplayName != nil {
		s.D.Set("heat_wave_cluster_display_name", *s.Res.HeatWaveClusterDisplayName)
	}

	if s.Res.HeatWaveMemorySize != nil {
		s.D.Set("heat_wave_memory_size", *s.Res.HeatWaveMemorySize)
	}

	if s.Res.HeatWaveNodeShape != nil {
		s.D.Set("heat_wave_node_shape", *s.Res.HeatWaveNodeShape)
	}

	heatWaveNodes := []interface{}{}
	for _, item := range s.Res.HeatWaveNodes {
		heatWaveNodes = append(heatWaveNodes, HeatWaveNodeToMap(item))
	}
	s.D.Set("heat_wave_nodes", heatWaveNodes)

	if s.Res.IsHeatWaveActive != nil {
		s.D.Set("is_heat_wave_active", *s.Res.IsHeatWaveActive)
	}

	if s.Res.IsHeatWaveEnabled != nil {
		s.D.Set("is_heat_wave_enabled", *s.Res.IsHeatWaveEnabled)
	}

	if s.Res.IsLakehouseEnabled != nil {
		s.D.Set("is_lakehouse_enabled", *s.Res.IsLakehouseEnabled)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeCreatedHeatWave != nil {
		s.D.Set("time_created_heat_wave", s.Res.TimeCreatedHeatWave.String())
	}

	return nil
}
