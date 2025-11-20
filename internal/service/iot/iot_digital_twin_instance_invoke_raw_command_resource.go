// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotDigitalTwinInstanceInvokeRawCommandResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotDigitalTwinInstanceInvokeRawCommandWithContext,
		ReadContext:   readIotDigitalTwinInstanceInvokeRawCommandWithContext,
		DeleteContext: deleteIotDigitalTwinInstanceInvokeRawCommandWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"digital_twin_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"request_data_format": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BINARY",
					"JSON",
					"TEXT",
				}, true),
			},
			"request_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"request_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"request_data_content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"request_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"response_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"response_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createIotDigitalTwinInstanceInvokeRawCommandWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinInstanceInvokeRawCommandResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotDigitalTwinInstanceInvokeRawCommandWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteIotDigitalTwinInstanceInvokeRawCommandWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type IotDigitalTwinInstanceInvokeRawCommandResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.InvokeRawCommandResponse
	DisableNotFoundRetries bool
}

func (s *IotDigitalTwinInstanceInvokeRawCommandResourceCrud) ID() string {
	if s.Res == nil || s.Res.Location == nil || *s.Res.Location == "" {
		return ""
	}
	urlStr := strings.TrimRight(*s.Res.Location, "/")
	parts := strings.Split(urlStr, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func (s *IotDigitalTwinInstanceInvokeRawCommandResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.InvokeRawCommandRequest{}
	err := s.populateTopLevelPolymorphicInvokeRawCommandRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.InvokeRawCommand(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IotDigitalTwinInstanceInvokeRawCommandResourceCrud) SetData() error {
	return nil
}

func (s *IotDigitalTwinInstanceInvokeRawCommandResourceCrud) populateTopLevelPolymorphicInvokeRawCommandRequest(request *oci_iot.InvokeRawCommandRequest) error {
	//discriminator
	requestDataFormatRaw, ok := s.D.GetOkExists("request_data_format")
	var requestDataFormat string
	if ok {
		requestDataFormat = requestDataFormatRaw.(string)
	} else {
		requestDataFormat = "" // default value
	}
	switch strings.ToLower(requestDataFormat) {
	case strings.ToLower("BINARY"):
		details := oci_iot.InvokeRawBinaryCommandDetails{}
		if requestData, ok := s.D.GetOkExists("request_data"); ok {
			tmp := requestData.(string)
			details.RequestData = &tmp
		}
		if requestDataContentType, ok := s.D.GetOkExists("request_data_content_type"); ok {
			tmp := requestDataContentType.(string)
			details.RequestDataContentType = &tmp
		}
		if digitalTwinInstanceId, ok := s.D.GetOkExists("digital_twin_instance_id"); ok {
			tmp := digitalTwinInstanceId.(string)
			request.DigitalTwinInstanceId = &tmp
		}
		if requestDuration, ok := s.D.GetOkExists("request_duration"); ok {
			tmp := requestDuration.(string)
			details.RequestDuration = &tmp
		}
		if requestEndpoint, ok := s.D.GetOkExists("request_endpoint"); ok {
			tmp := requestEndpoint.(string)
			details.RequestEndpoint = &tmp
		}
		if responseDuration, ok := s.D.GetOkExists("response_duration"); ok {
			tmp := responseDuration.(string)
			details.ResponseDuration = &tmp
		}
		if responseEndpoint, ok := s.D.GetOkExists("response_endpoint"); ok {
			tmp := responseEndpoint.(string)
			details.ResponseEndpoint = &tmp
		}
		request.InvokeRawCommandDetails = details
	case strings.ToLower("JSON"):
		details := oci_iot.InvokeRawJsonCommandDetails{}
		if requestData, ok := s.D.GetOkExists("request_data"); ok {
			// Parse the JSON string into a map
			requestDataMap, err := JsonStringToMap(requestData.(string))
			if err != nil {
				return err
			}
			details.RequestData = requestDataMap
		}
		if requestDataContentType, ok := s.D.GetOkExists("request_data_content_type"); ok {
			tmp := requestDataContentType.(string)
			details.RequestDataContentType = &tmp
		}
		if digitalTwinInstanceId, ok := s.D.GetOkExists("digital_twin_instance_id"); ok {
			tmp := digitalTwinInstanceId.(string)
			request.DigitalTwinInstanceId = &tmp
		}
		if requestDuration, ok := s.D.GetOkExists("request_duration"); ok {
			tmp := requestDuration.(string)
			details.RequestDuration = &tmp
		}
		if requestEndpoint, ok := s.D.GetOkExists("request_endpoint"); ok {
			tmp := requestEndpoint.(string)
			details.RequestEndpoint = &tmp
		}
		if responseDuration, ok := s.D.GetOkExists("response_duration"); ok {
			tmp := responseDuration.(string)
			details.ResponseDuration = &tmp
		}
		if responseEndpoint, ok := s.D.GetOkExists("response_endpoint"); ok {
			tmp := responseEndpoint.(string)
			details.ResponseEndpoint = &tmp
		}
		request.InvokeRawCommandDetails = details
	case strings.ToLower("TEXT"):
		details := oci_iot.InvokeRawTextCommandDetails{}
		if requestData, ok := s.D.GetOkExists("request_data"); ok {
			tmp := requestData.(string)
			details.RequestData = &tmp
		}
		if requestDataContentType, ok := s.D.GetOkExists("request_data_content_type"); ok {
			tmp := requestDataContentType.(string)
			details.RequestDataContentType = &tmp
		}
		if digitalTwinInstanceId, ok := s.D.GetOkExists("digital_twin_instance_id"); ok {
			tmp := digitalTwinInstanceId.(string)
			request.DigitalTwinInstanceId = &tmp
		}
		if requestDuration, ok := s.D.GetOkExists("request_duration"); ok {
			tmp := requestDuration.(string)
			details.RequestDuration = &tmp
		}
		if requestEndpoint, ok := s.D.GetOkExists("request_endpoint"); ok {
			tmp := requestEndpoint.(string)
			details.RequestEndpoint = &tmp
		}
		if responseDuration, ok := s.D.GetOkExists("response_duration"); ok {
			tmp := responseDuration.(string)
			details.ResponseDuration = &tmp
		}
		if responseEndpoint, ok := s.D.GetOkExists("response_endpoint"); ok {
			tmp := responseEndpoint.(string)
			details.ResponseEndpoint = &tmp
		}
		request.InvokeRawCommandDetails = details
	default:
		return fmt.Errorf("unknown request_data_format '%v' was specified", requestDataFormat)
	}
	return nil
}
