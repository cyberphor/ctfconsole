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

## Level 1: Output
```pwsh
Write-Host "Hello world!"
```

## Level 2: Input
```pwsh
$name = Read-Host -Prompt "What is your name?"
Write-Host $name
```

## Level 3: Arithmetic
```pwsh
1+1
```

## Level 4: Conditionals
```pwsh
$name = Read-Host -Prompt "What is your name?"
if ($name -eq 'victor') {
    Write-Host "Long time, no see $name"
} else {
    Write-Host "Hello $name"
}
```

## Level 5: Loops
```pwsh
1..3 | ForEach-Object {
    Write-Host "Counting down $_"
}
```

## References
* None (yet)

## Copyright
This project is licensed under the terms of the [MIT license](/_misc/LICENSE).
