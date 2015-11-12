$(function () {		
	$('#formbutton').click(function () {  
		var mydata = { idnumber: $('#idnumber').val() };							
		var request = $.ajax ({
			type: "POST",								
			url: '/Manual/Capture',
			dataType: 'json',
			async: true,					
			data:  mydata,					
		});

		request.done(function(msg) {
			$("#mydiv").removeClass('hide')
			addRow(msg.ID);
			clearIdNumberField();
		});				
				
		request.fail(function(jqZHR, textStatus) {
			alert('Request Failed: ' + textStatus);			
        });
	});
	
	function clearIdNumberField() {
		$('#idnumber').val('');
	}
	
	var files;
	$('input[type=file]').on('change', prepareUpload);
	
	function prepareUpload(event) {
		files = event.target.files;		
	}
	
	$('#fileuploadbutton').click(function () {	
		var data = new FormData();
		$.each(files, function(key, value) {			
			data.append('idnumberlist', value);
		});	
				
		var request = $.ajax ({
			type: "POST",						
			url: '/File/HandleUpload',
			cache: false,
			dataType: 'json',
			async: true,					
			data:  data,
			processData: false,
			contentType: false,
		});

		request.done(function(msg) {
			$("#mydiv").removeClass('hide')		

			$.each(msg.List, function (key, value) {
				addRow(value);
			});		
		});
				
		request.fail(function(jqZHR, textStatus) {
			alert('Request Failed: ' + textStatus);
        });
	});
		
	var i = 1;
	function addRow(data) {
		var html = "<tr>"  
					+ "<td>"+data.IDNumber+"</td>"
					+ "<td>"+data.DateOfBirth+"</td>"
					+ "<td>"+data.Gender+"</td>"
					+ "<td>"+data.Citizenship+"</td>"
					+ "<td>"+data.Vadility+"</td>"
				+ "</tr>";
		$('#my_table > tbody > tr').eq(i - 1).after(html).fadeIn(400, null);
		i++;
	}
});