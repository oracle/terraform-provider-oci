// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"fmt"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DbHomeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbHomeDataSource,
		Schema: map[string]*schema.Schema{
			// Required
			// Legacy property, just a vanity for what is regularly "id" everywhere else
			// todo: reconcile this one off--preferably deprecate db_home_id in favor of "id" during the great deprecation
			"db_home_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// todo: codegen omits this property
			"last_patch_history_entry_id": {
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
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readDbHomeDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &DbHomeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return crud.ReadResource(sync)
}

type DbHomeDataSourceCrud struct {
	crud.BaseCrud
	Client *oci_database.DatabaseClient
	Res    *oci_database.DbHome
}

func (s *DbHomeDataSourceCrud) Get() error {
	request := oci_database.GetDbHomeRequest{}

	// todo: when deprecating "db_home_id" this should be wired to "id"/s.D.Id()
	tmp := s.D.Get("db_home_id").(string)
	request.DbHomeId = &tmp

	if len(tmp) == 0 {
		return fmt.Errorf("db_home_id must contain a valid ocid")
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbHome
	return nil
}

func (s *DbHomeDataSourceCrud) SetData() {
	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.LastPatchHistoryEntryId != nil {
		s.D.Set("last_patch_history_entry_id", *s.Res.LastPatchHistoryEntryId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())

}
