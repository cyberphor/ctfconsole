#!/usr/bin/env python3

import argparse
import code

parser = argparse.ArgumentParser()
parser.add_argument('--username')
parser.add_argument('--password')
args = parser.parse_args()

class challenges():
    def __init__(self):
        if args.username and args.password:
            self.username = args.username
            self.password = args.password
            self.registered = True
            self.logged_in = True
            self.score = 0
        else:
            self.username = ''
            self.password = ''
            self.registered = False
            self.logged_in = False
            self.score = 0
        self.player_db = open('./yellowteam-player.db','rb')
        self.ctf_db = open('./yellowteam-ctf.db','rb')

    def player(self):
        if self.registered:
            if self.logged_in:
                return self.username, self.score
            else: return '[x] Please login using ctf.login().'
        else: return '[x] Please register using ctf.register().'

    def register(self,username,password):
        if len(username) > 0:
            self.username = username
            if len(password) > 0:
                self.password = password
                self.registered = True
                return '[+] Successfully registered.'
            else: return '[x] Invalid password.'
        else: return '[x] Invalid username.'
   
    def login(self,username,password):
        if username == self.username:
            if password == self.password:
                self.logged_in = True
                return '[+] Successfully logged-in.'
            else: return '[x] Invalid credentials.'
        else: return '[x] Invalid credentials.'

    def question(self,challenge):
        return '[>] What is the sum of 1 + 1?'

    def answer(self,challenge,solution):
        return '[+] Correct!'

def start():
    ctf = challenges()
    code.interact(banner='[+] Yellow Team CTF Challenges', local=locals())

if __name__ == '__main__':
    start()
