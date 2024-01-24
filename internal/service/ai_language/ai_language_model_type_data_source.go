// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiLanguageModelTypeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAiLanguageModelType,
		Schema: map[string]*schema.Schema{
			"model_type": {
				Type:     schema.TypeString,
				Required: true,
			},
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

func readSingularAiLanguageModelType(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

type AiLanguageModelTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_language.AIServiceLanguageClient
	Res    *oci_ai_language.GetModelTypeResponse
}

func (s *AiLanguageModelTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiLanguageModelTypeDataSourceCrud) Get() error {
	request := oci_ai_language.GetModelTypeRequest{}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		tmp := modelType.(string)
		request.ModelType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_language")

	response, err := s.Client.GetModelType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiLanguageModelTypeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiLanguageModelTypeDataSource-", AiLanguageModelTypeDataSource(), s.D))

	buf, err := json.Marshal(s.Res.Capabilities)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Capabilities%s\n", string(buf))
	s.D.Set("capabilities", string(buf))

	s.D.Set("versions", s.Res.Versions)

	return nil
}
