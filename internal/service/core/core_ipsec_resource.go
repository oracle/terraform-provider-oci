// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreIpSecConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreIpSecConnection,
		Read:     readCoreIpSecConnection,
		Update:   updateCoreIpSecConnection,
		Delete:   deleteCoreIpSecConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpe_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"static_routes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"cpe_local_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpe_local_identifier_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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

func createCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreIpSecConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreIpSecConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.IpSecConnection
	DisableNotFoundRetries bool
}

func (s *CoreIpSecConnectionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreIpSecConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateProvisioning),
	}
}

func (s *CoreIpSecConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateAvailable),
	}
}

func (s *CoreIpSecConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateTerminating),
	}
}

func (s *CoreIpSecConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateTerminated),
	}
}

func (s *CoreIpSecConnectionResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateProvisioning),
	}
}

func (s *CoreIpSecConnectionResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.IpSecConnectionLifecycleStateAvailable),
	}
}

func (s *CoreIpSecConnectionResourceCrud) Create() error {
	request := oci_core.CreateIPSecConnectionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpeId, ok := s.D.GetOkExists("cpe_id"); ok {
		tmp := cpeId.(string)
		request.CpeId = &tmp
	}

	if cpeLocalIdentifier, ok := s.D.GetOkExists("cpe_local_identifier"); ok {
		tmp := cpeLocalIdentifier.(string)
		request.CpeLocalIdentifier = &tmp
	}

	if cpeLocalIdentifierType, ok := s.D.GetOkExists("cpe_local_identifier_type"); ok {
		request.CpeLocalIdentifierType = oci_core.CreateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum(cpeLocalIdentifierType.(string))
	}

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

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.StaticRoutes = []string{}
	if staticRoutes, ok := s.D.GetOkExists("static_routes"); ok {
		interfaces := staticRoutes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.StaticRoutes = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *CoreIpSecConnectionResourceCrud) Get() error {
	request := oci_core.GetIPSecConnectionRequest{}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *CoreIpSecConnectionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateIPSecConnectionRequest{}

	if cpeLocalIdentifier, ok := s.D.GetOkExists("cpe_local_identifier"); ok {
		tmp := cpeLocalIdentifier.(string)
		request.CpeLocalIdentifier = &tmp
	}

	if cpeLocalIdentifierType, ok := s.D.GetOkExists("cpe_local_identifier_type"); ok {
		request.CpeLocalIdentifierType = oci_core.UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum(cpeLocalIdentifierType.(string))
	}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.StaticRoutes = []string{}
	if staticRoutes, ok := s.D.GetOkExists("static_routes"); ok {
		interfaces := staticRoutes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.StaticRoutes = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateIPSecConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IpSecConnection
	return nil
}

func (s *CoreIpSecConnectionResourceCrud) Delete() error {
	request := oci_core.DeleteIPSecConnectionRequest{}

	tmp := s.D.Id()
	request.IpscId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteIPSecConnection(context.Background(), request)
	return err
}

func (s *CoreIpSecConnectionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpeId != nil {
		s.D.Set("cpe_id", *s.Res.CpeId)
	}

	if s.Res.CpeLocalIdentifier != nil {
		s.D.Set("cpe_local_identifier", *s.Res.CpeLocalIdentifier)
	}

	s.D.Set("cpe_local_identifier_type", s.Res.CpeLocalIdentifierType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("static_routes", s.Res.StaticRoutes)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreIpSecConnectionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeIPSecConnectionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	tmp := s.D.Id()
	changeCompartmentRequest.IpscId = &tmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeIPSecConnectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
