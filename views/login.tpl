<!doctype html>
<html  class="x-admin-sm">
<head>
	<meta charset="UTF-8">
	<title>个人博客系统-后台登录</title>
	<meta name="renderer" content="webkit|ie-comp|ie-stand">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
    <meta http-equiv="Cache-Control" content="no-siteapp" />
    <link rel="stylesheet" href="/static/css/font.css">
    <link rel="stylesheet" href="/static/css/login.css">
    <link rel="stylesheet" href="/static/css/xadmin.css">
    <script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
    <script src="/static/lib/layui/layui.js" charset="utf-8"></script>
    <!--[if lt IE 9]>
      <script src="https://cdn.staticfile.org/html5shiv/r29/html5.min.js"></script>
      <script src="https://cdn.staticfile.org/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body class="login-bg">
    
    <div class="login layui-anim layui-anim-up">
        <div class="message">个人博客系统-后台登录</div>
        <div id="darkbannerwrap"></div>
        
        <form method="post" class="layui-form">
            {{/* 错误信息 */}}
            {{ if .errMsg }}
                <span class="x-red">{{ .errMsg }}!</span>
                <hr class="hr15">
            {{ end }}
            <div style="position: relative;">
                <i class="iconfont" style="position: absolute; line-height: 50px; right: 12px;">&#xe6b8;</i>
                <input name="username" placeholder="用户名"  type="text" lay-verify="required" class="layui-input" value="{{ .username }}">
                <hr class="hr15">

            </div>
            <div style="position: relative">
                <i class="iconfont" style="position: absolute; line-height: 50px; right: 12px;">&#xe82b;</i>
                <input name="password" lay-verify="required" placeholder="密码"  type="password" class="layui-input">
                <hr class="hr15">
            </div>
            <div>
                <input style="display: inline-block !important;margin-right: 5px" type="checkbox" name="remember" value="on" checked="checked">记住密码
                <hr class="hr15">
            </div>

            <input value="登录" lay-submit lay-filter="login" style="width:100%;" type="submit">
            <hr class="hr20" >
        </form>
    </div>

    <!-- 底部结束 -->
</body>
</html>