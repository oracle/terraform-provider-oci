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

func DatabaseAutonomousDbVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDbVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_db_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_workload": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_dedicated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_default_for_free": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_default_for_paid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_free_tier_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_paid_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDbVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDbVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDbVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDbVersionsResponse
}

func (s *DatabaseAutonomousDbVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDbVersionsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDbVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
		request.DbWorkload = oci_database.AutonomousDatabaseSummaryDbWorkloadEnum(dbWorkload.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDbVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDbVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDbVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDbVersionsDataSource-", DatabaseAutonomousDbVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDbVersion := map[string]interface{}{}

		autonomousDbVersion["db_workload"] = r.DbWorkload

		if r.Details != nil {
			autonomousDbVersion["details"] = *r.Details
		}

		if r.IsDedicated != nil {
			autonomousDbVersion["is_dedicated"] = *r.IsDedicated
		}

		if r.IsDefaultForFree != nil {
			autonomousDbVersion["is_default_for_free"] = *r.IsDefaultForFree
		}

		if r.IsDefaultForPaid != nil {
			autonomousDbVersion["is_default_for_paid"] = *r.IsDefaultForPaid
		}

		if r.IsFreeTierEnabled != nil {
			autonomousDbVersion["is_free_tier_enabled"] = *r.IsFreeTierEnabled
		}

		if r.IsPaidEnabled != nil {
			autonomousDbVersion["is_paid_enabled"] = *r.IsPaidEnabled
		}

		if r.Version != nil {
			autonomousDbVersion["version"] = *r.Version
		}

		resources = append(resources, autonomousDbVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDbVersionsDataSource().Schema["autonomous_db_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_db_versions", resources); err != nil {
		return err
	}

	return nil
}
