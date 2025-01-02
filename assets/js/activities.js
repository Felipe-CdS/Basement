function findCookie() {
  const cookies = document.cookie.split("; ");

  for (let i = 0; i < cookies.length; i++) {
    let row = cookies[i];

    if (row.startsWith("start=")) {
      return row.split("=")[1];
    }
  }
}

function calcTime() {
  const diff = new Date() - new Date(startTime * 1e3);

  let hrs = Math.floor(diff / 3.6e6);
  let min = Math.floor((diff % 3.6e6) / 6e4);
  let sec = Math.floor((diff % 6e4) / 1e3);

  hrs = hrs.toString().padStart(2, "0");
  min = min.toString().padStart(2, "0");
  sec = sec.toString().padStart(2, "0");

  const display = `${hrs}:${min}:${sec}`;

  if (isNaN(hrs)) {
    display = `00:00:00`;
  }

  document.getElementById("counter-display").innerHTML = display;
}

function startTimerInterval() {
  startTime = findCookie();

  if (startTime == null) {
    document.getElementById("counter-display").innerHTML = "00:00:00";
    clearInterval(intervalHolder);
    return;
  }

  intervalHolder = window.setInterval(calcTime, 1000);
}

// Start when page loads
var startTime = findCookie();
var intervalHolder, blinkHolder;

var display = document.getElementById("counter-display");

startTimerInterval();
