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

func JmsFleetAdvancedFeatureConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fleet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsFleetAdvancedFeatureConfigurationResource(), fieldMap, readSingularJmsFleetAdvancedFeatureConfiguration)
}

func readSingularJmsFleetAdvancedFeatureConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAdvancedFeatureConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetAdvancedFeatureConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetFleetAdvancedFeatureConfigurationResponse
}

func (s *JmsFleetAdvancedFeatureConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetAdvancedFeatureConfigurationDataSourceCrud) Get() error {
	request := oci_jms.GetFleetAdvancedFeatureConfigurationRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetFleetAdvancedFeatureConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetAdvancedFeatureConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetAdvancedFeatureConfigurationDataSource-", JmsFleetAdvancedFeatureConfigurationDataSource(), s.D))

	if s.Res.AdvancedUsageTracking != nil {
		s.D.Set("advanced_usage_tracking", []interface{}{AdvancedUsageTrackingToMap(s.Res.AdvancedUsageTracking)})
	} else {
		s.D.Set("advanced_usage_tracking", nil)
	}

	if s.Res.AnalyticBucketName != nil {
		s.D.Set("analytic_bucket_name", *s.Res.AnalyticBucketName)
	}

	if s.Res.AnalyticNamespace != nil {
		s.D.Set("analytic_namespace", *s.Res.AnalyticNamespace)
	}

	if s.Res.CryptoEventAnalysis != nil {
		s.D.Set("crypto_event_analysis", []interface{}{CryptoEventAnalysisToMap(s.Res.CryptoEventAnalysis)})
	} else {
		s.D.Set("crypto_event_analysis", nil)
	}

	if s.Res.JavaMigrationAnalysis != nil {
		s.D.Set("java_migration_analysis", []interface{}{JavaMigrationAnalysisToMap(s.Res.JavaMigrationAnalysis)})
	} else {
		s.D.Set("java_migration_analysis", nil)
	}

	if s.Res.JfrRecording != nil {
		s.D.Set("jfr_recording", []interface{}{JfrRecordingToMap(s.Res.JfrRecording)})
	} else {
		s.D.Set("jfr_recording", nil)
	}

	if s.Res.Lcm != nil {
		s.D.Set("lcm", []interface{}{LcmToMap(s.Res.Lcm)})
	} else {
		s.D.Set("lcm", nil)
	}

	if s.Res.PerformanceTuningAnalysis != nil {
		s.D.Set("performance_tuning_analysis", []interface{}{PerformanceTuningAnalysisToMap(s.Res.PerformanceTuningAnalysis)})
	} else {
		s.D.Set("performance_tuning_analysis", nil)
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}
