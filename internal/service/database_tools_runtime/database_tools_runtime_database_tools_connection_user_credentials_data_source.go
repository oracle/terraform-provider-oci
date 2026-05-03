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

func DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"database_tools_connection_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_credential_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"enabled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"owner": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"related_resource": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"identifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"windows_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.ListUserCredentialsResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.ListUserCredentialsRequest{}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	if userKey, ok := s.D.GetOkExists("user_key"); ok {
		tmp := userKey.(string)
		request.UserKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.ListUserCredentials(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUserCredentials(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSource-", DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSource(), s.D))

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		s.D.Set("database_tools_connection_id", databaseToolsConnectionId.(string))
	}

	if userKey, ok := s.D.GetOkExists("user_key"); ok {
		s.D.Set("user_key", userKey.(string))
	}

	resources := []map[string]interface{}{}
	databaseToolsConnectionUserCredential := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UserCredentialSummaryToMap(item))
	}
	databaseToolsConnectionUserCredential["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSource().Schema["user_credential_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsConnectionUserCredential["items"] = items
	}

	resources = append(resources, databaseToolsConnectionUserCredential)
	if err := s.D.Set("user_credential_collection", resources); err != nil {
		return err
	}

	return nil
}

func UserCredentialRelatedResourceToMap(obj *oci_database_tools_runtime.CredentialRelatedResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	result["type"] = string(obj.Type)

	return result
}

func UserCredentialSummaryToMap(obj oci_database_tools_runtime.UserCredentialSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Enabled != nil {
		result["enabled"] = string(*obj.Enabled)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["key_type"] = string(obj.KeyType)

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.RelatedResource != nil {
		result["related_resource"] = []interface{}{UserCredentialRelatedResourceToMap(obj.RelatedResource)}
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	if obj.WindowsDomain != nil {
		result["windows_domain"] = string(*obj.WindowsDomain)
	}

	return result
}
