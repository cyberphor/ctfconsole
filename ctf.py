#!/usr/bin/env python3

import argparse
import code
import hashlib
import os
import sqlite3

class game():
    def __init__(self):
        self.database = 'scoreboard.db'
        self.connection = sqlite3.connect(self.database)
        self.cursor = self.connection.cursor()
        self.admin = self.administrator()
        self.get_challenge = self.admin.get_challenge
        if self.scoreboard_exists() == False:
            table = 'scoreboard', 'username', 'password', 'score'
            create_table = 'CREATE TABLE %s (%s TEXT, %s TEXT, %s INTEGER)' % (table)
            self.api(create_table)
       
    def api(self,action):
        records = self.cursor.execute(action).fetchall()
        if 'CREATE' in action or 'INSERT' in action or 'UPDATE' in action or 'DELETE' in action: 
            self.connection.commit()
        else:
            return records

    def scoreboard_exists(self):
        select = 'SELECT count(name) FROM sqlite_master '
        where = 'WHERE type = "table" and name = "scoreboard"'
        query = select + where
        records = self.api(query)[0][0]
        if records > 0:
            return True
        else:
            return False

    def get_player(self,username):
        select = 'SELECT username, score FROM players '
        where = 'WHERE username = "%s"' % (username)
        query = select + query
        records = self.api(query)
        return records

    def add_player(self,username,password):
        if len(self.get_player(username)) == 0:
            password = hashlib.sha512(password.encode('UTF-8')).hexdigest()
            score = 0
            record = self.players, username, password, score
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
        select = 'SELECT password FROM players '
        where = 'WHERE username = "%s"' % (username)
        query = select + where
        records = self.api(query)
        if len(records) > 0:
            if password == records[0][0]: 
                return True
            else:
                return False

    def authorization(self):
        print('if correct admin password, self.authorization = "granted"')

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
            update = 'UPDATE players SET score = '
            where = '"%s" WHERE username = "%s"' % (score,username)
            update_record = update + where
            self.api(update_record)
            return '[+] Updated the scoreboard.'
        else:
            return '[x] Invalid credentials.'

    class administrator():
        def __init__(self):
            self.database = 'challenges.db'
            self.connection = sqlite3.connect(self.database)
            self.cursor = self.connection.cursor()
            if self.table_exists('challenges') == False:
                table = 'challenges', 'number', 'challenge', 'solution', 'points'
                create_table = 'CREATE TABLE %s (%s INTEGER, %s TEXT, %s TEXT, %s INTEGER)' % (table)
                self.api(create_table)

        def api(self,action):
            records = self.cursor.execute(action).fetchall()
            if 'CREATE' in action or 'INSERT' in action or 'UPDATE' in action or 'DELETE' in action: 
                self.connection.commit()
            else:
                return records

        def table_exists(self,table):
            select = 'SELECT count(name) FROM sqlite_master '
            where = 'WHERE type = "table" and name = "%s"' % (table)
            query = select + where
            records = self.api(query)[0][0]
            if records > 0:
                return True
            else:
                return False

        def get_challenge(self,number):
            select = 'SELECT challenge, points FROM challenges '
            where = 'WHERE number = "%s"' % (number)
            query = select + where
            records = self.api(query)
            return records

        def get_challenge_data(self,number):
            select = 'SELECT challenge, data FROM data '
            where = 'WHERE number = "%s"' % (number)
            query = select + query
            records = self.api(query)
            return records

        def add_challenge(self,number,authorization):
            if len(self.get_challenge(number)) == 0:
                challenge = input('[>] Challenge: ')
                solution = input('[>] Solution: ')
                points = input('[>] Points: ')
                answer = input('[>] Are you sure? (y/n)')
                record = 'challenges', number, challenge, solution, points
                add_record = 'INSERT INTO %s VALUES ("%s", "%s", "%s", "%s")' % (record)
                if answer == 'y':
                    self.api(add_record)
                    return '[+] Added challenge #%s to the game.' % (number)
            else: 
                return '[x] Challenge #%s already exists.' % (number)

        def add_challenge_data(self,number,authorization):
            if len(self.get_challenge_data(number)) == 0:
                return
            else:
                return '[x] Data for challenge #%s already exists.' % (number)

def main():
    dashes = '-----------------------------------'
    motd = '[+] Welcome to the YellowTeam CTF!'
    banner = '\n'.join([dashes,motd,dashes])
    ctf = game()
    code.interact(banner=banner, local=locals())

if __name__ == '__main__':
    main()
