$(document).ready((function(){
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
	$('#button_password_rec').click(function(){
		$('.block-popup_login').fadeOut();
		$('.block-popup_recovery').fadeIn();
	})
	$('.block-popup_recovery span').click(function(){
		$('.block-popup_recovery, .overlay').fadeOut();
	})

	$('.profile_button').click(function(){
		window.location = '/profile';
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
            	//console.log(response.responseJSON.error);
            	document.getElementsByClassName('bet_error')[0].innerText = response.responseJSON.error;
            	document.getElementsByClassName('bet_error')[0].style.display = 'block';
            	$('.bet input[type="submit" i]').attr('disabled',false);
            }
        });
        //console.log('disabled');
        
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
                //console.log("success", data);
                $('.block-popup_reg').fadeOut();
                $('.reg_success').fadeIn();
                $('.reg_form input[type="submit" i]').attr('disabled',false);
            },
            error: function(response, status, error){
            	//console.log(response.responseJSON.error);
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
                //console.log("success", data);
                window.location = '/';
            },
            error: function(response, status, error){
            	//console.log(response.responseJSON.error);
            	document.getElementsByClassName('log_error')[0].innerText = response.responseJSON.error;
            	document.getElementsByClassName('log_error')[0].style.display = 'block';
            	$('.login_form input[type="submit" i]').attr('disabled',false);
            }
        });
        
        event.preventDefault();
    });

    $('.recovery_form').submit(function(event) {
    	$('.recovery_form input[type="submit" i]').attr('disabled',true);
        // get the form data
        // there are many ways to get this data using jQuery (you can use the class or id also)
        var formData = {
            'email'              : $('.rec_e').val(),
        };

        // process the form
        $.ajax({
            type        : 'POST', // define the type of HTTP verb we want to use (POST for our form)
            url         : 'forgot-password', // the url where we want to POST
            data        : formData, // our data object
            encode      : true,
            timeout     : 0,
            success: function(data) {
                //console.log("success", data);
                window.location = '/';
            }
        });
        
        event.preventDefault();
    });


}));