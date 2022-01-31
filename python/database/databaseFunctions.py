"""
    Date : 28/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
from werkzeug.security import check_password_hash, generate_password_hash


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
            'picture': '',
            'cipherPassword': ""
        }
    }

    query = """SELECT * FROM User;"""
    db, cursor = connectDatabase()
    data = cursor.execute(query).fetchall()
    db.close()

    for users in data:
        print(users)
        if check_password_hash(users[3], username):
            result['statement'] = True
            result['user']['id'] = users[0]
            result['user']['picture'] = users[2]
            result['user']['cipherPassword'] = users[4]
    
    return result


def addUser(username: str, password: str) -> None:
    query = '''INSERT INTO User (username, password) VALUES (?, ?);'''

    cipherUsername = generate_password_hash(username, 'sha256')
    cipherPassword = generate_password_hash(password, 'sha256')

    args = (cipherUsername, cipherPassword)

    db, cursor = connectDatabase()
    cursor.execute(query, args)
    db.commit()
    db.close()