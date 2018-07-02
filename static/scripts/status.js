function setDeviceId(deviceIdVal) {
    var deviceId = document.getElementById('device_id');
    if (deviceIdVal != "UNKNOWN_DEVICE_SN") {
        deviceId.innerHTML = "ID " + deviceIdVal;
    } else {
        deviceId.innerHTML = 'ID Not available';
    }
}

function setSoftwareInfo(data) {
    var domE = document.getElementById('swinfo');
    if (data) {
        txt = "<ul>"
        txt += "<li>" + "date: " + data['date'] + "</li>"
        txt += "<li>" + "version: " + data['version'] + "</li>"
        txt += "<li>" + "name: " + data['name'] + "</li>"
        txt += "<li>" + "kernel: " + data['os'] + " " + data['kernel'] + " " + data["platform"] + "</li>"
        txt += "</ul>"
        domE.innerHTML = txt;
    } else {
        domE.innerHTML = 'Not available';
    }
}

function secondsToStr(seconds) {

    function numberEnding(number) {
        return (number > 1) ? 's' : '';
    }

    var temp = Math.floor(seconds);
    var years = Math.floor(temp / 31536000);
    if (years) {
        return years + ' year' + numberEnding(years);
    }
    var days = Math.floor((temp %= 31536000) / 86400);
    if (days) {
        return days + ' day' + numberEnding(days);
    }
    var hours = Math.floor((temp %= 86400) / 3600);
    if (hours) {
        return hours + ' hour' + numberEnding(hours);
    }
    var minutes = Math.floor((temp %= 3600) / 60);
    if (minutes) {
        return minutes + ' minute' + numberEnding(minutes);
    }
    var seconds = temp % 60;
    if (seconds) {
        return seconds + ' second' + numberEnding(seconds);
    }
    return 'less than a second';
}

function setUptime(uptimeVal) {
    var uptime = document.getElementById('uptime');
    if (uptimeVal) {
        uptime.innerHTML = secondsToStr(uptimeVal);
    } else {
        uptime.innerHTML = 'Not available';
    }
}

function setRegistered(data) {
    var registered = document.getElementById('registered');
    date = new Date(data['date'] * 1000)
    if (data) {
        txt = "<ul>"
        txt += "<li>" + "date " + date.toDateString() + "</li>"
        txt += "<li>" + "status " + data['status'] + "</li>"
        txt += "</ul>"
        registered.innerHTML = txt;
    } else {
        registered.innerHTML = 'Not available';
    }
}

function setOnline(onlineVal) {
    var online = document.getElementById('online');
    if (onlineVal) {
        online.innerHTML = '<div class="device-online">Online</div>';
    } else {
        online.innerHTML = '<div class="device-offline">Offline</div>';
    }
}

function setHardwareInfo(data) {
    var hwinfo = document.getElementById('hwinfo');
    if (data) {
        txt = "<ul>"
        for (x in data) {
            if (x == "memory") {
                mbRam = data[x] / 1048576
                val = mbRam + " Mb"
            } else if (x == "mhz") {
                ghz = data[x] / 1000
                val = ghz + " Ghz"
            } else {
                val = data[x]
            }
            txt += "<li>" + x + ": " + val + "</li>"
        }
        txt += "</ul>"
        hwinfo.innerHTML = txt;
    } else {
        hwinfo.innerHTML = 'Not available';
    }
}


function setNetworkInfo(networks) {
    var hwinfo = document.getElementById('netinfo');
    if (networks) {
        items = "<div>"
        for (network in networks) {
            var network_item = networks[network]
            var addrs = network_item['addrs']
            txt = "<ul>"
            txt += "<li>" + 'mac' + ": " + network_item['mac'] + "</li>"
            txt += "<li>" + 'name' + ": " + network_item['name'] + "</li>"
            for (addr in addrs) {
                txt += "<li>" + 'ip' + ": " + addrs[addr]['IP'] + "</li>"
            }
            txt += "</ul>"
            items += txt
        }
        items += "</div>"
        hwinfo.innerHTML = items;
    } else {
        hwinfo.innerHTML = 'Not available';
    }
}


function load() {
    fetch('/api/status')
        .then(function (response) { return response.json(); })
        .then(function (json) {
            var example = document.getElementById('example');
            example.innerHTML = ""
            var data = json.data;
            for (x in data) {
                if (x == "sw_info") {
                    setSoftwareInfo(data[x])
                } else if (x == "id") {
                    setDeviceId(data[x])
                } else if (x == "uptime") {
                    setUptime(data[x])
                } else if (x == "registered") {
                    setRegistered(data[x])
                } else if (x == "hw_info") {
                    setHardwareInfo(data[x])
                } else if (x == "online") {
                    setOnline(data[x])
                } else if (x == "network_interfaces") {
                    setNetworkInfo(data[x])
                }else {
                    example.innerHTML += "<p>" + x + ": " + data[x] + "</p>"
                }
            }
        });
}

load()

var intervalID = setInterval(function () { load(); }, 5000);