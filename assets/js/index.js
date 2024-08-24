links = [
  "https://pixqc.com/",
  "https://supaiku.com/",
  "http://xahlee.org/",
  "https://www.gmkonan.dev/",
  "https://bones-ai.bearblog.dev/",
  "https://dancres.github.io/Pages/",
  "https://vlimki.dev/",
  "https://adueck.github.io/blog/",
  "https://world.hey.com/dhh/",
  "https://lelouch.dev/",
  "https://vin.gg/",
];

function addItems() {
  let elem = document.getElementById("from-others-list");

  links.map((e) => {
    let newItem = document.createElement("li");
    let newAnchor = document.createElement("a");

    newAnchor.innerHTML = e;
    newAnchor.setAttribute("href", e);
    newItem.appendChild(newAnchor);
    elem.appendChild(newItem);
  });
}

addItems();

document.getElementById("noise").volume = 0.15;
