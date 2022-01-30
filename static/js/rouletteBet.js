$(document).ready(function() {
    $('#bet-button-black').click(function () {
        var jsonData = {
            "test": "userId",
            "color": 3
        };

        $.ajax({
            type: "POST",
            url: "http://localhost:5454/addUser",
            data: jsonData,
            success: function (result) {
                console.log(result)
            }
          });
    });
    $('#bet-button-green').click(function () {
        var jsonData = {
            "test": "userId",
            "color": 2
        };

        $.ajax({
            type: "POST",
            url: "http://localhost:5454/addUser",
            data: jsonData,
            success: function (result) {
                console.log(result)
            }
          }); 
    });
    $('#bet-button-red').click(function () {
        var jsonData = {
            "test": "userId",
            "color": 1
        };

        $.ajax({
            type: "POST",
            url: "http://localhost:5454/addUser",
            data: jsonData,
            success: function (result) {
                console.log(result)
            }
          });
    });
});