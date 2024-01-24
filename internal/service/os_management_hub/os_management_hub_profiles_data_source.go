// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubProfiles,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"arch_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OsManagementHubProfileResource()),
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubProfiles(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.OnboardingClient
	Res    *oci_os_management_hub.ListProfilesResponse
}

func (s *OsManagementHubProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubProfilesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListProfilesRequest{}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		request.ArchType = oci_os_management_hub.ListProfilesArchTypeEnum(archType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_os_management_hub.ListProfilesOsFamilyEnum(osFamily.(string))
	}

	if profileId, ok := s.D.GetOkExists("id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	if profileType, ok := s.D.GetOkExists("profile_type"); ok {
		interfaces := profileType.([]interface{})
		tmp := make([]oci_os_management_hub.ProfileTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ProfileTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("profile_type") {
			request.ProfileType = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_os_management_hub.ProfileLifecycleStateEnum(state.(string))
	}

	if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
		request.VendorName = oci_os_management_hub.ListProfilesVendorNameEnum(vendorName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListProfiles(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProfiles(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubProfilesDataSource-", OsManagementHubProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	profile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProfileSummaryToMap(item))
	}
	profile["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubProfilesDataSource().Schema["profile_collection"].Elem.(*schema.Resource).Schema)
		profile["items"] = items
	}

	resources = append(resources, profile)
	if err := s.D.Set("profile_collection", resources); err != nil {
		return err
	}

	return nil
}
