require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");

$(() => {
  if ($("#click-chart").length > 0) {
    let labels = [];
    let data = [];
    for (let i = 0, len = clicks.length; i < len; i++) {
      let click = clicks[i];
      labels.push(click.date);
      data.push(click.count);
    }
    new Chart(document.getElementById("click-chart"), {
      type: "line",
      data: {
        labels,
        datasets: [
          {
            label: "Clicks",
            data,
            fill: false,
            borderColor: "rgb(75, 192, 192)",
            lineTension: 0.1
          }
        ]
      },
      options: {
        scales: {
          yAxes: [
            {
              ticks: {
                beginAtZero: true
              }
            }
          ]
        },
        legend: {
          display: false
        }
      }
    });
  }
});
