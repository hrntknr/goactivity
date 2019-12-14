package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hrntknr/goactivity/activity"
	"gopkg.in/yaml.v2"
)

func main() {
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	config := &Config{}
	if err := yaml.Unmarshal(buf, config); err != nil {
		log.Fatal(err)
	}

	if err := startServer(config); err != nil {
		log.Fatal(err)
	}
}

func startServer(config *Config) error {
	r := gin.Default()
	r.GET("/users/hrntknr", func(c *gin.Context) {
		act := activity.NewActivity()
		as := activity.NewActivityStream(act)
		sec := activity.NewSecurity(act)
		as.ID = "https://hirano.work/users/hrntknr"
		as.Type = "Person"
		as.Following = "https://hirano.work/users/hrntknr/following"
		as.Followers = "https://hirano.work/users/hrntknr/followers"
		as.Inbox = "https://hirano.work/users/hrntknr/inbox"
		as.Outbox = "https://hirano.work/users/hrntknr/outbox"
		as.PreferredUsername = "hrntknr"
		as.Name = "Takanori Hirano"
		as.Summary = "<p></p>"
		as.URL = "https://hirano.work/@hrntknr"
		as.Endpoints.SharedInbox = "https://hirano.work/inbox"
		as.Icon.Type = "Image"
		as.Icon.MediaType = "image/png"
		as.Icon.URL = "https://avatars1.githubusercontent.com/u/24761092"
		sec.PublicKey.ID = "https://hirano.work/users/hrntknr#main-key"
		sec.PublicKey.Owner = "https://hirano.work/users/hrntknr"
		sec.PublicKey.PublicKeyPem = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsZI0PrPFktz8X3/DZPBJ\nLMNchKVxXRFIMAFv507848V0pf2dBiVMDJ+ip46fIQvU9JnK0wAjvl1wyu0uuRmH\n+IgRqhwx79N+K+3S9MRkfLxQypXM2H/UU+y9joRKGDk4wSQP6+mUgR8fx+b7YFvY\nV+7mSTBemyX55uZfaAXo3iYwziPCIfv2g5Z4x982yO0up+2lkt9F79nl7k4dzuLk\n8qRRZpK9uUfAK50RQTeSh3xsUEu+DUteKk7mGPGJaWanBfSlB0uF3nuaFeugt4sB\nlli8FyFCS+hQ2U4ZOlHf3x+hkppph0bKukYrraJs46IB9Zgg3n5x2M/ZIvmbzRBE\nIQIDAQAB\n-----END PUBLIC KEY-----\n"

		c.Header("Content-Type", "application/activity+json")
		c.JSON(200, act)
		return
	})
	r.GET("/.well-known/webfinger", func(c *gin.Context) {
		resource := c.Query("resource")
		sResource := strings.Split(resource, ":")
		if len(sResource) != 2 {
			c.Status(400)
			return
		}
		switch sResource[0] {
		case "acct":
			sAct := strings.Split(sResource[1], "@")
			if sAct[1] != config.Server.Host {
				c.Status(404)
				return
			}
			c.Header("Content-Type", "application/jrd+json")
			c.JSON(200, gin.H{
				"subject": resource,
				"links": []map[string]string{
					{
						"rel":  "self",
						"type": "application/activity+json",
						"href": fmt.Sprintf("https://%s/users/%s", config.Server.Host, sAct[0]),
					},
				},
			})
			return
		}
		c.Status(404)
		return
	})
	return r.Run()
}
