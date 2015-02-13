	<table width="80%" height="500px">
		<tr>
			<td width="70%">
				<h3 id="error">
					
				</h3>
			</td>
			<td width="30%" >
				<form id="form_data" action="{{urlfor "UserController.LoginSubmit"}}" method="post">
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
