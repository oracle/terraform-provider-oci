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

func DataSafeCompatibleFormatsForSensitiveTypeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeCompatibleFormatsForSensitiveType,
		Schema: map[string]*schema.Schema{
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Computed
			"formats_for_sensitive_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
						"sensitive_type_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDataSafeCompatibleFormatsForSensitiveType(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeCompatibleFormatsForSensitiveTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeCompatibleFormatsForSensitiveTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetCompatibleFormatsForSensitiveTypesResponse
}

func (s *DataSafeCompatibleFormatsForSensitiveTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCompatibleFormatsForSensitiveTypeDataSourceCrud) Get() error {
	request := oci_data_safe.GetCompatibleFormatsForSensitiveTypesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.GetCompatibleFormatsForSensitiveTypesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetCompatibleFormatsForSensitiveTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeCompatibleFormatsForSensitiveTypeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCompatibleFormatsForSensitiveTypeDataSource-", DataSafeCompatibleFormatsForSensitiveTypeDataSource(), s.D))

	formatsForSensitiveType := []interface{}{}
	for _, item := range s.Res.FormatsForSensitiveType {
		formatsForSensitiveType = append(formatsForSensitiveType, FormatsForSensitiveTypeToMap(item))
	}
	s.D.Set("formats_for_sensitive_type", formatsForSensitiveType)

	return nil
}

func DataSafeFormatSummaryToMap(obj oci_data_safe.FormatSummary) map[string]interface{} {
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

func FormatsForSensitiveTypeToMap(obj oci_data_safe.FormatsForSensitiveType) map[string]interface{} {
	result := map[string]interface{}{}

	maskingFormats := []interface{}{}
	for _, item := range obj.MaskingFormats {
		maskingFormats = append(maskingFormats, FormatSummaryToMap(item))
	}
	result["masking_formats"] = maskingFormats

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	return result
}
