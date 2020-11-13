$(document).ready((function(){
	//let i = setInterval(function(){update_events()}, 2000);
	update_events();
}));

function update_events(){
	$(".events").empty();
	console.log("updted");
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
					let time = document.createElement("div");
					time.className = 'time';
					date = new Date(e.time * 1000);
					time.innerText = date.getHours()+":"+String(date.getMinutes()).padStart(2,'0');
					let odd1 = document.createElement("div");
					odd1.className = "odds odd1";
					odd1.innerText = String(e.odds_1).split('.')[0]+'.'+String(e.odds_1).split('.')[0].padEnd(2,'0');
					let odd2 = document.createElement("div");
					odd2.className = 'odds odd2';
					odd2.innerText = String(e.odds_2).split('.')[0]+'.'+String(e.odds_2).split('.')[0].padEnd(2,'0');
					let tochki = document.createElement("div");
					tochki.className = 'dvoetochie';
					tochki.innerText = ":";
					let t1 = document.createElement("div");
					t1.className = 'team1';
					t1.innerText = e.player_1;
					let t2 = document.createElement("div");
					t2.className = 'team2';
					t2.innerText = e.player_2;
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