<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
	<meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
	<title>{{.AppName}}</title>
	<link rel="stylesheet" href="/static/css/lib/common.css" />
	<link rel="stylesheet" href="/static/css/user/login.css" />
</head>
<body>
	<table width="80%" height="500px">
		<tr>
			<td width="70%">
				<h3 id="error">
					用户名不可以为空
				</h3>
			</td>
			<td width="30%" >
				<form id="form_data" action="/login" method="post" >
					<table id="login_box" width="260px" height="230px">
						<tr>
							<td colspan="2">
								<h2>Admin Login</h2>
							</td>
						</tr>
						<tr>
							<td>username:</td>
							<td align="right">
								<input type="text" name="username" />
							</td>
						</tr>
						<tr>
							<td>password:</td>
							<td align="right">
								<input type="password" name="password" />
							</td>
						</tr>
						<tr>
							<td colspan="2" align="left">
								<input id="login" type="button" value="submit" />
							</td>
						</tr>
					</table>
				</form>
			</td>
		</tr>
	</table>
	<div id="footer">
		{{.AppVer}}
	</div>
</body>
<script src="/static/js/lib/jquery-1.9.1.min.js"></script>
<script src="/static/js/user/login.js"></script>
</html>