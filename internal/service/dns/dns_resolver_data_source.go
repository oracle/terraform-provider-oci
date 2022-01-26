// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v56/dns"
)

func DnsResolverDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["resolver_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["scope"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DnsResolverResource(), fieldMap, readSingularDnsResolver)
}

func readSingularDnsResolver(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsResolverDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetResolverResponse
}

func (s *DnsResolverDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsResolverDataSourceCrud) Get() error {
	request := oci_dns.GetResolverRequest{}

	if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
		tmp := resolverId.(string)
		request.ResolverId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetResolverScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.GetResolver(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsResolverDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AttachedVcnId != nil {
		s.D.Set("attached_vcn_id", *s.Res.AttachedVcnId)
	}

	attachedViews := []interface{}{}
	for _, item := range s.Res.AttachedViews {
		attachedViews = append(attachedViews, AttachedViewToMap(item))
	}
	s.D.Set("attached_views", attachedViews)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultViewId != nil {
		s.D.Set("default_view_id", *s.Res.DefaultViewId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	endpoints := []interface{}{}
	for _, item := range s.Res.Endpoints {
		endpoints = append(endpoints, ResolverEndpointSummaryToMap(item))
	}
	s.D.Set("endpoints", endpoints)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, ResolverRuleToMap(item))
	}
	s.D.Set("rules", rules)

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
