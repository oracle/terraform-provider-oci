// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_waas "github.com/oracle/oci-go-sdk/waas"
)

func init() {
	RegisterDatasource("oci_waas_address_list", WaasAddressListDataSource())
}

func WaasAddressListDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["address_list_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(WaasAddressListResource(), fieldMap, readSingularWaasAddressList)
}

func readSingularWaasAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &WaasAddressListDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "waas")

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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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
