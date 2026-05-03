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

func DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["property_set_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResource(), fieldMap, readSingularDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext)
}

func readSingularDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.GetPropertySetResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.GetPropertySetRequest{}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if propertySetKey, ok := s.D.GetOkExists("property_set_key"); ok {
		request.PropertySetKey = oci_database_tools_runtime.GetPropertySetPropertySetKeyEnum(propertySetKey.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.GetPropertySet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSource-", DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSource(), s.D))
	switch v := (s.Res.PropertySet).(type) {
	case oci_database_tools_runtime.PropertySetApex:
		s.D.Set("key", "APEX")
		clearNonApexPropertySetDataSourceFields(s.D)

		if v.UserKey != nil {
			s.D.Set("user_key", *v.UserKey)
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	case oci_database_tools_runtime.PropertySetApexDocumentGenerator:
		s.D.Set("key", "APEX_DOCUMENT_GENERATOR")
		clearNonApexDocumentGeneratorPropertySetDataSourceFields(s.D)

		s.D.Set("autonomous_database_resource_principal_status", v.AutonomousDatabaseResourcePrincipalStatus)

		if v.CredentialKey != nil {
			s.D.Set("credential_key", *v.CredentialKey)
		}

		if v.FunctionId != nil {
			s.D.Set("function_id", *v.FunctionId)
		}

		if v.InvokeEndpoint != nil {
			s.D.Set("invoke_endpoint", *v.InvokeEndpoint)
		}

		if v.ObjectStorageBucketCompartmentId != nil {
			s.D.Set("object_storage_bucket_compartment_id", *v.ObjectStorageBucketCompartmentId)
		}

		if v.ObjectStorageEndpoint != nil {
			s.D.Set("object_storage_endpoint", *v.ObjectStorageEndpoint)
		}

		if v.ObjectStorageNamespace != nil {
			s.D.Set("object_storage_namespace", *v.ObjectStorageNamespace)
		}

		s.D.Set("print_server_type", v.PrintServerType)

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	case oci_database_tools_runtime.PropertySetApexFaIntegration:
		s.D.Set("key", "APEX_FA_INTEGRATION")
		clearNonApexFaIntegrationPropertySetDataSourceFields(s.D)

		s.D.Set("authentication_substitutions", v.AuthenticationSubstitutions)

		if v.InstanceDbmsCredentialEnabled != nil {
			s.D.Set("instance_dbms_credential_enabled", *v.InstanceDbmsCredentialEnabled)
		}

		if v.PrerequisitesCheck != nil {
			s.D.Set("prerequisites_check", []interface{}{ApexFaIntegrationPrerequisitesCheckToMap(v.PrerequisitesCheck)})
		} else {
			s.D.Set("prerequisites_check", nil)
		}

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	case oci_database_tools_runtime.PropertySetOracleDatabaseExternalAuthentication:
		s.D.Set("key", "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION")
		clearNonOracleDatabaseExternalAuthenticationPropertySetDataSourceFields(s.D)

		if v.IdentityProvider != nil {
			identityProviderArray := []interface{}{}
			if identityProviderMap := PropertySetOracleDatabaseExternalAuthenticationIdentityProviderToMap(&v.IdentityProvider); identityProviderMap != nil {
				identityProviderArray = append(identityProviderArray, identityProviderMap)
			}
			s.D.Set("identity_provider", identityProviderArray)
		} else {
			s.D.Set("identity_provider", nil)
		}

		if v.IsMutable != nil {
			s.D.Set("is_mutable", *v.IsMutable)
		}
	default:
		log.Printf("[WARN] Received 'key' of unknown type %v", s.Res.PropertySet)
		return nil
	}

	return nil
}

func clearNonApexPropertySetDataSourceFields(d *schema.ResourceData) {
	d.Set("authentication_substitutions", nil)
	d.Set("autonomous_database_resource_principal_status", nil)
	d.Set("credential_key", nil)
	d.Set("function_id", nil)
	d.Set("identity_provider", nil)
	d.Set("instance_dbms_credential_enabled", nil)
	d.Set("invoke_endpoint", nil)
	d.Set("object_storage_bucket_compartment_id", nil)
	d.Set("object_storage_endpoint", nil)
	d.Set("object_storage_namespace", nil)
	d.Set("prerequisites_check", nil)
	d.Set("print_server_type", nil)
}

func clearNonApexDocumentGeneratorPropertySetDataSourceFields(d *schema.ResourceData) {
	d.Set("authentication_substitutions", nil)
	d.Set("identity_provider", nil)
	d.Set("instance_dbms_credential_enabled", nil)
	d.Set("prerequisites_check", nil)
	d.Set("user_key", nil)
	d.Set("version", nil)
}

func clearNonApexFaIntegrationPropertySetDataSourceFields(d *schema.ResourceData) {
	d.Set("autonomous_database_resource_principal_status", nil)
	d.Set("credential_key", nil)
	d.Set("function_id", nil)
	d.Set("identity_provider", nil)
	d.Set("invoke_endpoint", nil)
	d.Set("object_storage_bucket_compartment_id", nil)
	d.Set("object_storage_endpoint", nil)
	d.Set("object_storage_namespace", nil)
	d.Set("print_server_type", nil)
	d.Set("user_key", nil)
	d.Set("version", nil)
}

func clearNonOracleDatabaseExternalAuthenticationPropertySetDataSourceFields(d *schema.ResourceData) {
	d.Set("authentication_substitutions", nil)
	d.Set("autonomous_database_resource_principal_status", nil)
	d.Set("credential_key", nil)
	d.Set("function_id", nil)
	d.Set("instance_dbms_credential_enabled", nil)
	d.Set("invoke_endpoint", nil)
	d.Set("object_storage_bucket_compartment_id", nil)
	d.Set("object_storage_endpoint", nil)
	d.Set("object_storage_namespace", nil)
	d.Set("prerequisites_check", nil)
	d.Set("print_server_type", nil)
	d.Set("user_key", nil)
	d.Set("version", nil)
}
