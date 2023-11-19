## ctfconsole
![GitHub](https://img.shields.io/github/license/cyberphor/ctfconsole)  
ctfconsole is a Capture The Flag (CTF) server. 

### Instructions
**Step 1.** Download ctfconsole from GitHub.
```
git clone https://github.com/cyberphor/ctfconsole &&\
cd ctfconsole
```

**Step 2.** Build and run the required containers using Docker Compose. 
```
docker compose up --profile "ctfconsole" up
```

**Step 3.** Once you see a message similar to the one below, open a browser to `http://localhost`.
```
ctfconsole_frontend  | You can now view ctfconsole in the browser.
```

![ctfconsole](/screenshot.png)  

### Copyright
This project is licensed under the terms of the [MIT license](/LICENSE).
