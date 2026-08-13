// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingAnalyticsWithContext,
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
			"group_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_database_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_analytics_collection": {
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
									"subsetting_analytic_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_subsetted": {
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

func readDataSafeSubsettingAnalyticsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingAnalyticsResponse
}

func (s *DataSafeSubsettingAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingAnalyticsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingAnalyticsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_data_safe.ListSubsettingAnalyticsGroupByEnum(groupBy.(string))
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if targetDatabaseGroupId, ok := s.D.GetOkExists("target_database_group_id"); ok {
		tmp := targetDatabaseGroupId.(string)
		request.TargetDatabaseGroupId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingAnalytics(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingAnalytics(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingAnalyticsDataSource-", DataSafeSubsettingAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingAnalyticsSummaryToMap(item))
	}
	subsettingAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingAnalyticsDataSource().Schema["subsetting_analytics_collection"].Elem.(*schema.Resource).Schema)
		subsettingAnalytic["items"] = items
	}

	resources = append(resources, subsettingAnalytic)
	if err := s.D.Set("subsetting_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettingAnalyticsDimensionsToMap(obj *oci_data_safe.SubsettingAnalyticsDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PolicyId != nil {
		result["policy_id"] = string(*obj.PolicyId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}

func SubsettingAnalyticsSummaryToMap(obj oci_data_safe.SubsettingAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SubsettingAnalyticsDimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	if obj.Count != nil {
		result["subsetting_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.TimeLastSubsetted != nil {
		result["time_last_subsetted"] = obj.TimeLastSubsetted.String()
	}

	return result
}
