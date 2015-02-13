define(function(require, exports, module) {

	require("validate_helper");
	
    exports.init = function(){
        init_style();
    }

    function init_style(){
    	$("#partner_submit").click(function(){
    		validate()
    	})
    }

    function validate(){
    	var form =  $("#addPartner")
		form.validate(function($form, e){ alert("submitted") })
    }



});
