<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <title>{{.AppName}}</title>
    <style>
        *{
            margin: 0px;
            padding: 0px;
        }
        body{
            background-color: #000000;
        }
        #list{
            color: #ffffff;
        }
        #list td{
            cursor: pointer;
            padding: 5px 10px;
        }
        #list td:hover{
            color: blue;
        }

    </style>
</head>
<body>

<table id="list">
    <tr>
        <td>
            用户管理
        </td>
    </tr>
    <tr>
        <td>
            目的地管理
        </td>
    </tr>
    <tr>
        <td>
            产品管理
        </td>
    </tr>
    <tr>
        <td>
            海报管理
        </td>
    </tr>
    <tr>
        <td>
            订单管理
        </td>
    </tr>
    <tr>
        <td data-url="/partner/list">
            合作方管理
        </td>
    </tr>
    <tr>
        <td>
            促销管理
        </td>
    </tr>
    <tr>
        <td>
            汇率管理
        </td>
    </tr>
    <tr>
        <td>
            标签管理
        </td>
    </tr>
    <tr>
        <td>
            账户管理
        </td>
    </tr>
    <tr>
        <td>
            系统配置
        </td>
    </tr>

</table>
<script type="application/javascript" src="/static/js/lib/jquery-1.9.1.min.js"></script>
<script>
    $("#list td").click(function(){
        var url = $(this).attr("data-url");
        window.top.frames["right"].location.href = url;
    })
</script>
</body>
</html>