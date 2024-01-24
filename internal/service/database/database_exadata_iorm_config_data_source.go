// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseExadataIormConfigDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["db_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseExadataIormConfigResource(), fieldMap, readSingularDatabaseExadataIormConfig)
}

func readSingularDatabaseExadataIormConfig(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataIormConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExadataIormConfigDataSource-", DatabaseExadataIormConfigDataSource(), s.D))

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
