// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseExadataIormConfigDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseExadataIormConfig,
		Schema: map[string]*schema.Schema{
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"db_plans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"flash_cache_limit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"share": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"objective": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseExadataIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataIormConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseExadataIormConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExadataIormConfigResponse
}

func (s *DatabaseExadataIormConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadataIormConfigDataSourceCrud) Get() error {
	request := oci_database.GetExadataIormConfigRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetExadataIormConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExadataIormConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	dbPlans := []interface{}{}
	for _, item := range s.Res.DbPlans {
		if configMap := dbIormConfigToMap(item); configMap != nil {
			dbPlans = append(dbPlans, configMap)
		}
	}
	s.D.Set("db_plans", dbPlans)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("objective", s.Res.Objective)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
