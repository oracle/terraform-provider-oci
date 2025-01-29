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

func TenantmanagercontrolplaneLinkDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularTenantmanagercontrolplaneLink,
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
			"parent_tenancy_id": {
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

func readSingularTenantmanagercontrolplaneLink(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneLinkDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LinkClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneLinkDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.LinkClient
	Res    *oci_tenantmanagercontrolplane.GetLinkResponse
}

func (s *TenantmanagercontrolplaneLinkDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneLinkDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetLinkRequest{}

	if linkId, ok := s.D.GetOkExists("link_id"); ok {
		tmp := linkId.(string)
		request.LinkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetLink(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneLinkDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ChildTenancyId != nil {
		s.D.Set("child_tenancy_id", *s.Res.ChildTenancyId)
	}

	if s.Res.ParentTenancyId != nil {
		s.D.Set("parent_tenancy_id", *s.Res.ParentTenancyId)
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
