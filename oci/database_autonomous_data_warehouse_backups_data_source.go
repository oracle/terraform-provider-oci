// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDataWarehouseBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAutonomousDataWarehouseBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"autonomous_data_warehouse_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
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
			"autonomous_data_warehouse_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(AutonomousDataWarehouseBackupResource()),
			},
		},
	}
}

func readAutonomousDataWarehouseBackups(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDataWarehouseBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDataWarehouseBackupsResponse
}

func (s *AutonomousDataWarehouseBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDataWarehouseBackupsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDataWarehouseBackupsRequest{}

	if autonomousDataWarehouseId, ok := s.D.GetOkExists("autonomous_data_warehouse_id"); ok {
		tmp := autonomousDataWarehouseId.(string)
		request.AutonomousDataWarehouseId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDataWarehouseBackupSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDataWarehouseBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDataWarehouseBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AutonomousDataWarehouseBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDataWarehouseBackup := map[string]interface{}{}

		if r.AutonomousDataWarehouseId != nil {
			autonomousDataWarehouseBackup["autonomous_data_warehouse_id"] = *r.AutonomousDataWarehouseId
		}

		if r.CompartmentId != nil {
			autonomousDataWarehouseBackup["compartment_id"] = *r.CompartmentId
		}

		if r.DisplayName != nil {
			autonomousDataWarehouseBackup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			autonomousDataWarehouseBackup["id"] = *r.Id
		}

		if r.IsAutomatic != nil {
			autonomousDataWarehouseBackup["is_automatic"] = *r.IsAutomatic
		}

		if r.LifecycleDetails != nil {
			autonomousDataWarehouseBackup["lifecycle_details"] = *r.LifecycleDetails
		}

		autonomousDataWarehouseBackup["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			autonomousDataWarehouseBackup["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			autonomousDataWarehouseBackup["time_started"] = r.TimeStarted.String()
		}

		autonomousDataWarehouseBackup["type"] = r.Type

		resources = append(resources, autonomousDataWarehouseBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AutonomousDataWarehouseBackupsDataSource().Schema["autonomous_data_warehouse_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_data_warehouse_backups", resources); err != nil {
		return err
	}

	return nil
}
