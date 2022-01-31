"""
    Date : 28/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import neded packages
from flask import Blueprint, redirect, request, session
from flask.helpers import flash
from flask.templating import render_template
from werkzeug.security import check_password_hash


# Import personnal modules
from python.database.databaseFunctions import addUser, getUserByUsername
from python.database.checkFunctions import checkUserUsername


# Definition of the blueprint
loginBP = Blueprint('loginBP', __name__)


# Definition of the login route
@loginBP.route('/login', methods=['GET', 'POST'])
def loginBP_login() -> str:
    if request.method == 'POST':
        username = request.form['identifier']
        password = request.form['password']

        user = getUserByUsername(username)
        print(user)

        if user['statement'] and check_password_hash(user['user']['cipherPassword'], password):
            session['id'] = user['user']['id']
            session['username'] = username

            flash("You have succesfully logged in.", "Success")
            return render_template('home.html')
        
        elif not user['statement']:
            flash("You have a wrong identifier, please try again.", "Error")
            return render_template('login.html')
        
        else:
            flash("You have a wrong password, please retry.", "Error")
            return render_template('login.html')


    return render_template('login.html')



@loginBP.route('/signup', methods=['GET', 'POST'])
def loginBP_signup():
    if request.method == 'GET':
        return render_template('signup.html')

    elif request.method == 'POST':
        username = request.form['username']
        password = request.form['password']

        if checkUserUsername(username):
            flash("Il y a déjà un joueur avec ce pseudo. Merci d'en choisir un autre.", "Error")
            
            return render_template('signup.html')
        
        else:
            addUser(username, password)
            flash("Vous vous êtes bien inscrit.", "Success")

            return redirect('/home')


# Definition of the logout route
@loginBP.route('/logout')
def loginBP_logout() -> str:
    session['id'] = None
    session['username'] = None

    return render_template('home.html')