// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"net"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreIpSecConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreIpSecConnections,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpe_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreIpSecConnectionResource()),
			},
		},
	}
}

func readCoreIpSecConnections(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpSecConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpSecConnectionsDataSourceCrud struct {
	D          *schema.ResourceData
	Client     *oci_core.VirtualNetworkClient
	Res        *oci_core.ListIPSecConnectionsResponse
	ResTunnels map[string][]PrivateIpSecConnectionTunnelResourceCrud
}

func (s *CoreIpSecConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpSecConnectionsDataSourceCrud) Get() error {
	request := oci_core.ListIPSecConnectionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if cpeId, ok := s.D.GetOkExists("cpe_id"); ok {
		tmp := cpeId.(string)
		request.CpeId = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListIPSecConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIPSecConnections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	s.ResTunnels = make(map[string][]PrivateIpSecConnectionTunnelResourceCrud, len(s.Res.Items))
	for _, value := range s.Res.Items {
		// only retrieve tunnels for tunnel config info if ipsec over fastconnect
		if value.TransportType == oci_core.IpSecConnectionTransportTypeFastconnect {
			err = s.GetTunnels(value.Id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *CoreIpSecConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpSecConnectionsDataSource-", CoreIpSecConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ipSecConnection := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CpeId != nil {
			ipSecConnection["cpe_id"] = *r.CpeId
		}

		if r.CpeLocalIdentifier != nil {
			ipSecConnection["cpe_local_identifier"] = *r.CpeLocalIdentifier
		}

		ipSecConnection["cpe_local_identifier_type"] = r.CpeLocalIdentifierType

		if r.DefinedTags != nil {
			ipSecConnection["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			ipSecConnection["display_name"] = *r.DisplayName
		}

		if r.DrgId != nil {
			ipSecConnection["drg_id"] = *r.DrgId
		}

		ipSecConnection["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			ipSecConnection["id"] = *r.Id
		}

		ipSecConnection["state"] = r.LifecycleState

		ipSecConnection["static_routes"] = r.StaticRoutes

		if r.TimeCreated != nil {
			ipSecConnection["time_created"] = r.TimeCreated.String()
		}

		ipSecConnection["transport_type"] = r.TransportType

		// set tunnel_configurations if ipsec over fast connect tunnel
		if r.TransportType == oci_core.IpSecConnectionTransportTypeFastconnect {
			tunnels := s.ResTunnels[*r.Id]
			if tunnels != nil && len(tunnels) > 0 {
				tmpList := make([]interface{}, len(tunnels))
				for key, value := range tunnels {
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
					tmpList[key] = t
				}
				ipSecConnection["tunnel_configuration"] = tmpList
			}
		}
		resources = append(resources, ipSecConnection)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreIpSecConnectionsDataSource().Schema["connections"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("connections", resources); err != nil {
		return err
	}

	return nil
}

func (s *CoreIpSecConnectionsDataSourceCrud) GetTunnels(Id *string) error {
	request := oci_core.ListIPSecConnectionTunnelsRequest{}

	request.IpscId = Id

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
	sort.Slice(resTunnels.Items, func(i, j int) bool {
		return bytes.Compare(net.ParseIP(*resTunnels.Items[i].VpnIp), net.ParseIP(*resTunnels.Items[j].VpnIp)) > 0
	})
	tunnelConfig := make([]PrivateIpSecConnectionTunnelResourceCrud, len(resTunnels.Items))
	for key, value := range resTunnels.Items {
		var tmp PrivateIpSecConnectionTunnelResourceCrud
		tmp.AssociatedVirtualCircuits = value.AssociatedVirtualCircuits
		tmp.OracleTunnelIp = value.VpnIp
		err := s.GetDrgRouteTableId(value, &tmp)
		if err != nil {
			return err
		}
		tunnelConfig[key] = tmp
	}
	s.ResTunnels[*Id] = tunnelConfig
	return nil
}

func (s *CoreIpSecConnectionsDataSourceCrud) GetDrgRouteTableId(tunnel oci_core.IpSecConnectionTunnel, t *PrivateIpSecConnectionTunnelResourceCrud) error {
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
