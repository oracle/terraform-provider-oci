// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterPublicApiEndpointDecommissionStatusDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularContainerengineClusterPublicApiEndpointDecommissionStatusWithContext,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_decommission_rollback_deadline": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularContainerengineClusterPublicApiEndpointDecommissionStatusWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ContainerengineClusterPublicApiEndpointDecommissionStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ContainerengineClusterPublicApiEndpointDecommissionStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetPublicApiEndpointDecommissionStatusResponse
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionStatusDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_containerengine.GetPublicApiEndpointDecommissionStatusRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetPublicApiEndpointDecommissionStatus(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineClusterPublicApiEndpointDecommissionStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClusterPublicApiEndpointDecommissionStatusDataSource-", ContainerengineClusterPublicApiEndpointDecommissionStatusDataSource(), s.D))

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeDecommissionRollbackDeadline != nil {
		s.D.Set("time_decommission_rollback_deadline", s.Res.TimeDecommissionRollbackDeadline.String())
	}

	return nil
}
