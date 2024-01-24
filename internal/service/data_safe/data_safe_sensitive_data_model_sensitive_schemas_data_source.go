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

func DataSafeSensitiveDataModelSensitiveSchemasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveDataModelSensitiveSchemas,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"schema_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sensitive_schema_collection": {
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
									"schema_name": {
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

func readDataSafeSensitiveDataModelSensitiveSchemas(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelSensitiveSchemasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelSensitiveSchemasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSensitiveSchemasResponse
}

func (s *DataSafeSensitiveDataModelSensitiveSchemasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelSensitiveSchemasDataSourceCrud) Get() error {
	request := oci_data_safe.ListSensitiveSchemasRequest{}

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

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSensitiveSchemas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSensitiveSchemas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveDataModelSensitiveSchemasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveDataModelSensitiveSchemasDataSource-", DataSafeSensitiveDataModelSensitiveSchemasDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveDataModelSensitiveSchema := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SensitiveSchemaSummaryToMap(item))
	}
	sensitiveDataModelSensitiveSchema["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveDataModelSensitiveSchemasDataSource().Schema["sensitive_schema_collection"].Elem.(*schema.Resource).Schema)
		sensitiveDataModelSensitiveSchema["items"] = items
	}

	resources = append(resources, sensitiveDataModelSensitiveSchema)
	if err := s.D.Set("sensitive_schema_collection", resources); err != nil {
		return err
	}

	return nil
}

func SensitiveSchemaSummaryToMap(obj oci_data_safe.SensitiveSchemaSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	return result
}
