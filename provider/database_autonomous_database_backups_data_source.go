// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDatabaseBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAutonomousDatabaseBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"autonomous_database_id": {
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
			"autonomous_database_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(AutonomousDatabaseBackupResource()),
			},
		},
	}
}

func readAutonomousDatabaseBackups(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDatabaseBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDatabaseBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseBackupsResponse
}

func (s *AutonomousDatabaseBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDatabaseBackupsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseBackupsRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
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
		request.LifecycleState = oci_database.AutonomousDatabaseBackupSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabaseBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AutonomousDatabaseBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabaseBackup := map[string]interface{}{}

		if r.AutonomousDatabaseId != nil {
			autonomousDatabaseBackup["autonomous_database_id"] = *r.AutonomousDatabaseId
		}

		if r.CompartmentId != nil {
			autonomousDatabaseBackup["compartment_id"] = *r.CompartmentId
		}

		if r.DisplayName != nil {
			autonomousDatabaseBackup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			autonomousDatabaseBackup["id"] = *r.Id
		}

		if r.IsAutomatic != nil {
			autonomousDatabaseBackup["is_automatic"] = *r.IsAutomatic
		}

		if r.LifecycleDetails != nil {
			autonomousDatabaseBackup["lifecycle_details"] = *r.LifecycleDetails
		}

		autonomousDatabaseBackup["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			autonomousDatabaseBackup["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			autonomousDatabaseBackup["time_started"] = r.TimeStarted.String()
		}

		autonomousDatabaseBackup["type"] = r.Type

		resources = append(resources, autonomousDatabaseBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AutonomousDatabaseBackupsDataSource().Schema["autonomous_database_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_database_backups", resources); err != nil {
		return err
	}

	return nil
}
