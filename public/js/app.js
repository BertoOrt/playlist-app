$(document).ready(function () {
  $('#searchForm').submit(function (e) {
    var apiURL = "https://itunes.apple.com/search?term=jack+johnson&limit=25";
    $.getJSON(apiURL,function (data) {
      console.log(data);
    })
    e.preventDefault()
  })
})
