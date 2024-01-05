// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dns "github.com/oracle/oci-go-sdk/v65/dns"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DnsRrsetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsRrsets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"view_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone_name_or_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"rrsets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"rtype": {
							Type:     schema.TypeString,
							Required: true,
						},
						"items": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"domain": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
									},
									"rdata": {
										Type:     schema.TypeString,
										Required: true,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											rtype := d.Get("rtype").(string)
											return normalizeRData(rtype, new) == normalizeRData(rtype, old)
										},
									},
									"rtype": {
										Type:     schema.TypeString,
										Required: true,
									},
									"ttl": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
									"is_protected": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"record_hash": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"rrset_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDnsRrsets(d *schema.ResourceData, m interface{}) error {
	sync := &DnsRrsetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsRrsetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.GetZoneRecordsResponse
}

func (s *DnsRrsetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsRrsetsDataSourceCrud) Get() error {
	request := oci_dns.GetZoneRecordsRequest{}

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

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.GetZoneRecordsScopeEnum(scope.(string))
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

	response, err := s.Client.GetZoneRecords(context.Background(), request)
	if err != nil {
		return err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.GetZoneRecords(context.Background(), request)
		if err != nil {
			return err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	s.Res = &response
	return nil
}

func (s *DnsRrsetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsRrsetsDataSource-", DnsRrsetsDataSource(), s.D))

	rrsetMap := map[string]map[string][]oci_dns.Record{}

	for _, record := range s.Res.Items {
		if _, ok := rrsetMap[*record.Domain]; !ok {
			rrsetMap[*record.Domain] = map[string][]oci_dns.Record{}
		}
		if _, ok := rrsetMap[*record.Domain][*record.Rtype]; !ok {
			rrsetMap[*record.Domain][*record.Rtype] = []oci_dns.Record{}
		}
		rrsetMap[*record.Domain][*record.Rtype] = append(rrsetMap[*record.Domain][*record.Rtype], record)
	}

	rrsets := []interface{}{}
	for domain, domainRrsets := range rrsetMap {
		for rtype, rrset := range domainRrsets {
			rrsetMap := map[string]interface{}{}
			rrsetMap["domain"] = domain
			rrsetMap["rtype"] = rtype
			records := []interface{}{}
			for _, record := range rrset {
				records = append(records, RecordToMap(record))
			}
			rrsetMap["items"] = records
			rrsets = append(rrsets, rrsetMap)
		}
	}
	s.D.Set("rrsets", rrsets)

	return nil
}
