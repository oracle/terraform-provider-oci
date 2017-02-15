// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"errors"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
)

const testProviderConfig = `

provider "baremetal" {
	tenancy_ocid = "ocid.tenancy.aaaa"
	user_ocid = "ocid.user.bbbbb"
	fingerprint = "xxxxxxxxxx"
	private_key_path = "/home/foo/private_key.pem"
	private_key_password = "password"
}

`

// This test runs the Provider sanity checks.
func TestProvider(t *testing.T) {
	// Real client for the sanity check. Makes this more of an acceptance test.
	client := &baremetal.Client{}
	if err := Provider(func(d *schema.ResourceData) (interface{}, error) {
		return client, nil
	}).(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,9F4D00DEF02B2B75

IbSQEhNjPeRt49jUhZbhAEaAIG4L9IokDksw/P/QdCPXzZT008xzYK/zmxkz7so1
ZwvIYHn07E0Ul6fIHR6kjw/+MD7AWluCN1FLHs3PHc4XF4THUCKFCC90FvGJ2PEs
kEh7oJ4azZA/PH51g4rSgWpYtH5B/S6ioE2eZ9jJ/prH+34pCuOpX4AvXEFl5zue
pjFm5FhsReAhZ/9eCvjgjIWDHKc7PRfinwSydVHQSzgDnuq+GTMzQh6eztS+EuAp
MLg7w0mazTqmPOuMT+mw9SHGaIePGzA9TcwB1y3QgkYsg3Ch20uN/sUymgQ4PEKI
njXLldWDYvFvv1Tv3/8IOjCEodQ4P/5oWz7msrLh3QF+EhF7lQPYO7132e9Hvz3C
hTmcygmVGrPCtOY1jzuqy+/Kmt4Gv8FQpSnO7i8wFvt5v0N26av18RO10CzYY1ut
EV6WvynimFUtg1Lo03cadh7bspNohSXfFLpbNTji5NwHrIa+UQqTw3h4/zSPZHJl
NwHwM2I8N5lcCsqmSbM01+uTRG3QZ5i1BS8fsArHaAcvPyLvOy4mZGKkpuNlLDXo
qrCCsb+0m9jHR2bzx5AGp4impdHm2Qi3vTV3dMe277wqKkU5qfd5yDbL2eTqAYzQ
hXpPmTjquOTNYdbvoNsOg4TCHZv7WCsGY0nNMPrRO7zXCDApA6cKDJzagbqhW5Zu
/yz7sDT2D3wzE2WXUbtIBLevXyF0OS3AL7AgfbcyAviByOfmEb7WCP9jmdCFaLwY
SgNh9AjeOgkEEr/cRg1kBAXt0kuE7By0w+/ODJHZYelG0wg5nxhseA9Kc596XIJl
NyjbL87CXGfXmMoSYYTA4rzbtCDMmee7xHtbWiYKF1VGxNaGkQ5nnZSJLhCaI6rH
AD0XYwxv92j4fIjHqonbY/dlIKPot1t3VRcdnebbZMjAcNZ63n+I/iVla3DJpWLO
1gT50A4H2uEAve+WWFWmDQe2rfg5wwUtVVkot+Tn3McB6RzNqgcs0c+7uNDnDcOB
WtQ1OfniE1TdoFCPfYcDw8ngimw7uMYwp4mZIYtwlk7Z5GFl4YpNQeLOgh368ao4
8HL7EnTZmiU5cMbuaA8cZmUbgBqiQY0DtLF22VquThi0QOeUMJxJ6N1QUPckD3AU
dikEn0gilOsDQ51fnOsgk9J2uCz8rd5bnyUXlIguj5pyz6S7agyYFhRrXessVzHd
3889QM9V82+px5mv4qCvMn6ReYOvC+KSY1hn4ljXsndOM+6hQzD5CZKeL948pXRn
G7nqbG9D44wLklOz6mkIvqLn3qxEFWapl9UK7yfzjoezGoqeNFweadZ10Kp2+Umu
Sa759/2YDCZLDzaVVoLDTHLzi9ejpAkUIXgEFaPNGzQ8DYiL8N2klRozLSlnDEMr
xTHuOMkklNO7SiTluAUBvXrjxfGqe/gwJOHxXQGHC8W6vyhR2BdVx9PKFVebWjlr
gzRMpGgWnjsaz0ldu3uO7ozRxZg8FgdToIzAIaTytpHKI8HvONvPJlYywOMC1gRi
KwX6p26xaVtCV8PbDpF3RHuEJV1NU6PDIhaIHhdL374BiX/KmcJ6yv7tbkczpK+V
-----END RSA PRIVATE KEY-----`

var testKeyFingerPrint = "b4:8a:7d:54:e6:81:04:b2:99:8e:b3:ed:10:e2:12:2b"
var testTenancyOCID = "ocid1.tenancy.oc1..aaaaaaaaq3hulfjvrouw3e6qx2ncxtp256aq7etiabqqtzunnhxjslzkfyxq"
var testUserOCID = "ocid1.user.oc1..aaaaaaaaflxvsdpjs5ztahmsf7vjxy5kdqnuzyqpvwnncbkfhavexwd4w5ra"

func TestProviderConfig(t *testing.T) {
	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)
	d.SetId("tenancy_ocid")

	d.Set("tenancy_ocid", testTenancyOCID)
	d.Set("user_ocid", testUserOCID)
	d.Set("fingerprint", testKeyFingerPrint)
	d.Set("private_key", testPrivateKey)
	d.Set("private_key_password", "password")

	client, err := providerConfig(d)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	_, ok := client.(*baremetal.Client)
	assert.True(t, ok)
}

// TestNoInstanceState determines if there is any state for a given name.
func testNoInstanceState(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ms := s.RootModule()
		rs, ok := ms.Resources[name]
		if !ok {
			return nil
		}

		is := rs.Primary
		if is == nil {
			return nil
		}

		return errors.New("State exists for primary resource " + name)
	}
}
