package finger

import (
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func fingerSha1() *schema.Resource {
	return &schema.Resource{
		Create: CreateSha1,

		Read:   CreateSha1,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "443",
				ForceNew: true,
			},
		},
	}
}

func CreateSha1(d *schema.ResourceData, _ interface{}) error {
	domain := d.Get("domain")
	port := d.Get("port")

	conf := &tls.Config{}

	addr := fmt.Sprintf("%s:%s", domain, port)

	conn, err := tls.Dial("tcp", addr, conf)
	if err != nil {
		return err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates

	// the root CA is the last Cert
	rootCA := certs[len(certs)-1]

	sum := sha1.Sum(rootCA.Raw)
	hex := hex.EncodeToString(sum[:])

	d.SetId(strings.ToUpper(hex))

	return nil
}
