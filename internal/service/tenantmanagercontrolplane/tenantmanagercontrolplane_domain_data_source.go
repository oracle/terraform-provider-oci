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

func TenantmanagercontrolplaneDomainDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularTenantmanagercontrolplaneDomain,
		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"owner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
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
			"txt_record": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularTenantmanagercontrolplaneDomain(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneDomainDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DomainClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneDomainDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.DomainClient
	Res    *oci_tenantmanagercontrolplane.GetDomainResponse
}

func (s *TenantmanagercontrolplaneDomainDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneDomainDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetDomainRequest{}

	if domainId, ok := s.D.GetOkExists("domain_id"); ok {
		tmp := domainId.(string)
		request.DomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneDomainDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DomainName != nil {
		s.D.Set("domain_name", *s.Res.DomainName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.OwnerId != nil {
		s.D.Set("owner_id", *s.Res.OwnerId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TxtRecord != nil {
		s.D.Set("txt_record", *s.Res.TxtRecord)
	}

	return nil
}
