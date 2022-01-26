"""
    Date : 23/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
from flask import Blueprint, render_template, request


# Create main blueprint
rouletteBP = Blueprint('rouletteBP', __name__)


# Definition of the routes
## Definition of main route
@rouletteBP.route('/roulette', methods=["GET", "POST"])
def main_home() -> str:
    if request.method == "GET":
        return render_template("roulette.html")

    elif request.method == "POST":
        resultDict = request.form.to_dict()
        print(resultDict)
        return "success"
