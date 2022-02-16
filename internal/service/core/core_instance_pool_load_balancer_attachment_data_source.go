// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreInstancePoolLoadBalancerAttachmentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreInstancePoolLoadBalancerAttachment,
		Schema: map[string]*schema.Schema{
			"instance_pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_pool_load_balancer_attachment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"backend_set_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnic_selection": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreInstancePoolLoadBalancerAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolLoadBalancerAttachmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreInstancePoolLoadBalancerAttachmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.GetInstancePoolLoadBalancerAttachmentResponse
}

func (s *CoreInstancePoolLoadBalancerAttachmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstancePoolLoadBalancerAttachmentDataSourceCrud) Get() error {
	request := oci_core.GetInstancePoolLoadBalancerAttachmentRequest{}

	if instancePoolId, ok := s.D.GetOkExists("instance_pool_id"); ok {
		tmp := instancePoolId.(string)
		request.InstancePoolId = &tmp
	}

	if instancePoolLoadBalancerAttachmentId, ok := s.D.GetOkExists("instance_pool_load_balancer_attachment_id"); ok {
		tmp := instancePoolLoadBalancerAttachmentId.(string)
		request.InstancePoolLoadBalancerAttachmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetInstancePoolLoadBalancerAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreInstancePoolLoadBalancerAttachmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BackendSetName != nil {
		s.D.Set("backend_set_name", *s.Res.BackendSetName)
	}

	if s.Res.LoadBalancerId != nil {
		s.D.Set("load_balancer_id", *s.Res.LoadBalancerId)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.VnicSelection != nil {
		s.D.Set("vnic_selection", *s.Res.VnicSelection)
	}

	return nil
}
