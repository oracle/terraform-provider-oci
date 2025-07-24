// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_document

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_document "github.com/oracle/oci-go-sdk/v65/aidocument"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiDocumentModelTypeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAiDocumentModelType,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_sub_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"capabilities": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularAiDocumentModelType(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentModelTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.ReadResource(sync)
}

type AiDocumentModelTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_document.AIServiceDocumentClient
	Res    *oci_ai_document.GetModelTypeResponse
}

func (s *AiDocumentModelTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiDocumentModelTypeDataSourceCrud) Get() error {
	request := oci_ai_document.GetModelTypeRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if modelSubType, ok := s.D.GetOkExists("model_sub_type"); ok {
		tmp := modelSubType.(string)
		request.ModelSubType = &tmp
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		tmp := modelType.(string)
		request.ModelType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_document")

	response, err := s.Client.GetModelType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiDocumentModelTypeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiDocumentModelTypeDataSource-", AiDocumentModelTypeDataSource(), s.D))

	buf, err := json.Marshal(s.Res.Capabilities)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Capabilities%s\n", string(buf))
	s.D.Set("capabilities", string(buf))

	s.D.Set("versions", s.Res.Versions)

	return nil
}
