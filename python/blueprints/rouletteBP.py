"""
    Date : 23/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
from flask import Blueprint, render_template, request, session


# Create main blueprint
rouletteBP = Blueprint('rouletteBP', __name__)


# Definition of the routes
## Definition of main route
@rouletteBP.route('/roulette', methods=["GET", "POST"])
def roulette_page() -> str:
    if request.method == "GET":
        if session.get('id'):

            return render_template("roulette.html", userId=session['id'])
        
        else:
            return render_template("roulette.html")

    elif request.method == "POST":
        resultDict = request.form.to_dict()

        return "success"


@rouletteBP.route("/addUser", methods=['GET', 'POST'])
def roulette_addUser():
    if request.method == 'POST':
        print(request.form)
        return {"request": "POST", "color": request.form['color']}
    elif request.method == 'GET':
        print(request.args)
        return {"request": 'GET'}
