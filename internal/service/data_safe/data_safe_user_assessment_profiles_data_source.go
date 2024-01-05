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

func DataSafeUserAssessmentProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUserAssessmentProfiles,
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
			"failed_login_attempts_greater_than_or_equal": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"failed_login_attempts_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inactive_account_time_greater_than_or_equal": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inactive_account_time_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_user_created": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"password_lock_time_greater_than_or_equal": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password_lock_time_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password_verification_function": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sessions_per_user_greater_than_or_equal": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sessions_per_user_less_than": {
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
			"user_count_greater_than_or_equal": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_count_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"composite_limit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connect_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu_per_call": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu_per_session": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"failed_login_attempts": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"idle_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"inactive_account_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_user_created": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"logical_reads_per_call": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logical_reads_per_session": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_grace_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_life_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_lock_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_reuse_max": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_reuse_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_rollover_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_verification_function": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_verification_function_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_sga": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sessions_per_user": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_assessment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeUserAssessmentProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListProfileSummariesResponse
}

func (s *DataSafeUserAssessmentProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentProfilesDataSourceCrud) Get() error {
	request := oci_data_safe.ListProfileSummariesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListProfileSummariesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if failedLoginAttemptsGreaterThanOrEqual, ok := s.D.GetOkExists("failed_login_attempts_greater_than_or_equal"); ok {
		tmp := failedLoginAttemptsGreaterThanOrEqual.(string)
		request.FailedLoginAttemptsGreaterThanOrEqual = &tmp
	}

	if failedLoginAttemptsLessThan, ok := s.D.GetOkExists("failed_login_attempts_less_than"); ok {
		tmp := failedLoginAttemptsLessThan.(string)
		request.FailedLoginAttemptsLessThan = &tmp
	}

	if inactiveAccountTimeGreaterThanOrEqual, ok := s.D.GetOkExists("inactive_account_time_greater_than_or_equal"); ok {
		tmp := inactiveAccountTimeGreaterThanOrEqual.(string)
		request.InactiveAccountTimeGreaterThanOrEqual = &tmp
	}

	if inactiveAccountTimeLessThan, ok := s.D.GetOkExists("inactive_account_time_less_than"); ok {
		tmp := inactiveAccountTimeLessThan.(string)
		request.InactiveAccountTimeLessThan = &tmp
	}

	if isUserCreated, ok := s.D.GetOkExists("is_user_created"); ok {
		tmp := isUserCreated.(bool)
		request.IsUserCreated = &tmp
	}

	if passwordLockTimeGreaterThanOrEqual, ok := s.D.GetOkExists("password_lock_time_greater_than_or_equal"); ok {
		tmp := passwordLockTimeGreaterThanOrEqual.(string)
		request.PasswordLockTimeGreaterThanOrEqual = &tmp
	}

	if passwordLockTimeLessThan, ok := s.D.GetOkExists("password_lock_time_less_than"); ok {
		tmp := passwordLockTimeLessThan.(string)
		request.PasswordLockTimeLessThan = &tmp
	}

	if passwordVerificationFunction, ok := s.D.GetOkExists("password_verification_function"); ok {
		tmp := passwordVerificationFunction.(string)
		request.PasswordVerificationFunction = &tmp
	}

	if profileName, ok := s.D.GetOkExists("profile_name"); ok {
		tmp := profileName.(string)
		request.ProfileName = &tmp
	}

	if sessionsPerUserGreaterThanOrEqual, ok := s.D.GetOkExists("sessions_per_user_greater_than_or_equal"); ok {
		tmp := sessionsPerUserGreaterThanOrEqual.(string)
		request.SessionsPerUserGreaterThanOrEqual = &tmp
	}

	if sessionsPerUserLessThan, ok := s.D.GetOkExists("sessions_per_user_less_than"); ok {
		tmp := sessionsPerUserLessThan.(string)
		request.SessionsPerUserLessThan = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	if userCountGreaterThanOrEqual, ok := s.D.GetOkExists("user_count_greater_than_or_equal"); ok {
		tmp := userCountGreaterThanOrEqual.(string)
		request.UserCountGreaterThanOrEqual = &tmp
	}

	if userCountLessThan, ok := s.D.GetOkExists("user_count_less_than"); ok {
		tmp := userCountLessThan.(string)
		request.UserCountLessThan = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListProfileSummaries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProfileSummaries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeUserAssessmentProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUserAssessmentProfilesDataSource-", DataSafeUserAssessmentProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		userAssessmentProfile := map[string]interface{}{
			"compartment_id":     *r.CompartmentId,
			"user_assessment_id": *r.UserAssessmentId,
		}

		if r.DefinedTags != nil {
			userAssessmentProfile["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.FailedLoginAttempts != nil {
			userAssessmentProfile["failed_login_attempts"] = *r.FailedLoginAttempts
		}

		userAssessmentProfile["freeform_tags"] = r.FreeformTags
		userAssessmentProfile["freeform_tags"] = r.FreeformTags

		if r.InactiveAccountTime != nil {
			userAssessmentProfile["inactive_account_time"] = *r.InactiveAccountTime
		}

		if r.IsUserCreated != nil {
			userAssessmentProfile["is_user_created"] = *r.IsUserCreated
		}

		if r.PasswordLockTime != nil {
			userAssessmentProfile["password_lock_time"] = *r.PasswordLockTime
		}

		if r.PasswordVerificationFunction != nil {
			userAssessmentProfile["password_verification_function"] = *r.PasswordVerificationFunction
		}

		if r.ProfileName != nil {
			userAssessmentProfile["profile_name"] = *r.ProfileName
		}

		if r.SessionsPerUser != nil {
			userAssessmentProfile["sessions_per_user"] = *r.SessionsPerUser
		}

		if r.TargetId != nil {
			userAssessmentProfile["target_id"] = *r.TargetId
		}

		if r.UserCount != nil {
			userAssessmentProfile["user_count"] = *r.UserCount
		}

		resources = append(resources, userAssessmentProfile)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeUserAssessmentProfilesDataSource().Schema["profiles"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("profiles", resources); err != nil {
		return err
	}

	return nil
}
