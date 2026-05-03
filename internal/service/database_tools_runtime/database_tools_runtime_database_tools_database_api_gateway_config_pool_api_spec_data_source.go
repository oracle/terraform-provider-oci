// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools_runtime

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["api_spec_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["database_tools_database_api_gateway_config_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["pool_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource(), fieldMap, readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext)
}

func readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest{}

	if apiSpecKey, ok := s.D.GetOkExists("api_spec_key"); ok {
		tmp := apiSpecKey.(string)
		request.ApiSpecKey = &tmp
	}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSource(), s.D))
	switch v := (s.Res.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDefault:
		s.D.Set("type", "DEFAULT")

		if v.Content != nil {
			s.D.Set("content", *v.Content)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if key := v.GetKey(); key != nil {
			s.D.Set("key", *key)
		}

		if timeCreated := v.GetTimeCreated(); timeCreated != nil {
			s.D.Set("time_created", timeCreated.String())
		}

		if timeUpdated := v.GetTimeUpdated(); timeUpdated != nil {
			s.D.Set("time_updated", timeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec)
		return nil
	}

	return nil
}
