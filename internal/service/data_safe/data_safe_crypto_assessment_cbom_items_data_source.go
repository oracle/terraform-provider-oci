// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentCbomItemsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentCbomItemsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"crypto_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"crypto_assessment_cbom_item_collection": {
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
									"algorithm": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compliance_driver": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"component_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"configuration_location": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"feature": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"format": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_size": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": {
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

func readDataSafeCryptoAssessmentCbomItemsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentCbomItemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentCbomItemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentCbomItemsResponse
}

func (s *DataSafeCryptoAssessmentCbomItemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentCbomItemsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentCbomItemsRequest{}

	if cryptoAssessmentId, ok := s.D.GetOkExists("crypto_assessment_id"); ok {
		tmp := cryptoAssessmentId.(string)
		request.CryptoAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListCryptoAssessmentCbomItems(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeCryptoAssessmentCbomItemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentCbomItemsDataSource-", DataSafeCryptoAssessmentCbomItemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessmentCbomItem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentCbomItemSummaryToMap(item))
	}
	cryptoAssessmentCbomItem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentCbomItemsDataSource().Schema["crypto_assessment_cbom_item_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessmentCbomItem["items"] = items
	}

	resources = append(resources, cryptoAssessmentCbomItem)
	if err := s.D.Set("crypto_assessment_cbom_item_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAssessmentCbomItemSummaryToMap(obj oci_data_safe.CryptoAssessmentCbomItemSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Algorithm != nil {
		result["algorithm"] = string(*obj.Algorithm)
	}

	if obj.ComplianceDriver != nil {
		result["compliance_driver"] = string(*obj.ComplianceDriver)
	}

	if obj.ComponentType != nil {
		result["component_type"] = string(*obj.ComponentType)
	}

	result["configuration_location"] = obj.ConfigurationLocation

	if obj.Feature != nil {
		result["feature"] = string(*obj.Feature)
	}

	if obj.Format != nil {
		result["format"] = string(*obj.Format)
	}

	if obj.KeySize != nil {
		result["key_size"] = string(*obj.KeySize)
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	return result
}
