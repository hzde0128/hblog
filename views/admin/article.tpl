
        <div class="layui-fluid layui-col-md-offset2 layui-col-md10">
            <div class="layui-row">
                <form class="layui-form" action="/admin/blog/add" method="post">
                    <div class="layui-form-item">
                        <label for="username" class="layui-form-label">
                            <span class="x-red">*</span>文章标题</label>
                        <div class="layui-input-inline">
                            <input type="text" id="title" name="title" required="" lay-verify="required" autocomplete="off" class="layui-input"></div>
                    </div>
                    <div class="layui-form-item">
                        <label for="username" class="layui-form-label">
                            <span class="x-red">*</span>文章分类</label>
                        <div class="layui-input-inline">
                            <select id="category" name="category" class="valid">
                                {{ range .category }}
                                <option value="{{ .CategoryName }}">{{ .CategoryName }}</option>
                                {{ end }}
                            </select>
                        </div>
                    </div>
                    <div class="layui-form-item">
                        <label for="username" class="layui-form-label">
                            <span class="x-red">*</span>文章标签</label>
                        <div class="layui-input-inline">
                                {{ range .tag }}
                                <input name="tag_id" value="{{ .Id }}" type="hidden" hidden="hidden">
                                <input type="checkbox" name="tag[]" title="{{ .TagName }}">
                                {{ end }}
                        </div>
                    </div>
        <div class="layui-form-item layui-form-text">
            <label for="desc" class="layui-form-label">
                <span class="x-red">*</span>文章内容</label>
            <div class="layui-input-block">
                <textarea placeholder="请输入内容" id="desc" name="content" class="layui-textarea"></textarea>
            </div>
        </div>
        <div class="layui-form-item">
            <label for="L_repass" class="layui-form-label"></label>
            <input type="submit" class="layui-btn" lay-filter="add" value="添加文章">
        </div>
        </form>
        </div>
        </div>
        <script>layui.use(['form', 'layer'],
            function() {
                $ = layui.jquery;
                var form = layui.form,
                layer = layui.layer;

                //自定义验证规则
                form.verify({
                    nikename: function(value) {
                        if (value.length < 5) {
                            return '昵称至少得5个字符啊';
                        }
                    },
                    pass: [/(.+){6,12}$/, '密码必须6到12位'],
                    repass: function(value) {
                        if ($('#L_pass').val() != $('#L_repass').val()) {
                            return '两次密码不一致';
                        }
                    }
                });

                //监听提交
                form.on('submit(add)',
                function(data) {
                    console.log(data);
                    //发异步，把数据提交给php
                    layer.alert("增加成功", {
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