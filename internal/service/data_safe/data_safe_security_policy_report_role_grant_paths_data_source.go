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

func DataSafeSecurityPolicyReportRoleGrantPathsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityPolicyReportRoleGrantPaths,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"granted_role": {
				Type:     schema.TypeString,
				Required: true,
			},
			"grantee": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_policy_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_grant_path_collection": {
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
									"depth_level": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"granted_role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"grantee": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
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

func readDataSafeSecurityPolicyReportRoleGrantPaths(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyReportRoleGrantPathsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityPolicyReportRoleGrantPathsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListRoleGrantPathsResponse
}

func (s *DataSafeSecurityPolicyReportRoleGrantPathsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityPolicyReportRoleGrantPathsDataSourceCrud) Get() error {
	request := oci_data_safe.ListRoleGrantPathsRequest{}

	if grantedRole, ok := s.D.GetOkExists("granted_role"); ok {
		tmp := grantedRole.(string)
		request.GrantedRole = &tmp
	}

	if grantee, ok := s.D.GetOkExists("grantee"); ok {
		tmp := grantee.(string)
		request.Grantee = &tmp
	}

	if securityPolicyReportId, ok := s.D.GetOkExists("security_policy_report_id"); ok {
		tmp := securityPolicyReportId.(string)
		request.SecurityPolicyReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListRoleGrantPaths(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRoleGrantPaths(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityPolicyReportRoleGrantPathsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityPolicyReportRoleGrantPathsDataSource-", DataSafeSecurityPolicyReportRoleGrantPathsDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityPolicyReportRoleGrantPath := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RoleGrantPathSummaryToMap(item))
	}
	securityPolicyReportRoleGrantPath["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityPolicyReportRoleGrantPathsDataSource().Schema["role_grant_path_collection"].Elem.(*schema.Resource).Schema)
		securityPolicyReportRoleGrantPath["items"] = items
	}

	resources = append(resources, securityPolicyReportRoleGrantPath)
	if err := s.D.Set("role_grant_path_collection", resources); err != nil {
		return err
	}

	return nil
}

func RoleGrantPathSummaryToMap(obj oci_data_safe.RoleGrantPathSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DepthLevel != nil {
		result["depth_level"] = int(*obj.DepthLevel)
	}

	if obj.GrantedRole != nil {
		result["granted_role"] = string(*obj.GrantedRole)
	}

	if obj.Grantee != nil {
		result["grantee"] = string(*obj.Grantee)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	return result
}
