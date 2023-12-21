## ctfconsole
![GitHub](https://img.shields.io/github/license/cyberphor/ctfconsole)  
ctfconsole is a Capture The Flag (CTF) server. 

### Instructions
**Step 1.** Download ctfconsole from GitHub.
```bash
git clone https://github.com/cyberphor/ctfconsole &&\
cd ctfconsole
```

**Step 2.** Build and run the required containers using Docker Compose. 
```bash
docker compose up --profile "ctfconsole" up
```

**Step 3.** Once you see a message similar to the one below, open a browser to `http://localhost`.
```bash
ctfconsole_frontend  | You can now view ctfconsole in the browser.
```

**Step 4.** When you're done using ctfconsole, enter the command below. 
```bash
docker compose --profile "ctfconsole" down 
```

![ctfconsole](/screenshot.png)  

### Copyright
This project is licensed under the terms of the [MIT license](/LICENSE).