$(document).ready((function(){
	get_events();
	let i = setInterval(function(){update_events()}, 15000);
}));

function get_events(){
	$(".events").empty();
	let xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			let data = jQuery.parseJSON(this.responseText).leagues;
			//console.log(data);
			$.each(data, function(key, item){
				counter = 0;
				let league_cont = document.createElement("div");
				league_cont.className = "container league"
				let title = document.createElement("div");
				title.innerHTML = "<p>" + item.league_name + "</p>";
				title.className = "title"
				league_cont.appendChild(title);
				//console.log(item.league_name + " : " + item.events);
				$.each(item.events, function(key, e){
					if (e.status != 'finished'){
						counter ++;
						let match = document.createElement("div");
						if (e.status == 'active'){
							match.className = "match is_live";
						}
						else{
							match.className = "match";
						}
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
						let live = document.createElement('div');
						live.className = 'live';
						live.innerText = 'Live';
						match.appendChild(time);
						match.appendChild(odd1);
						match.appendChild(t1);
						match.appendChild(tochki);
						match.appendChild(live);
						match.appendChild(odd2);
						match.appendChild(t2);
						league_cont.appendChild(match);
					}
				});
				if (counter != 0){
					$(".events").append(league_cont);
				}
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
					if (match.status != 'finished'){
						if (match.status == 'active'){
							match.className = 'match is_live';
						}
						else{
							match.className = 'match';
						}
						//console.log(match);
						match_dom = $(`#${match.flashscore_id}`);
						match_dom.children('.odds').css("opacity", "1.0").animate({opacity: 0}, 1200, function(){
							match_dom.children('.odds').css("visibility", "hidden");
						});
						match_dom.children('.odd1')[0].innerText = match.odds_1;
						match_dom.children('.odd2')[0].innerText = match.odds_2;
						match_dom.children('.odds').css({opacity: 0.0, visibility: "visible"}).animate({opacity: 1}, 1200);
					}
				});
			});
		}
	}
	xhttp.open("GET", "/api/events", true);
	xhttp.send();
}