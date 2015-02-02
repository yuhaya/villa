<div class="HEADBOX">
    <table class="HEAD">
        <tr>
            <td>
                <ul class="UL">
                    <li data-val="1">已合作</li>
                    <li data-val="0">洽谈中</li>
                    <li data-val="2">解除合作</li>
                    <li data-val="4">已删除</li>
                    <li data-val="3">未回复</li>
                </ul>
            </td>
            <td style="text-align: right">
                <input type="button" id="add_partner" value="ADD" class="add_btn"/>
            </td>
        </tr>
    </table>

</div>

<div id="con">
    {{ if ne .err "" }}
        {{.err}}
    {{ else }}
        <table class="TABLE">
            <tr>
                <th>Paterner Name</th>
                <th>Location</th>
                <th>Manager</th>
                <th>Reservation</th>
                <th>Website</th>
                <th>Commission</th>
                <th>Operation</th>
            </tr>
            <tr>
            {{ if eq .num 0 }}
                <td colspan="7">no data</td>
            {{ else }}

                {{range .partners}}
                    <td>{{.Name}}</td>
                    <td>xxx</td>
                    <td>
                        <p>{{.ManagerContact}}</p>
                        <p>{{.ManagerTelephone}}</p>
                        <p>{{.ManagerEmail}}</p>
                    </td>
                    <td>
                        <p>{{.ReservationContact}}</p>
                        <p>{{.ReservationTelephone}}</p>
                        <p>{{.ReservationEmail}}</p>
                    </td>
                    <td>{{.Website}}</td>
                    <td>{{.Commission}}</td>
                    <td style="width: 150px">
                        <input type="button" value="Edit" class="edit_btn"/>
                        <input type="button" value="Delete" class="del_btn"/>
                    </td>
                {{end}}

            {{ end }}
            </tr>
        </table>
    {{ end }}
</div>
<script type="application/javascript">
var VALS = {
    state:{{.state}}
}
</script>