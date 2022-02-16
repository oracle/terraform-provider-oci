// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseAutonomousPatchDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousPatch,
		Schema: map[string]*schema.Schema{
			"autonomous_patch_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"quarter": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_released": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"year": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseAutonomousPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousPatchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousPatchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousPatchResponse
}

func (s *DatabaseAutonomousPatchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousPatchDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousPatchRequest{}

	if autonomousPatchId, ok := s.D.GetOkExists("autonomous_patch_id"); ok {
		tmp := autonomousPatchId.(string)
		request.AutonomousPatchId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousPatchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("patch_model", s.Res.PatchModel)

	if s.Res.Quarter != nil {
		s.D.Set("quarter", *s.Res.Quarter)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	if s.Res.Year != nil {
		s.D.Set("year", *s.Res.Year)
	}

	return nil
}
