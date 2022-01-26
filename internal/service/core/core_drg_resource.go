// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v56/workrequests"
)

func CoreDrgResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrg,
		Read:     readCoreDrg,
		Update:   updateCoreDrg,
		Delete:   deleteCoreDrg,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"default_drg_route_tables": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ipsec_tunnel": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_peering_connection": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vcn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_circuit": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"default_export_drg_route_distribution_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundancy_status": {
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

func createCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDrg(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreDrgResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Drg
	RedundancyStatus       *oci_core.DrgRedundancyStatus
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreDrgResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDrgResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgLifecycleStateProvisioning),
	}
}

func (s *CoreDrgResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgLifecycleStateAvailable),
	}
}

func (s *CoreDrgResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgLifecycleStateTerminating),
	}
}

func (s *CoreDrgResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgLifecycleStateTerminated),
	}
}

func (s *CoreDrgResourceCrud) Create() error {
	request := oci_core.CreateDrgRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *CoreDrgResourceCrud) Get() error {
	request := oci_core.GetDrgRequest{}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg

	statusRequest := oci_core.GetDrgRedundancyStatusRequest{}
	statusRequest.DrgId = &tmp

	if redundancyStatusResponse, err := s.Client.GetDrgRedundancyStatus(context.Background(), statusRequest); err == nil {
		s.RedundancyStatus = &redundancyStatusResponse.DrgRedundancyStatus
	}

	return err
}

func (s *CoreDrgResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateDrgRequest{}

	if defaultDrgRouteTables, ok := s.D.GetOkExists("default_drg_route_tables"); ok && s.D.HasChange("default_drg_route_tables") {
		if tmpList := defaultDrgRouteTables.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "default_drg_route_tables", 0)
			tmp, err := s.mapToDefaultDrgRouteTables(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DefaultDrgRouteTables = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.DrgId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *CoreDrgResourceCrud) Delete() error {
	request := oci_core.DeleteDrgRequest{}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDrg(context.Background(), request)
	return err
}

func (s *CoreDrgResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultDrgRouteTables != nil {
		s.D.Set("default_drg_route_tables", []interface{}{DefaultDrgRouteTablesToMap(s.Res.DefaultDrgRouteTables)})
	} else {
		s.D.Set("default_drg_route_tables", nil)
	}

	if s.Res.DefaultExportDrgRouteDistributionId != nil {
		s.D.Set("default_export_drg_route_distribution_id", *s.Res.DefaultExportDrgRouteDistributionId)
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

	if s.RedundancyStatus != nil {
		s.D.Set("redundancy_status", s.RedundancyStatus.Status)
	}

	return nil
}

func (s *CoreDrgResourceCrud) mapToDefaultDrgRouteTables(fieldKeyFormat string) (oci_core.DefaultDrgRouteTables, error) {
	result := oci_core.DefaultDrgRouteTables{}

	if ipsecTunnel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipsec_tunnel")); ok {
		tmp := ipsecTunnel.(string)
		result.IpsecTunnel = &tmp
	}

	if remotePeeringConnection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remote_peering_connection")); ok {
		tmp := remotePeeringConnection.(string)
		result.RemotePeeringConnection = &tmp
	}

	if vcn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn")); ok {
		tmp := vcn.(string)
		result.Vcn = &tmp
	}

	if virtualCircuit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "virtual_circuit")); ok {
		tmp := virtualCircuit.(string)
		result.VirtualCircuit = &tmp
	}

	return result, nil
}

func DefaultDrgRouteTablesToMap(obj *oci_core.DefaultDrgRouteTables) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpsecTunnel != nil {
		result["ipsec_tunnel"] = string(*obj.IpsecTunnel)
	}

	if obj.RemotePeeringConnection != nil {
		result["remote_peering_connection"] = string(*obj.RemotePeeringConnection)
	}

	if obj.Vcn != nil {
		result["vcn"] = string(*obj.Vcn)
	}

	if obj.VirtualCircuit != nil {
		result["virtual_circuit"] = string(*obj.VirtualCircuit)
	}

	return result
}

func (s *CoreDrgResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeDrgCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DrgId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeDrgCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	// Workaround: Sleep for some time before polling the configuration. Because Update happens asynchronously, polling too
	// soon may result in service returning stale configuration values.
	time.Sleep(time.Second * 20)
	request := oci_core.GetDrgRequest{}
	tmp := s.D.Id()
	request.DrgId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	_, err = s.Client.GetDrg(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}
