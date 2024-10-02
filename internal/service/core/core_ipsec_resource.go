// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
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
			"tunnel_configuration": { // used for private ipsec connection on CREATE only. Will not be updated.
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oracle_tunnel_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"associated_virtual_circuits": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"drg_route_table_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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
			"transport_type": {
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

type PrivateIpSecConnectionTunnelResourceCrud struct {
	OracleTunnelIp            *string
	AssociatedVirtualCircuits []string
	DrgRouteTableId           *string
}

type CoreIpSecConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.IpSecConnection
	ResTunnels             []PrivateIpSecConnectionTunnelResourceCrud
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if tunnelConfigs, ok := s.D.GetOkExists("tunnel_configuration"); ok {
		tmpList := tunnelConfigs.([]interface{})
		if len(tmpList) > 0 {
			log.Printf("mapping first tunnel information")
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tunnel_configuration", 0)
			tmp := s.mapToTunnelConfiguration(fieldKeyFormat)
			request.TunnelConfiguration = append(request.TunnelConfiguration, tmp)
		}
		if len(tmpList) > 1 {
			log.Printf("mapping second tunnel information")
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tunnel_configuration", 1)
			tmp := s.mapToTunnelConfiguration(fieldKeyFormat)
			request.TunnelConfiguration = append(request.TunnelConfiguration, tmp)
		}
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

	// only retrieve tunnels for tunnel config info if ipsec over fastconnect
	if s.Res.TransportType == oci_core.IpSecConnectionTransportTypeFastconnect {
		err = s.GetTunnels()
		if err != nil {
			return err
		}
	}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}
	if _, ok := s.D.GetOk("tunnel_configuration"); ok && s.D.HasChange("tunnel_configuration") {
		oldValue, newValue := s.D.GetChange("tunnel_configuration")

		oldTunnels, ok1 := oldValue.([]interface{})
		newTunnels, ok2 := newValue.([]interface{})
		if !ok1 || !ok2 {
			return fmt.Errorf("Type conversion failed for tunnel, data type of tunnel is %s", reflect.TypeOf(newValue))
		}
		for index, _ := range newTunnels {
			oldTunnel := oldTunnels[index].(map[string]interface{})
			newTunnel := newTunnels[index].(map[string]interface{})

			if oldTunnel["oracle_tunnel_ip"] != newTunnel["oracle_tunnel_ip"] {
				return fmt.Errorf("oracle_tunnel_ip field cannot be updated after create ipsec connection")
			}
			if notequal := !reflect.DeepEqual(oldTunnel["associated_virtual_circuits"], newTunnel["associated_virtual_circuits"]); notequal {
				return fmt.Errorf("associated_virtual_circuits field cannot be updated after create ipsec connection")
			}
			if oldTunnel["drg_route_table_id"] != newTunnel["drg_route_table_id"] {
				return fmt.Errorf("drg_route_table_id cannot be updated through oci_core_ipsec, use oci_core_drg_attachment_management resource instead to update")
			}
		}
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

	s.D.Set("transport_type", s.Res.TransportType)

	// set tunnel_configurations if ipsec over fast connect tunnel
	if s.Res.TransportType == oci_core.IpSecConnectionTransportTypeFastconnect {
		if s.ResTunnels != nil && len(s.ResTunnels) > 0 {
			tmpList := []interface{}{}
			for _, value := range s.ResTunnels {
				t := make(map[string]interface{})
				if value.OracleTunnelIp != nil {
					t["oracle_tunnel_ip"] = *value.OracleTunnelIp
				} else {
					t["oracle_tunnel_ip"] = ""
				}
				t["associated_virtual_circuits"] = value.AssociatedVirtualCircuits
				if value.DrgRouteTableId != nil {
					t["drg_route_table_id"] = *value.DrgRouteTableId
				} else {
					t["drg_route_table_id"] = ""
				}
				tmpList = append(tmpList, t)
			}
			err := s.D.Set("tunnel_configuration", tmpList)
			if err != nil {
				log.Printf("[WARN] create_ipsec_connection tunnel configuration could not be set: %q", err)
			}
		}
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

func (s *CoreIpSecConnectionResourceCrud) mapToTunnelConfiguration(fieldKeyFormat string) oci_core.CreateIpSecConnectionTunnelDetails {
	result := oci_core.CreateIpSecConnectionTunnelDetails{}

	if oracleTunnelIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oracle_tunnel_ip")); ok {
		tmp := oracleTunnelIp.(string)
		result.OracleTunnelIp = &tmp
	}
	if drgRouteTableId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "drg_route_table_id")); ok {
		tmp := drgRouteTableId.(string)
		result.DrgRouteTableId = &tmp
	}
	if virtualCircuits, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "associated_virtual_circuits")); ok {
		interfaces := virtualCircuits.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.AssociatedVirtualCircuits = tmp
	}
	return result
}

func (s *CoreIpSecConnectionResourceCrud) GetTunnels() error {
	if inputTunnelConfigs, ok := s.D.GetOk("tunnel_configuration"); ok {
		tmpList := inputTunnelConfigs.([]interface{})
		tunnelConfig := make([]PrivateIpSecConnectionTunnelResourceCrud, len(tmpList))
		for index, _ := range tmpList {
			tunnel := tmpList[index].(map[string]interface{})
			var tmp PrivateIpSecConnectionTunnelResourceCrud
			oracleIp := tunnel["oracle_tunnel_ip"].(string)
			drgId := tunnel["drg_route_table_id"].(string)
			tmp.OracleTunnelIp = &oracleIp
			tmp.AssociatedVirtualCircuits = make([]string, len(tunnel["associated_virtual_circuits"].([]interface{})))
			for k, v := range tunnel["associated_virtual_circuits"].([]interface{}) {
				tmp.AssociatedVirtualCircuits[k] = v.(string)
			}
			tmp.DrgRouteTableId = &drgId
			tunnelConfig[index] = tmp
		}
		s.ResTunnels = tunnelConfig
	} else {
		request := oci_core.ListIPSecConnectionTunnelsRequest{}

		tmp := s.D.Id()
		request.IpscId = &tmp

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

		response, err := s.Client.ListIPSecConnectionTunnels(context.Background(), request)
		if err != nil {
			return err
		}

		resTunnels := &response
		request.Page = resTunnels.OpcNextPage

		for request.Page != nil {
			listResponse, err := s.Client.ListIPSecConnectionTunnels(context.Background(), request)
			if err != nil {
				return err
			}

			resTunnels.Items = append(resTunnels.Items, listResponse.Items...)
			request.Page = listResponse.OpcNextPage
		}
		for _, value := range resTunnels.Items {
			var tmp PrivateIpSecConnectionTunnelResourceCrud
			tmp.OracleTunnelIp = value.VpnIp
			tmp.AssociatedVirtualCircuits = value.AssociatedVirtualCircuits
			err := s.GetDrgRouteTableId(&value, &tmp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *CoreIpSecConnectionResourceCrud) GetDrgRouteTableId(tunnel *oci_core.IpSecConnectionTunnel, t *PrivateIpSecConnectionTunnelResourceCrud) error {
	request := oci_core.ListDrgAttachmentsRequest{}

	request.NetworkId = tunnel.Id
	request.CompartmentId = tunnel.CompartmentId
	request.AttachmentType = oci_core.ListDrgAttachmentsAttachmentTypeIpsecTunnel

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDrgAttachments(context.Background(), request)
	if err != nil {
		return err
	}
	resAtt := response
	for request.Page != nil {
		listResponse, err := s.Client.ListDrgAttachments(context.Background(), request)
		if err != nil {
			return err
		}

		resAtt.Items = append(resAtt.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}
	if len(resAtt.Items) == 1 {
		t.DrgRouteTableId = resAtt.Items[0].DrgRouteTableId
	}
	return nil
}
