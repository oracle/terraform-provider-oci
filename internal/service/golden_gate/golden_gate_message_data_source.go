// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateMessageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularGoldenGateMessage,
		Schema: map[string]*schema.Schema{
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"deployment_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_message_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_golden_gate_message", "oci_golden_gate_messages"),
	}
}

func readSingularGoldenGateMessage(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateMessageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateMessageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListMessagesResponse
}

func (s *GoldenGateMessageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateMessageDataSourceCrud) Get() error {
	request := oci_golden_gate.ListMessagesRequest{}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListMessages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateMessageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateMessageDataSource-", GoldenGateMessageDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MessageSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
