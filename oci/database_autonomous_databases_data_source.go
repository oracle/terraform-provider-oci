// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAutonomousDatabases,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
				Elem:     GetDataSourceItemSchema(AutonomousDatabaseResource()),
			},
		},
	}
}

func readAutonomousDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabasesResponse
}

func (s *AutonomousDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDatabasesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

func (s *AutonomousDatabasesDataSourceCrud) SetData() error {
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

		resources = append(resources, autonomousDatabase)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AutonomousDatabasesDataSource().Schema["autonomous_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_databases", resources); err != nil {
		return err
	}

	return nil
}
