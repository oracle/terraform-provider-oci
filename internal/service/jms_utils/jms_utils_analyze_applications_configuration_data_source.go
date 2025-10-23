// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_utils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_utils "github.com/oracle/oci-go-sdk/v65/jmsutils"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsUtilsAnalyzeApplicationsConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsUtilsAnalyzeApplicationsConfigurationResource(), fieldMap, readSingularJmsUtilsAnalyzeApplicationsConfiguration)
}

func readSingularJmsUtilsAnalyzeApplicationsConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsAnalyzeApplicationsConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

type JmsUtilsAnalyzeApplicationsConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_utils.JmsUtilsClient
	Res    *oci_jms_utils.GetAnalyzeApplicationsConfigurationResponse
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationDataSourceCrud) Get() error {
	request := oci_jms_utils.GetAnalyzeApplicationsConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_utils")

	response, err := s.Client.GetAnalyzeApplicationsConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsUtilsAnalyzeApplicationsConfigurationDataSource-", JmsUtilsAnalyzeApplicationsConfigurationDataSource(), s.D))

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	return nil
}
