// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourcemanagerPrivateEndpointReachableIpDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularResourcemanagerPrivateEndpointReachableIp,
		Schema: map[string]*schema.Schema{
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularResourcemanagerPrivateEndpointReachableIp(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerPrivateEndpointReachableIpDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.ReadResource(sync)
}

type ResourcemanagerPrivateEndpointReachableIpDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resourcemanager.ResourceManagerClient
	Res    *oci_resourcemanager.GetReachableIpResponse
}

func (s *ResourcemanagerPrivateEndpointReachableIpDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourcemanagerPrivateEndpointReachableIpDataSourceCrud) Get() error {
	request := oci_resourcemanager.GetReachableIpRequest{}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
		tmp := privateIp.(string)
		request.PrivateIp = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resourcemanager")

	response, err := s.Client.GetReachableIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourcemanagerPrivateEndpointReachableIpDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ResourcemanagerPrivateEndpointReachableIpDataSource-", ResourcemanagerPrivateEndpointReachableIpDataSource(), s.D))

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	return nil
}
