# goRecon
A Go command-line tool for basic web server reconnaissance. Project for study.

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
### **How to Use**

```bash
./gorecon [command] --host [domain]
```
| Command | Description | Example |
| --- | --- | --- |
| `ip` | Lookup all A and AAAA records for the domain | `./gorecon ip --host google.com.br` |
| `ns` | Lookup NS records for the domain | `./gorecon ns --host google.com.br` |
| `cnm` | Lookup the canonical name for the domain | `./gorecon cnm --host google.com.br` |
| `txt` | Lookup txt records for the domain | `./gorecon txt --host google.com.br` |


### **Example with image**
<img src="docs/example_basic.png" width="600" alt="goRecon Demo">

---

### **Author**

**DarkZer010** 
- Telegram: [@DarkZer0](https://t.me/DarkZer0_dz0)
- X (old twitter): [DarkZer0_dz0](https://x.com/DarkZer0_dz0)

License: MIT
