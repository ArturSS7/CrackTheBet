$(document).ready((function(){
	//let i = setInterval(function(){update_events()}, 2000);
	get_events();
	$('.login_button').click(function(){
		$('.block-popup_login, .overlay').fadeIn();
	})
	$('.block-popup_login span').click(function(){
		$('.block-popup_login, .overlay').fadeOut();
	})
	$('.reg_button').click(function(){
		$('.block-popup_reg, .overlay').fadeIn();
	})
	$('.block-popup_reg span').click(function(){
		$('.block-popup_reg, .overlay').fadeOut();
	})
	$('.reg_success span').click(function(){
		$('.reg_success, .overlay').fadeOut();
	})
	$('.bet span').click(function(){
		$('.bet, .overlay').fadeOut();
		$('.bet_error')[0].innerText = '';
	})

	$('.logout_button').click(function(){
		$.ajax({
            type        : 'GET', // define the type of HTTP verb we want to use (POST for our form)
            url         : 'logout',
            success: function(data){
            	location.reload();
            }
        });
	})

	$('.bet_form').submit(function(event) {
		$('.bet input[type="submit" i]').attr('disabled',true);
        // get the form data
        // there are many ways to get this data using jQuery (you can use the class or id also)
        var formData = {
            'amount'              : $('.bet_amount').val(),
            'id'             : $('.bet_descr_id')[0].id,
            'player'         : $('.bet_descr_team')[0].id
        };

        // process the form
        $.ajax({
            type        : 'POST', // define the type of HTTP verb we want to use (POST for our form)
            url         : 'api/bet', // the url where we want to POST
            data        : formData, // our data object
            dataType    : 'json', // what type of data do we expect back from the server
            encode      : true,
            success: function(data) {
                location.reload();
                //$('.bet, .overlay').fadeOut();
            },
            error: function(response, status, error){
            	console.log(response.responseJSON.error);
            	document.getElementsByClassName('bet_error')[0].innerText = response.responseJSON.error;
            	document.getElementsByClassName('bet_error')[0].style.display = 'block';
            	$('.bet input[type="submit" i]').attr('disabled',false);
            }
        });
        console.log('disabled');
        
        event.preventDefault();
    });

	$('.reg_form').submit(function(event) {
		$('.reg_form input[type="submit" i]').attr('disabled',true);
        // get the form data
        // there are many ways to get this data using jQuery (you can use the class or id also)
        var formData = {
            'username'              : $('.reg_u').val(),
            'email'             : $('.reg_e').val(),
            'password'    : $('.reg_p').val(),
            'password-repeat'    : $('.reg_pr').val()
        };

        // process the form
        $.ajax({
            type        : 'POST', // define the type of HTTP verb we want to use (POST for our form)
            url         : 'register', // the url where we want to POST
            data        : formData, // our data object
            dataType    : 'json', // what type of data do we expect back from the server
            encode      : true,
            success: function(data) {
                console.log("success", data);
                $('.block-popup_reg').fadeOut();
                $('.reg_success').fadeIn();
                $('.reg_form input[type="submit" i]').attr('disabled',false);
            },
            error: function(response, status, error){
            	console.log(response.responseJSON.error);
            	document.getElementsByClassName('reg_error')[0].innerText = response.responseJSON.error;
            	document.getElementsByClassName('reg_error')[0].style.display = 'block';
            	$('.reg_form input[type="submit" i]').attr('disabled',false);
            }
        });
        
        event.preventDefault();
    });

    $('.login_form').submit(function(event) {
    	$('.login_form input[type="submit" i]').attr('disabled',true);
        // get the form data
        // there are many ways to get this data using jQuery (you can use the class or id also)
        var formData = {
            'username'              : $('.log_u').val(),
            'password'    : $('.log_p').val()
        };

        // process the form
        $.ajax({
            type        : 'POST', // define the type of HTTP verb we want to use (POST for our form)
            url         : 'login', // the url where we want to POST
            data        : formData, // our data object
            dataType    : 'json', // what type of data do we expect back from the server
            encode      : true,
            success: function(data) {
                console.log("success", data);
                location.reload();
            },
            error: function(response, status, error){
            	console.log(response.responseJSON.error);
            	document.getElementsByClassName('log_error')[0].innerText = response.responseJSON.error;
            	document.getElementsByClassName('log_error')[0].style.display = 'block';
            	$('.login_form input[type="submit" i]').attr('disabled',false);
            }
        });
        
        event.preventDefault();
    });


}));



function get_events(){
	$(".events").empty();
	let xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			let data = jQuery.parseJSON(this.responseText).leagues;
			console.log(data);
			$.each(data, function(key, item){
				let league_cont = document.createElement("div");
				league_cont.className = "container league"
				let title = document.createElement("div");
				title.innerHTML = "<p>" + item.league_name + "</p>";
				title.className = "title"
				league_cont.appendChild(title);
				console.log(item.league_name + " : " + item.events);
				$.each(item.events, function(key, e){
					let match = document.createElement("div");
					match.className = "match";
					match.id = e.flashscore_id;
					let time = document.createElement("div");
					time.className = 'time';
					date = new Date(e.time * 1000);
					time.innerText = date.getHours()+":"+String(date.getMinutes()).padStart(2,'0');
					let t1 = document.createElement("div");
					t1.className = 'team1';
					t1.innerText = e.player_1;
					let t2 = document.createElement("div");
					t2.className = 'team2';
					t2.innerText = e.player_2;
					let odd1 = document.createElement("div");
					odd1.className = "odds odd1";
					odd1.innerText = String(e.odds_1).split('.')[0]+'.'+String(e.odds_1).split('.')[0].padEnd(2,'0');
					odd1.onclick = function(){
						$('.bet, .overlay').fadeIn();
						document.getElementsByClassName('bet_descr_id')[0].id = match.id;
						document.getElementsByClassName('bet_descr_team')[0].id = 1;
						document.getElementsByClassName('bet_descr_team')[0].innerText = t1.innerText;
					};
					let odd2 = document.createElement("div");
					odd2.className = 'odds odd2';
					odd2.innerText = String(e.odds_2).split('.')[0]+'.'+String(e.odds_2).split('.')[0].padEnd(2,'0');
					odd2.onclick = function(){
						$('.bet, .overlay').fadeIn();
						document.getElementsByClassName('bet_descr_id')[0].id = match.id;
						document.getElementsByClassName('bet_descr_team')[0].id = 2;
						document.getElementsByClassName('bet_descr_team')[0].innerText = t2.innerText;
					};
					let tochki = document.createElement("div");
					tochki.className = 'dvoetochie';
					tochki.innerText = ":";
					match.appendChild(time);
					match.appendChild(odd1);
					match.appendChild(t1);
					match.appendChild(tochki);
					match.appendChild(odd2);
					match.appendChild(t2);
					league_cont.appendChild(match);
				});
				$(".events").append(league_cont);
			});
		}
	};
	xhttp.open("GET", "/api/events", true);
	xhttp.send();
}

function update_events(){
	let xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			let data = jQuery.parseJSON(this.responseText).leagues;
			$.each(data, function(key1, league){
				$.each(league.events, function(key2, match) {
					match_dom = $(`#${match.flashscore_id}`);
					match_dom.children('.odds').css("opacity", "1.0").animate({opacity: 0}, 1200, function(){
						match_dom.children('.odds').css("visibility", "hidden");
					});
					match_dom.children('.odd1')[0].innerText = match.odds_1;
					match_dom.children('.odd2')[0].innerText = match.odds_2;
					match_dom.children('.odds').css({opacity: 0.0, visibility: "visible"}).animate({opacity: 1}, 1200);
				});
			});
		}
	}
	xhttp.open("GET", "/api/events", true);
	xhttp.send();
}