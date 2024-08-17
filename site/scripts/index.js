const response = fetch("http://worldtimeapi.org/api/timezone/America/New_York")
    .then(data => data.json().then(
      data => {
        document.querySelector("h1").innerHTML = data.datetime;
      }
    ));