// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     BackupResource(),
			},
		},
	}
}

func readBackups(d *schema.ResourceData, m interface{}) error {
	sync := &BackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type BackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListBackupsResponse
}

func (s *BackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BackupsDataSourceCrud) Get() error {
	request := oci_database.ListBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BackupsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		backup := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			backup["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CompartmentId != nil {
			backup["compartment_id"] = *r.CompartmentId
		}

		if r.DatabaseEdition != nil {
			backup["database_edition"] = *r.DatabaseEdition
		}

		if r.DatabaseId != nil {
			backup["database_id"] = *r.DatabaseId
		}

		if r.DbDataSizeInMBs != nil {
			backup["db_data_size_in_mbs"] = *r.DbDataSizeInMBs
		}

		if r.DisplayName != nil {
			backup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			backup["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			backup["lifecycle_details"] = *r.LifecycleDetails
		}

		backup["state"] = r.LifecycleState

		if r.TimeEnded != nil {
			backup["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeStarted != nil {
			backup["time_started"] = r.TimeStarted.String()
		}

		backup["type"] = r.Type

		resources = append(resources, backup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, BackupsDataSource().Schema["backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backups", resources); err != nil {
		panic(err)
	}

	return
}
