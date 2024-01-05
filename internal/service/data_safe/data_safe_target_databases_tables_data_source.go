// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetDatabasesTablesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeTargetDatabasesTables,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"schema_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"schema_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"table_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"table_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"schema_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"table_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeTargetDatabasesTables(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasesTablesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabasesTablesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListTablesResponse
}

func (s *DataSafeTargetDatabasesTablesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabasesTablesDataSourceCrud) Get() error {
	request := oci_data_safe.ListTablesRequest{}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		interfaces := schemaName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schema_name") {
			request.SchemaName = tmp
		}
	}

	if schemaNameContains, ok := s.D.GetOkExists("schema_name_contains"); ok {
		tmp := schemaNameContains.(string)
		request.SchemaNameContains = &tmp
	}

	if tableName, ok := s.D.GetOkExists("table_name"); ok {
		interfaces := tableName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("table_name") {
			request.TableName = tmp
		}
	}

	if tableNameContains, ok := s.D.GetOkExists("table_name_contains"); ok {
		tmp := tableNameContains.(string)
		request.TableNameContains = &tmp
	}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListTables(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTables(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeTargetDatabasesTablesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabasesTablesDataSource-", DataSafeTargetDatabasesTablesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		targetDatabasesTable := map[string]interface{}{}

		if r.SchemaName != nil {
			targetDatabasesTable["schema_name"] = *r.SchemaName
		}

		if r.TableName != nil {
			targetDatabasesTable["table_name"] = *r.TableName
		}

		resources = append(resources, targetDatabasesTable)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeTargetDatabasesTablesDataSource().Schema["tables"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tables", resources); err != nil {
		return err
	}

	return nil
}
