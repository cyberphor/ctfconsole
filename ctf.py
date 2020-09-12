#!/usr/bin/env python3

import argparse
import code
import hashlib
import os
import sqlite3

class game():
    def __init__(self,args):
        if args.scoreboard:
            if os.path.exists(args.scoreboard):
                self.scoreboard = args.scoreboard
            else:
                print('[x] CTF scoreboard not found.')
                exit()
        elif args.create_scoreboard:
            if os.path.exists(args.create_scoreboard) == False:
                self.scoreboard = args.create_scoreboard
            else:
                print('[x] CTF scoreboard already exists.')
                exit()
        else: 
            self.scoreboard = 'scoreboard.db'
        self.connection = sqlite3.connect(self.scoreboard)
        self.cursor = self.connection.cursor()
        self.admin = self.administrator(args)
        self.challenge = self.admin.get_challenge
        self.data = self.admin.get_challenge_data
        self.solve = self.admin.solve_challenge
        if self.scoreboard_exists() == False:
            create = 'CREATE TABLE scoreboard (username TEXT, '
            where = 'password TEXT, score INTEGER)'
            create_table = create + where
            self.api(create_table)
       
    def api(self,action):
        records = self.cursor.execute(action).fetchall()
        keywords = ['CREATE','INSERT','UPDATE','DELETE']
        if any(trigger in action for trigger in keywords): 
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
        select = 'SELECT username, score FROM scoreboard '
        where = 'WHERE username = "%s"' % (username)
        query = select + where
        records = self.api(query)
        return records

    def add_player(self,username,password):
        if len(self.get_player(username)) == 0:
            password = hashlib.sha512(password.encode('UTF-8')).hexdigest()
            score = 0
            record = username, password, score
            add = 'INSERT INTO scoreboard VALUES ('
            where = '"%s", "%s", "%s")' % (record)
            add_record = add + where
            self.api(add_record)
            return '[+] Added %s to the scoreboard.' % (username)
        else: 
            return '[x] The username %s is already taken.' % (username)

    def get_scores(self):
        query = 'SELECT username, score FROM scoreboard'
        records = self.api(query)
        return records

    def correct_password(self,username,password):
        password = hashlib.sha512(password.encode('UTF-8')).hexdigest()
        select = 'SELECT password FROM scoreboard '
        where = 'WHERE username = "%s"' % (username)
        query = select + where
        records = self.api(query)
        if records:
            if password == records[0][0]: 
                return True
            else:
                return False

    def authorization(self):
        print('if correct admin password, self.authorization = "granted"')

    def remove_player(self,username,password):
        if self.correct_password(username,password) == True:
            delete = 'DELETE FROM scoreboard '
            where = 'WHERE username = "%s"' % (username)
            delete_record = delete + where
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
        def __init__(self,args):
            if args.database and os.path.exists(args.database):
                self.database = args.database
            elif args.create_database:
                if os.path.exists(args.create_database) == False:
                    self.database = args.create_database
                else:
                    print('[x] CTF database already exists.')
                    exit()
            elif os.path.exists('challenges.db'):
                self.database = 'challenges.db'
            else:
                print('[x] CTF database not found.')
                exit()
            self.connection = sqlite3.connect(self.database)
            self.cursor = self.connection.cursor()
            if self.table_exists('challenges') == False:
                create = 'CREATE TABLE challenges (number INTEGER, '
                table = 'challenge TEXT, data TEXT, solution TEXT, '
                create_table = create + table + 'points INTEGER)'
                self.api(create_table)

        def api(self,action):
            records = self.cursor.execute(action).fetchall()
            keywords = ['CREATE','INSERT','UPDATE','DELETE']
            if any(trigger in action for trigger in keywords): 
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
            select = 'SELECT data FROM challenges '
            where = 'WHERE number = "%s"' % (number)
            query = select + where
            records = self.api(query)
            if records == []:
                return records
            else:
                return records[0][0]

        def add_challenge(self,number,authorization):
            if len(self.get_challenge(number)) == 0:
                challenge = input('[>] Challenge: ')
                data = input('[>] Data: ')
                solution = input('[>] Solution: ')
                points = input('[>] Points: ')
                record = number, challenge, data, solution, points
                add = 'INSERT INTO challenges VALUES ('
                where = '"%s", "%s", "%s", "%s", "%s")' % (record)
                add_record = add + where
                answer = input('[>] Are you sure? (y/n)')
                if answer == 'y':
                    self.api(add_record)
                    return '[+] Added challenge #%s to the game.' % (number)
            else: 
                return '[x] Challenge #%s already exists.' % (number)

        def solve_challenge(self,number,answer):
            if len(self.get_challenge(number)) == 0:
                return '[+] Correct!'
            else:
                return '[x] Challenge #%s does not exist.' % (number)

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--create-scoreboard')
    parser.add_argument('--scoreboard')
    parser.add_argument('--create-database')
    parser.add_argument('--database')
    args = parser.parse_args()
    dashes = '-----------------------------------'
    motd = '[+] Welcome to the YellowTeam CTF!'
    banner = '\n'.join([dashes,motd,dashes])
    ctf = game(args)
    code.interact(banner=banner,local=locals())

if __name__ == '__main__':
    main()

# REFERENCES
# https://stackoverflow.com/questions/3389574/check-if-multiple-strings-exist-in-another-string
