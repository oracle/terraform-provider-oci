// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeSdmMaskingPolicyDifferenceDifferenceColumn,
		Schema: map[string]*schema.Schema{
			"difference_column_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sdm_masking_policy_difference_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"column_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"difference_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"masking_columnkey": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"planned_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schema_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sensitive_columnkey": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_synced": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataSafeSdmMaskingPolicyDifferenceDifferenceColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetDifferenceColumnResponse
}

func (s *DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSourceCrud) Get() error {
	request1 := oci_data_safe.ListDifferenceColumnsRequest{}

	if sdmMaskingPolicyDifferenceId, ok := s.D.GetOkExists("sdm_masking_policy_difference_id"); ok {
		tmp := sdmMaskingPolicyDifferenceId.(string)
		request1.SdmMaskingPolicyDifferenceId = &tmp
	}

	request1.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	var response oci_data_safe.ListDifferenceColumnsResponse
	response, err := s.Client.ListDifferenceColumns(context.Background(), request1)
	if err != nil {
		return err
	}

	var differenceColumnKey string
	for _, item := range response.Items {
		differenceColumnKey = string(*item.Key)
		break
	}

	request2 := oci_data_safe.GetDifferenceColumnRequest{}
	request2.DifferenceColumnKey = &differenceColumnKey

	if sdmMaskingPolicyDifferenceId, ok := s.D.GetOkExists("sdm_masking_policy_difference_id"); ok {
		tmp := sdmMaskingPolicyDifferenceId.(string)
		request2.SdmMaskingPolicyDifferenceId = &tmp
	}

	request2.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response2, err2 := s.Client.GetDifferenceColumn(context.Background(), request2)
	if err2 != nil {
		return err
	}

	s.Res = &response2
	return nil
}

func (s *DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSource-", DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSource(), s.D))

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	s.D.Set("difference_type", s.Res.DifferenceType)

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.MaskingColumnkey != nil {
		s.D.Set("masking_columnkey", *s.Res.MaskingColumnkey)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("planned_action", s.Res.PlannedAction)

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.SensitiveColumnkey != nil {
		s.D.Set("sensitive_columnkey", *s.Res.SensitiveColumnkey)
	}

	if s.Res.SensitiveTypeId != nil {
		s.D.Set("sensitive_type_id", *s.Res.SensitiveTypeId)
	}

	s.D.Set("sync_status", s.Res.SyncStatus)

	if s.Res.TimeLastSynced != nil {
		s.D.Set("time_last_synced", s.Res.TimeLastSynced.String())
	}

	return nil
}
