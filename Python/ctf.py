#!/usr/bin/env python3

import argparse
import code
import hashlib
import sqlite3

class scoreboard():
    def __init__(self,name):
        self.database = name + '.db'
        self.connection = sqlite3.connect(self.database)
        self.cursor = self.connection.cursor()
        self.table1 = 'players'
        self.column1 = 'username'
        self.column2 = 'password'
        self.column3 = 'score'
        if self.table_exists() == False:
            table = self.table1, self.column1, self.column2, self.column3
            create_table = 'CREATE TABLE %s (%s TEXT, %s TEXT, %s INTEGER)' % (table)
            self.api(create_table)
            
    def api(self,action):
        records = self.cursor.execute(action).fetchall()
        if 'CREATE' in action: self.connection.commit()
        elif 'INSERT' in action: self.connection.commit()
        else: return records

    def table_exists(self):
        query = 'SELECT count(name) FROM sqlite_master WHERE type = "table" and name = "players"'
        records = self.api(query)[0][0]
        if records > 0:
            return True
        else:
            return False

    def player(self,username):
        query = 'SELECT username, score FROM players WHERE username = "%s"' % (username)
        records = self.api(query)
        return records

    def add_player(self):
        username = input('[>] Username: ')
        if len(self.player(username)) == 0:
            password = hashlib.sha512(input('[>] Password: ').encode('UTF-8')).hexdigest()
            score = 0
            record = self.table1, username, password, score
            add_record = 'INSERT INTO %s VALUES ("%s", "%s", "%s")' % (record)
            self.api(add_record)
            return self.get_scores()[-1]
        else: 
            return '[x] The username %s is already taken.' % (username)

    def get_scores(self):
        query = 'SELECT username, score FROM players'
        records = self.api(query)
        return records

def add_challenge():
    challenge = input('[>] Challenge: ')
    solution = input('[>] Solution: ')

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--add-challenge',action='store_true')
    args = parser.parse_args()
    name = 'YellowTeam'
    dashes = '-----------------------------------'
    motd = '[+] Welcome to the %s CTF!' % (name)
    banner = '\n'.join([dashes,motd,dashes])
    ctf = scoreboard(name)
    code.interact(banner=banner, local=locals())

if __name__ == '__main__':
    main()
