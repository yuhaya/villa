$(function(){
	$("#login").click(function(){
		if(check()){
			$("#form_data").submit();
		}
	});
	function check(){
		if($("input[name='username']").val() == ""){
			$("#error").text("用户名不可以为空！");
			return false;
		}
		if($("input[name='password']").val() == ""){
			$("#error").text("密码不可以为空！");
			return false;
		}
		return true;
	}
})