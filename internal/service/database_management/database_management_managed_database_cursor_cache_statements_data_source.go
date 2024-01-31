// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseCursorCacheStatementsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseCursorCacheStatements,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cursor_cache_statement_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"schema": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_text": {
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

func readDatabaseManagementManagedDatabaseCursorCacheStatements(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseCursorCacheStatementsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseCursorCacheStatementsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCursorCacheStatementsResponse
}

func (s *DatabaseManagementManagedDatabaseCursorCacheStatementsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseCursorCacheStatementsDataSourceCrud) Get() error {
	request := oci_database_management.ListCursorCacheStatementsRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	if sqlText, ok := s.D.GetOkExists("sql_text"); ok {
		tmp := sqlText.(string)
		request.SqlText = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListCursorCacheStatements(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCursorCacheStatements(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseCursorCacheStatementsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseCursorCacheStatementsDataSource-", DatabaseManagementManagedDatabaseCursorCacheStatementsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseCursorCacheStatement := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CursorCacheStatementSummaryToMap(item))
	}
	managedDatabaseCursorCacheStatement["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseCursorCacheStatementsDataSource().Schema["cursor_cache_statement_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseCursorCacheStatement["items"] = items
	}

	resources = append(resources, managedDatabaseCursorCacheStatement)
	if err := s.D.Set("cursor_cache_statement_collection", resources); err != nil {
		return err
	}

	return nil
}

func CursorCacheStatementSummaryToMap(obj oci_database_management.CursorCacheStatementSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Schema != nil {
		result["schema"] = string(*obj.Schema)
	}

	if obj.SqlId != nil {
		result["sql_id"] = string(*obj.SqlId)
	}

	if obj.SqlText != nil {
		result["sql_text"] = string(*obj.SqlText)
	}

	return result
}
