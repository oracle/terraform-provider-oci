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

func DataSafeSdmMaskingPolicyDifferencesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSdmMaskingPolicyDifferences,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"difference_access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sdm_masking_policy_difference_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeSdmMaskingPolicyDifferenceResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeSdmMaskingPolicyDifferences(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferencesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSdmMaskingPolicyDifferencesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSdmMaskingPolicyDifferencesResponse
}

func (s *DataSafeSdmMaskingPolicyDifferencesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSdmMaskingPolicyDifferencesDataSourceCrud) Get() error {
	request := oci_data_safe.ListSdmMaskingPolicyDifferencesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if differenceAccessLevel, ok := s.D.GetOkExists("difference_access_level"); ok {
		request.DifferenceAccessLevel = oci_data_safe.ListSdmMaskingPolicyDifferencesDifferenceAccessLevelEnum(differenceAccessLevel.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.SdmMaskingPolicyDifferenceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSdmMaskingPolicyDifferences(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSdmMaskingPolicyDifferences(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSdmMaskingPolicyDifferencesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSdmMaskingPolicyDifferencesDataSource-", DataSafeSdmMaskingPolicyDifferencesDataSource(), s.D))
	resources := []map[string]interface{}{}
	sdmMaskingPolicyDifference := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SdmMaskingPolicyDifferenceSummaryToMap(item))
	}
	sdmMaskingPolicyDifference["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSdmMaskingPolicyDifferencesDataSource().Schema["sdm_masking_policy_difference_collection"].Elem.(*schema.Resource).Schema)
		sdmMaskingPolicyDifference["items"] = items
	}

	resources = append(resources, sdmMaskingPolicyDifference)
	if err := s.D.Set("sdm_masking_policy_difference_collection", resources); err != nil {
		return err
	}

	return nil
}
