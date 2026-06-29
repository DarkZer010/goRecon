package app

import (
    "net"
    "fmt"

    "github.com/urfave/cli"
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
        		Name: "ns",
        		Usage: "Search for the name of the service.",
        		Flags: flags,
        		Action: searchServersName,
        	},
        	{
        		Name: "cnm",
        		Usage: "Search for the canonical name record",
        		Flags: flags,
        		Action: searchCNAME,
        	},
        	{
        		Name: "txt",
        		Usage: "Search for TXT records",
        		Flags: flags,
        		Action: searchTXT,
        	},
        }
        
        return app
}

func searchIPs(c *cli.Context) error {
    fmt.Println("----------IPs------------")
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		return fmt.Errorf("Failed to get the IP address of %s: %w", host, err)
	}
	
	for _, ip := range ips {
		fmt.Println(ip)
	}
	
	fmt.Println("-------------------------\n")

	return nil
}

func searchServersName(c *cli.Context) error {
    fmt.Println("----------Host NS------------")
    host := c.String("host")

    servers, err := net.LookupNS(host)
    if err != nil {
    	return fmt.Errorf("Failed to get the server name of %s: %w", host, err)
    }

    for _, server := range servers {
    	fmt.Println(server.Host)
    }

    fmt.Println("-----------------------------\n")

    return nil
}

func searchCNAME(c *cli.Context) error {
	fmt.Println("-------CNAME-------")
	host := c.String("host")

	cname, err := net.LookupCNAME(host)
	if err != nil {
		return fmt.Errorf("Failed to get the CNAME for %s: %w", host, err)
	}

	fmt.Println(cname)
	fmt.Println("-------------------\n")

	return nil
}

func searchTXT(c *cli.Context) error {
	fmt.Println("-----------------------------TXT records----------------------------")
	host := c.String("host")

	txts, erro := net.LookupTXT(host)
	if erro != nil {
		return fmt.Errorf("Failed to get TXT records of %s: %w", host, erro)	
	}

	for _, txt := range txts {
		fmt.Println(txt)
	}

	fmt.Println("--------------------------------------------------------------------\n")
	return nil
}
