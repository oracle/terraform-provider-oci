// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseAutonomousDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabases,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseAutonomousDatabaseResource()),
			},
		},
	}
}

func readDatabaseAutonomousDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseAutonomousDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabasesResponse
}

func (s *DatabaseAutonomousDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
		request.DbWorkload = oci_database.AutonomousDatabaseSummaryDbWorkloadEnum(dbWorkload.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDatabaseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabase := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ConnectionStrings != nil {
			autonomousDatabase["connection_strings"] = []interface{}{AutonomousDatabaseConnectionStringsToMap(r.ConnectionStrings)}
		} else {
			autonomousDatabase["connection_strings"] = nil
		}

		if r.CpuCoreCount != nil {
			autonomousDatabase["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.DataStorageSizeInTBs != nil {
			autonomousDatabase["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbName != nil {
			autonomousDatabase["db_name"] = *r.DbName
		}

		if r.DbVersion != nil {
			autonomousDatabase["db_version"] = *r.DbVersion
		}

		autonomousDatabase["db_workload"] = r.DbWorkload

		if r.DefinedTags != nil {
			autonomousDatabase["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousDatabase["display_name"] = *r.DisplayName
		}

		autonomousDatabase["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousDatabase["id"] = *r.Id
		}

		autonomousDatabase["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousDatabase["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ServiceConsoleUrl != nil {
			autonomousDatabase["service_console_url"] = *r.ServiceConsoleUrl
		}

		autonomousDatabase["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousDatabase["time_created"] = r.TimeCreated.String()
		}

		if r.UsedDataStorageSizeInTBs != nil {
			autonomousDatabase["used_data_storage_size_in_tbs"] = *r.UsedDataStorageSizeInTBs
		}

		resources = append(resources, autonomousDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDatabasesDataSource().Schema["autonomous_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_databases", resources); err != nil {
		return err
	}

	return nil
}
