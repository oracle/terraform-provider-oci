// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"
)

func JmsSummarizeLibraryInventoryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsSummarizeLibraryInventory,
		Schema: map[string]*schema.Schema{
			"fleet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"statically_detected_library_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dynamically_detected_library_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"uncorrelated_package_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"high_severity_library_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"medium_severity_library_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"low_severity_library_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularJmsSummarizeLibraryInventory(d *schema.ResourceData, m interface{}) error {
	sync := &JmsSummarizeLibraryInventoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsSummarizeLibraryInventoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.SummarizeLibraryInventoryResponse
}

func (s *JmsSummarizeLibraryInventoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsSummarizeLibraryInventoryDataSourceCrud) Get() error {
	request := oci_jms.SummarizeLibraryInventoryRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.SummarizeLibraryInventory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsSummarizeLibraryInventoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsSummarizeLibraryInventoryDataSource-", JmsSummarizeLibraryInventoryDataSource(), s.D))

	if s.Res.StaticallyDetectedLibraryCount != nil {
		s.D.Set("statically_detected_library_count", *s.Res.StaticallyDetectedLibraryCount)
	}

	if s.Res.DynamicallyDetectedLibraryCount != nil {
		s.D.Set("dynamically_detected_library_count", *s.Res.DynamicallyDetectedLibraryCount)
	}

	if s.Res.UncorrelatedPackageCount != nil {
		s.D.Set("uncorrelated_package_count", *s.Res.UncorrelatedPackageCount)
	}

	if s.Res.HighSeverityLibraryCount != nil {
		s.D.Set("high_severity_library_count", *s.Res.HighSeverityLibraryCount)
	}

	if s.Res.MediumSeverityLibraryCount != nil {
		s.D.Set("medium_severity_library_count", *s.Res.MediumSeverityLibraryCount)
	}

	if s.Res.LowSeverityLibraryCount != nil {
		s.D.Set("low_severity_library_count", *s.Res.LowSeverityLibraryCount)
	}

	return nil
}
