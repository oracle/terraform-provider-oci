// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneLinkTenancyNameDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularTenantmanagercontrolplaneLinkTenancyNameWithContext,
		Schema: map[string]*schema.Schema{
			"link_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"child_tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"child_tenancy_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"feature": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_tenancy_name": {
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
			"time_terminated": {
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

func readSingularTenantmanagercontrolplaneLinkTenancyNameWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &TenantmanagercontrolplaneLinkTenancyNameDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LinkClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type TenantmanagercontrolplaneLinkTenancyNameDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.LinkClient
	Res    *oci_tenantmanagercontrolplane.GetLinkWithTenancyNamesResponse
}

func (s *TenantmanagercontrolplaneLinkTenancyNameDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneLinkTenancyNameDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_tenantmanagercontrolplane.GetLinkWithTenancyNamesRequest{}

	if linkId, ok := s.D.GetOkExists("link_id"); ok {
		tmp := linkId.(string)
		request.LinkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetLinkWithTenancyNames(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneLinkTenancyNameDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ChildTenancyId != nil {
		s.D.Set("child_tenancy_id", *s.Res.ChildTenancyId)
	}

	if s.Res.ChildTenancyName != nil {
		s.D.Set("child_tenancy_name", *s.Res.ChildTenancyName)
	}

	if s.Res.Feature != nil {
		s.D.Set("feature", *s.Res.Feature)
	}

	if s.Res.ParentTenancyId != nil {
		s.D.Set("parent_tenancy_id", *s.Res.ParentTenancyId)
	}

	if s.Res.ParentTenancyName != nil {
		s.D.Set("parent_tenancy_name", *s.Res.ParentTenancyName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeTerminated != nil {
		s.D.Set("time_terminated", s.Res.TimeTerminated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
