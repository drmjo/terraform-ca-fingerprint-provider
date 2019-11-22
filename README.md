# Root CA Fingerprint Generator

This terraform module will generate a fingerprint of the root certificate of a server

# Usage

```
resource "finger_sha1" "mjo_io" {
    domain = "mjo.io"
}

output "fingerprint" {
  value = "${finger_sha1.mjo_io.id}"
}
```