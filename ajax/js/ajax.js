		// On document ready;
		$(document).ready(
				function() {
					console.log("document ready!");
					// On click on element with class .relay;
					$(".divelem").click(
							function() {
								// get the element clicked
								var elem = $(this);
								console.log("Handler for .click() called. "
										+ elem.attr("id"));
								// Send id of clicked element and get response from server via ajax
								// Using the core $.ajax() method
								$.ajax({

									// The URL for the request
									url : "/ajax/post.html",

									// The data to send (will be converted to a query string)
									//data : {
									//	id : elem.attr("id"),
									//},
									data: JSON.stringify({
										Id: elem.attr("id"),
										Val: "test"
									}),

									// Whether this is a POST or GET request
									method : "POST",

									// The type of data we expect back
									dataType : "json",

									// Code to run if the request succeeds;
									// The response is passed to the function
									success : function(data) {
										$("#data").html(
														"status: ok, data: "
														+ data.Id);
										console.log("Success.");
									},

									// Code to run if the request fails;
									error : function() {
										$("#data").html(
												"Comm failed!");
										console.log("Error.");
									},

									// Code to run regardless of success or failure
									complete : function() {
										console.log("Complete.");
									}
								});
								// ajax transaction end
							});
				});
