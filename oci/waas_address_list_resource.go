// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_waas "github.com/oracle/oci-go-sdk/v44/waas"
)

func init() {
	RegisterResource("oci_waas_address_list", WaasAddressListResource())
}

func WaasAddressListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createWaasAddressList,
		Read:     readWaasAddressList,
		Update:   updateWaasAddressList,
		Delete:   deleteWaasAddressList,
		Schema: map[string]*schema.Schema{
			// Required
			"addresses": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"address_count": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createWaasAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &WaasAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()

	return CreateResource(d, sync)
}

func readWaasAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &WaasAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()

	return ReadResource(sync)
}

func updateWaasAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &WaasAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()

	return UpdateResource(d, sync)
}

func deleteWaasAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &WaasAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type WaasAddressListResourceCrud struct {
	BaseCrud
	Client                 *oci_waas.WaasClient
	Res                    *oci_waas.AddressList
	DisableNotFoundRetries bool
}

func (s *WaasAddressListResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *WaasAddressListResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waas.LifecycleStatesCreating),
	}
}

func (s *WaasAddressListResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waas.LifecycleStatesActive),
	}
}

func (s *WaasAddressListResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waas.LifecycleStatesDeleting),
	}
}

func (s *WaasAddressListResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waas.LifecycleStatesDeleted),
	}
}

func (s *WaasAddressListResourceCrud) Create() error {
	request := oci_waas.CreateAddressListRequest{}

	if addresses, ok := s.D.GetOkExists("addresses"); ok {
		interfaces := addresses.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("addresses") {
			request.Addresses = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.CreateAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AddressList
	return nil
}

func (s *WaasAddressListResourceCrud) Get() error {
	request := oci_waas.GetAddressListRequest{}

	tmp := s.D.Id()
	request.AddressListId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.GetAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AddressList
	return nil
}

func (s *WaasAddressListResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waas.UpdateAddressListRequest{}

	tmp := s.D.Id()
	request.AddressListId = &tmp

	if addresses, ok := s.D.GetOkExists("addresses"); ok {
		interfaces := addresses.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("addresses") {
			request.Addresses = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.UpdateAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AddressList
	return nil
}

func (s *WaasAddressListResourceCrud) Delete() error {
	request := oci_waas.DeleteAddressListRequest{}

	tmp := s.D.Id()
	request.AddressListId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	_, err := s.Client.DeleteAddressList(context.Background(), request)
	return err
}

func (s *WaasAddressListResourceCrud) SetData() error {
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

func (s *WaasAddressListResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waas.ChangeAddressListCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AddressListId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	_, err := s.Client.ChangeAddressListCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
