function validateForm() {
    var password = document.getElementById('pwdInput').value;
    var confirmationPassword = document.getElementById('pwdConfirmInput').value;

    if (password != confirmationPassword) {
        var jsonData = {
            'content': "Les mots de passe ne sont pas identiques.",
            'flashColor': "Error"
        }
        var url = generateFlashUrl();

        var status = sendRequest(url, jsonData);

        if (status) {
            return false
        }
    }

    if (password.length < 4) {
        var jsonData = {
            'content': "Le mot de passe est trop court. Merci de choisir un mot de passe de plus de 4 caractères.",
            'flashColor': "Error"
        }
        var url = generateFlashUrl();

        var status = sendRequest(url, jsonData);

        if (status) {
            return false
        }
    }
}


function generateUniqueCode(number) {
    var result = '';
    var characters = '0123456789';
    var charactersLength = characters.length;
    
    for ( var i = 0; i < number; i++ ) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}


function sendRequest(url, jsonData) {
    $.ajax({
        type: "POST",
        url: url,
        data: jsonData,
        success: function (result) {
          console.log(result);
          location.reload()
        }
      });
    
      return true
}

function generateFlashUrl() {
    var randomCode = generateUniqueCode(6);
    var url = "http://localhost:5454/";

    url += randomCode + '/flashMessage';

    return url;
}