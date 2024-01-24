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

func DatabaseManagementExternalDbHomeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementExternalDbHome,
		Schema: map[string]*schema.Schema{
			"external_db_home_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"additional_details": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"component_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"home_directory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementExternalDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbHomeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalDbHomeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalDbHomeResponse
}

func (s *DatabaseManagementExternalDbHomeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalDbHomeDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalDbHomeRequest{}

	if externalDbHomeId, ok := s.D.GetOkExists("external_db_home_id"); ok {
		tmp := externalDbHomeId.(string)
		request.ExternalDbHomeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalDbHomeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	if s.Res.HomeDirectory != nil {
		s.D.Set("home_directory", *s.Res.HomeDirectory)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
