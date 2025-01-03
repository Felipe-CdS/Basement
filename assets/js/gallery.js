htmx.on("#upload-form", "htmx:xhr:progress", function (evt) {
  htmx
    .find("#progress")
    .setAttribute("value", (evt.detail.loaded / evt.detail.total) * 100);
});

// force close dialog when page changes
window.addEventListener(
  "popstate",
  function (event) {
    document.getElementById("uploadDialog").close();
  },
  false,
);

function drop(ev) {
  let files = ev.dataTransfer.files;
  let recipientArea = document.getElementById("recipient");

  recipientArea.innerHTML = "";

  for (const i of files) {
    let p = document.createElement("p");
    p.innerHTML = i.name;
    recipientArea.appendChild(p);
  }

  document.getElementById("files-input").files = files;
}

htmx.config.allowNestedOobSwaps = false;
