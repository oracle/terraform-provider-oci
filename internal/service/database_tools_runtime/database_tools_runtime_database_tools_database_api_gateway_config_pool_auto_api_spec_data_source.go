// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools_runtime

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["auto_api_spec_key"] = &schema.Schema{
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
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource(), fieldMap, readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext)
}

func readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest{}

	if autoApiSpecKey, ok := s.D.GetOkExists("auto_api_spec_key"); ok {
		tmp := autoApiSpecKey.(string)
		request.AutoApiSpecKey = &tmp
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

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSource(), s.D))
	v := s.Res.DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec
	s.D.Set("type", "DEFAULT")
	if autoApiSpecKey, ok := s.D.GetOkExists("auto_api_spec_key"); ok {
		s.D.Set("auto_api_spec_key", autoApiSpecKey.(string))
	}
	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		s.D.Set("database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId.(string))
	}
	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		s.D.Set("pool_key", poolKey.(string))
	}

	if v.GetAlias() != nil {
		s.D.Set("alias", *v.GetAlias())
	}
	if v.GetDatabaseObjectName() != nil {
		s.D.Set("database_object_name", *v.GetDatabaseObjectName())
	}
	s.D.Set("database_object_type", v.GetDatabaseObjectType())
	if v.GetDescription() != nil {
		s.D.Set("description", *v.GetDescription())
	}
	if v.GetDisplayName() != nil {
		s.D.Set("display_name", *v.GetDisplayName())
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
	s.D.Set("operations", autoApiSpecOperationsToStrings(v.GetOperations()))
	s.D.Set("roles", v.GetRoles())
	if v.GetScope() != nil {
		s.D.Set("scope", *v.GetScope())
	}
	s.D.Set("security_schemes", autoApiSpecSecuritySchemesToStrings(v.GetSecuritySchemes()))

	return nil
}
