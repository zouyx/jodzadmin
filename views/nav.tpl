{{define "nav"}}
<div class="navbar navbar-inverse" role="navigation">
    <div class="navbar-header">
        <div class="logo"><h1>jodz</h1></div>
    </div>
</div>
<div class="navbar-collapse collapse templatemo-sidebar">
    <ul class="templatemo-sidebar-menu">
        <li>
            <form class="navbar-form">
                <input type="text" class="form-control" id="templatemo_search_box" placeholder="Search...">
                <span class="btn btn-default">Go</span>
            </form>
        </li>
        <li><a href="#"><i class="fa fa-home"></i>主页</a></li>
        <li class="sub">
            <a href="javascript:;">
                <i class="fa fa-cubes"></i>服务管理<div class="pull-right"><span class="caret"></span></div>
            </a>
            <ul class="templatemo-submenu">
                <li><a href="#">同步器管理</a></li>
            </ul>
        </li>
    </ul>
</div><!--/.navbar-collapse -->
{{end}}