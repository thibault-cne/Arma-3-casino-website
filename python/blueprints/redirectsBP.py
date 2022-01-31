"""
    Date : 31/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import neded packages
from flask import Blueprint, redirect, session


# Import personnal modules
from python.database.connectDatabase import connectDatabase


# Definition of the blueprint
redirectsBP = Blueprint('redirectsBP', __name__)


@redirectsBP.route("/r/<string:mainPage>/<int:userID>")
def redirects_route(mainPage: str, userID: int) -> str:
    query = '''SELECT * FROM User WHERE id=?;'''
    arg = (userID, )

    db, cursor = connectDatabase()
    cursor.execute(query, arg)
    data = cursor.fetchall()
    data = data[0]
    db.close()

    if mainPage == 'profile':
        return redirect(f"/{mainPage}/{session['username']}")