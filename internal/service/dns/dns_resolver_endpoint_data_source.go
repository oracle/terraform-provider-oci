// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"
)

func DnsResolverEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["resolver_endpoint_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resolver_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["scope"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DnsResolverEndpointResource(), fieldMap, readSingularDnsResolverEndpoint)
}

func readSingularDnsResolverEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsResolverEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetResolverEndpointResponse
}

func (s *DnsResolverEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsResolverEndpointDataSourceCrud) Get() error {
	request := oci_dns.GetResolverEndpointRequest{}

	if resolverEndpointName, ok := s.D.GetOkExists("resolver_endpoint_name"); ok {
		tmp := resolverEndpointName.(string)
		request.ResolverEndpointName = &tmp
	}

	if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
		tmp := resolverId.(string)
		request.ResolverId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetResolverEndpointScopeEnum(scope.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.GetResolverEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DnsResolverEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsResolverEndpointDataSource-", DnsResolverEndpointDataSource(), s.D))
	switch v := (s.Res.ResolverEndpoint).(type) {
	case oci_dns.ResolverVnicEndpoint:
		s.D.Set("endpoint_type", "VNIC")

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ForwardingAddress != nil {
			s.D.Set("forwarding_address", *v.ForwardingAddress)
		}

		if v.IsForwarding != nil {
			s.D.Set("is_forwarding", *v.IsForwarding)
		}

		if v.IsListening != nil {
			s.D.Set("is_listening", *v.IsListening)
		}

		if v.ListeningAddress != nil {
			s.D.Set("listening_address", *v.ListeningAddress)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.Self != nil {
			s.D.Set("self", *v.Self)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'endpoint_type' of unknown type %v", *s.Res)
		return nil
	}

	return nil
}
