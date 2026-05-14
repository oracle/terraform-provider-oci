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

func DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"credential_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_tools_connection_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"credential_execute_grantee_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResource(),
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.ListCredentialExecuteGranteesResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.ListCredentialExecuteGranteesRequest{}

	if credentialKey, ok := s.D.GetOkExists("credential_key"); ok {
		tmp := credentialKey.(string)
		request.CredentialKey = &tmp
	}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.ListCredentialExecuteGrantees(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCredentialExecuteGrantees(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSource-", DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsConnectionCredentialExecuteGrantee := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CredentialExecuteGranteeSummaryToMap(item))
	}
	databaseToolsConnectionCredentialExecuteGrantee["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSource().Schema["credential_execute_grantee_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsConnectionCredentialExecuteGrantee["items"] = items
	}

	resources = append(resources, databaseToolsConnectionCredentialExecuteGrantee)
	if err := s.D.Set("credential_execute_grantee_collection", resources); err != nil {
		return err
	}

	return nil
}
