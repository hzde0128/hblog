
<!doctype html>
<html lang="zh-CN">
<head>
    <meta name="generator" content="Hugo 0.64.1" />

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>个人博客</title>
    <meta property="og:title" content="黄忠德的博客 | DevOps,运维工程师,运维开发,Python,Golang程序员,开源爱好者">
    <meta property="og:type" content="website">
    <meta name="Keywords" content="黄忠德的博客,微服务,microservice,DevOps,go,golang,go技术文档,go语言教程,linux,kubernetes,k8s,istio,服务网格,运维,hzde0128,黄忠德,jerry">
    <meta name="description" content="黄忠德, Jerry Wong, huangzhongde, , 黄忠德的的博客, 博客, 个人网站, 互联网, Web, 云原生, PaaS, Istio, Kubernetes, 微服务, Microservice">
    <meta property="og:url" content="https://huangzhongde.cn/">
    <link rel="shortcut icon" href="/static/img/favicon.ico" type="image/x-icon">

    <link rel="stylesheet" href='/static/css/normalize.css'>
    <link rel="stylesheet" href='/static/css/style.css'>
    <link rel="alternate" type="application/rss+xml+xml" href="https://huangzhongde.cn/index.xml" title="黄忠德的博客" />
    <script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>

</head>


<body>
<header id="header" class="clearfix">
    <div class="container">
        <div class="col-group">
            <div class="site-name ">

                <h1>
                    <a id="logo" href="https://huangzhongde.cn">
                        黄忠德的博客
                    </a>
                </h1>

                <p class="description">DevOps,运维工程师,运维开发,Python,Golang程序员,开源爱好者</p>
            </div>
            <div>
                <nav id="nav-menu" class="clearfix">
                    <a class="current" href="https://huangzhongde.cn">首页</a>
                    <a  href="https://huangzhongde.cn/archives/" title="归档">归档</a>
                    <a  href="https://huangzhongde.cn/categories/Linux/" title="Linux">Linux</a>
                    <a  href="https://huangzhongde.cn/categories/Golang/" title="Golang">Golang</a>
                </nav>
            </div>
        </div>
    </div>
</header>

<div id="body">
    <div class="container">
        <div class="col-group">
            <div class="col-8" id="main">
                <div class="res-cons">
                    {{ range .article }}
                    <article class="post">
                        <header>
                            <h1 class="post-title">
                                <a href="https://huangzhongde.cn/post/CentOS8_Fastdfs_introduce/" title="CentOS8下FastdDFS安装和使用">CentOS8下FastdDFS安装和使用</a>
                            </h1>
                        </header>
                        <date class="post-meta meta-date">
                            2020年3月16日
                        </date>

                        <div class="post-meta">
                            <span>|</span>

                            <span class="meta-category"><a href='https://huangzhongde.cn/categories/Linux'>Linux</a></span>

                        </div>

                        <div class="post-content">
                            CentOS8下FastdDFS安装 Fastdfs简介 FastDFS 是一个开源的高性能分布式文件系统（DFS）。 它的主要功能包括：文件存储，文件同步和文件访问，以及高容量和负载平衡。主要解决了海量数据存储问题，特别适合以中小文件（建议范围：4KB &lt; file_size &lt;500MB）为载体的在线服务。 FastDFS 系……
                        </div>
                        <p class="readmore"><a href="https://huangzhongde.cn/post/CentOS8_Fastdfs_introduce/">阅读全文</a></p>
                    </article>
                    {{ end }}

                    <ol class="page-navigator">

                        <li  class="current">
                            <a href="https://huangzhongde.cn/">1</a>
                        </li>

                        <li >
                            <a href="https://huangzhongde.cn/page/2/">2</a>
                        </li>

                        <li >
                            <a href="https://huangzhongde.cn/page/3/">3</a>
                        </li>

                        <li >
                            <a href="https://huangzhongde.cn/page/4/">4</a>
                        </li>

                        <li >
                            <a href="https://huangzhongde.cn/page/5/">5</a>
                        </li>

                        <li >
                            <a href="https://huangzhongde.cn/page/6/">6</a>
                        </li>

                        <li >
                            <a href="https://huangzhongde.cn/page/7/">7</a>
                        </li>

                        <li class="next">
                            <a href="https://huangzhongde.cn/page/2/">下一页</a>
                        </li>

                    </ol>

                </div>

            </div>

            <div id="secondary">
                <section class="widget">
                    <form id="search" action='https://huangzhongde.cn/search/' method="get" accept-charset="utf-8" target="_blank" _lpchecked="1">

                        <input type="text" name="q" maxlength="20" placeholder="Search">
                        <input type="hidden" name="sitesearch" value="https://huangzhongde.cn">
                        <button type="submit" class="submit icon-search"></button>
                    </form>
                </section>

                <section class="widget">
                    <h3 class="widget-title">最近文章</h3>
                    <ul class="widget-list">
                        <li>
                            <a href="https://huangzhongde.cn/post/2020-03-16-Microservice_RESTful/" title="微服务系列（6）RESTful介绍">微服务系列（6）RESTful介绍</a>
                        </li>

                        <li>
                            <a href="https://huangzhongde.cn/post/2020-03-15-Microservice_Consul/" title="微服务系列（4）Consul介绍">微服务系列（4）Consul介绍</a>
                        </li>

                        <li>
                            <a href="https://huangzhongde.cn/post/2020-03-15-Microservice_Micro/" title="微服务系列（5）Micro">微服务系列（5）Micro</a>
                        </li>

                        <li>
                            <a href="https://huangzhongde.cn/post/2020-03-14-Microservice_introduce/" title="微服务系列（1）微服务介绍">微服务系列（1）微服务介绍</a>
                        </li>

                        <li>
                            <a href="https://huangzhongde.cn/post/2020-03-14-Microservice_Protobuf/" title="微服务系列（2）Protobuf介绍">微服务系列（2）Protobuf介绍</a>
                        </li>

                        <li>
                            <a href="https://huangzhongde.cn/post/2020-03-14-Microservice_gRPC/" title="微服务系列（3）gRPC介绍">微服务系列（3）gRPC介绍</a>
                        </li>

                    </ul>
                </section>

                <section class="widget">
                    <h3 class="widget-title">分类</h3>
                    <ul class="widget-list">

                        <li><a href="https://huangzhongde.cn/categories/Golang/">Golang (21)</a></li>

                        <li><a href="https://huangzhongde.cn/categories/Kubernetes/">Kubernetes (8)</a></li>

                        <li><a href="https://huangzhongde.cn/categories/Linux/">Linux (16)</a></li>


                    </ul>
                </section>

                <section class="widget">
                    <h3 class="widget-title">标签</h3>
                    <div class="tagcloud">

                        <a href="https://huangzhongde.cn/tags/DevOps/">DevOps</a>


                        <a href="https://huangzhongde.cn/tags/etcd/">etcd</a>

                        <a href="https://huangzhongde.cn/tags/go/">go</a>

                        <a href="https://huangzhongde.cn/tags/golang/">golang</a>

                        <a href="https://huangzhongde.cn/tags/kubernetes/">kubernetes</a>

                        <a href="https://huangzhongde.cn/tags/linux/">linux</a>

                        <a href="https://huangzhongde.cn/tags/microservice/">microservice</a>


                    </div>
                </section>


                <section class="widget">
                    <h3 class="widget-title">友情链接</h3>
                    <ul class="widget-list">

                        <li>
                            <a target="_blank" href="https://huangzhongde.cn/" title="黄忠德的博客">黄忠德的博客</a>
                        </li>

                        <li>
                            <a target="_blank" href="https://cncf.io/" title="CNCF">CNCF</a>
                        </li>

                        <li>
                            <a target="_blank" href="https://kubernetes.io/" title="Kubernetes">Kubernetes</a>
                        </li>

                        <li>
                            <a target="_blank" href="https://itsio.io/" title="Istio">Istio</a>
                        </li>

                    </ul>
                </section>


                <section class="widget">
                    <h3 class="widget-title">其它</h3>
                    <ul class="widget-list">
                        <li><a href="https://huangzhongde.cn/index.xml">文章 RSS</a></li>
                    </ul>
                </section>
            </div>

        </div>
    </div>
</div>
<footer id="footer">

</footer>

</body>

</html>