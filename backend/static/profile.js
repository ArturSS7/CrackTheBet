

$(document).ready((function () {
    get_history();
}));

function get_history(){
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
        function generateTableHead(table, data) {
            let thead = table.createTHead();
            let row = thead.insertRow();
            for (let key of data) {
              let th = document.createElement("th");
              let text = document.createTextNode(key);
              th.appendChild(text);
              row.appendChild(th);
            }
          }
        
        function generateTable(table, data) {
            for (let element of data) {
              let row = table.insertRow();
              for (key in element) {
                let cell = row.insertCell();
                let text = document.createTextNode(element[key]);
                cell.appendChild(text);
              }
            }
          }
        if (this.readyState == 4 && this.status == 200) {
            let data = jQuery.parseJSON(this.responseText)
            var bets = data.bets;
            let table = document.querySelector("table");
	    var new_bets = bets.map(function(bet){

		var new_bet = {};
		new_bet['Date'] = bet['bruh_time']
		new_bet['Player1'] = bet['player_1'];
		new_bet['Player2'] = bet['player_2'];
		new_bet['Bet Player'] = bet['bet_player'];
		new_bet['Odds'] = bet['odds'];
    new_bet['Bet'] = bet['amount'];
		new_bet['Prize'] = bet['prize'];
    new_bet['Status'] = bet['status'];
    console.log(typeof new_bet['Prize']);
    if (new_bet['Prize'] == -1 || new_bet['Prize'] == '-1') {
      new_bet['Prize'] = "";
    }

    if (new_bet['Status'] == 'not_processed') {
      new_bet['Status'] = 'Not Processed';
    }


		return new_bet;
});
console.log(new_bets);
            data = Object.keys(new_bets[0]);
            generateTableHead(table, data);
            generateTable(table, new_bets);
        }
    }

    xhttp.open("GET", "/api/bets", false);
    xhttp.send();
}


  
