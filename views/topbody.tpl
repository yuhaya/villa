<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <title>{{.AppName}}</title>
    <style type="text/css">
        *{
            padding: 0px;
            margin: 0px;
        }
        body{
            background-color: #000000;
        }
        #main{
            width: 100%;
        }
        #login_out{
            width: 80px;
            height: 25px;
            line-height: 25px;
            border: 0px;
            color: #fff;
            background-color: #000000;
            cursor: pointer;
        }
    </style>
</head>
<body>
<table id="main">
    <tr>
        <td style="width:180px;color: #ffffff;text-align: center">
            <h2>Villa Manager</h2>
            <p>ID : <span>{{.Name}}</span></p>
        </td>
        <td>
            &nbsp;&nbsp;
        </td>
        <td style="width: 100px;text-align: right">
            <input type="button" value="LoginOut" id="login_out"/>&nbsp;&nbsp;&nbsp;&nbsp;
        </td>
    </tr>
</table>
<script src="/static/js/lib/jquery-1.9.1.min.js"></script>
<script type="application/javascript">
    var code = {{.Code}};
    if(!code){
        window.top.alert("系统异常！请重新登录！");
    }
    $("#login_out").click(function(){
        $.ajax({
            type:"post",
            url:"/loginout",
            dataType:"json",
            success:function(data){
                if(data.Code){
                    window.top.location.href = "{{urlfor "UserController.LoginPage"}}";
                }else{
                    window.top.alert("系统异常!");
                }
            }
        })
    });
</script>
</body>
</html>