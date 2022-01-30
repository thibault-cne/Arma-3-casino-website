"""
    Date : 30/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""


# Import personal modules
from python.database.connectDatabase import connectDatabase


def checkUserUsername(username: str) -> bool:
    query = '''SELECT * FROM User WHERE username=?;'''
    arg = (username, )

    db, cursor = connectDatabase()
    cursor.execute(query, arg)
    data = cursor.fetchall()
    db.close()

    if len(data) != 0:
        return False
    else:
        return True