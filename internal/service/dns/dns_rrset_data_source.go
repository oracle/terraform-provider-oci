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

func DnsRrsetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["domain"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["rtype"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["scope"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["view_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["zone_name_or_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["zone_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DnsRrsetResource(), fieldMap, readSingularDnsRrset)
}

func readSingularDnsRrset(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsRrsetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetRRSetResponse
}

func (s *DnsRrsetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsRrsetDataSourceCrud) Get() error {
	request := oci_dns.GetRRSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetRRSetScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	if zoneVersion, ok := s.D.GetOkExists("zone_version"); ok {
		tmp := zoneVersion.(string)
		request.ZoneVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.GetRRSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsRrsetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsRrsetDataSource-", DnsRrsetDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RecordToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
