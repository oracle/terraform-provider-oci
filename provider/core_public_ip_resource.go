// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func PublicIpResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPublicIp,
		Read:     readPublicIp,
		Update:   updatePublicIp,
		Delete:   deletePublicIp,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lifetime": {
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
			"private_ip_id": {
				Type:     schema.TypeString,
				Optional: true,
				// Computed: true, Commented out because we want to allow unsetting the value.
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scope": {
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

func createPublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &PublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readPublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &PublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updatePublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &PublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deletePublicIp(d *schema.ResourceData, m interface{}) error {
	sync := &PublicIpResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type PublicIpResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.PublicIp
	DisableNotFoundRetries bool
}

func (s *PublicIpResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *PublicIpResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateProvisioning),
		string(oci_core.PublicIpLifecycleStateAssigning),
	}
}

func (s *PublicIpResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	}
}

func (s *PublicIpResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateUnassigning),
		string(oci_core.PublicIpLifecycleStateTerminating),
	}
}

func (s *PublicIpResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateUnassigned),
		string(oci_core.PublicIpLifecycleStateTerminated),
	}
}

func (s *PublicIpResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateProvisioning),
		string(oci_core.PublicIpLifecycleStateAssigning),
		string(oci_core.PublicIpLifecycleStateUnassigning),
	}
}

func (s *PublicIpResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	}
}

func (s *PublicIpResourceCrud) Create() error {
	request := oci_core.CreatePublicIpRequest{}

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

	if lifetime, ok := s.D.GetOkExists("lifetime"); ok {
		request.Lifetime = oci_core.CreatePublicIpDetailsLifetimeEnum(lifetime.(string))
	}

	if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
		tmp := privateIpId.(string)
		request.PrivateIpId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreatePublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *PublicIpResourceCrud) Get() error {
	request := oci_core.GetPublicIpRequest{}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetPublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *PublicIpResourceCrud) Update() error {
	request := oci_core.UpdatePublicIpRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	// Wrapping in "HasChange" conditionals because the service will treat the PUT as a PATCH.
	if s.D.HasChange("display_name") {
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			request.DisplayName = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if s.D.HasChange("private_ip_id") {
		if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
			tmp := privateIpId.(string)
			request.PrivateIpId = &tmp
		}
	}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdatePublicIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PublicIp
	return nil
}

func (s *PublicIpResourceCrud) Delete() error {
	request := oci_core.DeletePublicIpRequest{}

	tmp := s.D.Id()
	request.PublicIpId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeletePublicIp(context.Background(), request)
	return err
}

func (s *PublicIpResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
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

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	s.D.Set("lifetime", s.Res.Lifetime)

	if s.Res.PrivateIpId != nil {
		s.D.Set("private_ip_id", *s.Res.PrivateIpId)
	}

	s.D.Set("scope", s.Res.Scope)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
