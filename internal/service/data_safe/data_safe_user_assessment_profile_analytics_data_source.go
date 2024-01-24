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

func DataSafeUserAssessmentProfileAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUserAssessmentProfileAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"profile_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"profile_aggregations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     schema.TypeMap,
						},
					},
				},
			},
		},
	}
}

func readDataSafeUserAssessmentProfileAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentProfileAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentProfileAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListProfileAnalyticsResponse
}

func (s *DataSafeUserAssessmentProfileAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentProfileAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListProfileAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListProfileAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if profileName, ok := s.D.GetOkExists("profile_name"); ok {
		tmp := profileName.(string)
		request.ProfileName = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListProfileAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProfileAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeUserAssessmentProfileAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUserAssessmentProfileAnalyticsDataSource-", DataSafeUserAssessmentProfileAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		userAssessmentProfileAnalytic := map[string]interface{}{}

		if r.DefinedTags != nil {
			userAssessmentProfileAnalytic["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		userAssessmentProfileAnalytic["freeform_tags"] = r.FreeformTags
		userAssessmentProfileAnalytic["freeform_tags"] = r.FreeformTags

		items := []interface{}{}
		for _, item := range r.Items {
			items = append(items, objectToMap(item))
		}
		userAssessmentProfileAnalytic["items"] = items

		resources = append(resources, userAssessmentProfileAnalytic)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeUserAssessmentProfileAnalyticsDataSource().Schema["profile_aggregations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("profile_aggregations", resources); err != nil {
		return err
	}

	return nil
}

func objectToMap(obj map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}
