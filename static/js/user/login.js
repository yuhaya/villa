$(function(){
	$("#login").click(function(){
		var name = $("input[name='username']").val();
		var pwd = $("input[name='password']").val();
		if($("input[name='username']").val() == ""){
			$("#error").text("用户名不可以为空！");
			return false;
		}
		if($("input[name='password']").val() == ""){
			$("#error").text("密码不可以为空！");
			return false;
		}
		var hashpwd = CryptoJS.SHA1(pwd).toString();
		var data = {
			username:name,
			password:hashpwd
		}
		$.ajax({
			type:"POST",
			url:$("#form_data").attr("action"),
			data:data,
			dataType:"json",
			success:function(data){
				if(data.code){
					location.href = "/";
				}else{
					$("#error").text(data.msg);
				}
			}
		})
	});

})