## ctfconsole
![GitHub](https://img.shields.io/github/license/cyberphor/ctfconsole)  
ctfconsole is a Capture The Flag (CTF) server written in Go. 

### Instructions
**Step 1.** Download ctfconsole from GitHub.
```
git clone https://github.com/cyberphor/ctfconsole
```

**Step 2.** Navigate to the repository downloaded and compile ctfconsole.
```
cd ctfconsole
go build .
```

**Step 3.** Start ctfconsole and browse to TCP port 8000 on your computer. 
```
.\ctfconsole.exe
```

![ctfconsole](/screenshot.png)  

### Third-Party Packages
```go
go get github.com/golang-jwt/jwt
go get github.com/mattn/go-sqlite3
```

### TODO
- [x] Update HTML filepaths
- [ ] Simplify identity management (admins, users > principal)
- [ ] Add GitHub Action to deploy binary as a Release

### Filepaths
- /views/pages.go (HTML Templates): `views/templates/*`
- /views/pages.go (Root directory of HTTP server): `filePath := http.Dir("views/")`
- /views/pages.go (Files HTTP server will serve): `http.Handle("/static/", fileServer)`
- /views/templates/header.gohtml: `/static/bootstrap.css`
- /views/static/bootstrap.css
- /views/static/bootstrap.js

### Copyright
This project is licensed under the terms of the [MIT license](/LICENSE).
