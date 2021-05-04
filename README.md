# ctfconsole
![yellow.gif](/_misc/yellow.gif)

## Table of Contents
* [Installation](#installation)
* [Usage](#usage)
  * [Starting ctfconsole](#starting-ctfconsole)
  * [Solving a Challenge](#solving-a-challenge)
  * [Adding Challenges for Your Own CTF Event](#adding-challenges-for-your-own-ctf-event)
* [What is a Yellow Team?](#what-is-a-yellow-team)
* [Copyright](#copyright)
* [References](#references)

## Installation
```bash
git clone https://github.com/cyberphor/ctfconsole.git
```

## Usage
### Starting ctfconsole
```
./ctfconsole.py
```
### Solving a Challenge
```
-----------------------------------
[+] Welcome to the YellowTeam CTF!
-----------------------------------
>>> ctf.challenge(1)
'What is the sum of ctf.data(1) + 1?'
>>> ctf.data(1)
22
>>> ctf.solve(1,ctf.data(1)+1)
'[+] Correct!'
```
### Adding Challenges for Your Own CTF Event
```
./ctfconsole.py --add-challenges YellowTeam-CTF.csv
[+] Added 75 CTF challenges.
 --> 23, 42, 69 already existed.
```

## What is a Yellow Team?
> These are the people that build and design software, systems, and integrations that make businesses more efficient. Application developers, software engineers and architects fall into this category. <br> - Hackernoon

## Copyright
This project is licensed under the terms of the [MIT license](/_misc/LICENSE).

## References
* [Hackernoon - "Introducing the InfoSec colour wheel — blending developers with red and blue security teams."](https://hackernoon.com/introducing-the-infosec-colour-wheel-blending-developers-with-red-and-blue-security-teams-6437c1a07700)
