"""
    Date : 23/01/2022
    Author : Thibault Chenevière
    Discord : @Thibô#0001
    Email : thibault.cheneviere@telecomnancy.eu
"""

# Import needed modules
from flask import Flask, flash, redirect




def create_app() -> Flask:
    """
        Function to create the app Flask object

        Entries :
            None
        
        Outputs :
            - app (Flask) : the app Flask object
    """
    app = Flask(__name__)

    app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
    app.config['SECRET_KEY'] = '29FTh4Swfr3DuMlNRcQcZxCk7IFBMooP'

    # Import blueprints
    ## Import main blueprint
    from python.blueprints.mainBP import mainBP
    app.register_blueprint(mainBP)

    ## Import roulette blueprint
    from python.blueprints.rouletteBP import rouletteBP
    app.register_blueprint(rouletteBP)

    # Error 404 handler
    @app.errorhandler(404)
    def pageNotFound(error):
        flash("HTTP 404 Not Found", "Red_flash")
        return redirect('/')
    
    return app
    

if __name__ == "__main__":
    app = create_app()
    app.run(debug=1, host="0.0.0.0", port=5454)