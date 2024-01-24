// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetDrsFileDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetDrsFile,
		Schema: map[string]*schema.Schema{
			"drs_file_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"bucket": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"drs_file_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsFleetDrsFile(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetDrsFileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetDrsFileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetDrsFileResponse
}

func (s *JmsFleetDrsFileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetDrsFileDataSourceCrud) Get() error {
	request := oci_jms.GetDrsFileRequest{}

	if drsFileKey, ok := s.D.GetOkExists("drs_file_key"); ok {
		tmp := drsFileKey.(string)
		request.DrsFileKey = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetDrsFile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetDrsFileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetDrsFileDataSource-", JmsFleetDrsFileDataSource(), s.D))

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	s.D.Set("checksum_type", s.Res.ChecksumType)

	if s.Res.ChecksumValue != nil {
		s.D.Set("checksum_value", *s.Res.ChecksumValue)
	}

	if s.Res.DrsFileName != nil {
		s.D.Set("drs_file_name", *s.Res.DrsFileName)
	}

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	return nil
}
