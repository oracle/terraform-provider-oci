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

func DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["credential_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["database_tools_connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["execute_grantee_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResource(), fieldMap, readSingularDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeWithContext)
}

func readSingularDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetCredentialExecuteGranteeResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetCredentialExecuteGranteeRequest{}

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if executeGranteeKey, ok := s.D.GetOkExists("execute_grantee_key"); ok {
		tmp := executeGranteeKey.(string)
		request.ExecuteGranteeKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetCredentialExecuteGrantee(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSource-", DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSource(), s.D))

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	return nil
}
