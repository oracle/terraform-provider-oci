// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"
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
	switch v := (s.Res.NetworkAddressList).(type) {
	case oci_waf.NetworkAddressListAddresses:
		s.D.Set("type", "ADDRESSES")

		s.D.Set("addresses", v.Addresses)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_waf.NetworkAddressListVcnAddresses:
		s.D.Set("type", "VCN_ADDRESSES")

		vcnAddresses := []interface{}{}
		for _, item := range v.VcnAddresses {
			vcnAddresses = append(vcnAddresses, PrivateAddressesToMap(item))
		}
		s.D.Set("vcn_addresses", vcnAddresses)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.NetworkAddressList)
		return nil
	}

	return nil
}
