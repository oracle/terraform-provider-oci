// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeGenerateOnPremConnectorConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeGenerateOnPremConnectorConfiguration,
		Read:     readDataSafeGenerateOnPremConnectorConfiguration,
		Delete:   deleteDataSafeGenerateOnPremConnectorConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"on_prem_connector_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDataSafeGenerateOnPremConnectorConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeGenerateOnPremConnectorConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeGenerateOnPremConnectorConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeGenerateOnPremConnectorConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeGenerateOnPremConnectorConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.GenerateOnPremConnectorConfigurationResponse
	DisableNotFoundRetries bool
}

func (s *DataSafeGenerateOnPremConnectorConfigurationResourceCrud) ID() string {
	return *s.Res.OpcRequestId
}

func (s *DataSafeGenerateOnPremConnectorConfigurationResourceCrud) Create() error {
	request := oci_data_safe.GenerateOnPremConnectorConfigurationRequest{}

	if onPremConnectorId, ok := s.D.GetOkExists("on_prem_connector_id"); ok {
		tmp := onPremConnectorId.(string)
		request.OnPremConnectorId = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.GenerateOnPremConnectorConfigurationDetails = oci_data_safe.GenerateOnPremConnectorConfigurationDetails{
			Password: &tmp,
		}
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GenerateOnPremConnectorConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeGenerateOnPremConnectorConfigurationResourceCrud) SetData() error {
	return nil
}
