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

func databaseToolsDatabaseApiGatewayConfigPoolApiSpecListItemResource() *schema.Resource {
	itemSchema := DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource().Schema
	itemSchema["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return &schema.Resource{
		Schema: itemSchema,
	}
}

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"database_tools_database_api_gateway_config_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_tools_database_api_gateway_config_pool_api_spec_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     databaseToolsDatabaseApiGatewayConfigPoolApiSpecListItemResource(),
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest{}

	if databaseToolsDatabaseApiGatewayConfigId, ok := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); ok {
		tmp := databaseToolsDatabaseApiGatewayConfigId.(string)
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if poolKey, ok := s.D.GetOkExists("pool_key"); ok {
		tmp := poolKey.(string)
		request.PoolKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools_runtime")

	response, err := s.Client.ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsDatabaseApiGatewayConfigPoolApiSpec := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		itemMap := DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSummaryToMap(item)
		if itemMap != nil {
			if key, ok := itemMap["key"].(string); ok {
				if databaseToolsDatabaseApiGatewayConfigId, configOk := s.D.GetOkExists("database_tools_database_api_gateway_config_id"); configOk {
					if poolKey, poolOk := s.D.GetOkExists("pool_key"); poolOk {
						itemMap["id"] = GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCompositeId(key, databaseToolsDatabaseApiGatewayConfigId.(string), poolKey.(string))
					}
				}
			}
			items = append(items, itemMap)
		}
	}
	databaseToolsDatabaseApiGatewayConfigPoolApiSpec["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(
			f.(*schema.Set),
			items,
			DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSource().
				Schema["database_tools_database_api_gateway_config_pool_api_spec_collection"].
				Elem.(*schema.Resource).
				Schema["items"].
				Elem.(*schema.Resource).
				Schema,
		)
		databaseToolsDatabaseApiGatewayConfigPoolApiSpec["items"] = items
	}

	resources = append(resources, databaseToolsDatabaseApiGatewayConfigPoolApiSpec)
	if err := s.D.Set("database_tools_database_api_gateway_config_pool_api_spec_collection", resources); err != nil {
		return err
	}

	return nil
}
