usage

```
resource "finger_sha1" "mjo_io" {
    domain = "mjo.io"
}

output "fingerprint" {
  value = "${finger_sha1.mjo_io.id}"
}
```