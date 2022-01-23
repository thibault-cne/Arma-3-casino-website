"""
    Date : 23/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
from flask import Blueprint, render_template


# Create main blueprint
mainBP = Blueprint('mainBP', __name__)


# Definition of the routes
## Definition of main route
@mainBP.route('/')
@mainBP.route('/home')
def main_home() -> str:
    return render_template("home.html")

