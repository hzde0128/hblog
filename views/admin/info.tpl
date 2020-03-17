{{ template "header.tpl" . }}
    <body>
        <div class="layui-fluid">
            <div class="layui-row layui-col-space15">
                <div class="layui-col-md12">
                    <div class="layui-card">
                        <div class="layui-card-body ">
                            <blockquote class="layui-elem-quote">欢迎管理员：
                                <span class="x-red">{{ .user }}</span>！当前时间:2020-03-17 10:50:53
                            </blockquote>
                        </div>
                    </div>
                </div>

                <div class="layui-col-md12">
                    <div class="layui-card">
                        <div class="layui-card-header">系统信息</div>
                        <div class="layui-card-body ">
                            <table class="layui-table">
                                <tbody>
                                    <tr>
                                        <th>软件版本</th>
                                        <td>v0.1.2</td></tr>
                                    <tr>
                                        <th>服务器地址</th>
                                        <td>127.0.0.1</td></tr>
                                    <tr>
                                        <th>操作系统</th>
                                        <td>Linux</td></tr>
                                    <tr>
                                        <th>运行环境</th>
                                        <td>Go 1.14.0</td></tr>
                                    <tr>
                                        <th>MYSQL版本</th>
                                        <td>5.5.53</td></tr>
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
    </body>
</html>