// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"database_tools_database_api_gateway_config_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.ListDatabaseToolsDatabaseApiGatewayConfigsResponse
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.ListDatabaseToolsDatabaseApiGatewayConfigsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_tools.ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]oci_database_tools.DatabaseApiGatewayConfigTypeEnum, 0, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				val := interfaces[i].(string)
				if e, ok := oci_database_tools.GetMappingDatabaseApiGatewayConfigTypeEnum(val); ok {
					tmp = append(tmp, e)
				}
			}
		}
		if len(tmp) != 0 {
			request.Type = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.ListDatabaseToolsDatabaseApiGatewayConfigs(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseToolsDatabaseApiGatewayConfigs(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSource-", DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsDatabaseApiGatewayConfig := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseToolsDatabaseApiGatewayConfigSummaryToMap(item))
	}
	databaseToolsDatabaseApiGatewayConfig["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSource().Schema["database_tools_database_api_gateway_config_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsDatabaseApiGatewayConfig["items"] = items
	}

	resources = append(resources, databaseToolsDatabaseApiGatewayConfig)
	if err := s.D.Set("database_tools_database_api_gateway_config_collection", resources); err != nil {
		return err
	}

	return nil
}
