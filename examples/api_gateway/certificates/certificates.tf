// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "certificate_certificate" {
  default = "-----BEGIN CERTIFICATE-----\nMIIFZTCCBE2gAwIBAgISA6/vHTH2+uHo4BFuKgs90rQ2MA0GCSqGSIb3DQEBCwUA\nMEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD\nExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA3MDMxMzIwMTNaFw0y\nMDEwMDExMzIwMTNaMBsxGTAXBgNVBAMTEHd3dy5wYXRyaWNoaS5jb20wggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDWOPcmSEgjKK9o13/FYtJ5U4pzGTbE\n1EuaFPLC9KVFRcIN+16f72BdNjDGxm5b9CzfkfjcgKJpunCGWE7r6ukFch8JDKxn\nyNu6XDrgjhymm0cWn+UxaipzronT9A8RDUllXTD8UJzJjYi/x6/bsds5u0C03CIc\ni0ig5jwJrCKvKGUhsjX1vx8mKibjvETjmkLhZJ+50PYuDUW4MFFqjIKxwIKh+szm\ne2B+lCJQGYhiAY9R7R0J41M7tzYZj4IzuYh7mDt/+t+JOuz1OUenMwAFMauXS9yW\nL0oY7sNEW1cnC2QWH4EHQfyOdN0HkYbsy/0Eia7KzcbgK8zvipjRzbO3AgMBAAGj\nggJyMIICbjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsG\nAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFFHRKpBm95oY9EIO164b+cou\n6qHiMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUFBwEB\nBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNyeXB0\nLm9yZy8wKQYDVR0RBCIwIIIMcGF0cmljaGkuY29tghB3d3cucGF0cmljaGkuY29t\nMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUH\nAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB\n9ASB8QDvAHUAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklTLe+LkF4AAAFzFQyZ\nUAAABAMARjBEAiBNyGhSkSA/QAV8O/z4p1sZxUbHX2LFaYhRn2gnRAmkRwIge9JC\nxsrSGMBtSOwN/mGBCr2yWsO7Xq0mFa7HvdyT5N0AdgBvU3asMfAxGdiZAKRRFf93\nFRwR2QLBACkGjbIImjfZEwAAAXMVDJl7AAAEAwBHMEUCICz3LqARn2byH91k3RZ/\nXloOaCPFI1hhzGk4LZbB7ZxOAiEA9CIbpEytmD0WctjD3HEsLHgpKtFCyFdOMrcS\n3W7+jh0wDQYJKoZIhvcNAQELBQADggEBAIBNvSrPD+01gTK64c6DkUix1xkdkzK6\nbMzjGU6lvcONhCSJGi4WZyl1m/nJfjuJ66xJSVlw0bcTbBmSYQCKWodkDPgR3HFM\nVJPXQuuLW7uNCMjqgz4h6o3WEWI1mKP5Vf5mPHVfgErgfk7TYtVWdqnB2/zJ72Tw\nvJwiSZjbEUBaxlw8BVdK224taStb1i1fM+xP+GOEaVlr6I59E1GpfOp1iwPBPzZi\nrb/w/03eCJn0P1+yTwBskvdsFRJaBbvd5mCGzAHcaAgMvbqv+u/7e01XgzPs8U7t\nHQ+a1q07HAmnobmXvkljO68T7MuUomFLmEX3RnRoNcXxqhNnMQnbQZ4=\n-----END CERTIFICATE-----"
}

variable "certificate_defined_tags_value" {
  default = "value"
}

variable "certificate_display_name" {
  default = "displayName"
}

variable "certificate_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "certificate_intermediate_certificates" {
  default = "-----BEGIN CERTIFICATE-----\nMIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/\nMSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT\nDkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow\nSjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT\nGkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC\nAQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF\nq6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8\nSMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0\nZ8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA\na6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj\n/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T\nAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG\nCCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv\nbTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k\nc3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw\nVAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC\nARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz\nMDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu\nY3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF\nAAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo\nuM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/\nwApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu\nX4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG\nPfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6\nKOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==\n-----END CERTIFICATE-----"
}

variable "certificate_private_key" {
  default = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA1jj3JkhIIyivaNd/xWLSeVOKcxk2xNRLmhTywvSlRUXCDfte\nn+9gXTYwxsZuW/Qs35H43ICiabpwhlhO6+rpBXIfCQysZ8jbulw64I4cpptHFp/l\nMWoqc66J0/QPEQ1JZV0w/FCcyY2Iv8ev27HbObtAtNwiHItIoOY8CawiryhlIbI1\n9b8fJiom47xE45pC4WSfudD2Lg1FuDBRaoyCscCCofrM5ntgfpQiUBmIYgGPUe0d\nCeNTO7c2GY+CM7mIe5g7f/rfiTrs9TlHpzMABTGrl0vcli9KGO7DRFtXJwtkFh+B\nB0H8jnTdB5GG7Mv9BImuys3G4CvM74qY0c2ztwIDAQABAoIBAQCgnYtor48ulUWX\nFOQeqW5nOyS2EXSH9ShN+WDenTEBFEFf3FVhEsgsewHtNz/tP9EZO0fcg7HCFeBi\nSfm6VqGpzJvKXC8zjVx5iMc4MDT5EbkmFHJyL0hu/bEnMnjNbynVjx64PvSfBbg/\nGkgXL23QBj6Im/gTcLbetGDVW7ORoRetRUxgOAGq+rrVrZI2k2n3q2Q+yfMhfcoC\ni9bi62tPeroDfrh4sy5EpkTL8ddIzy3Mz0EdRWpyu17lml+rUgceoLSjZVZVc0Cs\n5mre6Q8m5W0590Rd12yubLliq2E66cKom/wuIaGGy36+8zYbM1wiT8svylfzif4W\nAUSIGWRJAoGBAOwEMuxCoYvcJREG1R7QVc/egR8UrWlM6SXsbc5D96mHBT+DaaPw\nPXReFx9wIQ5J+uNRgTJ2GRfDS05t6i3kHtxDRGmz+qtaakAClmWcKjRBq6gdGg1j\nBcb86aVb7zr/x/x4jL9DQV5CJfC8Qs3Q966BbKzxP0zrPTPs0jLbmNC1AoGBAOhc\nXuZOqTeu4cSvWbcrvYZpoGIuhLCTV3x5OEPF847tEQ/U7m2gYQLpKU5oQPDIboCi\nYdSOs8tszefSW/438+VFPmUaClVVsEOU3hmkRjHPji6eX/pdDmWglZly4egoJNHM\nCtzDSiXU/uf9C1UaIjgnI0BW/jxdnC8Scn0eMXI7AoGAYfNwKuuQXho0a/eY9Zvj\nMU0X641KVvxreqi//a3pmDsIK47fhgFLkAMIt6xym5qKfFM0OnwW9+y+UUN+wCL/\nx8xHFVNzwi/ZDs3EG1GPlnZ2xmOlj068dBQ857ra83J6Tka9qxc/ht4PHvUTCJYk\nZREjyDarXcH3eZhcuGy45E0CgYEA1B222EHIwjdYjIeSu98nNbPpIJfcs6DeBZhX\nX68uZzRNFgnI2rTEuraE6bnMRsKB9dXGvxbCVRrvYQgsyIbllE/A5OR/uGTv8tHM\natYG6mPSJQaZEWGvyeBtkNZKGffDnO3KoLt6Tc0CIl9i3/5qbJ511L9VeV/vOx14\n/HT6qI8CgYA+BTF/aS9iQGrRI1LpMTQh5sbSnMz39Ri2VWulp/MLAOBCTKEpB9+W\neL+z52ZKn/kqo9hP/Ysl1Pt+JymPY1tP0TMJ4QCpVHnrmbx0xyNwZldR5kvgc/DN\neRzKmOYlhMF5LcqqdOaOOP4fbPAReq+FudAZuhrL4iReWnOR+I1Stw==\n-----END RSA PRIVATE KEY-----"
}

variable "certificate_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_apigateway_certificate" "test_certificate" {
  #Required
  certificate    = "${var.certificate_certificate}"
  compartment_id = "${var.compartment_id}"
  private_key    = "${var.certificate_private_key}"

  #Optional
  defined_tags              = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.certificate_defined_tags_value}")}"
  display_name              = "${var.certificate_display_name}"
  freeform_tags             = "${var.certificate_freeform_tags}"
  intermediate_certificates = "${var.certificate_intermediate_certificates}"
}

data "oci_apigateway_certificates" "test_certificates" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  display_name = "${var.certificate_display_name}"
  state        = "${var.certificate_state}"
}
