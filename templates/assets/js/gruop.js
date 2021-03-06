$(document).ready(() => {
  // const token = xxxxx;
  const csrf_token = document
    .querySelector("meta[http-equiv='csrf-token']")
    .getAttribute("content");

  let headers = {
    "Content-Type": "application/json",
    Accept: "application/json",
    // Authorization: `Bearer ${token}`,
    "X-CSRF-Token": csrf_token,
  };

  FetchGetAll();

  let modal = $("#myModal");

  let span = $(".close")[0];
  let [AddRow] = $(".add-row");

  span.onclick = () => {
    modal.css("display", "none");
  };
  AddRow.onclick = () => {
    modal.css("display", "none");
  };

  $(".add-row").click(() => {
    let door = $("#door").val();
    let group = $("#group").val();
    let groupId = $("#groupId").val();

    fetch(`http://localhost:1106/api/addGroup`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        "X-CSRF-Token": $('meta[http-equiv="csrf-token"]').attr("content"),
        "Content-Type": "application/json; charset=utf-8",
      },
      body: JSON.stringify({
        group: `${group}`,
        groupId: `${groupId}`,
        door: `${door}`,
      }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log("newGroup:", data);
        FetchGetAll();
      })
      .catch((err) => console.log(err));
  });

  $(document).on("click", ".delete-row", () => {
    need2del = $(".checkbox").filter((index, item) => {
      if ($(item).is(":checked")) {
        console.log("???", item.id.split("_")[1]);
        FetchDel(item.id.split("_")[1]);
      }
      return $(item).is(":checked");
    });
    console.log("need2del: ", need2del);
  });
  $(document).on("click", "#CheckAll", (e) => {
    console.log($(e.target).is(":checked"));
    $(".checkbox").prop("checked", $(e.target).is(":checked"));
  });

  $(document).on("click", ".myBtn", (event) => {
    if ($(event.target).hasClass("add")) {
      $("#editlabel").html("新增");
      $(".add-row").val("新增");
      $("#door").val("");
      $("#group").val("");
      $("#groupId").val("");
    } else {
      $("#editlabel").html("編輯");
      $(".add-row").val("編輯");
      const [, groupId] = event.target.id.split("_");
      fetch(`http://localhost:1106/api/findGroup?groupId=${groupId}`)
        .then((res) => res.json())
        .then((data) => {
          console.log("edit", data);
          $("#door").val(data.door);
          $("#group").val(data.group);
          $("#groupId").val(data.groupId);
        })
        .catch((err) => console.log(err));
    }
    modal.css("display", "block");
  });
});

let FetchGetAll = () => {
  return fetch("/api/groups")
    .then((res) => res.json())
    .then((data) => {
      console.log("allGroups:", data);

      $("#tb").empty();

      $("#tb").append(`<tr>
                        <th><input type="checkbox" name="CheckAll" id="CheckAll"></th>
                        <th>部門名稱</th>
                        <th>部門編號</th>
                        <th>開門權限</th>
                        <th>編輯</th>
                        </tr>`);

      $.each(data, function (i, v) {
        // ==     data.forEach(v => {   });
        $("#tb").append(
          $(
            `<tr>
                                  <td><input class="checkbox" type="checkbox" name="record" id="check_${v.groupId}"></td>
                                  <td>${v.group}</td>
                                  <td>${v.groupId}</td>
                                  <td>${v.door}</td>
                                  <td><button id="btn_${v.groupId}" class="myBtn">編輯</button></td>
                                  </tr>`
          )
        );
      });
      return data;
    })
    .catch((err) => console.log(err));
};

let FetchDel = (groupId) => {
  fetch(`http://localhost:1106/api/deleteGroup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      "X-CSRF-Token": $('meta[http-equiv="csrf-token"]').attr("content"),
    },
    body: JSON.stringify({
      groupId: `${groupId}`,
    }),
  });
  FetchGetAll();
};
