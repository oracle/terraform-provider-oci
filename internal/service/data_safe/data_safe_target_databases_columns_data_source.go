// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetDatabasesColumnsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeTargetDatabasesColumns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"column_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"column_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"datatype": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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
			"columns": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"character_length": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"column_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"length": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"precision": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scale": {
							Type:     schema.TypeInt,
							Computed: true,
						},
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

func readDataSafeTargetDatabasesColumns(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasesColumnsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabasesColumnsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListColumnsResponse
}

func (s *DataSafeTargetDatabasesColumnsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabasesColumnsDataSourceCrud) Get() error {
	request := oci_data_safe.ListColumnsRequest{}

	if columnName, ok := s.D.GetOkExists("column_name"); ok {
		interfaces := columnName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("column_name") {
			request.ColumnName = tmp
		}
	}

	if columnNameContains, ok := s.D.GetOkExists("column_name_contains"); ok {
		tmp := columnNameContains.(string)
		request.ColumnNameContains = &tmp
	}

	if datatype, ok := s.D.GetOkExists("datatype"); ok {
		interfaces := datatype.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("datatype") {
			request.Datatype = tmp
		}
	}

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

	response, err := s.Client.ListColumns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListColumns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeTargetDatabasesColumnsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabasesColumnsDataSource-", DataSafeTargetDatabasesColumnsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		targetDatabasesColumn := map[string]interface{}{}

		if r.CharacterLength != nil {
			targetDatabasesColumn["character_length"] = *r.CharacterLength
		}

		if r.ColumnName != nil {
			targetDatabasesColumn["column_name"] = *r.ColumnName
		}

		if r.DataType != nil {
			targetDatabasesColumn["data_type"] = *r.DataType
		}

		if r.Length != nil {
			targetDatabasesColumn["length"] = strconv.FormatInt(*r.Length, 10)
		}

		if r.Precision != nil {
			targetDatabasesColumn["precision"] = *r.Precision
		}

		if r.Scale != nil {
			targetDatabasesColumn["scale"] = *r.Scale
		}

		if r.SchemaName != nil {
			targetDatabasesColumn["schema_name"] = *r.SchemaName
		}

		if r.TableName != nil {
			targetDatabasesColumn["table_name"] = *r.TableName
		}

		resources = append(resources, targetDatabasesColumn)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeTargetDatabasesColumnsDataSource().Schema["columns"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("columns", resources); err != nil {
		return err
	}

	return nil
}
