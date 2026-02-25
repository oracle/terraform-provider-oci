// Copyright (c) 2025 Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"
)

func DnsZoneDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
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
	return tfresource.GetSingularDataSourceItemSchema(DnsZoneResource(), fieldMap, readSingularDnsZone)
}

func readSingularDnsZone(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZoneDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsZoneDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetZoneResponse
}

func (s *DnsZoneDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsZoneDataSourceCrud) Get() error {
	request := oci_dns.GetZoneRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetZoneScopeEnum(scope.(string))
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	if zoneNameOrId, ok := s.D.GetOkExists("zone_name_or_id"); ok {
		tmp := zoneNameOrId.(string)
		request.ZoneNameOrId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.GetZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsZoneDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	// Set ID to the zone OCID
	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DnssecConfig != nil {
		s.D.Set("dnssec_config", []interface{}{DnssecConfigToMap(s.Res.DnssecConfig)})
	} else {
		s.D.Set("dnssec_config", nil)
	}

	s.D.Set("dnssec_state", s.Res.DnssecState)

	externalDownstreams := []interface{}{}
	for _, item := range s.Res.ExternalDownstreams {
		externalDownstreams = append(externalDownstreams, ExternalDownstreamToMap(item))
	}
	s.D.Set("external_downstreams", externalDownstreams)

	externalMasters := []interface{}{}
	for _, item := range s.Res.ExternalMasters {
		externalMasters = append(externalMasters, ExternalMasterToMap(item))
	}
	s.D.Set("external_masters", externalMasters)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("scope", s.Res.Scope)

	if s.Res.IsProtected != nil {
		s.D.Set("is_protected", *s.Res.IsProtected)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	nameservers := []interface{}{}
	for _, item := range s.Res.Nameservers {
		nameservers = append(nameservers, NameserverToMap(item))
	}
	s.D.Set("nameservers", nameservers)

	s.D.Set("resolution_mode", s.Res.ResolutionMode)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	if s.Res.Serial != nil {
		s.D.Set("serial", *s.Res.Serial)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	s.D.Set("zone_type", s.Res.ZoneType)

	if s.Res.ViewId != nil {
		s.D.Set("view_id", *s.Res.ViewId)
	}

	zoneTransferServers := []interface{}{}
	for _, item := range s.Res.ZoneTransferServers {
		zoneTransferServers = append(zoneTransferServers, ZoneTransferServerToMap(item))
	}
	s.D.Set("zone_transfer_servers", zoneTransferServers)

	return nil
}
