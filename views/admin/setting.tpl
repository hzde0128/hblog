
        <div class="layui-fluid layui-col-md-offset2 layui-col-md10">
            <div class="layui-row">
                <form class="layui-form" action="/admin/setting" method="post">
                    <div class="layui-form-item">
                        <label for="L_repass" class="layui-form-label">网站标题</label>
                        <div class="layui-input-inline">
                            <input type="password" id="L_repass" name="oldpass" required="" lay-verify="required" autocomplete="off" class="layui-input">
                        </div>
                    </div>
                    <div class="layui-form-item">
                        <label for="L_pass" class="layui-form-label">URL设置</label>
                        <div class="layui-input-inline">
                            <input type="password" id="L_pass" name="newpass" required="" lay-verify="required" autocomplete="off" class="layui-input">
                            <div class="layui-form-mid layui-word-aux"></div>
                        </div>
                    <div class="layui-form-item">
                        <label for="L_repass" class="layui-form-label">版权信息</label>
                        <div class="layui-input-inline">
                            <input type="password" id="L_repass" name="repass" required="" lay-verify="required" autocomplete="off" class="layui-input">
                        </div>
                    </div>
                    <div class="layui-form-item">
                        <label for="L_repass" class="layui-form-label"></label>
                        <input type="submit" class="layui-btn" lay-filter="save" value="确认修改">
                    </div>
                    <div class="layui-form-item">
                        {{ if .errMsg }}
                            <span class="x-red">{{ .errMsg }}</span>
                        {{ end }}
                    </div>
                </form>
            </div>
        </div>
        <script>layui.use(['form', 'layer'],
            function() {
                $ = layui.jquery;
                var form = layui.form,
                layer = layui.layer;

                //监听提交
                form.on('submit(save)',
                function(data) {
                    console.log(data);
                    //发异步，把数据提交给php
                    layer.alert("修改成功", {
                        icon: 6
                    },
                    function() {
                        // 获得frame索引
                        var index = parent.layer.getFrameIndex(window.name);
                        //关闭当前frame
                        parent.layer.close(index);
                    });
                    return false;
                });

            });</script>
    </body>

</html>