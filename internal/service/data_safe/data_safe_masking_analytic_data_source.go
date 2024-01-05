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

func DataSafeMaskingAnalyticDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeMaskingAnalytic,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dimensions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"policy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"metric_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_data_safe_masking_analytic", "oci_data_safe_masking_analytics"),
	}
}

func readSingularDataSafeMaskingAnalytic(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingAnalyticDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingAnalyticDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListMaskingAnalyticsResponse
}

func (s *DataSafeMaskingAnalyticDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingAnalyticDataSourceCrud) Get() error {
	request := oci_data_safe.ListMaskingAnalyticsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_data_safe.ListMaskingAnalyticsGroupByEnum(groupBy.(string))
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListMaskingAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeMaskingAnalyticDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingAnalyticDataSource-", DataSafeMaskingAnalyticDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaskingAnalyticsSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func DataSafeMaskingAnalyticsDimensionsToMap(obj *oci_data_safe.MaskingAnalyticsDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PolicyId != nil {
		result["policy_id"] = string(*obj.PolicyId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}

func DataSafeMaskingAnalyticsSummaryToMap(obj oci_data_safe.MaskingAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{MaskingAnalyticsDimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	return result
}
