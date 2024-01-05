// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_waas "github.com/oracle/oci-go-sdk/v65/waas"
)

func WaasPurgeCacheResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createWaasPurgeCache,
		Read:   readWaasPurgeCache,
		Delete: deleteWaasPurgeCache,
		Schema: map[string]*schema.Schema{
			// Required
			"waas_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"resources": {
				Type:     schema.TypeList,
				Optional: true,
				//@Codegen: the field is not computed as it cannot be set by the service
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
		},
	}
}

func createWaasPurgeCache(d *schema.ResourceData, m interface{}) error {
	sync := &WaasPurgeCacheResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.CreateResource(d, sync)
}

func readWaasPurgeCache(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteWaasPurgeCache(d *schema.ResourceData, m interface{}) error {
	return nil
}

type WaasPurgeCacheResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waas.WaasClient
	Res                    *oci_waas.PurgeCache
	DisableNotFoundRetries bool
}

func (s *WaasPurgeCacheResourceCrud) ID() string {
	return tfresource.Timestamp()
}

func (s *WaasPurgeCacheResourceCrud) Create() error {
	request := oci_waas.PurgeCacheRequest{}

	if resources, ok := s.D.GetOkExists("resources"); ok {
		interfaces := resources.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resources") {
			request.Resources = tmp
		}
	}

	if waasPolicyId, ok := s.D.GetOkExists("waas_policy_id"); ok {
		tmp := waasPolicyId.(string)
		request.WaasPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.PurgeCache(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, delWorkRequestErr := waasPolicyWaitForWorkRequest(workId, "waas",
		oci_waas.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *WaasPurgeCacheResourceCrud) SetData() error {
	s.D.SetId(tfresource.GenerateDataSourceID())
	return nil
}
