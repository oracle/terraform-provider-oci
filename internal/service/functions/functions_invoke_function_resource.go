// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"
)

func FunctionsInvokeFunctionResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFunctionsInvokeFunction,
		Read:     readFunctionsInvokeFunction,
		Delete:   deleteFunctionsInvokeFunction,
		Schema: map[string]*schema.Schema{
			// Required
			"function_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"invoke_function_body": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"invoke_function_body_base64_encoded", "input_body_source_path"},
			},
			"invoke_function_body_base64_encoded": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"invoke_function_body", "input_body_source_path"},
			},
			"fn_intent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fn_invoke_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"input_body_source_path": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				StateFunc:     tfresource.GetSourceFileState,
				ConflictsWith: []string{"invoke_function_body", "invoke_function_body_base64_encoded"},
				ValidateFunc:  validateFunctionSourceValue,
			},
			"base64_encode_content": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			// Computed
			"invoke_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFunctionsInvokeFunction(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsInvokeFunctionResourceCrud{}
	sync.D = d

	endPoint, ok := d.GetOkExists("invoke_endpoint")
	if !ok {
		var err error
		endPoint, err = sync.getInvokeEndPoint(m)
		if err != nil {
			return err
		}
		// cache endpoint as it won't change for a function
		d.Set("invoke_endpoint", endPoint)
	}

	sync.Client, _ = m.(*client.OracleClients).FunctionsInvokeClientWithEndpoint(endPoint.(string))
	return tfresource.CreateResource(d, sync)
}

func (s *FunctionsInvokeFunctionResourceCrud) getInvokeEndPoint(m interface{}) (string, error) {
	functionsManagementClient := m.(*client.OracleClients).FunctionsManagementClient()
	functionGetRequest := oci_functions.GetFunctionRequest{}

	if functionId, ok := s.D.GetOkExists("function_id"); ok {
		tmp := functionId.(string)
		functionGetRequest.FunctionId = &tmp
	}

	functionGetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	functionGetResponse, err := functionsManagementClient.GetFunction(context.Background(), functionGetRequest)
	if err != nil {
		return "", err
	}

	if functionGetResponse.InvokeEndpoint == nil {
		return "", errors.New("No valid function invocation endpoint.")
	}

	endPoint := *functionGetResponse.InvokeEndpoint
	return endPoint, nil
}

func readFunctionsInvokeFunction(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFunctionsInvokeFunction(d *schema.ResourceData, m interface{}) error {
	return nil
}

type FunctionsInvokeFunctionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_functions.FunctionsInvokeClient
	Res                    *[]byte
	DisableNotFoundRetries bool
}

func (s *FunctionsInvokeFunctionResourceCrud) ID() string {
	return tfresource.Timestamp()
}

func (s *FunctionsInvokeFunctionResourceCrud) Create() error {
	request := oci_functions.InvokeFunctionRequest{}

	if invokeFunctionBody, ok := s.D.GetOkExists("invoke_function_body"); ok {
		tmp := []byte(invokeFunctionBody.(string))
		request.InvokeFunctionBody = ioutil.NopCloser(bytes.NewReader(tmp))
	}

	if invokeFunctionBodyBase64Encoded, ok := s.D.GetOkExists("invoke_function_body_base64_encoded"); ok {
		tmp, err := base64.StdEncoding.DecodeString(invokeFunctionBodyBase64Encoded.(string))
		if err != nil {
			return err
		}
		request.InvokeFunctionBody = ioutil.NopCloser(bytes.NewReader(tmp))
	}

	// Either of 'invoke_function_body' or 'input_body_source_path' or 'invoke_function_body_base64_encoded' will be available
	if sourcePath, ok := s.D.GetOkExists("input_body_source_path"); ok {
		sourceFile, err := os.Open(sourcePath.(string))
		if err != nil {
			return err
		}
		defer tfresource.SafeClose(sourceFile, &err)
		request.InvokeFunctionBody = ioutil.NopCloser(sourceFile)
	}

	if fnIntent, ok := s.D.GetOkExists("fn_intent"); ok {
		request.FnIntent = oci_functions.InvokeFunctionFnIntentEnum(fnIntent.(string))
	}

	if fnInvokeType, ok := s.D.GetOkExists("fn_invoke_type"); ok {
		request.FnInvokeType = oci_functions.InvokeFunctionFnInvokeTypeEnum(fnInvokeType.(string))
	}

	if functionId, ok := s.D.GetOkExists("function_id"); ok {
		tmp := functionId.(string)
		request.FunctionId = &tmp
	}

	if isDryRun, ok := s.D.GetOkExists("is_dry_run"); ok {
		tmp := isDryRun.(bool)
		request.IsDryRun = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")
	if endPoint, ok := s.D.GetOkExists("invoke_endpoint"); !ok {
		s.Client.Host = endPoint.(string)
	}

	response, err := s.Client.InvokeFunction(context.Background(), request)
	if err != nil {
		return err
	}

	if response.Content != nil {
		defer response.Content.Close()
		if contentBytes, err := ioutil.ReadAll(response.Content); err == nil {
			s.Res = &contentBytes
		} else {
			return err
		}
	}
	return nil
}

func (s *FunctionsInvokeFunctionResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())

	base64EncodeContent := false
	if tmp, ok := s.D.GetOkExists("base64_encode_content"); ok {
		base64EncodeContent = tmp.(bool)
	}

	if base64EncodeContent {
		// This use case is for v0.12, where content should be base64 encoded to avoid
		// being normalized before setting in state.
		s.D.Set("content", base64.StdEncoding.EncodeToString(*s.Res))
	} else {
		s.D.Set("content", string(*s.Res))
	}

	return nil
}

func validateFunctionSourceValue(i interface{}, k string) (s []string, es []error) {
	v, ok := i.(string)
	if !ok {
		es = append(es, fmt.Errorf("expected type of %s to be string", k))
		return nil, es
	}
	info, err := os.Stat(v)
	if err != nil {
		es = append(es, fmt.Errorf("cannot get file information for the specified source: %s", v))
		return nil, es
	}
	if info.Size() > 6*1024*1024 {
		es = append(es, fmt.Errorf("the specified source: %s file is too large", v))
		return nil, es
	}
	return
}
