"""
    Date : 30/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import neded packages
from flask import Blueprint, flash, redirect, request, session


# Import personnal modules


# Definition of the blueprint
flashMessageBP = Blueprint('flashMessageBP', __name__)


@flashMessageBP.route('/<string:randomCode>/flashMessage', methods=['GET', 'POST'])
def flashMessageBP_flash(randomCode: str) -> dict:
    assert len(randomCode) == 6
    
    if request.method == 'POST':
        flashMessage = request.form['content']
        flashColor = request.form['flashColor']

        flash(flashMessage, flashColor)
        return {'statement': "success"}
