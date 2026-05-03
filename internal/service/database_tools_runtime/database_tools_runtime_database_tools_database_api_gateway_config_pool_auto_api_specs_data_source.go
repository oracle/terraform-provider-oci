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

func DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsWithContext,
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
			"database_tools_database_api_gateway_config_pool_auto_api_spec_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource(),
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsRuntimeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools_runtime.DatabaseToolsRuntimeClient
	Res    *oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest{}

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

	response, err := s.Client.ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSource-", DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSummaryToMap(item))
	}
	databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSource().Schema["database_tools_database_api_gateway_config_pool_auto_api_spec_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec["items"] = items
	}

	resources = append(resources, databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec)
	if err := s.D.Set("database_tools_database_api_gateway_config_pool_auto_api_spec_collection", resources); err != nil {
		return err
	}

	return nil
}
