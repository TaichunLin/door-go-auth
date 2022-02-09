$(document).ready(() => {
  FetchAllUsers();

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
    let username = $("#username").val();
    let cardId = $("#cardId").val();
    let groupId = $("#groupId").val();

    fetch(`http://localhost:1106/api/addUser`, {
      method: "POST",
      headers: headers,
      body: JSON.stringify({
        cardId: `${cardId}`,
        groupId: { groupId: `${groupId}` },
        username: `${username}`,
      }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log("newUser:", data);
      })
      .catch((err) => console.log(err));
    FetchAllUsers();
  });

  $(document).on("click", ".delete-row", () => {
    need2del = $(".checkbox").filter((index, item) => {
      if ($(item).is(":checked")) {
        FetchDelUser(item.id.split("_")[1]);
      }
      return $(item).is(":checked");
    });
    console.log(need2del);
  });
  $(document).on("click", "#CheckAll", (e) => {
    console.log($(e.target).is(":checked"));
    $(".checkbox").prop("checked", $(e.target).is(":checked"));
  });

  $(document).on("click", ".myBtn", (event) => {
    if ($(event.target).hasClass("add")) {
      $("#editlabel").html("新增");
      $(".add-row").val("新增");
      $("#username").val("");
      $("#cardId").val("");

      $("#groupId").val("");
    } else {
      $("#editlabel").html("編輯");
      $(".add-row").val("編輯");
      const [, cardId] = event.target.id.split("_");
      fetch(`http://localhost:1106/api/findUser?cardId=${cardId}`)
        .then((res) => res.json())
        .then((data) => {
          console.log("edit", data);
          $("#username").val(data.username);
          $("#cardId").val(data.cardId);
          $("#groupId").val(data.groupId);
        })
        .catch((err) => console.log(err));
    }
    modal.css("display", "block");
  });
});

let FetchAllUsers = () => {
  return fetch("http://localhost:1106/api/users")
    .then((res) => res.json())
    .then((data) => {
      console.log("allUsers:", data);

      $("#tb").empty();

      $("#tb").append(`<tr>
                        <th><input type="checkbox" name="CheckAll" id="CheckAll"></th>
                        <th>姓名</th>
                        <th>卡號</th>
                        <th>部門名稱</th>
                        <th>部門編號</th>
                        <th>編輯</th>
                        </tr>`);

      $.each(data, function (i, v) {
        // ==     data.forEach(v => {   });
        $("#tb").append(
          $(
            `<tr>
                                  <td><input class="checkbox" type="checkbox" name="record" id="check_${v.cardId}"></td>
                                  <td>${v.username}</td>
                                  <td>${v.cardId}</td>
                                  <td>${v.group}</td>
                                  <td>${v.groupId}</td>
                                  <td><button id="btn_${v.cardId}" class="myBtn">編輯</button></td>
                                  </tr>`
          )
        );
      });
      return data;
    })
    .catch((err) => console.log(err));
};

const csrf_token = document
  .querySelector("meta[http-equiv='csrf-token']")
  .getAttribute("content");

let headers = {
  "Content-Type": "application/json; charset=utf-8",
  Accept: "application/json",
  // Authorization: `Bearer ${token}`,
  "X-CSRF-Token": csrf_token,
};

let FetchDelUser = (cardId) => {
  fetch(`http://localhost:1106/api/deleteUser`, {
    method: "POST",
    headers: headers,
    body: JSON.stringify({
      cardId: `${cardId}`,
    }),
  });
  FetchAllUsers();
};

// var csrf_token = $('meta[http-equiv="csrf-token"]').attr("content");

// function csrfSafeMethod(method) {
// these HTTP methods do not require CSRF protection
//   return /^(GET|HEAD|OPTIONS)$/.test(method);
// }

// $.ajaxSetup({
//   beforeSend: function (xhr, settings) {
//     if (!csrfSafeMethod(settings.type) && !this.crossDomain) {
//       xhr.setRequestHeader("X-CSRF-Token", csrf_token);
//     }
//   },
// });
