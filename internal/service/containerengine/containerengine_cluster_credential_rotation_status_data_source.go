// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterCredentialRotationStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularContainerengineClusterCredentialRotationStatus,
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
			"status_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_auto_completion_scheduled": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularContainerengineClusterCredentialRotationStatus(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterCredentialRotationStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterCredentialRotationStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetCredentialRotationStatusResponse
}

func (s *ContainerengineClusterCredentialRotationStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterCredentialRotationStatusDataSourceCrud) Get() error {
	request := oci_containerengine.GetCredentialRotationStatusRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.GetCredentialRotationStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineClusterCredentialRotationStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClusterCredentialRotationStatusDataSource-", ContainerengineClusterCredentialRotationStatusDataSource(), s.D))

	s.D.Set("status", s.Res.Status)

	s.D.Set("status_details", s.Res.StatusDetails)

	if s.Res.TimeAutoCompletionScheduled != nil {
		s.D.Set("time_auto_completion_scheduled", s.Res.TimeAutoCompletionScheduled.String())
	}

	return nil
}
