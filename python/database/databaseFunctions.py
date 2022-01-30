"""
    Date : 28/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
from werkzeug.security import check_password_hash


# Import personnal modules
from python.database.connectDatabase import connectDatabase


def getUserByUsername(username: str) -> dict:
    """
        Function to get the user with is unique username
    """
    result = {
        'statement': False,
        'user': {
            'id': "",
            'username': username,
            'cipherPassword': ""
        }
    }

    query = """SELECT * FROM User;"""
    db, cursor = connectDatabase()
    data = cursor.execute(query).fetchall()
    db.close()

    for users in data:
        if check_password_hash(users[1], username):
            result['statement'] = True
            result['id'] = users[0]
            result['cipherPassword'] = users[2]
    
    return result