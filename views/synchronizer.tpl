{{template "header" .}}
<body>
  {{template "nav" .}}
    <div class="template-page-wrapper">
      <div class="templatemo-content-wrapper">
          <div class="templatemo-content">
            <div class="table-responsive">
                <h4 class="margin-bottom-15">同步器管理</h4>
                <table class="table table-striped table-hover table-bordered">
                    <thead>
                    <tr>
                        <th>同步器id</th>
                        <th>同步器名称</th>
                        <th>操作</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                        {{range $key, $val := .names}}
                            <tr>
                                <td>{{$key}}</td>
                                <td>{{$val}}</td>
                                <td><a href="/management/synchronizer/config/{{$val}}" class="btn btn-default">配置管理</a></td>
                                <td><a href="/management/synchronizer/nodes/{{$val}}" class="btn btn-default">查看节点</a></td>
                            </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>
          </div>
        </div>
      </div>
    </div>
  {{template "footer" .}}
</body>
</html>