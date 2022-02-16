// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v58/waf"
)

func WafNetworkAddressListDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_address_list_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WafNetworkAddressListResource(), fieldMap, readSingularWafNetworkAddressList)
}

func readSingularWafNetworkAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &WafNetworkAddressListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

type WafNetworkAddressListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waf.WafClient
	Res    *oci_waf.GetNetworkAddressListResponse
}

func (s *WafNetworkAddressListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WafNetworkAddressListDataSourceCrud) Get() error {
	request := oci_waf.GetNetworkAddressListRequest{}

	if networkAddressListId, ok := s.D.GetOkExists("network_address_list_id"); ok {
		tmp := networkAddressListId.(string)
		request.NetworkAddressListId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waf")

	response, err := s.Client.GetNetworkAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WafNetworkAddressListDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	// Set common data between all NAL types
	s.D.SetId(*s.Res.GetId())

	if s.Res.GetCompartmentId() != nil {
		s.D.Set("compartment_id", *s.Res.GetCompartmentId())
	}

	if s.Res.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.GetDefinedTags()))
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

	if s.Res.GetLifecycleDetails() != nil {
		s.D.Set("lifecycle_details", *s.Res.GetLifecycleDetails())
	}

	s.D.Set("state", s.Res.GetLifecycleState())

	if s.Res.GetSystemTags() != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.GetSystemTags()))
	}

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.GetTimeUpdated().String())
	}

	switch v := (s.Res.NetworkAddressList).(type) {
	case oci_waf.NetworkAddressListAddresses:
		s.D.Set("type", "ADDRESSES")
		if v.Addresses != nil {
			s.D.Set("addresses", v.Addresses)
		}
	case oci_waf.NetworkAddressListVcnAddresses:
		s.D.Set("type", "VCN_ADDRESSES")

		if v.VcnAddresses != nil {
			vcnAddresses := []interface{}{}
			for _, item := range v.VcnAddresses {
				vcnAddresses = append(vcnAddresses, PrivateAddressesToMap(item))
			}
			s.D.Set("vcn_addresses", vcnAddresses)
		}
	default:
		log.Printf("[WARN] Received unknown 'type': %v", s.Res.NetworkAddressList)
		return nil
	}

	return nil
}
