// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeListUserGrantsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeListUserGrants,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"depth_level": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"depth_level_greater_than_or_equal_to": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"depth_level_less_than": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"grant_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"grant_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"grants": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"depth_level": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"grant_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"privilege_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"privilege_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeListUserGrants(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeListUserGrantsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeListUserGrantsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListGrantsResponse
}

func (s *DataSafeListUserGrantsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeListUserGrantsDataSourceCrud) Get() error {
	request := oci_data_safe.ListGrantsRequest{}

	if depthLevel, ok := s.D.GetOkExists("depth_level"); ok {
		tmp := depthLevel.(int)
		request.DepthLevel = &tmp
	}

	if depthLevelGreaterThanOrEqualTo, ok := s.D.GetOkExists("depth_level_greater_than_or_equal_to"); ok {
		tmp := depthLevelGreaterThanOrEqualTo.(int)
		request.DepthLevelGreaterThanOrEqualTo = &tmp
	}

	if depthLevelLessThan, ok := s.D.GetOkExists("depth_level_less_than"); ok {
		tmp := depthLevelLessThan.(int)
		request.DepthLevelLessThan = &tmp
	}

	if grantKey, ok := s.D.GetOkExists("grant_key"); ok {
		tmp := grantKey.(string)
		request.GrantKey = &tmp
	}

	if grantName, ok := s.D.GetOkExists("grant_name"); ok {
		tmp := grantName.(string)
		request.GrantName = &tmp
	}

	if privilegeCategory, ok := s.D.GetOkExists("privilege_category"); ok {
		tmp := privilegeCategory.(string)
		request.PrivilegeCategory = &tmp
	}

	if privilegeType, ok := s.D.GetOkExists("privilege_type"); ok {
		tmp := privilegeType.(string)
		request.PrivilegeType = &tmp
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	if userKey, ok := s.D.GetOkExists("user_key"); ok {
		tmp := userKey.(string)
		request.UserKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListGrants(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGrants(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeListUserGrantsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeListUserGrantsDataSource-", DataSafeListUserGrantsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		listUserGrant := map[string]interface{}{}

		if r.DepthLevel != nil {
			listUserGrant["depth_level"] = *r.DepthLevel
		}

		if r.GrantName != nil {
			listUserGrant["grant_name"] = *r.GrantName
		}

		if r.Key != nil {
			listUserGrant["key"] = *r.Key
		}

		listUserGrant["privilege_category"] = r.PrivilegeCategory

		listUserGrant["privilege_type"] = r.PrivilegeType

		resources = append(resources, listUserGrant)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeListUserGrantsDataSource().Schema["grants"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("grants", resources); err != nil {
		return err
	}

	return nil
}
