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

func JmsUtilsAnalyzeApplicationsConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsUtilsAnalyzeApplicationsConfiguration,
		Read:     readJmsUtilsAnalyzeApplicationsConfiguration,
		Update:   updateJmsUtilsAnalyzeApplicationsConfiguration,
		Delete:   deleteJmsUtilsAnalyzeApplicationsConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional
			"bucket": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createJmsUtilsAnalyzeApplicationsConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsAnalyzeApplicationsConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsUtilsAnalyzeApplicationsConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsAnalyzeApplicationsConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

func updateJmsUtilsAnalyzeApplicationsConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsAnalyzeApplicationsConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsUtilsAnalyzeApplicationsConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsAnalyzeApplicationsConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsUtilsAnalyzeApplicationsConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms_utils.JmsUtilsClient
	Res                    *oci_jms_utils.AnalyzeApplicationsConfiguration
	DisableNotFoundRetries bool
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationResourceCrud) ID() string {
	if v, ok := s.D.GetOkExists("compartment_id"); ok {
		return v.(string)
	}
	return ""
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationResourceCrud) Create() error {
	request := oci_jms_utils.UpdateAnalyzeApplicationsConfigurationRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else if id := s.D.Id(); id != "" {
		// fall back to the resource ID (we use compartment OCID as ID)
		request.CompartmentId = &id
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_utils")

	response, err := s.Client.UpdateAnalyzeApplicationsConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyzeApplicationsConfiguration
	return nil
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationResourceCrud) Get() error {
	request := oci_jms_utils.GetAnalyzeApplicationsConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else if id := s.D.Id(); id != "" {
		// fall back to the resource ID (we use compartment OCID as ID)
		request.CompartmentId = &id
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_utils")

	response, err := s.Client.GetAnalyzeApplicationsConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyzeApplicationsConfiguration
	return nil
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationResourceCrud) Update() error {
	request := oci_jms_utils.UpdateAnalyzeApplicationsConfigurationRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else if id := s.D.Id(); id != "" {
		// fall back to the resource ID (we use compartment OCID as ID)
		request.CompartmentId = &id
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms_utils")

	response, err := s.Client.UpdateAnalyzeApplicationsConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyzeApplicationsConfiguration
	return nil
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationResourceCrud) Delete() error {
	// We dont know the original default values, so we pretend that the existing state is to be retained
	return nil
}

func (s *JmsUtilsAnalyzeApplicationsConfigurationResourceCrud) SetData() error {
	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	return nil
}
