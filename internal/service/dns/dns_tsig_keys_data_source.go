// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"
)

func DnsTsigKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsTsigKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tsig_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DnsTsigKeyResource()),
			},
		},
	}
}

func readDnsTsigKeys(d *schema.ResourceData, m interface{}) error {
	sync := &DnsTsigKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsTsigKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.ListTsigKeysResponse
}

func (s *DnsTsigKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsTsigKeysDataSourceCrud) Get() error {
	request := oci_dns.ListTsigKeysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dns.TsigKeySummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.ListTsigKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTsigKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DnsTsigKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsTsigKeysDataSource-", DnsTsigKeysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tsigKey := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Algorithm != nil {
			tsigKey["algorithm"] = *r.Algorithm
		}

		if r.DefinedTags != nil {
			tsigKey["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		tsigKey["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			tsigKey["id"] = *r.Id
		}

		if r.Name != nil {
			tsigKey["name"] = *r.Name
		}

		if r.Self != nil {
			tsigKey["self"] = *r.Self
		}

		tsigKey["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			tsigKey["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, tsigKey)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DnsTsigKeysDataSource().Schema["tsig_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tsig_keys", resources); err != nil {
		return err
	}

	return nil
}
