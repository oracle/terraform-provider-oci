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

func DataSafeCompatibleFormatsForDataTypeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeCompatibleFormatsForDataType,
		Schema: map[string]*schema.Schema{
			// Computed
			"formats_for_data_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"data_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"masking_formats": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
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

func readSingularDataSafeCompatibleFormatsForDataType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeCompatibleFormatsForDataTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeCompatibleFormatsForDataTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetCompatibleFormatsForDataTypesResponse
}

func (s *DataSafeCompatibleFormatsForDataTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCompatibleFormatsForDataTypeDataSourceCrud) Get() error {
	request := oci_data_safe.GetCompatibleFormatsForDataTypesRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetCompatibleFormatsForDataTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeCompatibleFormatsForDataTypeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCompatibleFormatsForDataTypeDataSource-", DataSafeCompatibleFormatsForDataTypeDataSource(), s.D))

	formatsForDataType := []interface{}{}
	for _, item := range s.Res.FormatsForDataType {
		formatsForDataType = append(formatsForDataType, FormatsForDataTypeToMap(item))
	}
	s.D.Set("formats_for_data_type", formatsForDataType)

	return nil
}

func FormatSummaryToMap(obj oci_data_safe.FormatSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func FormatsForDataTypeToMap(obj oci_data_safe.FormatsForDataType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	maskingFormats := []interface{}{}
	for _, item := range obj.MaskingFormats {
		maskingFormats = append(maskingFormats, FormatSummaryToMap(item))
	}
	result["masking_formats"] = maskingFormats

	return result
}
