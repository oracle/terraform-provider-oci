// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseInfrastructureTargetVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseInfrastructureTargetVersion,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"target_db_version_history_entry": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_storage_version_history_entry": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularDatabaseInfrastructureTargetVersion(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseInfrastructureTargetVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseInfrastructureTargetVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetInfrastructureTargetVersionsResponse
}

func (s *DatabaseInfrastructureTargetVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseInfrastructureTargetVersionDataSourceCrud) Get() error {
	request := oci_database.GetInfrastructureTargetVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if targetResourceId, ok := s.D.GetOkExists("target_resource_id"); ok {
		tmp := targetResourceId.(string)
		request.TargetResourceId = &tmp
	}

	if targetResourceType, ok := s.D.GetOkExists("target_resource_type"); ok {
		request.TargetResourceType = oci_database.MaintenanceRunSummaryTargetResourceTypeEnum(targetResourceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetInfrastructureTargetVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseInfrastructureTargetVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseInfrastructureTargetVersionDataSource-", DatabaseInfrastructureTargetVersionDataSource(), s.D))

	s.D.Set("target_db_version_history_entry", s.Res.TargetDbVersionHistoryEntry)

	if s.Res.TargetResourceId != nil {
		s.D.Set("target_resource_id", *s.Res.TargetResourceId)
	}

	s.D.Set("target_resource_type", s.Res.TargetResourceType)

	s.D.Set("target_storage_version_history_entry", s.Res.TargetStorageVersionHistoryEntry)

	return nil
}
