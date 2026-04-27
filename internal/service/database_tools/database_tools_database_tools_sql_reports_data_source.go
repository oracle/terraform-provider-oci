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

func DatabaseToolsDatabaseToolsSqlReportsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseToolsDatabaseToolsSqlReportsWithContext,
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
			"database_tools_sql_report_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseToolsDatabaseToolsSqlReportResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsDatabaseToolsSqlReportsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsSqlReportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsDatabaseToolsSqlReportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.ListDatabaseToolsSqlReportsResponse
}

func (s *DatabaseToolsDatabaseToolsSqlReportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsSqlReportsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.ListDatabaseToolsSqlReportsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_tools.ListDatabaseToolsSqlReportsLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]oci_database_tools.DatabaseToolsSqlReportTypeEnum, 0, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				val := interfaces[i].(string)
				if e, ok := oci_database_tools.GetMappingDatabaseToolsSqlReportTypeEnum(val); ok {
					tmp = append(tmp, e)
				}
			}
		}
		if len(tmp) != 0 {
			request.Type = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.ListDatabaseToolsSqlReports(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseToolsSqlReports(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsDatabaseToolsSqlReportsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsDatabaseToolsSqlReportsDataSource-", DatabaseToolsDatabaseToolsSqlReportsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsSqlReport := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseToolsSqlReportSummaryToMap(item))
	}
	databaseToolsSqlReport["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsDatabaseToolsSqlReportsDataSource().Schema["database_tools_sql_report_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsSqlReport["items"] = items
	}

	resources = append(resources, databaseToolsSqlReport)
	if err := s.D.Set("database_tools_sql_report_collection", resources); err != nil {
		return err
	}

	return nil
}
