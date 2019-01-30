// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func NatGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createNatGateway,
		Read:     readNatGateway,
		Update:   updateNatGateway,
		Delete:   deleteNatGateway,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"block_traffic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"nat_ip": {
				Type:     schema.TypeString,
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

func createNatGateway(d *schema.ResourceData, m interface{}) error {
	sync := &NatGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readNatGateway(d *schema.ResourceData, m interface{}) error {
	sync := &NatGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateNatGateway(d *schema.ResourceData, m interface{}) error {
	sync := &NatGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteNatGateway(d *schema.ResourceData, m interface{}) error {
	sync := &NatGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type NatGatewayResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.NatGateway
	DisableNotFoundRetries bool
}

func (s *NatGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NatGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.NatGatewayLifecycleStateProvisioning),
	}
}

func (s *NatGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.NatGatewayLifecycleStateAvailable),
	}
}

func (s *NatGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.NatGatewayLifecycleStateTerminating),
	}
}

func (s *NatGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.NatGatewayLifecycleStateTerminated),
	}
}

func (s *NatGatewayResourceCrud) Create() error {
	request := oci_core.CreateNatGatewayRequest{}

	if blockTraffic, ok := s.D.GetOkExists("block_traffic"); ok {
		tmp := blockTraffic.(bool)
		request.BlockTraffic = &tmp
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

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateNatGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NatGateway
	return nil
}

func (s *NatGatewayResourceCrud) Get() error {
	request := oci_core.GetNatGatewayRequest{}

	tmp := s.D.Id()
	request.NatGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetNatGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NatGateway
	return nil
}

func (s *NatGatewayResourceCrud) Update() error {
	request := oci_core.UpdateNatGatewayRequest{}

	if blockTraffic, ok := s.D.GetOkExists("block_traffic"); ok {
		tmp := blockTraffic.(bool)
		request.BlockTraffic = &tmp
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

	tmp := s.D.Id()
	request.NatGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateNatGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NatGateway
	return nil
}

func (s *NatGatewayResourceCrud) Delete() error {
	request := oci_core.DeleteNatGatewayRequest{}

	tmp := s.D.Id()
	request.NatGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteNatGateway(context.Background(), request)
	return err
}

func (s *NatGatewayResourceCrud) SetData() error {
	if s.Res.BlockTraffic != nil {
		s.D.Set("block_traffic", *s.Res.BlockTraffic)
	}

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

	if s.Res.NatIp != nil {
		s.D.Set("nat_ip", *s.Res.NatIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
