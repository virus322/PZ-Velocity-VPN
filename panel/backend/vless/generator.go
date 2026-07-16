package vless

import (
	"fmt"
	"net/url"
)


func Generate(c Config) string {


	params := url.Values{}

	params.Set(
		"encryption",
		"none",
	)

	params.Set(
		"security",
		"reality",
	)

	params.Set(
		"sni",
		c.SNI,
	)

	params.Set(
		"fp",
		c.Fingerprint,
	)


	link := fmt.Sprintf(

		"vless://%s@%s:%d?%s#%s",

		c.UUID,
		c.Address,
		c.Port,
		params.Encode(),
		c.Remark,

	)


	return link

}