# Root CA Fingerprint Generator

This terraform module will generate a fingerprint of the root certificate of a server

# Install
```
mkdir -p ~/.terraform.d/plugins

curl -L \
  -o ~/.terraform.d/plugins/terraform-provider-finger \
  https://github.com/drmjo/terraform-provider-finger/releases/download/0.0.1/terraform-provider-finger.v0.0.1.linux

chmod +x ~/.terraform.d/plugins/terraform-provider-finger
```

# Usage

```
resource "finger_sha1" "mjo_io" {
    domain = "mjo.io"
}

output "fingerprint" {
  value = "${finger_sha1.mjo_io.id}"
}
```