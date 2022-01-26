// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v56/dns"
)

func DnsResolverEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsResolverEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resolver_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resolver_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DnsResolverEndpointResource(),
			},
		},
	}
}

func readDnsResolverEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolverEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsResolverEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.ListResolverEndpointsResponse
}

func (s *DnsResolverEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsResolverEndpointsDataSourceCrud) Get() error {
	request := oci_dns.ListResolverEndpointsRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resolverId, ok := s.D.GetOkExists("resolver_id"); ok {
		tmp := resolverId.(string)
		request.ResolverId = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.ListResolverEndpointsScopeEnum(scope.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dns.ResolverEndpointSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.ListResolverEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResolverEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DnsResolverEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsResolverEndpointsDataSource-", DnsResolverEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		result := map[string]interface{}{}
		switch v := (r).(type) {
		case oci_dns.ResolverVnicEndpointSummary:
			result["endpoint_type"] = "VNIC"

			if v.SubnetId != nil {
				result["subnet_id"] = string(*v.SubnetId)
			}

			if v.CompartmentId != nil {
				result["compartment_id"] = string(*v.CompartmentId)
			}

			if v.ForwardingAddress != nil {
				result["forwarding_address"] = string(*v.ForwardingAddress)
			}

			if v.IsForwarding != nil {
				result["is_forwarding"] = bool(*v.IsForwarding)
			}

			if v.IsListening != nil {
				result["is_listening"] = bool(*v.IsListening)
			}

			if v.ListeningAddress != nil {
				result["listening_address"] = string(*v.ListeningAddress)
			}

			if v.Name != nil {
				result["name"] = string(*v.Name)
			}

			if v.Self != nil {
				result["self"] = string(*v.Self)
			}

			result["state"] = string(v.LifecycleState)

			if v.TimeCreated != nil {
				result["time_created"] = v.TimeCreated.String()
			}

			if v.TimeUpdated != nil {
				result["time_updated"] = v.TimeUpdated.String()
			}
		default:
			log.Printf("[WARN] Received 'endpoint_type' of unknown type %v", r)
			return nil
		}

		resources = append(resources, result)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DnsResolverEndpointsDataSource().Schema["resolver_endpoints"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("resolver_endpoints", resources); err != nil {
		return err
	}

	return nil
}
