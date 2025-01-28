variable "byodkim_name" {
  default = "test-selector-2"
}

variable "dkim_name" {
  default = "test-selector-1"
}

variable "dkim_description" {
  default = "Test DKIM"
}

variable "private_key" {
  default = "MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDYOJHqR6O6TofL3eUshL+yDN9sTDSDCmSMe5hapcu1Cj6zjXLbPvEUM+SZd2MrBaAfh5HaiM0U3hQgPX//6zAewT95n5K94qnCo1yKBYpt81lnY5PRcePOE0wuFQ8EqqJNbIPFojYLpoyKazB7ulCiA1y6Zk5ZzUolxNslnRwjWtNidC62ureDbDsE28LVhHron7+s5StR5cMTRUxU1nU3KeOluUS2YwJQzvdAQiDn9PHnG0k7p9+Cam0mlNPgPUKIKmFigx+WBrrI5k1yxyYVjbwTQgl7J7bEz0kZZJLN0Tb1UVfIe+yTDz9iZgMDZ0+spo9FA4UEpQiHsr3Cih6XAgMBAAECggEADtEZlq9gFeVJwvuqt7lFFaNyMA6szooIx+O0Vgnkgrx3/X/k7gdLJ/JCaQ3YWttg1KhKo2oae6KYSHlOazAntQPTILeBq8d7kJ33OnIUwWIar04QlBq03KtP3hyNUemLg7i68vEhuPypxtZ/lQr4eZG3agKWveqmeR4bUmGIyxdO6GaXQMoGo8javYylS02X3uW/jTT4cXWcnslBibQ48CeGS7fuxGHHoZgU0qq3JDlO2KZtOjnmMvwqhAlnCgcd0EX1exsddWXhi2R7ONMRXT/0XG16AVFUumqU5c1E3qNpB+/OR7GO1eIxXsIIN/lBBLblIuETV0naoEeIR0cC4QKBgQD/bEe8cqJlICDDBucNLviTXU2xWHbg4rik1l+NxcN5oRuF0klEklCKPeoiruwFuOxjR4v+E5E6OY/X3lEdVbm7sQpX5rU5Ql9Mkuikniq0gji3aG1fgp1S+rltie1IwIvWWJII5SKrl5zk+K/1F2fiYIeS3xDpBXZAEuYkmDlXcQKBgQDYtZ4v1gGqz2avqSurVgIBKuSy5lQW4RCctGHvaaMMqmE5GipCId/2toIxCm/TTnkw2xy5iLN+TRu1ejQlnXyjbv5rWpg6Ud4ZfS+dspBZKQDndcUxBb8AOEEa3mfNnqtf2fpLNiBHHQUQbQ0n+Mvlk0P9RCIF8cakw8vTUbAihwKBgQDRovB5xehtcWIST68BTgTtBk5G+WGQj+SWPfBtLtySDl3FhuNlqWT7IXfygXjr9WvmzAklPrAahrX9qdUPa5Wma9cecBdbaJNJCOZyffJ09tBQECB8nHo9adUZB3Q+GvJquYUaUd4T0bsO3uGkGlLEJRAAxp5dcoMscZZ8gIfHsQKBgQCwlu6cVGtvYSpwIg5vw9pGUUmjboU/T49V4rebfB1diMlI2sVstmXmiLp2c+bD7cJbyXESzw5U8UiPDVHJBRg3pY60tFIv096ELrrKamOuA/e0pQdW9zO5Xh3PY0R7DbzgkqOe/jA3QrKJvBNoaxjbrrFMhGSlK2uPf+1r+bjpqQKBgQDvSKV73DAKgi0SWreJsBxzJ5qXTj5vcFfu4ySYzoCabfyH6sSokhRUHrzT0V3I+oF0N84IqoDX4CxnNmmu4DvZx30geMSlkFEzY2hPktEcO3nyEgBEFd8PVIp3NqL9u3N8rfQIfy8kudoeZWg3+9sjIFkEiX/HghFf4tt7qcGRwg=="
}

resource "oci_email_dkim" "testbyodkim" {
	#Required
	email_domain_id = oci_email_email_domain.test_email_domain.id
	#Optional
	description = var.dkim_description
	name = var.byodkim_name
        private_key = var.private_key
  depends_on = [oci_email_email_domain.test_email_domain]
}

resource "oci_email_dkim" "testdkim" {
        #Required
        email_domain_id = oci_email_email_domain.test_email_domain.id
        #Optional
        description = var.dkim_description
        name = var.dkim_name
  depends_on = [oci_email_email_domain.test_email_domain]
}

