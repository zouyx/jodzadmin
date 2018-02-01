{{template "header" .}}
<body>
  {{template "nav" .}}
    <div class="template-page-wrapper">
      <div class="templatemo-content-wrapper">
          <div class="templatemo-content">
            <div class="table-responsive">
                <h4 class="margin-bottom-15">同步器节点管理</h4>
                <table class="table table-striped table-hover table-bordered">
                    <thead>
                    <tr>
                        <th>同步器id</th>
                        <th>同步器名称</th>
                        <th>编辑</th>
                    </tr>
                    </thead>
                    <tbody>
                        {{range $key, $val := .names}}
                            <tr>
                                <td>{{$key}}</td>
                                <td>{{$val}}</td>
                                <td><a href="/management/synchronizer/" class="btn btn-default">Edit</a></td>
                                <td><a href="#" class="btn btn-default">Edit</a></td>
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