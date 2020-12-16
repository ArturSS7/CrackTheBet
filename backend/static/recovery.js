$(document).ready((function(){
        $('.recovery__form').submit(function(event) {
                document.getElementsByClassName('rec_error')[0].style.display = 'none';
                document.getElementsByClassName('rec_success')[0].style.display = 'none';
        		$('.recovery__form input[type="submit" i]').attr('disabled',true);
                // get the form data
                // there are many ways to get this data using jQuery (you can use the class or id also)
                var formData = {
                    'password'              : $('.pass').val(),
                    'repeat-password'             : $('.pass_rep').val(),
                    'token'    : $('.token').val()
                };

                // process the form
                $.ajax({
                    type        : 'POST', // define the type of HTTP verb we want to use (POST for our form)
                    url         : 'recovery', // the url where we want to POST
                    data        : formData, // our data object
                    dataType    : 'json', // what type of data do we expect back from the server
                    encode      : true,
                    success: function(data) {
                        //console.log("success", data);
                        document.getElementsByClassName('rec_success')[0].innerText = "password has been changed";
                        document.getElementsByClassName('rec_success')[0].style.display = 'block';
                        $('.recovery__form input[type="submit" i]').attr('disabled',false);
                    },
                    error: function(response, status, error){
                    	//console.log(response.responseJSON.error);
                    	document.getElementsByClassName('rec_error')[0].innerText = response.responseJSON.error;
                    	document.getElementsByClassName('rec_error')[0].style.display = 'block';
                    	$('.recovery__form input[type="submit" i]').attr('disabled',false);
                    }
                });
                
                event.preventDefault();
            });
}));