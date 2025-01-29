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

func TenantmanagercontrolplaneOrganizationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularTenantmanagercontrolplaneOrganization,
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_ucm_subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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
	}
}

func readSingularTenantmanagercontrolplaneOrganization(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneOrganizationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneOrganizationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.OrganizationClient
	Res    *oci_tenantmanagercontrolplane.GetOrganizationResponse
}

func (s *TenantmanagercontrolplaneOrganizationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneOrganizationDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetOrganizationRequest{}

	if organizationId, ok := s.D.GetOkExists("organization_id"); ok {
		tmp := organizationId.(string)
		request.OrganizationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetOrganization(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneOrganizationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultUcmSubscriptionId != nil {
		s.D.Set("default_ucm_subscription_id", *s.Res.DefaultUcmSubscriptionId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ParentName != nil {
		s.D.Set("parent_name", *s.Res.ParentName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
