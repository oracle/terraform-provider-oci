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

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentWithContext,
		Schema: map[string]*schema.Schema{
			"database_tools_database_api_gateway_config_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigContentResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigContentRequest{}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfigContent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSource(), s.D))

	return nil
}
