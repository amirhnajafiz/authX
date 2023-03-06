// url
const url = 'http://' + window.location.host;

// make http call for register
function register() {
    let stEl = document.getElementById("student-number");
    let psEl = document.getElementById("password");

    let data = {
        "student_number": stEl.value,
        "password": psEl.value
    }

    fetch(url+"/api/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })
        .then((response) => response.text())
        .then((data) => {
            document.getElementById("response").value = data
        })
        .catch((error) => {
            console.error(error);
            alert("Registration Failed");
        })
}

// clear function
function resetForm() {
    document.getElementById("student-number").value = '';
    document.getElementById("password").value = '';
    document.getElementById("response").value = '';
}

// copy token to clipboard
function copy() {
    let text = document.getElementById("response").value

    if (text === "") {
        return
    }

    navigator.clipboard.writeText(text);

    alert("API Key copied to clipboard!");
}