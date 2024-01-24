// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
)

func DnsViewDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["scope"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["view_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DnsViewResource(), fieldMap, readSingularDnsView)
}

func readSingularDnsView(d *schema.ResourceData, m interface{}) error {
	sync := &DnsViewDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsViewDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetViewResponse
}

func (s *DnsViewDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsViewDataSourceCrud) Get() error {
	request := oci_dns.GetViewRequest{}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetViewScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.GetView(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsViewDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
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
