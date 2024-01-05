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

func DnsResolversDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsResolvers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resolvers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DnsResolverResource()),
			},
		},
	}
}

func readDnsResolvers(d *schema.ResourceData, m interface{}) error {
	sync := &DnsResolversDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsResolversDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.ListResolversResponse
}

func (s *DnsResolversDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsResolversDataSourceCrud) Get() error {
	request := oci_dns.ListResolversRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.ListResolversScopeEnum(scope.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dns.ResolverSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.ListResolvers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResolvers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DnsResolversDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsResolversDataSource-", DnsResolversDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		resolver := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AttachedVcnId != nil {
			resolver["attached_vcn_id"] = *r.AttachedVcnId
		}

		if r.DefaultViewId != nil {
			resolver["default_view_id"] = *r.DefaultViewId
		}

		if r.DefinedTags != nil {
			resolver["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			resolver["display_name"] = *r.DisplayName
		}

		resolver["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			resolver["id"] = *r.Id
		}

		if r.IsProtected != nil {
			resolver["is_protected"] = *r.IsProtected
		}

		if r.Self != nil {
			resolver["self"] = *r.Self
		}

		resolver["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			resolver["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			resolver["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, resolver)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DnsResolversDataSource().Schema["resolvers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("resolvers", resources); err != nil {
		return err
	}

	return nil
}
