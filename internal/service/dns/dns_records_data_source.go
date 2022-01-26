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

func DnsRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsRecords,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),

			// Required
			"zone_name_or_id": {
				Type:       schema.TypeString,
				Required:   true,
				Deprecated: tfresource.ResourceDeprecatedForAnother("oci_dns_records", "oci_dns_rrset"),
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rtype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DnsRecordResource(),
			},
		},
	}
}

func readDnsRecords(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetZoneRecordsResponse
}

func (s *DnsRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsRecordsDataSourceCrud) Get() error {
	request := oci_dns.GetZoneRecordsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if domainContains, ok := s.D.GetOkExists("domain_contains"); ok {
		tmp := domainContains.(string)
		request.DomainContains = &tmp
	}

	if rtype, ok := s.D.GetOkExists("rtype"); ok {
		tmp := rtype.(string)
		request.Rtype = &tmp
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = oci_dns.GetZoneRecordsSortByEnum(tmp)
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		tmp := sortOrder.(string)
		request.SortOrder = oci_dns.GetZoneRecordsSortOrderEnum(tmp)
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

	response, err := s.Client.GetZoneRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.GetZoneRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}
	return nil
}

func (s *DnsRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsRecordsDataSource-", DnsRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		record := map[string]interface{}{}

		if r.Domain != nil {
			record["domain"] = *r.Domain
		}

		if r.IsProtected != nil {
			record["is_protected"] = *r.IsProtected
		}

		if r.Rdata != nil {
			record["rdata"] = *r.Rdata
		}

		if r.RecordHash != nil {
			record["record_hash"] = *r.RecordHash
		}

		if r.RrsetVersion != nil {
			record["rrset_version"] = *r.RrsetVersion
		}

		if r.Rtype != nil {
			record["rtype"] = *r.Rtype
		}

		if r.Ttl != nil {
			record["ttl"] = *r.Ttl
		}

		resources = append(resources, record)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DnsRecordsDataSource().Schema["records"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("records", resources); err != nil {
		return err
	}

	return nil
}
