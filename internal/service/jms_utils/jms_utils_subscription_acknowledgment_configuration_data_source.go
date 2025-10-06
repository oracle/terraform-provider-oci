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

func JmsUtilsSubscriptionAcknowledgmentConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsUtilsSubscriptionAcknowledgmentConfigurationResource(), fieldMap, readSingularJmsUtilsSubscriptionAcknowledgmentConfiguration)
}

func readSingularJmsUtilsSubscriptionAcknowledgmentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsSubscriptionAcknowledgmentConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

type JmsUtilsSubscriptionAcknowledgmentConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_utils.JmsUtilsClient
	Res    *oci_jms_utils.GetSubscriptionAcknowledgmentConfigurationResponse
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationDataSourceCrud) Get() error {
	request := oci_jms_utils.GetSubscriptionAcknowledgmentConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_utils")

	response, err := s.Client.GetSubscriptionAcknowledgmentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsUtilsSubscriptionAcknowledgmentConfigurationDataSource-", JmsUtilsSubscriptionAcknowledgmentConfigurationDataSource(), s.D))

	if s.Res.AcknowledgedBy != nil {
		s.D.Set("acknowledged_by", *s.Res.AcknowledgedBy)
	}

	if s.Res.IsAcknowledged != nil {
		s.D.Set("is_acknowledged", *s.Res.IsAcknowledged)
	}

	if s.Res.TimeAcknowledged != nil {
		s.D.Set("time_acknowledged", s.Res.TimeAcknowledged.String())
	}

	return nil
}
