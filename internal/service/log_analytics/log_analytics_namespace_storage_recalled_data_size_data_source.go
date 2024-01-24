// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceStorageRecalledDataSizeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsNamespaceStorageRecalledDataSize,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_data_ended": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time_data_started": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Computed
			"not_recalled_data_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recalled_data_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsNamespaceStorageRecalledDataSize(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageRecalledDataSizeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceStorageRecalledDataSizeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetRecalledDataSizeResponse
}

func (s *LogAnalyticsNamespaceStorageRecalledDataSizeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceStorageRecalledDataSizeDataSourceCrud) Get() error {
	request := oci_log_analytics.GetRecalledDataSizeRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if timeDataEnded, ok := s.D.GetOkExists("time_data_ended"); ok {
		tmp, err := time.Parse(time.RFC3339, timeDataEnded.(string))
		if err != nil {
			return err
		}
		request.TimeDataEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeDataStarted, ok := s.D.GetOkExists("time_data_started"); ok {
		tmp, err := time.Parse(time.RFC3339, timeDataStarted.(string))
		if err != nil {
			return err
		}
		request.TimeDataStarted = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetRecalledDataSize(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceStorageRecalledDataSizeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceStorageRecalledDataSizeDataSource-", LogAnalyticsNamespaceStorageRecalledDataSizeDataSource(), s.D))

	if s.Res.NotRecalledDataInBytes != nil {
		s.D.Set("not_recalled_data_in_bytes", strconv.FormatInt(*s.Res.NotRecalledDataInBytes, 10))
	}

	if s.Res.RecalledDataInBytes != nil {
		s.D.Set("recalled_data_in_bytes", strconv.FormatInt(*s.Res.RecalledDataInBytes, 10))
	}

	if s.Res.TimeDataEnded != nil {
		s.D.Set("time_data_ended", s.Res.TimeDataEnded.String())
	}

	if s.Res.TimeDataStarted != nil {
		s.D.Set("time_data_started", s.Res.TimeDataStarted.String())
	}

	return nil
}
