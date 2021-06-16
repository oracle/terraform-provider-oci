// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v42/core"
)

func init() {
	RegisterResource("oci_core_ipv6", CoreIpv6Resource())
}

func CoreIpv6Resource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreIpv6,
		Read:     readCoreIpv6,
		Update:   updateCoreIpv6,
		Delete:   deleteCoreIpv6,
		Schema: map[string]*schema.Schema{
			// Required
			"vnic_id": {
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
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
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

func createCoreIpv6(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpv6ResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return CreateResource(d, sync)
}

func readCoreIpv6(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpv6ResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

func updateCoreIpv6(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpv6ResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return UpdateResource(d, sync)
}

func deleteCoreIpv6(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpv6ResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreIpv6ResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Ipv6
	DisableNotFoundRetries bool
}

func (s *CoreIpv6ResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreIpv6ResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.Ipv6LifecycleStateProvisioning),
	}
}

func (s *CoreIpv6ResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.Ipv6LifecycleStateAvailable),
	}
}

func (s *CoreIpv6ResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.Ipv6LifecycleStateTerminating),
	}
}

func (s *CoreIpv6ResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.Ipv6LifecycleStateTerminated),
	}
}

func (s *CoreIpv6ResourceCrud) Create() error {
	request := oci_core.CreateIpv6Request{}

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

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateIpv6(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Ipv6
	return nil
}

func (s *CoreIpv6ResourceCrud) Get() error {
	request := oci_core.GetIpv6Request{}

	tmp := s.D.Id()
	request.Ipv6Id = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetIpv6(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Ipv6
	return nil
}

func (s *CoreIpv6ResourceCrud) Update() error {
	request := oci_core.UpdateIpv6Request{}

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
	request.Ipv6Id = &tmp

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateIpv6(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Ipv6
	return nil
}

func (s *CoreIpv6ResourceCrud) Delete() error {
	request := oci_core.DeleteIpv6Request{}

	tmp := s.D.Id()
	request.Ipv6Id = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteIpv6(context.Background(), request)
	return err
}

func (s *CoreIpv6ResourceCrud) SetData() error {
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

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	return nil
}
