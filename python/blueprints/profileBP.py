"""
    Date : 31/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import neded packages
from flask import Blueprint, flash, render_template, request, session

# Import personal modules
from python.database.databaseFunctions import getUserByUsername


# Import personnal modules


# Definition of the blueprint
profileBP = Blueprint('profileBP', __name__)


@profileBP.route('/profile/<string:username>', methods=['POST', 'GET'])
def profileBP_profile(username: str):
    if request.method == 'GET':
        data = getUserByUsername(username)
        if data['statement']:
            return render_template("profile.html", data=data)
        else:
            session['id'] = None
            session['username'] = None

            flash('Une erreur vient de survenir. Merci de se reconnecter.', 'Error')
            return render_template('home.html')