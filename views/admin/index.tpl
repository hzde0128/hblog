{{ template "header.tpl" . }}
    <body class="index">
        <!-- 顶部开始 -->
        <div class="container">
            <div class="logo">
                <a href="/admin/">个人博客管理系统</a></div>
            <div class="left_open">
                <a><i title="展开左侧栏" class="iconfont">&#xe699;</i></a>
            </div>
            <ul class="layui-nav right" lay-filter="">
                <li class="layui-nav-item">
                    <a href="javascript:;">{{ .user }}</a>
                    <dl class="layui-nav-child">
                        <!-- 二级菜单 -->
                        <dd>
                            <a onclick="xadmin.open('个人信息','/admin/user')">修改密码</a></dd>
                        <dd>
                            <a href="/admin/logout">退出</a></dd>
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
                                <a onclick='xadmin.add_tab("文章列表","/admin/blog/")'>
                                    <i class="iconfont">&#xe699;</i>
                                    <cite>文章列表</cite></a>
                            </li>
                            <li>
                                <a onclick='xadmin.add_tab("添加文章","/admin/blog/add")'>
                                    <i class="iconfont">&#xe69e;</i>
                                    <cite>添加文章</cite></a>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a onclick="xadmin.add_tab('分类管理','/admin/category/')">
                            <i class="iconfont left-nav-li">&#xe723;</i>
                            <cite>分类管理</cite>
                        </a>
                    </li>
                    <li>
                        <a onclick="xadmin.add_tab('标签管理','/admin/tag/')">
                            <i class="iconfont left-nav-li">&#xe6f4;</i>
                            <cite>标签管理</cite>
                        </a>
                    </li>
                    <li>
                        <a href="javascript:;">
                            <i class="iconfont left-nav-li" lay-tips="系统设置">&#xe6ae;</i>
                            <cite>系统设置</cite>
                            <i class="iconfont nav_right">&#xe6ae;</i></a>
                        <ul class="sub-menu">
                            <li>
                                <a onclick="xadmin.add_tab('系统设置','/admin/settings')">
                                    <i class="iconfont">&#xe6ae;</i>
                                    <cite>系统设置</cite></a>
                            </li>
                            <li>
                                <a onclick="xadmin.add_tab('修改密码','/admin/user')">
                                    <i class="iconfont">&#xe6b8;</i>
                                    <cite>修改密码</cite></a>
                            </li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
        <!-- <div class="x-slide_left"></div> -->
        <!-- 左侧菜单结束 -->
        <!-- 右侧主体开始 -->
        <div class="page-content">
            <div class="layui-tab tab" lay-filter="xbs_tab" lay-allowclose="false">
                <ul class="layui-tab-title">
                    <li class="home">
                        <i class="layui-icon">&#xe68e;</i>我的桌面</li></ul>
                <div class="layui-unselect layui-form-select layui-form-selected" id="tab_right">
                    <dl>
                        <dd data-type="this">关闭当前</dd>
                        <dd data-type="other">关闭其它</dd>
                        <dd data-type="all">关闭全部</dd></dl>
                </div>
                <div class="layui-tab-content">
                    <div class="layui-tab-item layui-show">

                    </div>
                </div>
                <div id="tab_show"></div>
            </div>
        </div>
        <div class="page-content-bg"></div>
        <style id="theme_style"></style>
        <!-- 右侧主体结束 -->
        <hr class="hr15">
        <hr class="hr15">
        <div class="layui-fluid">
            <div class="layui-row layui-col-space15">
                <div class="layui-col-md10 layui-col-md-offset2">
                    <div class="layui-card">
                        <div class="layui-card-body ">
                            <blockquote class="layui-elem-quote">欢迎管理员：
                                <span class="x-red">{{ .user }}</span>！当前时间:2020-03-17 10:50:53
                            </blockquote>
                        </div>
                    </div>
                </div>

                <div class="layui-col-md10 layui-col-md-offset2">
                    <div class="layui-card">
                        <div class="layui-card-header">系统信息</div>
                        <div class="layui-card-body ">
                            <table class="layui-table">
                                <tbody>
                                <tr>
                                    <th>软件版本</th>
                                    <td>{{ .version }}</td></tr>
                                <tr>
                                    <th>服务器地址</th>
                                    <td>127.0.0.1</td></tr>
                                <tr>
                                    <th>运行环境</th>
                                    <td>操作系统：{{ .os }} ({{ .arch }})，CPU数量：{{ .cpuNum }}</td></tr>
                                <tr>
                                    <th>Go版本</th>
                                    <td>{{ .goVer }}</td></tr>
                                <tr>
                                    <th>上传附件限制</th>
                                    <td>2M</td></tr>
                                <tr>
                                    <th>执行时间限制</th>
                                    <td>30s</td></tr>
                                <tr>
                                    <th>剩余空间</th>
                                    <td>86015.2M</td></tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
                <style id="welcome_style"></style>
            </div>
        </div>
        </div>
        <!-- 中部结束 -->
    </body>

</html>