// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeUserAssessmentUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUserAssessmentUsers,
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
			"are_all_schemas_accessible": {
				Type:     schema.TypeBool,
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
			"schema_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"user_profile": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"account_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"admin_roles": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"are_all_schemas_accessible": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"authentication_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"schema_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_login": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_password_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_user_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_profile": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func readDataSafeUserAssessmentUsers(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListUsersResponse
}

func (s *DataSafeUserAssessmentUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentUsersDataSourceCrud) Get() error {
	request := oci_data_safe.ListUsersRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListUsersAccessLevelEnum(accessLevel.(string))
	}

	if accountStatus, ok := s.D.GetOkExists("account_status"); ok {
		tmp := accountStatus.(string)
		request.AccountStatus = &tmp
	}

	if areAllSchemasAccessible, ok := s.D.GetOkExists("are_all_schemas_accessible"); ok {
		tmp := areAllSchemasAccessible.(bool)
		request.AreAllSchemasAccessible = &tmp
	}

	if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
		tmp := authenticationType.(string)
		request.AuthenticationType = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if schemaList, ok := s.D.GetOkExists("schema_list"); ok {
		interfaces := schemaList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schema_list") {
			request.SchemaList = tmp
		}
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

	if userProfile, ok := s.D.GetOkExists("user_profile"); ok {
		tmp := userProfile.(string)
		request.UserProfile = &tmp
	}

	if userRole, ok := s.D.GetOkExists("user_role"); ok {
		tmp := userRole.(string)
		request.UserRole = &tmp
	}

	if userType, ok := s.D.GetOkExists("user_type"); ok {
		tmp := userType.(string)
		request.UserType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeUserAssessmentUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUserAssessmentUsersDataSource-", DataSafeUserAssessmentUsersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		userAssessmentUser := map[string]interface{}{}

		userAssessmentUser["account_status"] = r.AccountStatus

		userAssessmentUser["admin_roles"] = r.AdminRoles

		if r.AreAllSchemasAccessible != nil {
			userAssessmentUser["are_all_schemas_accessible"] = *r.AreAllSchemasAccessible
		}

		userAssessmentUser["authentication_type"] = r.AuthenticationType

		if r.Key != nil {
			userAssessmentUser["key"] = *r.Key
		}

		userAssessmentUser["schema_list"] = r.SchemaList

		if r.TargetId != nil {
			userAssessmentUser["target_id"] = *r.TargetId
		}

		if r.TimeLastLogin != nil {
			userAssessmentUser["time_last_login"] = r.TimeLastLogin.String()
		}

		if r.TimePasswordChanged != nil {
			userAssessmentUser["time_password_changed"] = r.TimePasswordChanged.String()
		}

		if r.TimeUserCreated != nil {
			userAssessmentUser["time_user_created"] = r.TimeUserCreated.String()
		}

		userAssessmentUser["user_category"] = r.UserCategory

		if r.UserName != nil {
			userAssessmentUser["user_name"] = *r.UserName
		}

		if r.UserProfile != nil {
			userAssessmentUser["user_profile"] = *r.UserProfile
		}

		userAssessmentUser["user_types"] = r.UserTypes

		resources = append(resources, userAssessmentUser)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeUserAssessmentUsersDataSource().Schema["users"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("users", resources); err != nil {
		return err
	}

	return nil
}
