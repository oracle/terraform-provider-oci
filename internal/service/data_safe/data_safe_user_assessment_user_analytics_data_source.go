// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v58/datasafe"
)

func DataSafeUserAssessmentUserAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUserAssessmentUserAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_last_login_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_last_login_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_password_last_changed_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_password_last_changed_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_user_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_user_created_less_than": {
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
			"user_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": {
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
							Type:     schema.TypeString,
							Computed: true,
							//Elem: &schema.Resource{
							// Schema: map[string]*schema.Schema{
							//    "count": {
							//       Type:     schema.TypeInt,
							//       Computed: true,
							//    },
							//    "grant_count": {
							//       Type:     schema.TypeInt,
							//       Computed: true,
							//    },
							// },
							//},

						},
					},
				},
			},
		},
	}
}

func readDataSafeUserAssessmentUserAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentUserAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentUserAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListUserAnalyticsResponse
}

func (s *DataSafeUserAssessmentUserAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentUserAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListUserAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListUserAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if accountStatus, ok := s.D.GetOkExists("account_status"); ok {
		tmp := accountStatus.(string)
		request.AccountStatus = &tmp
	}

	if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
		tmp := authenticationType.(string)
		request.AuthenticationType = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if timeLastLoginGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_last_login_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLastLoginGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeLastLoginGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeLastLoginLessThan, ok := s.D.GetOkExists("time_last_login_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLastLoginLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeLastLoginLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if timePasswordLastChangedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_password_last_changed_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timePasswordLastChangedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimePasswordLastChangedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timePasswordLastChangedLessThan, ok := s.D.GetOkExists("time_password_last_changed_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timePasswordLastChangedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimePasswordLastChangedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if timeUserCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_user_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUserCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeUserCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeUserCreatedLessThan, ok := s.D.GetOkExists("time_user_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUserCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeUserCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	if userCategory, ok := s.D.GetOkExists("user_category"); ok {
		tmp := userCategory.(string)
		request.UserCategory = &tmp
	}

	if userKey, ok := s.D.GetOkExists("user_key"); ok {
		tmp := userKey.(string)
		request.UserKey = &tmp
	}

	if userName, ok := s.D.GetOkExists("user_name"); ok {
		tmp := userName.(string)
		request.UserName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListUserAnalytics(context.Background(), request)
	fmt.Println("response in get ****************", response)

	if err != nil {
		fmt.Println("test err in get ****************", err)
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUserAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeUserAssessmentUserAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUserAssessmentUserAnalyticsDataSource-", DataSafeUserAssessmentUserAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		userAssessmentUserAnalytic := map[string]interface{}{}

		items := []interface{}{}
		for _, item := range r.Items {
			items = append(items, objectToMap(item))
		}
		userAssessmentUserAnalytic["items"] = items

		resources = append(resources, userAssessmentUserAnalytic)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeUserAssessmentUserAnalyticsDataSource().Schema["user_aggregations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("user_aggregations", resources); err != nil {
		return err
	}

	return nil
}

func objectToMap(obj map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}
