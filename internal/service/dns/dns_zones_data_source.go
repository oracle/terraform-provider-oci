// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"
)

func DnsZonesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDnsZones,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope": {
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tsig_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"view_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DnsZoneResource()),
			},
		},
	}
}

func readDnsZones(d *schema.ResourceData, m interface{}) error {
	sync := &DnsZonesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

type DnsZonesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.ListZonesResponse
}

func (s *DnsZonesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DnsZonesDataSourceCrud) Get() error {
	request := oci_dns.ListZonesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_dns.ListZonesScopeEnum(scope.(string))
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = oci_dns.ListZonesSortByEnum(tmp)
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		tmp := sortOrder.(string)
		request.SortOrder = oci_dns.ListZonesSortOrderEnum(tmp)
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dns.ListZonesLifecycleStateEnum(state.(string))
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if tsigKeyId, ok := s.D.GetOkExists("tsig_key_id"); ok {
		tmp := tsigKeyId.(string)
		request.TsigKeyId = &tmp
	}

	if viewId, ok := s.D.GetOkExists("view_id"); ok {
		tmp := viewId.(string)
		request.ViewId = &tmp
	}

	if zoneType, ok := s.D.GetOkExists("zone_type"); ok {
		request.ZoneType = oci_dns.ListZonesZoneTypeEnum(zoneType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dns")

	response, err := s.Client.ListZones(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListZones(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DnsZonesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DnsZonesDataSource-", DnsZonesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		zone := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			zone["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		zone["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			zone["id"] = *r.Id
		}

		if r.IsProtected != nil {
			zone["is_protected"] = *r.IsProtected
		}

		if r.Name != nil {
			zone["name"] = *r.Name
		}

		zone["scope"] = r.Scope

		if r.Self != nil {
			zone["self"] = *r.Self
		}

		if r.Serial != nil {
			zone["serial"] = *r.Serial
		}

		zone["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			zone["time_created"] = r.TimeCreated.String()
		}

		if r.Version != nil {
			zone["version"] = *r.Version
		}

		if r.ViewId != nil {
			zone["view_id"] = *r.ViewId
		}

		zone["zone_type"] = r.ZoneType

		resources = append(resources, zone)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DnsZonesDataSource().Schema["zones"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("zones", resources); err != nil {
		return err
	}

	return nil
}
