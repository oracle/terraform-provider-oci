// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	oci_apigateway "github.com/oracle/oci-go-sdk/v27/apigateway"
	"github.com/oracle/oci-go-sdk/v27/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ApiGatewayCertificateRequiredOnlyResource = ApiGatewayCertificateResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateRepresentation)

	CertificateResourceConfig = ApiGatewayCertificateResourceDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation)

	apiGatewaycertificateSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id": Representation{repType: Required, create: `${oci_apigateway_certificate.test_certificate.id}`},
	}

	apiGatewaycertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, apiGatewaycertificateDataSourceFilterRepresentation}}
	apiGatewaycertificateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_apigateway_certificate.test_certificate.id}`}},
	}

	apiGatewaycertificateRepresentation = map[string]interface{}{
		"certificate":               Representation{repType: Required, create: `-----BEGIN CERTIFICATE-----\nMIIFZTCCBE2gAwIBAgISA6/vHTH2+uHo4BFuKgs90rQ2MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA3MDMxMzIwMTNaFw0y\nMDEwMDExMzIwMTNaMBsxGTAXBgNVBAMTEHd3dy5wYXRyaWNoaS5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWOPcmSEgjKK9o13/FYtJ5U4pzGTbE\n1EuaFPLC9KVFRcIN+16f72BdNjDGxm5b9CzfkfjcgKJpunCGWE7r6ukFch8JDKxn\nyNu6XDrgjhymm0cWn+UxaipzronT9A8RDUllXTD8UJzJjYi/x6/bsds5u0C03CIc\ni0ig5jwJrCKvKGUhsjX1vx8mKibjvETjmkLhZJ+50PYuDUW4MFFqjIKxwIKh+szm\ne2B+lCJQGYhiAY9R7R0J41M7tzYZj4IzuYh7mDt/+t+JOuz1OUenMwAFMauXS9yW\nL0oY7sNEW1cnC2QWH4EHQfyOdN0HkYbsy/0Eia7KzcbgK8zvipjRzbO3AgMBAAGj\nggJyMIICbjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsG\nAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFFHRKpBm95oY9EIO164b+cou\n6qHiMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUFBwEB\nBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZy8wKQYDVR0RBCIwIIIMcGF0cmljaGkuY29tghB3d3cucGF0cmljaGkuY29t\nMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUH\nAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB\n9ASB8QDvAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklTLe+LkF4AAAFzFQyZ\nUAAABAMARjBEAiBNyGhSkSA/QAV8O/z4p1sZxUbHX2LFaYhRn2gnRAmkRwIge9JC\nxsrSGMBtSOwN/mGBCr2yWsO7Xq0mFa7HvdyT5N0AdgBvU3asMfAxGdiZAKRRFf93\nFRwR2QLBACkGjbIImjfZEwAAAXMVDJl7AAAEAwBHMEUCICz3LqARn2byH91k3RZ/\nXloOaCPFI1hhzGk4LZbB7ZxOAiEA9CIbpEytmD0WctjD3HEsLHgpKtFCyFdOMrcS\n3W7+jh0wDQYJKoZIhvcNAQELBQADggEBAIBNvSrPD+01gTK64c6DkUix1xkdkzK6\nbMzjGU6lvcONhCSJGi4WZyl1m/nJfjuJ66xJSVlw0bcTbBmSYQCKWodkDPgR3HFM\nVJPXQuuLW7uNCMjqgz4h6o3WEWI1mKP5Vf5mPHVfgErgfk7TYtVWdqnB2/zJ72Tw\nvJwiSZjbEUBaxlw8BVdK224taStb1i1fM+xP+GOEaVlr6I59E1GpfOp1iwPBPzZi\nrb/w/03eCJn0P1+yTwBskvdsFRJaBbvd5mCGzAHcaAgMvbqv+u/7e01XgzPs8U7t\nHQ+a1q07HAmnobmXvkljO68T7MuUomFLmEX3RnRoNcXxqhNnMQnbQZ4=\n-----END CERTIFICATE-----`},
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"private_key":               Representation{repType: Required, create: `-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1jj3JkhIIyivaNd/xWLSeVOKcxk2xNRLmhTywvSlRUXCDfte\nn+9gXTYwxsZuW/Qs35H43ICiabpwhlhO6+rpBXIfCQysZ8jbulw64I4cpptHFp/l\nMWoqc66J0/QPEQ1JZV0w/FCcyY2Iv8ev27HbObtAtNwiHItIoOY8CawiryhlIbI1\n9b8fJiom47xE45pC4WSfudD2Lg1FuDBRaoyCscCCofrM5ntgfpQiUBmIYgGPUe0d\nCeNTO7c2GY+CM7mIe5g7f/rfiTrs9TlHpzMABTGrl0vcli9KGO7DRFtXJwtkFh+B\nB0H8jnTdB5GG7Mv9BImuys3G4CvM74qY0c2ztwIDAQABAoIBAQCgnYtor48ulUWX\nFOQeqW5nOyS2EXSH9ShN+WDenTEBFEFf3FVhEsgsewHtNz/tP9EZO0fcg7HCFeBi\nSfm6VqGpzJvKXC8zjVx5iMc4MDT5EbkmFHJyL0hu/bEnMnjNbynVjx64PvSfBbg/\nGkgXL23QBj6Im/gTcLbetGDVW7ORoRetRUxgOAGq+rrVrZI2k2n3q2Q+yfMhfcoC\ni9bi62tPeroDfrh4sy5EpkTL8ddIzy3Mz0EdRWpyu17lml+rUgceoLSjZVZVc0Cs\n5mre6Q8m5W0590Rd12yubLliq2E66cKom/wuIaGGy36+8zYbM1wiT8svylfzif4W\nAUSIGWRJAoGBAOwEMuxCoYvcJREG1R7QVc/egR8UrWlM6SXsbc5D96mHBT+DaaPw\nPXReFx9wIQ5J+uNRgTJ2GRfDS05t6i3kHtxDRGmz+qtaakAClmWcKjRBq6gdGg1j\nBcb86aVb7zr/x/x4jL9DQV5CJfC8Qs3Q966BbKzxP0zrPTPs0jLbmNC1AoGBAOhc\nXuZOqTeu4cSvWbcrvYZpoGIuhLCTV3x5OEPF847tEQ/U7m2gYQLpKU5oQPDIboCi\nYdSOs8tszefSW/438+VFPmUaClVVsEOU3hmkRjHPji6eX/pdDmWglZly4egoJNHM\nCtzDSiXU/uf9C1UaIjgnI0BW/jxdnC8Scn0eMXI7AoGAYfNwKuuQXho0a/eY9Zvj\nMU0X641KVvxreqi//a3pmDsIK47fhgFLkAMIt6xym5qKfFM0OnwW9+y+UUN+wCL/\nx8xHFVNzwi/ZDs3EG1GPlnZ2xmOlj068dBQ857ra83J6Tka9qxc/ht4PHvUTCJYk\nZREjyDarXcH3eZhcuGy45E0CgYEA1B222EHIwjdYjIeSu98nNbPpIJfcs6DeBZhX\nX68uZzRNFgnI2rTEuraE6bnMRsKB9dXGvxbCVRrvYQgsyIbllE/A5OR/uGTv8tHM\natYG6mPSJQaZEWGvyeBtkNZKGffDnO3KoLt6Tc0CIl9i3/5qbJ511L9VeV/vOx14\n/HT6qI8CgYA+BTF/aS9iQGrRI1LpMTQh5sbSnMz39Ri2VWulp/MLAOBCTKEpB9+W\neL+z52ZKn/kqo9hP/Ysl1Pt+JymPY1tP0TMJ4QCpVHnrmbx0xyNwZldR5kvgc/DN\neRzKmOYlhMF5LcqqdOaOOP4fbPAReq+FudAZuhrL4iReWnOR+I1Stw==\n-----END RSA PRIVATE KEY-----`},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":              Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"intermediate_certificates": Representation{repType: Optional, create: `-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----`},
	}

	ApiGatewayCertificateResourceDependencies = DefinedTagsDependencies
)

func TestApigatewayCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayCertificateResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apigateway_certificate.test_certificate"
	datasourceName := "data.oci_apigateway_certificates.test_certificates"
	singularDatasourceName := "data.oci_apigateway_certificate.test_certificate"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApigatewayCertificateDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate", "-----BEGIN CERTIFICATE-----\nMIIFZTCCBE2gAwIBAgISA6/vHTH2+uHo4BFuKgs90rQ2MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA3MDMxMzIwMTNaFw0y\nMDEwMDExMzIwMTNaMBsxGTAXBgNVBAMTEHd3dy5wYXRyaWNoaS5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWOPcmSEgjKK9o13/FYtJ5U4pzGTbE\n1EuaFPLC9KVFRcIN+16f72BdNjDGxm5b9CzfkfjcgKJpunCGWE7r6ukFch8JDKxn\nyNu6XDrgjhymm0cWn+UxaipzronT9A8RDUllXTD8UJzJjYi/x6/bsds5u0C03CIc\ni0ig5jwJrCKvKGUhsjX1vx8mKibjvETjmkLhZJ+50PYuDUW4MFFqjIKxwIKh+szm\ne2B+lCJQGYhiAY9R7R0J41M7tzYZj4IzuYh7mDt/+t+JOuz1OUenMwAFMauXS9yW\nL0oY7sNEW1cnC2QWH4EHQfyOdN0HkYbsy/0Eia7KzcbgK8zvipjRzbO3AgMBAAGj\nggJyMIICbjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsG\nAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFFHRKpBm95oY9EIO164b+cou\n6qHiMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUFBwEB\nBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZy8wKQYDVR0RBCIwIIIMcGF0cmljaGkuY29tghB3d3cucGF0cmljaGkuY29t\nMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUH\nAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB\n9ASB8QDvAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklTLe+LkF4AAAFzFQyZ\nUAAABAMARjBEAiBNyGhSkSA/QAV8O/z4p1sZxUbHX2LFaYhRn2gnRAmkRwIge9JC\nxsrSGMBtSOwN/mGBCr2yWsO7Xq0mFa7HvdyT5N0AdgBvU3asMfAxGdiZAKRRFf93\nFRwR2QLBACkGjbIImjfZEwAAAXMVDJl7AAAEAwBHMEUCICz3LqARn2byH91k3RZ/\nXloOaCPFI1hhzGk4LZbB7ZxOAiEA9CIbpEytmD0WctjD3HEsLHgpKtFCyFdOMrcS\n3W7+jh0wDQYJKoZIhvcNAQELBQADggEBAIBNvSrPD+01gTK64c6DkUix1xkdkzK6\nbMzjGU6lvcONhCSJGi4WZyl1m/nJfjuJ66xJSVlw0bcTbBmSYQCKWodkDPgR3HFM\nVJPXQuuLW7uNCMjqgz4h6o3WEWI1mKP5Vf5mPHVfgErgfk7TYtVWdqnB2/zJ72Tw\nvJwiSZjbEUBaxlw8BVdK224taStb1i1fM+xP+GOEaVlr6I59E1GpfOp1iwPBPzZi\nrb/w/03eCJn0P1+yTwBskvdsFRJaBbvd5mCGzAHcaAgMvbqv+u/7e01XgzPs8U7t\nHQ+a1q07HAmnobmXvkljO68T7MuUomFLmEX3RnRoNcXxqhNnMQnbQZ4=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1jj3JkhIIyivaNd/xWLSeVOKcxk2xNRLmhTywvSlRUXCDfte\nn+9gXTYwxsZuW/Qs35H43ICiabpwhlhO6+rpBXIfCQysZ8jbulw64I4cpptHFp/l\nMWoqc66J0/QPEQ1JZV0w/FCcyY2Iv8ev27HbObtAtNwiHItIoOY8CawiryhlIbI1\n9b8fJiom47xE45pC4WSfudD2Lg1FuDBRaoyCscCCofrM5ntgfpQiUBmIYgGPUe0d\nCeNTO7c2GY+CM7mIe5g7f/rfiTrs9TlHpzMABTGrl0vcli9KGO7DRFtXJwtkFh+B\nB0H8jnTdB5GG7Mv9BImuys3G4CvM74qY0c2ztwIDAQABAoIBAQCgnYtor48ulUWX\nFOQeqW5nOyS2EXSH9ShN+WDenTEBFEFf3FVhEsgsewHtNz/tP9EZO0fcg7HCFeBi\nSfm6VqGpzJvKXC8zjVx5iMc4MDT5EbkmFHJyL0hu/bEnMnjNbynVjx64PvSfBbg/\nGkgXL23QBj6Im/gTcLbetGDVW7ORoRetRUxgOAGq+rrVrZI2k2n3q2Q+yfMhfcoC\ni9bi62tPeroDfrh4sy5EpkTL8ddIzy3Mz0EdRWpyu17lml+rUgceoLSjZVZVc0Cs\n5mre6Q8m5W0590Rd12yubLliq2E66cKom/wuIaGGy36+8zYbM1wiT8svylfzif4W\nAUSIGWRJAoGBAOwEMuxCoYvcJREG1R7QVc/egR8UrWlM6SXsbc5D96mHBT+DaaPw\nPXReFx9wIQ5J+uNRgTJ2GRfDS05t6i3kHtxDRGmz+qtaakAClmWcKjRBq6gdGg1j\nBcb86aVb7zr/x/x4jL9DQV5CJfC8Qs3Q966BbKzxP0zrPTPs0jLbmNC1AoGBAOhc\nXuZOqTeu4cSvWbcrvYZpoGIuhLCTV3x5OEPF847tEQ/U7m2gYQLpKU5oQPDIboCi\nYdSOs8tszefSW/438+VFPmUaClVVsEOU3hmkRjHPji6eX/pdDmWglZly4egoJNHM\nCtzDSiXU/uf9C1UaIjgnI0BW/jxdnC8Scn0eMXI7AoGAYfNwKuuQXho0a/eY9Zvj\nMU0X641KVvxreqi//a3pmDsIK47fhgFLkAMIt6xym5qKfFM0OnwW9+y+UUN+wCL/\nx8xHFVNzwi/ZDs3EG1GPlnZ2xmOlj068dBQ857ra83J6Tka9qxc/ht4PHvUTCJYk\nZREjyDarXcH3eZhcuGy45E0CgYEA1B222EHIwjdYjIeSu98nNbPpIJfcs6DeBZhX\nX68uZzRNFgnI2rTEuraE6bnMRsKB9dXGvxbCVRrvYQgsyIbllE/A5OR/uGTv8tHM\natYG6mPSJQaZEWGvyeBtkNZKGffDnO3KoLt6Tc0CIl9i3/5qbJ511L9VeV/vOx14\n/HT6qI8CgYA+BTF/aS9iQGrRI1LpMTQh5sbSnMz39Ri2VWulp/MLAOBCTKEpB9+W\neL+z52ZKn/kqo9hP/Ysl1Pt+JymPY1tP0TMJ4QCpVHnrmbx0xyNwZldR5kvgc/DN\neRzKmOYlhMF5LcqqdOaOOP4fbPAReq+FudAZuhrL4iReWnOR+I1Stw==\n-----END RSA PRIVATE KEY-----"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Create, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate", "-----BEGIN CERTIFICATE-----\nMIIFZTCCBE2gAwIBAgISA6/vHTH2+uHo4BFuKgs90rQ2MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA3MDMxMzIwMTNaFw0y\nMDEwMDExMzIwMTNaMBsxGTAXBgNVBAMTEHd3dy5wYXRyaWNoaS5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWOPcmSEgjKK9o13/FYtJ5U4pzGTbE\n1EuaFPLC9KVFRcIN+16f72BdNjDGxm5b9CzfkfjcgKJpunCGWE7r6ukFch8JDKxn\nyNu6XDrgjhymm0cWn+UxaipzronT9A8RDUllXTD8UJzJjYi/x6/bsds5u0C03CIc\ni0ig5jwJrCKvKGUhsjX1vx8mKibjvETjmkLhZJ+50PYuDUW4MFFqjIKxwIKh+szm\ne2B+lCJQGYhiAY9R7R0J41M7tzYZj4IzuYh7mDt/+t+JOuz1OUenMwAFMauXS9yW\nL0oY7sNEW1cnC2QWH4EHQfyOdN0HkYbsy/0Eia7KzcbgK8zvipjRzbO3AgMBAAGj\nggJyMIICbjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsG\nAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFFHRKpBm95oY9EIO164b+cou\n6qHiMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUFBwEB\nBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZy8wKQYDVR0RBCIwIIIMcGF0cmljaGkuY29tghB3d3cucGF0cmljaGkuY29t\nMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUH\nAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB\n9ASB8QDvAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklTLe+LkF4AAAFzFQyZ\nUAAABAMARjBEAiBNyGhSkSA/QAV8O/z4p1sZxUbHX2LFaYhRn2gnRAmkRwIge9JC\nxsrSGMBtSOwN/mGBCr2yWsO7Xq0mFa7HvdyT5N0AdgBvU3asMfAxGdiZAKRRFf93\nFRwR2QLBACkGjbIImjfZEwAAAXMVDJl7AAAEAwBHMEUCICz3LqARn2byH91k3RZ/\nXloOaCPFI1hhzGk4LZbB7ZxOAiEA9CIbpEytmD0WctjD3HEsLHgpKtFCyFdOMrcS\n3W7+jh0wDQYJKoZIhvcNAQELBQADggEBAIBNvSrPD+01gTK64c6DkUix1xkdkzK6\nbMzjGU6lvcONhCSJGi4WZyl1m/nJfjuJ66xJSVlw0bcTbBmSYQCKWodkDPgR3HFM\nVJPXQuuLW7uNCMjqgz4h6o3WEWI1mKP5Vf5mPHVfgErgfk7TYtVWdqnB2/zJ72Tw\nvJwiSZjbEUBaxlw8BVdK224taStb1i1fM+xP+GOEaVlr6I59E1GpfOp1iwPBPzZi\nrb/w/03eCJn0P1+yTwBskvdsFRJaBbvd5mCGzAHcaAgMvbqv+u/7e01XgzPs8U7t\nHQ+a1q07HAmnobmXvkljO68T7MuUomFLmEX3RnRoNcXxqhNnMQnbQZ4=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "intermediate_certificates", "-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1jj3JkhIIyivaNd/xWLSeVOKcxk2xNRLmhTywvSlRUXCDfte\nn+9gXTYwxsZuW/Qs35H43ICiabpwhlhO6+rpBXIfCQysZ8jbulw64I4cpptHFp/l\nMWoqc66J0/QPEQ1JZV0w/FCcyY2Iv8ev27HbObtAtNwiHItIoOY8CawiryhlIbI1\n9b8fJiom47xE45pC4WSfudD2Lg1FuDBRaoyCscCCofrM5ntgfpQiUBmIYgGPUe0d\nCeNTO7c2GY+CM7mIe5g7f/rfiTrs9TlHpzMABTGrl0vcli9KGO7DRFtXJwtkFh+B\nB0H8jnTdB5GG7Mv9BImuys3G4CvM74qY0c2ztwIDAQABAoIBAQCgnYtor48ulUWX\nFOQeqW5nOyS2EXSH9ShN+WDenTEBFEFf3FVhEsgsewHtNz/tP9EZO0fcg7HCFeBi\nSfm6VqGpzJvKXC8zjVx5iMc4MDT5EbkmFHJyL0hu/bEnMnjNbynVjx64PvSfBbg/\nGkgXL23QBj6Im/gTcLbetGDVW7ORoRetRUxgOAGq+rrVrZI2k2n3q2Q+yfMhfcoC\ni9bi62tPeroDfrh4sy5EpkTL8ddIzy3Mz0EdRWpyu17lml+rUgceoLSjZVZVc0Cs\n5mre6Q8m5W0590Rd12yubLliq2E66cKom/wuIaGGy36+8zYbM1wiT8svylfzif4W\nAUSIGWRJAoGBAOwEMuxCoYvcJREG1R7QVc/egR8UrWlM6SXsbc5D96mHBT+DaaPw\nPXReFx9wIQ5J+uNRgTJ2GRfDS05t6i3kHtxDRGmz+qtaakAClmWcKjRBq6gdGg1j\nBcb86aVb7zr/x/x4jL9DQV5CJfC8Qs3Q966BbKzxP0zrPTPs0jLbmNC1AoGBAOhc\nXuZOqTeu4cSvWbcrvYZpoGIuhLCTV3x5OEPF847tEQ/U7m2gYQLpKU5oQPDIboCi\nYdSOs8tszefSW/438+VFPmUaClVVsEOU3hmkRjHPji6eX/pdDmWglZly4egoJNHM\nCtzDSiXU/uf9C1UaIjgnI0BW/jxdnC8Scn0eMXI7AoGAYfNwKuuQXho0a/eY9Zvj\nMU0X641KVvxreqi//a3pmDsIK47fhgFLkAMIt6xym5qKfFM0OnwW9+y+UUN+wCL/\nx8xHFVNzwi/ZDs3EG1GPlnZ2xmOlj068dBQ857ra83J6Tka9qxc/ht4PHvUTCJYk\nZREjyDarXcH3eZhcuGy45E0CgYEA1B222EHIwjdYjIeSu98nNbPpIJfcs6DeBZhX\nX68uZzRNFgnI2rTEuraE6bnMRsKB9dXGvxbCVRrvYQgsyIbllE/A5OR/uGTv8tHM\natYG6mPSJQaZEWGvyeBtkNZKGffDnO3KoLt6Tc0CIl9i3/5qbJ511L9VeV/vOx14\n/HT6qI8CgYA+BTF/aS9iQGrRI1LpMTQh5sbSnMz39Ri2VWulp/MLAOBCTKEpB9+W\neL+z52ZKn/kqo9hP/Ysl1Pt+JymPY1tP0TMJ4QCpVHnrmbx0xyNwZldR5kvgc/DN\neRzKmOYlhMF5LcqqdOaOOP4fbPAReq+FudAZuhrL4iReWnOR+I1Stw==\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttrSet(resourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Create,
						representationCopyWithNewProperties(apiGatewaycertificateRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate", "-----BEGIN CERTIFICATE-----\nMIIFZTCCBE2gAwIBAgISA6/vHTH2+uHo4BFuKgs90rQ2MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA3MDMxMzIwMTNaFw0y\nMDEwMDExMzIwMTNaMBsxGTAXBgNVBAMTEHd3dy5wYXRyaWNoaS5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWOPcmSEgjKK9o13/FYtJ5U4pzGTbE\n1EuaFPLC9KVFRcIN+16f72BdNjDGxm5b9CzfkfjcgKJpunCGWE7r6ukFch8JDKxn\nyNu6XDrgjhymm0cWn+UxaipzronT9A8RDUllXTD8UJzJjYi/x6/bsds5u0C03CIc\ni0ig5jwJrCKvKGUhsjX1vx8mKibjvETjmkLhZJ+50PYuDUW4MFFqjIKxwIKh+szm\ne2B+lCJQGYhiAY9R7R0J41M7tzYZj4IzuYh7mDt/+t+JOuz1OUenMwAFMauXS9yW\nL0oY7sNEW1cnC2QWH4EHQfyOdN0HkYbsy/0Eia7KzcbgK8zvipjRzbO3AgMBAAGj\nggJyMIICbjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsG\nAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFFHRKpBm95oY9EIO164b+cou\n6qHiMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUFBwEB\nBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZy8wKQYDVR0RBCIwIIIMcGF0cmljaGkuY29tghB3d3cucGF0cmljaGkuY29t\nMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUH\nAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB\n9ASB8QDvAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklTLe+LkF4AAAFzFQyZ\nUAAABAMARjBEAiBNyGhSkSA/QAV8O/z4p1sZxUbHX2LFaYhRn2gnRAmkRwIge9JC\nxsrSGMBtSOwN/mGBCr2yWsO7Xq0mFa7HvdyT5N0AdgBvU3asMfAxGdiZAKRRFf93\nFRwR2QLBACkGjbIImjfZEwAAAXMVDJl7AAAEAwBHMEUCICz3LqARn2byH91k3RZ/\nXloOaCPFI1hhzGk4LZbB7ZxOAiEA9CIbpEytmD0WctjD3HEsLHgpKtFCyFdOMrcS\n3W7+jh0wDQYJKoZIhvcNAQELBQADggEBAIBNvSrPD+01gTK64c6DkUix1xkdkzK6\nbMzjGU6lvcONhCSJGi4WZyl1m/nJfjuJ66xJSVlw0bcTbBmSYQCKWodkDPgR3HFM\nVJPXQuuLW7uNCMjqgz4h6o3WEWI1mKP5Vf5mPHVfgErgfk7TYtVWdqnB2/zJ72Tw\nvJwiSZjbEUBaxlw8BVdK224taStb1i1fM+xP+GOEaVlr6I59E1GpfOp1iwPBPzZi\nrb/w/03eCJn0P1+yTwBskvdsFRJaBbvd5mCGzAHcaAgMvbqv+u/7e01XgzPs8U7t\nHQ+a1q07HAmnobmXvkljO68T7MuUomFLmEX3RnRoNcXxqhNnMQnbQZ4=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "intermediate_certificates", "-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1jj3JkhIIyivaNd/xWLSeVOKcxk2xNRLmhTywvSlRUXCDfte\nn+9gXTYwxsZuW/Qs35H43ICiabpwhlhO6+rpBXIfCQysZ8jbulw64I4cpptHFp/l\nMWoqc66J0/QPEQ1JZV0w/FCcyY2Iv8ev27HbObtAtNwiHItIoOY8CawiryhlIbI1\n9b8fJiom47xE45pC4WSfudD2Lg1FuDBRaoyCscCCofrM5ntgfpQiUBmIYgGPUe0d\nCeNTO7c2GY+CM7mIe5g7f/rfiTrs9TlHpzMABTGrl0vcli9KGO7DRFtXJwtkFh+B\nB0H8jnTdB5GG7Mv9BImuys3G4CvM74qY0c2ztwIDAQABAoIBAQCgnYtor48ulUWX\nFOQeqW5nOyS2EXSH9ShN+WDenTEBFEFf3FVhEsgsewHtNz/tP9EZO0fcg7HCFeBi\nSfm6VqGpzJvKXC8zjVx5iMc4MDT5EbkmFHJyL0hu/bEnMnjNbynVjx64PvSfBbg/\nGkgXL23QBj6Im/gTcLbetGDVW7ORoRetRUxgOAGq+rrVrZI2k2n3q2Q+yfMhfcoC\ni9bi62tPeroDfrh4sy5EpkTL8ddIzy3Mz0EdRWpyu17lml+rUgceoLSjZVZVc0Cs\n5mre6Q8m5W0590Rd12yubLliq2E66cKom/wuIaGGy36+8zYbM1wiT8svylfzif4W\nAUSIGWRJAoGBAOwEMuxCoYvcJREG1R7QVc/egR8UrWlM6SXsbc5D96mHBT+DaaPw\nPXReFx9wIQ5J+uNRgTJ2GRfDS05t6i3kHtxDRGmz+qtaakAClmWcKjRBq6gdGg1j\nBcb86aVb7zr/x/x4jL9DQV5CJfC8Qs3Q966BbKzxP0zrPTPs0jLbmNC1AoGBAOhc\nXuZOqTeu4cSvWbcrvYZpoGIuhLCTV3x5OEPF847tEQ/U7m2gYQLpKU5oQPDIboCi\nYdSOs8tszefSW/438+VFPmUaClVVsEOU3hmkRjHPji6eX/pdDmWglZly4egoJNHM\nCtzDSiXU/uf9C1UaIjgnI0BW/jxdnC8Scn0eMXI7AoGAYfNwKuuQXho0a/eY9Zvj\nMU0X641KVvxreqi//a3pmDsIK47fhgFLkAMIt6xym5qKfFM0OnwW9+y+UUN+wCL/\nx8xHFVNzwi/ZDs3EG1GPlnZ2xmOlj068dBQ857ra83J6Tka9qxc/ht4PHvUTCJYk\nZREjyDarXcH3eZhcuGy45E0CgYEA1B222EHIwjdYjIeSu98nNbPpIJfcs6DeBZhX\nX68uZzRNFgnI2rTEuraE6bnMRsKB9dXGvxbCVRrvYQgsyIbllE/A5OR/uGTv8tHM\natYG6mPSJQaZEWGvyeBtkNZKGffDnO3KoLt6Tc0CIl9i3/5qbJ511L9VeV/vOx14\n/HT6qI8CgYA+BTF/aS9iQGrRI1LpMTQh5sbSnMz39Ri2VWulp/MLAOBCTKEpB9+W\neL+z52ZKn/kqo9hP/Ysl1Pt+JymPY1tP0TMJ4QCpVHnrmbx0xyNwZldR5kvgc/DN\neRzKmOYlhMF5LcqqdOaOOP4fbPAReq+FudAZuhrL4iReWnOR+I1Stw==\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttrSet(resourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate", "-----BEGIN CERTIFICATE-----\nMIIFZTCCBE2gAwIBAgISA6/vHTH2+uHo4BFuKgs90rQ2MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA3MDMxMzIwMTNaFw0y\nMDEwMDExMzIwMTNaMBsxGTAXBgNVBAMTEHd3dy5wYXRyaWNoaS5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWOPcmSEgjKK9o13/FYtJ5U4pzGTbE\n1EuaFPLC9KVFRcIN+16f72BdNjDGxm5b9CzfkfjcgKJpunCGWE7r6ukFch8JDKxn\nyNu6XDrgjhymm0cWn+UxaipzronT9A8RDUllXTD8UJzJjYi/x6/bsds5u0C03CIc\ni0ig5jwJrCKvKGUhsjX1vx8mKibjvETjmkLhZJ+50PYuDUW4MFFqjIKxwIKh+szm\ne2B+lCJQGYhiAY9R7R0J41M7tzYZj4IzuYh7mDt/+t+JOuz1OUenMwAFMauXS9yW\nL0oY7sNEW1cnC2QWH4EHQfyOdN0HkYbsy/0Eia7KzcbgK8zvipjRzbO3AgMBAAGj\nggJyMIICbjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsG\nAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFFHRKpBm95oY9EIO164b+cou\n6qHiMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUFBwEB\nBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZy8wKQYDVR0RBCIwIIIMcGF0cmljaGkuY29tghB3d3cucGF0cmljaGkuY29t\nMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUH\nAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB\n9ASB8QDvAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklTLe+LkF4AAAFzFQyZ\nUAAABAMARjBEAiBNyGhSkSA/QAV8O/z4p1sZxUbHX2LFaYhRn2gnRAmkRwIge9JC\nxsrSGMBtSOwN/mGBCr2yWsO7Xq0mFa7HvdyT5N0AdgBvU3asMfAxGdiZAKRRFf93\nFRwR2QLBACkGjbIImjfZEwAAAXMVDJl7AAAEAwBHMEUCICz3LqARn2byH91k3RZ/\nXloOaCPFI1hhzGk4LZbB7ZxOAiEA9CIbpEytmD0WctjD3HEsLHgpKtFCyFdOMrcS\n3W7+jh0wDQYJKoZIhvcNAQELBQADggEBAIBNvSrPD+01gTK64c6DkUix1xkdkzK6\nbMzjGU6lvcONhCSJGi4WZyl1m/nJfjuJ66xJSVlw0bcTbBmSYQCKWodkDPgR3HFM\nVJPXQuuLW7uNCMjqgz4h6o3WEWI1mKP5Vf5mPHVfgErgfk7TYtVWdqnB2/zJ72Tw\nvJwiSZjbEUBaxlw8BVdK224taStb1i1fM+xP+GOEaVlr6I59E1GpfOp1iwPBPzZi\nrb/w/03eCJn0P1+yTwBskvdsFRJaBbvd5mCGzAHcaAgMvbqv+u/7e01XgzPs8U7t\nHQ+a1q07HAmnobmXvkljO68T7MuUomFLmEX3RnRoNcXxqhNnMQnbQZ4=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "intermediate_certificates", "-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1jj3JkhIIyivaNd/xWLSeVOKcxk2xNRLmhTywvSlRUXCDfte\nn+9gXTYwxsZuW/Qs35H43ICiabpwhlhO6+rpBXIfCQysZ8jbulw64I4cpptHFp/l\nMWoqc66J0/QPEQ1JZV0w/FCcyY2Iv8ev27HbObtAtNwiHItIoOY8CawiryhlIbI1\n9b8fJiom47xE45pC4WSfudD2Lg1FuDBRaoyCscCCofrM5ntgfpQiUBmIYgGPUe0d\nCeNTO7c2GY+CM7mIe5g7f/rfiTrs9TlHpzMABTGrl0vcli9KGO7DRFtXJwtkFh+B\nB0H8jnTdB5GG7Mv9BImuys3G4CvM74qY0c2ztwIDAQABAoIBAQCgnYtor48ulUWX\nFOQeqW5nOyS2EXSH9ShN+WDenTEBFEFf3FVhEsgsewHtNz/tP9EZO0fcg7HCFeBi\nSfm6VqGpzJvKXC8zjVx5iMc4MDT5EbkmFHJyL0hu/bEnMnjNbynVjx64PvSfBbg/\nGkgXL23QBj6Im/gTcLbetGDVW7ORoRetRUxgOAGq+rrVrZI2k2n3q2Q+yfMhfcoC\ni9bi62tPeroDfrh4sy5EpkTL8ddIzy3Mz0EdRWpyu17lml+rUgceoLSjZVZVc0Cs\n5mre6Q8m5W0590Rd12yubLliq2E66cKom/wuIaGGy36+8zYbM1wiT8svylfzif4W\nAUSIGWRJAoGBAOwEMuxCoYvcJREG1R7QVc/egR8UrWlM6SXsbc5D96mHBT+DaaPw\nPXReFx9wIQ5J+uNRgTJ2GRfDS05t6i3kHtxDRGmz+qtaakAClmWcKjRBq6gdGg1j\nBcb86aVb7zr/x/x4jL9DQV5CJfC8Qs3Q966BbKzxP0zrPTPs0jLbmNC1AoGBAOhc\nXuZOqTeu4cSvWbcrvYZpoGIuhLCTV3x5OEPF847tEQ/U7m2gYQLpKU5oQPDIboCi\nYdSOs8tszefSW/438+VFPmUaClVVsEOU3hmkRjHPji6eX/pdDmWglZly4egoJNHM\nCtzDSiXU/uf9C1UaIjgnI0BW/jxdnC8Scn0eMXI7AoGAYfNwKuuQXho0a/eY9Zvj\nMU0X641KVvxreqi//a3pmDsIK47fhgFLkAMIt6xym5qKfFM0OnwW9+y+UUN+wCL/\nx8xHFVNzwi/ZDs3EG1GPlnZ2xmOlj068dBQ857ra83J6Tka9qxc/ht4PHvUTCJYk\nZREjyDarXcH3eZhcuGy45E0CgYEA1B222EHIwjdYjIeSu98nNbPpIJfcs6DeBZhX\nX68uZzRNFgnI2rTEuraE6bnMRsKB9dXGvxbCVRrvYQgsyIbllE/A5OR/uGTv8tHM\natYG6mPSJQaZEWGvyeBtkNZKGffDnO3KoLt6Tc0CIl9i3/5qbJ511L9VeV/vOx14\n/HT6qI8CgYA+BTF/aS9iQGrRI1LpMTQh5sbSnMz39Ri2VWulp/MLAOBCTKEpB9+W\neL+z52ZKn/kqo9hP/Ysl1Pt+JymPY1tP0TMJ4QCpVHnrmbx0xyNwZldR5kvgc/DN\neRzKmOYlhMF5LcqqdOaOOP4fbPAReq+FudAZuhrL4iReWnOR+I1Stw==\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttrSet(resourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_not_valid_after"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_apigateway_certificates", "test_certificates", Optional, Update, apiGatewaycertificateDataSourceRepresentation) +
					compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Optional, Update, apiGatewaycertificateRepresentation) +
					generateDataSourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApiGatewayCertificateResourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate", "-----BEGIN CERTIFICATE-----\nMIIFZTCCBE2gAwIBAgISA6/vHTH2+uHo4BFuKgs90rQ2MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA3MDMxMzIwMTNaFw0y\nMDEwMDExMzIwMTNaMBsxGTAXBgNVBAMTEHd3dy5wYXRyaWNoaS5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWOPcmSEgjKK9o13/FYtJ5U4pzGTbE\n1EuaFPLC9KVFRcIN+16f72BdNjDGxm5b9CzfkfjcgKJpunCGWE7r6ukFch8JDKxn\nyNu6XDrgjhymm0cWn+UxaipzronT9A8RDUllXTD8UJzJjYi/x6/bsds5u0C03CIc\ni0ig5jwJrCKvKGUhsjX1vx8mKibjvETjmkLhZJ+50PYuDUW4MFFqjIKxwIKh+szm\ne2B+lCJQGYhiAY9R7R0J41M7tzYZj4IzuYh7mDt/+t+JOuz1OUenMwAFMauXS9yW\nL0oY7sNEW1cnC2QWH4EHQfyOdN0HkYbsy/0Eia7KzcbgK8zvipjRzbO3AgMBAAGj\nggJyMIICbjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsG\nAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFFHRKpBm95oY9EIO164b+cou\n6qHiMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUFBwEB\nBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZy8wKQYDVR0RBCIwIIIMcGF0cmljaGkuY29tghB3d3cucGF0cmljaGkuY29t\nMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUH\nAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB\n9ASB8QDvAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklTLe+LkF4AAAFzFQyZ\nUAAABAMARjBEAiBNyGhSkSA/QAV8O/z4p1sZxUbHX2LFaYhRn2gnRAmkRwIge9JC\nxsrSGMBtSOwN/mGBCr2yWsO7Xq0mFa7HvdyT5N0AdgBvU3asMfAxGdiZAKRRFf93\nFRwR2QLBACkGjbIImjfZEwAAAXMVDJl7AAAEAwBHMEUCICz3LqARn2byH91k3RZ/\nXloOaCPFI1hhzGk4LZbB7ZxOAiEA9CIbpEytmD0WctjD3HEsLHgpKtFCyFdOMrcS\n3W7+jh0wDQYJKoZIhvcNAQELBQADggEBAIBNvSrPD+01gTK64c6DkUix1xkdkzK6\nbMzjGU6lvcONhCSJGi4WZyl1m/nJfjuJ66xJSVlw0bcTbBmSYQCKWodkDPgR3HFM\nVJPXQuuLW7uNCMjqgz4h6o3WEWI1mKP5Vf5mPHVfgErgfk7TYtVWdqnB2/zJ72Tw\nvJwiSZjbEUBaxlw8BVdK224taStb1i1fM+xP+GOEaVlr6I59E1GpfOp1iwPBPzZi\nrb/w/03eCJn0P1+yTwBskvdsFRJaBbvd5mCGzAHcaAgMvbqv+u/7e01XgzPs8U7t\nHQ+a1q07HAmnobmXvkljO68T7MuUomFLmEX3RnRoNcXxqhNnMQnbQZ4=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "intermediate_certificates", "-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subject_names.0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_not_valid_after"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CertificateResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"private_key",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckApigatewayCertificateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).apiGatewayClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_certificate" {
			noResourceFound = false
			request := oci_apigateway.GetCertificateRequest{}

			tmp := rs.Primary.ID
			request.CertificateId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")

			response, err := client.GetCertificate(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apigateway.CertificateLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ApigatewayCertificate") {
		resource.AddTestSweepers("ApigatewayCertificate", &resource.Sweeper{
			Name:         "ApigatewayCertificate",
			Dependencies: DependencyGraph["certificate"],
			F:            sweepApigatewayCertificateResource,
		})
	}
}

func sweepApigatewayCertificateResource(compartment string) error {
	apiGatewayClient := GetTestClients(&schema.ResourceData{}).apiGatewayClient()
	certificateIds, err := getApiGatewayCertificateIds(compartment)
	if err != nil {
		return err
	}
	for _, certificateId := range certificateIds {
		if ok := SweeperDefaultResourceId[certificateId]; !ok {
			deleteCertificateRequest := oci_apigateway.DeleteCertificateRequest{}

			deleteCertificateRequest.CertificateId = &certificateId

			deleteCertificateRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "apigateway")
			_, error := apiGatewayClient.DeleteCertificate(context.Background(), deleteCertificateRequest)
			if error != nil {
				fmt.Printf("Error deleting Certificate %s %s, It is possible that the resource is already deleted. Please verify manually \n", certificateId, error)
				continue
			}
			waitTillCondition(testAccProvider, &certificateId, apiGatewayCertificateSweepWaitCondition, time.Duration(3*time.Minute),
				apiGatewayCertificateSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getApiGatewayCertificateIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CertificateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	apiGatewayClient := GetTestClients(&schema.ResourceData{}).apiGatewayClient()

	listCertificatesRequest := oci_apigateway.ListCertificatesRequest{}
	listCertificatesRequest.CompartmentId = &compartmentId
	listCertificatesRequest.LifecycleState = oci_apigateway.CertificateLifecycleStateActive
	listCertificatesResponse, err := apiGatewayClient.ListCertificates(context.Background(), listCertificatesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Certificate list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, certificate := range listCertificatesResponse.Items {
		id := *certificate.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CertificateId", id)
	}
	return resourceIds, nil
}

func apiGatewayCertificateSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if certificateResponse, ok := response.Response.(oci_apigateway.GetCertificateResponse); ok {
		return certificateResponse.LifecycleState != oci_apigateway.CertificateLifecycleStateDeleted
	}
	return false
}

func apiGatewayCertificateSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.apiGatewayClient().GetCertificate(context.Background(), oci_apigateway.GetCertificateRequest{
		CertificateId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
