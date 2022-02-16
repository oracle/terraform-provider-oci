// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_waas "github.com/oracle/oci-go-sdk/v58/waas"
)

func WaasAddressListsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWaasAddressLists,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"states": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_lists": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(WaasAddressListResource()),
			},
		},
	}
}

func readWaasAddressLists(d *schema.ResourceData, m interface{}) error {
	sync := &WaasAddressListsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasAddressListsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.ListAddressListsResponse
}

func (s *WaasAddressListsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasAddressListsDataSourceCrud) Get() error {
	request := oci_waas.ListAddressListsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if ids, ok := s.D.GetOkExists("ids"); ok {
		interfaces := ids.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("ids") {
			request.Id = tmp
		}
	}

	if names, ok := s.D.GetOkExists("names"); ok {
		interfaces := names.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("names") {
			request.Name = tmp
		}
	}

	if states, ok := s.D.GetOkExists("states"); ok {
		interfaces := states.([]interface{})
		tmp := make([]oci_waas.LifecycleStatesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waas.LifecycleStatesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("states") {
			request.LifecycleState = tmp
		}
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.ListAddressLists(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAddressLists(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WaasAddressListsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WaasAddressListsDataSource-", WaasAddressListsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		addressList := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AddressCount != nil {
			addressList["address_count"] = *r.AddressCount
		}

		if r.DefinedTags != nil {
			addressList["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			addressList["display_name"] = *r.DisplayName
		}

		addressList["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			addressList["id"] = *r.Id
		}

		addressList["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			addressList["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, addressList)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, WaasAddressListsDataSource().Schema["address_lists"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("address_lists", resources); err != nil {
		return err
	}

	return nil
}
