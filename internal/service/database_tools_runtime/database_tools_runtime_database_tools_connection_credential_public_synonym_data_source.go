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

func DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["credential_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["database_tools_connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["public_synonym_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResource(), fieldMap, readSingularDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext)
}

func readSingularDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetCredentialPublicSynonymResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetCredentialPublicSynonymRequest{}

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if publicSynonymKey, ok := s.D.GetOkExists("public_synonym_key"); ok {
		tmp := publicSynonymKey.(string)
		request.PublicSynonymKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetCredentialPublicSynonym(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSource-", DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSource(), s.D))

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		s.D.Set("credential_key", credentialKey.(string))
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		s.D.Set("database_tools_connection_id", databaseToolsConnectionId.(string))
	}

	if publicSynonymKey, ok := s.D.GetOkExists("public_synonym_key"); ok {
		s.D.Set("public_synonym_key", publicSynonymKey.(string))
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	return nil
}
