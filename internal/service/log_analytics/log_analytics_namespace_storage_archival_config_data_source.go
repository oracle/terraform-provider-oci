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

func LogAnalyticsNamespaceStorageArchivalConfigDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LogAnalyticsNamespaceStorageArchivalConfigResource(), fieldMap, readSingularLogAnalyticsNamespaceStorageArchivalConfig)
}

func readSingularLogAnalyticsNamespaceStorageArchivalConfig(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageArchivalConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceStorageArchivalConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetStorageResponse
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigDataSourceCrud) Get() error {
	request := oci_log_analytics.GetStorageRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetStorage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceStorageArchivalConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceStorageArchivalConfigDataSource-", LogAnalyticsNamespaceStorageArchivalConfigDataSource(), s.D))

	if s.Res.ArchivingConfiguration != nil {
		s.D.Set("archiving_configuration", []interface{}{ArchivingConfigurationToMap(s.Res.ArchivingConfiguration)})
	} else {
		s.D.Set("archiving_configuration", nil)
	}

	if s.Res.IsArchivingEnabled != nil {
		s.D.Set("is_archiving_enabled", *s.Res.IsArchivingEnabled)
	}

	return nil
}
