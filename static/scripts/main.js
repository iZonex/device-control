// fetch('/wifi')
//     .then(function (response) { return response.json(); })
//     .then(function (json) {
//         var connectionStatus = document.getElementById('connection');
//         if (json.data.connected === true) {
//             connectionStatus.innerHTML = '<strong class="connection-status">Connected</strong>';
//         } else {
//             connectionStatus.innerHTML = '<strong class="connection-status">Disconnected</strong>';
//         }
//     });

function sendForm() {
    location.href = "/server";
};