# Yellow Team CTF Challenges
![yellow.gif](/_misc/yellow.gif)

## Table of Contents
* [Home](/README.md)
* [Getting Setup](/_misc/getting_setup.md)
* [PowerShell](/PowerShell/README.md)
    * [Level 1: Output](#level-1-output)
    * [Level 2: Input](#level-2-input)
    * [Level 3: Arithmetic](#level-3-arithmetic)
    * [Level 4: Conditionals](#level-4-conditionals)
    * [Level 5: Loops](#level-5-loops)
    * [References](#references)
* [Python](/Python/README.md)

# PowerShell
## Level 0
```bash
sudo apt install powershell
```

## Level 1: Output
```bash
vim level_1.ps1
```
```pwsh
#!/usr/bin/env pwsh

Write-Host "Hello world!"
```
```bash
chmod +x level_1.ps1
./level_1.ps1
```

## Level 2: Input
```bash
vim level_2.ps1
```
```pwsh
#!/usr/bin/env pwsh

$name = Read-Host -Prompt "What is your name?"
Write-Host $name
```
```bash
chmod +x level_2.ps1
./level_2.ps1
```

## Level 3: Arithmetic
```bash
vim level_3.ps1
```
```pwsh
#!/usr/bin/env pwsh

1+1
```
```bash
chmod +x level_3.ps1
./level_3.ps1
```

## Level 4: Conditionals
```bash
vim level_4.ps1
```
```pwsh
#!/usr/bin/env pwsh

$name = Read-Host -Prompt "What is your name?"
if ($name -eq 'victor') {
    Write-Host "Long time, no see $name"
} else {
    Write-Host "Hello $name"
}
```
```bash
chmod +x level_4.ps1
./level_4.ps1
```

## Level 5: Loops
```bash
vim level_5.ps1
```
```pwsh
#!/usr/bin/env pwsh

1..3 | ForEach-Object {
    Write-Host "Counting down $_"
}
```
```bash
chmod +x level_5.ps1
./level_5.ps1
```

## References
* None (yet)

## Copyright
This project is licensed under the terms of the [MIT license](/_misc/LICENSE).
