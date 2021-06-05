package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBgYGBITjkaxIAE2Bg4GQpKk3Iyk0NDWBkYq/N3JwR4s3MgKwnwZmQSYcatHQIEGP47gmgMw1jZQMKMDIwMtgwMYMzAAAgAAP//ptW1E5AAAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
