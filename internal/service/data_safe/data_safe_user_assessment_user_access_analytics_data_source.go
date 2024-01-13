// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

func DataSafeUserAssessmentUserAccessAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUserAssessmentUserAccessAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_access_analytics_collection": {
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
									"user_assessment_user_access_analytic_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": {
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

func readDataSafeUserAssessmentUserAccessAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentUserAccessAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentUserAccessAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListUserAccessAnalyticsResponse
}

func (s *DataSafeUserAssessmentUserAccessAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentUserAccessAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListUserAccessAnalyticsRequest{}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListUserAccessAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUserAccessAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeUserAssessmentUserAccessAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUserAssessmentUserAccessAnalyticsDataSource-", DataSafeUserAssessmentUserAccessAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	userAssessmentUserAccessAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UserAccessAnalyticsSummaryToMap(item))
	}
	userAssessmentUserAccessAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeUserAssessmentUserAccessAnalyticsDataSource().Schema["user_access_analytics_collection"].Elem.(*schema.Resource).Schema)
		userAssessmentUserAccessAnalytic["items"] = items
	}

	resources = append(resources, userAssessmentUserAccessAnalytic)
	if err := s.D.Set("user_access_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func UserAccessAnalyticsSummaryToMap(obj oci_data_safe.UserAccessAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["user_assessment_user_access_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	return result
}
