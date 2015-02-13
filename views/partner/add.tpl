<form action="{{urlfor "PartnerController.AddSubmit"}}" method="post" name="addPartner" id="addPartner">
	<table class="add_table">
		<tr>
			<td>名称</td>
			<td>
				<div class="control-group">
			    	<input type="text" name="Name" id="Name" data-rules="required|max_length[10]" data-rules="Name" class="validate">
			    	<p class="help-block"></p>
			  	</div>
			</td>
		</tr>
		<tr>
			<td>
				目的地
			</td>
			<td>
				<div class="control-group">
					<select name="DestinationId" id="DestinationId" data-rules="required|integer|max_length[1]|max_length[2]" data-rules="DestinationId" class="validate">
						<option >选择目的地</option>
						<option value=""></option>
					</select>
				<p class="help-block"></p>
			  	</div>
			</td>
		</tr>
		<tr>
			<td>
				别墅管理联系人
			</td>
			<td>
				<input type="text" name="ManagerContact">
			</td>
		</tr>
		<tr>
			<td>
				别墅管理联系人电话
			</td>
			<td>
				<input type="text" name="ManagerTelephone">
			</td>
		</tr>
		<tr>
			<td>
				别墅管理联系人邮箱
			</td>
			<td>
				<input type="text" name="ManagerEmail">
			</td>
		</tr>
		<tr>
			<td>
				别墅预订联系人
			</td>
			<td>
				<input type="text" name="ReservationContact">
			</td>
		</tr>
		<tr>
			<td>
				别墅预订联系人电话
			</td>
			<td>
				<input type="text" name="ReservationTelephone">
			</td>
		</tr>
		<tr>
			<td>
				别墅预订联系人邮箱
			</td>
			<td>
				<input type="text" name="ReservationEmail">
			</td>
		</tr>
		<tr>
			<td>
				合作开始时间
			</td>
			<td>
				<input type="date" name="ContractStartDate">
			</td>
		</tr>
		<tr>
			<td>
				合作结束时间
			</td>
			<td>
				<input type="date" name="ContractEndDate">
			</td>
		</tr>
		<tr>
			<td>
				提成比例
			</td>
			<td>
				<input type="text" name="Commission">
			</td>
		</tr>
		<tr>
			<td>
				隶属集团
			</td>
			<td>
				<input type="text" name="MembershipGroup">
			</td>
		</tr>
		<tr>
			<td>
				状态
			</td>
			<td>
				<select name="State" id="">
					<option value="0">洽谈中</option>
					<option value="1">已合作</option>
					<option value="2">解除合作</option>
					<option value="3">未回复的</option>
					<option value="4">删除的</option>
				</select>
			</td>
		</tr>
		<tr>
			<td>
				备注
			</td>
			<td>
				<textarea name="Memo" id="" cols="30" rows="5"></textarea>
			</td>
		</tr>
		<tr>
			<td>
				联系地址
			</td>
			<td>
				<input type="text" name="Address">
			</td>
		</tr>
		<tr>
			<td>
				网址
			</td>
			<td>
				<input type="text" name="Website">
			</td>
		</tr>
		<tr>
			<td>
				别墅可用性查询链接
			</td>
			<td>
				<input type="text" name="AvailablityLink">
			</td>
		</tr>
		<tr>
			<td>
				可用性账户
			</td>
			<td>
				<input type="text" name="AvailablityUsername">
			</td>
		</tr>
		<tr>
			<td>
				可用性密码
			</td>
			<td>
				<input type="text" name="AvailablityPassword">
			</td>
		</tr>
		<tr>
			<td>
				
			</td>
			<td>
				<input type="button" value="取消" class="btn btn-default" onclick="history.go(-1);">
				<input type="submit" value="添加" id="partner_submit" class="btn btn-success">
			</td>
		</tr>
	</table>
</form>