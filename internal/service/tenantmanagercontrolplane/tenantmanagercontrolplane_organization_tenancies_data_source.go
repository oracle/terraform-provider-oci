// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneOrganizationTenanciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneOrganizationTenancies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organization_tenancy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"governance_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_approved_for_transfer": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tenancy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_joined": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_left": {
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

func readTenantmanagercontrolplaneOrganizationTenancies(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneOrganizationTenanciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneOrganizationTenanciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.OrganizationClient
	Res    *oci_tenantmanagercontrolplane.ListOrganizationTenanciesResponse
}

func (s *TenantmanagercontrolplaneOrganizationTenanciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneOrganizationTenanciesDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListOrganizationTenanciesRequest{}

	if organizationId, ok := s.D.GetOkExists("organization_id"); ok {
		tmp := organizationId.(string)
		request.OrganizationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListOrganizationTenancies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOrganizationTenancies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneOrganizationTenanciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneOrganizationTenanciesDataSource-", TenantmanagercontrolplaneOrganizationTenanciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	organizationTenancy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OrganizationTenancySummaryToMap(item))
	}
	organizationTenancy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneOrganizationTenanciesDataSource().Schema["organization_tenancy_collection"].Elem.(*schema.Resource).Schema)
		organizationTenancy["items"] = items
	}

	resources = append(resources, organizationTenancy)
	if err := s.D.Set("organization_tenancy_collection", resources); err != nil {
		return err
	}

	return nil
}

func OrganizationTenancySummaryToMap(obj oci_tenantmanagercontrolplane.OrganizationTenancySummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["governance_status"] = string(obj.GovernanceStatus)

	if obj.IsApprovedForTransfer != nil {
		result["is_approved_for_transfer"] = bool(*obj.IsApprovedForTransfer)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["role"] = string(obj.Role)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.TimeJoined != nil {
		result["time_joined"] = obj.TimeJoined.String()
	}

	if obj.TimeLeft != nil {
		result["time_left"] = obj.TimeLeft.String()
	}

	return result
}
