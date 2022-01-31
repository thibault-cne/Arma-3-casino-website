"""
    Date : 30/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
from werkzeug.security import check_password_hash

# Import personal modules
from python.database.connectDatabase import connectDatabase


def checkUserUsername(username: str) -> bool:
    query = '''SELECT * FROM User'''

    db, cursor = connectDatabase()
    cursor.execute(query)
    data = cursor.fetchall()
    db.close()

    for users in data:
        if check_password_hash(users[3], username):
            return True
    
    return False