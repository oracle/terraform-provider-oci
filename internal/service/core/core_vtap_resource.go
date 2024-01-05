// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreVtapResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVtap,
		Read:     readCoreVtap,
		Update:   updateCoreVtap,
		Delete:   deleteCoreVtap,
		Schema: map[string]*schema.Schema{
			// Required
			"capture_filter_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_id": {
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encapsulation_protocol": {
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
			"is_vtap_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"max_packet_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_private_endpoint_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_private_endpoint_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_network_identifier": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},

			// Computed
			"lifecycle_state_details": {
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

func createCoreVtap(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVtapResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreVtap(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVtapResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreVtap(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVtapResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreVtap(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVtapResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreVtapResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Vtap
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreVtapResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVtapResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VtapLifecycleStateProvisioning),
	}
}

func (s *CoreVtapResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VtapLifecycleStateAvailable),
	}
}

func (s *CoreVtapResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VtapLifecycleStateTerminating),
	}
}

func (s *CoreVtapResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VtapLifecycleStateTerminated),
	}
}

func (s *CoreVtapResourceCrud) Create() error {
	request := oci_core.CreateVtapRequest{}

	if captureFilterId, ok := s.D.GetOkExists("capture_filter_id"); ok {
		tmp := captureFilterId.(string)
		request.CaptureFilterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if encapsulationProtocol, ok := s.D.GetOkExists("encapsulation_protocol"); ok {
		request.EncapsulationProtocol = oci_core.CreateVtapDetailsEncapsulationProtocolEnum(encapsulationProtocol.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isVtapEnabled, ok := s.D.GetOkExists("is_vtap_enabled"); ok {
		tmp := isVtapEnabled.(bool)
		request.IsVtapEnabled = &tmp
	}

	if maxPacketSize, ok := s.D.GetOkExists("max_packet_size"); ok {
		tmp := maxPacketSize.(int)
		request.MaxPacketSize = &tmp
	}

	if sourceId, ok := s.D.GetOkExists("source_id"); ok {
		tmp := sourceId.(string)
		request.SourceId = &tmp
	}

	if sourcePrivateEndpointIp, ok := s.D.GetOkExists("source_private_endpoint_ip"); ok {
		tmp := sourcePrivateEndpointIp.(string)
		request.SourcePrivateEndpointIp = &tmp
	}

	if sourcePrivateEndpointSubnetId, ok := s.D.GetOkExists("source_private_endpoint_subnet_id"); ok {
		tmp := sourcePrivateEndpointSubnetId.(string)
		request.SourcePrivateEndpointSubnetId = &tmp
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		request.SourceType = oci_core.CreateVtapDetailsSourceTypeEnum(sourceType.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetIp, ok := s.D.GetOkExists("target_ip"); ok {
		tmp := targetIp.(string)
		request.TargetIp = &tmp
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_core.CreateVtapDetailsTargetTypeEnum(targetType.(string))
	}

	if trafficMode, ok := s.D.GetOkExists("traffic_mode"); ok {
		request.TrafficMode = oci_core.CreateVtapDetailsTrafficModeEnum(trafficMode.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	if vxlanNetworkIdentifier, ok := s.D.GetOkExists("vxlan_network_identifier"); ok {
		tmp := vxlanNetworkIdentifier.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert vxlanNetworkIdentifier string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VxlanNetworkIdentifier = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVtap(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vtap
	return nil
}

func (s *CoreVtapResourceCrud) Get() error {
	request := oci_core.GetVtapRequest{}

	tmp := s.D.Id()
	request.VtapId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVtap(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vtap
	return nil
}

func (s *CoreVtapResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateVtapRequest{}

	if captureFilterId, ok := s.D.GetOkExists("capture_filter_id"); ok {
		tmp := captureFilterId.(string)
		request.CaptureFilterId = &tmp
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

	if encapsulationProtocol, ok := s.D.GetOkExists("encapsulation_protocol"); ok {
		request.EncapsulationProtocol = oci_core.UpdateVtapDetailsEncapsulationProtocolEnum(encapsulationProtocol.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isVtapEnabled, ok := s.D.GetOkExists("is_vtap_enabled"); ok {
		tmp := isVtapEnabled.(bool)
		request.IsVtapEnabled = &tmp
	}

	if maxPacketSize, ok := s.D.GetOkExists("max_packet_size"); ok {
		tmp := maxPacketSize.(int)
		request.MaxPacketSize = &tmp
	}

	if sourceId, ok := s.D.GetOkExists("source_id"); ok {
		tmp := sourceId.(string)
		request.SourceId = &tmp
	}

	if sourcePrivateEndpointIp, ok := s.D.GetOkExists("source_private_endpoint_ip"); ok {
		tmp := sourcePrivateEndpointIp.(string)
		request.SourcePrivateEndpointIp = &tmp
	}

	if sourcePrivateEndpointSubnetId, ok := s.D.GetOkExists("source_private_endpoint_subnet_id"); ok {
		tmp := sourcePrivateEndpointSubnetId.(string)
		request.SourcePrivateEndpointSubnetId = &tmp
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		request.SourceType = oci_core.UpdateVtapDetailsSourceTypeEnum(sourceType.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetIp, ok := s.D.GetOkExists("target_ip"); ok {
		tmp := targetIp.(string)
		request.TargetIp = &tmp
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_core.UpdateVtapDetailsTargetTypeEnum(targetType.(string))
	}

	if trafficMode, ok := s.D.GetOkExists("traffic_mode"); ok {
		request.TrafficMode = oci_core.UpdateVtapDetailsTrafficModeEnum(trafficMode.(string))
	}

	tmp := s.D.Id()
	request.VtapId = &tmp

	if vxlanNetworkIdentifier, ok := s.D.GetOkExists("vxlan_network_identifier"); ok {
		tmp := vxlanNetworkIdentifier.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert vxlanNetworkIdentifier string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VxlanNetworkIdentifier = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.UpdateVtap(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *CoreVtapResourceCrud) Delete() error {
	request := oci_core.DeleteVtapRequest{}

	tmp := s.D.Id()
	request.VtapId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVtap(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *CoreVtapResourceCrud) SetData() error {
	if s.Res.CaptureFilterId != nil {
		s.D.Set("capture_filter_id", *s.Res.CaptureFilterId)
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

	s.D.Set("encapsulation_protocol", s.Res.EncapsulationProtocol)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsVtapEnabled != nil {
		s.D.Set("is_vtap_enabled", *s.Res.IsVtapEnabled)
	}

	s.D.Set("lifecycle_state_details", s.Res.LifecycleStateDetails)

	if s.Res.MaxPacketSize != nil {
		s.D.Set("max_packet_size", *s.Res.MaxPacketSize)
	}

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	if s.Res.SourcePrivateEndpointIp != nil {
		s.D.Set("source_private_endpoint_ip", *s.Res.SourcePrivateEndpointIp)
	}

	if s.Res.SourcePrivateEndpointSubnetId != nil {
		s.D.Set("source_private_endpoint_subnet_id", *s.Res.SourcePrivateEndpointSubnetId)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TargetIp != nil {
		s.D.Set("target_ip", *s.Res.TargetIp)
	}

	s.D.Set("target_type", s.Res.TargetType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("traffic_mode", s.Res.TrafficMode)

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	if s.Res.VxlanNetworkIdentifier != nil {
		s.D.Set("vxlan_network_identifier", strconv.FormatInt(*s.Res.VxlanNetworkIdentifier, 10))
	}

	return nil
}

func (s *CoreVtapResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVtapCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VtapId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeVtapCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "vtap", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
