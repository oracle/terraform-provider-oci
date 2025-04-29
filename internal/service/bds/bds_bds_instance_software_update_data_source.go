// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceSoftwareUpdateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularBdsBdsInstanceSoftwareUpdate,
		Schema: map[string]*schema.Schema{
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_update_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"software_update_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"software_update_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_due": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_released": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularBdsBdsInstanceSoftwareUpdate(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceSoftwareUpdateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceSoftwareUpdateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetSoftwareUpdateResponse
}

func (s *BdsBdsInstanceSoftwareUpdateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceSoftwareUpdateDataSourceCrud) Get() error {
	request := oci_bds.GetSoftwareUpdateRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if softwareUpdateKey, ok := s.D.GetOkExists("software_update_key"); ok {
		tmp := softwareUpdateKey.(string)
		request.SoftwareUpdateKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetSoftwareUpdate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceSoftwareUpdateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceSoftwareUpdateDataSource-", BdsBdsInstanceSoftwareUpdateDataSource(), s.D))
	switch v := (s.Res.SoftwareUpdate).(type) {
	case oci_bds.BdsSoftwareUpdate:
		s.D.Set("software_update_type", "BDS")

		if v.TimeDue != nil {
			s.D.Set("time_due", v.TimeDue.Format(time.RFC3339Nano))
		}

		if v.SoftwareUpdateKey != nil {
			s.D.Set("software_update_key", *v.SoftwareUpdateKey)
		}

		if v.SoftwareUpdateVersion != nil {
			s.D.Set("software_update_version", *v.SoftwareUpdateVersion)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeReleased != nil {
			s.D.Set("time_released", v.TimeReleased.String())
		}
	default:
		log.Printf("[WARN] Received 'software_update_type' of unknown type %v", s.Res.SoftwareUpdate)
		return nil
	}

	return nil
}
