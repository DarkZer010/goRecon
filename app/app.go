package app

import (
    "github.com/urfave/cli"
    
    "net"
    "log"
    "fmt"
)

// Generating will return the command-line application ready to be executed.
func Gerar() *cli.App {
        app := cli.NewApp()
        app.Name = "gorecon"
        app.Usage = "Search for IP addresses and web server names."	

        flags := []cli.Flag{
        	cli.StringFlag{
        		Name: "host",
        		Value: "devbook.com.br",
        		},
        	}
        

        app.Commands = []cli.Command{
        	{
        		Name: "ip",
        		Usage: "Find the server's IP address.",
        		Flags: flags,
        	    Action: searchIPs,
        	},
        	{
        		Name: "servername",
        		Usage: "Search for the name of the service.",
        		Flags: flags,
        		Action: searchServersName,
        	},
        }
        
        return app
}

func searchIPs(c *cli.Context) {
    fmt.Println("----------IPs------------")
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	
	fmt.Println("-------------------------\n")
}

func searchServersName(c *cli.Context) {
    fmt.Println("----------Host NS------------")
    host := c.String("host")

    servers, erro := net.LookupNS(host)
    if erro != nil {
    	log.Fatal(erro)
    }

    for _, server := range servers {
    	fmt.Println(server)
    }

    fmt.Println("-------------------------------\n")
}
