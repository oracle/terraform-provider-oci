// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementOnboardingPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementOnboardingPolicies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"onboarding_policy_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"statements": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readFleetAppsManagementOnboardingPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementOnboardingPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementOnboardingPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementAdminClient
	Res    *oci_fleet_apps_management.ListOnboardingPoliciesResponse
}

func (s *FleetAppsManagementOnboardingPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementOnboardingPoliciesDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListOnboardingPoliciesRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListOnboardingPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOnboardingPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementOnboardingPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementOnboardingPoliciesDataSource-", FleetAppsManagementOnboardingPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	onboardingPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OnboardingPolicySummaryToMap(item))
	}
	onboardingPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementOnboardingPoliciesDataSource().Schema["onboarding_policy_collection"].Elem.(*schema.Resource).Schema)
		onboardingPolicy["items"] = items
	}

	resources = append(resources, onboardingPolicy)
	if err := s.D.Set("onboarding_policy_collection", resources); err != nil {
		return err
	}

	return nil
}

func OnboardingPolicySummaryToMap(obj oci_fleet_apps_management.OnboardingPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["statements"] = obj.Statements

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
