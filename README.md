# Yellow Team CTF
![yellow.gif](/_misc/yellow.gif)

## Table of Contents
* [Installation](#installation)
* [Usage](#usage)
* [What is a Yellow Team?](#what-is-a-yellow-team)
* [References](#references)

## Installation
```bash
git clone https://github.com/cyberphor/yellowteam-ctf.git
chmod +x ./YellowTeam-CTF/ctf.py
cp ./YellowTeam-CTF/ /usr/local/games/
```

## Usage
### Starting the CTF
```
ctf.py
```
### Solving a Challenge
```
-----------------------------------
[+] Welcome to the YellowTeam CTF!
-----------------------------------
>>> ctf.get_challenge(1)
'What is the sum of ctf.data(1) + 1?'
>>> ctf.data(1)
22
>>> ctf.solve(1,ctf.data(1)+1)
'[+] Correct!'
```
### Adding a Challenge 
```
>>> ctf.admin.add_challenge(42)
[>] Challenge: Carve-out and hash the JPEG found in ctf.data(42).
[>] Solution: 5eb63bbbe01eeed093cb22bb8f5acdc3
[>] Points: 5
[>] Are you sure? (y/n) y
'[+] Added challenge #42 to the game.'
```

## What is a Yellow Team?
> These are the people that build and design software, systems, and integrations that make businesses more efficient. Application developers, software engineers and architects fall into this category. <br> - Hackernoon

## References
* [Hackernoon - "Introducing the InfoSec colour wheel — blending developers with red and blue security teams."](https://hackernoon.com/introducing-the-infosec-colour-wheel-blending-developers-with-red-and-blue-security-teams-6437c1a07700)

## Copyright
This project is licensed under the terms of the [MIT license](/_misc/LICENSE).
