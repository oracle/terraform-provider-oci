// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_dns "github.com/oracle/oci-go-sdk/dns"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ZonesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readZones,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),

			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
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
			"zone_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ZoneResource(),
			},
		},
	}
}

func readZones(d *schema.ResourceData, m interface{}) error {
	sync := &ZonesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dnsClient

	return crud.ReadResource(sync)
}

type ZonesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dns.DnsClient
	Res    *oci_dns.ListZonesResponse
}

func (s *ZonesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ZonesDataSourceCrud) Get() error {
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

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = oci_dns.ListZonesSortByEnum(tmp)
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		tmp := sortOrder.(string)
		request.SortOrder = oci_dns.ListZonesSortOrderEnum(tmp)
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := state.(string)
		request.LifecycleState = oci_dns.ListZonesLifecycleStateEnum(tmp)
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

	if zoneType, ok := s.D.GetOkExists("zone_type"); ok {
		request.ZoneType = oci_dns.ListZonesZoneTypeEnum(zoneType.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "dns")

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

func (s *ZonesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		zone := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Id != nil {
			zone["id"] = *r.Id
		}

		if r.Name != nil {
			zone["name"] = *r.Name
		}

		if r.Self != nil {
			zone["self"] = *r.Self
		}

		if r.Serial != nil {
			zone["serial"] = *r.Serial
		}

		zone["time_created"] = r.TimeCreated.String()

		if r.Version != nil {
			zone["version"] = *r.Version
		}

		zone["zone_type"] = r.ZoneType

		resources = append(resources, zone)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ZonesDataSource().Schema["zones"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("zones", resources); err != nil {
		panic(err)
	}

	return
}
