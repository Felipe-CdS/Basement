function findCookie() {
  const cookies = document.cookie.split("; ");

  for (let i = 0; i < cookies.length; i++) {
    let row = cookies[i];

    if (row.startsWith("start=")) {
      startTime = row.split("=")[1];
      break;
    }
  }
}

function calcTime() {
  if (startTime == null) {
    document.getElementById("counter-display").style.display = "none";
    return;
  }

  const diff = new Date() - new Date(startTime * 1e3);

  let hrs = Math.floor(diff / 3.6e6);
  let min = Math.floor((diff % 3.6e6) / 6e4);
  let sec = Math.floor((diff % 6e4) / 1e3);

  hrs = hrs.toString().padStart(2, "0");
  min = min.toString().padStart(2, "0");
  sec = sec.toString().padStart(2, "0");

  const display = `${hrs}:${min}:${sec}`;

  document.getElementById("counter-display").style.display = "initial";
  document.getElementById("counter-display").innerHTML = display;
}

// Start when page loads
var startTime = null;
findCookie();
var intervalCalc = window.setInterval(calcTime, 1000);
