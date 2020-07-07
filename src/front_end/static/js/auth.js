var auth = localStorage.getItem("Authorization");
if (auth == null ) {
    window.location.href = "/index";
}

function isTokenExpired(data) {
    if (data.code== 200412) {
        localStorage.removeItem("Authorization");
        window.location.href = "/index";
    }
}
