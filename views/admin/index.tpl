
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
                                    <td>{{ .ip }}</td></tr>
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