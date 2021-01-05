// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v31/core"
)

func init() {
	RegisterResource("oci_core_network_security_group", CoreNetworkSecurityGroupResource())
}

func CoreNetworkSecurityGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreNetworkSecurityGroup,
		Read:     readCoreNetworkSecurityGroup,
		Update:   updateCoreNetworkSecurityGroup,
		Delete:   deleteCoreNetworkSecurityGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

			// Computed
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

func createCoreNetworkSecurityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return CreateResource(d, sync)
}

func readCoreNetworkSecurityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

func updateCoreNetworkSecurityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return UpdateResource(d, sync)
}

func deleteCoreNetworkSecurityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreNetworkSecurityGroupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.NetworkSecurityGroup
	DisableNotFoundRetries bool
}

func (s *CoreNetworkSecurityGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreNetworkSecurityGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.NetworkSecurityGroupLifecycleStateProvisioning),
	}
}

func (s *CoreNetworkSecurityGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.NetworkSecurityGroupLifecycleStateAvailable),
	}
}

func (s *CoreNetworkSecurityGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.NetworkSecurityGroupLifecycleStateTerminating),
	}
}

func (s *CoreNetworkSecurityGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.NetworkSecurityGroupLifecycleStateTerminated),
	}
}

func (s *CoreNetworkSecurityGroupResourceCrud) Create() error {
	request := oci_core.CreateNetworkSecurityGroupRequest{}

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

	response, err := s.Client.CreateNetworkSecurityGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkSecurityGroup
	return nil
}

func (s *CoreNetworkSecurityGroupResourceCrud) Get() error {
	request := oci_core.GetNetworkSecurityGroupRequest{}

	tmp := s.D.Id()
	request.NetworkSecurityGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetNetworkSecurityGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkSecurityGroup
	return nil
}

func (s *CoreNetworkSecurityGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateNetworkSecurityGroupRequest{}

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
	request.NetworkSecurityGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateNetworkSecurityGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkSecurityGroup
	return nil
}

func (s *CoreNetworkSecurityGroupResourceCrud) Delete() error {
	request := oci_core.DeleteNetworkSecurityGroupRequest{}

	tmp := s.D.Id()
	request.NetworkSecurityGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteNetworkSecurityGroup(context.Background(), request)
	return err
}

func (s *CoreNetworkSecurityGroupResourceCrud) SetData() error {
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

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *CoreNetworkSecurityGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeNetworkSecurityGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NetworkSecurityGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeNetworkSecurityGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
