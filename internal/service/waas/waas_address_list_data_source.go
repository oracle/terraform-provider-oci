// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waas "github.com/oracle/oci-go-sdk/v65/waas"
)

func WaasAddressListDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["address_list_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WaasAddressListResource(), fieldMap, readSingularWaasAddressList)
}

func readSingularWaasAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &WaasAddressListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasAddressListDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.GetAddressListResponse
}

func (s *WaasAddressListDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasAddressListDataSourceCrud) Get() error {
	request := oci_waas.GetAddressListRequest{}

	if addressListId, ok := s.D.GetOkExists("address_list_id"); ok {
		tmp := addressListId.(string)
		request.AddressListId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.GetAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaasAddressListDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AddressCount != nil {
		s.D.Set("address_count", *s.Res.AddressCount)
	}

	s.D.Set("addresses", s.Res.Addresses)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
