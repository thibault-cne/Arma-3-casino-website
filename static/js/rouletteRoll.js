$(document).ready(function() {
	//setup multiple rows of colours, can also add and remove while spinning but overall this is easier.
	initWheel();
 
 	setInterval(rollWheel, 12 * 1000);
});

function initWheel(){
	var $wheel = $('.roulette-wrapper .wheel'),
  		row = "";
      
  row += "<div class='roulette-row'>";
  row += "  <div class='rcard red'>1<\/div>";
  row += "  <div class='rcard black'>14<\/div>";
  row += "  <div class='rcard red'>2<\/div>";
  row += "  <div class='rcard black'>13<\/div>";
  row += "  <div class='rcard red'>3<\/div>";
  row += "  <div class='rcard black'>12<\/div>";
  row += "  <div class='rcard red'>4<\/div>";
  row += "  <div class='rcard green'>0<\/div>";
  row += "  <div class='rcard black'>11<\/div>";
  row += "  <div class='rcard red'>5<\/div>";
  row += "  <div class='rcard black'>10<\/div>";
  row += "  <div class='rcard red'>6<\/div>";
  row += "  <div class='rcard black'>9<\/div>";
  row += "  <div class='rcard red'>7<\/div>";
  row += "  <div class='rcard black'>8<\/div>";
  row += "<\/div>";
  
	for(var x = 0; x < 29; x++){
  	$wheel.append(row);
  }
}

function spinWheel(roll){
  var $wheel = $('.roulette-wrapper .wheel'),
  		order = [0, 11, 5, 10, 6, 9, 7, 8, 1, 14, 2, 13, 3, 12, 4],
      position = order.indexOf(roll);
            
  //determine position where to land
  var rows = 12,
  		card = 75 + 3 * 2,
      landingPosition = (rows * 15 * card) + (position * card);
  	
  var randomize = Math.floor(Math.random() * 75) - (75/2);
    
  landingPosition = landingPosition + randomize;
    
  var object = {
		x: Math.floor(Math.random() * 50) / 100,
    y: Math.floor(Math.random() * 20) / 100
	};
  
  $wheel.css({
		'transition-timing-function':'cubic-bezier(0,'+ object.x +','+ object.y + ',1)',
		'transition-duration':'5s',
		'transform':'translate3d(-'+landingPosition+'px, 0px, 0px)'
	});
  
  setTimeout(function(){
		$wheel.css({
			'transition-timing-function':'',
			'transition-duration':'',
		});
    
    var resetTo = -(position * card + randomize);
		$wheel.css('transform', 'translate3d('+resetTo+'px, 0px, 0px)');
  }, 6000);
}

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}

function rollWheel() {
  var outcome = getRandomInt(15);
  var jsonData = {
    'result': outcome
  };

  $.ajax({
    type: "POST",
    url: "http://localhost:5454/roulette",
    data: jsonData,
    success: function (result) {
      console.log(result);

      endRoll();
    }
  });

  function endRoll() {
    var blackPlayers = $('.players-name-black');
    var redPlayers = $('.players-name-red');
    var greenPlayers = $('.players-name-green');

    blackPlayers.text("");
    redPlayers.text("");
    greenPlayers.text("");
  }

  spinWheel(outcome);
}