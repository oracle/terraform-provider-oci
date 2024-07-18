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

func DataSafeSensitiveDataModelSensitiveTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSensitiveDataModelSensitiveTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_data_model_sensitive_type_collection": {
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
									"sensitive_data_model_sensitive_type_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sensitive_type_id": {
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

func readDataSafeSensitiveDataModelSensitiveTypes(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelSensitiveTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelSensitiveTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSensitiveDataModelSensitiveTypesResponse
}

func (s *DataSafeSensitiveDataModelSensitiveTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelSensitiveTypesDataSourceCrud) Get() error {
	request := oci_data_safe.ListSensitiveDataModelSensitiveTypesRequest{}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSensitiveDataModelSensitiveTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSensitiveDataModelSensitiveTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSensitiveDataModelSensitiveTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveDataModelSensitiveTypesDataSource-", DataSafeSensitiveDataModelSensitiveTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	sensitiveDataModelSensitiveType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SensitiveDataModelSensitiveTypeSummaryToMap(item))
	}
	sensitiveDataModelSensitiveType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSensitiveDataModelSensitiveTypesDataSource().Schema["sensitive_data_model_sensitive_type_collection"].Elem.(*schema.Resource).Schema)
		sensitiveDataModelSensitiveType["items"] = items
	}

	resources = append(resources, sensitiveDataModelSensitiveType)
	if err := s.D.Set("sensitive_data_model_sensitive_type_collection", resources); err != nil {
		return err
	}

	return nil
}

func SensitiveDataModelSensitiveTypeSummaryToMap(obj oci_data_safe.SensitiveDataModelSensitiveTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["sensitive_data_model_sensitive_type_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	return result
}
