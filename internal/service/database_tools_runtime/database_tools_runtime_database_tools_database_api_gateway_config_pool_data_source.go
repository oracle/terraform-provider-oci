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

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_database_api_gateway_config_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["pool_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResource(), fieldMap, readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext)
}

func readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolRequest{}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSource(), s.D))
	switch v := (s.Res.DatabaseToolsDatabaseApiGatewayConfigPool).(type) {
	case oci_database_tools_runtime.DatabaseToolsDatabaseApiGatewayConfigPoolDefault:
		s.D.Set("type", "DEFAULT")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		s.D.Set("database_actions_status", v.DatabaseActionsStatus)

		if v.DatabaseToolsConnectionId != nil {
			s.D.Set("database_tools_connection_id", *v.DatabaseToolsConnectionId)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.Key != nil {
			s.D.Set("key", *v.Key)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.InitialPoolSize != nil {
			s.D.Set("initial_pool_size", *v.InitialPoolSize)
		}

		if v.JwtProfileAudience != nil {
			s.D.Set("jwt_profile_audience", *v.JwtProfileAudience)
		}

		if v.JwtProfileIssuer != nil {
			s.D.Set("jwt_profile_issuer", *v.JwtProfileIssuer)
		}

		if v.JwtProfileJwkUrl != nil {
			s.D.Set("jwt_profile_jwk_url", *v.JwtProfileJwkUrl)
		}

		if v.JwtProfileRoleClaimName != nil {
			s.D.Set("jwt_profile_role_claim_name", *v.JwtProfileRoleClaimName)
		}

		if v.MaxPoolSize != nil {
			s.D.Set("max_pool_size", *v.MaxPoolSize)
		}

		if v.MinPoolSize != nil {
			s.D.Set("min_pool_size", *v.MinPoolSize)
		}

		if v.PoolRouteValue != nil {
			s.D.Set("pool_route_value", *v.PoolRouteValue)
		}

		s.D.Set("rest_enabled_sql_status", v.RestEnabledSqlStatus)
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsDatabaseApiGatewayConfigPool)
		return nil
	}

	return nil
}
