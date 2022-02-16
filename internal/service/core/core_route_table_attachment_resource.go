// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreRouteTableAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRouteTableAttachment,
		Delete:   deleteRouteTableAttachment,
		Read:     readRouteTableAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createRouteTableAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.CreateResource(d, sync)
}

func deleteRouteTableAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

func readRouteTableAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &RouteTableAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.ReadResource(sync)
}

type RouteTableAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Subnet
	DisableNotFoundRetries bool
}

func getRouteTableAttachmentId(subnetId string, routeTableId string) string {
	return subnetId + "/" + routeTableId
}

func parseRouteTableAttachmentId(id string) (subnetId string, routTableId string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) < 2 {
		err = fmt.Errorf("illegal id %s encountered", id)
	}
	subnetId, routTableId = parts[0], parts[1]

	return
}

func (s *RouteTableAttachmentResourceCrud) ID() string {
	return getRouteTableAttachmentId(*s.Res.Id, *s.Res.RouteTableId)
}

func (s *RouteTableAttachmentResourceCrud) Create() error {
	request := oci_core.UpdateSubnetRequest{}

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	response, err := s.Client.UpdateSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subnet
	return nil
}

func (s *RouteTableAttachmentResourceCrud) Get() error {

	subnetId, _, err := parseRouteTableAttachmentId(s.D.Id())
	if err != nil {
		return err
	}
	request := oci_core.GetSubnetRequest{}
	request.SubnetId = &subnetId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")
	response, err := s.Client.GetSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subnet
	return nil
}

func (s *RouteTableAttachmentResourceCrud) Delete() error {

	var subnetIdStr = s.D.Get("subnet_id").(string)
	var retryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	subnetRequest := oci_core.GetSubnetRequest{}
	subnetRequest.SubnetId = &subnetIdStr
	subnetRequest.RequestMetadata.RetryPolicy = retryPolicy
	subnetResponse, err := s.Client.GetSubnet(context.Background(), subnetRequest)
	if err != nil {
		return err
	}

	vcnRequest := oci_core.GetVcnRequest{}
	vcnRequest.VcnId = subnetResponse.VcnId
	vcnRequest.RequestMetadata.RetryPolicy = retryPolicy
	vcnResponse, err := s.Client.GetVcn(context.Background(), vcnRequest)
	if err != nil {
		return err
	}

	updateSubnetRequest := oci_core.UpdateSubnetRequest{}
	updateSubnetRequest.RouteTableId = vcnResponse.DefaultRouteTableId
	updateSubnetRequest.SubnetId = &subnetIdStr
	updateSubnetRequest.RequestMetadata.RetryPolicy = retryPolicy
	updateSubnetResponse, err := s.Client.UpdateSubnet(context.Background(), updateSubnetRequest)
	if err != nil {
		return err
	}

	s.Res = &updateSubnetResponse.Subnet
	return nil
}

func (s *RouteTableAttachmentResourceCrud) SetData() error {

	s.D.Set("subnet_id", *s.Res.Id)

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	return nil
}
