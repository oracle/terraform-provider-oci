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

func JmsUtilsSubscriptionAcknowledgmentConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsUtilsSubscriptionAcknowledgmentConfiguration,
		Read:     readJmsUtilsSubscriptionAcknowledgmentConfiguration,
		Update:   updateJmsUtilsSubscriptionAcknowledgmentConfiguration,
		Delete:   deleteJmsUtilsSubscriptionAcknowledgmentConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_acknowledged": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Computed
			"acknowledged_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_acknowledged": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createJmsUtilsSubscriptionAcknowledgmentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsUtilsSubscriptionAcknowledgmentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

func updateJmsUtilsSubscriptionAcknowledgmentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsUtilsSubscriptionAcknowledgmentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms_utils.JmsUtilsClient
	Res                    *oci_jms_utils.SubscriptionAcknowledgmentConfiguration
	DisableNotFoundRetries bool
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud) ID() string {
	if v, ok := s.D.GetOkExists("compartment_id"); ok {
		return v.(string)
	}
	return ""
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud) Create() error {
	request := oci_jms_utils.UpdateSubscriptionAcknowledgmentConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else if id := s.D.Id(); id != "" {
		// fall back to the resource ID (we use compartment OCID as ID)
		request.CompartmentId = &id
	}

	if isAcknowledged, ok := s.D.GetOkExists("is_acknowledged"); ok {
		tmp := isAcknowledged.(bool)
		request.IsAcknowledged = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_utils")

	response, err := s.Client.UpdateSubscriptionAcknowledgmentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SubscriptionAcknowledgmentConfiguration
	return nil
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud) Get() error {
	request := oci_jms_utils.GetSubscriptionAcknowledgmentConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else if id := s.D.Id(); id != "" {
		// fall back to the resource ID (we use compartment OCID as ID)
		request.CompartmentId = &id
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_utils")

	response, err := s.Client.GetSubscriptionAcknowledgmentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SubscriptionAcknowledgmentConfiguration
	return nil
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud) Update() error {
	request := oci_jms_utils.UpdateSubscriptionAcknowledgmentConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else if id := s.D.Id(); id != "" {
		// fall back to the resource ID (we use compartment OCID as ID)
		request.CompartmentId = &id
	}

	if isAcknowledged, ok := s.D.GetOkExists("is_acknowledged"); ok {
		tmp := isAcknowledged.(bool)
		request.IsAcknowledged = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_utils")

	response, err := s.Client.UpdateSubscriptionAcknowledgmentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SubscriptionAcknowledgmentConfiguration
	return nil
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud) Delete() error {
	// We dont know the original default values, so we pretend that the existing state is to be retained
	return nil
}

func (s *JmsUtilsSubscriptionAcknowledgmentConfigurationResourceCrud) SetData() error {
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
