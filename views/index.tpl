<!doctype html>
<html lang="en">
	<head>
        <title>{{ .basic.Title }}</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

        <!-- Main CSS --> 
        <link rel="stylesheet" href="/static/css/style.css">

        <!-- Font Awesome -->
        <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
    </head>
  
    <body>

        <!-- 头部信息 -->
        <div class="header-wrap d-none d-md-block">
            <div class="container">
                <div class="row">
                    
                    <!-- 左侧头部 -->
                    <header class="col-6 text-left">
                        <h1>{{ .basic.Title }}</h1>
                    </header>
                    
                    <!-- Right header box -->
                    <div class="col-6 text-right">               
                        <p class="header-social-icons social-icons">
                            <a href="#"><i class="fa fa-facebook fa-2x"></i></a>
                            <a href="#"><i class="fa fa-twitter fa-2x"></i></a>
                            <a href="#"><i class="fa fa-youtube fa-2x"></i></a>
                            <a href="#"><i class="fa fa-instagram fa-2x"></i></a>
                        </p>
                    </div>
                </div>
            </div>
        </div>

        
        <!-- 主导航 -->
        <nav class="navbar navbar-expand-md navbar-dark bg-primary">
            <div class="container">
                <!-- 导航条目 -->
                <div class="collapse navbar-collapse" id="mainNavbar">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item active">
                                <a class="nav-link" href="/">主页<span class="sr-only">(current)</span></a>
                        </li>

                        <li class="nav-item">
                                <a class="nav-link" href="#">分类</a>
                        </li>

                        <li class="nav-item">
                                <a class="nav-link" href="#">关于</a>
                        </li>
                    </ul>
                    <ul class="navbar-nav ml-auto">
                        <li class="nav-item">
                            <a class="nav-link" href="/login">登录后台</a></li>
                    </ul>
                </div>
                
            </div>
        </nav>
		<div class="tlinks">Collect from 网站模板</div>

        <!-- 主体内容区域 -->
        <main class="container">
            <div class="row">
                
                <!-- 文章列表 -->
                <div class="col-sm-8">
                    {{ range .article }}
                    <article>
                        <h2 class="article-title">{{ .Title }}</h2>

                        <p class="article-meta">发布于<time datetime="2020-03-18">{{ .CreateAt.Format "2006-01-02" }}</time></p>

                        <p>{{ .Content }}</p>

                        <a href="#" class="btn btn-primary">阅读全文</a>
                        <a href="#" class="btn btn-secondary">发表评论</a>

                    </article>
                    {{ end }}
                    <!-- Example pagination Bootstrap component -->
                    <nav>
                        <ul class="pagination">
                                <li class="page-item">
                                    <a class="page-link" href="#" aria-label="Previous">
                                        <span aria-hidden="true">&laquo;</span>
                                        <span class="sr-only">Previous</span>
                                    </a>
                                </li>
                                <li class="page-item"><a class="page-link" href="#">1</a></li>
                                <li class="page-item active"><a class="page-link" href="#">2</a></li>
                                <li class="page-item"><a class="page-link" href="#">3</a></li>
                                <li class="page-item">
                                    <a class="page-link" href="#" aria-label="Next">
                                        <span aria-hidden="true">&raquo;</span>
                                        <span class="sr-only">Next</span>
                                    </a>
                                </li>
                        </ul>
                    </nav>
                </div>

                <!-- 侧边栏 -->
                <aside class="col-sm-4">

                    <div class="sidebar-box">
                        <h4>站内搜索</h4>
                        <form class="form-inline my-2 my-lg-0">
                            <input class="form-control mr-sm-2" type="text"  placeholder="Search" aria-label="Search">
                            <button class="btn my-2 my-sm-0" type="submit">搜索</button>
                        </form>
                    </div>

                    <div class="sidebar-box">
                        <h4>分类</h4>
                        <div class="list-group list-group-root">
                            <a class="list-group-item active" href="/">主页</a>
                            <a class="list-group-item" href="/category">分类</a>
                            <a class="list-group-item" href="#">Sed</a>
                            <a class="list-group-item" href="#">关于</a>
                        </div>
                    </div>

                    <div class="sidebar-box sidebar-box-bg">
                        <h4>关于</h4>
                        <p>Aenean nec massa a tortor auctor sodales sed a dolor. Duis vitae lorem sem. Proin at velit vel arcu pretium luctus. <a href="#" class="readmore">Read More &raquo;</a></p>    
                    </div>

                    <div class="sidebar-box">
                        <h4>友情链接</h4>
                        <ul>
                            <li><a href="https://huangzhongde.cn" title="myblog">黄忠德的博客</a></li>
                        </ul>
                    </div>
                </aside> 
                
            </div>
        </main>

        <!-- 底部 -->
        <footer class="footer">
            <div class="footer-lists">
                <div class="container">
                </div>
            </div>
            
            
            <div class="footer-bottom">
                <p class="text-center">@2020 版权所有 <a href="{{ .basic.BaseUrl }}">{{ .basic.Title }}</a></p>
                    <p class="text-center"><a href="#">返回顶部</a></p>
            </div>
            
        </footer>

        <!-- Bootcamp JavaScript -->
        <!-- jQuery first, then Popper.js, then Bootstrap JS -->
        <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.3/umd/popper.min.js" integrity="sha384-vFJXuSJphROIrBnz7yo7oB41mKfc8JzQZiCq4NCceLEaO4IHwicKwpJf9c9IpFgh" crossorigin="anonymous"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/js/bootstrap.min.js" integrity="sha384-alpBpkh1PFOepccYVYDB4do5UnbKysX5WZXm3XxPqe5iKTfUKjNkCk9SaVuEZflJ" crossorigin="anonymous"></script>

    </body>
</html>