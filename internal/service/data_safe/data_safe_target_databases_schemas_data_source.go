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

func DataSafeTargetDatabasesSchemasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeTargetDatabasesSchemas,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"is_oracle_maintained": {
				Type:     schema.TypeBool,
				Optional: true,
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
			"target_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_oracle_maintained": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"schema_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeTargetDatabasesSchemas(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasesSchemasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabasesSchemasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSchemasResponse
}

func (s *DataSafeTargetDatabasesSchemasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabasesSchemasDataSourceCrud) Get() error {
	request := oci_data_safe.ListSchemasRequest{}

	if isOracleMaintained, ok := s.D.GetOkExists("is_oracle_maintained"); ok {
		tmp := isOracleMaintained.(bool)
		request.IsOracleMaintained = &tmp
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

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSchemas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSchemas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeTargetDatabasesSchemasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabasesSchemasDataSource-", DataSafeTargetDatabasesSchemasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		targetDatabasesSchema := map[string]interface{}{}

		if r.IsOracleMaintained != nil {
			targetDatabasesSchema["is_oracle_maintained"] = *r.IsOracleMaintained
		}

		if r.SchemaName != nil {
			targetDatabasesSchema["schema_name"] = *r.SchemaName
		}

		resources = append(resources, targetDatabasesSchema)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeTargetDatabasesSchemasDataSource().Schema["schemas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("schemas", resources); err != nil {
		return err
	}

	return nil
}
