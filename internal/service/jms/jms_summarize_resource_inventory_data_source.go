// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_jms "github.com/oracle/oci-go-sdk/v58/jms"
)

func JmsSummarizeResourceInventoryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsSummarizeResourceInventory,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
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
			"active_fleet_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"application_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"installation_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"jre_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"managed_instance_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularJmsSummarizeResourceInventory(d *schema.ResourceData, m interface{}) error {
	sync := &JmsSummarizeResourceInventoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsSummarizeResourceInventoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.SummarizeResourceInventoryResponse
}

func (s *JmsSummarizeResourceInventoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsSummarizeResourceInventoryDataSourceCrud) Get() error {
	request := oci_jms.SummarizeResourceInventoryRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	response, err := s.Client.SummarizeResourceInventory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsSummarizeResourceInventoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsSummarizeResourceInventoryDataSource-", JmsSummarizeResourceInventoryDataSource(), s.D))

	if s.Res.ActiveFleetCount != nil {
		s.D.Set("active_fleet_count", *s.Res.ActiveFleetCount)
	}

	if s.Res.ApplicationCount != nil {
		s.D.Set("application_count", *s.Res.ApplicationCount)
	}

	if s.Res.InstallationCount != nil {
		s.D.Set("installation_count", *s.Res.InstallationCount)
	}

	if s.Res.JreCount != nil {
		s.D.Set("jre_count", *s.Res.JreCount)
	}

	if s.Res.ManagedInstanceCount != nil {
		s.D.Set("managed_instance_count", *s.Res.ManagedInstanceCount)
	}

	return nil
}
