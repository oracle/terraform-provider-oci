// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"regexp"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectstoragePARTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *ResourceObjectstoragePARTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenizeWithHttpReplay("object_storage_resource")
	s.Providers = testAccProviders
	testAccPreCheck(s.T())
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
	data "oci_objectstorage_namespace" "t" {
	}
	
	resource "oci_objectstorage_bucket" "t" {
		compartment_id = "${var.compartment_id}"
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		name = "{{.token}}"
		access_type="ObjectRead"
	}

	resource "oci_objectstorage_object" "t" {
		namespace = "${data.oci_objectstorage_namespace.t.namespace}"
		bucket = "${oci_objectstorage_bucket.t.name}"
		object = "-tf-object"
		content = "123"
	}`, nil)

	s.ResourceName = "oci_objectstorage_preauthrequest.t"
}

func (s *ResourceObjectstoragePARTestSuite) TestAccResourceObjectstoragePAR_basic() {

	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: s.Config + `
				resource "oci_objectstorage_preauthrequest" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					name = "-tf-par"
					access_type = "ObjectRead"
					time_expires = "` + expirationTimeForPar.Format(time.RFC3339Nano) + `"
					object = "-tf-object"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "namespace"),
					resource.TestCheckResourceAttr(s.ResourceName, "bucket", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttr(s.ResourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "access_uri"),
					// regex match example: /p/QJ1Geyhs3WKZvJr8jhw0TeqqqKd4OE1i9ZsGcJ5bzi8/n/internalbriangustafson/b/2018-02-05-130953-145201650/o/
					resource.TestMatchResourceAttr(s.ResourceName, "access_uri", regexp.MustCompile("/p/.*/n/.*/b/"+s.Token+"/o/")),
					resource.TestCheckResourceAttr(s.ResourceName, "object", "-tf-object"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
			// verify access_uri is still available after subsequent refreshes (api only returns this value on create)
			{
				Config: s.Config + `
				resource "oci_objectstorage_preauthrequest" "t" {
					namespace = "${data.oci_objectstorage_namespace.t.namespace}"
					bucket = "${oci_objectstorage_bucket.t.name}"
					name = "-tf-par"
					access_type = "ObjectRead"
					time_expires = "` + expirationTimeForPar.Format(time.RFC3339Nano) + `"
					object = "-tf-object"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(s.ResourceName, "access_uri", regexp.MustCompile("/p/.*/n/.*/b/"+s.Token+"/o/")),
				),
			},
		},
	})
}

func TestUnitResourceObjectstoragePAR_parseIds(t *testing.T) {
	t.Run("Parse Composite Ids", func(t *testing.T) {
		tests := []struct {
			parId       string
			expectError bool
			error       string
			parsedParId string
		}{
			{`n/dxterraformdev/b/bucket/p/dJoeW0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJoeW0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W/0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W/0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W/0iJzmjVX4x6rAKn/UUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W/0iJzmjVX4x6rAKn/UUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0iJzmj/n/VX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W0iJzmj/n/VX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0iJzm/p/jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0in/JzmjVX4x/b/6rAKnUUF/p/8Wx4XAYzwI5YcACNtzyY=:object`, false, "", "dJo/W0in/JzmjVX4x/b/6rAKnUUF/p/8Wx4XAYzwI5YcACNtzyY=:object"},
			{`n/dxterraformdev/b/bucket/p/dJo/W0in/JzmjVX4x/b/6rAKnUUF/p/8Wx4XAY%2FzwI5YcACNtzyY=:object`, false, "", "8Wx4XAY/zwI5YcACNtzyY=:object"},
			{`dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`n/dxterraformdev/p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId n/dxterraformdev/p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`n/dxterraformdev/b/bucket/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId n/dxterraformdev/b/bucket/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId p/dJo/W0iJzmjVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
			{`/b/bucket/p/dJo/W0iJzm/p/jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object`, true, "illegal compositeId /b/bucket/p/dJo/W0iJzm/p/jVX4x6rAKnUUF8Wx4XAYzwI5YcACNtzyY=:object encountered", ""},
		}

		for _, test := range tests {
			if _, _, parId, err := parsePreauthenticatedRequestCompositeId(test.parId); err != nil {

				if test.expectError && err == nil {
					t.Fatalf("expected an error but got none")
				}
				if !test.expectError && err != nil {
					t.Fatalf("did not expect an error but got one %s ", err.Error())
				}
				if test.expectError && err != nil && err.Error() != test.error {
					t.Fatalf("unexpected error %s, expected: %s ", err.Error(), test.error)
				}

				if !test.expectError && err == nil && parId != test.parsedParId {
					t.Fatalf("parId parsed incorrectly, got: %s, expected: %s ", parId, test.parsedParId)
				}
			}
		}
	})
}

func TestResourceObjectstoragePARTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceObjectstoragePARTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceObjectstoragePARTestSuite))
}
