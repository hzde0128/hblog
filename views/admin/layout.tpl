<!DOCTYPE html>
<html class="x-admin-sm">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}-{{ .basic.Title }}-后台管理</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
    <link rel="stylesheet" href="/static/css/font.css">
    <link rel="stylesheet" href="/static/css/xadmin.css">
    <script src="/static/lib/layui/layui.js" charset="utf-8"></script>
    <script type="text/javascript" src="/static/js/xadmin.js"></script>
    <!-- 让IE8/9支持媒体查询，从而兼容栅格 -->
    <!--[if lt IE 9]>
    <script src="https://cdn.staticfile.org/html5shiv/r29/html5.min.js"></script>
    <script src="https://cdn.staticfile.org/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>

<body class="index">
<!-- 顶部开始 -->
<div class="container">
    <div class="logo">
        <a href="/admin/">{{ .basic.Title }}</a></div>
    <ul class="layui-nav right" lay-filter="">
        <li class="layui-nav-item">
            <a href="javascript:;">{{ .user }}</a>
            <dl class="layui-nav-child">
                <!-- 二级菜单 -->
                <dd><a href="/admin/user">修改密码</a></dd>
                <dd><a href="/admin/logout">退出</a></dd>
            </dl>
        </li>
        <li class="layui-nav-item to-index">
            <a href="/">前台首页</a></li>
    </ul>
</div>
<!-- 顶部结束 -->

<!-- 中部开始 -->
<!-- 左侧菜单开始 -->
<div class="left-nav">
    <div id="side-nav">
        <ul id="nav">
            <li>
                <a href="javascript:;">
                    <i class="iconfont left-nav-li" lay-tips="文章管理">&#xe705;</i>
                    <cite>文章管理</cite>
                    <i class="iconfont nav_right">&#xe697;</i></a>
                <ul class="sub-menu">
                    <li>
                        <div style="position: relative">
                            <i class="iconfont" style="position: absolute; left: 35px; line-height: 40px;">&#xe699;</i>
                            <a style="padding-left: 60px" href="/admin/blog/">文章列表</a>
                        </div>
                    </li>
                    <li>
                        <div style="position: relative">
                            <i class="iconfont" style="position: absolute; left: 35px; line-height: 40px;">&#xe69e;</i>
                            <a style="padding-left: 60px" href="/admin/blog/add">发表文章</a>
                        </div>
                    </li>
                </ul>
            </li>
            <li>
                <div style="position: relative">
                    <i class="iconfont" style="position: absolute; left: 20px; line-height: 40px;">&#xe723;</i>
                    <a style="padding-left: 42px" href="/admin/category/">分类管理</a>
                </div>
            </li>
            <li>
                <div style="position: relative">
                    <i class="iconfont" style="position: absolute; left: 20px; line-height: 40px;">&#xe6f4;</i>
                    <a style="padding-left: 42px" href="/admin/tag/">标签管理</a>
                </div>
            </li>
            <li>
                <a href="javascript:;">
                    <i class="iconfont left-nav-li" lay-tips="系统设置">&#xe6ae;</i>
                    <cite>系统设置</cite>
                    <i class="iconfont nav_right">&#xe6ae;</i></a>
                <ul class="sub-menu">
                    <li>
                        <div style="position: relative">
                            <i class="iconfont" style="position: absolute; left: 35px; line-height: 40px;">&#xe6ae;</i>
                            <a style="padding-left: 60px" href="/admin/setting">系统设置</a>
                        </div>
                    </li>
                    <li>
                        <div style="position: relative">
                            <i class="iconfont" style="position: absolute; left: 35px; line-height: 40px;">&#xe6b8;</i>
                            <a style="padding-left: 60px" href="/admin/user">修改密码</a>
                        </div>
                    </li>
                </ul>
            </li>
        </ul>
    </div>
</div>
<!-- <div class="x-slide_left"></div> -->
<!-- 左侧菜单结束 -->


<!-- 右侧主体开始 -->
<div class="page-content-bg"></div>
<style id="theme_style"></style>
<!-- 右侧主体结束 -->

{{ .LayoutContent }}

</body>

</html>