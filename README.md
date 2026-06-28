# goRecon
A Go command-line tool for basic web server reconnaissance. Quickly lookup IPs and Name Servers for any domain straight from your terminal.

Built with `github.com/urfave/cli v1`.

### **Installation (TERMUX/LINUX)**

```bash
apt install git -y
apt install golang -y
git clone https://github.com/DarkZer010/goRecon.git
cd goRecon
go mod tidy
go build
```

Fechou. Aqui tá o `README.md` completo já com `Usage` e a tabela de comandos no formato que tu pediu 👇

Copia e cola direto:
goRecon
A Go command-line tool for basic web server reconnaissance. Quickly lookup IPs and Name Servers for any domain straight from your terminal.

Built with `github.com/urfave/cli v1`.

**Installation (TERMUX/LINUX)**

```bash
apt install git -y
apt install golang -y
git clone https://github.com/DarkZer010/goRecon.git
cd goRecon
go mod tidy
go build
chmod +x goRecon
*How to Use*
./goRecon [command] --host [domain]
*Flag:*
`--host` : Target domain. Default: `devbook.com.br`

*Available Commands*
Command	Description	Example
`ip`	Lookup all A and AAAA records for the domain	`./goRecon ip --host google.com.br`
`servername`	Lookup NS records for the domain	`./goRecon servername --host google.com.br`
*Example Output*
$./goRecon ip --host google.com.br
Informações básica do Servidor:

----------IPs------------
142.250.218.195
2a00:1450:4001:824::2003


$./goRecon servername --host google.com.br
Informações básica do Servidor:

----------Host NS------------
ns1.google.com.
ns2.google.com.
ns3.google.com.
ns4.google.com.

*Tech Stack*
- Go 1.26.3
- urfave/cli v1.22.17

*Roadmap*
- [ ] `cname` command to lookup CNAME records
- [ ] `mx` command to lookup Mail Exchange records
- [ ] `-v` flag for verbose mode

*Author*
Made by @DarkZer010

License: MIT

**O que mudei:**
1. `How to Use` com o comando base `./goRecon [command] --host [domain]`
2. Tabela `Available Commands` exatamente no formato que tu quer, com 3 colunas
3. Mantive `Termux/LINUX` porque é teu público

Só dar `git add README.md` > `git commit` > `git push` que fica top no perfil.

Bora adicionar o `cname` agora pra tabela ficar com 3 linhas?
