// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDataWarehousesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAutonomousDataWarehouses,
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
			"autonomous_data_warehouses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(AutonomousDataWarehouseResource()),
			},
		},
	}
}

func readAutonomousDataWarehouses(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehousesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDataWarehousesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDataWarehousesResponse
}

func (s *AutonomousDataWarehousesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDataWarehousesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDataWarehousesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDataWarehouseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDataWarehouses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDataWarehouses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AutonomousDataWarehousesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDataWarehouse := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ConnectionStrings != nil {
			autonomousDataWarehouse["connection_strings"] = []interface{}{AutonomousDataWarehouseConnectionStringsToMap(r.ConnectionStrings)}
		} else {
			autonomousDataWarehouse["connection_strings"] = nil
		}

		if r.CpuCoreCount != nil {
			autonomousDataWarehouse["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.DataStorageSizeInTBs != nil {
			autonomousDataWarehouse["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbName != nil {
			autonomousDataWarehouse["db_name"] = *r.DbName
		}

		if r.DbVersion != nil {
			autonomousDataWarehouse["db_version"] = *r.DbVersion
		}

		if r.DefinedTags != nil {
			autonomousDataWarehouse["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousDataWarehouse["display_name"] = *r.DisplayName
		}

		autonomousDataWarehouse["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousDataWarehouse["id"] = *r.Id
		}

		autonomousDataWarehouse["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousDataWarehouse["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ServiceConsoleUrl != nil {
			autonomousDataWarehouse["service_console_url"] = *r.ServiceConsoleUrl
		}

		autonomousDataWarehouse["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousDataWarehouse["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, autonomousDataWarehouse)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AutonomousDataWarehousesDataSource().Schema["autonomous_data_warehouses"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_data_warehouses", resources); err != nil {
		return err
	}

	return nil
}
