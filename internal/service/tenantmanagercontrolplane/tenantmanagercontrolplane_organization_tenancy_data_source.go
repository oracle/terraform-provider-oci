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

func TenantmanagercontrolplaneOrganizationTenancyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularTenantmanagercontrolplaneOrganizationTenancy,
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"time_joined": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_left": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularTenantmanagercontrolplaneOrganizationTenancy(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneOrganizationTenancyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneOrganizationTenancyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.OrganizationClient
	Res    *oci_tenantmanagercontrolplane.GetOrganizationTenancyResponse
}

func (s *TenantmanagercontrolplaneOrganizationTenancyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneOrganizationTenancyDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetOrganizationTenancyRequest{}

	if organizationId, ok := s.D.GetOkExists("organization_id"); ok {
		tmp := organizationId.(string)
		request.OrganizationId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetOrganizationTenancy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneOrganizationTenancyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneOrganizationTenancyDataSource-", TenantmanagercontrolplaneOrganizationTenancyDataSource(), s.D))

	s.D.Set("governance_status", s.Res.GovernanceStatus)

	if s.Res.IsApprovedForTransfer != nil {
		s.D.Set("is_approved_for_transfer", *s.Res.IsApprovedForTransfer)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeJoined != nil {
		s.D.Set("time_joined", s.Res.TimeJoined.String())
	}

	if s.Res.TimeLeft != nil {
		s.D.Set("time_left", s.Res.TimeLeft.String())
	}

	return nil
}
