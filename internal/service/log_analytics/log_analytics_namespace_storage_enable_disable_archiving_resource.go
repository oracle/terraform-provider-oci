// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceStorageEnableDisableArchivingResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsNamespaceStorageEnableDisableArchiving,
		Read:     readLogAnalyticsNamespaceStorageEnableDisableArchiving,
		Update:   updateLogAnalyticsNamespaceStorageEnableDisableArchiving,
		Delete:   deleteLogAnalyticsNamespaceStorageEnableDisableArchiving,
		Schema: map[string]*schema.Schema{
			// Required
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_archiving_tenant": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
			"message": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsNamespaceStorageEnableDisableArchiving(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.Res = &LogAnalyticsNamespaceStorageEnableDisableArchivingResponse{}

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceStorageEnableDisableArchiving(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateLogAnalyticsNamespaceStorageEnableDisableArchiving(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.Res = &LogAnalyticsNamespaceStorageEnableDisableArchivingResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsNamespaceStorageEnableDisableArchiving(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.Res = &LogAnalyticsNamespaceStorageEnableDisableArchivingResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsNamespaceStorageEnableDisableArchivingResponse struct {
	enableResponse  *oci_log_analytics.EnableArchivingResponse
	disableResponse *oci_log_analytics.DisableArchivingResponse
}

type LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *LogAnalyticsNamespaceStorageEnableDisableArchivingResponse
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceStorageEnableDisableArchivingResource-", LogAnalyticsNamespaceStorageEnableDisableArchivingResource(), s.D)
}

func (s *LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_archiving_tenant"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_log_analytics.EnableArchivingRequest{}

		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			request.NamespaceName = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

		response, err := s.Client.EnableArchiving(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_log_analytics.DisableArchivingRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.DisableArchiving(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud) Update() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_archiving_tenant"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_log_analytics.EnableArchivingRequest{}

		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			request.NamespaceName = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

		response, err := s.Client.EnableArchiving(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.enableResponse = &response
		return nil
	}

	request := oci_log_analytics.DisableArchivingRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.DisableArchiving(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res.disableResponse = &response
	return nil
}

func (s *LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud) Delete() error {
	return nil
}

func (s *LogAnalyticsNamespaceStorageEnableDisableArchivingResourceCrud) SetData() error {
	return nil
}
