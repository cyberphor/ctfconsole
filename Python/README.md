# Yellow Team CTF Challenges
![yellow.gif](/_misc/yellow.gif)

## Table of Contents
* [Home](/README.md)
* [PowerShell](/PowerShell/README.md)
* [Python](/Python/README.md)
    * [Level 1: Output](#level-1-output)
    * [Level 2: Input](#level-2-input)
    * [Level 3: Arithmetic](#level-3-arithmetic)
    * [Level 4: Conditionals](#level-4-conditionals)
    * [Level 5: Loops](#level-5-loops)
* [References](#references)

## Level 1: Output
```python
#!/usr/bin/env python3

print("Hello world!")
```

## Level 2: Input
```python
#!/usr/bin/env python3

name = input("What is your name? ")
print(name)
```

## Level 3: Arithmetic
```python
#!/usr/bin/env python3

name = input("What is your name? ")
print("Hello " + name + "!")
```

## Level 4: Conditionals
```python
#!/usr/bin/env python3

name = input("What is your name? ")
if name == "neo":
    print("You are the one.")
else:
    print("You are not the one, " + name)
```

## Level 5: Loops
```python
#!/usr/bin/env python3

name = input("What is your name? ")
for letter in name:
    print("Spelling your name, one letter at a time: " + letter)
```

## References
* None (yet)

## Copyright
This project is licensed under the terms of the [MIT license](/_misc/LICENSE).
