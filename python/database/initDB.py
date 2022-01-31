"""
    Date : 24/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
import pathlib, sys

# Add parent dir to the path
cwd = pathlib.Path(__file__).parents[2]
parentdir = str(cwd)
sys.path.append(parentdir)

# Import personal modules
from python.database.connectDatabase import connectDatabase


if __name__ == "__main__":
    db, cursor = connectDatabase()
    
    # Script for the database
    query = '''
    DROP TABLE IF EXISTS User;
    DROP TABLE IF EXISTS Account;
    DROP TABLE IF EXISTS Roulette;

    CREATE TABLE User
    (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        steamId INTEGER UNIQUE,
        picture TEXT UNIQUE,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );

    CREATE TABLE Account
    (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        amount INTEGER,
        FOREIGN KEY (id) REFERENCES User(id)
    );

    CREATE TABLE Roulette
    (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        zero INTEGER,
        red INTEGER,
        black INTEGER,
        FOREIGN KEY (id) REFERENCES User(id)
    )
    '''

    cursor.executescript(query)
    db.commit()
    db.close()