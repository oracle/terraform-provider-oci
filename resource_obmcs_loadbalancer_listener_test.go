// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

type ResourceLoadBalancerListenerTestSuite struct {
	suite.Suite
	Client      client.BareMetalClient
	Providers   map[string]terraform.ResourceProvider
	TimeCreated baremetal.Time
}

func (s *ResourceLoadBalancerListenerTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": Provider(
			func(d *schema.ResourceData) (interface{}, error) {
				return s.Client, nil
			},
		),
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}
}

func (s *ResourceLoadBalancerListenerTestSuite) TestCreateResourceLoadBalancerListenerrMaximal() {
	resourceName := "baremetal_load_balancer_listener.t"
	config := `
resource "baremetal_load_balancer_listener" "t" {
  load_balancer_id         = "stub_load_balancer_id"
  name                     = "stub_name"
  default_backend_set_name = "stub_backend_set_name"
  port                     = 1234
  protocol                 = "stub_protocol"

  ssl_configuration {
      certificate_name        = "stub_certificate_name"
      verify_depth            = 6
      verify_peer_certificate = false
  }
}
`
	config += testProviderConfig

	loadBalancerID := "stub_load_balancer_id"
	res := &baremetal.Listener{
		Name: "stub_name",
		DefaultBackendSetName: "stub_backend_set_name",
		Port:     1234,
		Protocol: "stub_protocol",
		SSLConfig: &baremetal.SSLConfiguration{
			CertificateName:       "stub_certificate_name",
			VerifyDepth:           6,
			VerifyPeerCertificate: false,
		},
	}
	res.RequestID = "stub_opc_request_id"

	deletedRes := &baremetal.Listener{}
	*deletedRes = *res

	workReqID := "stub_work_req_id"

	s.Client.On(
		"CreateListener",
		loadBalancerID,
		res.Name,
		res.DefaultBackendSetName,
		res.Protocol,
		res.Port,
		res.SSLConfig,
		(*baremetal.LoadBalancerOptions)(nil),
	).Return(workReqID, nil)

	workReqCreated := &baremetal.WorkRequest{
		ID:             workReqID,
		LoadBalancerID: loadBalancerID,
		State:          baremetal.WorkRequestSucceeded,
	}

	s.Client.On(
		"GetWorkRequest",
		workReqID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqCreated, nil)

	lb := baremetal.LoadBalancer{
		Listeners: map[string]baremetal.Listener{
			res.Name: *res,
		},
	}
	s.Client.On(
		"GetLoadBalancer",
		loadBalancerID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(&lb, nil)

	s.Client.On(
		"DeleteListener",
		loadBalancerID,
		res.Name,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqID, nil)

	s.Client.On(
		"GetWorkRequest",
		workReqID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(workReqCreated, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer_id", loadBalancerID),
					resource.TestCheckResourceAttr(resourceName, "name", res.Name),
					resource.TestCheckResourceAttr(resourceName, "default_backend_set_name", res.DefaultBackendSetName),
					resource.TestCheckResourceAttr(resourceName, "port", strconv.Itoa(res.Port)),
					resource.TestCheckResourceAttr(resourceName, "protocol", res.Protocol),

					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.certificate_name", res.SSLConfig.CertificateName),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_depth", strconv.Itoa(res.SSLConfig.VerifyDepth)),
					resource.TestCheckResourceAttr(resourceName, "ssl_configuration.0.verify_peer_certificate", strconv.FormatBool(res.SSLConfig.VerifyPeerCertificate)),
				),
			},
		},
	})
}

func TestResourceLoadBalancerListenerTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceLoadBalancerListenerTestSuite))
}
