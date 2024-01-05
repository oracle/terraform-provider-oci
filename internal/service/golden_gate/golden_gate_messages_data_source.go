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

func GoldenGateMessagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateMessages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_messages_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
		},
	}
}

func readGoldenGateMessages(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateMessagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateMessagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListMessagesResponse
}

func (s *GoldenGateMessagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateMessagesDataSourceCrud) Get() error {
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

func (s *GoldenGateMessagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateMessagesDataSource-", GoldenGateMessagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	message := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MessageSummaryToMap(item))
	}
	message["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateMessagesDataSource().Schema["deployment_messages_collection"].Elem.(*schema.Resource).Schema)
		message["items"] = items
	}

	resources = append(resources, message)
	if err := s.D.Set("deployment_messages_collection", resources); err != nil {
		return err
	}

	return nil
}

func MessageSummaryToMap(obj oci_golden_gate.MessageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeploymentMessage != nil {
		result["deployment_message"] = string(*obj.DeploymentMessage)
	}

	result["deployment_message_status"] = string(obj.DeploymentMessageStatus)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
