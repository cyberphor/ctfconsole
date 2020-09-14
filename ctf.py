#!/usr/bin/env python3

import argparse
import code
import hashlib
import os
import sqlite3

class game():
    def __init__(self,args):
        if args.use_scoreboard:
            if os.path.exists(args.use_scoreboard):
                self.scoreboard = args.use_scoreboard
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
            self.scoreboard = 'scoreboard.sqlite'
        self.connection = sqlite3.connect(self.scoreboard)
        self.cursor = self.connection.cursor()
        self.admin = self.administrator(args)
        self.challenge = self.admin.get_challenge
        self.data = self.admin.get_challenge_data
        self.solve = self.admin.solve_challenge
        if self.scoreboard_exists() == False:
            create_table = '''CREATE TABLE scoreboard
                (username TEXT, password TEXT, score INTEGER)'''
            self.api(create_table,None)
       
    def api(self,action,parameters):
        if parameters == None:
            records = self.cursor.execute(action).fetchall()
        else:
            records = self.cursor.execute(action,parameters).fetchall()
        keywords = ['CREATE','INSERT','UPDATE','DELETE']
        if any(trigger in action for trigger in keywords): 
            self.connection.commit()
        else:
            return records

    def scoreboard_exists(self):
        query = '''SELECT count(name) FROM sqlite_master 
            WHERE type = "table" and name = "scoreboard"'''
        records = self.api(query,None)[0][0]
        if records > 0:
            return True
        else:
            return False

    def get_player(self,username):
        query = '''SELECT username, score FROM scoreboard 
            WHERE username = ?'''
        username = (username,)
        record = self.api(query,username)
        return record

    def add_player(self,username,password):
        if len(self.get_player(username)) == 0:
            password = hashlib.sha512(password.encode('UTF-8')).hexdigest()
            add = '''INSERT INTO scoreboard VALUES (?, ?, ?)'''
            record = (username, password, 0)
            self.api(add,record)
            return '[+] Added %s to the scoreboard.' % (username)
        else: 
            return '[x] The username %s is already taken.' % (username)

    def get_scores(self):
        query = '''SELECT username, score FROM scoreboard'''
        records = self.api(query,None)
        return records

    def correct_password(self,username,password):
        password = hashlib.sha512(password.encode('UTF-8')).hexdigest()
        query = '''SELECT password FROM scoreboard 
            WHERE username = ?'''
        username = (username,)
        record = self.api(query,username)
        if record:
            if password == record[0][0]: 
                return True
            else:
                return False

    def authorization(self):
        print('if correct admin password, self.authorization = "granted"')

    def remove_player(self,username,password):
        if self.correct_password(username,password) == True:
            delete = '''DELETE FROM scoreboard
                WHERE username = ?'''
            username = (username,)
            self.api(delete,username)
            return '[!] Removed %s from the scoreboard.' % (username)
        else:
            return '[x] Invalid credentials.'

    def update_score(self,username,password,new_points):
        if self.correct_password(username,password) == True:
            update = '''UPDATE players SET score = ?
                WHERE username = ?'''
            score = self.get_player(username)[0][1] + new_points
            record = (username, score)
            self.api(update,record)
            return '[+] Updated the scoreboard.'
        else:
            return '[x] Invalid credentials.'

    class administrator():
        def __init__(self,args):
            if args.use_database and os.path.exists(args.use_database):
                self.database = args.use_database
            elif args.create_database:
                if os.path.exists(args.create_database) == False:
                    self.database = args.create_database
                else:
                    print('[x] CTF database already exists.')
                    exit()
            elif os.path.exists('challenges.sqlite'):
                self.database = 'challenges.sqlite'
            else:
                print('[x] CTF database not found.')
                exit()
            self.connection = sqlite3.connect(self.database)
            self.cursor = self.connection.cursor()
            if self.challenges_exist() == False:
                create_table = '''CREATE TABLE challenges 
                    (number INTEGER, points INTEGER, challenge TEXT, 
                    solution TEXT, data TEXT)'''
                self.api(create_table,None)
            if args.add_challenges:
                self.read_csv_file(args.add_challenges)
                exit()

        def api(self,action,parameters):
            if parameters == None:
                records = self.cursor.execute(action).fetchall()
            else:
                records = self.cursor.execute(action,parameters).fetchall()
            keywords = ['CREATE','INSERT','UPDATE','DELETE']
            if any(trigger in action for trigger in keywords): 
                self.connection.commit()
            else:
                return records

        def challenges_exist(self):
            query = '''SELECT count(name) FROM sqlite_master 
                WHERE type = "table" and name = "challenges"'''
            records = self.api(query,None)[0][0]
            if records > 0:
                return True
            else:
                return False

        def get_challenge(self,number):
            query = '''SELECT challenge FROM challenges 
                WHERE number = ?'''
            number = (number,)
            record = self.api(query,number)
            if len(record) > 0:
                return record[0][0]
            else: 
                return record

        def get_challenge_data(self,number):
            query = '''SELECT data FROM challenges 
                WHERE number = ?'''
            number = (number,)
            record = self.api(query,number)
            if len(record) > 0:
                data = eval(record[0][0])
                return data
            else:
                return record

        def add_challenge(self,number,points,challenge,solution,data):
            if len(self.get_challenge(number)) == 0:
                if type(solution).__name__ != 'str':
                    solution = str(solution)
                if type(data).__name__ != 'str':
                    data = str(data)
                add = '''INSERT INTO challenges VALUES (?, ?, ?, ?, ?)'''
                record = (number, points, challenge, solution, data)
                self.api(add,record)
                print('[+] Added challenge #%s to the game.' % (number))
            else: 
                print('[x] Challenge #%s already exists.' % (number))

        def read_csv_file(self,filename):
            if os.path.exists(filename):
                with open(filename) as filedata:
                    for line in filedata.readlines():
                        if line[0] != '#':
                            self.add_challenge(*eval(line))

        def solve_challenge(self,number,answer):
            if len(self.get_challenge(number)) > 0:
                query = '''SELECT solution FROM challenges 
                    WHERE number = ?'''
                number = (number,)
                record = self.api(query,number)[0][0]
                solution = eval(record)
                if answer == solution:
                    return '[+] Correct!'
                else:
                    return '[x] Incorrect.'
            else:
                return '[x] Challenge #%s does not exist.' % (number)

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--create-scoreboard')
    parser.add_argument('--use-scoreboard')
    parser.add_argument('--create-database')
    parser.add_argument('--use-database')
    parser.add_argument('--add-challenges')
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
# https://stackoverflow.com/questions/19112735/how-to-print-a-list-of-tuples-with-no-brackets-in-python
