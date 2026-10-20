// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpGenerateVmwareBinaryDownloadInfoDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpGenerateVmwareBinaryDownloadInfo,
		Schema: map[string]*schema.Schema{
			"sddc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vmware_binary_file_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readOcvpGenerateVmwareBinaryDownloadInfo(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpGenerateVmwareBinaryDownloadInfoDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()

	return tfresource.ReadResource(sync)
}

type OcvpGenerateVmwareBinaryDownloadInfoDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.GenerateVmwareBinaryDownloadInfoResponse
}

func (s *OcvpGenerateVmwareBinaryDownloadInfoDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpGenerateVmwareBinaryDownloadInfoDataSourceCrud) Get() error {
	request := oci_ocvp.GenerateVmwareBinaryDownloadInfoRequest{}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	if vmwareBinaryFileName, ok := s.D.GetOkExists("vmware_binary_file_name"); ok {
		tmp := vmwareBinaryFileName.(string)
		request.GenerateVmwareBinaryDownloadInfoDetails = oci_ocvp.GenerateVmwareBinaryDownloadInfoDetails{
			VmwareBinaryFileName: &tmp,
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.GenerateVmwareBinaryDownloadInfo(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpGenerateVmwareBinaryDownloadInfoDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpGenerateVmwareBinaryDownloadInfoDataSource-", OcvpGenerateVmwareBinaryDownloadInfoDataSource(), s.D))

	if s.Res.FileName != nil {
		s.D.Set("file_name", *s.Res.FileName)
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.String())
	}

	if s.Res.Url != nil {
		s.D.Set("url", *s.Res.Url)
	}

	return nil
}
