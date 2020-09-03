#!/usr/bin/env python3

import argparse
import hashlib
import os
import sqlite3
import sys

class scoreboard():
    def __init__(self,database_name):
        self.database = database_name
        self.connection = sqlite3.connect(self.database)
        self.cursor = self.connection.cursor()
        self.table1 = 'players'
        self.column1 = 'username'
        self.column2 = 'password'
        self.column3 = 'score'
        what = "SELECT count(name) FROM sqlite_master " 
        where = "WHERE type = 'table' and name = '%s'" % (self.table1)
        query = what + where
        self.cursor.execute(query)
        if self.cursor.fetchone()[0] == 0:
            table = self.table1, self.column1, self.column2, self.column3
            create_table = "CREATE TABLE %s (%s TEXT, %s TEXT, %s INTEGER)" % (table)
            self.api(create_table)

    def api(self,action):
        self.cursor.execute(action)
        self.connection.commit()

    def add_player(self):
        print('[>] Username: ')

    def get_scores(self):
        print('[+] Scoreboard: ')

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--add-player', action='store_true')
    parser.add_argument('--get-scores', action='store_true')
    args = parser.parse_args()
    database_name = 'ctf-player.db'
    if len(sys.argv) > 1:
        ctf = scoreboard(database_name)
        if args.add_player:
            ctf.add_player()
        elif args.get_scores: 
            ctf.get_scores()
    else:
        print('[x] No option specified.')

# REFERENCES
# https://www.digitalocean.com/community/tutorials/how-to-use-the-sqlite3-module-in-python-3
# https://pythonexamples.org/python-sqlite3-check-if-table-exists/
# https://www.pythonforbeginners.com/system/python-sys-argv
