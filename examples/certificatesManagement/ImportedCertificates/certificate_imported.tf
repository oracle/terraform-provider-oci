// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


provider "oci" {
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region = "r1"
}

locals {
  imported_certificate_pem = <<-EOT
  -----BEGIN CERTIFICATE-----
  MIID1DCCArygAwIBAgIUUSNkzHmCRJWIZeCu/+RNSolSWgYwDQYJKoZIhvcNAQEL
  BQAwSTEkMCIGA1UEAwwbQ29kZXggSW1wb3J0ZWQgQ2VydCBSb290IENBMRQwEgYD
  VQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwHhcNMjYwNDI5MjAxNjU4WhcN
  MjgwODAxMjAxNjU4WjBBMRwwGgYDVQQDDBNpbXBvcnRlZC1jZXJ0LmxvY2FsMRQw
  EgYDVQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwggEiMA0GCSqGSIb3DQEB
  AQUAA4IBDwAwggEKAoIBAQCuc7mFa9Zo29aQBg3FwPQXofPDaucPgdqngvYJXJFY
  wgTPidgZUgiRi1npOY8PCDm7dn5B/vQuLK/GMrI3LqyOvzYNo36ilDHkt1zit8MP
  Amxw2LBUg+6WOHTOjD2U7GogTBTXu8NhKhgdlehx6u6Nd+HvrPEaD9rb2jOZUcBZ
  mi1zmctZTd+4kZqWJx42ktHzRLrWaRltu7gvr/cIjQ879tKQOnCncxFbLJyFUtGB
  bTO1/2kL83l69GQGfb7qdJVy0r6MxlOYyNnI+kbclyagyhbTyPWUnn8Wy2ztvBYs
  CoRc165Al8n4prGFGvtCapgGRRpOIMatHUT4JL+6q6fHAgMBAAGjgbswgbgwOQYD
  VR0RBDIwMIITaW1wb3J0ZWQtY2VydC5sb2NhbIIZaW1wb3J0ZWQtY2VydC5leGFt
  cGxlLmNvbTAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAU
  BggrBgEFBQcDAQYIKwYBBQUHAwIwHQYDVR0OBBYEFLBPB1H6lhA/vh9muRT88YBZ
  jk6qMB8GA1UdIwQYMBaAFNs69dZIK4qugHpj93QD4Gv3p/yaMA0GCSqGSIb3DQEB
  CwUAA4IBAQDGidVB+XQDB1dvjJ1v+kxKFjmqHH//9MAq+EYVrPXpNnUk03OfPuPv
  oj8ybEwIU8lMX+acDFMnJaDGMzCOce2bCgmvGSNboUGpRE1jzYhPozQK1MjHt111
  QUe7nOG/A+/qmdnYjNg1ZTRjaxoufGYJQSdFowhAGxyi3rtXVXeCVgsYPKljURI1
  cKmcp1YPIz9oaFZdGTNCUvxY113Y/GHsArROfI48fZTzxUIvZ+3zCa3LFqQeRLe5
  OmBV10uveTHS5e6GeHol1yEtpxFEhs0w8JtjxdAjKaBrnB88s4rvsHIyXmkvW9L3
  zriS1Nol4DlLdoH9POvN3y4PpZ6asH4y
  -----END CERTIFICATE-----
  EOT

  imported_private_key_pem = <<-EOT
  -----BEGIN PRIVATE KEY-----
  MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCuc7mFa9Zo29aQ
  Bg3FwPQXofPDaucPgdqngvYJXJFYwgTPidgZUgiRi1npOY8PCDm7dn5B/vQuLK/G
  MrI3LqyOvzYNo36ilDHkt1zit8MPAmxw2LBUg+6WOHTOjD2U7GogTBTXu8NhKhgd
  lehx6u6Nd+HvrPEaD9rb2jOZUcBZmi1zmctZTd+4kZqWJx42ktHzRLrWaRltu7gv
  r/cIjQ879tKQOnCncxFbLJyFUtGBbTO1/2kL83l69GQGfb7qdJVy0r6MxlOYyNnI
  +kbclyagyhbTyPWUnn8Wy2ztvBYsCoRc165Al8n4prGFGvtCapgGRRpOIMatHUT4
  JL+6q6fHAgMBAAECggEABiQWvMaeBpTuiaVc6MAxF3/g99kUhyw4CTuH865pg2lZ
  CAE/NSz5Fj6EjLw09/9/IMU8DpCuAc4a5ew079JZyrCac4J0275Bv15hC64SVL3u
  GU7XbjIyvOEB4592L37fAYIMHcgDSbFUpf3378KoYRx+PJSm7Bbwylp1BImnLsA6
  npTSC/gfJzPtpgKlPuADcn2eHlZkqxF1GEdzNbee69sEf7UoY54y5Z7CaaObRB55
  jB4qmSHD5Zk3Le7h7oXFKp39BFRPiKJlnBPmZ10Ca2pdPH4wXDpNj7qm5gG/qwWN
  SeaaXvhUbJ+Xtj4LJ04CrRLjXYL5z2NSVnZuaLiCgQKBgQD0OF/X8mWj1/OzIz3H
  /Van334ef0VMDMUcW4T+vGdL/4lnYHY85ZNleh/f0P9XnypcjwQwStNPWwiq+83w
  iJC0Gmsf71eUyDPIJI2rOp3COdVeWo6+NfDp2iNoXYksneGbL4pETn1fuiCFdgO+
  MLPqAvvx243nR0+FZBkMVclnQQKBgQC23drIXsm0WeTwcaiA/7m+KWg4soUCJc3d
  qUY2pGKAYZunBzDT6DPy6vbDGKNFib0NPPwHrI52UYKpfWta0jixecndm9e8KF1W
  Kr+OuzqNvt2RNQv1oaDYXkFqPx8V20HYCbGykLUrcTvl8kQ/P/fuJGb1lJpFjrxb
  Zj37ndGVBwKBgQDfWxe9BoISwRSvB1IfsAEq82oDNl8IWL2FW+Zcm67VgNYH5977
  fvO5YWH5bsm5N2ak1F/+UtDDyPw6dKU6kYvWTDshL+Knvhott9Chg7B4zC/qZWQb
  bnXxSoILl1A3sAV/ypDoO+jKPg40QWTW8u/429XUsvorRbriwlNuRp82gQKBgBlp
  POTJHzhboY+0C5lZb+M3986hKBCDVlcuFZ60Oyfh6HJOcn+P7eAcvIuh2dZo7Idx
  aQhb9zQD4adR0WA+hDhNVzrfVBxGc6nsqM/ZvqpB0qJB/9ZzTlmix25PNS+hV0YB
  GttYdoSB6LPsdYVgi0BheJk5vDJcQlplG8niRWRzAoGAIvTSa/DG8XR4EsEUFqtX
  POFZD+rtbo3billYf5JenBtAuhUw0Ub1Xs5A+z+zpMbEjlv26qFH+5JUhc+bFxit
  tCbAF7RXkhU1c5MTy+R4dAq8ZDfy44eZbNnSycfxhHku8zP7a3rwfoqxQOqGHsML
  ERgQfPPismBqinuoMVUCHLU=
  -----END PRIVATE KEY-----
  EOT

  imported_cert_chain_pem = <<-EOT
  -----BEGIN CERTIFICATE-----
  MIIDgzCCAmugAwIBAgIUPkRQkm0NaovEUyojIoqkYWHtYQswDQYJKoZIhvcNAQEL
  BQAwSTEkMCIGA1UEAwwbQ29kZXggSW1wb3J0ZWQgQ2VydCBSb290IENBMRQwEgYD
  VQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwHhcNMjYwNDI5MjAxNjU4WhcN
  MzYwNDI2MjAxNjU4WjBJMSQwIgYDVQQDDBtDb2RleCBJbXBvcnRlZCBDZXJ0IFJv
  b3QgQ0ExFDASBgNVBAoMC09yYWNsZSBUZXN0MQswCQYDVQQGEwJVUzCCASIwDQYJ
  KoZIhvcNAQEBBQADggEPADCCAQoCggEBAM5yng/WTdPusublw9tqzpfqN5D6AAZR
  XelergeDwsh3xRul0mqwE26S9SwmRZO0GnJJ4Gmf/Wnoj/2+LGgGs4Z5/qjgg6FG
  4Uy6QzDTapbGldqzoKGEFXvJ4AQgwpotuosp7Xtm9OG0LMU8V+X/OePBEfixQaln
  n0fXifVZDHQDQZOTF0XJAnDnu8oe8q9U35DHRjwoghAcjk4lN9E1025FFEVvrFTR
  jK0H6Y5W/ZJNIMUk9XPRpatLQphLrQAKRF3mH0xMHxJttUPrCQSRYyBj9BXK4aIv
  5Z3+AQP40mceptckAi91vyCzEfISJPr3KZltpIbRYc2kEPi55pZSRWkCAwEAAaNj
  MGEwHQYDVR0OBBYEFNs69dZIK4qugHpj93QD4Gv3p/yaMB8GA1UdIwQYMBaAFNs6
  9dZIK4qugHpj93QD4Gv3p/yaMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQD
  AgGGMA0GCSqGSIb3DQEBCwUAA4IBAQCBuL130bVAtzwKbS5IAqM7NlqrJazVOOhi
  nNQDvY1C8CrRqZJIdMyTNURIiILXrdGlAjX/7xD3NLR5Ea5Ka6QGN8tVsTrtLdYw
  32rGrsUQ38FbEj8dD4N5wPDhrqsYw/lDn2yka2VjTIDBnwVB7r/g9TOyElPIoCyS
  oV1HruBPRqeF78GvUxdIJky15JM6LSAVR+7nqbcMZOEFyxMRdCXztyY5jrQmZJug
  TFgSRG/Q92we69f7VsZ36+b+qztokiJkoCdfLE6irdWdVlu7KeGK0vRnenXWE/Jx
  L/T/I7wj7Ol+eF5VGVybNyEEZL1TZjIWzbg8rtO8RcJqcdt6vRpM
  -----END CERTIFICATE-----
  EOT
}

resource "oci_certificates_management_certificate" "imported_certificate" {
  compartment_id = "ocid1.tenancy.region1..aaaaaaaalt555yh2rr3axqttiylafzm3rsbczigzko33wjeaj5t63ld3bl4q"
  name           = "imported-cert"
  description    = "Imported certificate smoke test for local provider validation"

  certificate_config {
    config_type     = "IMPORTED"
    certificate_pem = trimspace(local.imported_certificate_pem)
    private_key_pem = trimspace(local.imported_private_key_pem)
    cert_chain_pem  = trimspace(local.imported_cert_chain_pem)
    version_name    = "imported-cert-v1"
  }
}

output "imported_certificate_id" {
  value = oci_certificates_management_certificate.imported_certificate.id
}