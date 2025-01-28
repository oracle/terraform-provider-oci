// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUserAssessmentPasswordExpiryDateAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"time_password_expiry_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_aggregations": {
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
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
		},
	}
}

func readDataSafeUserAssessmentPasswordExpiryDateAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListPasswordExpiryDateAnalyticsResponse
}

func (s *DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListPasswordExpiryDateAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListPasswordExpiryDateAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if timePasswordExpiryLessThan, ok := s.D.GetOkExists("time_password_expiry_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timePasswordExpiryLessThan.(string))
		if err != nil {
			return err
		}
		request.TimePasswordExpiryLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	if userCategory, ok := s.D.GetOkExists("user_category"); ok {
		tmp := userCategory.(string)
		request.UserCategory = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListPasswordExpiryDateAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSource-", DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		userAssessmentPasswordExpiryDateAnalytic := map[string]interface{}{}

		items := []interface{}{}
		for _, item := range r.Items {
			items = append(items, objectToMap(item))
		}
		userAssessmentPasswordExpiryDateAnalytic["items"] = items

		resources = append(resources, userAssessmentPasswordExpiryDateAnalytic)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSource().Schema["user_aggregations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("user_aggregations", resources); err != nil {
		return err
	}

	return nil
}
