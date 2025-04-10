// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CorePrivateIpResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCorePrivateIp,
		Read:     readCorePrivateIp,
		Update:   updateCorePrivateIp,
		Delete:   deleteCorePrivateIp,
		Schema: map[string]*schema.Schema{
			// Required

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
			"hostname_label": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"lifetime": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_table_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vnic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_primary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_reserved": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCorePrivateIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePrivateIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCorePrivateIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePrivateIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCorePrivateIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePrivateIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCorePrivateIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePrivateIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CorePrivateIpResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.PrivateIp
	DisableNotFoundRetries bool
}

func (s *CorePrivateIpResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CorePrivateIpResourceCrud) Create() error {
	request := oci_core.CreatePrivateIpRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if lifetime, ok := s.D.GetOkExists("lifetime"); ok {
		request.Lifetime = oci_core.CreatePrivateIpDetailsLifetimeEnum(lifetime.(string))
	}

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if vlanId, ok := s.D.GetOkExists("vlan_id"); ok {
		tmp := vlanId.(string)
		request.VlanId = &tmp
	}

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreatePrivateIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateIp
	return nil
}

func (s *CorePrivateIpResourceCrud) Get() error {
	request := oci_core.GetPrivateIpRequest{}

	tmp := s.D.Id()
	request.PrivateIpId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetPrivateIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateIp
	return nil
}

func (s *CorePrivateIpResourceCrud) Update() error {
	request := oci_core.UpdatePrivateIpRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if lifetime, ok := s.D.GetOkExists("lifetime"); ok {
		request.Lifetime = oci_core.UpdatePrivateIpDetailsLifetimeEnum(lifetime.(string))
	}

	tmp := s.D.Id()
	request.PrivateIpId = &tmp

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok && vnicId != "" {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	} else {
		// call detach
		request_get := oci_core.GetPrivateIpRequest{}
		request_get.PrivateIpId = &tmp

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

		response, err := s.Client.GetPrivateIp(context.Background(), request_get)
		if err != nil {
			return err
		}
		if response.Lifetime == "RESERVED" {
			if response.IpState == "ASSIGNED" {
				err := s.PrivateIpVnicDetach()
				if err != nil {
					return err
				}
				response, err := s.Client.GetPrivateIp(context.Background(), request_get)
				s.Res = &response.PrivateIp
				if err != nil {
					return err
				}
				return nil
			}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdatePrivateIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PrivateIp
	return nil
}

func (s *CorePrivateIpResourceCrud) Delete() error {
	request := oci_core.DeletePrivateIpRequest{}

	tmp := s.D.Id()
	request.PrivateIpId = &tmp

	// use detach API
	request_get := oci_core.GetPrivateIpRequest{}
	request_get.PrivateIpId = &tmp

	request_get.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err_get := s.Client.GetPrivateIp(context.Background(), request_get)
	if err_get != nil {
		return err_get
	}
	if response.Lifetime == "RESERVED" {
		if response.IpState == "ASSIGNED" {
			err := s.PrivateIpVnicDetach()
			if err != nil {
				return err
			}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeletePrivateIp(context.Background(), request)
	return err
}

func (s *CorePrivateIpResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

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

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	s.D.Set("ip_state", s.Res.IpState)

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	s.D.Set("lifetime", s.Res.Lifetime)

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	/*if s.Res.IsReserved != nil {
		s.D.Set("is_reserved", *s.Res.IsReserved)
	}*/

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VlanId != nil {
		s.D.Set("vlan_id", *s.Res.VlanId)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	return nil
}

func (s *CorePrivateIpResourceCrud) PrivateIpVnicDetach() error {
	request := oci_core.PrivateIpVnicDetachRequest{}

	idTmp := s.D.Id()
	request.PrivateIpId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.PrivateIpVnicDetach(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
