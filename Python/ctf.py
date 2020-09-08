#!/usr/bin/env python3

import argparse
import code
import hashlib
import sqlite3

class scoreboard():
    def __init__(self):
        self.database = 'scoreboard.db'
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
        if 'CREATE' in action or 'INSERT' in action or 'UPDATE' in action or 'DELETE' in action: 
            self.connection.commit()
        else:
            return records

    def table_exists(self):
        query = 'SELECT count(name) FROM sqlite_master WHERE type = "table" and name = "players"'
        records = self.api(query)[0][0]
        if records > 0:
            return True
        else:
            return False

    def get_player(self,username):
        query = 'SELECT username, score FROM players WHERE username = "%s"' % (username)
        records = self.api(query)
        return records

    def add_player(self,username,password):
        if len(self.get_player(username)) == 0:
            password = hashlib.sha512(password.encode('UTF-8')).hexdigest()
            score = 0
            record = self.table1, username, password, score
            add_record = 'INSERT INTO %s VALUES ("%s", "%s", "%s")' % (record)
            self.api(add_record)
            return '[+] Added %s to the scoreboard.' % (username)
        else: 
            return '[x] The username %s is already taken.' % (username)

    def get_scores(self):
        query = 'SELECT username, score FROM players'
        records = self.api(query)
        return records

    def correct_password(self,username,password):
        password = hashlib.sha512(password.encode('UTF-8')).hexdigest()
        query = 'SELECT password FROM players WHERE username = "%s"' % (username)
        records = self.api(query)
        if len(records) > 0:
            if password == records[0][0]: 
                return True
            else:
                return False

    def remove_player(self,username,password):
        if self.correct_password(username,password) == True:
            delete_record = 'DELETE FROM players WHERE username = "%s"' % (username)
            self.api(delete_record)
            return '[!] Removed %s from the scoreboard.' % (username)
        else:
            return '[x] Invalid credentials.'

    def update_score(self,username,password,new_points):
        if self.correct_password(username,password) == True:
            score = self.get_player(username)[0][1] + new_points
            update_record = 'UPDATE players SET score = "%s" WHERE username = "%s"' % (score,username)
            self.api(update_record)
            return '[+] Updated the scoreboard.'
        else:
            return '[x] Invalid credentials.'

def add_challenge():
    challenge = input('[>] Challenge: ')
    solution = input('[>] Solution: ')

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--add-challenge',action='store_true')
    args = parser.parse_args()
    dashes = '-----------------------------------'
    motd = '[+] Welcome to the YellowTeam CTF!'
    banner = '\n'.join([dashes,motd,dashes])
    ctf = scoreboard()
    code.interact(banner=banner, local=locals())

if __name__ == '__main__':
    main()
